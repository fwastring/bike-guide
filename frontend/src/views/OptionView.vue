<template>
  <div class="h-screen flex flex-col">
    <main class="flex-1 overflow-y-auto bg-gray-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-10 lg:px-14 py-8 space-y-8">
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
  await searchStore.fetchRouteOptions()
})

const handleEdit = () => {
  router.push('/search')
}

const handleExpand = (index: number) => {
  router.push(`/route/${index}`)
}

</script>