import { Fragment, useEffect, useState } from "react";
import { Dialog, Disclosure, Transition } from "@headlessui/react";
import { XMarkIcon } from "@heroicons/react/24/outline";
import { useAutoAnimate } from "@formkit/auto-animate/react";
import create from "zustand";
import {
  FunnelIcon,
  MagnifyingGlassIcon,
  MinusIcon,
  PlusIcon,
  StarIcon,
} from "@heroicons/react/20/solid";
import { useSearchParams } from "react-router-dom";
import clsx from "clsx";
import { useQueries, useQuery } from "@tanstack/react-query";
import { useCartStore } from "@/shared/ui/Layout";

const fakeApi = (): Promise<Product[]> =>
  new Promise((resolve) => {
    return setTimeout(() => {
      resolve([
        {
          id: 1,
          name: "Бутылка",
          description:
            "Очень классная бутыл, купил себе, маме, отцу, брату, сестренке, жене, теще, свату",
          product_categories: "",
          product_sizes: "",
          colors: "",
          weight: "5",
          product_comments: "",
          price: "48",
          rating: 4,
          imageSrc:
            "https://tailwindui.com/img/ecommerce-images/category-page-04-image-card-01.jpg",
        },
        {
          id: 2,
          name: "Бутылка",
          description:
            "Очень классная бутыл, купил себе, маме, отцу, брату, сестренке, жене, теще, свату",
          product_categories: "",
          product_sizes: "",
          colors: "",
          weight: "5",
          product_comments: "",
          price: "48",
          rating: 4,
          imageSrc:
            "https://tailwindui.com/img/ecommerce-images/category-page-04-image-card-02.jpg",
        },
        {
          id: 3,
          name: "Блокнот",
          description:
            "Очень классная бутыл, купил себе, маме, отцу, брату, сестренке, жене, теще, свату",
          product_categories: "",
          product_sizes: "",
          colors: "",
          weight: "5",
          product_comments: "",
          price: "48",
          rating: 4,
          imageSrc:
            "https://tailwindui.com/img/ecommerce-images/category-page-04-image-card-03.jpg",
        },
        {
          id: 4,
          name: "ручка",
          description:
            "Очень классная бутыл, купил себе, маме, отцу, брату, сестренке, жене, теще, свату",
          product_categories: "",
          product_sizes: "",
          colors: "",
          weight: "5",
          product_comments: "",
          price: "48",
          rating: 4,
          imageSrc:
            "https://tailwindui.com/img/ecommerce-images/category-page-04-image-card-04.jpg",
        },
      ]);
    }, 1000);
  });

const subCategories = [
  { name: "Торты", href: "#" },
  { name: "Свечи", href: "#" },
];
const filters = [
  {
    id: "Средвства ухода",
    name: "Средвства ухода",
    options: [
      { value: "Мыло", label: "Мыло", checked: false },
      { value: "Шампунь", label: "Шампунь", checked: false },
      { value: "Гель для душа", label: "Гель для душа", checked: false },
    ],
  },
  {
    id: "Вкусняшки",
    name: "Вкусняшки",
    options: [
      { value: "Торты", label: "Торты", checked: false },
      { value: "Моти", label: "Моти", checked: false },
      { value: "Пирожное", label: "Пирожное", checked: false },
    ],
  },
];

const useMobileFilterOpen = create<{ open: boolean; handleClick: () => void }>(
  (set) => ({
    open: false,
    handleClick: () => set((state) => ({ open: !state.open })),
  })
);

const Orders = () => {
  const handleOpen = useMobileFilterOpen((state) => state.handleClick);

  return (
    <div className="bg-white">
      <div>
        <MobileFilter />

        <main className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="flex flex-col items-baseline justify-between border-b border-gray-200 pt-12 pb-6 md:flex-row">
            <h1 className="w-full text-4xl font-bold tracking-tight text-blue-800">
              Товары и услуги
            </h1>

            <div className="mt-4 flex w-full items-center justify-between md:mt-0">
              <Search />

              <div className="flex items-center">
                <button
                  type="button"
                  className="-m-2 ml-4 p-2 text-gray-400 hover:text-gray-500 sm:ml-6 lg:hidden"
                  onClick={handleOpen}
                >
                  <span className="sr-only">Filters</span>
                  <FunnelIcon className="h-5 w-5" aria-hidden="true" />
                </button>
              </div>
            </div>
          </div>

          <section aria-labelledby="products-heading" className="pt-6 pb-24">
            <h2 id="products-heading" className="sr-only">
              Products
            </h2>

            <div className="grid grid-cols-1 gap-x-8 gap-y-10 lg:grid-cols-4">
              <Filter />

              <Products />
            </div>
          </section>
        </main>
      </div>
    </div>
  );
};

