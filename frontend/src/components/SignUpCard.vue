<script setup lang="ts">
import { ref } from 'vue'
import InputField from './InputField.vue'
import BaseButton from './BaseButton.vue'

const name = ref('')
const email = ref('')
const password = ref('')
const confirmPassword = ref('')
const error = ref('')

const emit = defineEmits(['signup'])

const handleSignup = () => {
  if (!name.value || !email.value || !password.value || !confirmPassword.value) {
    error.value = 'Please fill in all fields'
    return
  }
  if (password.value !== confirmPassword.value) {
    error.value = 'Passwords do not match'
    return
  }
  emit('signup', {
    name: name.value,
    email: email.value,
    password: password.value
  })
}
</script>

<template>
  <div class="signup-card">
    <h2>Sign Up</h2>
    <form @submit.prevent="handleSignup" class="signup-form">
      <InputField
        v-model="name"
        label="Full Name"
        placeholder="Enter your full name"
      />
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
        placeholder="Create a password"
      />
      <InputField
        v-model="confirmPassword"
        label="Confirm Password"
        type="password"
        placeholder="Confirm your password"
      />
      <span v-if="error" class="error-message">{{ error }}</span>
      <BaseButton
        title="Sign Up"
        variant="primary"
        @click="handleSignup"
      />
    </form>
    <div class="login-link">
      <p>Already have an account? <RouterLink to="/login">Log in</RouterLink></p>
    </div>
  </div>
</template>

<style scoped>
.signup-card {
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

.signup-form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.error-message {
  color: #dc3545;
  text-align: center;
}

.login-link {
  text-align: center;
  margin-top: 1rem;
}

.login-link a {
  color: #007bff;
  text-decoration: none;
}

.login-link a:hover {
  text-decoration: underline;
}
</style> 