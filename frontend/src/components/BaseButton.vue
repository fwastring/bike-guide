<script setup lang="ts">
import { RouterLink } from 'vue-router'

interface Props {
  title: string
  link?: string
  variant?: 'primary' | 'secondary' | 'tertiary'
  type?: 'button' | 'submit'
  disabled?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  variant: 'primary',
  type: 'button',
  disabled: false
})

const emit = defineEmits(['click'])
</script>

<template>
  <RouterLink
    v-if="link"
    :to="link"
    class="inline-block px-8 py-3 rounded-full font-medium text-center transition-all duration-200"
    :class="{
      'bg-purple-500 text-white hover:bg-purple-600': variant === 'primary',
      'bg-blue-50 text-purple-500 hover:bg-blue-100': variant === 'secondary',
      'text-purple-500 hover:text-purple-600 underline px-2 py-1': variant === 'tertiary'
    }"
  >
    {{ title }}
  </RouterLink>
  <button
    v-else
    :type="type"
    :disabled="disabled"
    class="inline-block px-8 py-3 rounded-full font-medium text-center transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed"
    :class="{
      'bg-purple-500 text-white hover:bg-purple-600': variant === 'primary',
      'bg-blue-50 text-purple-500 hover:bg-blue-100': variant === 'secondary',
      'text-purple-500 hover:text-purple-600 underline px-2 py-1': variant === 'tertiary'
    }"
    @click="emit('click')"
  >
    {{ title }}
  </button>
</template> 