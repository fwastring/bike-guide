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
import WindBackground from '../components/WindBackground.vue'
import { useI18n } from 'vue-i18n'
import SearchCard from '../components/SearchCard.vue'

const searchStore = useSearchStore()
const searchSection = ref<HTMLElement | null>(null)
const router = useRouter()
const { t } = useI18n()

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
    <main class="flex-1 overflow-y-auto">
      <!-- Hero Section -->
      <section class="relative flex flex-col lg:flex-row items-center justify-between max-w-7xl mx-auto w-full py-6 sm:py-10">
        <!-- Desktop wind background -->
        <WindBackground class="hidden lg:block" />
        
        <!-- Videos: on mobile, above text; on desktop, right of text -->
        <div class="flex items-center justify-center w-full lg:w-2/5 px-8 xl:px-4 order-1 lg:order-2  lg:mb-0 relative z-10">
          <!-- Mobile wind background -->
          <WindBackground class="lg:hidden" />
          <div class="flex z-10 items-center">
            <video
              src="@/assets/rider1.webm"
              autoplay
              loop
              muted
              playsinline
              class="h-20 sm:h-28 lg:h-50 w-auto mr-2 lg:mr-4"
              style="background: transparent; object-fit: contain; mix-blend-mode: screen;"
            ></video>
            <video
              src="@/assets/rider2.webm"
              autoplay
              loop
              muted
              playsinline
              class="h-28 sm:h-32 lg:h-50 w-auto"
              style="background: transparent; object-fit: contain; mix-blend-mode: screen;"
            ></video>
          </div>
        </div>
        <!-- Text: smaller font, more margin -->
        <div class="relative z-20 max-w-[500px] sm:max-w-[600px] md:max-w-[600px] lg:max-w-3xl py-16 md:py-32 text-left w-full lg:w-3/5 px-8 xl:px-16 order-2 lg:order-1">
          <h1 class="font-pramukh text-7xl sm:text-8xl lg:text-8xl text-gray-900 mb-2 sm:mb-4">{{ t('landing.hero.title') }}</h1>
          
          <p class="font-poppins tracking-tight font-medium sm:text-lg text-gray-600 mb-2">{{ t('landing.hero.description') }}</p>
          <div class="w-full max-w-3xl mx-auto">
            <SearchCard @search="handleSearch" />
          </div>
        </div>
      </section>
      <!-- Onboarding Section -->
      <section id="onboarding">
        <OnboardingView />
      </section>
      <!-- About Section -->
      <AboutView />
      <Footer />
    </main>
  </div>
</template>

<style>
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@400;500;600;700&display=swap');

@font-face {
  font-family: 'Pramukh Rounded';
  src: url('@/assets/fonts/PramukhRounded-Regular.woff2') format('woff2'),
       url('@/assets/fonts/PramukhRounded-Regular.woff') format('woff');
  font-weight: 400;
  font-style: normal;
  font-display: swap;
}


@font-face {
  font-family: 'Pramukh Rounded';
  src: url('@/assets/fonts/PramukhRounded-SemiBold.woff2') format('woff2'),
       url('@/assets/fonts/PramukhRounded-SemiBold.woff') format('woff');
  font-weight: 600;
  font-style: normal;
  font-display: swap;
}

@font-face {
  font-family: 'Pramukh Rounded';
  src: url('@/assets/fonts/PramukhRounded-Bold.woff2') format('woff2'),
       url('@/assets/fonts/PramukhRounded-Bold.woff') format('woff');
  font-weight: 700;
  font-style: normal;
  font-display: swap;
}

.font-pramukh {
  font-family: 'Pramukh Rounded', sans-serif;
}

.font-poppins {
  font-family: 'Poppins', sans-serif;
}
</style>
