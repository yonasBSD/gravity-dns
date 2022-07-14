package dhcp

import (
	"context"
	"net"
	"sync"

	"beryju.io/gravity/pkg/roles"
	"beryju.io/gravity/pkg/roles/dhcp/types"
	log "github.com/sirupsen/logrus"

	"github.com/insomniacslk/dhcp/dhcpv4/server4"
	"github.com/insomniacslk/dhcp/dhcpv6/server6"
)

type DHCPRole struct {
	scopes     map[string]*Scope
	leases     map[string]*Lease
	leasesSync sync.RWMutex

	cfg *DHCPRoleConfig

	s4  *server4.Server
	s6  *server6.Server
	log *log.Entry
	i   roles.Instance
	ctx context.Context
}

func New(instance roles.Instance) *DHCPRole {
	r := &DHCPRole{
		log:        instance.GetLogger().WithField("role", types.KeyRole),
		i:          instance,
		scopes:     make(map[string]*Scope),
		leases:     make(map[string]*Lease),
		leasesSync: sync.RWMutex{},
	}
	r.i.AddEventListener(types.EventTopicDHCPCreateLease, r.eventCreateLease)
	return r
}

func (r *DHCPRole) Start(ctx context.Context, config []byte) error {
	r.ctx = ctx
	r.cfg = r.decodeDHCPRoleConfig(config)

	r.loadInitialScopes()
	r.loadInitialScopes()

	go r.startWatchScopes()
	go r.startWatchLeases()

	return r.startServer4()
}

func (r *DHCPRole) startServer4() error {
	laddr := net.UDPAddr{
		IP:   net.ParseIP("0.0.0.0"),
		Port: r.cfg.Port,
	}
	server, err := server4.NewServer(
		"", // TODO: specify interface to DHCP?
		&laddr,
		r.recoverMiddleware4(
			r.loggingMiddleware4(
				r.handler4,
			),
		),
	)
	if err != nil {
		return err
	}
	r.s4 = server
	r.log.WithField("port", r.cfg.Port).Info("starting DHCP Server")
	return r.s4.Serve()
}

func (r *DHCPRole) Stop() {
	if r.s4 != nil {
		r.s4.Close()
	}
}