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
          <div class="border border-gray-300 rounded-2xl overflow-hidden transition-all duration-300 bg-[#FBFBFB]  flex flex-col md:flex-row" role="listitem">
            <div class="md:w-1/4 flex items-center justify-center">
              <img 
                src="@/assets/jonathan.jpeg" 
                alt="Jonathan Ahlström" 
                class="w-32 h-32 object-cover rounded-full p-2" 
              />
            </div>
            <div class="p-6 md:p-8 md:w-2/3">
              <h2 class="font-poppins text-xl font-bold text-gray-900 mb-4">Jonathan Ahlström</h2>
              <p class="text-gray-600 mb-6">Cycling enthusiast & tech lover. Building the future of bike navigation.</p>
              <div class="flex flex-wrap gap-2 justify-center md:justify-start" role="list" aria-label="Jonathan's social links">
                <a 
                  href="https://www.linkedin.com/in/jonathanahlstrom/" 
                  target="_blank" 
                  rel="noopener" 
                  class="inline-flex items-center px-3 py-1.5 rounded-md bg-gradient-to-r from-[#8E7DBE] to-[#A7A3E2] text-white hover:from-[#7D6CAE] hover:to-[#9692D2] transition-colors duration-200"
                  role="listitem"
                  aria-label="Visit Jonathan's LinkedIn profile"
                >
                  <Linkedin class="w-4 h-4 mr-1.5" aria-hidden="true" />
                  <span class="text-sm">LinkedIn</span>
                </a>
                <a 
                  href="https://github.com/jonahlstrom" 
                  target="_blank" 
                  rel="noopener" 
                  class="inline-flex items-center px-3 py-1.5 rounded-md bg-gradient-to-r from-[#8E7DBE] to-[#A7A3E2] text-white hover:from-[#7D6CAE] hover:to-[#9692D2] transition-colors duration-200"
                  role="listitem"
                  aria-label="Visit Jonathan's GitHub profile"
                >
                  <Github class="w-4 h-4 mr-1.5" aria-hidden="true" />
                  <span class="text-sm">GitHub</span>
                </a>
                <a 
                  href="mailto:jonathan@example.com" 
                  class="inline-flex items-center px-3 py-1.5 rounded-md bg-gradient-to-r from-[#8E7DBE] to-[#A7A3E2] text-white hover:from-[#7D6CAE] hover:to-[#9692D2] transition-colors duration-200"
                  role="listitem"
                  aria-label="Contact Jonathan via email"
                >
                  <Mail class="w-4 h-4 mr-1.5" aria-hidden="true" />
                  <span class="text-sm">Email</span>
                </a>
              </div>
            </div>
          </div>

          <!-- Team Member: Fredrik -->
          <div class="border border-gray-300 rounded-2xl overflow-hidden transition-all duration-300 bg-[#FBFBFB]  flex flex-col md:flex-row" role="listitem">
            <div class="md:w-1/4 flex items-center justify-center">
              <img 
                src="@/assets/fredrik.jpeg" 
                alt="Fredrik Wastring" 
                class="w-32 h-32 object-cover rounded-full p-2" 
              />
            </div>
            <div class="p-6 md:p-8 md:w-2/3">
              <h2 class="font-poppins text-xl font-bold text-gray-900 mb-4">Fredrik Wastring</h2>
              <p class="text-gray-600 mb-6">Adventure seeker & code wizard. Making bike routes awesome.</p>
              <div class="flex flex-wrap gap-2 justify-center md:justify-start" role="list" aria-label="Fredrik's social links">
                <a 
                  href="[Friend_LinkedIn]" 
                  target="_blank" 
                  rel="noopener" 
                  class="inline-flex items-center px-3 py-1.5 rounded-md bg-gradient-to-r from-[#8E7DBE] to-[#A7A3E2] text-white hover:from-[#7D6CAE] hover:to-[#9692D2] transition-colors duration-200"
                  role="listitem"
                  aria-label="Visit Fredrik's LinkedIn profile"
                >
                  <Linkedin class="w-4 h-4 mr-1.5" aria-hidden="true" />
                  <span class="text-sm">LinkedIn</span>
                </a>
                <a 
                  href="[Friend_GitHub]" 
                  target="_blank" 
                  rel="noopener" 
                  class="inline-flex items-center px-3 py-1.5 rounded-md bg-gradient-to-r from-[#8E7DBE] to-[#A7A3E2] text-white hover:from-[#7D6CAE] hover:to-[#9692D2] transition-colors duration-200"
                  role="listitem"
                  aria-label="Visit Fredrik's GitHub profile"
                >
                  <Github class="w-4 h-4 mr-1.5" aria-hidden="true" />
                  <span class="text-sm">GitHub</span>
                </a>
                <a 
                  href="mailto:fredrik@example.com" 
                  class="inline-flex items-center px-3 py-1.5 rounded-md bg-gradient-to-r from-[#8E7DBE] to-[#A7A3E2] text-white hover:from-[#7D6CAE] hover:to-[#9692D2] transition-colors duration-200"
                  role="listitem"
                  aria-label="Contact Fredrik via email"
                >
                  <Mail class="w-4 h-4 mr-1.5" aria-hidden="true" />
                  <span class="text-sm">Email</span>
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