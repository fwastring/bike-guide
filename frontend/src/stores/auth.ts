import { defineStore } from 'pinia'

interface User {
  id: string
  email: string
  name: string
  picture: string
}

interface AuthState {
  user: User | null
  isAuthenticated: boolean
}

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    user: null,
    isAuthenticated: false,
  }),

  actions: {
    setUser(user: User | null) {
      this.user = user
      this.isAuthenticated = !!user
    },

    async checkAuth() {
      try {
        const response = await fetch('/api/auth/check', {
          credentials: 'include',
        })
        
        if (response.ok) {
          const user = await response.json()
          this.setUser(user)
          return true
        }
        
        this.setUser(null)
        return false
      } catch (error) {
        console.error('Auth check failed:', error)
        this.setUser(null)
        return false
      }
    },

    async logout() {
      try {
        await fetch('/api/auth/logout', {
          method: 'POST',
          credentials: 'include',
        })
        this.setUser(null)
      } catch (error) {
        console.error('Logout failed:', error)
      }
    },
  },
}) 