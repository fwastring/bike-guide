<script setup lang="ts">
import { RouterLink } from 'vue-router'

interface Props {
  title: string
  link?: string
  variant?: 'primary' | 'secondary' | 'tertiary'
}

const props = withDefaults(defineProps<Props>(), {
  variant: 'primary'
})

const emit = defineEmits(['click'])
</script>

<template>
  <RouterLink
    v-if="link"
    :to="link"
    class="button"
    :class="variant"
  >
    {{ title }}
  </RouterLink>
  <button
    v-else
    class="button"
    :class="variant"
    @click="emit('click')"
  >
    {{ title }}
  </button>
</template>

<style scoped>
.button {
  padding: 0.75rem 1.5rem;
  border-radius: 4px;
  font-weight: 600;
  text-decoration: none;
  transition: all 0.3s ease;
  cursor: pointer;
  border: none;
  display: inline-block;
}

.primary {
  background-color: #007bff;
  color: white;
}

.primary:hover {
  background-color: #0056b3;
}

.secondary {
  background-color: #6c757d;
  color: white;
}

.secondary:hover {
  background-color: #5a6268;
}

.tertiary {
  background-color: transparent;
  color: #007bff;
  border: 1px solid #007bff;
}

.tertiary:hover {
  background-color: rgba(0, 123, 255, 0.1);
}
</style> 