import { createBrowserRouter } from "react-router-dom";

import Home from "@/pages/home";
import Layout from "@/shared/ui/Layout";
import SignIn from "./sign-in";
import SignUp from "@/pages/sign-up";
import FindOrder from "@/pages/find-order";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <Layout />,
    children: [
      { path: "/", element: <Home /> },
      {
        path: "/cabinet",
        children: [
          { path: "/cabinet/sign-in", element: <SignIn /> },
          { path: "/cabinet/sign-up", element: <SignUp /> },
          { path: "/cabinet/find-order", element: <FindOrder /> },
        ],
      },
    ],
  },
]);
