<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import Footer from '../components/Footer.vue'
import ResultCard from '../components/ResultCard.vue'

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

    <main class="flex-1 p-8 bg-gray-50">
      <div class="max-w-3xl mx-auto">
        <ResultCard
          :title="routeDetails.title"
          :distance="routeDetails.distance"
          :elevation="routeDetails.elevation"
          :difficulty="routeDetails.difficulty"
          :description="routeDetails.description"
          :highlights="routeDetails.highlights"
          :mapUrl="routeDetails.mapUrl"
        />
      </div>
    </main>
    <Footer />
  </div>
</template>

