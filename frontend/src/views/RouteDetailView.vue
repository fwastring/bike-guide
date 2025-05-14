<template>
  <div class="min-h-screen flex flex-col">
    <main class="flex-1 p-8 bg-gray-50">
      <div class="max-w-3xl mx-auto">
        <button
          @click="router.back()"
          class="mb-6 flex items-center text-[#8E7DBE] hover:text-[#7D6CAE]"
        >
          <span class="mr-2">←</span> Back to routes
        </button>
        <template v-if="currentRoute">
          <ResultCard
            :title="currentRoute.title"
            subtitle="Route Details"
            :distance="currentRoute.distance"
            elevationGain="—"
            avgGrade="—"
            lowestElev="—"
            highestElev="—"
            elevDifference="—"
            attempts="0 attempts"
            stars="0 stars"
            mapImageUrl="/placeholder-route.jpg"
            elevationChartUrl="/placeholder-elevation.jpg"
            :mapUrl="currentRoute.imageUrl || ''"
            gpxUrl=""
          />
        </template>
        <div v-else class="text-center py-12">
          <p class="text-gray-600">Route not found.</p>
          <button
            @click="router.push('/options')"
            class="mt-4 text-[#8E7DBE] hover:text-[#7D6CAE]"
          >
            Return to route options
          </button>
        </div>
      </div>
    </main>
    <Footer />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useSearchStore } from '@/stores/searchStore'
import Footer from '../components/Footer.vue'
import ResultCard from '../components/ResultCard.vue'

const router = useRouter()
const route = useRoute()
const searchStore = useSearchStore()

const routeId = parseInt(route.params.id as string)
const currentRoute = computed(() => searchStore.routeOptions[routeId])
</script> 