<script setup lang="ts">
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'
import SignUpCard from '../components/SignUpCard.vue'
</script>

<template>
  <div class="signup">
    <Header />
    <main class="main-content">
      <SignUpCard @signup="handleSignup" />
    </main>
    <Footer />
  </div>
</template>

<style scoped>
.signup {
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
  background-color: #f5f5f5;
}
</style> 