<script setup lang="ts">
import BaseButton from './BaseButton.vue'
import IconButton from './IconButton.vue'
import { MessageSquare } from 'lucide-vue-next'

interface Props {
  location: string
  difficulty: string
  distance: string
}

defineProps<Props>()
const emit = defineEmits(['edit'])
</script>

<template>
  <div class="w-full">
    <!-- Large Address Display -->
    <h1 class="font-pramukh text-4xl sm:text-5xl lg:text-6xl text-gray-900 mb-6">
      {{ location }}
    </h1>

    <!-- Info Boxes -->
    <div class="flex flex-wrap gap-4 mb-6">
      <div class="bg-gray-100 px-4 py-2 rounded-full">
        <span class="text-sm font-medium text-gray-600">Difficulty</span>
        <span class="ml-2 text-gray-900">{{ difficulty }}</span>
      </div>
      
      <div class="bg-gray-100 px-4 py-2 rounded-full">
        <span class="text-sm font-medium text-gray-600">Distance</span>
        <span class="ml-2 text-gray-900">{{ distance }}</span>
      </div>
    </div>

  </div>
</template>

<style scoped>
.font-pramukh {
  font-family: 'Pramukh Rounded', sans-serif;
}
</style> 