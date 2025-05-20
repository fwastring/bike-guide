<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
  <footer class="bg-gray-800 text-white py-16 mt-auto">
    <div class="max-w-7xl mx-auto px-4">
      <div class="flex flex-col items-center space-y-4">
        <h3 class="text-xl font-semibold">Bike Guide</h3>
        <nav class="flex gap-6">
          <RouterLink to="/" class="text-gray-300 hover:text-white transition-colors">Search</RouterLink>
          <RouterLink to="/help" class="text-gray-300 hover:text-white transition-colors">Help</RouterLink>
        </nav>
        <p class="text-gray-400">&copy; {{ new Date().getFullYear() }} Bike Guide</p>
      </div>
      <div class="text-center mt-8 pt-8 border-t border-gray-700">

      </div>
    </div>
  </footer>
</template> 