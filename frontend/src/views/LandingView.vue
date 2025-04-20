<script setup lang="ts">
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'
import BaseButton from '../components/BaseButton.vue'
</script>

<template>
  <div class="min-h-screen flex flex-col">
    <Header />
    <main class="flex-1 flex items-center justify-center p-8">
      <section class="text-center max-w-3xl">
        <h1 class="text-5xl font-bold text-gray-800 mb-4">Find Your Perfect Bike Route</h1>
        <p class="text-xl text-gray-600 mb-8">Discover the best cycling routes tailored to your preferences and skill level</p>
        <div class="flex gap-4 justify-center">
          <BaseButton
            title="Get Started"
            link="/onboarding"
            variant="primary"
          />
          <BaseButton
            title="Learn More"
            link="/about"
            variant="secondary"
          />
        </div>
      </section>
    </main>
    <Footer />
  </div>
</template>

