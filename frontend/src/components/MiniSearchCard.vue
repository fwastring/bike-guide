<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import InputField from './InputField.vue'
import LocationButton from './LocationButton.vue'
import BaseButton from './BaseButton.vue'

const emit = defineEmits(['search'])
const address = ref('')
const difficulty = 'Medium'
const distance = '20-30 km'
const errorMessage = ref('')

// Responsive: isMobile is true if window width < 640px (sm)
const windowWidth = ref(window.innerWidth)
const updateWidth = () => { windowWidth.value = window.innerWidth }
onMounted(() => window.addEventListener('resize', updateWidth))
onUnmounted(() => window.removeEventListener('resize', updateWidth))
const isMobile = computed(() => windowWidth.value < 640)

const handleSearch = () => {
  if (!address.value.trim()) {
    errorMessage.value = 'Please enter a location to search for routes.'
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
  <div class="relative w-full">
    <form @submit.prevent="handleSearch" class="flex flex-col sm:flex-row items-center gap-4 w-full sm:max-w-3xl mx-auto">
      <div class="relative flex-1 w-full">
        <InputField
          v-model="address"
          placeholder="Vägvägen 7"
          :mobile="isMobile"
          :errorBorder="!!errorMessage"
          class="w-full"
        />
        <div class="absolute right-4 top-1/2 -translate-y-1/2 flex items-center">
          <LocationButton @location="handleLocation" class="!bg-transparent !shadow-none !p-0 !w-7 !h-7" />
        </div>
      </div>
      <template v-if="isMobile">
        <div class="flex flex-row gap-2 w-full">
          <BaseButton
            title="Help"
            variant="secondary"
            class="h-10 px-6  text-base rounded-full whitespace-nowrap w-1/2"
            type="button"
          />
          <BaseButton
            title="Find Route"
            variant="primary"
            class="h-10 px-6 text-base rounded-full whitespace-nowrap w-2/3"
            type="submit"
            hover="bg-primary-dark"
            cursor="pointer"
          />
        </div>
      </template>
      <template v-else>
        <BaseButton
          title="Find Route"
          variant="primary"
          class="h-14 px-10 text-lg rounded-full whitespace-nowrap w-full sm:w-auto"
          type="submit"
        />
      </template>
    </form>
    <transition
      enter-active-class="transition-opacity duration-500"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition-opacity duration-300"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div v-if="errorMessage" class="absolute left-1/2 -translate-x-1/2 mt-2 w-full flex justify-center pointer-events-none z-20">
        <span class="text-blue-500 text-md font-semibold  bg-opacity-80 px-3 py-1 ">{{ errorMessage }}</span>
      </div>
    </transition>
  </div>
</template> 