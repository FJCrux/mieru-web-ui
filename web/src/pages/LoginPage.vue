<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { NForm, NFormItem, NInput, NButton, useMessage } from 'naive-ui'
import { api, ApiError } from '../api/client'

const router = useRouter()
const message = useMessage()
const username = ref('')
const password = ref('')
const loading = ref(false)

async function login() {
  loading.value = true
  try {
    await api.post('/api/login', { username: username.value, password: password.value })
    router.push('/')
  } catch (e) {
    message.error(e instanceof ApiError ? e.message : 'Login failed')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="wrap">
    <div class="card">
      <div class="brand">
        <span class="dot" />
        <span>mieru-web-ui</span>
      </div>
      <p class="sub">Sign in to your panel</p>
      <n-form @keyup.enter="login">
        <n-form-item label="Username">
          <n-input v-model:value="username" placeholder="admin" autofocus size="large" />
        </n-form-item>
        <n-form-item label="Password">
          <n-input v-model:value="password" type="password" show-password-on="click" size="large" />
        </n-form-item>
        <n-button type="primary" block size="large" :loading="loading" @click="login">Log in</n-button>
      </n-form>
    </div>
  </div>
</template>

<style scoped>
.wrap {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 16px;
  background:
    radial-gradient(900px 500px at 50% -10%, rgba(99, 226, 183, 0.09), transparent 60%),
    #101014;
}
.card {
  width: 360px;
  background: #18181c;
  border: 1px solid #26262c;
  border-radius: 14px;
  padding: 28px;
  box-shadow: 0 24px 60px rgba(0, 0, 0, 0.45);
}
.brand {
  display: flex;
  align-items: center;
  gap: 10px;
  font-weight: 700;
  font-size: 18px;
}
.brand .dot {
  width: 9px;
  height: 9px;
  border-radius: 50%;
  background: #63e2b7;
  box-shadow: 0 0 10px #63e2b7aa;
}
.sub {
  margin: 6px 0 20px;
  font-size: 13px;
  opacity: 0.6;
}
</style>