const Search = () => {
  const [searchParams, setSearchParams] = useSearchParams();

  const [value, setValue] = useState("");

  useEffect(() => {
    const time = setTimeout(() => {
      setSearchParams({ search: value });
    }, 1000);

    return () => clearTimeout(time);
  }, [value]);

  useEffect(() => {
    const search = searchParams.get("search");
    if (search) setValue(search);
  }, []);

  return (
    <div className="flex flex-1 items-center justify-center px-2 lg:ml-6 lg:justify-end">
      <div className="w-full max-w-lg lg:max-w-xs">
        <label htmlFor="search" className="sr-only">
          Search
        </label>
        <div className="relative">
          <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
            <MagnifyingGlassIcon
              className="h-5 w-5 text-gray-400"
              aria-hidden="true"
            />
          </div>
          <input
            id="search"
            name="search"
            value={value}
            onChange={(e) => setValue(e.target.value)}
            className="block w-full rounded-md border border-gray-300 bg-white py-2 pl-10 pr-3 leading-5 placeholder-gray-500 shadow-sm focus:border-blue-600 focus:placeholder-gray-400 focus:outline-none focus:ring-1 focus:ring-blue-600 sm:text-sm"
            placeholder="Search"
            type="search"
          />
        </div>
      </div>
    </div>
  );
};

const Products = () => {
  const { data, isLoading } = useQuery({
    queryKey: ["products"],
    queryFn: fakeApi,
  });

  const [parent] = useAutoAnimate<HTMLDivElement | null>();

  return (
    <div className="lg:col-span-3">
      <div className="h-full bg-white">
        {isLoading ? (
          <div className="flex h-full w-full items-center justify-center text-blue-600">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="10 10 80 80"
              className={clsx("animate-spin")}
              width={40}
              height={40}
              preserveAspectRatio="xMidYMid"
              stroke="currentColor"
            >
              <circle
                cx="50"
                cy="50"
                fill="none"
                stroke-width="10"
                r="35"
                stroke-dasharray="164.93361431346415 56.97787143782138"
              />
            </svg>
          </div>
        ) : (
          <div
            className="grid grid-cols-1 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 xl:gap-x-8"
            ref={parent}
          >
            {data?.map((product) => (
              <ProductItem key={product.id} product={product} />
            ))}
          </div>
        )}
      </div>
    </div>
  );
};

