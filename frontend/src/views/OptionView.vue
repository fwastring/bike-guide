<template>
  <div class="h-screen flex flex-col">
    <main class="flex-1 overflow-y-auto bg-gray-50">
        <!-- Mini Search Card at the top -->
      <div class="w-full bg-white py-4 shadow-sm">
        <div class="max-w-lg mx-auto px-6">
          <MiniSearchCard @search="handleSearch" :mobile="true" />
        </div>
      </div>
      
      <div class="max-w-3xl mx-auto px-6 py-8 space-y-8">
        <!-- Search Summary -->
        <div>
          <SummaryCard 
            :location="searchStore.searchParams.address" 
            :difficulty="searchStore.searchParams.difficulty"
            :distance="searchStore.searchParams.distance" 
            @edit="handleEdit" 
          />
        </div>

        <!-- Route Options -->
        <div>
          <div v-if="searchStore.routeOptions.length === 0" class="text-center py-12">
            <p class="text-sm text-gray-600">No routes found matching your criteria.</p>
          </div>
          <OptionListCard v-else :options="searchStore.routeOptions" @expand="handleExpand" />
        </div>
      </div>
      
      <Footer class="mb-20 sm:mb-0" />
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
import MiniSearchCard from '../components/MiniSearchCard.vue'

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

const handleSearch = async (searchData: { address: string; difficulty: string; distance: string }) => {
  searchStore.setSearchParams(searchData)
  await searchStore.fetchRouteOptions()
}
</script>