package rules

import (
	"net/netip"
	"testing"

	C "github.com/metacubex/mihomo/constant"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseRuleCompatibility(t *testing.T) {
	tests := []struct {
		name     string
		ruleType string
		payload  string
		target   string
		params   []string
		wantType C.RuleType
		wantData string
	}{
		{
			name:     "domain normalized",
			ruleType: "DOMAIN",
			payload:  "Example.COM",
			target:   "Proxy",
			wantType: C.Domain,
			wantData: "example.com",
		},
		{
			name:     "match without payload",
			ruleType: "MATCH",
			target:   "DIRECT",
			wantType: C.MATCH,
		},
		{
			name:     "destination cidr",
			ruleType: "IP-CIDR",
			payload:  "10.0.0.0/8",
			target:   "Proxy",
			params:   []string{"no-resolve"},
			wantType: C.IPCIDR,
			wantData: "10.0.0.0/8",
		},
		{
			name:     "source cidr alias",
			ruleType: "SRC-IP-CIDR",
			payload:  "192.168.0.0/16",
			target:   "DIRECT",
			wantType: C.SrcIPCIDR,
			wantData: "192.168.0.0/16",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule, err := ParseRule(tt.ruleType, tt.payload, tt.target, tt.params, nil)
			require.NoError(t, err)
			assert.Equal(t, tt.wantType, rule.RuleType())
			assert.Equal(t, tt.target, rule.Adapter())
			assert.Equal(t, tt.wantData, rule.Payload())
		})
	}
}

func TestParseRuleSourceParamCompatibility(t *testing.T) {
	rule, err := ParseRule("IP-CIDR", "10.0.0.0/8", "Proxy", []string{"src"}, nil)
	require.NoError(t, err)
	assert.Equal(t, C.SrcIPCIDR, rule.RuleType())

	resolveCalled := false
	matched, adapter := rule.Match(&C.Metadata{
		SrcIP: netip.MustParseAddr("10.1.2.3"),
		DstIP: netip.MustParseAddr("203.0.113.1"),
	}, C.RuleMatchHelper{
		ResolveIP: func() {
			resolveCalled = true
		},
	})
	assert.True(t, matched)
	assert.Equal(t, "Proxy", adapter)
	assert.False(t, resolveCalled)
}

func TestParseRuleErrorsCompatibility(t *testing.T) {
	_, err := ParseRule("DOMAIN", "", "Proxy", nil, nil)
	require.EqualError(t, err, "missing subsequent parameters: DOMAIN")

	_, err = ParseRule("UNKNOWN", "value", "Proxy", nil, nil)
	require.EqualError(t, err, "unsupported rule type: UNKNOWN")

	_, err = ParseRule("IP-CIDR", "invalid", "Proxy", nil, nil)
	require.Error(t, err)
}
