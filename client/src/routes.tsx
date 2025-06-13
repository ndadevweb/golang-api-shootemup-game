import { createBrowserRouter } from "react-router"
import SignupPage from "./pages/auth/Signup"

export const router = createBrowserRouter([
    {
        path: "/",
        Component: SignupPage,
    },
]);
