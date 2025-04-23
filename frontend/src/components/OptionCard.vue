<script setup lang="ts">
import { Expand } from 'lucide-vue-next'
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
  <div class=" p-4 sm:p-6 rounded-3xl  w-full ">
    <!-- Header with difficulty -->
    <div class="flex justify-end mb-4">
      <div :class="['px-3 py-1 rounded-full text-sm font-medium', getDifficultyColor(difficulty)]">
        {{ difficulty }}
      </div>
    </div>

    <div class="flex flex-col sm:flex-row gap-8">
      <!-- Image -->
      <div class="w-full sm:w-32 h-32 rounded-2xl bg-gray-200 overflow-hidden flex-shrink-0">
        <img 
          :src="imageUrl" 
          :alt="title"
          class="w-full h-full object-cover"
        />
      </div>

      <!-- Content -->
      <div class="flex-1">
        <div>
          <h3 class="text-xl font-semibold text-gray-800 mb-1">{{ title }}</h3>
          <p class="text-sm text-gray-600">{{ description }}</p>
        </div>

        <!-- Stats -->
        <div class="mt-4 pt-4">
          <div class="flex flex-col sm:flex-row gap-8 items-start sm:items-center justify-between">
            <div class="grid grid-cols-2 sm:flex gap-8">
              <div>
                <p class="text-sm text-gray-600 mb-1 border-t border-gray-300 pt-2">Distance</p>
                <p class="text-lg font-semibold text-gray-800">{{ distance }}</p>
              </div>
              <div>
                <p class="text-sm text-gray-600 mb-1 border-t border-gray-300 pt-2">Time</p>
                <p class="text-lg font-semibold text-gray-800">{{ duration }}</p>
              </div>
            </div>

            <!-- Expand Button -->
            <button 
              @click="emit('expand')"
              class="inline-flex border border-gray-300 rounded-full px-4 py-2 items-center gap-2 text-[#8E7DBE] hover:text-[#7D6CAE] transition-colors"
            >
              <span class="text-sm font-medium">Expand</span>
              <Expand class="w-4 h-4" />
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template> 