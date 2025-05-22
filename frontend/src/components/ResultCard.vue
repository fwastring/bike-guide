<script setup lang="ts">
import { computed, ref, onMounted } from 'vue'
import BaseButton from './BaseButton.vue'
import MapCard from '@/components/MapCard.vue'
import type { Feature, LineString, GeoJsonProperties, Position } from 'geojson';
import { getPOIs, getRoutes } from '@/services/routeService';
import { useSearchStore } from '@/stores/searchStore';
import { useRoute, useRouter } from 'vue-router';
export type RouteResponse = {
  response: Feature<LineString, GeoJsonProperties>
};

interface Props {
  title: string
  subtitle: string
  distance: string
  elevationGain: string
  avgGrade: string
  lowestElev: string
  highestElev: string
  elevDifference: string
  attempts: string
  stars: string
  mapImageUrl: string
  elevationChartUrl: string
  mapUrl: string
  gpxUrl: string
}

export type Node = {
  ID: number
  Lat: number
  Lon: number
}

export type OverpassResponse = {
  Timestamp: string
  Nodes: Node[]
};

export type Poi = {
  ID: number
  Lat: number
  Lon: number
}

// Initialize routesRef with the correct type and an empty array for routes
// or a default Feature structure if absolutely needed.
// const routesRef = ref<RouteResponse>({ routes: [] });

const poiRef = ref<Poi[]>(
  [{
    ID: 1231,
    Lon: 123123,
    Lat: 12312231
  }]
)
const routesRef = ref<RouteResponse>({
  response: {
    type: "Feature",
    geometry: {
      type: "LineString",
      coordinates: [] as Position[] // Explicitly type empty coordinates array
    },
    properties: {}
  }
});

const router = useRouter()
const route = useRoute()
const searchStore = useSearchStore()

const routeId = parseInt(route.params.id as string)
const currentRoute = computed(() => searchStore.routeOptions[routeId])

async function findRoutes() {
  // The getRoutes service should be typed to return RouteResponse
  const newRoutes: RouteResponse = await getRoutes(currentRoute.value.address, currentRoute.value.distance)
  const pois: OverpassResponse = await getPOIs(currentRoute.value.address)
  routesRef.value = newRoutes
  console.log(newRoutes)
  let allPois = []
  for (let [id, node] of Object.entries(pois.Nodes)) {
    allPois.push(node)
  }
  poiRef.value = allPois
}

defineProps<Props>()

onMounted(async () => {
  await findRoutes()
})
</script>

<template>
  <div class="bg-[#FBFBFB] p-4 sm:p-6 rounded-3xl border border-gray-300 max-w-prose border-[0.5px] w-full">
    <h2 class="text-lg sm:text-xl font-semibold text-gray-800 mb-1">{{ title }}</h2>
    <div class="text-gray-500 text-sm mb-4">{{ subtitle }}</div>

    <div class="grid grid-cols-2 sm:grid-cols-6 gap-2 sm:gap-4 mb-2">
      <div class="text-center">
        <h3 class="text-gray-600 mb-1 text-xs">Distance</h3>
        <p class="text-base font-semibold text-gray-800">{{ distance }}</p>
      </div>
      <div class="text-center">
        <h3 class="text-gray-600 mb-1 text-xs">Elevation Gain</h3>
        <p class="text-base font-semibold text-gray-800">{{ elevationGain }}</p>
      </div>
      <div class="text-center">
        <h3 class="text-gray-600 mb-1 text-xs">Avg Grade</h3>
        <p class="text-base font-semibold text-gray-800">{{ avgGrade }}</p>
      </div>
      <div class="text-center">
        <h3 class="text-gray-600 mb-1 text-xs">Lowest Elev</h3>
        <p class="text-base font-semibold text-gray-800">{{ lowestElev }}</p>
      </div>
      <div class="text-center">
        <h3 class="text-gray-600 mb-1 text-xs">Highest Elev</h3>
        <p class="text-base font-semibold text-gray-800">{{ highestElev }}</p>
      </div>
      <div class="text-center">
        <h3 class="text-gray-600 mb-1 text-xs">Elev Difference</h3>
        <p class="text-base font-semibold text-gray-800">{{ elevDifference }}</p>
      </div>
    </div>
    <div class="flex flex-col sm:flex-row justify-between items-center text-xs text-gray-500 mb-4 gap-1">
      <span>{{ attempts }}</span>
      <span>{{ stars }}</span>
    </div>

    <div class="mb-4">
        <MapCard :response="routesRef" :markers="poiRef"/>
    </div>
    <div class="mb-2">
      <img :src="elevationChartUrl" alt="Elevation chart" class="w-full h-24 object-contain" />
    </div>
    <div class="flex gap-3 justify-center mt-4">
      <BaseButton
        title="View on Map"
        :link="mapUrl"
        variant="primary"
        class="text-sm"
      />
      <BaseButton
        title="Download GPX"
        :link="gpxUrl"
        variant="primary"
        class="text-sm"
      />
      <BaseButton
        title="Save Route"
        variant="secondary"
        class="text-sm"
      />
    </div>
  </div>
</template>
