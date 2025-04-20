<script setup lang="ts">
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'
import BaseButton from '../components/BaseButton.vue'
</script>

<template>
  <div class="min-h-screen flex flex-col">
    <Header />
    <main class="flex-1 p-8 bg-gray-50">
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
    <Footer />
  </div>
</template>

<style scoped>
.about {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.main-content {
  flex: 1;
  padding: 2rem;
  background-color: #f5f5f5;
}

.about-container {
  max-width: 800px;
  margin: 0 auto;
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

h1 {
  text-align: center;
  margin-bottom: 2rem;
  color: #333;
}

.about-section {
  margin-bottom: 2rem;
}

h2 {
  margin-bottom: 1rem;
  color: #333;
}

p {
  line-height: 1.6;
  color: #666;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  padding: 0.5rem 0;
  color: #666;
}

li::before {
  content: "•";
  color: #007bff;
  font-weight: bold;
  display: inline-block;
  width: 1em;
  margin-left: -1em;
}

.cta-buttons {
  display: flex;
  gap: 1rem;
  justify-content: center;
  margin-top: 2rem;
}
</style> 