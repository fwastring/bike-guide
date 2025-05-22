<script setup lang="ts">
import { RouterLink, useRoute } from 'vue-router'
import BaseButton from './BaseButton.vue'
import Logo from '@/assets/bike-guide logo.png'
import { ref, computed, onMounted, onUnmounted, watch } from 'vue'
import { Menu } from 'lucide-vue-next'
import { useI18n } from 'vue-i18n'
import MiniSearchCard from './MiniSearchCard.vue'
import { useScrollStore } from '@/stores/scrollStore'
import { useRouter } from 'vue-router'
import { useSearchStore } from '@/stores/searchStore'

const { t } = useI18n()
const isMenuOpen = ref(false)
const route = useRoute()
const router = useRouter()
const scrollStore = useScrollStore()
const searchStore = useSearchStore()

// Watch for scroll store changes
watch(() => scrollStore.hasScrolled, (newValue, oldValue) => {
  console.log('Header detected scroll state change:', { old: oldValue, new: newValue })
})

watch(() => scrollStore.scrollPosition, (newValue, oldValue) => {
  console.log('Header detected scroll position change:', { old: oldValue, new: newValue })
})

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value
}

onMounted(() => {
  console.log('Header mounted, initializing scroll listener')
  scrollStore.initScrollListener()
})

onUnmounted(() => {
  console.log('Header unmounted, removing scroll listener')
  scrollStore.removeScrollListener()
})

// Watch for route changes to reset scroll state
watch(() => route.path, (newPath, oldPath) => {
  console.log('Route changed in Header:', { old: oldPath, new: newPath })
  if (newPath === '/') {
    scrollStore.resetScrollState()
  }
})

const showSearch = computed(() => {
  const shouldShow = route.path !== '/' || scrollStore.hasScrolled
  console.log('Computing showSearch:', {
    path: route.path,
    hasScrolled: scrollStore.hasScrolled,
    scrollPosition: scrollStore.scrollPosition,
    shouldShow
  })
  return shouldShow
})

const handleSearch = async (searchData: { address: string; difficulty: string; distance: string }) => {
  searchStore.setSearchParams(searchData)
  await searchStore.fetchRouteOptions()
}
</script>

<template>
  <header class="fixed top-0 left-0 right-0 w-full bg-white border-b border-gray-100 z-50" role="banner">
    <nav class="max-w-7xl mx-auto flex items-center justify-center md:justify-between h-10 sm:h-14 px-4" role="navigation" aria-label="Main navigation">
      <!-- Left: Logo and Title -->
      <div class="flex items-center gap-2 min-w-max md:min-w-0">
        <RouterLink to="/" class="flex items-center gap-1 sm:gap-2 no-underline hover:opacity-80 transition-opacity" aria-label="Bike Guide Home">
          <img :src="Logo" alt="Bike Guide Logo" class="h-6 sm:h-8 w-auto" />
          <span class="text-xl sm:text-3xl tracking-tight font-bold text-gray-800">Bike Guide</span>
        </RouterLink>
      </div>

      <!-- Center: MiniSearchCard -->
      <div 
        v-show="showSearch" 
        class="hidden md:block flex-1 max-w-xl mx-4 transition-all duration-300"
        :class="{ 'opacity-0': !scrollStore.hasScrolled && route.path === '/' }"
        role="search"
        aria-label="Quick route search"
      >
        <MiniSearchCard @search="handleSearch" mobile="true" />
      </div>
      
      <!-- Right: Navigation Links and Log In -->
      <div class="hidden md:flex items-center gap-6" role="navigation" aria-label="User navigation">
        
        <RouterLink to="/help" class="text-gray-800 no-underline px-2 py-1 font-medium hover:opacity-80 transition-opacity" aria-label="Get help">{{ t('navigation.help') }}</RouterLink>
        <RouterLink to="/login" aria-label="Log in to your account">
          <BaseButton
            title="Log in"
            variant="secondary"
            class="h-10 px-6 text-sm rounded-full whitespace-nowrap"
          />
        </RouterLink>
      </div>
    </nav>
    
    <!-- Mobile Menu Dropdown -->
    <div v-if="isMenuOpen" class="md:hidden bg-white border-b border-gray-100 shadow-md" role="navigation" aria-label="Mobile navigation">
      <div class="flex flex-col p-4 space-y-3">
        <RouterLink to="/about" class="text-gray-800 no-underline px-2 py-2 font-medium hover:bg-gray-100 rounded-md transition-colors" aria-label="About Bike Guide">{{ t('navigation.about') }}</RouterLink>
        <RouterLink to="/help" class="text-gray-800 no-underline px-2 py-2 font-medium hover:bg-gray-100 rounded-md transition-colors" aria-label="Get help">{{ t('navigation.help') }}</RouterLink>
        <RouterLink to="/experimental" class="text-gray-800 no-underline px-2 py-2 font-medium hover:bg-gray-100 rounded-md transition-colors" aria-label="Experimental features">{{ t('navigation.experimental') }}</RouterLink>
        <RouterLink to="/login" aria-label="Log in to your account">
          <BaseButton
            title="Log in"
            variant="secondary"
            class="w-full h-10 px-6 text-sm rounded-full whitespace-nowrap"
          />
        </RouterLink>
      </div>
    </div>
  </header>
  <!-- Add padding to account for fixed header -->
  <div class="h-14"></div>
</template>
