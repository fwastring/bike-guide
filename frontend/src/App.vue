<script setup lang="ts">
import 'leaflet/dist/leaflet.css'
import { RouterView } from 'vue-router'
import Header from './components/Header.vue'
</script>

<template>
  <div class="w-full m-0 p-0 bg-[#FBFBFB]">
    <Header />

    <RouterView />
  </div>
</template>

