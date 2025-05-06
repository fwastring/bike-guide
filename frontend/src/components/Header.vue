<script setup lang="ts">
import { RouterLink } from 'vue-router'
import BaseButton from './BaseButton.vue'
import Logo from '@/assets/bike-guide logo.png'
import { ref } from 'vue'
import { Menu, Globe } from 'lucide-vue-next'
import { useI18n } from 'vue-i18n'

const { t, locale } = useI18n()
const isMenuOpen = ref(false)

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value
}

const toggleLanguage = () => {
  locale.value = locale.value === 'en' ? 'sv' : 'en'
}
</script>

<template>
  <header class="w-full bg-white border-b border-gray-100">
    <nav class="max-w-7xl mx-auto flex items-center justify-between h-14 px-4">
      <!-- Hamburger Menu for Mobile -->
      <button @click="toggleMenu" class="md:hidden text-gray-800">
        <Menu class="w-6 h-6" />
      </button>
      
      <!-- Left: Logo and Title -->
      <div class="flex items-center gap-2 min-w-max">
        <RouterLink to="/" class="flex items-center gap-2 no-underline hover:opacity-80 transition-opacity">
          <img :src="Logo" alt="Bike Guide Logo" class="h-8 w-auto" />
          <span class="text-2xl font-bold text-gray-800">Bike Guide</span>
        </RouterLink>
      </div>
      
      <!-- Center: Navigation Links with Dividers - Hidden on Mobile -->
      <div class="hidden md:flex items-center gap-6 mx-auto">
        <RouterLink to="/about" class="text-gray-800 no-underline px-2 py-1 font-medium hover:opacity-80 transition-opacity">{{ t('navigation.about') }}</RouterLink>
        <span class="h-6 w-px bg-gray-200 mx-1"></span>
        <RouterLink to="/help" class="text-gray-800 no-underline px-2 py-1 font-medium hover:opacity-80 transition-opacity">{{ t('navigation.help') }}</RouterLink>
        <span class="h-6 w-px bg-gray-200 mx-1"></span>
        <BaseButton 
          :title="t('navigation.experimental')"
          variant="secondary" 
          link="/experimental"
          class="px-2 py-1 font-medium !rounded-full !bg-transparent !border-0 text-gray-800 hover:bg-gray-100 transition-colors"
        />
      </div>
      
      <!-- Right: Language Switcher and Log In -->
      <div class="flex items-center gap-4 min-w-max">
        <!-- Language Switcher -->
        <button 
          @click="toggleLanguage" 
          class="text-gray-800 hover:opacity-80 transition-opacity"
          :title="t('language.switch')"
        >
          <Globe class="w-6 h-6" />
        </button>
        
        <!-- Log In -->
        <RouterLink to="/login" class="flex items-center gap-2">
          <span class="text-gray-800 no-underline font-medium mr-0 hover:opacity-80 transition-opacity hidden sm:inline">{{ t('navigation.login') }}</span>
          <span class="inline-flex items-center justify-center w-8 h-8 rounded-full bg-gray-100">
            <svg xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 24 24" class="w-6 h-6 text-gray-700">
              <circle cx="12" cy="8" r="4" />
              <ellipse cx="12" cy="17" rx="7" ry="4" />
            </svg>
          </span>
        </RouterLink>
      </div>
    </nav>
    
    <!-- Mobile Menu Dropdown -->
    <div v-if="isMenuOpen" class="md:hidden bg-white border-b border-gray-100 shadow-md">
      <div class="flex flex-col p-4 space-y-3">
        <RouterLink to="/about" class="text-gray-800 no-underline px-2 py-2 font-medium hover:bg-gray-100 rounded-md transition-colors">{{ t('navigation.about') }}</RouterLink>
        <RouterLink to="/help" class="text-gray-800 no-underline px-2 py-2 font-medium hover:bg-gray-100 rounded-md transition-colors">{{ t('navigation.help') }}</RouterLink>
        <RouterLink to="/experimental" class="text-gray-800 no-underline px-2 py-2 font-medium hover:bg-gray-100 rounded-md transition-colors">{{ t('navigation.experimental') }}</RouterLink>
        <button 
          @click="toggleLanguage" 
          class="text-gray-800 no-underline px-2 py-2 font-medium hover:bg-gray-100 rounded-md transition-colors text-left"
        >
          {{ t('language.switch') }}
        </button>
      </div>
    </div>
  </header>
</template>
