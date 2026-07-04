<script setup lang="ts">
import { NConfigProvider, NMessageProvider, NDialogProvider, darkTheme } from 'naive-ui'

// Dark-only theme.
const themeOverrides = {
  common: {
    primaryColor: '#63e2b7',
    primaryColorHover: '#7fe7c4',
    primaryColorPressed: '#5acea7',
  },
}
</script>

<template>
  <n-config-provider :theme="darkTheme" :theme-overrides="themeOverrides">
    <n-message-provider>
      <n-dialog-provider>
        <router-view />
      </n-dialog-provider>
    </n-message-provider>
  </n-config-provider>
</template>

<style>
html,
body,
#app {
  margin: 0;
  height: 100%;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
}
/* Consistent page heading across pages. */
.page-title {
  margin: 0 0 18px;
  font-size: 20px;
  font-weight: 650;
  letter-spacing: 0.2px;
}
/* Dark background behind the app (covers areas outside Naive layout, e.g. the
   login and public share pages). */
html,
body {
  background: #101014;
  color: #ffffffd1;
}
</style>
