<script setup lang="ts">
import { ref } from 'vue'
import InputField from './InputField.vue'
import BaseButton from './BaseButton.vue'

const email = ref('')
const password = ref('')
const isSubmitting = ref(false)
const errors = ref({
  email: '',
  password: ''
})

const emit = defineEmits(['submit'])

const handleSubmit = () => {
  // Reset errors
  errors.value = {
    email: '',
    password: ''
  }

  // Basic validation
  if (!email.value) {
    errors.value.email = 'Email is required'
  }
  if (!password.value) {
    errors.value.password = 'Password is required'
  }

  if (Object.values(errors.value).some(error => error)) {
    return
  }

  isSubmitting.value = true
  emit('submit', {
    email: email.value,
    password: password.value
  })
}
</script>

<template>
  <div class="bg-[#FBFBFB] p-4 sm:p-6 rounded-3xl border border-gray-300 max-w-prose border-[0.5px] w-full">
    <h2 class="text-lg sm:text-xl font-semibold text-gray-800 mb-4 text-center">Log In</h2>
    <form @submit.prevent="handleSubmit" class="space-y-3">
      <InputField
        v-model="email"
        label="Email"
        type="email"
        placeholder="Enter your email"
        :error="errors.email"
        required
      />
      <InputField
        v-model="password"
        label="Password"
        type="password"
        placeholder="Enter your password"
        :error="errors.password"
        required
      />
      <div class="flex justify-center">
        <BaseButton
          title="Log In"
          variant="primary"
          type="submit"
          :disabled="isSubmitting"
          class="w-full flex items-center justify-center px-4 py-2 text-sm font-medium"
        />
      </div>
    </form>
    <div class="text-center mt-3">
      <p class="text-sm text-gray-600">
        Don't have an account? 
        <RouterLink to="/signup" class="text-purple-500 hover:text-purple-600 hover:underline">Sign up</RouterLink>
      </p>
    </div>
  </div>
</template> 