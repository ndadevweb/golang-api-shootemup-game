import { redirect } from 'react-router'
import useAuthStore from '@/state/useAuthStore'
import { ROUTES } from '@/routes'

export function signoutLoader() {
    useAuthStore.getState().signout()

    return redirect(ROUTES.homepage)
}

