<script setup lang="ts">
import { ref, computed, onMounted, watch, nextTick } from 'vue'
import { NModal, NTabs, NTabPane, NInput, NButton, NSpace, NSelect, NAlert, NText, NRadioGroup, NRadioButton, useMessage, useDialog } from 'naive-ui'
import { useI18n } from 'vue-i18n'
import QRCode from 'qrcode'
import { api, ApiError } from '../api/client'
import type { ShareLinks, ShareToken, SubscriptionStatus, SubscriptionToken } from '../api/client'

const props = defineProps<{ username: string }>()
const emit = defineEmits<{ close: [] }>()

const message = useMessage()
const dialog = useDialog()
const { t } = useI18n()
const links = ref<ShareLinks | null>(null)
const error = ref('')
const linkCanvas = ref<HTMLCanvasElement | null>(null)
const subCanvas = ref<HTMLCanvasElement | null>(null)
const advCanvas = ref<HTMLCanvasElement | null>(null)
const activeTab = ref('subscription')

// Permanent subscription; the format targets a client family, not a URI
// scheme: the same token URL serves Clash YAML or mieru link lists.
const subUrl = ref('')
const subCreatedAt = ref(0)
const subLoaded = ref(false)
const subBusy = ref(false)
const subFormat = ref<'clash' | 'mierus' | 'mieru'>('clash')

const subVariantUrl = computed(() => {
  if (!subUrl.value) return ''
  if (subFormat.value === 'clash') return subUrl.value
  return `${subUrl.value}?format=${subFormat.value}`
})
const subFormatHint = computed(() => t(`share.subHint_${subFormat.value}`))

// Expiring shareable link.
const ttlMinutes = ref(60)
const ttlOptions = computed(() => [
  { label: t('share.ttl15m'), value: 15 },
  { label: t('share.ttl1h'), value: 60 },
  { label: t('share.ttl24h'), value: 1440 },
])
const linkUrl = ref('')
const linkExpiresAt = ref(0)
const generating = ref(false)

// Raw formats for manual import.
const advFormat = ref<'mierus' | 'mieru' | 'json'>('mierus')
const advText = computed(() => {
  if (!links.value) return ''
  if (advFormat.value === 'mierus') return links.value.mierusUrls[0] ?? ''
  if (advFormat.value === 'mieru') return links.value.mieruUrl
  return links.value.clientConfigJson
})
const advHint = computed(() => t(`share.advHint_${advFormat.value}`))

async function drawQR(canvas: HTMLCanvasElement | null, text: string) {
  if (!canvas || !text) return
  try {
    await QRCode.toCanvas(canvas, text, { width: 280, margin: 1 })
  } catch {
    /* some links exceed QR capacity; the copyable text still works */
  }
}

async function renderQRs() {
  await nextTick()
  if (activeTab.value === 'subscription') {
    await drawQR(subCanvas.value, subVariantUrl.value)
  } else if (activeTab.value === 'link') {
    await drawQR(linkCanvas.value, linkUrl.value)
  } else if (activeTab.value === 'advanced' && advFormat.value !== 'json') {
    await drawQR(advCanvas.value, advText.value)
  }
}

onMounted(async () => {
  try {
    links.value = await api.get<ShareLinks>(`/api/users/${encodeURIComponent(props.username)}/share`)
  } catch (e) {
    error.value = e instanceof ApiError ? e.message : t('share.failed')
  }
  try {
    const st = await api.get<SubscriptionStatus>(`/api/users/${encodeURIComponent(props.username)}/subscription`)
    if (st.exists && st.url) {
      subUrl.value = st.url
      subCreatedAt.value = st.createdAt ?? 0
    }
  } catch {
    /* subscription state is secondary; the tab shows the create button */
  } finally {
    subLoaded.value = true
  }
  await renderQRs()
})

watch([activeTab, subFormat, advFormat], renderQRs)

async function generateLink() {
  generating.value = true
  try {
    const res = await api.post<ShareToken>(
      `/api/users/${encodeURIComponent(props.username)}/share-token`,
      { ttlMinutes: ttlMinutes.value },
    )
    linkUrl.value = res.url
    linkExpiresAt.value = res.expiresAt
    await renderQRs()
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : t('share.linkFailed'))
  } finally {
    generating.value = false
  }
}

function expiryText(): string {
  if (!linkExpiresAt.value) return ''
  return t('share.expires', { date: new Date(linkExpiresAt.value * 1000).toLocaleString() })
}

function copy(text: string) {
  navigator.clipboard.writeText(text).then(
    () => message.success(t('common.copied')),
    () => message.error(t('common.copyFailed')),
  )
}

function downloadConfig() {
  if (!links.value) return
  const blob = new Blob([links.value.clientConfigJson], { type: 'application/json' })
  const a = document.createElement('a')
  a.href = URL.createObjectURL(blob)
  a.download = `mieru-${props.username}.json`
  a.click()
  URL.revokeObjectURL(a.href)
}

async function createSubscription() {
  subBusy.value = true
  try {
    const res = await api.post<SubscriptionToken>(
      `/api/users/${encodeURIComponent(props.username)}/subscription`, {},
    )
    subUrl.value = res.url
    subCreatedAt.value = res.createdAt
    await renderQRs()
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : t('share.subFailed'))
  } finally {
    subBusy.value = false
  }
}

