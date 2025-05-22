<script setup lang="ts">
  import { ref, onMounted } from 'vue'
import { RouterView } from 'vue-router'
import Header from './components/Header.vue'
import MapCard from './components/MapCard.vue'
import { getPOIs, getRoutes } from './services/routeService' // Assuming getRoutes returns Promise<RouteResponse>
import { useAuthStore } from './stores/auth'
import { useRouter } from 'vue-router'
// Import GeoJSON types
import type { Feature, LineString, GeoJsonProperties, Position } from 'geojson';

export type RouteResponse = {
  response: Feature<LineString, GeoJsonProperties>
};

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

const authStore = useAuthStore()
const router = useRouter()
const address = ref("")

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


onMounted(async () => {
  await authStore.checkAuth()
})

async function findRoutes() {
  // The getRoutes service should be typed to return RouteResponse
  const newRoutes: RouteResponse = await getRoutes(address.value)
  const pois: OverpassResponse = await getPOIs(address.value)
  routesRef.value = newRoutes
  console.log(newRoutes)
  let allPois = []
  for (let [id, node] of Object.entries(pois.Nodes)) {
    allPois.push(node)
  }
  poiRef.value = allPois
}

const handleLogout = async () => {
  await authStore.logout()
  router.push({ name: 'login' })
}

// Your custom LatLngExpression type from Leaflet (if used elsewhere) might still be relevant,
// but the GeoJSON coordinates themselves are number arrays ([lng, lat]).
// import type { LatLngExpression } from 'leaflet'
// Type LatLngExpression is not directly used in the RouteResponse from GeoJSON.
</script>

<template>
  <div class="bg-[#FBFBFB]">
    <Header />
      <!-- <MapCard :response="routesRef" :markers="poiRef"/> -->
      <!-- <input v-model="address" class="h-8 w-full border-2"/> -->
      <!-- <button @click="findRoutes" class="h-8 w-full border-2"> -->
      <!--   Send it -->
      <!-- </button> -->
    <nav v-if="authStore.isAuthenticated" class="navbar">
      <div class="nav-brand">Bike Guide</div>
      <div class="nav-user">
        <img :src="authStore.user?.picture" :alt="authStore.user?.name" class="user-avatar" />
        <span class="user-name">{{ authStore.user?.name }}</span>
        <button @click="handleLogout" class="logout-btn">Logout</button>
      </div>
    </nav>
    <RouterView />
  </div>
</template>

<style>
#app {
  width: 100%;
  height: 100vh;
  margin: 0;
  padding: 0;
}

.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 2rem;
  background-color: white;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.nav-brand {
  font-size: 1.5rem;
  font-weight: bold;
  color: #333;
}

.nav-user {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.user-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
}

.user-name {
  color: #333;
}

.logout-btn {
  padding: 0.5rem 1rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  background-color: white;
  color: #333;
  cursor: pointer;
  transition: background-color 0.2s;
}

.logout-btn:hover {
  background-color: #f5f5f5;
}
</style>
