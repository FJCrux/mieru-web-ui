<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { NCard, NForm, NFormItem, NInput, NButton, NSpace, NAlert, NText, useMessage, useDialog } from 'naive-ui'
import { useRouter } from 'vue-router'
import { api, ApiError } from '../api/client'
import type { Settings } from '../api/client'

const message = useMessage()
const dialog = useDialog()
const router = useRouter()

const publicHost = ref('')
const panelUrl = ref('')
const basePath = ref('')
const sharePath = ref('')
const restartPending = ref(false)
const currentPassword = ref('')
const newUsername = ref('')
const newPassword = ref('')

async function load() {
  try {
    const s = await api.get<Settings>('/api/settings')
    publicHost.value = s.publicHost
    panelUrl.value = s.panelUrl
    basePath.value = s.basePath
    sharePath.value = s.sharePath
    restartPending.value = s.restartPending
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : 'Failed to load settings')
  }
}

async function saveServer() {
  try {
    await api.put('/api/settings', {
      publicHost: publicHost.value,
      panelUrl: panelUrl.value,
      basePath: basePath.value,
      sharePath: sharePath.value,
    })
    message.success('Saved')
    await load()
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : 'Save failed')
  }
}

function confirmRestart() {
  dialog.warning({
    title: 'Restart panel',
    content:
      'The panel briefly goes offline and comes back with the current settings applied. ' +
      'This works when the container has a restart policy (the default compose does).',
    positiveText: 'Restart',
    negativeText: 'Cancel',
    onPositiveClick: restartPanel,
  })
}

async function restartPanel() {
  try {
    // Persist current settings first so the restart picks them up.
    await api.put('/api/settings', {
      publicHost: publicHost.value,
      panelUrl: panelUrl.value,
      basePath: basePath.value,
      sharePath: sharePath.value,
    })
    await api.post('/api/panel/restart')
    message.info('Panel is restarting, reconnecting...')
    setTimeout(() => window.location.reload(), 5000)
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : 'Restart failed')
  }
}

async function changePassword() {
  try {
    await api.put('/api/settings/password', {
      currentPassword: currentPassword.value,
      newUsername: newUsername.value,
      newPassword: newPassword.value,
    })
    message.success('Credentials changed, log in again')
    router.push('/login')
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : 'Change failed')
  }
}

onMounted(load)
</script>

<template>
  <h2 class="page-title">Settings</h2>
  <n-space vertical :size="16">
    <n-card title="Server">
      <n-form>
        <n-form-item label="Public host — address clients connect to, used in client configs">
          <n-input v-model:value="publicHost" placeholder="203.0.113.10 or vpn.example.com" style="max-width: 380px" />
        </n-form-item>
        <n-form-item label="Panel URL — this panel's external address, used for share links and host checking">
          <n-input v-model:value="panelUrl" placeholder="https://vpn.example.com" style="max-width: 380px" />
        </n-form-item>
        <n-button type="primary" @click="saveServer">Save</n-button>
      </n-form>
    </n-card>

    <n-card title="Paths">
      <n-alert v-if="restartPending" type="warning" title="Restart required" style="margin-bottom: 14px">
        Path settings have changed but the panel is still running the old ones.
        Use <b>Save &amp; restart panel</b> below to apply them.
      </n-alert>
      <n-alert type="info" :show-icon="false" style="margin-bottom: 14px">
        Serve the admin panel under a secret prefix (obscurity), and place share
        links on a separate public prefix so they don't reveal it. Changes take
        effect after restarting the panel (restart the container:
        <n-text code>docker compose restart</n-text>). Leave the base path empty
        to serve at the root.
      </n-alert>
      <n-form>
        <n-form-item label="Panel base path (secret admin prefix)">
          <n-input v-model:value="basePath" placeholder="empty = root, e.g. /a7Fq2xK9" style="max-width: 380px" />
        </n-form-item>
        <n-form-item label="Share path (public prefix for share links)">
          <n-input v-model:value="sharePath" placeholder="/s" style="max-width: 380px" />
        </n-form-item>
        <n-space>
          <n-button type="primary" @click="saveServer">Save</n-button>
          <n-button type="warning" secondary @click="confirmRestart">Save &amp; restart panel</n-button>
        </n-space>
      </n-form>
    </n-card>

    <n-card title="Panel admin">
      <n-form style="max-width: 380px">
        <n-form-item label="Current password">
          <n-input v-model:value="currentPassword" type="password" show-password-on="click" />
        </n-form-item>
        <n-form-item label="New username (optional)">
          <n-input v-model:value="newUsername" placeholder="keep current" />
        </n-form-item>
        <n-form-item label="New password (min 8 characters)">
          <n-input v-model:value="newPassword" type="password" show-password-on="click" />
        </n-form-item>
        <n-button type="primary" @click="changePassword">Change credentials</n-button>
      </n-form>
      <p class="hint">Changing credentials logs out every session, including this one.</p>
    </n-card>
  </n-space>
</template>

<style scoped>
.hint {
  font-size: 12px;
  opacity: 0.6;
  margin-top: 12px;
}
</style>
