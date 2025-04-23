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
  <div class="bg-[#FBFBFB] p-4 sm:p-6 rounded-3xl border border-gray-300 border-[0.5px] w-full">
    <div class="space-y-2">
      <div class="space-y-1">
        <div class="flex items-center justify-between">
          <span class="text-sm text-gray-600">Location:</span>
        </div>
        <p class="text-lg font-semibold text-gray-800">{{ location }}</p>
      </div>
      
      <div class="space-y-1">
        <span class="text-sm text-gray-600">Difficulty:</span>
        <p class="text-lg font-semibold text-gray-800">{{ difficulty }}</p>
      </div>
      
      <div class="space-y-1">
        <span class="text-sm text-gray-600">Distance:</span>
        <p class="text-lg font-semibold text-gray-800">{{ distance }}</p>
      </div>
    </div>

    <div class="flex justify-between items-center mt-6">
      <BaseButton
        title="Edit"
        variant="secondary"
        class="w-full flex items-center justify-center px-4 py-2 text-sm font-medium mr-3"
        @click="emit('edit')"
      />
      
    </div>
  </div>
</template> 