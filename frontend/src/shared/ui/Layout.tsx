import React, { Fragment, useEffect, useState } from "react";
import { Outlet, useNavigate } from "react-router-dom";
import { Dialog, Menu, Popover, Transition } from "@headlessui/react";
import {
  ShoppingBagIcon,
  BellIcon,
  XMarkIcon,
} from "@heroicons/react/24/outline";
import create from "zustand";
import { StarIcon } from "@heroicons/react/20/solid";
import clsx from "clsx";
import { useAutoAnimate } from "@formkit/auto-animate/react";
import { useMutation, useQuery } from "@tanstack/react-query";
import axios from "axios";
import { toast } from "react-toastify";
import { instance } from "@/shared/api/axios.instance";

const user = {
  name: "Whitney Francis",
  email: "whitney@example.com",
  imageUrl:
    "https://images.unsplash.com/photo-1517365830460-955ce3ccd263?ixlib=rb-=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=8&w=256&h=256&q=80",
};

const Layout = () => {
  return (
    <div className="min-h-full">
      <header className="bg-white shadow">
        <div className="mx-auto max-w-7xl px-2 sm:px-4 lg:px-8">
          <Popover className="flex h-16 justify-between">
            <div className="flex px-2 lg:px-0">
              <div className="flex flex-shrink-0 items-center">
                <a href="/">
                  <svg
                    className="h-[20px] w-[150px] md:h-[26px] md:w-[200px]"
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
              </div>
            </div>

            <div className="ml-4 flex items-center">
              <Cart />

              <Avatar />
            </div>
          </Popover>
        </div>
      </header>
      <Outlet />
    </div>
  );
};

export const useCartStore = create<{
  open: boolean;
  handleOpen: (value: boolean) => void;
  products: Product[];
  addProduct: (product: Product) => void;
  removeProduct: (id: number) => void;
}>((set, get) => ({
  open: false,
  handleOpen: (value) => set({ open: value }),
  products: [],
  addProduct: (product: Product) => {
    set((state) => {
      const products = [...state.products];
      const index = products.findIndex((p) => p.id === product.id);
      if (index === -1) {
        products.push(product);
      }
      return { products };
    });
  },
  removeProduct: (id: number) => {
    set((state) => {
      const products = [...state.products];

      return { products: products.filter((p) => p.id !== id) };
    });
  },
}));

const images = [
  "https://tailwindui.com/img/ecommerce-images/category-page-04-image-card-01.jpg",
  "https://tailwindui.com/img/ecommerce-images/category-page-04-image-card-02.jpg",
  "https://tailwindui.com/img/ecommerce-images/category-page-04-image-card-03.jpg",
  "https://tailwindui.com/img/ecommerce-images/category-page-04-image-card-04.jpg",
];

const checkout = (data: {
  client_id: number;
  shop_id: number;
  product_id: number;
}) => instance.post("/api/orders", data);

