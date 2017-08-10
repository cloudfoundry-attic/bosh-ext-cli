package cmd

import (
	"fmt"
	"net/http"

	boshui "github.com/cloudfoundry/bosh-cli/ui"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"

	"github.com/cppforlife/bosh-lint/web"
)

type WebCmd struct {
	cmdRunner boshsys.CmdRunner
	ui        boshui.UI

	logTag string
	logger boshlog.Logger

	allowedCmds map[string][]apiOpt
}

type apiOpt struct {
	Name         string
	WithoutValue bool
	Positional   bool
	EqualsSign   bool
}

func NewWebCmd(cmdRunner boshsys.CmdRunner, ui boshui.UI, logger boshlog.Logger) WebCmd {
	return WebCmd{
		cmdRunner: cmdRunner,
		ui:        ui,

		logTag: "WebCmd",
		logger: logger,

		allowedCmds: map[string][]apiOpt{
			"deployments": []apiOpt{},
			"instances": []apiOpt{
				apiOpt{Name: "deployment"},
				apiOpt{Name: "ps", WithoutValue: true},
				apiOpt{Name: "details", WithoutValue: true},
			},
			"tasks": []apiOpt{
				apiOpt{Name: "recent", EqualsSign: true},
				apiOpt{Name: "all", WithoutValue: true},
			},
			"task": []apiOpt{
				apiOpt{Name: "id", Positional: true},
			},
			"events": []apiOpt{
				apiOpt{Name: "action"},
				apiOpt{Name: "deployment"},
				apiOpt{Name: "instance"},
				apiOpt{Name: "object-name"},
				apiOpt{Name: "object-type"},
				apiOpt{Name: "task"},
				apiOpt{Name: "event-user"},
				apiOpt{Name: "before"},
				apiOpt{Name: "after"},
				apiOpt{Name: "before-id"},
			},
		},
	}
}

func (c WebCmd) Run(opts WebOpts) error {
	http.HandleFunc("/", c.serveUI)
	http.HandleFunc("/api/command", c.serveAPICommand)

	addr := fmt.Sprintf("%s:%d", opts.ListenAddr, opts.ListenPort)
	c.ui.PrintLinef("Starting server on http://%s", addr)

	return http.ListenAndServe(addr, nil)
}

func (c WebCmd) serveUI(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving UI")
	fmt.Fprintf(w, web.Layout)
}

func (c WebCmd) serveAPICommand(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving API command")

	r.ParseForm()

	c.logger.Debug(c.logTag, "Form submitted: %#v", r.Form)

	cmdName := r.Form.Get("command")
	if len(cmdName) == 0 {
		c.logger.Error(c.logTag, "Empty command")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	apiOpts, found := c.allowedCmds[cmdName]
	if !found {
		c.logger.Error(c.logTag, "Disallowed cmd '%s'", cmdName)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmd := boshsys.Command{
		Name: "bosh",
		Args: []string{cmdName, "--json"},
	}

	for _, opt := range apiOpts {
		if len(r.Form[opt.Name]) == 0 {
			continue
		}

		val := r.Form[opt.Name][0]

		if opt.WithoutValue {
			if len(val) > 0 {
				c.logger.Error(c.logTag, "Expected opt '%s' to not have value", opt.Name)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			cmd.Args = append(cmd.Args, "--"+opt.Name)
		} else {
			if len(val) == 0 {
				continue
			}

			if opt.Positional {
				cmd.Args = append(cmd.Args, val)
			} else if opt.EqualsSign {
				cmd.Args = append(cmd.Args, "--"+opt.Name+"="+val)
			} else {
				cmd.Args = append(cmd.Args, "--"+opt.Name, val)
			}
		}
	}

	stdout, _, _, err := c.cmdRunner.RunComplexCommand(cmd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	_, err = w.Write([]byte(stdout))
	if err != nil {
		c.logger.Error(c.logTag, "Failed to write API events response")
	}
}
