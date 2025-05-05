<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import Footer from '../components/Footer.vue'
import Rider1 from '@/assets/rider 1.png'
import Rider2 from '@/assets/rider 2.png'
import { useSearchStore } from '@/stores/searchStore'
import MiniSearchCard from '../components/MiniSearchCard.vue'
import OnboardingView from './OnboardingView.vue'
import AboutView from './AboutView.vue'

const searchStore = useSearchStore()
const searchSection = ref<HTMLElement | null>(null)
const router = useRouter()

const handleSearch = async (searchData: { address: string; difficulty: string; distance: string }) => {
  searchStore.setSearchParams(searchData)
  await router.push('/options')
}

const scrollToSearch = () => {
  if (searchSection.value) {
    searchSection.value.scrollIntoView({ behavior: 'smooth' })
  }
}
</script>

<template>
  <div class="h-screen flex flex-col">
    <main class="flex-1 overflow-y-auto bg-white">
      <!-- Hero Section -->
      <section class="flex flex-col sm:flex-row items-center justify-between max-w-7xl mx-auto w-full px-4 py-6 sm:py-10 gap-6 sm:gap-4">
        <div class="max-w-2xl py-20 md:py-40 text-left w-full sm:w-3/5">
          <h1 class="text-4xl sm:text-5xl lg:text-6xl font-bold text-gray-900 mb-2 sm:mb-4">Your Perfect Ride Awaits</h1>
          <p class="text-xl sm:text-2xl lg:text-3xl text-gray-500 mb-2">Discover cycling routes that match your style</p>
          <p class="text-lg sm:text-xl text-gray-600 mb-2">Tell us your preferences, and we'll find the best bike routes for you — whether you're a casual rider or an experienced cyclist.</p>
          <div class="w-full mt-6">
            <MiniSearchCard @search="handleSearch" />
          </div>
        </div>
        <div class="flex items-center justify-center w-full sm:w-2/5 mt-8 sm:mt-0">
          <div class="flex items-center">
            <img 
              :src="Rider1" 
              alt="Cyclist 1" 
              class="h-20 sm:h-28 lg:h-32 w-auto mr-2 sm:mr-4" 
            />
            <img 
              :src="Rider2" 
              alt="Cyclist 2" 
              class="h-28 sm:h-32 lg:h-40 w-auto" 
            />
          </div>
        </div>
      </section>
      <!-- Onboarding Section -->
        <OnboardingView />
      <!-- About Section -->
        <AboutView />
      <Footer />
    </main>
  </div>
</template>
