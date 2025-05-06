<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import BaseButton from '../components/BaseButton.vue'
import { useRouter } from 'vue-router'
import Header from '../components/Header.vue'
import { MapPin, Map, Bike } from 'lucide-vue-next'

const router = useRouter()

const onboardingSteps = [
  {
    title: 'Find Route',
    description: 'Set your parameters for your personalized bike route.',
    icon: MapPin,
  },
  {
    title: 'Choose optimal route',
    description: 'Select one out of several suggested routes based on your preference',
    icon: Map,
  },
  {
    title: 'Go biking!',
    description: 'Save the route to your phone and start biking!',
    icon: Bike,
  }
]

const handleGetStarted = () => {
  router.push('/search')
}

const parallaxY = ref(0)

function handleScroll() {
  // Adjust the divisor for more/less parallax effect
  parallaxY.value = window.scrollY * 0.3
}

onMounted(() => {
  window.addEventListener('scroll', handleScroll)
})
onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
})
</script>

<template>
  <section class="relative py-16 md:py-24 bg-gray-50 overflow-hidden">
    <img
      src="@/assets/mapbackground.jpg"
      alt="Map background"
      class="absolute inset-0 w-full h-full object-cover opacity-10 pointer-events-none select-none z-0 will-change-transform transition-transform duration-300"
      aria-hidden="true"
      :style="{ transform: `translateY(${parallaxY}px)` }"
    />
    <div
      class="absolute inset-0 w-full h-full z-0 pointer-events-none select-none bg-blue-500 opacity-20 mix-blend-multiply will-change-transform transition-transform duration-300"
      :style="{ transform: `translateY(${parallaxY}px)` }"
    ></div>
    <div class="relative z-10 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-16">
        <h2 class="font-pramukh text-6xl sm:text-6xl lg:text-7xl text-gray-600 mb-2 sm:mb-4">How It Works</h2>
        <p class="font-poppins text-lg sm:text-xl text-gray-600 max-w-3xl mx-auto">Find your perfect cycling route in three simple steps</p>
      </div>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <!-- Step 1 -->
        <div class="bg-white rounded-lg p-6 shadow-sm">
          <div class="w-16 h-16 bg-blue-100 rounded-full flex items-center justify-center mb-4 mx-auto">
            <MapPin class="w-8 h-8 text-blue-600" />
          </div>
          <h3 class="font-poppins text-xl font-semibold text-gray-900 mb-2 text-center">Find Route</h3>
          <p class="font-poppins text-gray-600 text-center">Set your parameters for your personalized bike route.</p>
        </div>
        <!-- Step 2 -->
        <div class="bg-white rounded-lg p-6 shadow-sm">
          <div class="w-16 h-16 bg-blue-100 rounded-full flex items-center justify-center mb-4 mx-auto">
            <Map class="w-8 h-8 text-blue-600" />
          </div>
          <h3 class="font-poppins text-xl font-semibold text-gray-900 mb-2 text-center">Choose optimal route</h3>
          <p class="font-poppins text-gray-600 text-center">Select one out of several suggested routes based on your preference</p>
        </div>
        <!-- Step 3 -->
        <div class="bg-white rounded-lg p-6 shadow-sm">
          <div class="w-16 h-16 bg-blue-100 rounded-full flex items-center justify-center mb-4 mx-auto">
            <Bike class="w-8 h-8 text-blue-600" />
          </div>
          <h3 class="font-poppins text-xl font-semibold text-gray-900 mb-2 text-center">Go biking!</h3>
          <p class="font-poppins text-gray-600 text-center">Save the route to your phone and start biking!</p>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.font-pramukh {
  font-family: 'Pramukh Rounded', sans-serif;
}

.font-poppins {
  font-family: 'Poppins', sans-serif;
}
</style>

