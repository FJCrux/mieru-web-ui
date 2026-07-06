# Sharing client configs

Each mieru user needs the server address plus their credentials. The panel
generates these for you.

## From the Users page

Click **Share** on a user to get:

- a **mierus://** link and QR code (simple format),
- a **mieru://** link and QR code (full config),
- the client config JSON to download.

Set **Public host** in Settings first - it's the address clients connect to.

## Expiring links

The **Link** tab creates a tokenized URL (`https://host/s/<token>`) that expires
(15 min / 1 hour / 24 hours). Send that instead of the raw credentials:

- The token is unguessable and single-purpose.
- The link auto-expires; a leaked old link stops working.
- Resetting the user's password invalidates their links immediately.
- The link lives on a separate path (`SHARE_PATH`, default `/s`) and never
  reveals the admin panel's location.

The recipient opens the link and sees a page with the QR and connection string -
no login, no admin UI.

## Subscriptions

The **Subscription** tab issues a permanent URL (`https://host/sub/<token>`)
that clients refresh automatically. Unlike expiring links it survives password
changes - the config is built from the current password on every fetch.

- The default response is a ready-to-import **Clash profile** (proxies, a
  PROXY group, minimal rules). Works with any client on mihomo >= 1.19.0:
  Mihomo Party, Clash Verge Rev, FlClash, ClashMi.
- `?format=mierus` / `?format=mieru` return the native links as a
  base64-encoded line list (the URI-subscription convention) for clients like
  NekoBox/Exclave with the mieru plugin.
- `?flavor=proxies` returns a proxies-only document for `proxy-provider`.
- The response carries `subscription-userinfo` (used traffic and the quota
  cap), so clients show a usage bar.
- **Rotate** issues a new token and kills the old URL everywhere; **Revoke**
  disables it. Deleting the user does too.
- The path is configurable (`SUB_PATH`, default `/sub`). With `SUB_PORT` the
  subscriptions get a dedicated listener on their own port (the panel and
  share pages stay off it); a dedicated port requires TLS (`SUB_TLS_CERT` /
  `SUB_TLS_KEY`, falling back to the panel certs) unless `SUB_INSECURE_HTTP=1`
  is set for a TLS-terminating reverse proxy in front.

## Users created outside the panel

The panel can only build links for users it created (it stores their password;
mita keeps only a hash). For a user added via the CLI, reset their password in
the panel to enable sharing.
