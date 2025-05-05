<script setup lang="ts">
import { useRouter } from 'vue-router'
import { onMounted } from 'vue'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const handleGoogleLogin = () => {
  // Redirect to backend Google auth endpoint which will handle the OAuth flow
  window.location.href = '/api/auth/google'
}

// Check for authorization code in URL (for callback)
onMounted(async () => {
  const urlParams = new URLSearchParams(window.location.search)
  const code = urlParams.get('code')
  if (code) {
    // If we have a code, we're in the callback phase
    // The backend will handle the code exchange and set the session cookie
    window.location.href = '/api/auth/google/callback?code=' + code
  } else {
    // Check if we're already authenticated
    const isAuthenticated = await authStore.checkAuth()
    if (isAuthenticated) {
      router.push('/map')
    }
  }
})
</script>

<template>
  <div class="login-container">
    <div class="login-box">
      <h1>Bike Guide</h1>
      <p>Welcome to Bike Guide! Please log in to continue.</p>
      <button class="google-login-btn" @click="handleGoogleLogin">
        <img src="../assets/google-icon.svg" alt="Google Icon" />
        Sign in with Google
      </button>
    </div>
  </div>
</template>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f5f5f5;
}

.login-box {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  text-align: center;
  max-width: 400px;
  width: 90%;
}

h1 {
  color: #333;
  margin-bottom: 1rem;
}

p {
  color: #666;
  margin-bottom: 2rem;
}

.google-login-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  background-color: white;
  color: #333;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.2s;
}

.google-login-btn:hover {
  background-color: #f5f5f5;
}

.google-login-btn img {
  width: 24px;
  height: 24px;
}
</style> 