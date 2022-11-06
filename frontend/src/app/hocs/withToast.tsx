import React, { ComponentType } from "react";

import { ToastContainer } from "react-toastify";

import "react-toastify/dist/ReactToastify.css";

export const withToast =
  <T,>(Component: ComponentType<T>) =>
  (props: T) =>
    (
      <>
        <Component {...props} />
        <ToastContainer />
      </>
    );

export default withToast;
