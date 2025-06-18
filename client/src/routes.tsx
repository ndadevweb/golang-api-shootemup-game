import { createBrowserRouter } from 'react-router'
import { onlyAuthenticatedLoader } from './features/auth/loaders/onlyAuthenticated'
import { onlyGuestLoader } from './features/auth/loaders/onlyGuest'
import HomePage from './pages/homepage'
import SignupPage from './pages/auth/Signup'
import SigninPage from './pages/auth/Signin'
import DashboardPage from './pages/dashboard'

export const ROUTES = {
    'homepage': '/',
    'signup': '/signup',
    'signin': '/signin',

    'dashboard': '/dashboard',
}

export const router = createBrowserRouter([
    {
        path: ROUTES.homepage,
        Component: HomePage,
        loader: onlyGuestLoader
    },
    {
        path: ROUTES.signup,
        Component: SignupPage,
        loader: onlyGuestLoader
    },
    {
        path: ROUTES.signin,
        Component: SigninPage,
        loader: onlyGuestLoader
    },

    {
        path: ROUTES.dashboard,
        Component: DashboardPage,
        loader: onlyAuthenticatedLoader
    }
]);