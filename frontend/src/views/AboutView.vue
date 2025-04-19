<script setup lang="ts">
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'
import BaseButton from '../components/BaseButton.vue'
</script>

<template>
  <div class="about">
    <Header />
    <main class="main-content">
      <div class="about-container">
        <h1>About Bike Guide</h1>
        
        <section class="about-section">
          <h2>Our Mission</h2>
          <p>
            Bike Guide is dedicated to helping cyclists of all levels find the perfect routes for their adventures.
            Whether you're a beginner looking for a casual ride or an experienced cyclist seeking new challenges,
            we've got you covered.
          </p>
        </section>

        <section class="about-section">
          <h2>Features</h2>
          <ul>
            <li>Personalized route recommendations based on your skill level</li>
            <li>Detailed route information including distance, elevation, and difficulty</li>
            <li>User-friendly interface for easy navigation</li>
            <li>Save your favorite routes for future reference</li>
          </ul>
        </section>

        <section class="about-section">
          <h2>Get Started</h2>
          <p>
            Ready to find your next cycling adventure? Create an account and start exploring our curated routes today!
          </p>
          <div class="cta-buttons">
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