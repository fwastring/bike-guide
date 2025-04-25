<template>
  <div class="map-container border-[0.5px] border-gray-300 rounded-2xl">
    <l-map
      v-model:zoom="zoom"
      :center="center"
      :options="mapOptions"
      @ready="onMapReady"
      class="map-view"
    >
      <l-tile-layer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        :attribution="attribution"
      ></l-tile-layer>
      <l-control-zoom position="topright"></l-control-zoom>
    </l-map>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { LMap, LTileLayer, LControlZoom } from '@vue-leaflet/vue-leaflet'
import * as L from 'leaflet'
import 'leaflet/dist/leaflet.css'

const zoom = ref(13)
const center = ref<[number, number]>([55.7047, 13.1910])
const map = ref<L.Map | null>(null)

const attribution = '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
const mapOptions = {
  zoomControl: false,
  minZoom: 4,
  maxBounds: [
    [55.0, 10.0], // Southwest corner (southern Sweden)
    [69.0, 24.0]  // Northeast corner (northern Sweden)
  ],
  maxBoundsViscosity: 1.0
}

// Custom icon configuration
const customIcon = L.icon({
  iconUrl: 'https://raw.githubusercontent.com/pointhi/leaflet-color-markers/master/img/marker-icon-2x-blue.png',
  shadowUrl: 'https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/images/marker-shadow.png',
  iconSize: [25, 41],
  iconAnchor: [12, 41],
  popupAnchor: [1, -34],
  shadowSize: [41, 41]
})

const onMapReady = (mapInstance: L.Map) => {
  map.value = mapInstance
  L.Marker.prototype.options.icon = customIcon
}

onMounted(() => {
  if (map.value) {
    map.value.setMinZoom(4)
  }
})
</script>

<style>
.map-container {
  width: 100%;
  height: 400px;
  position: relative;
}

.map-view {
  width: 100%;
  height: 100%;
  border-radius: 1rem;
  z-index: 0;
}

/* Ensure the map tiles are visible */
.leaflet-container {
  width: 100% !important;
  height: 100% !important;
  z-index: 0;
}

.leaflet-pane {
  z-index: 0;
}

.leaflet-tile-pane {
  z-index: 0;
}

.leaflet-control-container {
  z-index: 1;
}
</style> 