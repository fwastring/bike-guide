<script setup lang="ts">
import { Expand, CheckIcon } from 'lucide-vue-next'
import IconButton from './IconButton.vue'

interface Props {
  title: string
  description: string
  distance: string
  duration: string
  difficulty: 'Easy' | 'Medium' | 'Hard'
  imageUrl: string
}

defineProps<Props>()
const emit = defineEmits(['expand'])

const getDifficultyColor = (difficulty: string) => {
  switch (difficulty) {
    case 'Easy':
      return 'bg-[#ECFDF3] text-[#027A48]'
    case 'Medium':
      return 'bg-[#FFF6ED] text-[#B93815]'
    case 'Hard':
      return 'bg-[#FEF3F2] text-[#B42318]'
    default:
      return 'bg-gray-100 text-gray-600'
  }
}
</script>

<template>
  <div class="p-4 sm:p-6 rounded-3xl w-full">
    <div class="flex flex-col sm:flex-row gap-4 sm:gap-8">
      <!-- Image -->
      <!-- <div class="w-full aspect-[4/3] sm:w-32 sm:h-32 h-auto rounded-2xl bg-gray-200 overflow-hidden flex-shrink-0"> -->
      <!--   <img  -->
      <!--     :src="imageUrl"  -->
      <!--     :alt="title" -->
      <!--     class="w-full h-full object-cover" -->
      <!--   /> -->
      <!-- </div> -->

      <!-- Content -->
      <div class="flex-1 flex flex-col justify-between">
        <div>
          <div class="flex items-center justify-between mb-2">
            <h3 class="text-xl font-semibold text-gray-800">{{ title }}</h3>
            <div :class="['px-3 py-1 rounded-full text-sm font-medium', getDifficultyColor(difficulty)]">
              {{ difficulty }}
            </div>
          </div>
          <p class="text-sm text-gray-600 mb-4">{{ description }}</p>
        </div>

        <!-- Stats and Expand Button -->
        <div class="flex flex-col sm:flex-row sm:items-center gap-4 sm:gap-8 mt-2">
          <div class="flex flex-row gap-8 sm:gap-8 w-full">
            <div>
              <p class="text-sm text-gray-600 mb-1">Distance</p>
              <p class="text-lg font-semibold text-gray-800">{{ distance }}</p>
            </div>
            <div>
              <p class="text-sm text-gray-600 mb-1">Time</p>
              <p class="text-lg font-semibold text-gray-800">{{ duration }}</p>
            </div>
          </div>
          <!-- Expand Button -->
          <button
            @click="emit('expand')"
            class="inline-flex border border-gray-300 rounded-full px-4 py-2 items-center gap-2 text-[#8E7DBE] hover:text-[#7D6CAE] transition-colors w-full sm:w-auto justify-center sm:ml-auto mt-2 sm:mt-4"
          >
            <span class="text-sm font-medium">Choose</span>
            <CheckIcon class="w-4 h-4" />
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
