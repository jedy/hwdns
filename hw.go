package hwdns

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/jedy/libhw"
)

// Provider wraps the provider implementation as a Caddy module.
type Provider struct{ *libhw.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID: "dns.providers.hwdns",
		New: func() caddy.Module {
			return &Provider{new(libhw.Provider)}
		},
	}
}

// Provision implements the Provisioner interface to initialize the client
func (p *Provider) Provision(ctx caddy.Context) error {
	repl := caddy.NewReplacer()
	p.Provider.AccessKeyId = repl.ReplaceAll(p.Provider.AccessKeyId, "")
	p.Provider.SecretAccessKey = repl.ReplaceAll(p.Provider.SecretAccessKey, "")
	p.Provider.Region = repl.ReplaceAll(p.Provider.Region, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	hwdns {
//	    access_key_id <string>
//	    secret_access_key <string>
//	    region <string>
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "access_key_id":
				if d.NextArg() {
					p.Provider.AccessKeyId = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "secret_access_key":
				if d.NextArg() {
					p.Provider.SecretAccessKey = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "region":
				if d.NextArg() {
					p.Provider.Region = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}

	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
