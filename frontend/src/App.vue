<script setup lang="ts">
import 'leaflet/dist/leaflet.css'
import { RouterView } from 'vue-router'
</script>

<template>
  <div class="w-full h-screen m-0 p-0 bg-[#FBFBFB]">
    <RouterView />
  </div>
</template>

<style>
#app {
  width: 100%;
  height: 100vh;
  margin: 0;
  padding: 0;
}
</style>
