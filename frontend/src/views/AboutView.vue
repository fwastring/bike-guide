<script setup lang="ts">
import Footer from '../components/Footer.vue'
import { useI18n } from 'vue-i18n'
import { Linkedin, Github, Mail, ArrowRight } from 'lucide-vue-next'

const { t } = useI18n()
</script>

<template>
  <div class=" bg-gradient-to-b from-gray-50   to-white">
    <main class="py-16 lg:py-24">
      <!-- Hero Section -->
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 ">
        <div class="text-center mb-16 md:mb-24">
          <h1 class="font-pramukh text-5xl sm:text-6xl lg:text-7xl text-gray-900 mb-4">
            {{ t('about.title') }}
          </h1>
          <p class="font-poppins text-sm sm:text-lg text-gray-600 max-w-2xl mx-auto">
            {{ t('about.description') }}
          </p>
        </div>

        <!-- Team Section -->
        <div class="grid grid-cols-1 gap-8 max-w-4xl mx-auto" role="list" aria-label="Team members">
          <!-- Team Member: Jonathan -->
          <div class="border-[0.5px] border-gray-300 rounded-2xl overflow-hidden transition-all duration-300 bg-[#FBFBFB]  flex flex-col md:flex-row" role="listitem">
            <div class="md:w-1/4 pt-4 md:pt-0 flex items-center justify-center">
              <img 
                src="@/assets/jonathan.jpeg" 
                alt="Jonathan Ahlström" 
                class="w-32 h-32 object-cover rounded-full p-2" 
              />
            </div>
            <div class="p-6 md:p-8 md:w-2/3">
              <h2 class="font-poppins text-xl font-bold text-gray-900 mb-4 text-center md:text-left">{{ t('about.team.jonathan.name') }}</h2>
              <div class="w-full h-[0.5px] bg-gray-200 mx-auto md:mx-0 mb-6"></div>
              <p class="text-gray-600 mb-6 text-center md:text-left">{{ t('about.team.jonathan.description') }}</p>
              <div class="flex flex-col sm:flex-row gap-2 justify-center md:justify-start items-center sm:items-start" role="list" aria-label="Jonathan's social links">
                <a 
                  href="https://www.linkedin.com/in/jonathan-ahlström-8700a71a1" 
                  target="_blank" 
                  rel="noopener" 
                  class="inline-flex items-center justify-center w-full sm:w-32 px-6 py-2 rounded-full bg-blue-50 text-purple-500 hover:bg-blue-100 transition-colors duration-200"
                  role="listitem"
                  aria-label="Visit Jonathan's LinkedIn profile"
                >
                  <Linkedin class="w-4 h-4 mr-1.5" aria-hidden="true" />
                  <span class="text-sm">{{ t('about.links.linkedin') }}</span>
                </a>
                <a 
                  href="https://github.com/joah-s" 
                  target="_blank" 
                  rel="noopener" 
                  class="inline-flex items-center justify-center w-full sm:w-32 px-6 py-2 rounded-full bg-blue-50 text-purple-500 hover:bg-blue-100 transition-colors duration-200"
                  role="listitem"
                  aria-label="Visit Jonathan's GitHub profile"
                >
                  <Github class="w-4 h-4 mr-1.5" aria-hidden="true" />
                  <span class="text-sm">{{ t('about.links.github') }}</span>
                </a>
                <a 
                  href="mailto:jonathan.ahlstrom@outlook.com" 
                  class="inline-flex items-center justify-center w-full sm:w-32 px-6 py-2 rounded-full bg-blue-50 text-purple-500 hover:bg-blue-100 transition-colors duration-200"
                  role="listitem"
                  aria-label="Contact Jonathan via email"
                >
                  <Mail class="w-4 h-4 mr-1.5" aria-hidden="true" />
                  <span class="text-sm">{{ t('about.links.email') }}</span>
                </a>
              </div>
            </div>
          </div>

          <!-- Team Member: Fredrik -->
          <div class="border-[0.5px] border-gray-300 rounded-2xl overflow-hidden transition-all duration-300 bg-[#FBFBFB]  flex flex-col md:flex-row" role="listitem">
            <div class="md:w-1/4 flex items-center pt-4 md:pt-0 justify-center">
              <img 
                src="@/assets/fredrik.jpeg" 
                alt="Fredrik Wastring" 
                class="w-32 h-32 object-cover rounded-full p-2" 
              />
            </div>
            <div class="p-6 md:p-8 md:w-2/3">
              <h2 class="font-poppins text-xl font-bold text-gray-900 mb-4 text-center md:text-left">{{ t('about.team.fredrik.name') }}</h2>
              <div class="w-full h-[0.5px] bg-gray-200 mx-auto md:mx-0 mb-6"></div>
              <p class="text-gray-600 mb-6 text-center md:text-left">{{ t('about.team.fredrik.description') }}</p>
              <div class="flex flex-col sm:flex-row gap-2 justify-center md:justify-start items-center sm:items-start" role="list" aria-label="Fredrik's social links">
                <a 
                  href="https://www.linkedin.com/in/fredrik-wastring-95357b195/" 
                  target="_blank" 
                  rel="noopener" 
                  class="inline-flex items-center justify-center w-full sm:w-32 px-6 py-2 rounded-full bg-blue-50 text-purple-500 hover:bg-blue-100 transition-colors duration-200"
                  role="listitem"
                  aria-label="Visit Fredrik's LinkedIn profile"
                >
                  <Linkedin class="w-4 h-4 mr-1.5" aria-hidden="true" />
                  <span class="text-sm">{{ t('about.links.linkedin') }}</span>
                </a>
                <a 
                  href="https://github.com/fwastring" 
                  target="_blank" 
                  rel="noopener" 
                  class="inline-flex items-center justify-center w-full sm:w-32 px-6 py-2 rounded-full bg-blue-50 text-purple-500 hover:bg-blue-100 transition-colors duration-200"
                  role="listitem"
                  aria-label="Visit Fredrik's GitHub profile"
                >
                  <Github class="w-4 h-4 mr-1.5" aria-hidden="true" />
                  <span class="text-sm">{{ t('about.links.github') }}</span>
                </a>
                <a 
                  href="mailto:fredrik@wastring.com" 
                  class="inline-flex items-center justify-center w-full sm:w-32 px-6 py-2 rounded-full bg-blue-50 text-purple-500 hover:bg-blue-100 transition-colors duration-200"
                  role="listitem"
                  aria-label="Contact Fredrik via email"
                >
                  <Mail class="w-4 h-4 mr-1.5" aria-hidden="true" />
                  <span class="text-sm">{{ t('about.links.email') }}</span>
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

/* Custom focus styles */
a:focus-visible,
button:focus-visible {
  outline: 2px solid #8E7DBE;
  outline-offset: 2px;
  box-shadow: 0 0 0 4px rgba(142, 125, 190, 0.2);
  border-radius: 0.375rem;
  transition: outline-offset 0.2s ease;
}

/* Remove default focus styles */
a:focus:not(:focus-visible),
button:focus:not(:focus-visible) {
  outline: none;
}

/* Ensure focus styles work in high contrast mode */
@media (forced-colors: active) {
  a:focus-visible,
  button:focus-visible {
    outline: 2px solid CanvasText;
    outline-offset: 2px;
  }
}
</style>