# Outbound routing and GeoIP

The **Outbounds** page controls where traffic exits. Rules are evaluated top to
bottom; the first match wins, and the default when nothing matches is DIRECT.
A rule matches by domain, CIDR, or GeoIP category, and does one of PROXY,
DIRECT, or REJECT.

All of this is optional - with no rules, mita connects directly.

## GeoIP and GeoSite datasets

The panel reads two xray-format dataset kinds, both selectable on the Outbounds
page (or mounted into the datasets dir yourself):

- **GeoIP** (`geoip.dat`) - lists of **IP ranges**. A category expands to CIDRs.
- **GeoSite** (`geosite.dat`) - lists of **domains**. A category expands to
  domain suffixes matched by mita. Regex and keyword entries are skipped
  (mita matches domains exactly or by dot-separated suffix only).

```yaml
# docker-compose.yml, under the service:
volumes:
  - ./geoip:/data/geoip     # geoip.dat and geosite.dat files
```

Mounted files are named `<name>.dat` for GeoIP and `<name>.site.dat` for
GeoSite. Built-in presets:

- **RuNet Freedom** ([russia-v2ray-rules-dat](https://github.com/runetfreedom/russia-v2ray-rules-dat)) -
  both a GeoIP and a GeoSite dataset. RU-focused categories in one file: `ru`,
  `ru-blocked`, `ru-blocked-community`, plus per-service entries (`telegram`,
  `youtube`, `discord`, `cloudflare`, ...).
- **Loyalsoldier** - per-country GeoIP codes (`ru`, `cn`, `us`, ...).
- **v2fly** - the community GeoSite list (`telegram`, `google`, `netflix`, ...).

In a rule, pick categories from the combined list (IP categories are tagged
`ip:`, domain categories `site:`); the panel expands GeoIP to CIDRs and GeoSite
to domains. Categories from all added datasets of the same kind are merged.

Large categories expand to many CIDRs (RU-blocked is ~90k), which enlarges the
server config. The panel keeps the category name in the UI and expands it only
when pushing to mita.

## Keep local traffic local

Everything inside your country goes out directly; the rest is proxied through an
upstream.

1. Add an outbound proxy (or a [chained panel](chaining-panels.md)) named e.g.
   `exit`.
2. Rule 1: match GeoIP `ru` → DIRECT.
3. Rule 2: match `*` → PROXY via `exit`.

## Proxy only blocked sites

Everything is direct except a blocklist, which is proxied.

1. Add the `ru-blocked` dataset.
2. Rule 1: match GeoIP `ru-blocked` → PROXY via `exit`.
3. Rule 2: match `*` → DIRECT.
