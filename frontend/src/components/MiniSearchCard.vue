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
const distance = '20-30 km'
const errorMessage = ref('')
const isMobile = ref(window.innerWidth < 640)

// Responsive: isMobile is true if window width < 640px (sm)
const windowWidth = ref(window.innerWidth)
const updateWidth = () => { windowWidth.value = window.innerWidth }
onMounted(() => window.addEventListener('resize', updateWidth))
onUnmounted(() => window.removeEventListener('resize', updateWidth))

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
    <form @submit.prevent="handleSearch" class="flex flex-col sm:flex-row items-center gap-3 w-full sm:max-w-3xl mx-auto">
      <div class="relative flex-1 w-full">
        <InputField
          v-model="address"
          placeholder="Vägvägen 7"
          :mobile="isMobile"
          :errorBorder="!!errorMessage"
          :compact="true"
          class="w-full"
        />
        <div class="absolute right-3 top-1/2 -translate-y-1/2 flex items-center">
          <LocationButton @location="handleLocation" class="!bg-transparent !shadow-none !p-0 !w-5 !h-5" />
        </div>
      </div>
      <BaseButton
        title="Search"
        variant="secondary"
        class="h-10 px-6 text-sm rounded-full whitespace-nowrap w-full sm:w-auto"
        type="submit"
      />
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
        <span class="text-blue-500 text-sm font-semibold bg-opacity-80 px-3 py-1">{{ errorMessage }}</span>
      </div>
    </transition>
  </div>
</template> 