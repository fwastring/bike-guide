<script setup lang="ts">
import { ref } from 'vue'
import IconButton from './IconButton.vue'
import { MapPin } from 'lucide-vue-next'

const emit = defineEmits(['location'])
const loading = ref(false)

const handleClick = async () => {
  loading.value = true
  try {
    const position = await new Promise<GeolocationPosition>((resolve, reject) => {
      navigator.geolocation.getCurrentPosition(resolve, reject)
    })
    
    emit('location', {
      lat: position.coords.latitude,
      lng: position.coords.longitude
    })
  } catch (error) {
    console.error('Error getting location:', error)
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <IconButton
    title="Use current location"
    :icon="MapPin"
    variant="ghost"
    size="sm"
    :disabled="loading"
    @click="handleClick"
  />
</template> 