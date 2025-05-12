<script setup lang="ts">
import Footer from '../components/Footer.vue'
import { useI18n } from 'vue-i18n'
import { Linkedin, Github } from 'lucide-vue-next'

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
              <div class="flex flex-col sm:flex-row items-center sm:items-start gap-6 sm:gap-8">
                <img 
                  src="@/assets/jonathan.jpeg" 
                  alt="Jonathan Ahlström" 
                  class="w-32 h-32 rounded-full object-cover shadow-lg" 
                />
                <div class="flex flex-col items-center sm:items-start text-center sm:text-left">
                  <h2 class="text-2xl text-gray-900 mb-2">Jonathan Ahlström</h2>
                  <div class="flex gap-4">
                    <a 
                      href="https://www.linkedin.com/in/jonathanahlstrom/" 
                      target="_blank" 
                      rel="noopener" 
                      class="text-blue-600 font-poppins hover:text-blue-700 transition-colors flex items-center gap-2"
                    >
                      <div class="w-8 h-8 rounded-full bg-blue-100 flex items-center justify-center">
                        <Linkedin class="w-4 h-4" />
                      </div>
                      <span>{{ t('about.links.linkedin') }}</span>
                    </a>
                    <a 
                      href="https://github.com/jonahlstrom" 
                      target="_blank" 
                      rel="noopener" 
                      class="text-gray-700 font-poppins hover:text-gray-800 transition-colors flex items-center gap-2"
                    >
                      <div class="w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center">
                        <Github class="w-4 h-4" />
                      </div>
                      <span>{{ t('about.links.github') }}</span>
                    </a>
                  </div>
                </div>
              </div>

              <!-- Fredrik -->
              <div class="flex flex-col sm:flex-row items-center sm:items-start gap-6 sm:gap-8">
                <img 
                  src="@/assets/fredrik.jpeg" 
                  alt="Fredrik" 
                  class="w-32 h-32 rounded-full object-cover shadow-lg" 
                />
                <div class="flex flex-col items-center sm:items-start text-center sm:text-left">
                  <h2 class="text-2xl text-gray-900 mb-2">Fredrik Wastring</h2>
                  <div class="flex gap-4">
                    <a 
                      href="[Friend_LinkedIn]" 
                      target="_blank" 
                      rel="noopener" 
                      class="text-blue-600 font-poppins hover:text-blue-700 transition-colors flex items-center gap-2"
                    >
                      <div class="w-8 h-8 rounded-full bg-blue-100 flex items-center justify-center">
                        <Linkedin class="w-4 h-4" />
                      </div>
                      <span>{{ t('about.links.linkedin') }}</span>
                    </a>
                    <a 
                      href="[Friend_GitHub]" 
                      target="_blank" 
                      rel="noopener" 
                      class="text-gray-700 font-poppins hover:text-gray-800 transition-colors flex items-center gap-2"
                    >
                      <div class="w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center">
                        <Github class="w-4 h-4" />
                      </div>
                      <span>{{ t('about.links.github') }}</span>
                    </a>
                  </div>
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

