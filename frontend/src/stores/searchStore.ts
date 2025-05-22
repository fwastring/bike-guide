import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { RouteOption } from '@/types/route'

export interface SearchParams {
  address: string
  difficulty: string
  distance: string
}

export const useSearchStore = defineStore('search', () => {
  const searchParams = ref<SearchParams>({
    address: '',
    difficulty: '',
    distance: ''
  })

  const routeOptions = ref<RouteOption[]>([])
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const setSearchParams = (params: SearchParams) => {
    searchParams.value = params
  }

  const fetchRouteOptions = async () => {
    isLoading.value = true
    error.value = null

    try {
      // TODO: Replace with actual API call
      // Simulated API call for now
      await new Promise(resolve => setTimeout(resolve, 1000))

      routeOptions.value = [
        {
          title: 'Scenic Mountain Trail',
          address: searchParams.value.address,
          description: 'A beautiful mountain trail with breathtaking views',
          distance: '10',
          duration: '2h',
          difficulty: 'Easy' as const,
          imageUrl: '/images/mountain-trail.jpg'
        },
        {
          title: 'Forest Adventure',
          address: searchParams.value.address,
          description: 'Explore the dense forest and its wildlife',
          distance: '15',
          duration: '3h',
          difficulty: 'Medium' as const,
          imageUrl: '/images/forest-adventure.jpg'
        },
        {
          title: 'Coastal Route',
          address: searchParams.value.address,
          description: 'Ride along the stunning coastline',
          distance: '20',
          duration: '4h',
          difficulty: 'Hard' as const,
          imageUrl: '/images/coastal-route.jpg'
        }
      ]
    } catch (e) {
      error.value = 'Failed to fetch route options'
      console.error(e)
    } finally {
      isLoading.value = false
    }
  }

  return {
    searchParams,
    routeOptions,
    isLoading,
    error,
    setSearchParams,
    fetchRouteOptions
  }
})
