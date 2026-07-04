<script setup lang="ts">
import { h, ref, computed, onMounted } from 'vue'
import {
  NCard, NButton, NSpace, NInput, NInputNumber, NSelect, NAlert, NTag, NText,
  NDataTable, NDrawer, NDrawerContent, NForm, NFormItem, NDynamicInput, useMessage,
} from 'naive-ui'
import type { DataTableColumns } from 'naive-ui'
import { api, ApiError } from '../api/client'
import type { EgressConfig, EgressProxy, EgressRule, GeoDataset, GeoCategory, GeoipState, Peer } from '../api/client'

const message = useMessage()
const loading = ref(false)
const proxies = ref<EgressProxy[]>([])
const rules = ref<EgressRule[]>([])
const peers = ref<Peer[]>([])

const actionOptions = ['PROXY', 'DIRECT', 'REJECT'].map((a) => ({ label: a, value: a }))
// Rules can route through user-defined proxies and chained upstream peers.
const proxyOptions = computed(() => [
  ...proxies.value.map((p) => ({ label: p.name, value: p.name })),
  ...peers.value.map((p) => ({ label: `${p.name} (peer)`, value: p.name })),
])

// --- geoip datasets (xray geoip.dat format) ---
const datasets = ref<GeoDataset[]>([])
const categories = ref<GeoCategory[]>([])
const geoBusy = ref(false)
const dsName = ref('')
const dsUrl = ref('')
const presets = [
  { name: 'geoip', url: 'https://github.com/Loyalsoldier/geoip/releases/latest/download/geoip.dat', label: 'Loyalsoldier (countries)' },
  { name: 'ru-blocked', url: 'https://github.com/runetfreedom/russia-blocked-geoip/releases/latest/download/ru-blocked.dat', label: 'RU blocked (Roskomnadzor)' },
]
const catOptions = computed(() =>
  categories.value.map((c) => ({ label: `${c.code} (${c.cidrs})`, value: c.code })),
)

function usePreset(p: { name: string; url: string }) {
  dsName.value = p.name
  dsUrl.value = p.url
}

async function loadGeo() {
  try {
    const res = await api.get<GeoipState>('/api/geoip')
    datasets.value = res.datasets
    categories.value = res.categories
  } catch {
    /* geoip is optional */
  }
}
async function addDataset() {
  if (!dsName.value.trim() || !dsUrl.value.trim()) {
    message.error('Name and URL are required')
    return
  }
  geoBusy.value = true
  try {
    const res = await api.post<GeoipState>('/api/geoip/datasets', { name: dsName.value.trim(), url: dsUrl.value.trim() })
    datasets.value = res.datasets
    categories.value = res.categories
    dsName.value = ''
    dsUrl.value = ''
    message.success('Dataset added')
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : 'Download failed')
  } finally {
    geoBusy.value = false
  }
}
async function deleteDataset(name: string) {
  try {
    const res = await api.del<GeoipState>(`/api/geoip/datasets/${encodeURIComponent(name)}`)
    datasets.value = res.datasets
    categories.value = res.categories
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : 'Delete failed')
  }
}
function fmtMB(b: number): string {
  return `${(b / 1024 / 1024).toFixed(1)} MB`
}

// --- proxy editor drawer ---
const proxyDrawer = ref(false)
const proxyIndex = ref<number | null>(null)
const proxyForm = ref<EgressProxy>({ name: '', host: '', port: 1080, username: '', password: '' })

function openProxy(i: number | null) {
  proxyIndex.value = i
  proxyForm.value = i === null
    ? { name: '', host: '', port: 1080, username: '', password: '' }
    : { ...proxies.value[i] }
  proxyDrawer.value = true
}
function saveProxy() {
  if (!proxyForm.value.name || !proxyForm.value.host) {
    message.error('Name and host are required')
    return
  }
  if (proxyIndex.value === null) proxies.value.push({ ...proxyForm.value })
  else proxies.value[proxyIndex.value] = { ...proxyForm.value }
  proxyDrawer.value = false
}

const proxyColumns: DataTableColumns<EgressProxy> = [
  { title: 'Name', key: 'name' },
  { title: 'Address', key: 'addr', render: (p) => `${p.host}:${p.port}` },
  { title: 'Auth', key: 'auth', render: (p) => (p.username ? 'yes' : '—') },
  {
    title: 'Actions',
    key: 'actions',
    render: (_p, i) =>
      h(NSpace, null, {
        default: () => [
          h(NButton, { size: 'small', onClick: () => openProxy(i) }, { default: () => 'Edit' }),
          h(NButton, { size: 'small', type: 'error', quaternary: true, onClick: () => proxies.value.splice(i, 1) }, { default: () => 'Delete' }),
        ],
      }),
  },
]

function addRule() {
  rules.value.push({ domains: [], cidrs: [], geo: [], action: 'PROXY', proxies: [] })
}

async function load() {
  loading.value = true
  try {
    const cfg = await api.get<EgressConfig>('/api/config/egress')
    proxies.value = cfg.proxies
    rules.value = cfg.rules.map((r) => ({ ...r, geo: r.geo ?? [] }))
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : 'Failed to load egress config')
  } finally {
    loading.value = false
  }
}

async function save() {
  try {
    await api.put('/api/config/egress', { proxies: proxies.value, rules: rules.value })
    message.success('Outbound config applied')
    await load()
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : 'Save failed')
  }
}

async function loadPeers() {
  try {
    peers.value = await api.get<Peer[]>('/api/chain')
  } catch {
    /* chain is optional */
  }
}

onMounted(() => {
  load()
  loadGeo()
  loadPeers()
})
</script>

