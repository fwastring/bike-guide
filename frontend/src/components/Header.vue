<script setup lang="ts">
import { RouterLink } from 'vue-router'
import BaseButton from './BaseButton.vue'
import Logo from '@/assets/bike-guide logo.png'
</script>

<template>
  <header class="w-full p-4 bg-white">
    <nav class="max-w-7xl mx-auto flex justify-between items-center">
      <RouterLink to="/" class="flex items-center gap-2 no-underline hover:opacity-80 transition-opacity">
        <img :src="Logo" alt="Bike Guide Logo" class="h-8 w-auto" />
        <span class="text-2xl font-bold text-gray-800">Bike Guide</span>
      </RouterLink>
      <div class="flex gap-4 items-center">
        <RouterLink to="/about" class="text-gray-800 no-underline px-4 py-2 rounded hover:bg-gray-100 transition-colors">About</RouterLink>
        <RouterLink to="/search" class="text-gray-800 no-underline px-4 py-2 rounded hover:bg-gray-100 transition-colors">Search</RouterLink>
        <BaseButton 
          title="Experimental" 
          variant="secondary" 
          link="/experimental"
          class="ml-2"
        />
      </div>
    </nav>
  </header>
</template>