const ProductItem = ({ product }: { product: Product }) => {
  const cart = useCartStore();
  const [open, setOpen] = useState(false);

  return (
    <>
      <button
        onClick={() => setOpen(true)}
        key={product.id}
        className="group items-center"
      >
        <div className="aspect-w-1 aspect-h-1 xl:aspect-w-7 xl:aspect-h-8 w-full overflow-hidden rounded-lg bg-gray-200">
          <img
            src={product.imageSrc}
            className="h-full w-full object-cover object-center group-hover:opacity-75"
          />
        </div>
        <h3 className="mt-4 text-sm text-gray-700">{product.name}</h3>
        <p className="mt-1 flex items-center text-lg font-medium text-gray-900">
          {product.price}
          <img
            className="object-fit h-2.5 w-2"
            src="https://upload.wikimedia.org/wikipedia/commons/f/f8/Tenge_symbol.svg"
            alt=""
          />
        </p>
      </button>
      <Transition.Root show={open} as={Fragment}>
        <Dialog as="div" className="relative z-10" onClose={setOpen}>
          <Transition.Child
            as={Fragment}
            enter="ease-out duration-300"
            enterFrom="opacity-0"
            enterTo="opacity-100"
            leave="ease-in duration-200"
            leaveFrom="opacity-100"
            leaveTo="opacity-0"
          >
            <div className="fixed inset-0 hidden bg-gray-500 bg-opacity-75 transition-opacity md:block" />
          </Transition.Child>

          <div className="fixed inset-0 z-10 overflow-y-auto">
            <div className="flex min-h-full items-stretch justify-center text-center md:items-center md:px-2 lg:px-4">
              <Transition.Child
                as={Fragment}
                enter="ease-out duration-300"
                enterFrom="opacity-0 translate-y-4 md:translate-y-0 md:scale-95"
                enterTo="opacity-100 translate-y-0 md:scale-100"
                leave="ease-in duration-200"
                leaveFrom="opacity-100 translate-y-0 md:scale-100"
                leaveTo="opacity-0 translate-y-4 md:translate-y-0 md:scale-95"
              >
                <Dialog.Panel className="flex w-full transform text-left text-base transition md:my-8 md:max-w-2xl md:px-4 lg:max-w-4xl">
                  <div className="relative flex w-full items-center overflow-hidden bg-white px-4 pt-14 pb-8 shadow-2xl sm:px-6 sm:pt-8 md:p-6 lg:p-8">
                    <button
                      type="button"
                      className="absolute top-4 right-4 text-gray-400 hover:text-gray-500 sm:top-8 sm:right-6 md:top-6 md:right-6 lg:top-8 lg:right-8"
                      onClick={() => setOpen(false)}
                    >
                      <span className="sr-only">Close</span>
                      <XMarkIcon className="h-6 w-6" aria-hidden="true" />
                    </button>

                    <div className="grid w-full grid-cols-1 items-start gap-y-8 gap-x-6 sm:grid-cols-12 lg:gap-x-8">
                      <div className="aspect-w-2 aspect-h-3 overflow-hidden rounded-lg bg-gray-100 sm:col-span-4 lg:col-span-5">
                        <img
                          src={product.imageSrc}
                          className="object-cover object-center"
                        />
                      </div>
                      <div className="flex h-full flex-col sm:col-span-8 lg:col-span-7">
                        <h2 className="text-2xl font-bold text-gray-900 sm:pr-12">
                          {product.name}
                        </h2>

                        <p>{product.description}</p>

                        <section
                          aria-labelledby="information-heading"
                          className="mt-2"
                        >
                          <h3 id="information-heading" className="sr-only">
                            Product information
                          </h3>

                          <p className="flex items-center text-2xl text-gray-900">
                            {product.price}{" "}
                            <img
                              className="object-fit h-2.5 w-2"
                              src="https://upload.wikimedia.org/wikipedia/commons/f/f8/Tenge_symbol.svg"
                              alt=""
                            />
                          </p>

                          {/* Reviews */}
                          <div className="mt-6">
                            <div className="flex items-center">
                              <div className="flex items-center">
                                {[0, 1, 2, 3, 4].map((rating) => (
                                  <StarIcon
                                    key={rating}
                                    className={clsx(
                                      product.rating > rating
                                        ? "text-yellow-500"
                                        : "text-gray-200",
                                      "h-5 w-5 flex-shrink-0"
                                    )}
                                    aria-hidden="true"
                                  />
                                ))}
                              </div>
                            </div>
                          </div>
                          <p className="mt-6 text-sm font-medium">
                            Weight: {product.weight}
                          </p>
                        </section>

                        <section
                          aria-labelledby="options-heading"
                          className="mt-auto"
                        >
                          <div>
                            {cart.products.findIndex(
                              (e) => e.id === product.id
                            ) !== -1 ? (
                              <button
                                disabled
                                className="mt-6 flex w-full items-center justify-center rounded-md border border-transparent bg-gray-600 py-3 px-8 text-base font-medium text-white"
                              >
                                Добавлено в корзину
                              </button>
                            ) : (
                              <button
                                onClick={(e) => {
                                  e.preventDefault();
                                  cart.addProduct(product);
                                  setOpen(false);
                                }}
                                className="mt-6 flex w-full items-center justify-center rounded-md border border-transparent bg-indigo-600 py-3 px-8 text-base font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                              >
                                Добавить в корзину
                              </button>
                            )}
                          </div>
                        </section>
                      </div>
                    </div>
                  </div>
                </Dialog.Panel>
              </Transition.Child>
            </div>
          </div>
        </Dialog>
      </Transition.Root>
    </>
  );
};

const Filter = () => {
  return (
    <form className="hidden lg:block">
      <h3 className="sr-only">Categories</h3>
      <ul
        role="list"
        className="space-y-4 border-b border-gray-200 pb-6 text-sm font-medium text-gray-900"
      >
        {subCategories.map((category) => (
          <li key={category.name}>
            <a href={category.href}>{category.name}</a>
          </li>
        ))}
      </ul>

      {filters.map((section) => (
        <Disclosure
          as="div"
          key={section.id}
          className="border-b border-gray-200 py-6"
        >
          {({ open }) => (
            <>
              <h3 className="-my-3 flow-root">
                <Disclosure.Button className="flex w-full items-center justify-between bg-white py-3 text-sm text-gray-400 hover:text-gray-500">
                  <span className="font-medium text-gray-900">
                    {section.name}
                  </span>
                  <span className="ml-6 flex items-center">
                    {open ? (
                      <MinusIcon className="h-5 w-5" aria-hidden="true" />
                    ) : (
                      <PlusIcon className="h-5 w-5" aria-hidden="true" />
                    )}
                  </span>
                </Disclosure.Button>
              </h3>
              <Disclosure.Panel className="pt-6">
                <div className="space-y-4">
                  {section.options.map((option, optionIdx) => (
                    <div key={option.value} className="flex items-center">
                      <input
                        id={`filter-${section.id}-${optionIdx}`}
                        name={`${section.id}[]`}
                        defaultValue={option.value}
                        type="checkbox"
                        defaultChecked={option.checked}
                        className="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                      />
                      <label
                        htmlFor={`filter-${section.id}-${optionIdx}`}
                        className="ml-3 text-sm text-gray-600"
                      >
                        {option.label}
                      </label>
                    </div>
                  ))}
                </div>
              </Disclosure.Panel>
            </>
          )}
        </Disclosure>
      ))}
    </form>
  );
};

