<script setup lang="ts">
interface Props {
  label?: string
  type?: string
  placeholder?: string
  modelValue?: string
  error?: string
  mobile?: boolean
  errorBorder?: boolean
  compact?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  type: 'text',
  placeholder: '',
  modelValue: '',
  error: '',
  mobile: false,
  errorBorder: false,
  compact: false
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
      :class="[
        'w-full border-[0.5px] bg-white border text-gray-800 placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent rounded-full',
        error ? 'border-red-500' : 'border-gray-300',
        compact ? 'px-4 py-2 text-sm h-10' : mobile
          ? 'px-3 py-2 text-sm h-10'
          : 'px-6 pr-16 py-4 text-base h-14 border-2',
        errorBorder ? 'ring-1 outline-none ring-blue-500 ' : ''
      ]"
    />
    <span v-if="error" class="text-sm text-red-500">{{ error }}</span>
  </div>
</template> 