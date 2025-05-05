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
  <div class="bg-[#FBFBFB] p-4 sm:p-6 rounded-2xl border border-gray-300 border-[0.5px] w-full">
    <div class="space-y-1 text-md">
      <div class="flex items-baseline">
        <span class="font-bold text-gray-900">Location: &nbsp;</span>
        <span class="text-gray-900">{{ location }}</span>
      </div>
      
      <div class="flex items-baseline">
        <span class="font-bold text-gray-900">Difficulty: &nbsp;</span>
        <span class="text-gray-900">{{ difficulty }}</span>
      </div>
      
      <div class="flex items-baseline">
        <span class="font-bold text-gray-900">Distance: &nbsp;</span>
        <span class="text-gray-900">{{ distance }}</span>
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