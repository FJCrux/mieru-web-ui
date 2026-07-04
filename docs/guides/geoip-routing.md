# Outbound routing and GeoIP

The **Outbounds** page controls where traffic exits. Rules are evaluated top to
bottom; the first match wins, and the default when nothing matches is DIRECT.
A rule matches by domain, CIDR, or GeoIP category, and does one of PROXY,
DIRECT, or REJECT.

All of this is optional - with no rules, mita connects directly.

## GeoIP datasets

GeoIP uses the same `geoip.dat` format as xray/3x-ui. Add a dataset by URL on
the Outbounds page (presets for Loyalsoldier countries and the RU-blocked list
are built in), or mount your own:

```yaml
# docker-compose.yml, under the service:
volumes:
  - ./geoip:/data/geoip     # drop any geoip.dat files here
```

Then reference a category (e.g. `ru`, `cn`, `ru-blocked`) in a rule.

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
