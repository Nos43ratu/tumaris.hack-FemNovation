import React from "react";
import { LockClosedIcon } from "@heroicons/react/20/solid";
import { useMutation } from "@tanstack/react-query";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { instance } from "@/shared/api/axios.instance";

const SignUp = () => {
  const [email, setEmail] = React.useState("");
  const [password, setPassword] = React.useState("");
  const navigate = useNavigate();

  const mutation = useMutation({
    mutationFn: (data: { email: string; password: string }) =>
      instance.post("/api/sign-in", data),
    onSuccess: (data) => {
      navigate("/cabinet/orders");
    },
  });

  const handleSignIn = (e: any) => {
    e.preventDefault();
    mutation.mutate({ email, password });
  };

  return (
    <div className="flex min-h-full items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <div className="w-full max-w-md space-y-8">
        <div className="flex flex-col items-center">
          <a href="/">
            <svg
              width="200"
              height="26"
              viewBox="0 0 556 74"
              fill="#2563eb"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path d="M52.9806 73.036H50.2007L19.3127 34.9408L2.32434 53.3706V73.036H0.265137V1.06696H2.32434V50.3848L47.8327 1.06696H50.7155L20.7542 33.3964L52.9806 73.036Z" />
              <path d="M79.4562 44.1042C79.4562 52.7529 81.7556 59.4453 86.3545 64.1814C91.022 68.9176 97.5429 71.2857 105.917 71.2857C114.291 71.2857 120.777 68.9176 125.376 64.1814C130.044 59.4453 132.378 52.7529 132.378 44.1042V1.06696H134.437V44.1042C134.437 53.4393 131.931 60.6465 126.921 65.7258C121.979 70.8052 114.977 73.3449 105.917 73.3449C96.8565 73.3449 89.8209 70.8052 84.8101 65.7258C79.8681 60.6465 77.397 53.4393 77.397 44.1042V1.06696H79.4562V44.1042Z" />
              <path d="M204.929 7.55344C202.596 6.04336 199.781 4.87648 196.487 4.0528C193.261 3.16048 190.035 2.71432 186.808 2.71432C180.494 2.71432 175.414 4.01848 171.57 6.6268C167.795 9.16648 165.908 12.5642 165.908 16.8198C165.908 20.5264 166.834 23.5122 168.687 25.7774C170.541 28.0425 172.806 29.7928 175.483 31.0283C178.228 32.1952 181.935 33.4994 186.603 34.9408C191.545 36.4509 195.491 37.8923 198.443 39.2651C201.394 40.5693 203.865 42.5255 205.856 45.1338C207.915 47.7422 208.945 51.1742 208.945 55.4298C208.945 58.9991 207.95 62.1566 205.959 64.9022C203.968 67.5791 201.12 69.6726 197.413 71.1827C193.775 72.6242 189.485 73.3449 184.543 73.3449C180.013 73.3449 175.517 72.4182 171.056 70.565C166.663 68.7117 162.956 66.2406 159.936 63.1518L161.377 61.7104C164.123 64.6619 167.589 66.9957 171.776 68.7117C175.963 70.4277 180.219 71.2857 184.543 71.2857C191.339 71.2857 196.761 69.8442 200.811 66.9614C204.861 64.0785 206.886 60.2346 206.886 55.4298C206.886 51.6546 205.925 48.6345 204.003 46.3694C202.149 44.0356 199.85 42.251 197.104 41.0154C194.359 39.7113 190.618 38.3385 185.882 36.897C181.008 35.387 177.096 33.9798 174.144 32.6757C171.261 31.3029 168.825 29.3466 166.834 26.807C164.844 24.2673 163.848 20.9039 163.848 16.7169C163.848 11.8434 165.942 7.96528 170.129 5.0824C174.316 2.13088 179.876 0.655121 186.808 0.655121C190.172 0.655121 193.57 1.1356 197.002 2.09656C200.502 3.05752 203.488 4.32736 205.959 5.90608L204.929 7.55344Z" />
              <path d="M228.196 1.06696H279.675V3.12615H254.965V73.036H252.906V3.12615H228.196V1.06696Z" />
              <path d="M327.843 0.655121C334.502 0.655121 340.645 2.30248 346.273 5.5972C351.97 8.82328 356.466 13.2162 359.761 18.7761C363.124 24.3359 364.806 30.3762 364.806 36.897C364.806 43.4178 363.124 49.4925 359.761 55.121C356.466 60.6808 351.97 65.1081 346.273 68.4028C340.645 71.6975 334.502 73.3449 327.843 73.3449C321.185 73.3449 315.008 71.6975 309.311 68.4028C303.682 65.1081 299.186 60.6808 295.823 55.121C292.528 49.4925 290.881 43.4178 290.881 36.897C290.881 30.3762 292.528 24.3359 295.823 18.7761C299.186 13.2162 303.682 8.82328 309.311 5.5972C315.008 2.30248 321.185 0.655121 327.843 0.655121ZM327.843 2.71432C321.597 2.71432 315.797 4.25872 310.443 7.34752C305.089 10.4363 300.834 14.6234 297.676 19.9086C294.519 25.1253 292.94 30.7881 292.94 36.897C292.94 43.0746 294.519 48.8061 297.676 54.0914C300.834 59.3766 305.089 63.5637 310.443 66.6525C315.797 69.7413 321.597 71.2857 327.843 71.2857C334.09 71.2857 339.89 69.7413 345.244 66.6525C350.598 63.5637 354.853 59.3766 358.011 54.0914C361.168 48.8061 362.747 43.0746 362.747 36.897C362.747 30.7881 361.168 25.1253 358.011 19.9086C354.853 14.6234 350.598 10.4363 345.244 7.34752C339.89 4.25872 334.09 2.71432 327.843 2.71432Z" />
              <path d="M393.18 1.06696H396.784L428.186 60.6808L459.383 1.06696H462.987V73.036H460.928L460.825 3.0232L428.186 65.1081H427.98L395.239 3.0232V73.036H393.18V1.06696Z" />
              <path d="M553.367 73.036L543.277 51.1055H499.622L489.635 73.036H487.163L520.111 1.06696H522.376L555.735 73.036H553.367ZM500.548 49.0463H542.247L521.243 3.43503L500.548 49.0463Z" />
            </svg>
          </a>
          <h2 className="mt-6 text-center text-2xl font-bold tracking-tight text-gray-900">
            Зарегистрируйтесь
          </h2>
          <p className="mt-2 text-center text-sm text-gray-600">
            или{" "}
            <a
              href="/sign-in"
              className="font-medium text-indigo-600 hover:text-indigo-500"
            >
              войдите
            </a>
          </p>
        </div>
        <form className="mt-8 space-y-6">
          <input type="hidden" name="remember" defaultValue="true" />
          <div className="-space-y-px rounded-md shadow-sm">
            <div>
              <label htmlFor="email-address" className="sr-only">
                Email
              </label>
              <input
                id="email-address"
                name="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                type="email"
                autoComplete="email"
                required
                className="relative block w-full appearance-none rounded-none rounded-t-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
                placeholder="Email"
              />
            </div>
            <div>
              <label htmlFor="password" className="sr-only">
                Пароль
              </label>
              <input
                id="password"
                name="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                type="password"
                autoComplete="current-password"
                required
                className="relative block w-full appearance-none rounded-none rounded-b-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
                placeholder="Пароль"
              />
            </div>
          </div>

          <div>
            <button
              onClick={handleSignIn}
              className="group relative flex w-full justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
            >
              <span className="absolute inset-y-0 left-0 flex items-center pl-3">
                <LockClosedIcon
                  className="h-5 w-5 text-indigo-500 group-hover:text-indigo-400"
                  aria-hidden="true"
                />
              </span>
              Зарегистрироваться
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default SignUp;
