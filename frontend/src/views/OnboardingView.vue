<script setup lang="ts">
import { ref } from 'vue'
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
</script>

<template>
  <div class="min-h-screen bg-blue-50 flex items-center py-16 px-4 sm:px-8">
    <div class="max-w-7xl mx-auto w-full">
      <section class="flex flex-col sm:flex-row items-center justify-between max-w-7xl mx-auto w-full px-4 py-6 sm:py-10 gap-6 sm:gap-4">
        <div class="max-w-2xl text-left w-full">
          <h1 class="text-4xl sm:text-5xl lg:text-6xl font-bold text-gray-900 mb-2 sm:mb-4">But how does it work?</h1>
          <p class="text-xl sm:text-2xl lg:text-3xl text-gray-500 mb-2">TextTextTextTextText</p>
          <p class="text-lg sm:text-xl text-gray-600 mb-2">TextTextTextTextTextTextTextText</p>
        </div>
        
      </section>
      
      <div class="flex flex-col sm:flex-row gap-6 sm:gap-8 mb-12">
        <div 
          v-for="(step, index) in onboardingSteps" 
          :key="index"
          class="flex-1 bg-[#FBFBFB] rounded-3xl p-8 relative border border-gray-200"
        >
          <div class="mb-8">
            <component 
              :is="step.icon" 
              class="w-16 h-16 mx-auto mb-6 text-[#8E7DBE]"
            />
            <h3 class="text-2xl font-bold text-center text-gray-900 mb-3">
              {{ step.title }}
            </h3>
            <p class="text-center text-gray-600">
              {{ step.description }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

