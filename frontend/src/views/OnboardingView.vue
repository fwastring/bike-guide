<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import BaseButton from '../components/BaseButton.vue'
import { useRouter } from 'vue-router'
import Header from '../components/Header.vue'
import { MapPin, Map, Bike, ArrowRight } from 'lucide-vue-next'
import { useI18n } from 'vue-i18n'

const router = useRouter()
const { t } = useI18n()

const handleGetStarted = () => {
  router.push('/search')
}

// Animation on scroll effect
const isVisible = ref([false, false, false])

const checkVisibility = () => {
  const cards = document.querySelectorAll('.step-card')
  cards.forEach((card, index) => {
    const rect = card.getBoundingClientRect()
    if (rect.top <= window.innerHeight * 0.8) {
      isVisible.value[index] = true
    }
  })
}

onMounted(() => {
  checkVisibility()
  window.addEventListener('scroll', checkVisibility)
})

onUnmounted(() => {
  window.removeEventListener('scroll', checkVisibility)
})
</script>

<template>
  <section class="relative py-20 bg-gradient-to-b from-[#8E7DBE]/10 to-[#A7A3E2]/10 overflow-hidden">
    <!-- Animated background pattern -->
    <div class="absolute inset-0 z-0">
      <div class="absolute inset-0 opacity-10">
        <img
          src="@/assets/mapbackground.jpg"
          alt="Map background"
          class="absolute inset-0 w-full h-full object-cover pointer-events-none select-none will-change-transform"
          aria-hidden="true"
        />
      </div>
      
      <!-- Abstract shapes -->
      <div class="absolute top-20 left-10 w-64 h-64 bg-[#8E7DBE] rounded-full mix-blend-multiply filter blur-3xl opacity-20 animate-blob"></div>
      <div class="absolute top-40 right-10 w-72 h-72 bg-[#A7A3E2] rounded-full mix-blend-multiply filter blur-3xl opacity-20 animate-blob animation-delay-2000"></div>
      <div class="absolute bottom-40 left-1/3 w-80 h-80 bg-[#7D6CAE] rounded-full mix-blend-multiply filter blur-3xl opacity-20 animate-blob animation-delay-4000"></div>
    </div>
    
    <div class="relative z-10 max-w-7xl mx-auto px-4 sm:px-10 lg:px-12">
      <!-- Hero section -->
      <div class="text-center mb-16 md:mb-24">
     
        
        <h1 class="font-pramukh text-5xl sm:text-6xl lg:text-7xl bg-clip-text text-transparent bg-gradient-to-r from-[#8E7DBE] to-[#A7A3E2] mb-2 leading-tight max-w-5xl mx-auto">
          {{ t('onboarding.title') }}
        </h1>
        
        <p class="font-poppins text-sm sm:text-lg sm:text-xl text-gray-600 max-w-3xl mx-auto mb-10">
          {{ t('onboarding.subtitle') }}
        </p>
        
        
      </div>
      
      <!-- Steps grid -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 lg:gap-10 max-w-6xl mx-auto">
        <!-- Step 1 -->
        <div 
          class="border border-gray-300 rounded-2xl overflow-hidden transition-all duration-300 bg-[#FBFBFB] hover:scale-[1.02] p-8 sm:p-10 transition-all duration-500 transform hover:-translate-y-2"
          :class="{ 'animate-fadeIn': isVisible[0] }"
        >
          <div class="mb-8 relative">
            <div class="absolute -top-4 -left-4 w-20 h-20 bg-gradient-to-br from-indigo-400 to-indigo-600 rounded-2xl transform -rotate-6"></div>
            <div class="relative w-16 h-16 bg-indigo-50 rounded-2xl flex items-center justify-center">
              <MapPin class="w-8 h-8 text-indigo-600" />
            </div>
          </div>
          
          <span class="inline-block px-3 py-1 bg-indigo-100 text-indigo-800 text-xs font-semibold rounded-full mb-3">
            {{ t('onboarding.steps.find.step', 'Step 1') }}
          </span>
          
          <h3 class="font-poppins text-xl sm:text-2xl font-bold text-gray-900 mb-4">
            {{ t('onboarding.steps.find.title') }}
          </h3>
          
          <p class="font-poppins text-gray-600 text-sm sm:text-lg tracking-tight">
            {{ t('onboarding.steps.find.description') }}
          </p>
        </div>
        
        <!-- Step 2 -->
        <div 
          class="border border-gray-300 rounded-2xl overflow-hidden transition-all duration-300 bg-[#FBFBFB] hover:scale-[1.02] p-8 sm:p-10 transition-all duration-500 transform hover:-translate-y-2"
          :class="{ 'animate-fadeIn animation-delay-300': isVisible[1] }"
        >
          <div class="mb-8 relative">
            <div class="absolute -top-4 -left-4 w-20 h-20 bg-gradient-to-br from-indigo-400 to-indigo-600 rounded-2xl transform -rotate-6"></div>
            <div class="relative w-16 h-16 bg-indigo-50 rounded-2xl flex items-center justify-center">
              <Map class="w-8 h-8 text-indigo-600" />
            </div>
          </div>
          
          <span class="inline-block px-3 py-1 bg-indigo-100 text-indigo-800 text-xs font-semibold rounded-full mb-3">
            {{ t('onboarding.steps.choose.step', 'Step 2') }}
          </span>
          
          <h3 class="font-poppins text-xl sm:text-2xl font-bold text-gray-900 mb-4">
            {{ t('onboarding.steps.choose.title') }}
          </h3>
          
          <p class="font-poppins text-gray-600 text-sm sm:text-lg tracking-tight">
            {{ t('onboarding.steps.choose.description') }}
          </p>
        </div>
        
        <!-- Step 3 -->
        <div 
          class="border border-gray-300 rounded-2xl overflow-hidden transition-all duration-300 bg-[#FBFBFB] hover:scale-[1.02] p-8 sm:p-10 transition-all duration-500 transform hover:-translate-y-2"
          :class="{ 'animate-fadeIn animation-delay-600': isVisible[2] }"
        >
          <div class="mb-8 relative">
            <div class="absolute -top-4 -left-4 w-20 h-20 bg-gradient-to-br from-purple-400 to-purple-600 rounded-2xl transform -rotate-6"></div>
            <div class="relative w-16 h-16 bg-purple-50 rounded-2xl flex items-center justify-center">
              <Bike class="w-8 h-8 text-purple-600" />
            </div>
          </div>
          
          <span class="inline-block px-3 py-1 bg-purple-100 text-purple-800 text-xs font-semibold rounded-full mb-3">
            {{ t('onboarding.steps.go.step', 'Step 3') }}
          </span>
          
          <h3 class="font-poppins text-xl sm:text-2xl font-bold text-gray-900 mb-4">
            {{ t('onboarding.steps.go.title') }}
          </h3>
          
          <p class="font-poppins tracking-tight sm:text-lg text-gray-600 text-sm sm:text-lg">
            {{ t('onboarding.steps.go.description') }}
          </p>
        </div>
      </div>
      
      
    </div>
  </section>
</template>

<style scoped>
.font-pramukh {
  font-family: 'Pramukh Rounded', sans-serif;
}

.font-poppins {
  font-family: 'Poppins', sans-serif;
}

@keyframes blob {
  0% {
    transform: translate(0px, 0px) scale(1);
  }
  33% {
    transform: translate(30px, -50px) scale(1.1);
  }
  66% {
    transform: translate(-20px, 20px) scale(0.9);
  }
  100% {
    transform: translate(0px, 0px) scale(1);
  }
}

.animate-blob {
  animation: blob 7s infinite;
}

.animation-delay-2000 {
  animation-delay: 2s;
}

.animation-delay-4000 {
  animation-delay: 4s;
}

.animation-delay-300 {
  animation-delay: 0.3s;
}

.animation-delay-600 {
  animation-delay: 0.6s;
}

@keyframes fadeIn {
  0% {
    opacity: 0;
    transform: translateY(20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fadeIn {
  animation: fadeIn 0.8s ease-out forwards;
}
</style>