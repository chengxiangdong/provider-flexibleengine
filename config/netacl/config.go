// Package netacl contains the configuration for the netacl package.
package netacl

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/FrangipaneTeam/provider-flexibleengine/pkg/tools"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// flexibleengine_fw_firewall_group_v2
	p.AddResourceConfigurator("flexibleengine_fw_firewall_group_v2", func(r *config.Resource) {
		r.References["ingress_firewall_policy_id"] = config.Reference{
			Type: "flexibleengine_fw_policy_v2",
		}
	})

	// flexibleengine_fw_policy_v2
	p.AddResourceConfigurator("flexibleengine_fw_policy_v2", func(r *config.Resource) {
		r.References["firewall_rules"] = config.Reference{
			Type: tools.GenerateType("netacl", "Policy"),
		}
	})

	// flexibleengine_fw_rule_v2
	p.AddResourceConfigurator("flexibleengine_fw_rule_v2", func(r *config.Resource) {
	})

	// flexibleengine_network_acl
	p.AddResourceConfigurator("flexibleengine_network_acl", func(r *config.Resource) {
		r.References["inbound_rules"] = config.Reference{
			Type:              tools.GenerateType("netacl", "ACLRule"),
			SelectorFieldName: "InboundRuleSelector",
			RefFieldName:      "InboundRuleRefs",
		}
		r.References["outbound_rules"] = config.Reference{
			Type:              tools.GenerateType("netacl", "ACLRule"),
			SelectorFieldName: "OutboundRuleSelector",
			RefFieldName:      "OutboundRuleRefs",
		}
		r.References["subnets"] = config.Reference{
			Type:              tools.GenerateType("vpc", "VPCSubnet"),
			Extractor:         tools.GenerateExtractor(true, "id"),
			SelectorFieldName: "SubnetSelector",
			RefFieldName:      "SubnetRefs",
		}
	})
}
