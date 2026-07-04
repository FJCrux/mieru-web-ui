<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { NCard, NGrid, NGi, NStatistic, NTag, NDataTable, NCollapse, NCollapseItem, NCode, NButton, NSpace, NAlert, useMessage } from 'naive-ui'
import type { DataTableColumns } from 'naive-ui'
import { api, ApiError } from '../api/client'
import type { Dashboard, SessionInfo } from '../api/client'

const message = useMessage()
const dash = ref<Dashboard | null>(null)
const sessions = ref<SessionInfo[]>([])
const mitaLogs = ref<string[]>([])
let timer: number | undefined

function fmtBytes(n: number): string {
  if (n < 1024) return `${n} B`
  const units = ['KiB', 'MiB', 'GiB', 'TiB']
  let v = n
  let i = -1
  do {
    v /= 1024
    i++
  } while (v >= 1024 && i < units.length - 1)
  return `${v.toFixed(1)} ${units[i]}`
}

function fmtUptime(s: number): string {
  if (s <= 0) return '—'
  const d = Math.floor(s / 86400)
  const h = Math.floor((s % 86400) / 3600)
  const m = Math.floor((s % 3600) / 60)
  return d > 0 ? `${d}d ${h}h` : h > 0 ? `${h}h ${m}m` : `${m}m`
}

const traffic = computed(() => {
  const t = dash.value?.metrics?.['traffic'] ?? {}
  return { down: t['DownloadBytes'] ?? 0, up: t['UploadBytes'] ?? 0 }
})

const statusType = computed(() =>
  dash.value?.mitaStatus === 'RUNNING' ? 'success' : dash.value?.mitaStatus === 'UNREACHABLE' ? 'error' : 'warning',
)

const sessionColumns: DataTableColumns<SessionInfo> = [
  { title: 'ID', key: 'id' },
  { title: 'Protocol', key: 'protocol' },
  { title: 'Local', key: 'localAddr' },
  { title: 'Remote', key: 'remoteAddr' },
  { title: 'State', key: 'state' },
]

async function poll() {
  try {
    dash.value = await api.get<Dashboard>('/api/dashboard')
    sessions.value = await api.get<SessionInfo[]>('/api/sessions')
  } catch {
    /* transient errors surface on next poll */
  }
}

async function loadLogs() {
  try {
    const res = await api.get<{ lines: string[] }>('/api/mita/logs?lines=200')
    mitaLogs.value = res.lines
  } catch (e) {
    if (e instanceof ApiError && e.status === 501) mitaLogs.value = ['(panel is not supervising mita)']
  }
}

async function restartMita() {
  try {
    await api.post('/api/mita/restart')
    message.success('Restart requested')
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : 'Restart failed')
  }
}

function startPolling() {
  poll()
  timer = window.setInterval(() => {
    if (!document.hidden) poll()
  }, 5000)
}

onMounted(startPolling)
onUnmounted(() => window.clearInterval(timer))
</script>

<template>
  <h2 class="page-title">Dashboard</h2>
  <n-space vertical :size="16">
    <n-alert v-if="dash?.warnings?.length" type="warning" title="Hardening">
      <ul class="warnlist">
        <li v-for="(wmsg, i) in dash.warnings" :key="i">{{ wmsg }}</li>
      </ul>
    </n-alert>
    <n-grid :cols="5" :x-gap="12" :y-gap="12" item-responsive responsive="screen">
      <n-gi span="5 m:1">
        <div class="tile">
          <div class="tile-label">mita</div>
          <n-tag :type="statusType" size="small" round>{{ dash?.mitaStatus ?? '…' }}</n-tag>
          <div class="tile-sub">v{{ dash?.mitaVersion || '-' }} · up {{ fmtUptime(dash?.mitaUptimeSeconds ?? 0) }}</div>
        </div>
      </n-gi>
      <n-gi span="5 m:1">
        <div class="tile"><div class="tile-label">Users</div><div class="tile-value">{{ dash?.userCount ?? 0 }}</div></div>
      </n-gi>
      <n-gi span="5 m:1">
        <div class="tile"><div class="tile-label">Active sessions</div><div class="tile-value">{{ dash?.sessionCount ?? 0 }}</div></div>
      </n-gi>
      <n-gi span="5 m:1">
        <div class="tile"><div class="tile-label">Download</div><div class="tile-value">{{ fmtBytes(traffic.down) }}</div></div>
      </n-gi>
      <n-gi span="5 m:1">
        <div class="tile"><div class="tile-label">Upload</div><div class="tile-value">{{ fmtBytes(traffic.up) }}</div></div>
      </n-gi>
    </n-grid>

    <n-card title="Active sessions">
      <n-data-table :columns="sessionColumns" :data="sessions" :row-key="(s: SessionInfo) => s.id" size="small" />
    </n-card>

    <n-card title="mita">
      <template #header-extra>
        <n-button size="small" type="warning" secondary @click="restartMita">Restart mita</n-button>
      </template>
      <n-collapse @item-header-click="loadLogs">
        <n-collapse-item title="Logs (last 200 lines)" name="logs">
          <n-code :code="mitaLogs.join('\n') || '(empty)'" language="text" word-wrap />
        </n-collapse-item>
        <n-collapse-item title="Raw metrics" name="metrics">
          <n-code :code="JSON.stringify(dash?.metrics ?? {}, null, 2)" language="json" word-wrap />
        </n-collapse-item>
      </n-collapse>
    </n-card>
  </n-space>
</template>

<style scoped>
.tile {
  background: #1a1a1f;
  border: 1px solid #26262c;
  border-radius: 12px;
  padding: 16px 18px;
  min-height: 84px;
}
.tile-label {
  font-size: 12px;
  text-transform: uppercase;
  letter-spacing: 0.6px;
  opacity: 0.5;
  margin-bottom: 8px;
}
.tile-value {
  font-size: 24px;
  font-weight: 650;
}
.tile-sub {
  margin-top: 8px;
  font-size: 12px;
  opacity: 0.55;
}
.warnlist {
  margin: 0;
  padding-left: 18px;
}
.warnlist li {
  margin: 2px 0;
}
</style>
