<script setup lang="ts">
import { h, ref, onMounted } from 'vue'
import {
  NButton, NCard, NDataTable, NDrawer, NDrawerContent, NForm, NFormItem, NInput,
  NInputNumber, NSpace, NSwitch, NTag, NPopconfirm, useMessage,
} from 'naive-ui'
import type { DataTableColumns } from 'naive-ui'
import { api, ApiError } from '../api/client'
import type { UserInfo, Quota } from '../api/client'
import ShareModal from '../components/ShareModal.vue'

const message = useMessage()
const users = ref<UserInfo[]>([])
const loading = ref(false)

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

const columns: DataTableColumns<UserInfo> = [
  { title: 'Name', key: 'name' },
  {
    title: 'Quotas',
    key: 'quotas',
    render: (u) =>
      u.quotas.length
        ? u.quotas.map((q) => h(NTag, { size: 'small', style: 'margin-right:4px' }, { default: () => `${q.megabytes} MB / ${q.days}d` }))
        : '—',
  },
  { title: 'Traffic', key: 'traffic', render: trafficOf },
  {
    title: 'Share',
    key: 'share',
    render: (u) =>
      h(
        NButton,
        { size: 'small', disabled: !u.hasSecret, onClick: () => (shareUser.value = u.name) },
        { default: () => (u.hasSecret ? 'Share' : 'No password stored') },
      ),
  },
  {
    title: 'Actions',
    key: 'actions',
    render: (u) =>
      h(NSpace, null, {
        default: () => [
          h(NButton, { size: 'small', onClick: () => openEdit(u) }, { default: () => 'Edit' }),
          h(
            NPopconfirm,
            { onPositiveClick: () => removeUser(u.name) },
            {
              trigger: () => h(NButton, { size: 'small', type: 'error', quaternary: true }, { default: () => 'Delete' }),
              default: () => `Delete user ${u.name}?`,
            },
          ),
        ],
      }),
  },
]

async function load() {
  loading.value = true
  try {
    users.value = await api.get<UserInfo[]>('/api/users')
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : 'Failed to load users')
  } finally {
    loading.value = false
  }
}

function openCreate() {
  editingName.value = null
  form.value = { name: '', password: '', allowPrivateIP: false }
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
      message.success('User created')
    } else {
      await api.put(`/api/users/${encodeURIComponent(editingName.value)}`, body)
      message.success('User updated')
    }
    drawerOpen.value = false
    await load()
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : 'Save failed')
  }
}

async function removeUser(name: string) {
  try {
    await api.del(`/api/users/${encodeURIComponent(name)}`)
    message.success('User deleted')
    await load()
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : 'Delete failed')
  }
}

onMounted(load)
</script>

<template>
  <n-card title="Users">
    <template #header-extra>
      <n-button type="primary" @click="openCreate">Add user</n-button>
    </template>
    <n-data-table :columns="columns" :data="users" :loading="loading" :row-key="(u: UserInfo) => u.name" />
  </n-card>

  <n-drawer v-model:show="drawerOpen" :width="420">
    <n-drawer-content :title="editingName === null ? 'Add user' : `Edit ${editingName}`">
      <n-form>
        <n-form-item v-if="editingName === null" label="Name">
          <n-input v-model:value="form.name" placeholder="username" />
        </n-form-item>
        <n-form-item :label="editingName === null ? 'Password' : 'New password (leave empty to keep)'">
          <n-input v-model:value="form.password" type="password" show-password-on="click" />
        </n-form-item>
        <n-form-item label="Allow connections to private IPs">
          <n-switch v-model:value="form.allowPrivateIP" />
        </n-form-item>
        <n-form-item label="Quotas (traffic cap per rolling window)">
          <n-space vertical style="width: 100%">
            <n-space v-for="(q, i) in quotas" :key="i" align="center">
              <n-input-number v-model:value="q.megabytes" :min="1" :step="1024">
                <template #suffix>MB</template>
              </n-input-number>
              <span>per</span>
              <n-input-number v-model:value="q.days" :min="1" :max="365">
                <template #suffix>days</template>
              </n-input-number>
              <n-button size="tiny" quaternary type="error" @click="quotas.splice(i, 1)">✕</n-button>
            </n-space>
            <n-button size="small" dashed @click="addQuota">Add quota</n-button>
          </n-space>
        </n-form-item>
      </n-form>
      <template #footer>
        <n-space>
          <n-button @click="drawerOpen = false">Cancel</n-button>
          <n-button type="primary" @click="save">Save</n-button>
        </n-space>
      </template>
    </n-drawer-content>
  </n-drawer>

  <ShareModal v-if="shareUser" :username="shareUser" @close="shareUser = null" />
</template>
