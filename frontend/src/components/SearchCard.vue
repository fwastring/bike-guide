<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import InputField from './InputField.vue'
import LocationButton from './LocationButton.vue'
import BaseButton from './BaseButton.vue'
import { useScrollTo } from '@/composables/useScrollTo'

const router = useRouter()
const emit = defineEmits(['search'])
const address = ref('')
const difficulty = 'Medium'
const distance = '20-30 km'
const errorMessage = ref('')
const isMobile = ref(window.innerWidth < 640)
const { scrollToSection } = useScrollTo()

// Responsive: isMobile is true if window width < 640px (sm)
const windowWidth = ref(window.innerWidth)
const updateWidth = () => { windowWidth.value = window.innerWidth }
onMounted(() => window.addEventListener('resize', updateWidth))
onUnmounted(() => window.removeEventListener('resize', updateWidth))

const scrollToOnboarding = () => {
  scrollToSection('onboarding')
}

const handleSearch = () => {
  if (!address.value.trim()) {
    errorMessage.value = 'Please enter a location to search for routes.'
    return
  }
  errorMessage.value = ''
  emit('search', { address: address.value, difficulty, distance })
}

const handleLocation = async (location: { lat: number; lng: number }) => {
  try {
    // Get user's current position
    const position = await new Promise<GeolocationPosition>((resolve, reject) => {
      navigator.geolocation.getCurrentPosition(resolve, reject)
    })

    const { latitude, longitude } = position.coords

    // Reverse geocode the coordinates to get the address
    const response = await fetch(
      `https://nominatim.openstreetmap.org/reverse?format=json&lat=${latitude}&lon=${longitude}&zoom=18&addressdetails=1`
    )
    const data = await response.json()
    
    if (data.display_name) {
      address.value = data.display_name
      errorMessage.value = ''
    } else {
      throw new Error('Could not get address from coordinates')
    }
  } catch (error) {
    console.error('Error getting location:', error)
    errorMessage.value = 'Could not get your location. Please enter an address manually.'
  }
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
          @keyup.enter="handleSearch"
        />
        <div class="absolute right-4 top-1/2 -translate-y-1/2 flex items-center">
          <LocationButton @location="handleLocation" class="!bg-transparent !shadow-none !p-0 !w-7 !h-7" type="button" />
        </div>
      </div>
      <template v-if="isMobile">
        <div class="flex flex-row gap-2 w-full">
          <BaseButton
            title="Help"
            variant="secondary"
            class="h-10 px-6 text-base rounded-full whitespace-nowrap w-1/2"
            type="button"
            @click="scrollToOnboarding"
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