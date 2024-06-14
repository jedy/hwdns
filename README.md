# Huawei Cloud DNS module for Caddy

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with Huawei Cloud accounts.

## Caddy module name

```
dns.providers.hwdns
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
  "module": "acme",
  "challenges": {
    "dns": {
      "provider": {
        "name": "hwdns",
        "access_key_id":"YOUR_HUAWEICLOUD_ACCESS_KEY_ID",
        "secret_access_key":"YOUR_HUAWEICLOUD_SECRET_ACCESS_KEY",
        "region": "PREFERRED_REGION_ID"
      }
    }
  }
}
```

or with the Caddyfile:

```
# globally

acme_dns hwdns {
  access_key_id {env.HUAWEICLOUD_ACCESS_KEY_ID}
  secret_access_key {env.HUAWEICLOUD_SECRET_ACCESS_KEY}
  region PREFERRED_REGION_ID
}
```

```
# one site

tls {
  # you may need to add this when DNS challenge times out
  resolvers ns1.huaweicloud-dns.org
  dns hwdns {
    access_key_id {env.HUAWEICLOUD_ACCESS_KEY_ID}
    secret_access_key {env.HUAWEICLOUD_SECRET_ACCESS_KEY}
    region PREFERRED_REGION_ID
  }
}
```

You can replace `{env.HUAWEICLOUD_ACCESS_KEY_ID}`,`{env.HUAWEICLOUD_SECRET_ACCESS_KEY}` with the actual auth token in the `""` if you prefer to put it directly in your config instead of an environment variable.

You can omit the "region" if you don't have any specific requirements.
