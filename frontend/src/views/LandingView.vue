<script setup lang="ts">
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'
import BaseButton from '../components/BaseButton.vue'
</script>

<template>
  <div class="landing">
    <Header />
    <main class="main-content">
      <section class="hero">
        <h1>Find Your Perfect Bike Route</h1>
        <p>Discover the best cycling routes tailored to your preferences and skill level</p>
        <div class="cta-buttons">
          <BaseButton
            title="Get Started"
            link="/onboarding"
            variant="primary"
          />
          <BaseButton
            title="Learn More"
            link="/about"
            variant="secondary"
          />
        </div>
      </section>
    </main>
    <Footer />
  </div>
</template>

<style scoped>
.landing {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.main-content {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem;
}

.hero {
  text-align: center;
  max-width: 800px;
}

.hero h1 {
  font-size: 3rem;
  margin-bottom: 1rem;
  color: #333;
}

.hero p {
  font-size: 1.25rem;
  color: #666;
  margin-bottom: 2rem;
}

.cta-buttons {
  display: flex;
  gap: 1rem;
  justify-content: center;
}
</style> 