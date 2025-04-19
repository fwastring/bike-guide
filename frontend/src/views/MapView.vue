<template>
  <div class="map-container">
    <Sidebar />
    <div class="map-wrapper">
      <l-map
        ref="map"
        :zoom="13"
        :center="[55.7047, 13.1910]"
        :options="{ 
          zoomControl: false, 
          minZoom: 4,
          maxBounds: [
            [55.0, 10.0],  // Southwest corner (southern Sweden)
            [69.0, 24.0]   // Northeast corner (northern Sweden)
          ],
          maxBoundsViscosity: 1.0
        }"
        style="height: 100%; width: 100%"
      >
        <l-tile-layer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
          layer-type="base"
          name="OpenStreetMap"
        ></l-tile-layer>
        <l-control-zoom position="topright"></l-control-zoom>
      </l-map>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { LMap, LTileLayer, LControlZoom } from '@vue-leaflet/vue-leaflet'
import * as L from 'leaflet'
import 'leaflet/dist/leaflet.css'
import Sidebar from '@/components/Sidebar.vue'

const mapRef = ref<L.Map | null>(null)

// Fix for default marker icons
onMounted(() => {
  const iconRetinaUrl = 'https://unpkg.com/leaflet@1.7.1/dist/images/marker-icon-2x.png'
  const iconUrl = 'https://unpkg.com/leaflet@1.7.1/dist/images/marker-icon.png'
  const shadowUrl = 'https://unpkg.com/leaflet@1.7.1/dist/images/marker-shadow.png'
  const iconDefault = L.icon({
    iconRetinaUrl,
    iconUrl,
    shadowUrl,
    iconSize: [25, 41],
    iconAnchor: [12, 41],
    popupAnchor: [1, -34],
    tooltipAnchor: [16, -28],
    shadowSize: [41, 41]
  })
  L.Marker.prototype.options.icon = iconDefault

  // Set up zoom constraints
  if (mapRef.value) {
    mapRef.value.setMinZoom(4)
  }
})
</script>

<style>
.map-container {
  height: 100vh;
  width: 100%;
  position: relative;
}

.map-wrapper {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 1;
}

/* Ensure the map tiles are visible */
.leaflet-container {
  height: 100% !important;
  width: 100% !important;
  background-color: #fff;
}
</style> 