<template>
  <h2 class="page-title">Outbounds</h2>
  <n-space vertical :size="16">
    <n-card title="Outbound proxies" :loading="loading">
      <template #header-extra>
        <n-button type="primary" @click="openProxy(null)">Add proxy</n-button>
      </template>
      <n-alert type="info" :show-icon="false" style="margin-bottom: 12px">
        SOCKS5 hops for chaining. Reference them by name in the rules below.
      </n-alert>
      <n-data-table :columns="proxyColumns" :data="proxies" :row-key="(p: EgressProxy) => p.name" size="small" />
    </n-card>

    <n-card title="GeoIP datasets (optional)">
      <n-alert type="info" :show-icon="false" style="margin-bottom: 12px">
        Add <n-tag size="small">geoip.dat</n-tag> datasets (the same format xray/3x-ui use), then match
        their categories in the rules below. Common use: keep your country DIRECT and send everything
        else through a proxy, or send only RU-blocked ranges through a proxy. You can also mount your own
        <n-tag size="small">.dat</n-tag> files into the datasets directory (see docs). Large categories
        expand to many CIDRs, which enlarges the server config.
      </n-alert>
      <n-space style="margin-bottom: 10px">
        <n-button v-for="p in presets" :key="p.name" size="small" tertiary @click="usePreset(p)">
          {{ p.label }}
        </n-button>
      </n-space>
      <n-space align="center" style="margin-bottom: 12px">
        <n-input v-model:value="dsName" placeholder="name" style="width: 130px" />
        <n-input v-model:value="dsUrl" placeholder="geoip.dat URL" style="width: 420px" />
        <n-button type="primary" :loading="geoBusy" @click="addDataset">Add</n-button>
      </n-space>
      <n-space>
        <n-tag v-for="d in datasets" :key="d.name" closable @close="deleteDataset(d.name)">
          {{ d.name }} · {{ fmtMB(d.bytes) }}
        </n-tag>
        <n-text v-if="!datasets.length" depth="3">No datasets yet.</n-text>
      </n-space>
      <div v-if="categories.length" class="cats">{{ categories.length }} categories available</div>
    </n-card>

    <n-card title="Routing rules">
      <template #header-extra>
        <n-button @click="addRule">Add rule</n-button>
      </template>
      <n-alert type="info" :show-icon="false" style="margin-bottom: 12px">
        Evaluated top to bottom; the first match wins. If nothing matches, traffic goes out directly.
        Use <n-tag size="small">*</n-tag> to match everything.
      </n-alert>
      <n-space vertical>
        <n-card v-for="(rule, i) in rules" :key="i" size="small" embedded>
          <n-space vertical>
            <n-space align="center" wrap>
              <span class="lbl">#{{ i + 1 }}</span>
              <n-select v-model:value="rule.action" :options="actionOptions" style="width: 130px" />
              <n-select
                v-if="rule.action === 'PROXY'"
                v-model:value="rule.proxies"
                multiple
                :options="proxyOptions"
                placeholder="via proxies"
                style="min-width: 200px"
              />
              <n-button size="tiny" quaternary type="error" @click="rules.splice(i, 1)">Remove</n-button>
            </n-space>
            <n-select
              v-model:value="rule.geo"
              multiple
              filterable
              :options="catOptions"
              placeholder="match geo categories (e.g. ru, ru-blocked)"
            />
            <n-space>
              <n-dynamic-input v-model:value="rule.domains" placeholder="domain (e.g. *.example.com or *)" :min="0">
                <template #default="{ value, index }">
                  <n-input :value="value" @update:value="(v: string) => (rule.domains[index] = v)" placeholder="domain" />
                </template>
              </n-dynamic-input>
              <n-dynamic-input v-model:value="rule.cidrs" placeholder="CIDR (e.g. 10.0.0.0/8 or *)" :min="0">
                <template #default="{ value, index }">
                  <n-input :value="value" @update:value="(v: string) => (rule.cidrs[index] = v)" placeholder="CIDR" />
                </template>
              </n-dynamic-input>
            </n-space>
          </n-space>
        </n-card>
      </n-space>
      <n-button type="primary" style="margin-top: 16px" @click="save">Apply</n-button>
    </n-card>

    <n-drawer v-model:show="proxyDrawer" :width="400">
      <n-drawer-content :title="proxyIndex === null ? 'Add proxy' : 'Edit proxy'">
        <n-form>
          <n-form-item label="Name (referenced by rules)">
            <n-input v-model:value="proxyForm.name" placeholder="upstream1" />
          </n-form-item>
          <n-form-item label="Host">
            <n-input v-model:value="proxyForm.host" placeholder="1.2.3.4 or proxy.example.com" />
          </n-form-item>
          <n-form-item label="Port">
            <n-input-number v-model:value="proxyForm.port" :min="1" :max="65535" />
          </n-form-item>
          <n-form-item label="Username (optional)">
            <n-input v-model:value="proxyForm.username" />
          </n-form-item>
          <n-form-item label="Password (optional)">
            <n-input v-model:value="proxyForm.password" type="password" show-password-on="click" />
          </n-form-item>
        </n-form>
        <template #footer>
          <n-space>
            <n-button @click="proxyDrawer = false">Cancel</n-button>
            <n-button type="primary" @click="saveProxy">Save</n-button>
          </n-space>
        </template>
      </n-drawer-content>
    </n-drawer>
  </n-space>
</template>

<style scoped>
.lbl {
  font-weight: 600;
  opacity: 0.6;
}
.cats {
  margin-top: 10px;
  font-size: 12px;
  opacity: 0.6;
}
</style>
