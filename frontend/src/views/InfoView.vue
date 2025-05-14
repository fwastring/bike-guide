<script setup lang="ts">
import Footer from '../components/Footer.vue'
import AboutView from './AboutView.vue';
</script>

<template>
  <div class="h-screen flex flex-col">
    <main class="flex-1 overflow-y-auto">
      <AboutView />
      <Footer />
    </main>
  </div>
</template>

