<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
  <header class="w-full p-4 bg-white">
    <nav class="max-w-7xl mx-auto flex justify-between items-center">
      <RouterLink to="/" class="text-2xl font-bold text-gray-800 no-underline">Bike Guide</RouterLink>
      <div class="flex gap-4">
        <RouterLink to="/about" class="text-gray-800 no-underline px-4 py-2 rounded hover:bg-gray-100 transition-colors">About</RouterLink>
        <RouterLink to="/search" class="text-gray-800 no-underline px-4 py-2 rounded hover:bg-gray-100 transition-colors">Search</RouterLink>
      </div>
    </nav>
  </header>
</template>


