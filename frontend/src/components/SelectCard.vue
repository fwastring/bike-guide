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
  <div class="select-card">
    <h2>{{ title }}</h2>
    <div class="options">
      <div
        v-for="option in options"
        :key="option.id"
        class="option"
        :class="{ selected: selectedOption === option.id }"
        @click="handleSelect(option.id)"
      >
        <h3>{{ option.title }}</h3>
        <p>{{ option.description }}</p>
      </div>
    </div>
    <div class="actions">
      <BaseButton
        title="Continue"
        variant="primary"
        :disabled="!selectedOption"
        @click="$emit('continue')"
      />
    </div>
  </div>
</template>

<style scoped>
.select-card {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  max-width: 600px;
  width: 100%;
}

h2 {
  text-align: center;
  margin-bottom: 1.5rem;
  color: #333;
}

.options {
  display: grid;
  gap: 1rem;
  margin-bottom: 2rem;
}

.option {
  padding: 1rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.option:hover {
  border-color: #007bff;
  background-color: rgba(0, 123, 255, 0.05);
}

.option.selected {
  border-color: #007bff;
  background-color: rgba(0, 123, 255, 0.1);
}

.option h3 {
  margin: 0 0 0.5rem 0;
  color: #333;
}

.option p {
  margin: 0;
  color: #666;
}

.actions {
  display: flex;
  justify-content: center;
}
</style> 