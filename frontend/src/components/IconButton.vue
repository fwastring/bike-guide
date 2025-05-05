<script setup lang="ts">
import type { Component } from 'vue'

interface Props {
  title: string
  icon: Component
  variant?: 'primary' | 'secondary' | 'ghost'
  size?: 'sm' | 'md' | 'lg'
  disabled?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  variant: 'primary',
  size: 'md',
  disabled: false
})

const sizeClasses = {
  sm: 'p-1.5',
  md: 'p-2',
  lg: 'p-3'
}

const variantClasses = {
  primary: 'bg-purple-500 text-white hover:bg-purple-600 focus:ring-purple-500',
  secondary: 'bg-gray-100 text-gray-800 hover:bg-gray-200 focus:ring-gray-500',
  ghost: 'bg-transparent hover:bg-gray-100 text-gray-800 focus:ring-gray-500'
}

const iconSizeClasses = {
  sm: 'h-4 w-4',
  md: 'h-5 w-5',
  lg: 'h-6 w-6'
}
</script>

<template>
  <button
    :disabled="disabled"
    :title="title"
    class="rounded-full flex items-center justify-center transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed"
    :class="[sizeClasses[size], variantClasses[variant]]"
  >
    <component :is="icon" :class="iconSizeClasses[size]" />
  </button>
</template> 