function rotateSubscription() {
  dialog.warning({
    title: t('share.subRotate'),
    content: t('share.subRotateConfirm'),
    positiveText: t('share.subRotate'),
    negativeText: t('common.cancel'),
    onPositiveClick: createSubscription,
  })
}

function revokeSubscription() {
  dialog.warning({
    title: t('share.subRevoke'),
    content: t('share.subRevokeConfirm'),
    positiveText: t('share.subRevoke'),
    negativeText: t('common.cancel'),
    onPositiveClick: async () => {
      try {
        await api.del(`/api/users/${encodeURIComponent(props.username)}/subscription`)
        subUrl.value = ''
        subCreatedAt.value = 0
      } catch (e) {
        message.error(e instanceof ApiError ? e.message : t('share.subFailed'))
      }
    },
  })
}

function subActiveSince(): string {
  if (!subCreatedAt.value) return ''
  return t('share.subActiveSince', { date: new Date(subCreatedAt.value * 1000).toLocaleDateString() })
}
</script>

<template>
  <n-modal :show="true" preset="card" :title="t('share.title', { name: username })" style="width: 520px" @close="emit('close')" @mask-click="emit('close')">
    <n-alert v-if="error" type="warning" :show-icon="false">{{ error }}</n-alert>
    <n-tabs v-else v-model:value="activeTab">
      <!-- Permanent subscription: the primary way to hand out access. -->
      <n-tab-pane name="subscription" :tab="t('share.subTab')">
        <n-space vertical align="center">
          <template v-if="!subUrl">
            <n-text depth="3" style="font-size: 13px; text-align: center">
              {{ t('share.subIntro') }}
            </n-text>
            <n-button type="primary" :loading="subBusy" :disabled="!subLoaded" @click="createSubscription">
              {{ t('share.subCreate') }}
            </n-button>
          </template>
          <template v-else>
            <n-radio-group v-model:value="subFormat" size="small">
              <n-radio-button value="clash">{{ t('share.subClientClash') }}</n-radio-button>
              <n-radio-button value="mierus">{{ t('share.subClientNeko') }}</n-radio-button>
              <n-radio-button value="mieru">{{ t('share.subClientMieru') }}</n-radio-button>
            </n-radio-group>
            <n-text depth="3" style="font-size: 12px; text-align: center; max-width: 420px">
              {{ subFormatHint }}
            </n-text>
            <canvas ref="subCanvas" />
            <n-input :value="subVariantUrl" readonly size="small" style="width: 460px; max-width: 100%" />
            <n-button size="small" type="primary" @click="copy(subVariantUrl)">{{ t('share.copyLink') }}</n-button>
            <n-text depth="3" style="font-size: 12px; text-align: center; max-width: 440px">
              {{ t('share.subUpdates') }}
            </n-text>
            <n-space align="center" style="margin-top: 8px">
              <n-text depth="3" style="font-size: 12px">{{ subActiveSince() }}</n-text>
              <n-button size="tiny" quaternary @click="rotateSubscription">{{ t('share.subRotate') }}</n-button>
              <n-button size="tiny" quaternary type="error" @click="revokeSubscription">{{ t('share.subRevoke') }}</n-button>
            </n-space>
          </template>
        </n-space>
      </n-tab-pane>

      <!-- One-off expiring page with the credentials. -->
      <n-tab-pane name="link" :tab="t('share.linkTab')">
        <n-space vertical align="center">
          <n-text depth="3" style="font-size: 12px; text-align: center">
            {{ t('share.linkText') }}
          </n-text>
          <n-space align="center">
            <n-select v-model:value="ttlMinutes" :options="ttlOptions" style="width: 140px" />
            <n-button type="primary" :loading="generating" @click="generateLink">{{ t('share.generateLink') }}</n-button>
          </n-space>
          <template v-if="linkUrl">
            <canvas ref="linkCanvas" />
            <n-input :value="linkUrl" readonly size="small" style="width: 460px; max-width: 100%" />
            <n-space align="center">
              <n-button size="small" @click="copy(linkUrl)">{{ t('share.copyLink') }}</n-button>
              <n-text depth="3" style="font-size: 12px">{{ expiryText() }}</n-text>
            </n-space>
          </template>
        </n-space>
      </n-tab-pane>

      <!-- Raw formats for manual import into any client. -->
      <n-tab-pane name="advanced" :tab="t('share.advTab')" :disabled="!links">
        <n-space vertical align="center">
          <n-radio-group v-model:value="advFormat" size="small">
            <n-radio-button value="mierus">mierus://</n-radio-button>
            <n-radio-button value="mieru">mieru://</n-radio-button>
            <n-radio-button value="json">JSON</n-radio-button>
          </n-radio-group>
          <n-text depth="3" style="font-size: 12px; text-align: center; max-width: 420px">
            {{ advHint }}
          </n-text>
          <canvas v-show="advFormat !== 'json'" ref="advCanvas" />
          <n-input :value="advText" type="textarea" readonly size="small" style="width: 460px; max-width: 100%" :autosize="{ minRows: 2, maxRows: advFormat === 'json' ? 10 : 4 }" />
          <n-space>
            <n-button size="small" @click="copy(advText)">{{ t('common.copy') }}</n-button>
            <n-button v-if="advFormat === 'json'" size="small" @click="downloadConfig">{{ t('share.downloadJson') }}</n-button>
          </n-space>
        </n-space>
      </n-tab-pane>
    </n-tabs>
  </n-modal>
</template>
