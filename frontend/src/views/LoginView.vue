<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import LogInCard from '../components/LogInCard.vue'
import Header from '../components/Header.vue'
const router = useRouter()

const handleLogin = (credentials: { email: string; password: string }) => {
  // Here you would typically make an API call to authenticate the user
  console.log('Logging in with:', credentials)
  // For now, just redirect to home
  router.push('/')
}
</script>

<template>
  <Header />
  <div class="min-h-screen bg-gray-50 flex flex-col items-center justify-center p-4">
    
    <div class="w-full max-w-md">
      <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-gray-800 mb-2">Welcome Back</h1>
        <p class="text-gray-600">Log in to your account to continue</p>
      </div>

      <LogInCard @submit="handleLogin" />

      <div class="mt-8 text-center">
        <p class="text-gray-600">Or continue with</p>
        <div class="flex justify-center gap-4 mt-4">
          <button
            class="flex items-center justify-center w-12 h-12 rounded-full bg-white shadow-md hover:shadow-lg transition-shadow"
          >
            <svg class="w-6 h-6 text-gray-600" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12 0c-6.627 0-12 5.373-12 12s5.373 12 12 12 12-5.373 12-12-5.373-12-12-12zm-2 16h-2v-6h2v6zm-1-6.891c-.607 0-1.1-.496-1.1-1.109 0-.612.492-1.109 1.1-1.109s1.1.497 1.1 1.109c0 .613-.493 1.109-1.1 1.109zm8 6.891h-1.998v-2.861c0-1.881-2.002-1.722-2.002 0v2.861h-2v-6h2v1.093c.872-1.616 4-1.736 4 1.548v3.359z"/>
            </svg>
          </button>
          <button
            class="flex items-center justify-center w-12 h-12 rounded-full bg-white shadow-md hover:shadow-lg transition-shadow"
          >
            <svg class="w-6 h-6 text-gray-600" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12 0c-6.627 0-12 5.373-12 12s5.373 12 12 12 12-5.373 12-12-5.373-12-12-12zm-2 16h-2v-6h2v6zm-1-6.891c-.607 0-1.1-.496-1.1-1.109 0-.612.492-1.109 1.1-1.109s1.1.497 1.1 1.109c0 .613-.493 1.109-1.1 1.109zm8 6.891h-1.998v-2.861c0-1.881-2.002-1.722-2.002 0v2.861h-2v-6h2v1.093c.872-1.616 4-1.736 4 1.548v3.359z"/>
            </svg>
          </button>
          <button
            class="flex items-center justify-center w-12 h-12 rounded-full bg-white shadow-md hover:shadow-lg transition-shadow"
          >
            <svg class="w-6 h-6 text-gray-600" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12 0c-6.627 0-12 5.373-12 12s5.373 12 12 12 12-5.373 12-12-5.373-12-12-12zm-2 16h-2v-6h2v6zm-1-6.891c-.607 0-1.1-.496-1.1-1.109 0-.612.492-1.109 1.1-1.109s1.1.497 1.1 1.109c0 .613-.493 1.109-1.1 1.109zm8 6.891h-1.998v-2.861c0-1.881-2.002-1.722-2.002 0v2.861h-2v-6h2v1.093c.872-1.616 4-1.736 4 1.548v3.359z"/>
            </svg>
          </button>
        </div>
      </div>
    </div>
  </div>
</template> 