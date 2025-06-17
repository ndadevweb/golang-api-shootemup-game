import { create } from 'zustand'

interface AuthState {
  jwt: string | null;
  setJwt: (jwt: string) => void;
  signout: () => void;
}

const useAuthStore = create<AuthState>((set) => ({
  jwt: null,
  setJwt: (jwt) => set({ jwt }),
  signout: () => set({ jwt: null }),
}))

export default useAuthStore