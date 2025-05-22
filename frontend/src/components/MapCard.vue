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
      <l-polyline v-if="coordsForDisplay && coordsForDisplay.length > 0" color="red" :lat-lngs="coordsForDisplay" :key="polylineKey" />
        <l-marker v-if="marker" :lat-lng="marker"></l-marker>
    </l-map>
  </div>
</template>

<script setup lang="ts">
  import { ref, onMounted, computed, watch } from 'vue'
import { LMap, LTileLayer, LControlZoom, LPolyline, LMarker } from '@vue-leaflet/vue-leaflet'
import { type Poi, type RouteResponse } from '@/components/ResultCard.vue' // This now uses the GeoJSON types
import * as L from 'leaflet'
import 'leaflet/dist/leaflet.css'
import type { Position } from 'geojson'; // For clarity on coordinate types
import 'leaflet/dist/leaflet.css'

const props = defineProps<{
  response: RouteResponse,
  markers: Poi[]
}>()

const marker = computed(() => {
  if (
    props.markers && props.markers.length > 0
  ) {
    return props.markers.map(marker => [marker.Lat, marker.Lon] as [number, number])[0]

  }
})


const zoom = ref(13)
const center = ref<[number, number]>([55.7047, 13.1910])
const map = ref<L.Map | null>(null)
const polylineKey = ref(0); // To force re-render of LPolyline if needed

const coordsForDisplay = computed(() => {
  console.log(props.response)
  if (
    props.response &&
    props.response.response.geometry.coordinates.length > 0
  ) {
    const geoJsonCoordinates = props.response.response.geometry.coordinates as Position[];

    // Swap to [lat, lng] for Leaflet
    const leafletCoordinates = geoJsonCoordinates.map(coord => [coord[1], coord[0]] as [number, number]);

    console.log('Coordinates passed to LPolyline (should be [lat,lng]):', JSON.stringify(leafletCoordinates));
    return leafletCoordinates;

  }
  return []
})

// ... (rest of your MapCard logic, including the watch for 'coords' to call map.fitBounds)
// Remember that map.fitBounds typically expects [latitude, longitude]
// so the swapping logic in your watch function remains important:
watch(coordsForDisplay, (newCoords) => { // newCoords are now already [lat, lng]
  if (map.value && newCoords && newCoords.length > 0) {
    // No need to swap again here, as coordsForDisplay already provides [lat, lng]
    map.value.fitBounds(newCoords as L.LatLngBoundsExpression);
    polylineKey.value++;
  }
}, { deep: true })



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