const Cart = () => {
  const navigate = useNavigate();
  const cart = useCartStore();
  const [parent] = useAutoAnimate<HTMLUListElement | null>();

  const [userData, setUserData] = useState<null | UserData>(null);
  useEffect(() => {
    const user = localStorage.getItem("userData");
    user && setUserData(JSON.parse(user));
  }, []);

  const mutate = useMutation({
    mutationFn: (data: {
      client_id: number;
      shop_id: number;
      product_id: number;
    }) => {
      console.log(data);
      return checkout(data);
    },
  });

  const handleCheckout = async () => {
    await Promise.all(
      cart.products.map((product) => {
        return mutate.mutate({
          client_id: userData.id,
          shop_id: product.shop_id,
          product_id: product.id,
        });
      })
    );
    navigate("/cabinet");
  };

  return (
    <>
      <div className="flex items-center">
        <button
          onClick={() => cart.handleOpen(true)}
          className="inline-flex items-center justify-center rounded-md p-2 text-gray-400 hover:bg-gray-100 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-blue-500"
        >
          <ShoppingBagIcon className="block h-6 w-6" aria-hidden="true" />
        </button>
      </div>
      <Transition.Root show={cart.open} as={Fragment}>
        <Dialog
          as="div"
          className="relative z-10"
          onClose={() => cart.handleOpen(false)}
        >
          <Transition.Child
            as={Fragment}
            enter="ease-in-out duration-500"
            enterFrom="opacity-0"
            enterTo="opacity-100"
            leave="ease-in-out duration-500"
            leaveFrom="opacity-100"
            leaveTo="opacity-0"
          >
            <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" />
          </Transition.Child>

          <div className="fixed inset-0 overflow-hidden">
            <div className="absolute inset-0 overflow-hidden">
              <div className="pointer-events-none fixed inset-y-0 right-0 flex max-w-full">
                <Transition.Child
                  as={Fragment}
                  enter="transform transition ease-in-out duration-500 sm:duration-700"
                  enterFrom="translate-x-full"
                  enterTo="translate-x-0"
                  leave="transform transition ease-in-out duration-500 sm:duration-700"
                  leaveFrom="translate-x-0"
                  leaveTo="translate-x-full"
                >
                  <Dialog.Panel className="pointer-events-auto w-screen max-w-md">
                    <div className="flex h-full flex-col overflow-y-scroll bg-white shadow-xl">
                      <div className="flex flex-1 flex-col overflow-y-auto py-6 px-4 sm:px-6">
                        <div className="flex items-start justify-between">
                          <Dialog.Title className="text-lg font-medium text-gray-900">
                            Корзина
                          </Dialog.Title>
                          <div className="ml-3 flex h-7 items-center">
                            <button
                              type="button"
                              className="-m-2 p-2 text-gray-400 hover:text-gray-500"
                              onClick={() => cart.handleOpen(false)}
                            >
                              <span className="sr-only">Close panel</span>
                              <XMarkIcon
                                className="h-6 w-6"
                                aria-hidden="true"
                              />
                            </button>
                          </div>
                        </div>

                        <div className="mt-8 h-full">
                          <div className="flow-root h-full">
                            {cart.products.length === 0 ? (
                              <div className="flex h-full w-full items-center justify-center">
                                <span className="text-base font-medium text-gray-900">
                                  Ваша корзина пуста =_=
                                </span>
                              </div>
                            ) : (
                              <ul
                                role="list"
                                className="-my-6 divide-y divide-gray-200"
                                ref={parent}
                              >
                                {cart.products.map((product) => (
                                  <li key={product.id} className="flex py-6">
                                    <div className="h-24 w-24 flex-shrink-0 overflow-hidden rounded-md border border-gray-200">
                                      <img
                                        src={images[product.id ?? 0]}
                                        className="h-full w-full object-cover object-center"
                                      />
                                    </div>

                                    <div className="ml-4 flex flex-1 flex-col">
                                      <div>
                                        <div className="flex justify-between text-base font-medium text-gray-900">
                                          <h3>
                                            <span>{product.name}</span>
                                          </h3>
                                          <p className="ml-4 flex items-center ">
                                            {product.price}{" "}
                                            <img
                                              className="object-fit h-2.5 w-2"
                                              src="https://upload.wikimedia.org/wikipedia/commons/f/f8/Tenge_symbol.svg"
                                              alt=""
                                            />
                                          </p>
                                        </div>
                                        <div className="mt-1">
                                          <p className="text-xs line-clamp-2">
                                            {product.description}
                                          </p>
                                        </div>
                                        <div
                                          aria-labelledby="information-heading"
                                          className="mt-2 flex justify-between"
                                        >
                                          <div className="flex items-center justify-between">
                                            <div className="flex items-center">
                                              {[0, 1, 2, 3, 4].map((rating) => (
                                                <StarIcon
                                                  key={rating}
                                                  className={clsx(
                                                    product.rating > rating
                                                      ? "text-yellow-500"
                                                      : "text-gray-200",
                                                    "h-4 w-4 flex-shrink-0"
                                                  )}
                                                  aria-hidden="true"
                                                />
                                              ))}
                                            </div>
                                          </div>
                                          <div className="flex text-sm">
                                            <button
                                              type="button"
                                              onClick={() => {
                                                cart.removeProduct(product.id);
                                              }}
                                              className="font-medium text-indigo-600 hover:text-indigo-500"
                                            >
                                              удалить
                                            </button>
                                          </div>
                                        </div>
                                      </div>
                                    </div>
                                  </li>
                                ))}
                              </ul>
                            )}
                          </div>
                        </div>
                      </div>

                      {cart.products.length > 0 && (
                        <div className="border-t border-gray-200 py-6 px-4 sm:px-6">
                          <div className="flex justify-between text-base font-medium text-gray-900">
                            <p>Общая стоимость</p>
                            <p className="flex items-center">
                              {cart?.products?.reduce(
                                (prev, curr) => prev + parseInt(curr.price),
                                0
                              )}{" "}
                              <img
                                className="object-fit h-2.5 w-2"
                                src="https://upload.wikimedia.org/wikipedia/commons/f/f8/Tenge_symbol.svg"
                                alt=""
                              />
                            </p>
                          </div>
                          <div className=" mt-6 flex">
                            <button
                              onClick={() => {
                                handleCheckout();
                                cart.handleOpen(false);
                              }}
                              className="flex w-full items-center justify-center rounded-md border border-transparent bg-indigo-600 px-6 py-3 text-base font-medium text-white shadow-sm hover:bg-indigo-700"
                            >
                              Купить
                            </button>
                          </div>
                        </div>
                      )}
                    </div>
                  </Dialog.Panel>
                </Transition.Child>
              </div>
            </div>
          </div>
        </Dialog>
      </Transition.Root>
    </>
  );
};

