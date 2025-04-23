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
  <div class="bg-[#FBFBFB] p-4 sm:p-6 rounded-3xl border border-gray-300 max-w-prose border-[0.5px] w-full">
    <h2 class="text-lg sm:text-xl font-semibold text-gray-800 mb-4 text-center">Sign Up</h2>
    <form @submit.prevent="handleSignup" class="space-y-3">
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
      <span v-if="error" class="text-sm text-red-500 text-center">{{ error }}</span>
      <div class="flex justify-center">
        <BaseButton
          title="Sign Up"
          variant="primary"
          type="submit"
          class="w-full flex items-center justify-center px-4 py-2 text-sm font-medium"
        />
      </div>
    </form>
    <div class="text-center mt-3">
      <p class="text-sm text-gray-600">
        Already have an account? 
        <RouterLink to="/login" class="text-purple-500 hover:text-purple-600 hover:underline">Log in</RouterLink>
      </p>
    </div>
  </div>
</template>

<style scoped>
/* Remove the old styles since we're using Tailwind classes directly */
</style> 