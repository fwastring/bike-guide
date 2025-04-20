<script setup lang="ts">
import { ref } from 'vue'
import SearchCard from '../components/SearchCard.vue'
import MapCard from '../components/MapCard.vue'
import Header from '../components/Header.vue'

const handleSearch = (searchData: { address: string; difficulty: string; distance: string }) => {
  console.log('Searching with:', searchData)
}
</script>

<template>
<Header />
  <div class="flex max-w-7xl mx-auto p-6 md:space-x-6">
    
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