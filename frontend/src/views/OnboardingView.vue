<script setup lang="ts">
import { ref } from 'vue'
import SelectCard from '../components/SelectCard.vue'
import BaseButton from '../components/BaseButton.vue'
import { useRouter } from 'vue-router'
import Header from '../components/Header.vue'
const router = useRouter()
const currentStep = ref(1)

const bikeTypes = [
  {
    id: 'road',
    title: 'Road Bike',
    description: 'Lightweight and fast, perfect for paved roads and long distances.'
  },
  {
    id: 'mountain',
    title: 'Mountain Bike',
    description: 'Durable and versatile, great for off-road trails and rough terrain.'
  },
  {
    id: 'hybrid',
    title: 'Hybrid Bike',
    description: 'A mix of road and mountain bike features, good for commuting and light trails.'
  },
  {
    id: 'gravel',
    title: 'Gravel Bike',
    description: 'Versatile bike that can handle both paved roads and gravel paths.'
  }
]

const experienceLevels = [
  {
    id: 'beginner',
    title: 'Beginner',
    description: 'New to cycling or getting back into it after a long break.'
  },
  {
    id: 'intermediate',
    title: 'Intermediate',
    description: 'Regular cyclist with some experience on different terrains.'
  },
  {
    id: 'advanced',
    title: 'Advanced',
    description: 'Experienced cyclist comfortable with various conditions and distances.'
  }
]

const selectedBikeType = ref<string | null>(null)
const selectedExperience = ref<string | null>(null)

const handleBikeTypeSelect = (type: string) => {
  selectedBikeType.value = type
}

const handleExperienceSelect = (level: string) => {
  selectedExperience.value = level
}

const handleContinue = () => {
  if (currentStep.value === 1 && selectedBikeType.value) {
    currentStep.value = 2
  } else if (currentStep.value === 2 && selectedExperience.value) {
    // Here you would typically save the selections and proceed
    router.push('/search')
  }
}
</script>

<template>
  <Header />
  <div class="min-h-screen bg-gray-50 flex flex-col items-center justify-center p-4">
    <div class="w-full max-w-4xl">
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-center text-gray-800 mb-4">Find Your Perfect Bike</h1>
        <p class="text-center text-gray-600">Let's help you find the right bike for your needs</p>
      </div>

      <div v-if="currentStep === 1">
        <SelectCard
          title="What type of bike are you looking for?"
          :options="bikeTypes"
          @select="handleBikeTypeSelect"
          @continue="handleContinue"
        />
      </div>

      <div v-if="currentStep === 2">
        <SelectCard
          title="What's your experience level?"
          :options="experienceLevels"
          @select="handleExperienceSelect"
          @continue="handleContinue"
        />
      </div>

      <div class="flex justify-center mt-4">
        <div class="flex gap-2">
          <div
            v-for="step in 2"
            :key="step"
            class="w-3 h-3 rounded-full"
            :class="step <= currentStep ? 'bg-blue-600' : 'bg-gray-300'"
          ></div>
        </div>
      </div>
    </div>
  </div>
</template> 