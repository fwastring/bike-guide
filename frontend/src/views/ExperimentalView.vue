<template>
  <div class="h-screen flex flex-col">
    <main class="flex-1 overflow-y-auto">
      <div class="p-8 bg-gray-50">
        <div class="max-w-7xl mx-auto space-y-16">
          <!-- Typography Section -->
          <section>
            <h2 class="text-3xl font-bold text-gray-800 mb-8">Typography</h2>
            <div class="space-y-8">
              <!-- Headings -->
              <div class="space-y-4">
                <h3 class="text-xl font-semibold text-gray-700">Headings</h3>
                <div class="space-y-2">
                  <h1 class="text-5xl font-bold text-gray-800">Heading 1 (text-5xl)</h1>
                  <h2 class="text-4xl font-bold text-gray-800">Heading 2 (text-4xl)</h2>
                  <h3 class="text-3xl font-bold text-gray-800">Heading 3 (text-3xl)</h3>
                  <h4 class="text-2xl font-bold text-gray-800">Heading 4 (text-2xl)</h4>
                  <h5 class="text-xl font-bold text-gray-800">Heading 5 (text-xl)</h5>
                  <h6 class="text-lg font-bold text-gray-800">Heading 6 (text-lg)</h6>
                </div>
              </div>

              <!-- Body Text -->
              <div class="space-y-4">
                <h3 class="text-xl font-semibold text-gray-700">Body Text</h3>
                <div class="space-y-2">
                  <p class="text-base text-gray-600">Regular body text (text-base)</p>
                  <p class="text-sm text-gray-600">Small text (text-sm)</p>
                  <p class="text-xs text-gray-600">Extra small text (text-xs)</p>
                </div>
              </div>

              <!-- Font Weights -->
              <div class="space-y-4">
                <h3 class="text-xl font-semibold text-gray-700">Font Weights</h3>
                <div class="space-y-2">
                  <p class="text-base font-light text-gray-600">Light (font-light)</p>
                  <p class="text-base font-normal text-gray-600">Normal (font-normal)</p>
                  <p class="text-base font-medium text-gray-600">Medium (font-medium)</p>
                  <p class="text-base font-semibold text-gray-600">Semibold (font-semibold)</p>
                  <p class="text-base font-bold text-gray-600">Bold (font-bold)</p>
                </div>
              </div>
            </div>
          </section>

          <!-- Components Section -->
          <section class="space-y-8">
            <h2 class="text-3xl font-bold text-gray-800 mb-8">Components</h2>
            <div class="grid grid-cols-1 ">
              <div class="space-y-8">
                <div>
                  <h3 class="text-xl font-semibold text-gray-700 mb-4">Buttons</h3>
                  <div class="flex gap-4 flex-wrap">
                    <BaseButton title="Primary Button" variant="primary" />
                    <BaseButton title="Secondary Button" variant="secondary" />
                    <BaseButton title="Tertiary Button" variant="tertiary" />
                  </div>
                </div>

                <div>
                  <h3 class="text-xl font-semibold text-gray-700 mb-4">Input Field</h3>
                  <InputField label="Default Input" placeholder="Enter text here" />
                </div>

                <div>
                  <h3 class="text-xl font-semibold text-gray-700 mb-4">Login Card</h3>
                  <LogInCard />
                </div>

                <div>
                  <h3 class="text-xl font-semibold text-gray-700 mb-4">Sign Up Card</h3>
                  <SignUpCard />
                </div>

                <div>
                  <h3 class="text-xl font-semibold text-gray-700 mb-4">Search Card</h3>
                  <div class="space-y-6">
                    <div>
                      <h4 class="text-base font-medium text-gray-600 mb-3">Regular Search Card</h4>
                      <SearchCard @search="handleSearch" />
                    </div>
                    <div>
                      <h4 class="text-base font-medium text-gray-600 mb-3">Mini Search Card</h4>
                      <MiniSearchCard @search="handleSearch" />
                    </div>
                  </div>
                </div>

                <div>
                  <h3 class="text-xl font-semibold text-gray-700 mb-4">Summary Card</h3>
                  <SummaryCard
                    location="Vägvägen 7"
                    difficulty="Medium"
                    distance="20-30 km"
                    @edit="handleEdit"
                  />
                </div>

                <div>
                  <h3 class="text-xl font-semibold text-gray-700 mb-4">Option Card</h3>
                  <OptionCard
                    title="Adelgatan Lund"
                    description="BienVélo loop from Lund centralstation"
                    distance="27.5 km"
                    duration="54 min"
                    difficulty="Easy"
                    imageUrl="/placeholder-route.jpg"
                    @expand="handleExpand"
                  />
                </div>

                <div>
                  <h3 class="text-xl font-semibold text-gray-700 mb-4">Option List Card</h3>
                  <OptionListCard
                    :options="routeOptions"
                    @expand="handleExpand"
                  />
                </div>

                <div>
                  <h3 class="text-xl font-semibold text-gray-700 mb-4">Legacy Search Card</h3>
                  <div class="w-full max-w-3xl mx-auto px-4">
                    <LegacySearchCard @search="handleSearch" />
                  </div>
                </div>
              </div>
            </div>
          </section>

          <!-- Color Scheme Section -->
          <section>
            <h2 class="text-3xl font-bold text-gray-800 mb-8">Color Scheme</h2>
            <div class="space-y-8">
              <!-- Primary Colors -->
              <div class="space-y-4">
                <h3 class="text-xl font-semibold text-gray-700">Primary Colors</h3>
                <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
                  <div class="space-y-2">
                    <div class="h-20 bg-blue-600 rounded-lg"></div>
                    <p class="text-sm text-gray-600">Blue-600 (Primary)</p>
                  </div>
                  <div class="space-y-2">
                    <div class="h-20 bg-blue-700 rounded-lg"></div>
                    <p class="text-sm text-gray-600">Blue-700 (Primary Dark)</p>
                  </div>
                  <div class="space-y-2">
                    <div class="h-20 bg-blue-500 rounded-lg"></div>
                    <p class="text-sm text-gray-600">Blue-500 (Primary Light)</p>
                  </div>
                  <div class="space-y-2">
                    <div class="h-20 bg-blue-50 rounded-lg"></div>
                    <p class="text-sm text-gray-600">Blue-50 (Primary Lightest)</p>
                  </div>
                </div>
              </div>

              <!-- Neutral Colors -->
              <div class="space-y-4">
                <h3 class="text-xl font-semibold text-gray-700">Neutral Colors</h3>
                <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
                  <div class="space-y-2">
                    <div class="h-20 bg-gray-800 rounded-lg"></div>
                    <p class="text-sm text-gray-600">Gray-800 (Dark)</p>
                  </div>
                  <div class="space-y-2">
                    <div class="h-20 bg-gray-600 rounded-lg"></div>
                    <p class="text-sm text-gray-600">Gray-600 (Medium)</p>
                  </div>
                  <div class="space-y-2">
                    <div class="h-20 bg-gray-300 rounded-lg"></div>
                    <p class="text-sm text-gray-600">Gray-300 (Light)</p>
                  </div>
                  <div class="space-y-2">
                    <div class="h-20 bg-gray-50 rounded-lg"></div>
                    <p class="text-sm text-gray-600">Gray-50 (Lightest)</p>
                  </div>
                </div>
              </div>

              <!-- Status Colors -->
              <div class="space-y-4">
                <h3 class="text-xl font-semibold text-gray-700">Status Colors</h3>
                <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
                  <div class="space-y-2">
                    <div class="h-20 bg-green-500 rounded-lg"></div>
                    <p class="text-sm text-gray-600">Green-500 (Success)</p>
                  </div>
                  <div class="space-y-2">
                    <div class="h-20 bg-red-500 rounded-lg"></div>
                    <p class="text-sm text-gray-600">Red-500 (Error)</p>
                  </div>
                  <div class="space-y-2">
                    <div class="h-20 bg-yellow-500 rounded-lg"></div>
                    <p class="text-sm text-gray-600">Yellow-500 (Warning)</p>
                  </div>
                  <div class="space-y-2">
                    <div class="h-20 bg-blue-500 rounded-lg"></div>
                    <p class="text-sm text-gray-600">Blue-500 (Info)</p>
                  </div>
                </div>
              </div>

              <!-- Result Card Example -->
              <div class="space-y-4 mt-8">
                <h3 class="text-xl font-semibold text-gray-700">Result Card</h3>
                <ResultCard
                  title="Scenic Coastal Route"
                  distance="20km"
                  elevation="150m"
                  difficulty="Moderate"
                  description="A beautiful ride along the coast with stunning ocean views. Perfect for intermediate cyclists looking for a scenic route."
                  :highlights="[
                    'Stunning coastal views',
                    'Well-maintained paths',
                    'Multiple rest stops',
                    'Café at the halfway point'
                  ]"
                  mapUrl="https://maps.example.com/route/1"
                  gpxUrl="https://maps.example.com/route/1.gpx"
                />
              </div>
            </div>
          </section>
        </div>
      </div>
      <Footer />

    </main>
  </div>
