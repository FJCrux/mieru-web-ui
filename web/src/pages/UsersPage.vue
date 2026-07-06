<script setup lang="ts">
import { h, ref, computed, onMounted, onUnmounted } from 'vue'
import {
  NButton, NCard, NDataTable, NDrawer, NDrawerContent, NForm, NFormItem, NInput,
  NInputGroup, NInputNumber, NSpace, NSwitch, NTag, NPopconfirm, NIcon, useMessage,
} from 'naive-ui'
import type { DataTableColumns } from 'naive-ui'
import { PeopleOutline, RefreshOutline } from '@vicons/ionicons5'
import { useI18n } from 'vue-i18n'
import { api, ApiError } from '../api/client'
import type { UserInfo, Quota } from '../api/client'
import ShareModal from '../components/ShareModal.vue'
import HelpLabel from '../components/HelpLabel.vue'

const message = useMessage()
const { t } = useI18n()
const users = ref<UserInfo[]>([])
const loading = ref(false)
// Ticks every few seconds so relative "last active" labels and the online
// badge stay fresh between polls.
const now = ref(Date.now())
let poll: number | undefined
let clock: number | undefined

// A user counts as online when they moved data within the last 5 minutes
// (mita can't attribute live sessions to users, so activity is inferred from
// per-user traffic counters).
const ONLINE_WINDOW_MS = 5 * 60 * 1000

function relativeActive(ms: number): string {
  const diff = Math.max(0, now.value - ms)
  const mins = Math.floor(diff / 60000)
  if (mins < 1) return t('users.momentsAgo')
  if (mins < 60) return t('users.minutesAgo', { n: mins })
  const hours = Math.floor(mins / 60)
  if (hours < 24) return t('users.hoursAgo', { n: hours })
  return t('users.daysAgo', { n: Math.floor(hours / 24) })
}

// Total bytes this user has ever transferred. mita keeps the cumulative
// counter even after the timestamped history that drives "last active" has
// rolled off, so traffic > 0 with no timestamp still means "used it before".
function totalTraffic(u: UserInfo): number {
  return Object.entries(u.metrics).reduce((sum, [k, v]) => {
    const key = k.toLowerCase()
    return key.includes('download') || key.includes('upload') ? sum + v : sum
  }, 0)
}

const drawerOpen = ref(false)
const editingName = ref<string | null>(null) // null = create mode
const form = ref({ name: '', password: '', allowPrivateIP: false })
const quotas = ref<Quota[]>([])

const shareUser = ref<string | null>(null)

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

function trafficOf(u: UserInfo): string {
  const entries = Object.entries(u.metrics)
  if (!entries.length) return '—'
  const down = entries.find(([k]) => k.toLowerCase().includes('download'))?.[1] ?? 0
  const up = entries.find(([k]) => k.toLowerCase().includes('upload'))?.[1] ?? 0
  return `↓ ${fmtBytes(down)} / ↑ ${fmtBytes(up)}`
}

