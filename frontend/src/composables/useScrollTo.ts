import { useRouter } from 'vue-router'

export function useScrollTo() {
  const router = useRouter()

  const scrollToSection = (sectionId: string) => {
    const section = document.getElementById(sectionId)
    if (section) {
      section.scrollIntoView()
    } else if (router.currentRoute.value.path !== '/') {
      // Only route if we're not on the home page and couldn't find the section
      router.push(`/#${sectionId}`)
    }
  }

  return {
    scrollToSection
  }
} 