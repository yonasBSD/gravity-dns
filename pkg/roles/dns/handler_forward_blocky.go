package dns

import (
	"fmt"
	"net/netip"
	"strings"
	"time"

	"beryju.io/gravity/pkg/extconfig"
	"beryju.io/gravity/pkg/roles/dns/types"
	"beryju.io/gravity/pkg/roles/dns/utils"
	"github.com/creasty/defaults"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"

	"github.com/0xERR0R/blocky/config"
	blockylog "github.com/0xERR0R/blocky/log"
	"github.com/0xERR0R/blocky/server"
	"github.com/miekg/dns"
)

const BlockyForwarderType = "forward_blocky"

type BlockyForwarder struct {
	*IPForwarderHandler
	c   map[string]interface{}
	b   *server.Server
	log *zap.Logger
	st  time.Time
	cfg *config.Config
}

func init() {
	HandlerRegistry.Add(BlockyForwarderType, func(z *Zone, rawConfig map[string]interface{}) Handler {
		return NewBlockyForwarder(z, rawConfig)
	})
}

func HTTPByteSource(url string) config.BytesSource {
	return config.BytesSource{
		Type: config.BytesSourceTypeHttp,
		From: url,
	}
}

func TextByteSource(content string) config.BytesSource {
	return config.BytesSource{
		Type: config.BytesSourceTypeText,
		From: content,
	}
}

func NewBlockyForwarder(z *Zone, rawConfig map[string]interface{}) *BlockyForwarder {
	bfwd := &BlockyForwarder{
		IPForwarderHandler: NewIPForwarderHandler(z, rawConfig),
		c:                  rawConfig,
		st:                 time.Now(),
	}
	bfwd.log = z.log.With(zap.String("handler", bfwd.Identifier()))
	bfwd.startBlockyListServer()
	waitForStart := func() {
		err := bfwd.setup()
		if err != nil {
			bfwd.log.Warn("failed to setup blocky, queries will fallthrough", zap.Error(err))
		}
	}
	cfg, err := bfwd.getConfig()
	if err != nil {
		bfwd.log.Warn("Failed to build blocky config", zap.Error(err))
	}
	bfwd.cfg = cfg

	if extconfig.Get().Debug {
		bfwd.log.Info("starting blocky sync")
		waitForStart()
	} else {
		bfwd.log.Info("starting blocky async")
		go waitForStart()
	}
	return bfwd
}

func (bfwd *BlockyForwarder) Identifier() string {
	return BlockyForwarderType
}

// Generate the configuration for Blocky based on parameters specified.
// Used if no full custom config is specified
func (bfwd *BlockyForwarder) generateConfig(cfg *config.Config) error {
	forwarders := []string{}
	switch v := bfwd.c["to"].(type) {
	case string:
		forwarders = strings.Split(v, ";")
	case []interface{}:
		for _, fwd := range v {
			if f, ok := fwd.(string); ok {
				forwarders = append(forwarders, f)
			}
		}
	}
	upstreams := make([]config.Upstream, len(forwarders))
	for idx, fwd := range forwarders {
		us, err := config.ParseUpstream(fwd)
		if err != nil {
			bfwd.log.Warn("failed to parse upstream", zap.Error(err))
			continue
		}
		upstreams[idx] = us
	}

	blockLists := []config.BytesSource{
		HTTPByteSource(blockyListBase + "AdGuardSDNSFilter.txt"),
		HTTPByteSource(blockyListBase + "AdguardDNS.txt"),
		HTTPByteSource(blockyListBase + "Easylist.txt"),
		HTTPByteSource(blockyListBase + "StevenBlack.hosts.txt"),
		HTTPByteSource(blockyListBase + "adaway.org.txt"),
		HTTPByteSource(blockyListBase + "big.oisd.nl.txt"),
	}
	if bll, ok := bfwd.c["blocklists"]; ok {
		blockLists = bfwd.getLists(bll)
	}

	allowLists := []config.BytesSource{}
	if all, ok := bfwd.c["allowlists"]; ok {
		allowLists = bfwd.getLists(all)
	}
	bootstrap, err := netip.ParseAddrPort(extconfig.Get().FallbackDNS)
	if err != nil {
		return err
	}
	cfg.BootstrapDNS = config.BootstrapDNS{
		{
			Upstream: config.Upstream{
				Net:  config.NetProtocolTcpUdp,
				Host: bootstrap.Addr().String(),
				Port: bootstrap.Port(),
			},
		},
	}
	cfg.Upstreams = config.Upstreams{
		Groups: map[string][]config.Upstream{
			"default": upstreams,
		},
		Timeout: config.Duration(types.DefaultUpstreamTimeout),
	}
	cfg.Blocking = config.Blocking{
		BlockType: "zeroIP",
		Denylists: map[string][]config.BytesSource{
			"default": blockLists,
		},
		Allowlists: map[string][]config.BytesSource{
			"default": allowLists,
		},
		Loading: config.SourceLoading{
			Downloads: config.Downloader{
				Attempts: 3,
			},
		},
		ClientGroupsBlock: map[string][]string{
			"default": {"default"},
		},
	}
	return nil
}

