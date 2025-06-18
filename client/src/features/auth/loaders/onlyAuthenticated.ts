import { redirect } from 'react-router-dom'
import useAuthStore from '@/state/useAuthStore'
import { ROUTES } from '@/routes'

export function onlyAuthenticatedLoader() {
    const jwt = useAuthStore.getState().jwt

    if(!jwt) {
        return redirect(ROUTES.homepage)
    }

    return null
}