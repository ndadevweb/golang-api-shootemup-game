import { createBrowserRouter } from "react-router"
import HomePage from './pages/homepage'
import SignupPage from "./pages/auth/Signup"
import SigninPage from "./pages/auth/Signin"

export const ROUTES = {
    'homepage': '/',
    'signup': '/signup',
    'signin': '/signin'
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
    }
]);