<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
  <footer class="footer">
    <div class="footer-content">
      <div class="footer-section">
        <h3>Bike Guide</h3>
        <p>Your ultimate cycling companion</p>
      </div>
      <div class="footer-section">
        <h4>Quick Links</h4>
        <RouterLink to="/about">About</RouterLink>
        <RouterLink to="/search">Search</RouterLink>
      </div>
      <div class="footer-section">
        <h4>Legal</h4>
        <RouterLink to="/privacy">Privacy Policy</RouterLink>
        <RouterLink to="/terms">Terms of Service</RouterLink>
      </div>
    </div>
    <div class="footer-bottom">
      <p>&copy; {{ new Date().getFullYear() }} Bike Guide. All rights reserved.</p>
    </div>
  </footer>
</template>

<style scoped>
.footer {
  background-color: #333;
  color: #fff;
  padding: 2rem 0;
  margin-top: auto;
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 2rem;
  padding: 0 1rem;
}

.footer-section h3,
.footer-section h4 {
  margin-bottom: 1rem;
}

.footer-section a {
  display: block;
  color: #fff;
  text-decoration: none;
  margin-bottom: 0.5rem;
  opacity: 0.8;
  transition: opacity 0.3s;
}

.footer-section a:hover {
  opacity: 1;
}

.footer-bottom {
  text-align: center;
  padding-top: 2rem;
  margin-top: 2rem;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}
</style> 