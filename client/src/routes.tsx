import { createBrowserRouter } from 'react-router'
import { onlyAuthenticatedLoader } from './features/auth/loaders/onlyAuthenticated'
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
    },
    {
        path: ROUTES.signup,
        Component: SignupPage
    },
    {
        path: ROUTES.signin,
        Component: SigninPage
    },

    {
        path: ROUTES.dashboard,
        Component: DashboardPage,
        loader: onlyAuthenticatedLoader
    }
]);