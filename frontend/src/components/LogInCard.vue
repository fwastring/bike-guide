<script setup lang="ts">
import { ref } from 'vue'
import InputField from './InputField.vue'
import BaseButton from './BaseButton.vue'

const email = ref('')
const password = ref('')
const error = ref('')

const emit = defineEmits(['login'])

const handleLogin = () => {
  if (!email.value || !password.value) {
    error.value = 'Please fill in all fields'
    return
  }
  emit('login', { email: email.value, password: password.value })
}
</script>

<template>
  <div class="login-card">
    <h2>Login</h2>
    <form @submit.prevent="handleLogin" class="login-form">
      <InputField
        v-model="email"
        label="Email"
        type="email"
        placeholder="Enter your email"
      />
      <InputField
        v-model="password"
        label="Password"
        type="password"
        placeholder="Enter your password"
      />
      <span v-if="error" class="error-message">{{ error }}</span>
      <BaseButton
        title="Login"
        variant="primary"
        @click="handleLogin"
      />
    </form>
    <div class="signup-link">
      <p>Don't have an account? <RouterLink to="/signup">Sign up</RouterLink></p>
    </div>
  </div>
</template>

<style scoped>
.login-card {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  max-width: 400px;
  width: 100%;
}

h2 {
  text-align: center;
  margin-bottom: 1.5rem;
  color: #333;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.error-message {
  color: #dc3545;
  text-align: center;
}

.signup-link {
  text-align: center;
  margin-top: 1rem;
}

.signup-link a {
  color: #007bff;
  text-decoration: none;
}

.signup-link a:hover {
  text-decoration: underline;
}
</style> 