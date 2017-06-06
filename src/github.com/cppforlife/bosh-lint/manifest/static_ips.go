package manifest

import (
	"fmt"

	check "github.com/cppforlife/bosh-lint/check"
)

type StaticIPs struct {
	context   check.Context
	netAssocs []NetworkAssociation
	check.Config
}

func NewStaticIPs(context check.Context, netAssocs []NetworkAssociation, config check.Config) StaticIPs {
	return StaticIPs{context, netAssocs, config}
}

func (c StaticIPs) Description() check.Description {
	return check.Description{
		Context_:   c.context,
		Purpose_:   "if static IPs are present",
		Reasoning_: "It's recommended to not use static IPs (unless it's a VIP network).",
	}
}

func (c StaticIPs) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	switch {
	case len(c.netAssocs) == 1:
		netAssoc := c.netAssocs[0]

		// We currently do not allow machines without non-VIP networks
		if len(netAssoc.StaticIPs) > 0 {
			sugs = append(sugs, check.Simple{
				Context_:    c.context.Nested(fmt.Sprintf("Network '%s'", netAssoc.Name)),
				Problem_:    "Static IPs are specified",
				Resolution_: "Remove static IPs",
			})
		}

	case len(c.netAssocs) > 1:
		for _, netAssoc := range c.netAssocs {
			if len(netAssoc.StaticIPs) > 0 && len(netAssoc.Default) > 0 {
				sugs = append(sugs, check.Simple{
					Context_:    c.context.Nested(fmt.Sprintf("Network '%s'", netAssoc.Name)),
					Problem_:    "Static IPs are specified",
					Resolution_: "Remove static IPs",
				})
			}
		}
	}

	return sugs, nil
}
