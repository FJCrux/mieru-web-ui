// Panel base path injected by the server (empty for root, e.g. "/panel").
export const BASE: string = (window as unknown as { __BASE__?: string }).__BASE__ ?? ''
