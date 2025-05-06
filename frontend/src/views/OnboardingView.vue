<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import BaseButton from '../components/BaseButton.vue'
import { useRouter } from 'vue-router'
import Header from '../components/Header.vue'
import { MapPin, Map, Bike } from 'lucide-vue-next'
import { useI18n } from 'vue-i18n'

const router = useRouter()
const { t } = useI18n()

const handleGetStarted = () => {
  router.push('/search')
}

</script>

<template>
  <section class="relative py-20 md:py-32 bg-gray-50 overflow-hidden">
    <img
      src="@/assets/mapbackground.jpg"
      alt="Map background"
      class="absolute inset-0 w-full h-full object-cover opacity-10 pointer-events-none select-none z-0 will-change-transform transition-transform duration-300"
      aria-hidden="true"
    />
    <div
      class="absolute inset-0 w-full h-full z-0 pointer-events-none select-none bg-blue-500 opacity-20 mix-blend-multiply will-change-transform transition-transform duration-300"
    ></div>
    <div class="relative z-10 max-w-7xl mx-auto px-8 sm:px-10 lg:px-12">
      <div class="text-center mb-20">
        <h2 class="font-pramukh text-6xl sm:text-6xl lg:text-7xl text-gray-600 mb-4 sm:mb-6">{{ t('onboarding.title') }}</h2>
        <p class="font-poppins text-lg sm:text-xl text-gray-600 max-w-3xl mx-auto">{{ t('onboarding.subtitle') }}</p>
      </div>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-10 max-w-6xl mx-auto">
        <!-- Step 1 -->
        <div class="bg-white rounded-2xl p-8 shadow-md hover:shadow-lg transition-shadow">
          <div class="w-20 h-20 bg-blue-100 rounded-full flex items-center justify-center mb-6 mx-auto">
            <MapPin class="w-10 h-10 text-blue-600" />
          </div>
          <h3 class="font-poppins text-2xl font-semibold text-gray-900 mb-3 text-center">{{ t('onboarding.steps.find.title') }}</h3>
          <p class="font-poppins text-gray-600 text-center">{{ t('onboarding.steps.find.description') }}</p>
        </div>
        <!-- Step 2 -->
        <div class="bg-white rounded-2xl p-8 shadow-md hover:shadow-lg transition-shadow">
          <div class="w-20 h-20 bg-blue-100 rounded-full flex items-center justify-center mb-6 mx-auto">
            <Map class="w-10 h-10 text-blue-600" />
          </div>
          <h3 class="font-poppins text-2xl font-semibold text-gray-900 mb-3 text-center">{{ t('onboarding.steps.choose.title') }}</h3>
          <p class="font-poppins text-gray-600 text-center">{{ t('onboarding.steps.choose.description') }}</p>
        </div>
        <!-- Step 3 -->
        <div class="bg-white rounded-2xl p-8 shadow-md hover:shadow-lg transition-shadow">
          <div class="w-20 h-20 bg-blue-100 rounded-full flex items-center justify-center mb-6 mx-auto">
            <Bike class="w-10 h-10 text-blue-600" />
          </div>
          <h3 class="font-poppins text-2xl font-semibold text-gray-900 mb-3 text-center">{{ t('onboarding.steps.go.title') }}</h3>
          <p class="font-poppins text-gray-600 text-center">{{ t('onboarding.steps.go.description') }}</p>
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
</style>
