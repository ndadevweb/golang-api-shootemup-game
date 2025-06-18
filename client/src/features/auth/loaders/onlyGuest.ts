import { redirect } from 'react-router-dom'
import useAuthStore from '@/state/useAuthStore'
import { ROUTES } from '@/routes'

export function onlyGuestLoader() {
    const jwt = useAuthStore.getState().jwt

    if(jwt) {
        return redirect(ROUTES.dashboard)
    }

    return null
}