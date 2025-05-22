<script setup lang="ts">
import OnboardingView from './OnboardingView.vue'
import AboutView from './AboutView.vue'
import Footer from '../components/Footer.vue'
</script>

<template>
  <div class="h-screen flex flex-col">
    <main class="flex-1 overflow-y-auto">
      <OnboardingView />
      <AboutView />
      <Footer />
    </main>
  </div>
</template> 