const getUser = (email: string): Promise<ApiResponse<UserData>> =>
  instance.get("/api/users?email=" + "Mdidara@quirduck.khs");

const Avatar = () => {
  const navigate = useNavigate();
  const [userData, setUserData] = useState<null | UserData>(null);

  useEffect(() => {
    const user = localStorage.getItem("userData");
    user && setUserData(JSON.parse(user));
  }, []);
  const { data } = useQuery(["user"], () => getUser(userData?.email ?? ""), {
    enabled: userData !== null,
  });

  console.log(data);

  const mutation = useMutation({
    mutationFn: () => instance.get("/api/sign-out"),
    onSuccess: (data: any) => {
      return navigate("/");
    },
  });

  const handleSignOut = () => {
    window.localStorage.removeItem("userData");
    mutation.mutate();
  };

  return (
    <Menu as="div" className="relative ml-3">
      <div>
        <Menu.Button className="flex rounded-full bg-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800">
          <span className="sr-only">Open user menu</span>
          <img className="h-8 w-8 rounded-full" src={user.imageUrl} alt="" />
        </Menu.Button>
      </div>
      <Transition
        as={Fragment}
        enter="transition ease-out duration-100"
        enterFrom="transform opacity-0 scale-95"
        enterTo="transform opacity-100 scale-100"
        leave="transition ease-in duration-75"
        leaveFrom="transform opacity-100 scale-100"
        leaveTo="transform opacity-0 scale-95"
      >
        <Menu.Items className="absolute right-0 z-10 mt-2 w-48 origin-top-right rounded-md bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none">
          <Menu.Item>
            {({ active }) => (
              <a
                href="/cabinet"
                className={clsx(
                  active ? "bg-gray-100" : "",
                  "block px-4 py-2 text-sm text-gray-700"
                )}
              >
                Мой профиль
              </a>
            )}
          </Menu.Item>
          <Menu.Item>
            {({ active }) => (
              <a
                href="/cabinet/orders"
                className={clsx(
                  active ? "bg-gray-100" : "",
                  "block px-4 py-2 text-sm text-gray-700"
                )}
              >
                Найти товар
              </a>
            )}
          </Menu.Item>
          <Menu.Item>
            {({ active }) => (
              <button
                onClick={handleSignOut}
                className={clsx(
                  active ? "bg-gray-100" : "",
                  "block w-full px-4 py-2 text-left text-sm text-gray-700"
                )}
              >
                Выйти
              </button>
            )}
          </Menu.Item>
        </Menu.Items>
      </Transition>
    </Menu>
  );
};

export default Layout;
