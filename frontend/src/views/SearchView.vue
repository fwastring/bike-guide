<script setup lang="ts">
import { ref } from 'vue'
import SearchCard from '../components/SearchCard.vue'
import MapCard from '../components/MapCard.vue'
import Header from '../components/Header.vue'
import { useSearchStore } from '@/stores/searchStore'

const searchStore = useSearchStore()

const handleSearch = (searchData: { address: string; difficulty: string; distance: string }) => {
  searchStore.setSearchParams(searchData)
}
</script>

<template>
  <div class="min-h-screen flex max-w-7xl mx-auto p-6 md:space-x-6">
    
    <!-- Search Panel -->
    <div class="w-full md:w-1/3">
      <SearchCard @search="handleSearch" />
    </div>

    <!-- Map Panel -->
    <div class="w-2/3 hidden md:block">
      <MapCard />
    </div>
  </div>
</template> 