const MobileFilter = () => {
  const mobileFilter = useMobileFilterOpen((state) => state);

  return (
    <Transition.Root show={mobileFilter.open} as={Fragment}>
      <Dialog
        as="div"
        className="relative z-40 lg:hidden"
        onClose={mobileFilter.handleClick}
      >
        <Transition.Child
          as={Fragment}
          enter="transition-opacity ease-linear duration-300"
          enterFrom="opacity-0"
          enterTo="opacity-100"
          leave="transition-opacity ease-linear duration-300"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
        >
          <div className="fixed inset-0 bg-black bg-opacity-25" />
        </Transition.Child>

        <div className="fixed inset-0 z-40 flex">
          <Transition.Child
            as={Fragment}
            enter="transition ease-in-out duration-300 transform"
            enterFrom="translate-x-full"
            enterTo="translate-x-0"
            leave="transition ease-in-out duration-300 transform"
            leaveFrom="translate-x-0"
            leaveTo="translate-x-full"
          >
            <Dialog.Panel className="relative ml-auto flex h-full w-full max-w-xs flex-col overflow-y-auto bg-white py-4 pb-12 shadow-xl">
              <div className="flex items-center justify-between px-4">
                <h2 className="text-lg font-medium text-gray-900">Фильтр</h2>
                <button
                  type="button"
                  className="-mr-2 flex h-10 w-10 items-center justify-center rounded-md bg-white p-2 text-gray-400"
                  onClick={mobileFilter.handleClick}
                >
                  <span className="sr-only">Close menu</span>
                  <XMarkIcon className="h-6 w-6" aria-hidden="true" />
                </button>
              </div>

              {/* Filters */}
              <form className="mt-4 border-t border-gray-200">
                <h3 className="sr-only">Categories</h3>
                <ul role="list" className="px-2 py-3 font-medium text-gray-900">
                  {subCategories.map((category) => (
                    <li key={category.name}>
                      <a href={category.href} className="block px-2 py-3">
                        {category.name}
                      </a>
                    </li>
                  ))}
                </ul>

                {filters.map((section) => (
                  <Disclosure
                    as="div"
                    key={section.id}
                    className="border-t border-gray-200 px-4 py-6"
                  >
                    {({ open }) => (
                      <>
                        <h3 className="-mx-2 -my-3 flow-root">
                          <Disclosure.Button className="flex w-full items-center justify-between bg-white px-2 py-3 text-gray-400 hover:text-gray-500">
                            <span className="font-medium text-gray-900">
                              {section.name}
                            </span>
                            <span className="ml-6 flex items-center">
                              {open ? (
                                <MinusIcon
                                  className="h-5 w-5"
                                  aria-hidden="true"
                                />
                              ) : (
                                <PlusIcon
                                  className="h-5 w-5"
                                  aria-hidden="true"
                                />
                              )}
                            </span>
                          </Disclosure.Button>
                        </h3>
                        <Disclosure.Panel className="pt-6">
                          <div className="space-y-6">
                            {section.options.map((option, optionIdx) => (
                              <div
                                key={option.value}
                                className="flex items-center"
                              >
                                <input
                                  id={`filter-mobile-${section.id}-${optionIdx}`}
                                  name={`${section.id}[]`}
                                  defaultValue={option.value}
                                  type="checkbox"
                                  defaultChecked={option.checked}
                                  className="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-500"
                                />
                                <label
                                  htmlFor={`filter-mobile-${section.id}-${optionIdx}`}
                                  className="ml-3 min-w-0 flex-1 text-gray-500"
                                >
                                  {option.label}
                                </label>
                              </div>
                            ))}
                          </div>
                        </Disclosure.Panel>
                      </>
                    )}
                  </Disclosure>
                ))}
              </form>
            </Dialog.Panel>
          </Transition.Child>
        </div>
      </Dialog>
    </Transition.Root>
  );
};

export default Orders;
