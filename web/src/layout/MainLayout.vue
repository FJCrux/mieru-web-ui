<script setup lang="ts">
import { h, ref, onMounted, onUnmounted, computed } from 'vue'
import { RouterLink, useRoute, useRouter } from 'vue-router'
import {
  NLayout, NLayoutSider, NLayoutHeader, NLayoutContent, NMenu, NButton, NSpace,
  NText, NIcon, NTag, NDropdown,
} from 'naive-ui'
import type { MenuOption } from 'naive-ui'
import {
  GridOutline, PeopleOutline, GitNetworkOutline, SwapHorizontalOutline,
  LinkOutline, SettingsOutline, PersonCircleOutline, LogOutOutline,
} from '@vicons/ionicons5'
import { api } from '../api/client'
import type { Dashboard } from '../api/client'

const route = useRoute()
const router = useRouter()
const username = ref('')
const status = ref('')
const version = ref('')
let timer: number | undefined

function icon(c: unknown) {
  return () => h(NIcon, null, { default: () => h(c as never) })
}

const menuOptions: MenuOption[] = [
  { label: () => h(RouterLink, { to: '/' }, { default: () => 'Dashboard' }), key: '/', icon: icon(GridOutline) },
  { label: () => h(RouterLink, { to: '/users' }, { default: () => 'Users' }), key: '/users', icon: icon(PeopleOutline) },
  { label: () => h(RouterLink, { to: '/network' }, { default: () => 'Network' }), key: '/network', icon: icon(GitNetworkOutline) },
  { label: () => h(RouterLink, { to: '/outbounds' }, { default: () => 'Outbounds' }), key: '/outbounds', icon: icon(SwapHorizontalOutline) },
  { label: () => h(RouterLink, { to: '/chain' }, { default: () => 'Chain' }), key: '/chain', icon: icon(LinkOutline) },
  { label: () => h(RouterLink, { to: '/settings' }, { default: () => 'Settings' }), key: '/settings', icon: icon(SettingsOutline) },
]

const statusType = computed(() =>
  status.value === 'RUNNING' ? 'success' : status.value === 'IDLE' ? 'warning' : status.value ? 'error' : 'default',
)

const userMenu = [
  { label: 'Log out', key: 'logout', icon: icon(LogOutOutline) },
]

async function poll() {
  try {
    const d = await api.get<Dashboard>('/api/dashboard')
    status.value = d.mitaStatus
    version.value = d.mitaVersion
  } catch {
    /* handled by router on 401 */
  }
}

onMounted(async () => {
  try {
    const me = await api.get<{ username: string }>('/api/me')
    username.value = me.username
  } catch {
    /* 401 already redirected to /login */
  }
  poll()
  timer = window.setInterval(() => {
    if (!document.hidden) poll()
  }, 8000)
})
onUnmounted(() => window.clearInterval(timer))

async function onUser(key: string) {
  if (key === 'logout') {
    try {
      await api.post('/api/logout')
    } finally {
      router.push('/login')
    }
  }
}
</script>

<template>
  <n-layout position="absolute" has-sider>
    <n-layout-sider bordered :collapsed-width="64" :width="220" collapse-mode="width" show-trigger="bar">
      <div class="brand">
        <span class="dot" />
        <span class="name">mieru</span>
      </div>
      <n-menu :value="route.path" :options="menuOptions" :indent="20" />
    </n-layout-sider>

    <n-layout>
      <n-layout-header bordered class="header">
        <n-space align="center" :size="10">
          <n-tag :type="statusType" size="small" round>{{ status || '…' }}</n-tag>
          <n-text v-if="version" depth="3" style="font-size: 12px">mita v{{ version }}</n-text>
        </n-space>
        <n-dropdown :options="userMenu" @select="onUser" trigger="click">
          <n-button text>
            <template #icon><n-icon :component="PersonCircleOutline" /></template>
            {{ username || 'admin' }}
          </n-button>
        </n-dropdown>
      </n-layout-header>

      <n-layout-content class="content" content-style="padding: 24px; max-width: 1100px; margin: 0 auto;">
        <router-view />
      </n-layout-content>
    </n-layout>
  </n-layout>
</template>

<style scoped>
.brand {
  height: 56px;
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 0 20px;
  font-weight: 700;
  font-size: 17px;
  letter-spacing: 0.3px;
}
.brand .dot {
  width: 9px;
  height: 9px;
  border-radius: 50%;
  background: #63e2b7;
  box-shadow: 0 0 10px #63e2b7aa;
}
.header {
  height: 56px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
}
.content {
  overflow: auto;
}
</style>
