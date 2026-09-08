package config

import (
	"net/netip"
	"testing"

	C "github.com/metacubex/mihomo/constant"
	T "github.com/metacubex/mihomo/tunnel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigCompatibilityDefaults(t *testing.T) {
	rawCfg, err := UnmarshalRawConfig([]byte("{}"))
	require.NoError(t, err)

	assert.False(t, rawCfg.AllowLan)
	assert.Equal(t, "*", rawCfg.BindAddress)
	assert.Equal(t, []netip.Prefix{
		netip.MustParsePrefix("0.0.0.0/0"),
		netip.MustParsePrefix("::/0"),
	}, rawCfg.LanAllowedIPs)
	assert.True(t, rawCfg.IPv6)
	assert.Equal(t, T.Rule, rawCfg.Mode)
	assert.Equal(t, 24, rawCfg.GeoUpdateInterval)
	assert.Equal(t, "memconservative", rawCfg.GeodataLoader)
	assert.True(t, rawCfg.ETagSupport)

	assert.False(t, rawCfg.DNS.Enable)
	assert.False(t, rawCfg.DNS.IPv6)
	assert.True(t, rawCfg.DNS.UseHosts)
	assert.True(t, rawCfg.DNS.UseSystemHosts)
	assert.Equal(t, uint(100), rawCfg.DNS.IPv6Timeout)
	assert.Equal(t, C.DNSMapping, rawCfg.DNS.EnhancedMode)
	assert.Equal(t, "198.18.0.1/16", rawCfg.DNS.FakeIPRange)
	assert.Equal(t, C.FilterBlackList, rawCfg.DNS.FakeIPFilterMode)

	assert.False(t, rawCfg.Tun.Enable)
	assert.Equal(t, C.TunGvisor, rawCfg.Tun.Stack)
	assert.True(t, rawCfg.Tun.AutoRoute)
	assert.True(t, rawCfg.Tun.AutoDetectInterface)
	assert.Equal(t, []string{"0.0.0.0:53"}, rawCfg.Tun.DNSHijack)

	assert.True(t, rawCfg.Profile.StoreSelected)
	assert.Equal(t, []string{"*"}, rawCfg.ExternalControllerCors.AllowOrigins)
	assert.True(t, rawCfg.ExternalControllerCors.AllowPrivateNetwork)
}

func TestConfigCompatibilityPartialOverridesKeepDefaults(t *testing.T) {
	rawCfg, err := UnmarshalRawConfig([]byte(`
allow-lan: true
dns:
  enable: true
  nameserver:
    - 1.1.1.1
tun:
  enable: true
`))
	require.NoError(t, err)

	assert.True(t, rawCfg.AllowLan)
	assert.Equal(t, "*", rawCfg.BindAddress)

	assert.True(t, rawCfg.DNS.Enable)
	assert.Equal(t, []string{"1.1.1.1"}, rawCfg.DNS.NameServer)
	assert.True(t, rawCfg.DNS.UseHosts)
	assert.True(t, rawCfg.DNS.UseSystemHosts)
	assert.Equal(t, "198.18.0.1/16", rawCfg.DNS.FakeIPRange)

	assert.True(t, rawCfg.Tun.Enable)
	assert.Equal(t, C.TunGvisor, rawCfg.Tun.Stack)
	assert.True(t, rawCfg.Tun.AutoRoute)
	assert.True(t, rawCfg.Tun.AutoDetectInterface)
}
