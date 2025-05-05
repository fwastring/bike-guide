<script setup lang="ts">
import { ref } from 'vue'
import InputField from './InputField.vue'
import LocationButton from './LocationButton.vue'
import BaseButton from './BaseButton.vue'

const emit = defineEmits(['search'])
const address = ref('')
const difficulty = 'Medium'
const distance = '20-30 km'
const errorMessage = ref('')

const handleSearch = () => {
  if (!address.value.trim()) {
    errorMessage.value = 'Please enter a location'
    return
  }
  
  errorMessage.value = ''
  emit('search', { address: address.value, difficulty, distance })
}

const handleLocation = async (location: { lat: number; lng: number }) => {
  // Dummy location logic; replace with real geolocation if needed
  address.value = 'Current Location (Demo)'
  errorMessage.value = ''
}
</script>

<template>
  <form @submit.prevent="handleSearch" class="flex flex-col sm:flex-row items-center gap-4 w-full sm:max-w-3xl mx-auto">
    <div class="relative flex-1 w-full">
      <InputField
        v-model="address"
        placeholder="Vägvägen 7"
        :big="true"
        :errorBorder="!!errorMessage"
        class="w-full"
      />
      <div class="absolute right-4 top-1/2 -translate-y-1/2 flex items-center">
        <LocationButton @location="handleLocation" class="!bg-transparent !shadow-none !p-0 !w-7 !h-7" />
      </div>
    </div>
    <BaseButton
      title="Find Route"
      variant="primary"
      class="h-14 px-10 text-lg rounded-full whitespace-nowrap w-full sm:w-auto"
      type="submit"
    />
  </form>
  <div v-if="errorMessage" class="mt-6 sm:mt-2 text-center md:text-left w-full">
    <span class="text-red-500 text-md">{{ errorMessage }}</span>
  </div>
</template> 