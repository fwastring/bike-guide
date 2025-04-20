<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'
import BaseButton from '../components/BaseButton.vue'

const route = useRoute()
const routeDetails = ref({
  title: 'Scenic Coastal Route',
  distance: '20km',
  elevation: '150m',
  difficulty: 'Moderate',
  description: 'A beautiful ride along the coast with stunning ocean views. Perfect for intermediate cyclists looking for a scenic route.',
  highlights: [
    'Stunning coastal views',
    'Well-maintained paths',
    'Multiple rest stops',
    'Café at the halfway point'
  ],
  mapUrl: 'https://maps.example.com/route/1'
})

onMounted(() => {
  // In a real app, we would fetch the route details based on route.query.routeId
  console.log('Route ID:', route.query.routeId)
})
</script>

<template>
  <div class="min-h-screen flex flex-col">
    <Header />
    <main class="flex-1 p-8 bg-gray-50">
      <div class="max-w-3xl mx-auto bg-white p-8 rounded-lg shadow-md">
        <h1 class="text-3xl font-bold text-gray-800 mb-8">{{ routeDetails.title }}</h1>
        
        <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-8">
          <div class="bg-gray-50 p-4 rounded-lg text-center">
            <h3 class="text-gray-600 mb-1">Distance</h3>
            <p class="text-xl font-semibold text-gray-800">{{ routeDetails.distance }}</p>
          </div>
          <div class="bg-gray-50 p-4 rounded-lg text-center">
            <h3 class="text-gray-600 mb-1">Elevation</h3>
            <p class="text-xl font-semibold text-gray-800">{{ routeDetails.elevation }}</p>
          </div>
          <div class="bg-gray-50 p-4 rounded-lg text-center">
            <h3 class="text-gray-600 mb-1">Difficulty</h3>
            <p class="text-xl font-semibold text-gray-800">{{ routeDetails.difficulty }}</p>
          </div>
        </div>

        <div class="mb-8">
          <h2 class="text-2xl font-semibold text-gray-800 mb-4">Description</h2>
          <p class="text-gray-600 leading-relaxed">{{ routeDetails.description }}</p>
        </div>

        <div class="mb-8">
          <h2 class="text-2xl font-semibold text-gray-800 mb-4">Highlights</h2>
          <ul class="space-y-2">
            <li v-for="(highlight, index) in routeDetails.highlights" 
                :key="index"
                class="text-gray-600 border-b border-gray-200 pb-2 last:border-0">
              {{ highlight }}
            </li>
          </ul>
        </div>

        <div class="flex gap-4 justify-center">
          <BaseButton
            title="View on Map"
            :link="routeDetails.mapUrl"
            variant="primary"
          />
          <BaseButton
            title="Save Route"
            variant="secondary"
          />
        </div>
      </div>
    </main>
    <Footer />
  </div>
</template>

<style scoped>
.result {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.main-content {
  flex: 1;
  padding: 2rem;
  background-color: #f5f5f5;
}

.route-details {
  max-width: 800px;
  margin: 0 auto;
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

h1 {
  margin-bottom: 2rem;
  color: #333;
}

.route-stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 1rem;
  margin-bottom: 2rem;
}

.stat {
  text-align: center;
  padding: 1rem;
  background-color: #f8f9fa;
  border-radius: 4px;
}

.stat h3 {
  margin: 0 0 0.5rem 0;
  color: #666;
}

.stat p {
  margin: 0;
  font-size: 1.25rem;
  font-weight: bold;
  color: #333;
}

.description,
.highlights {
  margin-bottom: 2rem;
}

h2 {
  margin-bottom: 1rem;
  color: #333;
}

.highlights ul {
  list-style-type: none;
  padding: 0;
}

.highlights li {
  padding: 0.5rem 0;
  border-bottom: 1px solid #eee;
}

.highlights li:last-child {
  border-bottom: none;
}

.actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
}
</style> 