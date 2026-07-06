<script setup lang="ts">
import { h, ref, computed, onMounted } from 'vue'
import {
  NCard, NButton, NSpace, NInput, NInputNumber, NSelect, NAlert, NTag, NText,
  NDataTable, NDrawer, NDrawerContent, NForm, NFormItem, NDynamicInput,
  NRadioGroup, NRadioButton, useMessage,
} from 'naive-ui'
import type { DataTableColumns } from 'naive-ui'
import { useI18n } from 'vue-i18n'
import { api, ApiError } from '../api/client'
import type { EgressConfig, EgressProxy, EgressRule, GeoDataset, GeoCategory, GeoSiteCategory, GeoipState, Peer } from '../api/client'
import HelpLabel from '../components/HelpLabel.vue'

const message = useMessage()
const { t } = useI18n()
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

// --- geoip / geosite datasets (xray geoip.dat / geosite.dat format) ---
const datasets = ref<GeoDataset[]>([])
const categories = ref<GeoCategory[]>([])
const siteDatasets = ref<GeoDataset[]>([])
const siteCategories = ref<GeoSiteCategory[]>([])
const geoBusy = ref(false)
const dsKind = ref<'geoip' | 'geosite'>('geoip')
const dsName = ref('')
const dsUrl = ref('')

type Preset = { name: string; url: string; label: string; kind: 'geoip' | 'geosite' }
const presets: Preset[] = [
  // runetfreedom's combined dats carry many RU-relevant categories in one
  // file, superseding the older single-category ru-blocked lists.
  { name: 'runet', url: 'https://github.com/runetfreedom/russia-v2ray-rules-dat/releases/latest/download/geoip.dat', label: 'RuNet Freedom', kind: 'geoip' },
  { name: 'geoip', url: 'https://github.com/Loyalsoldier/geoip/releases/latest/download/geoip.dat', label: 'Loyalsoldier (countries)', kind: 'geoip' },
  { name: 'runet', url: 'https://github.com/runetfreedom/russia-v2ray-rules-dat/releases/latest/download/geosite.dat', label: 'RuNet Freedom', kind: 'geosite' },
  { name: 'v2fly', url: 'https://github.com/v2fly/domain-list-community/releases/latest/download/dlc.dat', label: 'v2fly (services)', kind: 'geosite' },
]
const shownPresets = computed(() => presets.filter((p) => p.kind === dsKind.value))

// geoip categories expand to CIDRs, geosite to domains; both are offered as
// rule match options, tagged so the two never get confused.
const catOptions = computed(() => [
  ...categories.value.map((c) => ({ label: `${c.code} · ${c.cidrs} IP`, value: `ip:${c.code}` })),
  ...siteCategories.value.map((c) => ({ label: `${c.code} · ${c.domains} ${t('outbounds.domainsShort')}`, value: `site:${c.code}` })),
])

function usePreset(p: Preset) {
  dsKind.value = p.kind
  dsName.value = p.name
  dsUrl.value = p.url
}

function applyGeoState(res: GeoipState) {
  datasets.value = res.datasets
  categories.value = res.categories
  siteDatasets.value = res.siteDatasets ?? []
  siteCategories.value = res.siteCategories ?? []
}

async function loadGeo() {
  try {
    applyGeoState(await api.get<GeoipState>('/api/geoip'))
  } catch {
    /* geoip is optional */
  }
}
async function addDataset() {
  if (!dsName.value.trim() || !dsUrl.value.trim()) {
    message.error(t('outbounds.nameUrlRequired'))
    return
  }
  geoBusy.value = true
  try {
    applyGeoState(await api.post<GeoipState>('/api/geoip/datasets', {
      name: dsName.value.trim(), url: dsUrl.value.trim(), kind: dsKind.value,
    }))
    dsName.value = ''
    dsUrl.value = ''
    message.success(t('outbounds.datasetAdded'))
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : t('outbounds.downloadFailed'))
  } finally {
    geoBusy.value = false
  }
}
async function deleteDataset(name: string, kind: 'geoip' | 'geosite') {
  try {
    const q = kind === 'geosite' ? '?kind=geosite' : ''
    applyGeoState(await api.del<GeoipState>(`/api/geoip/datasets/${encodeURIComponent(name)}${q}`))
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : t('common.deleteFailed'))
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
    message.error(t('outbounds.nameHostRequired'))
    return
  }
  if (proxyIndex.value === null) proxies.value.push({ ...proxyForm.value })
  else proxies.value[proxyIndex.value] = { ...proxyForm.value }
  proxyDrawer.value = false
}

