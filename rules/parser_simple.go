package rules

import (
	C "github.com/metacubex/mihomo/constant"
	RC "github.com/metacubex/mihomo/rules/common"
)

func parseSimpleRule(ruleType, payload, target string) (C.Rule, bool, error) {
	var (
		rule C.Rule
		err  error
	)

	switch ruleType {
	case "DOMAIN":
		rule = RC.NewDomain(payload, target)
	case "DOMAIN-SUFFIX":
		rule = RC.NewDomainSuffix(payload, target)
	case "DOMAIN-KEYWORD":
		rule = RC.NewDomainKeyword(payload, target)
	case "DOMAIN-REGEX":
		rule, err = RC.NewDomainRegex(payload, target)
	case "DOMAIN-WILDCARD":
		rule, err = RC.NewDomainWildcard(payload, target)
	case "SRC-PORT":
		rule, err = RC.NewPort(payload, target, C.SrcPort)
	case "DST-PORT":
		rule, err = RC.NewPort(payload, target, C.DstPort)
	case "IN-PORT":
		rule, err = RC.NewPort(payload, target, C.InPort)
	case "DSCP":
		rule, err = RC.NewDSCP(payload, target)
	case "PROCESS-NAME":
		rule, err = RC.NewProcess(payload, target, C.ProcessName)
	case "PROCESS-PATH":
		rule, err = RC.NewProcess(payload, target, C.ProcessPath)
	case "PROCESS-NAME-REGEX":
		rule, err = RC.NewProcess(payload, target, C.ProcessNameRegex)
	case "PROCESS-PATH-REGEX":
		rule, err = RC.NewProcess(payload, target, C.ProcessPathRegex)
	case "PROCESS-NAME-WILDCARD":
		rule, err = RC.NewProcess(payload, target, C.ProcessNameWildcard)
	case "PROCESS-PATH-WILDCARD":
		rule, err = RC.NewProcess(payload, target, C.ProcessPathWildcard)
	case "NETWORK":
		rule, err = RC.NewNetworkType(payload, target)
	case "UID":
		rule, err = RC.NewUid(payload, target)
	case "IN-TYPE":
		rule, err = RC.NewInType(payload, target)
	case "IN-USER":
		rule, err = RC.NewInUser(payload, target)
	case "IN-NAME":
		rule, err = RC.NewInName(payload, target)
	case "REMATCH-NAME":
		rule, err = RC.NewRematchName(payload, target)
	case "MATCH":
		rule = RC.NewMatch(target)
	default:
		return nil, false, nil
	}

	return rule, true, err
}
