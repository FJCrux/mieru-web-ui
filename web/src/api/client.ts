// Thin fetch wrapper: JSON in/out, CSRF header on mutations,
// redirect to /login on 401.
import router from '../router'
import { BASE } from '../base'

export class ApiError extends Error {
  status: number
  constructor(status: number, message: string) {
    super(message)
    this.status = status
  }
}

async function request<T>(method: string, path: string, body?: unknown): Promise<T> {
  const headers: Record<string, string> = {
    'X-Requested-With': 'XMLHttpRequest',
  }
  if (body !== undefined) headers['Content-Type'] = 'application/json'

  const res = await fetch(BASE + path, {
    method,
    headers,
    body: body !== undefined ? JSON.stringify(body) : undefined,
    credentials: 'same-origin',
  })

  if (res.status === 401 && router.currentRoute.value.path !== '/login') {
    router.push('/login')
    throw new ApiError(401, 'unauthorized')
  }

  const text = await res.text()
  let data: unknown = {}
  if (text) {
    try {
      data = JSON.parse(text)
    } catch {
      data = { error: text }
    }
  }
  if (!res.ok) {
    const msg =
      typeof data === 'object' && data !== null && 'error' in data
        ? String((data as { error: unknown }).error)
        : `HTTP ${res.status}`
    throw new ApiError(res.status, msg)
  }
  return data as T
}

export const api = {
  get: <T>(path: string) => request<T>('GET', path),
  post: <T>(path: string, body?: unknown) => request<T>('POST', path, body),
  put: <T>(path: string, body?: unknown) => request<T>('PUT', path, body),
  del: <T>(path: string) => request<T>('DELETE', path),
}

// --- API types mirrored from the Go handlers ---

export interface Quota {
  days: number
  megabytes: number
}

export interface UserInfo {
  name: string
  quotas: Quota[]
  allowPrivateIP: boolean
  hasSecret: boolean
  metrics: Record<string, number>
}

export interface PortBinding {
  port?: number
  portRange?: string
  protocol: string
}

export interface NetworkConfig {
  portBindings: PortBinding[]
  mtu: number
  loggingLevel: string
  portsManaged?: boolean
}

export interface Dashboard {
  mitaStatus: string
  mitaVersion: string
  restarts: number
  mitaUptimeSeconds: number
  sessionCount: number
  userCount: number
  metrics: Record<string, Record<string, number>>
  insecureAccess: boolean
  warnings: string[]
}

export interface SessionInfo {
  id: string
  protocol: string
  localAddr: string
  remoteAddr: string
  state: string
}

export interface ShareLinks {
  clientConfigJson: string
  mieruUrl: string
  mierusUrls: string[]
}

export interface ShareToken {
  url: string
  expiresAt: number
}

export interface EgressProxy {
  name: string
  host: string
  port: number
  username: string
  password: string
}

export interface EgressRule {
  domains: string[]
  cidrs: string[]
  geo: string[]
  action: string // PROXY, DIRECT, REJECT
  proxies: string[]
}

export interface EgressConfig {
  proxies: EgressProxy[]
  rules: EgressRule[]
}

export interface GeoDataset {
  name: string
  bytes: number
}

export interface GeoCategory {
  code: string
  cidrs: number
}

export interface GeoipState {
  datasets: GeoDataset[]
  categories: GeoCategory[]
}

export interface Peer {
  name: string
  socks5Port: number
  running: boolean
  restarts: number
}

export interface ChainKey {
  name: string
  key: string
}

export interface Settings {
  publicHost: string
  panelUrl: string
  basePath: string
  sharePath: string
  restartPending: boolean
}
