# Chaining two panels

Route one panel's egress through another mieru server, so traffic exits from the
second server. Both run the same image. This is optional.

## Setup

1. On the **upstream** panel (B): set its Public host in Settings, open **Chain**,
   click **Generate peer key**, and copy the `mieru://` key.
2. On the **downstream** panel (A): open **Chain → Upstream panels**, give it a
   name, paste the key, and **Connect**. A starts a client to B; the peer now
   appears as a proxy.
3. On A, open **Outbounds** and route traffic through the peer - a `* -> PROXY
   [name]` rule sends everything through B, or combine with
   [GeoIP](geoip-routing.md) to send only some traffic.

## How it works

mieru's egress speaks SOCKS5, and a mita server is not a SOCKS5 endpoint. So for
each peer, panel A runs a supervised `mieru` client connected to B that exposes a
local SOCKS5 port; that port is registered as an egress proxy. The client is
restarted automatically and survives panel restarts.

## Notes

- The peer key contains B's credentials - share it privately.
- Remove a peer from the Chain page to stop its client and drop the egress proxy.
- Each peer generates a dedicated user on B; delete it there to revoke access.
