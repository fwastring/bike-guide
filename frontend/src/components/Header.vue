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
  <header class="fixed top-0 left-0 right-0 w-full bg-white border-b border-gray-100 z-50">
    <nav class="max-w-7xl mx-auto flex items-center justify-center md:justify-between h-14 px-4">
      <!-- Left: Logo and Title -->
      <div class="flex items-center gap-2 min-w-max md:min-w-0">
        <RouterLink to="/" class="flex items-center gap-2 no-underline hover:opacity-80 transition-opacity">
          <img :src="Logo" alt="Bike Guide Logo" class="h-8 w-auto" />
          <span class="text-2xl font-bold text-gray-800">Bike Guide</span>
        </RouterLink>
      </div>

      <!-- Center: MiniSearchCard -->
      <div 
        v-show="showSearch" 
        class="hidden md:block flex-1 max-w-xl mx-4 transition-all duration-300"
        :class="{ 'opacity-0': !scrollStore.hasScrolled && route.path === '/' }"
      >
        <MiniSearchCard @search="handleSearch" mobile="true" />
      </div>
      
      <!-- Right: Navigation Links and Log In -->
      <div class="hidden md:flex items-center gap-6">
        <RouterLink to="/about" class="text-gray-800 no-underline px-2 py-1 font-medium hover:opacity-80 transition-opacity">{{ t('navigation.about') }}</RouterLink>
        <RouterLink to="/help" class="text-gray-800 no-underline px-2 py-1 font-medium hover:opacity-80 transition-opacity">{{ t('navigation.help') }}</RouterLink>
        <RouterLink to="/login">
          <BaseButton
            title="Log in"
            variant="secondary"
            class="h-10 px-6 text-sm rounded-full whitespace-nowrap"
          />
        </RouterLink>
      </div>
    </nav>
    
    <!-- Mobile Menu Dropdown -->
    <div v-if="isMenuOpen" class="md:hidden bg-white border-b border-gray-100 shadow-md">
      <div class="flex flex-col p-4 space-y-3">
        <RouterLink to="/about" class="text-gray-800 no-underline px-2 py-2 font-medium hover:bg-gray-100 rounded-md transition-colors">{{ t('navigation.about') }}</RouterLink>
        <RouterLink to="/help" class="text-gray-800 no-underline px-2 py-2 font-medium hover:bg-gray-100 rounded-md transition-colors">{{ t('navigation.help') }}</RouterLink>
        <RouterLink to="/experimental" class="text-gray-800 no-underline px-2 py-2 font-medium hover:bg-gray-100 rounded-md transition-colors">{{ t('navigation.experimental') }}</RouterLink>
        <RouterLink to="/login">
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