const proxyColumns = computed<DataTableColumns<EgressProxy>>(() => [
  { title: t('common.name'), key: 'name' },
  { title: t('outbounds.colAddress'), key: 'addr', render: (p) => `${p.host}:${p.port}` },
  { title: t('outbounds.colAuth'), key: 'auth', render: (p) => (p.username ? t('outbounds.authYes') : '—') },
  {
    title: t('common.actions'),
    key: 'actions',
    render: (_p, i) =>
      h(NSpace, null, {
        default: () => [
          h(NButton, { size: 'small', onClick: () => openProxy(i) }, { default: () => t('common.edit') }),
          h(NButton, { size: 'small', type: 'error', quaternary: true, onClick: () => proxies.value.splice(i, 1) }, { default: () => t('common.delete') }),
        ],
      }),
  },
])

function addRule() {
  rules.value.push({ domains: [], cidrs: [], geo: [], sites: [], action: 'PROXY', proxies: [] })
}

// The rule editor offers geoip and geosite categories in one select, tagged
// "ip:" / "site:"; on the wire they stay in separate geo/sites arrays.
function ruleCats(rule: EgressRule): string[] {
  return [...(rule.geo ?? []).map((c) => `ip:${c}`), ...(rule.sites ?? []).map((c) => `site:${c}`)]
}
function setRuleCats(rule: EgressRule, vals: string[]) {
  rule.geo = vals.filter((v) => v.startsWith('ip:')).map((v) => v.slice(3))
  rule.sites = vals.filter((v) => v.startsWith('site:')).map((v) => v.slice(5))
}

async function load() {
  loading.value = true
  try {
    const cfg = await api.get<EgressConfig>('/api/config/egress')
    proxies.value = cfg.proxies
    rules.value = cfg.rules.map((r) => ({ ...r, geo: r.geo ?? [], sites: r.sites ?? [] }))
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : t('outbounds.loadFailed'))
  } finally {
    loading.value = false
  }
}

