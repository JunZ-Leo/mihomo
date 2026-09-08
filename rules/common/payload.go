package common

import "strings"

type rulePayload struct {
	ruleType string
	payload  string
	target   string
	params   []string
}

func hasCommaPayload(ruleType string) bool {
	switch ruleType {
	case "NOT", "OR", "AND", "SUB-RULE", "DOMAIN-REGEX", "PROCESS-NAME-REGEX", "PROCESS-PATH-REGEX":
		return true
	default:
		return false
	}
}

func splitRuleFields(ruleRaw string) []string {
	fields := strings.Split(ruleRaw, ",")
	for i := range fields {
		fields[i] = strings.Trim(fields[i], " ")
	}
	return fields
}

func parseRulePayload(ruleRaw string, needTarget bool) rulePayload {
	fields := splitRuleFields(ruleRaw)
	parsed := rulePayload{ruleType: strings.ToUpper(fields[0])}
	if len(fields) == 1 {
		return parsed
	}

	if parsed.ruleType == "MATCH" {
		parsed.target = fields[1]
		return parsed
	}

	if hasCommaPayload(parsed.ruleType) {
		if needTarget {
			parsed.target = fields[len(fields)-1]
			fields = fields[:len(fields)-1]
		}
		parsed.payload = strings.Join(fields[1:], ",")
		return parsed
	}

	parsed.payload = fields[1]
	if len(fields) <= 2 {
		return parsed
	}
	if needTarget {
		parsed.target = fields[2]
		if len(fields) > 3 {
			parsed.params = fields[3:]
		}
	} else {
		parsed.params = fields[2:]
	}
	return parsed
}

// ParseRulePayload parses `type,payload,target(,params...)` or
// `type,payload(,params...)` while preserving the established rule syntax.
func ParseRulePayload(ruleRaw string, needTarget bool) (ruleType, payload, target string, params []string) {
	parsed := parseRulePayload(ruleRaw, needTarget)
	return parsed.ruleType, parsed.payload, parsed.target, parsed.params
}
