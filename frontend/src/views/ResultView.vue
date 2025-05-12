<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import Footer from '../components/Footer.vue'
import ResultCard from '../components/ResultCard.vue'

const route = useRoute()
const routeDetails = ref({
  title: 'Utslängd från VG',
  subtitle: 'Ride Segment Lund, Skåne County',
  distance: '0.25km',
  elevationGain: '7m',
  avgGrade: '2.9%',
  lowestElev: '34m',
  highestElev: '41m',
  elevDifference: '7m',
  attempts: '27,877 Attempts by 2,772 People',
  stars: 'Starred by 21 People',
  mapImageUrl: '/map-segment.png', // Placeholder, replace with real image or map
  elevationChartUrl: '/elevation-chart.png', // Placeholder, replace with real image or chart
  mapUrl: 'https://maps.example.com/route/1',
  gpxUrl: '/route1.gpx',
})

onMounted(() => {
  // In a real app, we would fetch the route details based on route.query.routeId
  console.log('Route ID:', route.query.routeId)
})
</script>

<template>
  <div class="min-h-screen flex flex-col">

    <main class="flex-1 p-8 bg-gray-50">
      <div class="max-w-3xl ">
        <ResultCard
          :title="routeDetails.title"
          :subtitle="routeDetails.subtitle"
          :distance="routeDetails.distance"
          :elevationGain="routeDetails.elevationGain"
          :avgGrade="routeDetails.avgGrade"
          :lowestElev="routeDetails.lowestElev"
          :highestElev="routeDetails.highestElev"
          :elevDifference="routeDetails.elevDifference"
          :attempts="routeDetails.attempts"
          :stars="routeDetails.stars"
          :mapImageUrl="routeDetails.mapImageUrl"
          :elevationChartUrl="routeDetails.elevationChartUrl"
          :mapUrl="routeDetails.mapUrl"
          :gpxUrl="routeDetails.gpxUrl"
        />
      </div>
    </main>
    <Footer />
  </div>
</template>

