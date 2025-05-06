import { useRouter } from 'vue-router'

export function useScrollTo() {
  const router = useRouter()

  const scrollToSection = (sectionId: string) => {
    if (router.currentRoute.value.path === '/') {
      const section = document.getElementById(sectionId)
      if (section) {
        const header = document.querySelector('header')
        const headerHeight = header?.offsetHeight || 56

        const elementPosition = section.getBoundingClientRect().top
        const offsetPosition = elementPosition + window.pageYOffset - headerHeight

        window.scrollTo({
          top: offsetPosition,
          behavior: 'smooth'
        })
      }
    } else {
      router.push(`/#${sectionId}`)
    }
  }

  return {
    scrollToSection
  }
} 