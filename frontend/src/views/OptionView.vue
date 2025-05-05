<template>
  <div class="h-screen flex flex-col">
    <main class="flex-1 overflow-y-auto bg-gray-50">
      <div class="flex flex-col md:flex-row max-w-7xl mx-auto p-6 md:space-x-6">
        <!-- Left Panel: Summary -->
        <div class="w-full md:w-1/3">
          <!-- Search Summary -->
          <div class="mb-8">
            <SummaryCard :location="searchStore.searchParams.address" :difficulty="searchStore.searchParams.difficulty"
              :distance="searchStore.searchParams.distance" @edit="handleEdit" />
          </div>

          <!-- Loading State -->
          <div v-if="searchStore.isLoading" class="flex justify-center items-center py-12">
            <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
          </div>

          <!-- Error State -->
          <div v-else-if="searchStore.error" class="text-center py-12">
            <p class="text-red-600 text-sm">{{ searchStore.error }}</p>
            <button @click="searchStore.fetchRouteOptions()"
              class="mt-4 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 text-sm font-medium">
              Try Again
            </button>
          </div>

          <!-- Map (hidden on mobile) -->
          <div class="hidden md:block">
            <MapCard />
          </div>
        </div>

        <!-- Right Panel: Route Options (full width on mobile) -->
        <div class="w-full md:w-2/3">
          <div class="mb-8"> 
            <div v-if="searchStore.routeOptions.length === 0" class="text-center py-12">
              <p class="text-sm text-gray-600">No routes found matching your criteria.</p>
            </div>
            <OptionListCard v-else :options="searchStore.routeOptions" @expand="handleExpand" />
          </div>
        </div>
      </div>
      <Footer />
    </main>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useSearchStore } from '@/stores/searchStore'
import Footer from '../components/Footer.vue'
import SummaryCard from '../components/SummaryCard.vue'
import OptionListCard from '../components/OptionListCard.vue'
import MapCard from '../components/MapCard.vue'


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