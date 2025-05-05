<script setup lang="ts">
import 'leaflet/dist/leaflet.css'
import { RouterView } from 'vue-router'
import Header from './components/Header.vue'
</script>

<template>
  <div class="bg-[#FBFBFB]">
    <Header />
    <RouterView />
  </div>
</template>

