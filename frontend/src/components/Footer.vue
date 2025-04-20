<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
  <footer class="bg-gray-800 text-white py-8 mt-auto">
    <div class="max-w-7xl mx-auto px-4 grid grid-cols-1 md:grid-cols-3 gap-8">
      <div class="space-y-2">
        <h3 class="text-xl font-semibold">Bike Guide</h3>
        <p class="text-gray-300">Your ultimate cycling companion</p>
      </div>
      <div class="space-y-2">
        <h4 class="text-lg font-medium">Quick Links</h4>
        <RouterLink to="/about" class="block text-gray-300 hover:text-white transition-colors">About</RouterLink>
        <RouterLink to="/search" class="block text-gray-300 hover:text-white transition-colors">Search</RouterLink>
      </div>
      <div class="space-y-2">
        <h4 class="text-lg font-medium">Legal</h4>
        <RouterLink to="/privacy" class="block text-gray-300 hover:text-white transition-colors">Privacy Policy</RouterLink>
        <RouterLink to="/terms" class="block text-gray-300 hover:text-white transition-colors">Terms of Service</RouterLink>
      </div>
    </div>
    <div class="max-w-7xl mx-auto px-4 pt-8 mt-8 border-t border-gray-700 text-center">
      <p>&copy; {{ new Date().getFullYear() }} Bike Guide. All rights reserved.</p>
    </div>
  </footer>
</template> 