async function save() {
  try {
    await api.put('/api/config/egress', { proxies: proxies.value, rules: rules.value })
    message.success(t('outbounds.applied'))
    await load()
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : t('common.saveFailed'))
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
  <h2 class="page-title">{{ t('outbounds.title') }}</h2>
  <n-space vertical :size="16">
    <n-card :title="t('outbounds.proxiesTitle')" :loading="loading">
      <template #header-extra>
        <n-button type="primary" @click="openProxy(null)">{{ t('outbounds.addProxy') }}</n-button>
      </template>
      <n-alert type="info" :show-icon="false" style="margin-bottom: 12px">
        {{ t('outbounds.proxiesAlert') }}
      </n-alert>
      <n-data-table :columns="proxyColumns" :data="proxies" :row-key="(p: EgressProxy) => p.name" size="small" :scroll-x="520" />
    </n-card>

    <n-card :title="t('outbounds.geoTitle')">
      <n-alert type="info" :show-icon="false" style="margin-bottom: 12px">
        {{ t('outbounds.geoAlert') }}
      </n-alert>
      <n-radio-group v-model:value="dsKind" size="small" style="margin-bottom: 10px">
        <n-radio-button value="geoip">{{ t('outbounds.kindGeoip') }}</n-radio-button>
        <n-radio-button value="geosite">{{ t('outbounds.kindGeosite') }}</n-radio-button>
      </n-radio-group>
      <n-text depth="3" style="display: block; font-size: 12px; margin-bottom: 10px">
        {{ dsKind === 'geoip' ? t('outbounds.kindGeoipHint') : t('outbounds.kindGeositeHint') }}
      </n-text>
      <n-space style="margin-bottom: 10px">
        <n-button v-for="p in shownPresets" :key="p.kind + p.name" size="small" tertiary @click="usePreset(p)">
          {{ p.label }}
        </n-button>
      </n-space>
      <n-space align="center" style="margin-bottom: 12px">
        <n-input v-model:value="dsName" :placeholder="t('outbounds.dsNamePlaceholder')" style="width: 130px" />
        <n-input v-model:value="dsUrl" :placeholder="dsKind === 'geoip' ? 'geoip.dat URL' : 'geosite.dat URL'" style="width: 420px" />
        <n-button type="primary" :loading="geoBusy" @click="addDataset">{{ t('common.add') }}</n-button>
      </n-space>
      <n-space vertical :size="8">
        <n-space align="center" :size="6">
          <span class="ds-lbl">{{ t('outbounds.kindGeoip') }}:</span>
          <n-tag v-for="d in datasets" :key="d.name" closable type="warning" @close="deleteDataset(d.name, 'geoip')">
            {{ d.name }} · {{ fmtMB(d.bytes) }}
          </n-tag>
          <n-text v-if="!datasets.length" depth="3" style="font-size: 12px">{{ t('outbounds.noDatasets') }}</n-text>
          <n-text v-if="categories.length" depth="3" style="font-size: 12px">· {{ t('outbounds.categoriesAvailable', { n: categories.length }) }}</n-text>
        </n-space>
        <n-space align="center" :size="6">
          <span class="ds-lbl">{{ t('outbounds.kindGeosite') }}:</span>
          <n-tag v-for="d in siteDatasets" :key="d.name" closable type="info" @close="deleteDataset(d.name, 'geosite')">
            {{ d.name }} · {{ fmtMB(d.bytes) }}
          </n-tag>
          <n-text v-if="!siteDatasets.length" depth="3" style="font-size: 12px">{{ t('outbounds.noDatasets') }}</n-text>
          <n-text v-if="siteCategories.length" depth="3" style="font-size: 12px">· {{ t('outbounds.siteCategoriesAvailable', { n: siteCategories.length }) }}</n-text>
        </n-space>
      </n-space>
    </n-card>

    <n-card :title="t('outbounds.rulesTitle')">
      <template #header-extra>
        <n-button @click="addRule">{{ t('outbounds.addRule') }}</n-button>
      </template>
      <n-alert type="info" :show-icon="false" style="margin-bottom: 12px">
        {{ t('outbounds.rulesAlert') }}
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
                :placeholder="t('outbounds.viaProxies')"
                style="min-width: 200px"
              />
              <n-button size="tiny" quaternary type="error" @click="rules.splice(i, 1)">{{ t('common.remove') }}</n-button>
            </n-space>
            <div class="field">
              <span class="field-lbl">{{ t('outbounds.matchCategories') }}</span>
              <n-select
                :value="ruleCats(rule)"
                @update:value="(v: string[]) => setRuleCats(rule, v)"
                multiple
                filterable
                :options="catOptions"
                :placeholder="t('outbounds.matchGeoPlaceholder')"
              />
            </div>
            <n-space :size="16">
              <div class="field" style="flex: 1; min-width: 220px">
                <span class="field-lbl">{{ t('outbounds.matchDomains') }}</span>
                <n-dynamic-input v-model:value="rule.domains" :min="0">
                  <template #create-button-default>{{ t('outbounds.addDomain') }}</template>
                  <template #default="{ value, index }">
                    <n-input :value="value" @update:value="(v: string) => (rule.domains[index] = v)" :placeholder="t('outbounds.domainPlaceholder')" />
                  </template>
                </n-dynamic-input>
              </div>
              <div class="field" style="flex: 1; min-width: 220px">
                <span class="field-lbl">{{ t('outbounds.matchCidrs') }}</span>
                <n-dynamic-input v-model:value="rule.cidrs" :min="0">
                  <template #create-button-default>{{ t('outbounds.addCidr') }}</template>
                  <template #default="{ value, index }">
                    <n-input :value="value" @update:value="(v: string) => (rule.cidrs[index] = v)" :placeholder="t('outbounds.cidrPlaceholder')" />
                  </template>
                </n-dynamic-input>
              </div>
            </n-space>
          </n-space>
        </n-card>
      </n-space>
      <n-button type="primary" style="margin-top: 16px" @click="save">{{ t('common.apply') }}</n-button>
    </n-card>

    <n-drawer v-model:show="proxyDrawer" :width="400">
      <n-drawer-content :title="proxyIndex === null ? t('outbounds.addProxy') : t('outbounds.editProxy')">
        <n-form>
          <n-form-item>
            <template #label>
              <HelpLabel :label="t('outbounds.proxyName')" :help="t('outbounds.proxyNameHelp')" />
            </template>
            <n-input v-model:value="proxyForm.name" placeholder="upstream1" />
          </n-form-item>
          <n-form-item>
            <template #label>
              <HelpLabel :label="t('outbounds.host')" :help="t('outbounds.hostHelp')" />
            </template>
            <n-input v-model:value="proxyForm.host" placeholder="1.2.3.4 / proxy.example.com" />
          </n-form-item>
          <n-form-item>
            <template #label>
              <HelpLabel :label="t('common.port')" :help="t('outbounds.portHelp')" />
            </template>
            <n-input-number v-model:value="proxyForm.port" :min="1" :max="65535" />
          </n-form-item>
          <n-form-item>
            <template #label>
              <HelpLabel :label="t('outbounds.usernameOpt')" :help="t('outbounds.socksAuthHelp')" />
            </template>
            <n-input v-model:value="proxyForm.username" />
          </n-form-item>
          <n-form-item>
            <template #label>
              <HelpLabel :label="t('outbounds.passwordOpt')" :help="t('outbounds.socksAuthHelp')" />
            </template>
            <n-input v-model:value="proxyForm.password" type="password" show-password-on="click" />
          </n-form-item>
        </n-form>
        <template #footer>
          <n-space>
            <n-button @click="proxyDrawer = false">{{ t('common.cancel') }}</n-button>
            <n-button type="primary" @click="saveProxy">{{ t('common.save') }}</n-button>
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
.ds-lbl {
  font-size: 12px;
  font-weight: 600;
  opacity: 0.55;
  min-width: 64px;
}
.field-lbl {
  display: block;
  font-size: 12px;
  opacity: 0.55;
  margin-bottom: 4px;
}
</style>
