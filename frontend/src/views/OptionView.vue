<template>
  <div class="min-h-screen bg-gray-50">
    <Header />
    <main class="container mx-auto px-4 py-8">
      <div class="max-w-7xl mx-auto">
        <!-- Search Summary -->
        <div class="mb-8">
          <SummaryCard
            :location="searchStore.searchParams.address"
            :difficulty="searchStore.searchParams.difficulty"
            :distance="searchStore.searchParams.distance"
            @edit="handleEdit"
          />
        </div>

        <!-- Loading State -->
        <div v-if="searchStore.isLoading" class="flex justify-center items-center py-12">
          <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
        </div>

        <!-- Error State -->
        <div v-else-if="searchStore.error" class="text-center py-12">
          <p class="text-red-600">{{ searchStore.error }}</p>
          <button 
            @click="searchStore.fetchRouteOptions()"
            class="mt-4 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
          >
            Try Again
          </button>
        </div>

        <!-- Route Options -->
        <div v-else class="mb-8">
          <h2 class="text-2xl font-bold text-gray-800 mb-4">Available Routes</h2>
          <div v-if="searchStore.routeOptions.length === 0" class="text-center py-12">
            <p class="text-gray-600">No routes found matching your criteria.</p>
          </div>
          <OptionListCard
            v-else
            :options="searchStore.routeOptions"
            @expand="handleExpand"
          />
        </div>
      </div>
    </main>
    <Footer />
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useSearchStore } from '@/stores/searchStore'
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'
import SummaryCard from '../components/SummaryCard.vue'
import OptionListCard from '../components/OptionListCard.vue'

const router = useRouter()
const searchStore = useSearchStore()

onMounted(async () => {
  // If no search params are set, redirect back to search
  if (!searchStore.searchParams.address) {
    router.push('/search')
    return
  }
  
  await searchStore.fetchRouteOptions()
})

const handleEdit = () => {
  router.push('/search')
}

const handleExpand = (index: number) => {
  router.push(`/route/${index}`)
}
</script> 