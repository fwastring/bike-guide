<script setup lang="ts">
import Footer from '../components/Footer.vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
</script>

<template>
  <div class="flex flex-col min-h-screen bg-gray-50">
    <main class="flex-1">
      <section class="relative py-16 md:py-24">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <!-- Content Grid -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-16 items-center">
            <!-- Left: Text Content -->
            <div class="flex flex-col justify-center order-2 md:order-1">
              <h1 class="font-pramukh text-5xl sm:text-6xl lg:text-7xl text-gray-900 mb-6">{{ t('about.title') }}</h1>
              <p class="font-poppins text-base sm:text-lg text-gray-600">
                {{ t('about.description') }}
              </p>
            </div>

            <!-- Right: Team Photos -->
            <div class="flex flex-col gap-16 order-1 md:order-2">
              <!-- Jonathan -->
              <div class="flex flex-col items-center text-center">
                <img 
                  src="@/assets/jonathan.jpeg" 
                  alt="Jonathan Ahlström" 
                  class="w-48 h-48 rounded-full object-cover shadow-lg mb-6" 
                />
                <h2 class="font-pramukh text-3xl text-gray-900 mb-2">Jonathan Ahlström</h2>
                <div class="flex gap-4 mt-2">
                  <a 
                    href="https://www.linkedin.com/in/jonathanahlstrom/" 
                    target="_blank" 
                    rel="noopener" 
                    class="text-blue-600 font-poppins hover:underline"
                  >
                    {{ t('about.links.linkedin') }}
                  </a>
                  <a 
                    href="https://github.com/jonahlstrom" 
                    target="_blank" 
                    rel="noopener" 
                    class="text-gray-700 font-poppins hover:underline"
                  >
                    {{ t('about.links.github') }}
                  </a>
                </div>
              </div>

              <!-- Fredrik -->
              <div class="flex flex-col items-center text-center">
                <img 
                  src="@/assets/fredrik.jpeg" 
                  alt="Fredrik" 
                  class="w-48 h-48 rounded-full object-cover shadow-lg mb-6" 
                />
                <h2 class="font-pramukh text-3xl text-gray-900 mb-2">Fredrik</h2>
                <div class="flex gap-4 mt-2">
                  <a 
                    href="[Friend_LinkedIn]" 
                    target="_blank" 
                    rel="noopener" 
                    class="text-blue-600 font-poppins hover:underline"
                  >
                    {{ t('about.links.linkedin') }}
                  </a>
                  <a 
                    href="[Friend_GitHub]" 
                    target="_blank" 
                    rel="noopener" 
                    class="text-gray-700 font-poppins hover:underline"
                  >
                    {{ t('about.links.github') }}
                  </a>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </main>
  </div>
</template>

<style scoped>
.font-pramukh {
  font-family: 'Pramukh Rounded', sans-serif;
}
.font-poppins {
  font-family: 'Poppins', sans-serif;
}
</style>

