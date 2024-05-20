package dhcp_test

import (
	"net"
	"testing"
	"time"

	"beryju.io/gravity/pkg/instance"
	"beryju.io/gravity/pkg/roles/dhcp"
	"beryju.io/gravity/pkg/roles/dhcp/types"
	"beryju.io/gravity/pkg/tests"
	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/stretchr/testify/assert"
)

var DHCPRequestPayload = []byte{1, 1, 6, 0, 136, 9, 170, 249, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 68, 144, 187, 102, 50, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 99, 130, 83, 99, 53, 1, 3, 55, 9, 1, 121, 3, 6, 15, 108, 114, 119, 252, 57, 2, 5, 220, 61, 7, 1, 68, 144, 187, 102, 50, 4, 50, 4, 10, 120, 20, 64, 51, 4, 0, 118, 167, 0, 12, 14, 106, 101, 110, 115, 45, 105, 112, 104, 111, 110, 101, 45, 49, 50, 255, 0, 0, 0, 0}

func TestDHCPRequest(t *testing.T) {
	defer tests.Setup(t)()
	rootInst := instance.New()
	ctx := tests.Context()
	inst := rootInst.ForRole("dhcp", ctx)
	role := dhcp.New(inst)

	tests.PanicIfError(inst.KV().Put(
		ctx,
		inst.KV().Key(
			types.KeyRole,
			types.KeyScopes,
			"test",
		).String(),
		tests.MustJSON(dhcp.Scope{
			SubnetCIDR: "10.100.0.0/24",
			Default:    true,
			TTL:        86400,
			IPAM: map[string]string{
				"type":        "internal",
				"range_start": "10.100.0.100",
				"range_end":   "10.100.0.250",
			},
		}),
	))
	tests.PanicIfError(role.Start(ctx, []byte(tests.MustJSON(dhcp.RoleConfig{
		Port: 1067,
	}))))
	defer role.Stop()

	req, err := dhcpv4.FromBytes(DHCPRequestPayload)
	assert.NoError(t, err)
	req4 := role.NewRequest4(req)
	res := role.Handler4(req4)
	assert.NotNil(t, res)
	assert.Equal(t, "10.100.0.100", res.YourIPAddr.String())
	ones, bits := res.SubnetMask().Size()
	assert.Equal(t, 24, ones)
	assert.Equal(t, 32, bits)
	assert.Equal(t, "44:90:bb:66:32:04", res.ClientHWAddr.String())
	assert.Equal(t, 86400*time.Second, res.IPAddressLeaseTime(1*time.Second))
}

func TestDHCPRequestDNS(t *testing.T) {
	defer tests.Setup(t)()
	rootInst := instance.New()
	ctx := tests.Context()
	inst := rootInst.ForRole("dhcp", ctx)
	role := dhcp.New(inst)

	tests.PanicIfError(inst.KV().Put(
		ctx,
		inst.KV().Key(
			types.KeyRole,
			types.KeyScopes,
			"test",
		).String(),
		tests.MustJSON(dhcp.Scope{
			SubnetCIDR: "10.100.0.0/24",
			Default:    true,
			TTL:        86400,
			DNS: &dhcp.ScopeDNS{
				Zone:              "test.gravity.beryju.io",
				AddZoneInHostname: true,
			},
			IPAM: map[string]string{
				"type":        "internal",
				"range_start": "10.100.0.100",
				"range_end":   "10.100.0.250",
			},
		}),
	))

	tests.PanicIfError(role.Start(ctx, []byte(tests.MustJSON(dhcp.RoleConfig{
		Port: 1067,
	}))))
	defer role.Stop()

	req, err := dhcpv4.FromBytes(DHCPRequestPayload)
	assert.NoError(t, err)
	req4 := role.NewRequest4(req)
	res := role.Handler4(req4)
	assert.NotNil(t, res)
	assert.Equal(t, "10.100.0.100", res.YourIPAddr.String())
	ones, bits := res.SubnetMask().Size()
	assert.Equal(t, 24, ones)
	assert.Equal(t, 32, bits)
	assert.Equal(t, "44:90:bb:66:32:04", res.ClientHWAddr.String())
	assert.Equal(t, 86400*time.Second, res.IPAddressLeaseTime(1*time.Second))
	assert.Equal(t, "test.gravity.beryju.io", res.DomainName())
}

func TestDHCPRequestDNS_ChangedScope(t *testing.T) {
	defer tests.Setup(t)()
	rootInst := instance.New()
	ctx := tests.Context()
	inst := rootInst.ForRole("dhcp", ctx)
	role := dhcp.New(inst)

	tests.PanicIfError(inst.KV().Put(
		ctx,
		inst.KV().Key(
			types.KeyRole,
			types.KeyScopes,
			"test",
		).String(),
		tests.MustJSON(dhcp.Scope{
			SubnetCIDR: "10.100.0.0/24",
			Default:    true,
			TTL:        86400,
			DNS: &dhcp.ScopeDNS{
				Zone:              "test.gravity.beryju.io",
				AddZoneInHostname: true,
			},
			IPAM: map[string]string{
				"type":        "internal",
				"range_start": "10.100.0.100",
				"range_end":   "10.100.0.250",
			},
		}),
	))
	tests.PanicIfError(inst.KV().Put(
		ctx,
		inst.KV().Key(
			types.KeyRole,
			types.KeyScopes,
			"test2",
		).String(),
		tests.MustJSON(dhcp.Scope{
			SubnetCIDR: "10.200.0.0/24",
			TTL:        86400,
			DNS: &dhcp.ScopeDNS{
				Zone:              "test2.gravity.beryju.io",
				AddZoneInHostname: true,
			},
			IPAM: map[string]string{
				"type":        "internal",
				"range_start": "10.200.0.100",
				"range_end":   "10.200.0.250",
			},
		}),
	))

	tests.PanicIfError(role.Start(ctx, []byte(tests.MustJSON(dhcp.RoleConfig{
		Port: 1067,
	}))))
	defer role.Stop()

	// First ensure the lease is created as we expect
	req, err := dhcpv4.FromBytes(DHCPRequestPayload)
	assert.NoError(t, err)
	req4 := role.NewRequest4(req)
	res := role.Handler4(req4)
	assert.NotNil(t, res)
	assert.Equal(t, "10.100.0.100", res.YourIPAddr.String())
	ones, bits := res.SubnetMask().Size()
	assert.Equal(t, 24, ones)
	assert.Equal(t, 32, bits)
	assert.Equal(t, "44:90:bb:66:32:04", res.ClientHWAddr.String())
	assert.Equal(t, 86400*time.Second, res.IPAddressLeaseTime(1*time.Second))
	assert.Equal(t, "test.gravity.beryju.io", res.DomainName())

	// Now we're requesting an IP from another subnet, so the lease should move
	req.GatewayIPAddr = net.ParseIP("10.200.0.1")
	req4 = role.NewRequest4(req)
	res = role.Handler4(req4)
	assert.NotNil(t, res)
	assert.Equal(t, "10.200.0.100", res.YourIPAddr.String())
	ones, bits = res.SubnetMask().Size()
	assert.Equal(t, 24, ones)
	assert.Equal(t, 32, bits)
	assert.Equal(t, "44:90:bb:66:32:04", res.ClientHWAddr.String())
	assert.Equal(t, 86400*time.Second, res.IPAddressLeaseTime(1*time.Second))
	assert.Equal(t, "test2.gravity.beryju.io", res.DomainName())
}
