<script setup lang="ts">
import Footer from '../components/Footer.vue'
</script>

<template>
  <div class="flex flex-col min-h-screen bg-gray-50">
    <main class="flex-1 flex flex-col items-center justify-center py-16 md:py-24">
      <div class="max-w-2xl w-full mx-auto bg-white rounded-3xl shadow-lg px-8 py-12 flex flex-col items-center">
        <h1 class="font-pramukh text-5xl sm:text-6xl text-gray-900 mb-4 text-center">Meet the Developers</h1>
        <p class="font-poppins text-lg text-gray-600 mb-8 text-center max-w-xl">
          Bike Guide was created by two friends passionate about cycling and technology. The project was built from scratch with a focus on user experience, modern web technologies, and a love for adventure. Every feature, from route search to the animated landing page, was designed and coded collaboratively, with countless hours spent on design, development, and testing to make the perfect cycling companion.
        </p>
        <div class="flex flex-row gap-10 mb-8">
          <div class="flex flex-col items-center">
            <img  alt="Jonathan Ahlström" class="w-32 h-32 rounded-full object-cover border-4 border-blue-200 shadow-lg mb-3" />
            <span class="font-pramukh text-xl text-gray-900">Jonathan Ahlström</span>
            <div class="flex gap-3 mt-1">
              <a href="https://www.linkedin.com/in/jonathanahlstrom/" target="_blank" rel="noopener" class="text-blue-600 font-poppins hover:underline">LinkedIn</a>
              <a href="https://github.com/jonahlstrom" target="_blank" rel="noopener" class="text-gray-700 font-poppins hover:underline">GitHub</a>
            </div>
          </div>
          <div class="flex flex-col items-center">
            <img  alt="[Friend's Name]" class="w-32 h-32 rounded-full object-cover border-4 border-blue-200 shadow-lg mb-3" />
            <span class="font-pramukh text-xl text-gray-900">[Friend's Name]</span>
            <div class="flex gap-3 mt-1">
              <a href="[Friend_LinkedIn]" target="_blank" rel="noopener" class="text-blue-600 font-poppins hover:underline">LinkedIn</a>
              <a href="[Friend_GitHub]" target="_blank" rel="noopener" class="text-gray-700 font-poppins hover:underline">GitHub</a>
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

