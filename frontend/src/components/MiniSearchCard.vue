<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import InputField from './InputField.vue'
import LocationButton from './LocationButton.vue'
import BaseButton from './BaseButton.vue'

const router = useRouter()
const emit = defineEmits(['search'])
const address = ref('')
const difficulty = 'Medium'
const distance = '10-20 km'
const isMobile = ref(window.innerWidth < 640)

// Responsive: isMobile is true if window width < 640px (sm)
const windowWidth = ref(window.innerWidth)
const updateWidth = () => { windowWidth.value = window.innerWidth }
onMounted(() => window.addEventListener('resize', updateWidth))
onUnmounted(() => window.removeEventListener('resize', updateWidth))

const handleSearch = () => {
  if (!address.value.trim()) {
    return
  }
  emit('search', { address: address.value, difficulty, distance })
  address.value = '' // Clear the input field after successful search
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
    }
  } catch (error) {
    console.error('Error getting location:', error)
  }
}
</script>

<template>
  <div class="relative w-full">
    <form @submit.prevent="handleSearch" class="flex flex-col sm:flex-row items-center gap-3 w-full sm:max-w-3xl mx-auto px-2">
      <div class="relative flex-1 w-full">
        <InputField
          v-model="address"
          placeholder="Vägvägen 7"
          :mobile="isMobile"
          :compact="true"
          class="w-full"
          @keyup.enter="handleSearch"
        />
        <div class="absolute right-3 top-1/2 -translate-y-1/2 flex items-center">
          <LocationButton 
            @location="handleLocation" 
            class="!bg-transparent !shadow-none !p-0 !w-5 !h-5"
            type="button"
          />
        </div>
      </div>
      
    </form>
  </div>
</template> 