const columns = computed<DataTableColumns<UserInfo>>(() => [
  { title: t('users.name'), key: 'name' },
  {
    title: t('users.colStatus'),
    key: 'status',
    render: (u) => {
      const online = u.lastActiveUnixMs > 0 && now.value - u.lastActiveUnixMs < ONLINE_WINDOW_MS
      let label: string
      if (online) {
        label = t('users.online')
      } else if (u.lastActiveUnixMs > 0) {
        label = relativeActive(u.lastActiveUnixMs) // known last-seen time
      } else if (totalTraffic(u) > 0) {
        label = t('users.offline') // used it before, but no recent activity data
      } else {
        label = t('users.neverActive') // never transferred anything
      }
      return h(
        NTag,
        { type: online ? 'success' : 'default', size: 'small', round: true },
        { default: () => label },
      )
    },
  },
  {
    title: t('users.colQuotas'),
    key: 'quotas',
    render: (u) =>
      u.quotas.length
        ? u.quotas.map((q) => h(NTag, { size: 'small', style: 'margin-right:4px' }, { default: () => `${q.megabytes} MB / ${q.days}d` }))
        : '—',
  },
  { title: t('users.colTraffic'), key: 'traffic', render: trafficOf },
  {
    title: t('users.colShare'),
    key: 'share',
    render: (u) =>
      h(NSpace, { align: 'center', wrap: false }, {
        default: () => [
          h(
            NButton,
            { size: 'small', disabled: !u.hasSecret, onClick: () => (shareUser.value = u.name) },
            { default: () => (u.hasSecret ? t('users.share') : t('users.noPassword')) },
          ),
          ...(u.hasSubscription
            ? [h(NTag, { size: 'small', round: true, type: 'info' }, { default: () => t('users.subActive') })]
            : []),
        ],
      }),
  },
  {
    title: t('common.actions'),
    key: 'actions',
    render: (u) =>
      h(NSpace, null, {
        default: () => [
          h(NButton, { size: 'small', onClick: () => openEdit(u) }, { default: () => t('common.edit') }),
          h(
            NPopconfirm,
            { onPositiveClick: () => removeUser(u.name) },
            {
              trigger: () => h(NButton, { size: 'small', type: 'error', quaternary: true }, { default: () => t('common.delete') }),
              default: () => t('users.confirmDelete', { name: u.name }),
            },
          ),
        ],
      }),
  },
])

async function load(silent = false) {
  if (!silent) loading.value = true
  try {
    users.value = await api.get<UserInfo[]>('/api/users')
    now.value = Date.now()
  } catch (e) {
    // Background polls fail quietly; only the explicit load surfaces errors.
    if (!silent) message.error(e instanceof ApiError ? e.message : t('users.loadFailed'))
  } finally {
    if (!silent) loading.value = false
  }
}

// generatePassword returns a strong, URL-safe random password (no ambiguous
// characters), built from the crypto RNG.
function generatePassword(len = 20): string {
  const alphabet = 'abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789'
  const bytes = new Uint32Array(len)
  crypto.getRandomValues(bytes)
  let out = ''
  for (let i = 0; i < len; i++) out += alphabet[bytes[i] % alphabet.length]
  return out
}

function regeneratePassword() {
  form.value.password = generatePassword()
}

function openCreate() {
  editingName.value = null
  // Pre-fill a strong random password so an admin can just create-and-share;
  // the field stays editable and can be regenerated or typed over.
  form.value = { name: '', password: generatePassword(), allowPrivateIP: false }
  quotas.value = []
  drawerOpen.value = true
}

function openEdit(u: UserInfo) {
  editingName.value = u.name
  form.value = { name: u.name, password: '', allowPrivateIP: u.allowPrivateIP }
  quotas.value = u.quotas.map((q) => ({ ...q }))
  drawerOpen.value = true
}

function addQuota() {
  quotas.value.push({ days: 30, megabytes: 10240 })
}

async function save() {
  const body = {
    password: form.value.password,
    quotas: quotas.value,
    allowPrivateIP: form.value.allowPrivateIP,
  }
  try {
    if (editingName.value === null) {
      await api.post('/api/users', { name: form.value.name, ...body })
      message.success(t('users.created'))
    } else {
      await api.put(`/api/users/${encodeURIComponent(editingName.value)}`, body)
      message.success(t('users.updated'))
    }
    drawerOpen.value = false
    await load()
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : t('common.saveFailed'))
  }
}

async function removeUser(name: string) {
  try {
    await api.del(`/api/users/${encodeURIComponent(name)}`)
    message.success(t('users.deleted'))
    await load()
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : t('common.deleteFailed'))
  }
}

onMounted(() => {
  load()
  // Refresh the list (and thus activity) periodically, and tick a clock so
  // relative labels advance without a fetch. Skip polling while a hidden tab.
  poll = window.setInterval(() => {
    if (!document.hidden && !drawerOpen.value) load(true)
  }, 10000)
  clock = window.setInterval(() => (now.value = Date.now()), 5000)
})
onUnmounted(() => {
  window.clearInterval(poll)
  window.clearInterval(clock)
})
</script>

