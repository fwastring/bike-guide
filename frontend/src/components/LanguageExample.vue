<template>
  <div class="language-example">
    <h1>{{ t('common.welcome') }}</h1>
    
    <div class="language-switcher">
      <button @click="switchLanguage('en')">English</button>
      <button @click="switchLanguage('sv')">Svenska</button>
    </div>

    <nav>
      <ul>
        <li>{{ t('navigation.home') }}</li>
        <li>{{ t('navigation.about') }}</li>
        <li>{{ t('navigation.contact') }}</li>
      </ul>
    </nav>

    <div class="bike-info">
      <h2>{{ t('bike.title') }}</h2>
      <p>{{ t('bike.type') }}: {{ bikeType }}</p>
      <p>{{ t('bike.brand') }}: {{ bikeBrand }}</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'

const { t, locale } = useI18n()

const bikeType = 'Mountain Bike'
const bikeBrand = 'Trek'

const switchLanguage = (lang: string) => {
  locale.value = lang
}
</script>

<style scoped>
.language-example {
  padding: 20px;
}

.language-switcher {
  margin: 20px 0;
}

.language-switcher button {
  margin-right: 10px;
  padding: 5px 10px;
}

nav ul {
  list-style: none;
  padding: 0;
}

nav li {
  margin: 10px 0;
}

.bike-info {
  margin-top: 20px;
  padding: 15px;
  border: 1px solid #ccc;
  border-radius: 4px;
}
</style> 