func (bfwd *BlockyForwarder) getConfig() (*config.Config, error) {
	cfg := config.Config{}
	if ccfg, ok := bfwd.c["config"].(string); ok {
		err := yaml.UnmarshalStrict([]byte(ccfg), &cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to parse specified config: %w", err)
		}
	} else {
		err := defaults.Set(&cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to set config defaults: %w", err)
		}
		err = bfwd.generateConfig(&cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to generate config: %w", err)
		}
	}
	// Blocky uses a custom registry, so this doesn't work as expected
	// cfg.Prometheus.Enable = true
	cfg.Log.Level = logrus.DebugLevel
	cfg.QueryLog.Type = config.QueryLogTypeConsole
	if !extconfig.Get().Debug {
		cfg.Log.Format = blockylog.FormatTypeJson
		// Only log errors from blocky to prevent double-logging all queries
		cfg.Log.Level = logrus.FatalLevel
		cfg.QueryLog.Type = config.QueryLogTypeNone
	}
	return &cfg, nil
}

func (bfwd *BlockyForwarder) getLists(raw interface{}) []config.BytesSource {
	list := []config.BytesSource{}
	rawLists := []string{}
	switch v := raw.(type) {
	case string:
		rawLists = strings.Split(v, ";")
	case []interface{}:
		for _, rl := range v {
			if r, ok := rl.(string); ok {
				rawLists = append(rawLists, r)
			}
		}
	}
	for _, rl := range rawLists {
		if strings.HasPrefix(rl, "http") {
			list = append(list, HTTPByteSource(rl))
		} else {
			list = append(list, TextByteSource(rl))
		}
	}
	return list
}

func (bfwd *BlockyForwarder) setup() error {
	bfwd.log.Debug("blocky config", zap.Any("config", bfwd.cfg))

	srv, err := server.NewServer(bfwd.z.inst.Context(), bfwd.cfg)
	if err != nil {
		return fmt.Errorf("can't start server: %w", err)
	}
	bfwd.log.Info("finished blocky setup", zap.Duration("took", time.Since(bfwd.st)))
	bfwd.b = srv
	return nil
}

func (bfwd *BlockyForwarder) Handle(w *utils.FakeDNSWriter, r *utils.DNSRequest) *dns.Msg {
	if bfwd.b == nil {
		bfwd.log.Debug("Blocky not started/setup yet, falling through to next handler")
		return nil
	}
	bs := sentry.TransactionFromContext(r.Context()).StartChild("gravity.dns.handler.forward_blocky.handle")
	bfwd.b.OnRequest(r.Context(), w, r.Msg)
	bs.Finish()
	// fall to next handler when no record is found
	if w.Msg().Rcode == dns.RcodeNameError {
		return nil
	}
	for _, query := range r.Question {
		for idx, ans := range w.Msg().Answer {
			go bfwd.cacheToEtcd(r, query, ans, idx)
		}
	}
	return w.Msg()
}
