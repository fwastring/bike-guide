import { defineStore } from 'pinia'
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'

export const useScrollStore = defineStore('scroll', () => {
  const hasScrolled = ref(false)
  const scrollPosition = ref(0)
  let ticking = false
  const route = useRoute()

  // Watch for state changes
  watch(hasScrolled, (newValue, oldValue) => {
    console.log('hasScrolled state changed:', { old: oldValue, new: newValue })
  })

  watch(scrollPosition, (newValue, oldValue) => {
    console.log('scrollPosition changed:', { old: oldValue, new: newValue })
  })

  const checkScroll = () => {
    const currentScroll = window.scrollY
    const oldPosition = scrollPosition.value
    scrollPosition.value = currentScroll
    
    // Different threshold for landing page
    const scrollThreshold = route.path === '/' ? 200 : 100
    const newHasScrolled = currentScroll > scrollThreshold
    const oldHasScrolled = hasScrolled.value
    
    if (newHasScrolled !== oldHasScrolled) {
      console.log('Scroll state changed:', {
        position: currentScroll,
        oldPosition,
        threshold: scrollThreshold,
        oldHasScrolled,
        newHasScrolled,
        path: route.path
      })
      hasScrolled.value = newHasScrolled
    }
  }

  const handleScroll = () => {
    if (!ticking) {
      window.requestAnimationFrame(() => {
        checkScroll()
        ticking = false
      })
      ticking = true
    }
  }

  const initScrollListener = () => {
    console.log('Initializing scroll listener')
    // Remove any existing listeners first
    window.removeEventListener('scroll', handleScroll)
    // Add the new listener
    window.addEventListener('scroll', handleScroll, { passive: true })
    // Initial check
    checkScroll()
  }

  const removeScrollListener = () => {
    console.log('Removing scroll listener')
    window.removeEventListener('scroll', handleScroll)
  }

  // Reset scroll state when route changes
  const resetScrollState = () => {
    console.log('Resetting scroll state')
    const oldHasScrolled = hasScrolled.value
    const oldPosition = scrollPosition.value
    hasScrolled.value = false
    scrollPosition.value = 0
    console.log('State reset:', {
      oldHasScrolled,
      newHasScrolled: false,
      oldPosition,
      newPosition: 0
    })
    checkScroll() // Check new position immediately
  }

  // Watch for route changes
  watch(() => route.path, (newPath, oldPath) => {
    console.log('Route changed in store:', { old: oldPath, new: newPath })
    if (newPath === '/') {
      resetScrollState()
    }
  })

  return {
    hasScrolled,
    scrollPosition,
    initScrollListener,
    removeScrollListener,
    resetScrollState
  }
}) 