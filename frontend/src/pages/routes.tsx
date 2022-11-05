import { createBrowserRouter } from "react-router-dom";

import Home from "@/pages/home";
import SignIn from "./sign-in";
import SignUp from "@/pages/sign-up";
import Order from "@/pages/order";
import Cabinet from "./cabinet";
import Layout from "@/shared/ui/Layout";

export const router = createBrowserRouter([
  {
    path: "/",
    children: [
      { path: "/", element: <Home /> },
      {
        path: "/cabinet",
        element: <Layout />,
        children: [{ path: "/cabinet/orders", element: <Order /> }],
      },
    ],
  },
  { path: "/sign-in", element: <SignIn /> },
  { path: "/sign-up", element: <SignUp /> },
]);
