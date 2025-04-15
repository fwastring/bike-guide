<template>
  <div class="sidebar bg-white shadow-lg h-full w-80 flex flex-col">
    <!-- Header -->
    <div class="p-4 border-b border-gray-200">
      <h2 class="text-xl font-semibold text-gray-800">Bike Guide</h2>
    </div>

    <!-- Content Area -->
    <div class="flex-1 overflow-y-auto p-4 space-y-4">
      <!-- Search Input -->
      <div class="space-y-2">
        <label for="search" class="block text-sm font-medium text-gray-700">Search</label>
        <input
          type="text"
          id="search"
          v-model="searchQuery"
          class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="Search locations..."
        />
      </div>

      <!-- Filter Section -->
      <div class="space-y-2">
        <h3 class="text-sm font-medium text-gray-700">Filters</h3>
        <div class="space-y-2">
          <label class="flex items-center space-x-2">
            <input type="checkbox" v-model="filters.bikeLanes" class="rounded text-blue-500" />
            <span class="text-sm text-gray-600">Bike Lanes</span>
          </label>
          <label class="flex items-center space-x-2">
            <input type="checkbox" v-model="filters.parking" class="rounded text-blue-500" />
            <span class="text-sm text-gray-600">Bike Parking</span>
          </label>
        </div>
      </div>

      <!-- Results Display -->
      <div class="space-y-2">
        <h3 class="text-sm font-medium text-gray-700">Results</h3>
        <div class="space-y-2">
          <div v-for="(result, index) in results" :key="index" 
               class="p-3 bg-gray-50 rounded-md hover:bg-gray-100 cursor-pointer">
            <h4 class="font-medium text-gray-800">{{ result.title }}</h4>
            <p class="text-sm text-gray-600">{{ result.description }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Footer -->
    <div class="p-4 border-t border-gray-200">
      <button 
        class="w-full px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        @click="handleSubmit"
      >
        Apply Filters
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'

const searchQuery = ref('')
const filters = reactive({
  bikeLanes: false,
  parking: false
})

const results = ref([
  {
    title: 'Example Location 1',
    description: 'Description of the location'
  },
  {
    title: 'Example Location 2',
    description: 'Another location description'
  }
])

const handleSubmit = () => {
  // Handle form submission
  console.log('Search:', searchQuery.value)
  console.log('Filters:', filters)
}
</script>

<style scoped>
.sidebar {
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  z-index: 1000;
}
</style> 