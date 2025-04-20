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
  <div class="bg-white p-8 rounded-lg shadow-md max-w-md w-full">
    <h2 class="text-2xl font-bold text-center text-gray-800 mb-6">Log In</h2>
    <form @submit.prevent="handleSubmit" class="space-y-4">
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
        />
      </div>
    </form>
    <div class="text-center mt-4">
      <p class="text-gray-600">
        Don't have an account? 
        <RouterLink to="/signup" class="text-blue-600 hover:text-blue-700 hover:underline">Sign up</RouterLink>
      </p>
    </div>
  </div>
</template> 