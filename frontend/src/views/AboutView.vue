<script setup lang="ts">
import Footer from '../components/Footer.vue'
import { useI18n } from 'vue-i18n'
import { Linkedin, Github, Mail, ArrowRight } from 'lucide-vue-next'

const { t } = useI18n()
</script>

<template>
  <div class=" bg-gradient-to-b from-gray-50  to-white">
    <main class="py-16 lg:py-24">
      <!-- Hero Section -->
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16 md:mb-24">
          <h1 class="font-pramukh text-4xl md:text-5xl lg:text-6xl text-gray-900 mb-4">
            {{ t('about.title') }}
          </h1>
          <p class="font-poppins text-lg text-gray-600 max-w-2xl mx-auto">
            {{ t('about.description') }}
          </p>
        </div>

        <!-- Team Section -->
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-12 max-w-2xl mx-auto">
          <!-- Team Member: Jonathan -->
          <div class="border border-gray-300 rounded-2xl overflow-hidden transition-all duration-300 bg-[#FBFBFB] hover:scale-[1.02]">
            <div class="aspect-w-16 aspect-h-9 bg-blue-50">
              <img 
                src="@/assets/jonathan.jpeg" 
                alt="Jonathan Ahlström" 
                class="w-full h-full object-cover" 
              />
            </div>
            <div class="p-6 md:p-8">
              <div class="flex items-center justify-center mb-4">
                <h2 class="text-2xl md:text-3xl font-pramukh text-gray-900 ">Jonathan Ahlström</h2>
              </div>
              <div class="flex flex-col space-y-2.5 mt-6">
                <a 
                  href="https://www.linkedin.com/in/jonathanahlstrom/" 
                  target="_blank" 
                  rel="noopener" 
                  class="inline-flex items-center justify-center px-4 py-2 rounded-md bg-blue-50 text-blue-600 hover:bg-blue-100 transition-colors duration-200"
                >
                  <Linkedin class="w-4 h-4 mr-2" />
                  <span class="font-medium text-sm">{{ t('about.links.linkedin') }}</span>
                </a>
                <a 
                  href="https://github.com/jonahlstrom" 
                  target="_blank" 
                  rel="noopener" 
                  class="inline-flex items-center justify-center px-4 py-2 rounded-md bg-gray-50 text-gray-600 hover:bg-gray-100 transition-colors duration-200"
                >
                  <Github class="w-4 h-4 mr-2" />
                  <span class="font-medium text-sm">{{ t('about.links.github') }}</span>
                </a>
                <a 
                  href="mailto:jonathan@example.com" 
                  class="inline-flex items-center justify-center px-4 py-2 rounded-md bg-emerald-50 text-emerald-600 hover:bg-emerald-100 transition-colors duration-200"
                >
                  <Mail class="w-4 h-4 mr-2" />
                  <span class="font-medium text-sm">Contact</span>
                </a>
              </div>
            </div>
          </div>

          <!-- Team Member: Fredrik -->
          <div class="border border-gray-300 rounded-2xl overflow-hidden transition-all duration-300 bg-[#FBFBFB] hover:scale-[1.02]">
            <div class="aspect-w-16 aspect-h-9  bg-blue-50">
              <img 
                src="@/assets/fredrik.jpeg" 
                alt="Fredrik Wastring" 
                class="w-full h-full object-cover" 
              />
            </div>
            <div class="p-6 md:p-8">
              <div class="flex items-center justify-center mb-4">
                <h2 class="text-2xl md:text-3xl font-pramukh text-gray-900 ">Fredrik Wastring</h2>
              </div>
              <div class="flex flex-col space-y-2.5 mt-6">
                <a 
                  href="[Friend_LinkedIn]" 
                  target="_blank" 
                  rel="noopener" 
                  class="inline-flex items-center justify-center px-4 py-2 rounded-md bg-blue-50 text-blue-600 hover:bg-blue-100 transition-colors duration-200"
                >
                  <Linkedin class="w-4 h-4 mr-2" />
                  <span class="font-medium text-sm">{{ t('about.links.linkedin') }}</span>
                </a>
                <a 
                  href="[Friend_GitHub]" 
                  target="_blank" 
                  rel="noopener" 
                  class="inline-flex items-center justify-center px-4 py-2 rounded-md bg-gray-50 text-gray-600 hover:bg-gray-100 transition-colors duration-200"
                >
                  <Github class="w-4 h-4 mr-2" />
                  <span class="font-medium text-sm">{{ t('about.links.github') }}</span>
                </a>
                <a 
                  href="mailto:fredrik@example.com" 
                  class="inline-flex items-center justify-center px-4 py-2 rounded-md bg-emerald-50 text-emerald-600 hover:bg-emerald-100 transition-colors duration-200"
                >
                  <Mail class="w-4 h-4 mr-2" />
                  <span class="font-medium text-sm">Contact</span>
                </a>
              </div>
            </div>
          </div>
        </div>

      </div>
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