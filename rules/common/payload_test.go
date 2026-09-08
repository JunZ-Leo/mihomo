package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRulePayloadCompatibility(t *testing.T) {
	tests := []struct {
		name       string
		raw        string
		needTarget bool
		ruleType   string
		payload    string
		target     string
		params     []string
	}{
		{
			name:       "domain with target and params",
			raw:        " DOMAIN-SUFFIX, example.com , Proxy , no-resolve ",
			needTarget: true,
			ruleType:   "DOMAIN-SUFFIX",
			payload:    "example.com",
			target:     "Proxy",
			params:     []string{"no-resolve"},
		},
		{
			name:       "domain without target",
			raw:        "DOMAIN,example.com,no-resolve",
			needTarget: false,
			ruleType:   "DOMAIN",
			payload:    "example.com",
			params:     []string{"no-resolve"},
		},
		{
			name:       "match",
			raw:        "MATCH,DIRECT",
			needTarget: true,
			ruleType:   "MATCH",
			target:     "DIRECT",
		},
		{
			name:       "logic rule keeps commas in payload",
			raw:        "AND,((DOMAIN,example.com),(NETWORK,TCP)),Proxy",
			needTarget: true,
			ruleType:   "AND",
			payload:    "((DOMAIN,example.com),(NETWORK,TCP))",
			target:     "Proxy",
		},
		{
			name:       "regex keeps commas in payload",
			raw:        `DOMAIN-REGEX,^example\\.(com|net),Proxy`,
			needTarget: true,
			ruleType:   "DOMAIN-REGEX",
			payload:    `^example\\.(com|net)`,
			target:     "Proxy",
		},
		{
			name:     "type only",
			raw:      "DOMAIN",
			ruleType: "DOMAIN",
		},
		{
			name: "empty rule",
			raw:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ruleType, payload, target, params := ParseRulePayload(tt.raw, tt.needTarget)
			assert.Equal(t, tt.ruleType, ruleType)
			assert.Equal(t, tt.payload, payload)
			assert.Equal(t, tt.target, target)
			assert.Equal(t, tt.params, params)
		})
	}
}

func TestParseParamsCompatibility(t *testing.T) {
	tests := []struct {
		name      string
		params    []string
		isSrc     bool
		noResolve bool
	}{
		{name: "empty"},
		{name: "no resolve", params: []string{"no-resolve"}, noResolve: true},
		{name: "source implies no resolve", params: []string{"src"}, isSrc: true, noResolve: true},
		{name: "unknown ignored", params: []string{"unknown"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isSrc, noResolve := ParseParams(tt.params)
			assert.Equal(t, tt.isSrc, isSrc)
			assert.Equal(t, tt.noResolve, noResolve)
		})
	}
}
