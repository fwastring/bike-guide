<script setup lang="ts">
import BaseButton from './BaseButton.vue'

interface Props {
  title: string
  subtitle: string
  distance: string
  elevationGain: string
  avgGrade: string
  lowestElev: string
  highestElev: string
  elevDifference: string
  attempts: string
  stars: string
  mapImageUrl: string
  elevationChartUrl: string
  mapUrl: string
  gpxUrl: string
}

defineProps<Props>()
</script>

<template>
  <div class="bg-[#FBFBFB] p-4 sm:p-6 rounded-3xl border border-gray-300 max-w-prose border-[0.5px] w-full">
    <h2 class="text-lg sm:text-xl font-semibold text-gray-800 mb-1">{{ title }}</h2>
    <div class="text-gray-500 text-sm mb-4">{{ subtitle }}</div>

    <div class="grid grid-cols-2 sm:grid-cols-6 gap-2 sm:gap-4 mb-2">
      <div class="text-center">
        <h3 class="text-gray-600 mb-1 text-xs">Distance</h3>
        <p class="text-base font-semibold text-gray-800">{{ distance }}</p>
      </div>
      <div class="text-center">
        <h3 class="text-gray-600 mb-1 text-xs">Elevation Gain</h3>
        <p class="text-base font-semibold text-gray-800">{{ elevationGain }}</p>
      </div>
      <div class="text-center">
        <h3 class="text-gray-600 mb-1 text-xs">Avg Grade</h3>
        <p class="text-base font-semibold text-gray-800">{{ avgGrade }}</p>
      </div>
      <div class="text-center">
        <h3 class="text-gray-600 mb-1 text-xs">Lowest Elev</h3>
        <p class="text-base font-semibold text-gray-800">{{ lowestElev }}</p>
      </div>
      <div class="text-center">
        <h3 class="text-gray-600 mb-1 text-xs">Highest Elev</h3>
        <p class="text-base font-semibold text-gray-800">{{ highestElev }}</p>
      </div>
      <div class="text-center">
        <h3 class="text-gray-600 mb-1 text-xs">Elev Difference</h3>
        <p class="text-base font-semibold text-gray-800">{{ elevDifference }}</p>
      </div>
    </div>
    <div class="flex flex-col sm:flex-row justify-between items-center text-xs text-gray-500 mb-4 gap-1">
      <span>{{ attempts }}</span>
      <span>{{ stars }}</span>
    </div>

    <div class="mb-4">
      <img :src="mapImageUrl" alt="Map segment" class="rounded-xl w-full border border-gray-200" />
    </div>
    <div class="mb-2">
      <img :src="elevationChartUrl" alt="Elevation chart" class="w-full h-24 object-contain" />
    </div>
    <div class="flex gap-3 justify-center mt-4">
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