<script setup lang="ts">
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'
import BaseButton from '../components/BaseButton.vue'
import Rider1 from '@/assets/rider 1.png'
import Rider2 from '@/assets/rider 2.png'
import OnboardingView from './OnboardingView.vue'
import SearchView from './SearchView.vue'
</script>

<template>
  <div class="sm:h-screen flex flex-col">
    <main class="flex-1 overflow-y-auto sm:snap-y sm:snap-mandatory">
      <div class="sm:snap-start">
        <div class="flex items-center sm:h-screen p-4 sm:p-8">
          <div class="max-w-7xl mx-auto w-full">
            <div class="flex flex-col sm:flex-row justify-between items-center w-full gap-8 sm:gap-4">
              <div class="max-w-2xl text-center sm:text-left">
                <h1 class="text-4xl sm:text-5xl lg:text-6xl font-bold text-gray-900 mb-4 sm:mb-6">Your Perfect Ride Awaits</h1>
                <p class="text-xl sm:text-2xl lg:text-3xl text-gray-500 mb-4">Discover cycling routes that match your style</p>
                <p class="text-lg sm:text-xl text-gray-600 mb-8">Tell us your preferences, and we'll find the best bike routes for you — whether you're a casual rider or an experienced cyclist.</p>
                <div class="flex flex-col sm:flex-row gap-4">
                  <BaseButton
                    title="Find Route"
                    link="/onboarding"
                    variant="primary"
                    class="px-6 sm:px-8 py-2 sm:py-3 text-base sm:text-lg"
                  />
                  <BaseButton
                    title="Help"
                    link="/onboarding"
                    variant="secondary"
                    class="px-6 sm:px-8 py-2 sm:py-3 text-base sm:text-lg bg-gray-100"
                  />
                </div>
              </div>
              <div class="relative flex items-end justify-center sm:justify-end mt-8 sm:mt-0">
                <img 
                  :src="Rider1" 
                  alt="Cyclist 1" 
                  class="h-32 sm:h-40 lg:h-48 w-auto mr-4 sm:mr-8" 
                />
                <img 
                  :src="Rider2" 
                  alt="Cyclist 2" 
                  class="h-40 sm:h-48 lg:h-64 w-auto" 
                />
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="sm:snap-start">
        <OnboardingView />
      </div>
      <div class="sm:snap-start">
        <SearchView />
      </div>
      <div class="sm:snap-start">
        <Footer />
      </div>
    </main>
  </div>
</template>