<template>
  <n-card :title="t('users.title')">
    <template #header-extra>
      <n-button type="primary" @click="openCreate">{{ t('users.addUser') }}</n-button>
    </template>
    <n-data-table :columns="columns" :data="users" :loading="loading" :row-key="(u: UserInfo) => u.name" :scroll-x="760">
      <template #empty>
        <div class="empty">
          <n-icon :component="PeopleOutline" :size="30" class="empty-icon" />
          <div class="empty-title">{{ t('users.emptyTitle') }}</div>
          <div class="empty-sub">{{ t('users.emptySub') }}</div>
          <n-button type="primary" size="small" @click="openCreate">{{ t('users.addUser') }}</n-button>
        </div>
      </template>
    </n-data-table>
    <p v-if="users.length" class="legend">{{ t('users.statusLegend') }}</p>
  </n-card>

  <n-drawer v-model:show="drawerOpen" :width="420">
    <n-drawer-content :title="editingName === null ? t('users.addUser') : t('users.editUser', { name: editingName })">
      <n-form>
        <n-form-item v-if="editingName === null">
          <template #label>
            <HelpLabel :label="t('users.name')" :help="t('users.nameHelp')" />
          </template>
          <n-input v-model:value="form.name" placeholder="username" />
        </n-form-item>
        <n-form-item>
          <template #label>
            <HelpLabel
              :label="editingName === null ? t('users.password') : t('users.passwordKeep')"
              :help="t('users.passwordHelp')"
            />
          </template>
          <n-input-group>
            <n-input
              v-model:value="form.password"
              type="text"
              :placeholder="editingName === null ? '' : t('users.passwordKeepPlaceholder')"
            />
            <n-button ghost :title="t('users.regenerate')" @click="regeneratePassword">
              <template #icon><n-icon :component="RefreshOutline" /></template>
            </n-button>
          </n-input-group>
        </n-form-item>
        <n-form-item>
          <template #label>
            <HelpLabel :label="t('users.allowPrivate')" :help="t('users.allowPrivateHelp')" />
          </template>
          <n-switch v-model:value="form.allowPrivateIP" />
        </n-form-item>
        <n-form-item>
          <template #label>
            <HelpLabel :label="t('users.quotas')" :help="t('users.quotasHelp')" />
          </template>
          <n-space vertical style="width: 100%">
            <n-space v-for="(q, i) in quotas" :key="i" align="center">
              <n-input-number v-model:value="q.megabytes" :min="1" :step="1024">
                <template #suffix>MB</template>
              </n-input-number>
              <span>{{ t('users.per') }}</span>
              <n-input-number v-model:value="q.days" :min="1" :max="365">
                <template #suffix>{{ t('users.days') }}</template>
              </n-input-number>
              <n-button size="tiny" quaternary type="error" @click="quotas.splice(i, 1)">✕</n-button>
            </n-space>
            <n-button size="small" dashed @click="addQuota">{{ t('users.addQuota') }}</n-button>
          </n-space>
        </n-form-item>
      </n-form>
      <template #footer>
        <n-space>
          <n-button @click="drawerOpen = false">{{ t('common.cancel') }}</n-button>
          <n-button type="primary" @click="save">{{ t('common.save') }}</n-button>
        </n-space>
      </template>
    </n-drawer-content>
  </n-drawer>

  <ShareModal v-if="shareUser" :username="shareUser" @close="shareUser = null" />
</template>

<style scoped>
.empty {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  padding: 28px 12px;
}
.empty-icon {
  opacity: 0.35;
}
.empty-title {
  font-size: 15px;
  font-weight: 600;
}
.empty-sub {
  font-size: 13px;
  opacity: 0.6;
  margin-bottom: 6px;
  text-align: center;
}
.legend {
  margin: 12px 2px 0;
  font-size: 12px;
  opacity: 0.5;
}
</style>
