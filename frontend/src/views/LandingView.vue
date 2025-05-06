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
      <section class="relative flex flex-col lg:flex-row items-center justify-between max-w-7xl mx-auto w-full py-6 sm:py-10 gap-6 sm:gap-4">
        <!-- Wind background, only in hero area -->
        <div class="absolute left-0 right-0 top-0 h-full z-0 pointer-events-none px-0">
          <svg class="w-full h-full" viewBox="0 220 1440 80">
            <!-- Layer 1: Main wind -->
            <path
              d="M0,240 C480,270 960,230 1440,250"
              fill="none"
              stroke="#a0aec0"
              stroke-width="8"
              opacity="0.7"
            >
              <animate
                attributeName="d"
                values="
                  M0,240 C480,270 960,230 1440,250;
                  M0,250 C480,260 960,240 1440,260;
                  M0,240 C480,270 960,230 1440,250
                "
                dur="4s"
                repeatCount="indefinite"
              />
            </path>
            <!-- Layer 2: Subtle, thinner, faster -->
            <path
              d="M0,260 C360,280 1080,250 1440,270"
              fill="none"
              stroke="#cbd5e1"
              stroke-width="4"
              opacity="0.5"
            >
              <animate
                attributeName="d"
                values="
                  M0,260 C360,280 1080,250 1440,270;
                  M0,270 C360,270 1080,260 1440,280;
                  M0,260 C360,280 1080,250 1440,270
                "
                dur="2.5s"
                repeatCount="indefinite"
              />
            </path>
            <!-- Layer 3: Wavy, slow, thick, more transparent -->
            <path
              d="M0,280 C600,290 840,270 1440,290"
              fill="none"
              stroke="#94a3b8"
              stroke-width="12"
              opacity="0.25"
            >
              <animate
                attributeName="d"
                values="
                  M0,280 C600,290 840,270 1440,290;
                  M0,290 C600,280 840,280 1440,300;
                  M0,280 C600,290 840,270 1440,290
                "
                dur="6s"
                repeatCount="indefinite"
              />
            </path>
            <!-- Layer 4: Small, quick gust -->
            <path
              d="M0,300 C700,310 900,290 1440,310"
              fill="none"
              stroke="#e0e7ef"
              stroke-width="3"
              opacity="0.3"
            >
              <animate
                attributeName="d"
                values="
                  M0,300 C700,310 900,290 1440,310;
                  M0,310 C700,300 900,300 1440,320;
                  M0,300 C700,310 900,290 1440,310
                "
                dur="1.8s"
                repeatCount="indefinite"
              />
            </path>
          </svg>
          <!-- Left gradient (existing) -->
          <div class="absolute left-0 top-0 h-full w-1/2 bg-gradient-to-r from-white/100 to-white/0 pointer-events-none"></div>
          <!-- Right gradient (new) -->
          <div class="absolute right-0 top-0 h-full w-1/2 bg-gradient-to-l from-white/100 to-white/0 pointer-events-none"></div>
        </div>
        <!-- Videos: on mobile, above text; on desktop, right of text -->
        <div class="flex items-center justify-center w-full lg:w-2/5 px-8 xl:px-4 order-1 lg:order-2 mb-6 lg:mb-0">
          <div class="flex items-center">
            <video
              src="@/assets/rider1.webm"
              autoplay
              loop
              muted
              playsinline
              class="h-20 sm:h-28 lg:h-50 w-auto mr-2 lg:mr-4 z-10"
              style="background: transparent;"
            ></video>
            <video
              src="@/assets/rider2.webm"
              autoplay
              loop
              muted
              playsinline
              class="h-28 sm:h-32 lg:h-50 w-auto z-10"
              style="background: transparent;"
            ></video>
          </div>
        </div>
        <!-- Text: smaller font, more margin -->
        <div class="relative z-20 max-w-3xl py-16 md:py-32 text-left w-full lg:w-3/5 px-8 xl:px-16 order-2 lg:order-1">
          <h1 class="text-2xl sm:text-3xl lg:text-4xl font-bold text-gray-900 mb-2 sm:mb-4">Your Perfect Ride Awaits</h1>
          <p class="text-lg sm:text-xl lg:text-2xl text-gray-500 mb-2">Discover cycling routes that match your style</p>
          <p class="text-base sm:text-lg text-gray-600 mb-2">Tell us your preferences, and we'll find the best bike routes for you — whether you're a casual rider or an experienced cyclist.</p>
          <div class="w-full mt-6">
            <MiniSearchCard @search="handleSearch" />
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
