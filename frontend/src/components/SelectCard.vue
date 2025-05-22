<script setup lang="ts">
import { ref } from 'vue'
import BaseButton from './BaseButton.vue'

interface Option {
  id: string
  title: string
  description: string
}

interface Props {
  title: string
  options: Option[]
}

const props = defineProps<Props>()
const selectedOption = ref<string | null>(null)

const emit = defineEmits(['select'])

const handleSelect = (optionId: string) => {
  selectedOption.value = optionId
  emit('select', optionId)
}
</script>

<template>
  <div class="bg-white p-8 rounded-lg shadow-md max-w-2xl w-full" role="region" aria-label="Selection interface">
    <h2 class="text-2xl font-bold text-center text-gray-800 mb-6" id="selection-title">{{ title }}</h2>
    <div class="grid gap-4 mb-8" role="radiogroup" aria-labelledby="selection-title">
      <div
        v-for="option in options"
        :key="option.id"
        class="p-4 border rounded-md cursor-pointer transition-colors"
        :class="{
          'border-[#8E7DBE] bg-[#8E7DBE]/10': selectedOption === option.id,
          'border-gray-200 hover:border-[#8E7DBE] hover:bg-[#8E7DBE]/10': selectedOption !== option.id
        }"
        @click="handleSelect(option.id)"
        role="radio"
        :aria-checked="selectedOption === option.id"
        :aria-label="option.title"
        tabindex="0"
        @keydown.enter="handleSelect(option.id)"
        @keydown.space="handleSelect(option.id)"
      >
        <h3 class="font-semibold text-gray-800 mb-1">{{ option.title }}</h3>
        <p class="text-gray-600">{{ option.description }}</p>
      </div>
    </div>
    <div class="flex justify-center">
      <BaseButton
        title="Continue"
        variant="primary"
        :disabled="!selectedOption"
        @click="$emit('select')"
        :aria-label="selectedOption ? 'Continue with selected option' : 'Please select an option to continue'"
      />
    </div>
  </div>
</template> 