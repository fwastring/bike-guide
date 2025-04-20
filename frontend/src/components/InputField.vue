<script setup lang="ts">
interface Props {
  label?: string
  type?: string
  placeholder?: string
  modelValue?: string
  error?: string
}

const props = withDefaults(defineProps<Props>(), {
  type: 'text',
  placeholder: '',
  modelValue: '',
  error: ''
})

const emit = defineEmits(['update:modelValue'])
</script>

<template>
  <div class="flex flex-col gap-2 w-full">
    <label v-if="label" class="font-semibold text-gray-800">{{ label }}</label>
    <input
      :type="type"
      :placeholder="placeholder"
      :value="modelValue"
      @input="emit('update:modelValue', ($event.target as HTMLInputElement).value)"
      class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-colors"
      :class="{ 'border-red-500': error }"
    />
    <span v-if="error" class="text-sm text-red-500">{{ error }}</span>
  </div>
</template> 