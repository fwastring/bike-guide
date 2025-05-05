<script setup lang="ts">
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'
import BaseButton from '../components/BaseButton.vue'
</script>

<template>
  <div class="flex flex-col">
   
    <main class="">
      <div class="max-w-3xl mx-auto bg-white p-8 rounded-lg shadow-md">
        <h1 class="text-3xl font-bold text-center text-gray-800 mb-8">About Bike Guide</h1>
        
        <section class="mb-8">
          <h2 class="text-2xl font-semibold text-gray-800 mb-4">Our Mission</h2>
          <p class="text-gray-600 leading-relaxed">
            Bike Guide is dedicated to helping cyclists of all levels find the perfect routes for their adventures.
            Whether you're a beginner looking for a casual ride or an experienced cyclist seeking new challenges,
            we've got you covered.
          </p>
        </section>

        <section class="mb-8">
          <h2 class="text-2xl font-semibold text-gray-800 mb-4">Features</h2>
          <ul class="space-y-2">
            <li class="flex items-start">
              <span class="text-blue-500 mr-2">•</span>
              <span class="text-gray-600">Personalized route recommendations based on your skill level</span>
            </li>
            <li class="flex items-start">
              <span class="text-blue-500 mr-2">•</span>
              <span class="text-gray-600">Detailed route information including distance, elevation, and difficulty</span>
            </li>
            <li class="flex items-start">
              <span class="text-blue-500 mr-2">•</span>
              <span class="text-gray-600">User-friendly interface for easy navigation</span>
            </li>
            <li class="flex items-start">
              <span class="text-blue-500 mr-2">•</span>
              <span class="text-gray-600">Save your favorite routes for future reference</span>
            </li>
          </ul>
        </section>

        <section class="mb-8">
          <h2 class="text-2xl font-semibold text-gray-800 mb-4">Get Started</h2>
          <p class="text-gray-600 mb-6">
            Ready to find your next cycling adventure? Create an account and start exploring our curated routes today!
          </p>
          <div class="flex gap-4 justify-center">
            <BaseButton
              title="Sign Up"
              link="/signup"
              variant="primary"
            />
            <BaseButton
              title="Learn More"
              link="/search"
              variant="secondary"
            />
          </div>
        </section>
      </div>
    </main>

  </div>
</template>

