package cmd

import (
	"fmt"
	"net/http"

	boshui "github.com/cloudfoundry/bosh-cli/ui"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

type VisualizeEventsCmd struct {
	cmdRunner boshsys.CmdRunner
	ui        boshui.UI

	logTag string
	logger boshlog.Logger
}

func NewVisualizeEventsCmd(cmdRunner boshsys.CmdRunner, ui boshui.UI, logger boshlog.Logger) VisualizeEventsCmd {
	return VisualizeEventsCmd{cmdRunner, ui, "VisualizeEventsCmd", logger}
}

func (c VisualizeEventsCmd) Run() error {
	http.HandleFunc("/", c.serveUI)
	http.HandleFunc("/api/events", c.serveAPIEvents)
	http.HandleFunc("/api/tasks", c.serveAPITasks)
	http.HandleFunc("/api/task", c.serveAPITask)

	c.ui.PrintLinef("Starting server on http://localhost:9090")

	return http.ListenAndServe(":9090", nil)
}

func (c VisualizeEventsCmd) serveUI(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving UI")
	fmt.Fprintf(w, eventsUI)
}

func (c VisualizeEventsCmd) serveAPIEvents(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving API events")

	r.ParseForm()

	c.logger.Debug(c.logTag, "Form submitted: %#v", r.Form)

	cmd := boshsys.Command{
		Name: "bosh",
		Args: []string{"events", "--json"},
	}

	allowedKeys := []string{"action", "deployment", "instance",
		"object-name", "object-type", "task", "event-user", "before", "after", "before-id"}

	for _, key := range allowedKeys {
		if len(r.Form[key]) > 0 && len(r.Form[key][0]) > 0 {
			cmd.Args = append(cmd.Args, "--"+key, r.Form[key][0])
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

func (c VisualizeEventsCmd) serveAPITasks(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving API tasks")

	r.ParseForm()

	c.logger.Debug(c.logTag, "Form submitted: %#v", r.Form)

	cmd := boshsys.Command{
		Name: "bosh",
		Args: []string{"tasks", "--json"},
	}

	allowedKeys := []string{"recent", "all"}

	for _, key := range allowedKeys {
		if len(r.Form[key]) > 0 {
			if len(r.Form[key][0]) > 0 {
				cmd.Args = append(cmd.Args, "--"+key+"="+r.Form[key][0])
			} else {
				cmd.Args = append(cmd.Args, "--"+key)
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
		c.logger.Error(c.logTag, "Failed to write API tasks response")
	}
}

func (c VisualizeEventsCmd) serveAPITask(w http.ResponseWriter, r *http.Request) {
	c.logger.Debug(c.logTag, "Serving API task")

	c.logger.Debug(c.logTag, "Form submitted: %#v", r.Form)

	cmd := boshsys.Command{
		Name: "bosh",
		Args: []string{"task", "--json", r.URL.Query().Get("id")},
	}

	stdout, _, _, err := c.cmdRunner.RunComplexCommand(cmd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	_, err = w.Write([]byte(stdout))
	if err != nil {
		c.logger.Error(c.logTag, "Failed to write API task response")
	}
}
