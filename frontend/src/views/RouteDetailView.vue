<template>
  <div class="min-h-screen bg-gray-50">
    <Header />
    <main class="container mx-auto px-4 py-8">
      <div class="max-w-7xl mx-auto">
        <!-- Back Button -->
        <button
          @click="router.back()"
          class="mb-6 flex items-center text-blue-600 hover:text-blue-700"
        >
          <span class="mr-2">←</span> Back to routes
        </button>

        <div v-if="currentRoute" class="bg-white rounded-2xl shadow-sm overflow-hidden">
          <!-- Route Image -->
          <img
            :src="currentRoute.imageUrl"
            :alt="currentRoute.title"
            class="w-full h-64 object-cover"
          />

          <!-- Route Content -->
          <div class="p-6">
            <h1 class="text-3xl font-bold text-gray-800 mb-4">{{ currentRoute.title }}</h1>
            
            <div class="grid grid-cols-3 gap-4 mb-6">
              <div class="text-center">
                <p class="text-sm text-gray-500">Distance</p>
                <p class="text-lg font-semibold">{{ currentRoute.distance }}</p>
              </div>
              <div class="text-center">
                <p class="text-sm text-gray-500">Duration</p>
                <p class="text-lg font-semibold">{{ currentRoute.duration }}</p>
              </div>
              <div class="text-center">
                <p class="text-sm text-gray-500">Difficulty</p>
                <p class="text-lg font-semibold">{{ currentRoute.difficulty }}</p>
              </div>
            </div>

            <p class="text-gray-600 mb-6">{{ currentRoute.description }}</p>

            <button
              class="w-full bg-blue-600 text-white py-3 px-6 rounded-lg hover:bg-blue-700 transition-colors"
            >
              Start Navigation
            </button>
          </div>
        </div>

        <div v-else class="text-center py-12">
          <p class="text-gray-600">Route not found.</p>
          <button
            @click="router.push('/options')"
            class="mt-4 text-blue-600 hover:text-blue-700"
          >
            Return to route options
          </button>
        </div>
      </div>
    </main>
    <Footer />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useSearchStore } from '@/stores/searchStore'
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'

const router = useRouter()
const route = useRoute()
const searchStore = useSearchStore()

const routeId = parseInt(route.params.id as string)
const currentRoute = computed(() => searchStore.routeOptions[routeId])
</script> 