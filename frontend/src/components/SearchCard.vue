<script setup lang="ts">
import { ref } from 'vue'
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

const address = ref('')
const difficulty = ref('Medium')
const distance = ref('20-30 km')

const difficultyOptions = ['Easy', 'Medium', 'Hard']
const distanceOptions = ['5-10 km', '10-20 km', '20-30 km', '30-40 km', '40+ km']

const handleSearch = () => {
  emit('search', {
    address: address.value,
    difficulty: difficulty.value,
    distance: distance.value
  })
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
  <div class="bg-[#FBFBFB] p-6 sm:p-8 rounded-4xl border border-gray-300 max-w-prose">
    <h2 class="text-xl sm:text-2xl font-semibold text-gray-800 mb-6">Find Your Route</h2>
    
    <!-- Address Input -->
    <div class="space-y-2 mb-6">
      <label class="block text-base sm:text-lg font-medium text-gray-800">Search for a place or address</label>
      <div class="relative">
        <input
          v-model="address"
          type="text"
          placeholder="Vägvägen 7"
          class="w-full pl-4 pr-10 py-3 bg-white border border-gray-300 rounded-full text-base text-gray-800 placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent"
        />
        <div class="absolute right-3 top-1/2 -translate-y-1/2">
          <LocationButton @location="handleLocation" />
        </div>
      </div>
    </div>

    <!-- Add Destination -->
    <button class="flex items-center gap-2 text-base text-gray-800 mb-6 hover:text-purple-500 transition-colors">
      <CirclePlus class="h-5 w-5" />
      <span>Add destination</span>
    </button>

    <div class="grid grid-cols-1 sm:grid-cols-2 gap-6 mb-8">
      <!-- Difficulty Dropdown -->
      <div class="space-y-2">
        <label class="block text-base sm:text-lg font-medium text-gray-800">Difficulty</label>
        <div class="relative">
          <select
            v-model="difficulty"
            class="w-full appearance-none px-4 py-3 bg-white border border-gray-300 rounded-full text-base text-gray-800 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent"
          >
            <option v-for="option in difficultyOptions" :key="option" :value="option">
              {{ option }}
            </option>
          </select>
          <div class="absolute right-4 top-1/2 -translate-y-1/2 pointer-events-none text-gray-400">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
            </svg>
          </div>
        </div>
      </div>

      <!-- Distance Dropdown -->
      <div class="space-y-2">
        <label class="block text-base sm:text-lg font-medium text-gray-800">Distance</label>
        <div class="relative">
          <select
            v-model="distance"
            class="w-full appearance-none px-4 py-3 bg-white border border-gray-300 rounded-full text-base text-gray-800 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent"
          >
            <option v-for="option in distanceOptions" :key="option" :value="option">
              {{ option }}
            </option>
          </select>
          <div class="absolute right-4 top-1/2 -translate-y-1/2 pointer-events-none text-gray-400">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 24 24" fill="none" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
            </svg>
          </div>
        </div>
      </div>
    </div>

    <div class="border-t border-gray-300 pt-4 flex justify-between space-x-4">
      <BaseButton
        title="Find route!"
        variant="primary"
        class="w-full flex items-center justify-center px-6 py-3 text-base font-medium"
        @click="handleSearch"
      >
        <span>Find route!</span>
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 ml-2" viewBox="0 0 24 24" fill="none" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3" />
        </svg>
      </BaseButton>
      <IconButton title="Settings" :icon="Settings" variant="secondary" size="lg" />
    </div>
  </div>
</template>