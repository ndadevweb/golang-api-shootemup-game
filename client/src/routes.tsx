import { createBrowserRouter } from "react-router"
import SignupPage from "./pages/auth/Signup"
import SigninPage from "./pages/auth/Signin"

export const router = createBrowserRouter([
    {
        path: "/",
        Component: SignupPage,
    },
    {
        path: "/signin",
        Component: SigninPage
    }
]);
