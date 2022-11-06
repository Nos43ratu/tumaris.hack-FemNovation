import { RouterProvider } from "react-router";

import { router } from "@/pages/routes";

import { withHocs } from "./hocs";

export const App = () => <RouterProvider router={router} />;

export default withHocs(App);
