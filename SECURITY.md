# Security policy

## Reporting a vulnerability

Please report security issues privately via GitHub's "Report a vulnerability"
(Security → Advisories) rather than opening a public issue. I'll aim to respond
within a few days.

## Threat model and expectations

- **Run the panel behind TLS or on loopback.** Over plain HTTP on a public
  address, credentials and session cookies travel in the clear. The dashboard
  warns when it detects this. Use a reverse proxy with TLS, `PANEL_TLS_CERT`/
  `PANEL_TLS_KEY`, or bind to `127.0.0.1` and reach it through an SSH tunnel.
- **The panel stores mieru user passwords in plaintext** in its SQLite DB
  (file mode 0600). This is required to regenerate client configs and share
  links, since mita only keeps hashes. Protect the `/data` volume; anyone who
  reads it can impersonate your proxy users.
- **Share links and QR codes are secrets.** A `mierus://` / `mieru://` link
  contains a user's credentials. Send it over a trusted channel, and if one
  leaks, reset that user's password (which invalidates the old link).
- **Panel admin** is a single account: bcrypt-hashed password, opaque session
  tokens stored hashed, `HttpOnly` + `SameSite=Lax` cookies, a CSRF header
  check on mutations, and login rate limiting. Pick a strong password and
  change the bootstrap one on first login.
- **Defense in depth** (see [docs/guides/hardening.md](docs/guides/hardening.md)):
  set `PANEL_URL` to reject requests with an unexpected Host, `PANEL_BASE_PATH`
  to serve the panel under a secret prefix (everything else 404s), and keep the
  panel on loopback behind a reverse proxy. Share links live on a separate path
  so they don't reveal the admin location.
