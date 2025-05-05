<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import BaseButton from './BaseButton.vue'
import LocationButton from './LocationButton.vue'
import IconButton from './IconButton.vue'
import { Settings } from 'lucide-vue-next'
import { CirclePlus } from 'lucide-vue-next'

interface SearchData {
  address: string
  difficulty: string
  distance: string
}

const emit = defineEmits(['search'])
const router = useRouter()

const address = ref('')
const difficulty = ref('Medium')
const distance = ref('20-30 km')

const difficultyOptions = ['Easy', 'Medium', 'Hard']
const distanceOptions = ['5-10 km', '10-20 km', '20-30 km', '30-40 km', '40+ km']

const handleSearch = async () => {
  const searchData = {
    address: address.value,
    difficulty: difficulty.value,
    distance: distance.value
  }
  emit('search', searchData)
  try {
    await router.push('/options')
  } catch (error) {
    console.error('Navigation failed:', error)
  }
}

const handleLocation = async (location: { lat: number; lng: number }) => {
  try {
    const response = await fetch(
      `https://nominatim.openstreetmap.org/reverse?format=json&lat=${location.lat}&lon=${location.lng}`
    )
    const data = await response.json()
    address.value = data.display_name
  } catch (error) {
    console.error('Error getting address:', error)
  }
}
</script>

<template>
  <div class="bg-[#FBFBFB] p-4 sm:p-6 rounded-3xl border border-gray-300 max-w-prose border-[0.5px]">
    <h2 class="text-lg sm:text-xl font-semibold text-gray-800 mb-4">Find Your Route</h2>
    
    <!-- Address Input -->
    <div class="space-y-1.5 mb-4">
      <label class="block text-sm sm:text-base font-medium text-gray-800">Search for a place or address</label>
      <div class="relative">
        <input
          v-model="address"
          type="text"
          placeholder="Vägvägen 7"
          class="w-full pl-3 pr-8 py-2 bg-white border border-gray-300 rounded-full text-sm text-gray-800 placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent border-[0.5px]"
        />
        <div class="absolute right-2 top-1/2 -translate-y-1/2">
          <LocationButton @location="handleLocation" />
        </div>
      </div>
    </div>

    <!-- Add Destination -->
    <button class="flex items-center gap-1.5 text-sm text-gray-800 mb-4 hover:text-purple-500 transition-colors">
      <CirclePlus class="h-4 w-4" />
      <span>Add destination</span>
    </button>

    <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-6">
      <!-- Difficulty Dropdown -->
      <div class="space-y-1.5">
        <label class="block text-sm sm:text-base font-medium text-gray-800">Difficulty</label>
        <div class="relative">
          <select
            v-model="difficulty"
            class="w-full appearance-none px-3 py-2 bg-white border border-gray-300 rounded-full text-sm text-gray-800 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent border-[0.5px]"
          >
            <option v-for="option in difficultyOptions" :key="option" :value="option">
              {{ option }}
            </option>
          </select>
          <div class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none text-gray-400">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 24 24" fill="none" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
            </svg>
          </div>
        </div>
      </div>

      <!-- Distance Dropdown -->
      <div class="space-y-1.5">
        <label class="block text-sm sm:text-base font-medium text-gray-800">Distance</label>
        <div class="relative">
          <select
            v-model="distance"
            class="w-full appearance-none px-3 py-2 bg-white border border-gray-300 rounded-full text-sm text-gray-800 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent border-[0.5px]"
          >
            <option v-for="option in distanceOptions" :key="option" :value="option">
              {{ option }}
            </option>
          </select>
          <div class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none text-gray-400">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" viewBox="0 0 24 24" fill="none" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
            </svg>
          </div>
        </div>
      </div>
    </div>

    <div class="flex justify-between mb-4 space-x-3">
      <BaseButton
        title="Find route!"
        variant="primary"
        class="w-full flex items-center justify-center px-4 py-2 text-sm font-medium"
        @click="handleSearch"
      >
        <span>Find route!</span>
        <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 ml-1.5" viewBox="0 0 24 24" fill="none" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3" />
        </svg>
      </BaseButton>
      <IconButton title="Settings" :icon="Settings" variant="secondary" size="md" />
    </div>
  </div>
</template>