<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
  <header class="header">
    <nav class="nav">
      <RouterLink to="/" class="logo">Bike Guide</RouterLink>
      <div class="nav-links">
        <RouterLink to="/about">About</RouterLink>
        <RouterLink to="/search">Search</RouterLink>
      </div>
    </nav>
  </header>
</template>

<style scoped>
.header {
  width: 100%;
  padding: 1rem;
  background-color: #ffffff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.nav {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo {
  font-size: 1.5rem;
  font-weight: bold;
  text-decoration: none;
  color: #333;
}

.nav-links {
  display: flex;
  gap: 1rem;
}

.nav-links a {
  text-decoration: none;
  color: #333;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  transition: background-color 0.3s;
}

.nav-links a:hover {
  background-color: #f0f0f0;
}
</style> 