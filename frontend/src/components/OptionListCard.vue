<script setup lang="ts">
import OptionCard from './OptionCard.vue'

interface RouteOption {
  title: string
  description: string
  distance: string
  duration: string
  difficulty: 'Easy' | 'Medium' | 'Hard'
  imageUrl: string
}

interface Props {
  options: RouteOption[]
}

defineProps<Props>()
const emit = defineEmits(['expand'])

const handleExpand = (index: number) => {
  emit('expand', index)
}
</script>

<template>
  <div class="bg-[#FBFBFB] border border-gray-300 border-[0.5px] rounded-3xl w-full max-w-2xl">
    <div class="divide-y divide-gray-200">
      <div 
        v-for="(option, index) in options" 
        :key="index"
        class="first:rounded-t-3xl last:rounded-b-3xl overflow-hidden"
      >
        <OptionCard
          :title="option.title"
          :description="option.description"
          :distance="option.distance"
          :duration="option.duration"
          :difficulty="option.difficulty"
          :imageUrl="option.imageUrl"
          @expand="() => handleExpand(index)"
          class="border-none rounded-none"
        />
      </div>
    </div>
  </div>
</template> 