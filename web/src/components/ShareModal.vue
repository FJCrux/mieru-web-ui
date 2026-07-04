<script setup lang="ts">
import { ref, onMounted, watch, nextTick } from 'vue'
import { NModal, NTabs, NTabPane, NInput, NButton, NSpace, NSelect, NAlert, NText, useMessage } from 'naive-ui'
import QRCode from 'qrcode'
import { api, ApiError } from '../api/client'
import type { ShareLinks, ShareToken } from '../api/client'

const props = defineProps<{ username: string }>()
const emit = defineEmits<{ close: [] }>()

const message = useMessage()
const links = ref<ShareLinks | null>(null)
const error = ref('')
const mierusCanvas = ref<HTMLCanvasElement | null>(null)
const mieruCanvas = ref<HTMLCanvasElement | null>(null)
const linkCanvas = ref<HTMLCanvasElement | null>(null)
const activeTab = ref('link')

// Expiring shareable link.
const ttlMinutes = ref(60)
const ttlOptions = [
  { label: '15 minutes', value: 15 },
  { label: '1 hour', value: 60 },
  { label: '24 hours', value: 1440 },
]
const linkUrl = ref('')
const linkExpiresAt = ref(0)
const generating = ref(false)

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
  if (activeTab.value === 'link') {
    await drawQR(linkCanvas.value, linkUrl.value)
    return
  }
  if (!links.value) return
  if (links.value.mierusUrls.length) await drawQR(mierusCanvas.value, links.value.mierusUrls[0])
  await drawQR(mieruCanvas.value, links.value.mieruUrl)
}

onMounted(async () => {
  try {
    links.value = await api.get<ShareLinks>(`/api/users/${encodeURIComponent(props.username)}/share`)
    await renderQRs()
  } catch (e) {
    error.value = e instanceof ApiError ? e.message : 'Failed to build share links'
  }
})

watch(activeTab, renderQRs)

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
    message.error(e instanceof ApiError ? e.message : 'Failed to create link')
  } finally {
    generating.value = false
  }
}

function expiryText(): string {
  if (!linkExpiresAt.value) return ''
  return `Expires ${new Date(linkExpiresAt.value * 1000).toLocaleString()}`
}

function copy(text: string) {
  navigator.clipboard.writeText(text).then(
    () => message.success('Copied'),
    () => message.error('Copy failed'),
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
</script>

<template>
  <n-modal :show="true" preset="card" :title="`Share: ${username}`" style="width: 480px" @close="emit('close')" @mask-click="emit('close')">
    <n-alert v-if="error" type="warning" :show-icon="false">{{ error }}</n-alert>
    <n-tabs v-else v-model:value="activeTab">
      <n-tab-pane name="link" tab="Link">
        <n-space vertical align="center">
          <n-text depth="3" style="font-size: 12px; text-align: center">
            Generate a link that expires, so you don't have to send raw credentials around.
            Anyone with the link can view this user's config until it expires.
          </n-text>
          <n-space align="center">
            <n-select v-model:value="ttlMinutes" :options="ttlOptions" style="width: 140px" />
            <n-button type="primary" :loading="generating" @click="generateLink">Generate link</n-button>
          </n-space>
          <template v-if="linkUrl">
            <canvas ref="linkCanvas" />
            <n-input :value="linkUrl" readonly />
            <n-space align="center">
              <n-button size="small" @click="copy(linkUrl)">Copy link</n-button>
              <n-text depth="3" style="font-size: 12px">{{ expiryText() }}</n-text>
            </n-space>
          </template>
        </n-space>
      </n-tab-pane>
      <n-tab-pane name="mierus" tab="mierus://" :disabled="!links">
        <n-space vertical align="center">
          <canvas ref="mierusCanvas" />
          <n-input :value="links?.mierusUrls[0]" type="textarea" readonly :autosize="{ minRows: 2, maxRows: 4 }" />
          <n-button size="small" @click="copy(links!.mierusUrls[0])">Copy link</n-button>
        </n-space>
      </n-tab-pane>
      <n-tab-pane name="mieru" tab="mieru://" :disabled="!links">
        <n-space vertical align="center">
          <canvas ref="mieruCanvas" />
          <n-input :value="links?.mieruUrl" type="textarea" readonly :autosize="{ minRows: 2, maxRows: 4 }" />
          <n-button size="small" @click="copy(links!.mieruUrl)">Copy link</n-button>
        </n-space>
      </n-tab-pane>
      <n-tab-pane name="json" tab="Config file" :disabled="!links">
        <n-space vertical>
          <n-input :value="links?.clientConfigJson" type="textarea" readonly :autosize="{ minRows: 8, maxRows: 14 }" />
          <n-space>
            <n-button size="small" @click="copy(links!.clientConfigJson)">Copy</n-button>
            <n-button size="small" @click="downloadConfig">Download JSON</n-button>
          </n-space>
        </n-space>
      </n-tab-pane>
    </n-tabs>
  </n-modal>
</template>
