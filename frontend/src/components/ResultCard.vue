<script setup lang="ts">
import BaseButton from './BaseButton.vue'

interface Props {
  title: string
  distance: string
  elevation: string
  difficulty: string
  description: string
  highlights: string[]
  mapUrl: string
  gpxUrl?: string
}

defineProps<Props>()
</script>

<template>
  <div class="bg-[#FBFBFB] p-4 sm:p-6 rounded-3xl border border-gray-300 max-w-prose border-[0.5px] w-full">
    <h2 class="text-lg sm:text-xl font-semibold text-gray-800 mb-4">{{ title }}</h2>
    
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
      <div class="bg-gray-50 p-4 rounded-lg text-center">
        <h3 class="text-gray-600 mb-1 text-sm">Distance</h3>
        <p class="text-lg font-semibold text-gray-800">{{ distance }}</p>
      </div>
      <div class="bg-gray-50 p-4 rounded-lg text-center">
        <h3 class="text-gray-600 mb-1 text-sm">Elevation</h3>
        <p class="text-lg font-semibold text-gray-800">{{ elevation }}</p>
      </div>
      <div class="bg-gray-50 p-4 rounded-lg text-center">
        <h3 class="text-gray-600 mb-1 text-sm">Difficulty</h3>
        <p class="text-lg font-semibold text-gray-800">{{ difficulty }}</p>
      </div>
    </div>

    <div class="mb-6">
      <h3 class="text-base font-semibold text-gray-800 mb-2">Description</h3>
      <p class="text-gray-600 text-sm leading-relaxed">{{ description }}</p>
    </div>

    <div class="mb-6">
      <h3 class="text-base font-semibold text-gray-800 mb-2">Highlights</h3>
      <ul class="space-y-2">
        <li v-for="(highlight, index) in highlights" 
            :key="index"
            class="text-gray-600 text-sm border-b border-gray-200 pb-2 last:border-0">
          {{ highlight }}
        </li>
      </ul>
    </div>

    <div class="flex gap-3 justify-center">
      <BaseButton
        title="View on Map"
        :link="mapUrl"
        variant="primary"
        class="text-sm"
      />
      <BaseButton
        title="Download GPX"
        :link="gpxUrl"
        variant="primary"
        class="text-sm"
      />
      <BaseButton
        title="Save Route"
        variant="secondary"
        class="text-sm"
      />
    </div>
  </div>
</template> 