</template>

<script setup lang="ts">
import Header from '../components/Header.vue'
import Footer from '../components/Footer.vue'
import BaseButton from '../components/BaseButton.vue'
import InputField from '../components/InputField.vue'
import LogInCard from '../components/LogInCard.vue'
import SignUpCard from '../components/SignUpCard.vue'
import SearchCard from '../components/SearchCard.vue'
import SummaryCard from '../components/SummaryCard.vue'
import OptionCard from '../components/OptionCard.vue'
import OptionListCard from '../components/OptionListCard.vue'
import ResultCard from '../components/ResultCard.vue'
import type { RouteOption } from '@/types/route'
import { useRouter } from 'vue-router'
import { useSearchStore } from '@/stores/searchStore'
import LegacySearchCard from '../components/LegacySearchCard.vue'
import MiniSearchCard from '../components/MiniSearchCard.vue'

const handleSearch = (searchData: { address: string; difficulty: string; distance: string }) => {
  console.log('Searching with:', searchData)
}

const handleEdit = () => {
  console.log('Edit clicked')
}

const handleExpand = (index: number) => {
  console.log('Expand clicked for option:', index)
}

const routeOptions: RouteOption[] = [
  {
    title: 'Scenic Mountain Trail',
    description: 'A beautiful mountain trail with breathtaking views',
    distance: '10km',
    duration: '2 hours',
    difficulty: 'Easy' as const,
    imageUrl: '/images/mountain-trail.jpg'
  },
  {
    title: 'Forest Adventure',
    description: 'Explore the dense forest and its wildlife',
    distance: '15km',
    duration: '3 hours',
    difficulty: 'Medium' as const,
    imageUrl: '/images/forest-adventure.jpg'
  },
  {
    title: 'Coastal Route',
    description: 'Ride along the stunning coastline',
    distance: '20km',
    duration: '4 hours',
    difficulty: 'Hard' as const,
    imageUrl: '/images/coastal-route.jpg'
  }
]
</script> 