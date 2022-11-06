import { ComponentType } from "react";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

export const queryClient = new QueryClient();

export const withQueryClient =
  <T,>(Component: ComponentType<T>) =>
  (props: T) =>
    (
      <QueryClientProvider client={queryClient}>
        <Component {...props} />
      </QueryClientProvider>
    );

export default withQueryClient;
