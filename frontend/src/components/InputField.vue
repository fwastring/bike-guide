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
  <div class="flex flex-col gap-1.5 w-full">
    <label v-if="label" class="text-sm sm:text-base font-medium text-gray-800">{{ label }}</label>
    <input
      :type="type"
      :placeholder="placeholder"
      :value="modelValue"
      @input="emit('update:modelValue', ($event.target as HTMLInputElement).value)"
      class="w-full px-3 py-2 bg-white border border-gray-300 rounded-full text-sm text-gray-800 placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent border-[0.5px]"
      :class="{ 'border-red-500': error }"
    />
    <span v-if="error" class="text-sm text-red-500">{{ error }}</span>
  </div>
</template> 