import React, { Fragment, useEffect, useMemo, useState } from "react";
import { ScaleIcon, XMarkIcon } from "@heroicons/react/24/outline";
import {
  BanknotesIcon,
  CheckCircleIcon,
  CheckIcon,
  ChevronRightIcon,
  HandThumbUpIcon,
  QuestionMarkCircleIcon,
  StarIcon,
  UserIcon,
} from "@heroicons/react/20/solid";
import { instance } from "@/shared/api/axios.instance";
import { useMutation, useQuery } from "@tanstack/react-query";
import { useAutoAnimate } from "@formkit/auto-animate/react";
import clsx from "clsx";
import { Dialog, Transition } from "@headlessui/react";
import order from "@/pages/order";
import axios from "axios";
import { queryClient } from "@/app/app";

const eventTypes = {
  created: { icon: UserIcon, bgColorClass: "bg-yellow-400" },
  payed: { icon: CheckIcon, bgColorClass: "bg-blue-500" },
  packed: { icon: HandThumbUpIcon, bgColorClass: "bg-green-500" },
  delivered: { icon: CheckIcon, bgColorClass: "bg-green-500" },
  canceled: { icon: XMarkIcon, bgColorClass: "bg-red-500" },
};

function classNames(...classes) {
  return classes.filter(Boolean).join(" ");
}

const Cabinet = () => {
  return (
    <main className="flex-1 pb-8">
      {/* Page header */}
      <Head />

      <div className="mt-8">
        <Info />

        <h2 className="mx-auto mt-8 max-w-6xl px-4 text-lg font-medium leading-6 text-gray-900 sm:px-6 lg:px-8">
          Мои заказы
        </h2>

        <OrderList />
      </div>
    </main>
  );
};

const getOrder = (id: number): Promise<ApiResponse<Order[]>> =>
  instance.get(`/api/users/${id}/orders`);

const status = {
  color: {
    "0": "bg-yellow-100 text-yellow-800",
    "1": "bg-green-100 text-green-800",
    "2": "bg-red-100 text-red-800",
    "3": "bg-gray-100 text-gray-800",
    "4": "bg-gray-100 text-gray-800",
    "5": "bg-gray-100 text-gray-800",
    "-10": "bg-gray-100 text-gray-800",
    "-1": "bg-gray-100 text-gray-800",
    "-2": "bg-gray-100 text-gray-800",
  },
  text: {
    0: "Ожидает ответа",
    1: "Ожидает оплаты",
    2: "В работе",
    3: "Передан курьеру",
    4: "Ожидает доставки",
    5: "Доставлен",
    "-10": "Ошибка системы",
    "-1": "Отменен",
    "-2": "Отменен",
  },
};

const Status = ({ type }: { type: OrderStatus }) => (
  <span
    className={`${status.color[type]} inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium capitalize`}
  >
    {status.text[type]}
  </span>
);
const OrderList = () => {
  const [parent] = useAutoAnimate();
  const [userData, setUserData] = useState<null | UserData>(null);

  useEffect(() => {
    const user = localStorage.getItem("userData");
    user && setUserData(JSON.parse(user));
  }, []);
  const { data, isLoading } = useQuery(
    ["orders"],
    () => getOrder(userData?.id),
    {
      enabled: !!userData,
    }
  );

  const [activeOrder, setActiveOrder] = useState<null | number>(null);

  const orders = data?.data?.response?.data;

  return (
    <>
      <div className="shadow sm:hidden">
        <ul
          role="list"
          ref={parent}
          className="mt-2 divide-y divide-gray-200 overflow-hidden shadow sm:hidden"
        >
          {isLoading ? (
            <div className="absolute flex h-full w-full items-center justify-center text-blue-600">
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
                  strokeWidth="10"
                  r="35"
                  strokeDasharray="164.93361431346415 56.97787143782138"
                />
              </svg>
            </div>
          ) : (
            orders.map((order) => (
              <li key={order.id}>
                <button
                  onClick={() => setActiveOrder(order.id)}
                  className="block w-full bg-white px-4 py-4 hover:bg-gray-50"
                >
                  <span className="flex w-full items-center space-x-4">
                    <span className="flex flex-1 space-x-2 truncate">
                      <BanknotesIcon
                        className="h-5 w-5 flex-shrink-0 text-gray-400"
                        aria-hidden="true"
                      />
                      <span className="flex w-full justify-between truncate text-sm text-gray-500">
                        <span className="truncate">
                          {order?.products?.name}
                        </span>
                        <span className="flex items-center">
                          <span className="font-medium text-gray-900">
                            {order?.products?.price} <Tenge />
                          </span>
                        </span>
                        <time dateTime={order.created_at.Time}>
                          {formatTime(order.created_at.Time)}
                        </time>
                      </span>
                    </span>
                    <ChevronRightIcon
                      className="h-5 w-5 flex-shrink-0 text-gray-400"
                      aria-hidden="true"
                    />
                  </span>
                </button>
              </li>
            ))
          )}
        </ul>
      </div>

      {/* Activity table (small breakpoint and up) */}
      <div className="hidden sm:block">
        <div className="mx-auto max-w-6xl px-4 sm:px-6 lg:px-8">
          <div className="mt-2 flex flex-col">
            <div className="min-w-full overflow-hidden overflow-x-auto align-middle shadow sm:rounded-lg">
              <table className="min-w-full divide-y divide-gray-200">
                <thead>
                  <tr>
                    <th
                      className="bg-gray-50 px-6 py-3 text-left text-sm font-semibold text-gray-900"
                      scope="col"
                    >
                      Заказ
                    </th>
                    <th
                      className="bg-gray-50 px-6 py-3 text-right text-sm font-semibold text-gray-900"
                      scope="col"
                    >
                      Цена
                    </th>
                    <th
                      className="hidden bg-gray-50 px-6 py-3 text-left text-sm font-semibold text-gray-900 md:block"
                      scope="col"
                    >
                      Статус
                    </th>
                    <th
                      className="bg-gray-50 px-6 py-3 text-right text-sm font-semibold text-gray-900"
                      scope="col"
                    >
                      Дата
                    </th>
                  </tr>
                </thead>
                <tbody
                  className="divide-y divide-gray-200 bg-white"
                  ref={parent}
                >
                  {isLoading ? (
                    <tr>
                      <td align="center" className="py-10 text-blue-600">
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
                            strokeWidth="10"
                            r="35"
                            strokeDasharray="164.93361431346415 56.97787143782138"
                          />
                        </svg>
                      </td>
                    </tr>
                  ) : (
                    orders.map((order) => (
                      <tr
                        key={order.id}
                        className="bg-white"
                        onClick={() => setActiveOrder(order.id)}
                      >
                        <td className="w-full max-w-0 whitespace-nowrap px-6 py-4 text-sm text-gray-900">
                          <div className="flex">
                            <button className="group inline-flex space-x-2 truncate text-sm">
                              <BanknotesIcon
                                className="h-5 w-5 flex-shrink-0 text-gray-400 group-hover:text-gray-500"
                                aria-hidden="true"
                              />
                              <p className="truncate text-gray-500 group-hover:text-gray-900">
                                {order?.products?.name}
                              </p>
                            </button>
                          </div>
                        </td>
                        <td className="whitespace-nowrap px-6 py-4 text-right text-sm text-gray-500">
                          <p className="flex items-center font-medium text-gray-900">
                            {order?.products?.price} <Tenge />
                          </p>
                        </td>
                        <td className="hidden whitespace-nowrap px-6 py-4 text-sm text-gray-500 md:block">
                          <Status type={order.status} />
                        </td>
                        <td className="whitespace-nowrap px-6 py-4 text-right text-sm text-gray-500">
                          <time dateTime={order.created_at.Time}>
                            {formatTime(order.created_at.Time)}
                          </time>
                        </td>
                      </tr>
                    ))
                  )}
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
      <OrderItem
        open={!!activeOrder}
        setOpen={() => setActiveOrder(null)}
        id={activeOrder}
      />
    </>
  );
};

const getOrderItem = (id: number): Promise<ApiResponse<Order>> =>
  instance.get(`/api/orders/${id}`);

const OrderItem = ({
  open,
  setOpen,
  id,
}: {
  id: number;
  open: boolean;
  setOpen: () => void;
}) => {
  const { data: orderData, isLoading: orderIsloading } = useQuery<
    ApiResponse<Order>
  >(["order", id], () => getOrderItem(id), {
    enabled: !!id,
  });

  const timeLine = useMemo(() => {
    if (!orderData) return [];

    const data = orderData.data.response.data;

    return [
      {
        id: 1,
        type: eventTypes.created,
        content: "Заказ ",
        target: "Создан",
        active: !data?.cancel_reason && !!data?.created_at?.Valid,
        date: data?.created_at?.Valid
          ? formatTime(data?.created_at?.Time)
          : null,
        datetime: data?.created_at?.Valid
          ? formatTime(data?.created_at?.Time)
          : null,
      },
      {
        id: 2,
        type: eventTypes.payed,
        content: "Заказ",
        target: "Оплачен",
        active: !data?.cancel_reason && !!data?.payed_at?.Valid,
        date: data?.payed_at?.Valid ? formatTime(data?.payed_at?.Time) : null,
        datetime: data?.payed_at?.Valid
          ? formatTime(data?.payed_at?.Time)
          : null,
      },
      {
        id: 3,
        type: eventTypes.packed,
        content: "Взят в обработку",
        target: "Исполнителем",
        active: !data?.cancel_reason && !!data?.packed_at?.Valid,
        date: data?.packed_at?.Valid ? formatTime(data?.packed_at?.Time) : null,
        datetime: data?.packed_at?.Valid
          ? formatTime(data?.packed_at?.Time)
          : null,
      },
      {
        id: 4,
        type: eventTypes.created,
        content: "Доставлен",
        target: "на указанный адрес",
        active: !data?.cancel_reason && !!data?.delivered_at?.Valid,
        date: data?.delivered_at?.Valid
          ? formatTime(data?.delivered_at?.Time)
          : null,
        datetime: data?.delivered_at?.Valid
          ? formatTime(data?.delivered_at?.Time)
          : null,
      },
      {
        id: 5,
        type: eventTypes.canceled,
        content: "Заказ",
        target: "Отменен",
        active: !!data?.cancel_reason,
        date: "",
        datetime: "",
      },
    ];
  }, [orderData]);

  const [userData, setUserData] = useState<null | UserData>(null);

  useEffect(() => {
    const user = localStorage.getItem("userData");

    user && setUserData(JSON.parse(user));
  }, []);

  const mutation = useMutation({
    mutationFn: () =>
      instance.post("/api/orders/" + id, { status: -1, cancel_reason: "a" }),
    onSuccess: () => {
      queryClient.refetchQueries(["orders"]);
      setOpen();
    },
  });

  const calcelOrder = () => {
    return mutation.mutate();
  };

  return (
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
              <Dialog.Panel className="flex w-full transform text-left text-base transition md:my-8 md:max-w-2xl md:px-4 lg:max-w-7xl">
                <div className="relative flex w-full items-center overflow-hidden bg-white px-4 pt-14 pb-8 shadow-2xl sm:px-6 sm:pt-8 md:p-6 lg:p-8">
                  <button
                    type="button"
                    className="absolute top-4 right-4 text-gray-400 hover:text-gray-500 sm:top-8 sm:right-6 md:top-6 md:right-6 lg:top-8 lg:right-8"
                    onClick={() => setOpen()}
                  >
                    <span className="sr-only">Close</span>
                    <XMarkIcon className="h-6 w-6" aria-hidden="true" />
                  </button>

                  {orderIsloading ? (
                    <div>loading</div>
                  ) : (
                    <main className="py-10">
                      <div className="mx-auto mt-8 grid max-w-3xl grid-cols-1 gap-6 sm:px-6 lg:max-w-7xl lg:grid-flow-col-dense lg:grid-cols-3">
                        <div className="space-y-6 lg:col-span-2 lg:col-start-1">
                          {/* Description list*/}
                          <section aria-labelledby="applicant-information-title">
                            <div className="bg-white shadow sm:rounded-lg">
                              <div className="px-4 py-5 sm:px-6">
                                <h2
                                  id="applicant-information-title"
                                  className="text-lg font-medium leading-6 text-gray-900"
                                >
                                  Информация о заказе
                                </h2>
                              </div>
                              <div className="border-t border-gray-200 px-4 py-5 sm:px-6">
                                <dl className="grid grid-cols-1 gap-x-4 gap-y-8 sm:grid-cols-2">
                                  <div className="sm:col-span-1">
                                    <dt className="text-sm font-medium text-gray-500">
                                      Название товара
                                    </dt>
                                    <dd className="mt-1 text-sm text-gray-900">
                                      {
                                        orderData?.data?.response?.data
                                          ?.products?.name
                                      }
                                    </dd>
                                  </div>
                                  <div className="sm:col-span-1">
                                    <dt className="text-sm font-medium text-gray-500">
                                      Email
                                    </dt>
                                    <dd className="mt-1 text-sm text-gray-900">
                                      {userData?.email}
                                    </dd>
                                  </div>
                                  <div className="sm:col-span-1">
                                    <dt className="text-sm font-medium text-gray-500">
                                      Цена
                                    </dt>
                                    <dd className="mt-1 flex items-center text-sm text-gray-900">
                                      {
                                        orderData?.data?.response?.data
                                          ?.products?.price
                                      }{" "}
                                      <Tenge />
                                    </dd>
                                  </div>
                                  <div className="sm:col-span-1">
                                    <dt className="text-sm font-medium text-gray-500">
                                      Вес
                                    </dt>
                                    <dd className="mt-1 text-sm text-gray-900">
                                      {
                                        orderData?.data?.response?.data
                                          ?.products?.weight
                                      }
                                    </dd>
                                  </div>
                                  <div className="sm:col-span-2">
                                    <dt className="text-sm font-medium text-gray-500">
                                      Описание
                                    </dt>
                                    <dd className="mt-1 text-sm text-gray-900">
                                      {
                                        orderData?.data?.response?.data
                                          ?.products?.description
                                      }
                                    </dd>
                                  </div>
                                </dl>
                              </div>
                              <div>
                                <a className="block bg-gray-50 px-4 py-4 text-center text-sm font-medium text-gray-500 hover:text-gray-700 sm:rounded-b-lg">
                                  Подробнее
                                </a>
                              </div>
                            </div>
                          </section>
                        </div>

                        <section
                          aria-labelledby="timeline-title"
                          className="lg:col-span-1 lg:col-start-3"
                        >
                          <div className="bg-white px-4 py-5 shadow sm:rounded-lg sm:px-6">
                            <h2
                              id="timeline-title"
                              className="text-lg font-medium text-gray-900"
                            >
                              Временная линия
                            </h2>

                            <div className="mt-6 flow-root">
                              <ul role="list" className="-mb-8">
                                {timeLine.map((item, itemIdx) => (
                                  <li key={item.id}>
                                    <div className="relative pb-8">
                                      {itemIdx !== timeLine.length - 1 ? (
                                        <span
                                          className="absolute top-4 left-4 -ml-px h-full w-0.5 bg-gray-200"
                                          aria-hidden="true"
                                        />
                                      ) : null}
                                      <div className="relative flex space-x-3">
                                        <div>
                                          <span
                                            className={classNames(
                                              item?.active
                                                ? item.type.bgColorClass
                                                : "bg-gray-400",
                                              "flex h-8 w-8 items-center justify-center rounded-full ring-8 ring-white"
                                            )}
                                          >
                                            <item.type.icon
                                              className="h-5 w-5 text-white"
                                              aria-hidden="true"
                                            />
                                          </span>
                                        </div>
                                        <div className="flex min-w-0 flex-1 justify-between space-x-4 pt-1.5">
                                          <div>
                                            <p className="text-sm text-gray-500">
                                              {item.content}{" "}
                                              <a
                                                href="#"
                                                className="font-medium text-gray-900"
                                              >
                                                {item.target}
                                              </a>
                                            </p>
                                          </div>
                                          <div className="whitespace-nowrap text-right text-sm text-gray-500">
                                            <time dateTime={item.datetime}>
                                              {item.date}
                                            </time>
                                          </div>
                                        </div>
                                      </div>
                                    </div>
                                  </li>
                                ))}
                              </ul>
                            </div>
                            <div className="justify-stretch mt-6 flex flex-col">
                              <button
                                type="button"
                                onClick={calcelOrder}
                                className="inline-flex items-center justify-center rounded-md border border-transparent bg-red-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
                              >
                                Отменить заказ
                              </button>
                            </div>
                          </div>
                        </section>
                      </div>
                    </main>
                  )}
                </div>
              </Dialog.Panel>
            </Transition.Child>
          </div>
        </div>
      </Dialog>
    </Transition.Root>
  );
};

const Head = () => {
  const [userData, setUserData] = useState<null | UserData>(null);

  useEffect(() => {
    const user = localStorage.getItem("userData");

    user && setUserData(JSON.parse(user));
  }, []);

  return (
    <div className="bg-white shadow">
      <div className="px-4 sm:px-6 lg:mx-auto lg:max-w-7xl lg:px-8">
        <div className="py-6 md:flex md:items-center md:justify-between lg:border-t lg:border-gray-200">
          <div className="min-w-0 flex-1">
            {/* Profile */}
            <div className="flex items-center">
              <img
                className="hidden h-16 w-16 rounded-full sm:block"
                src="https://images.unsplash.com/photo-1517365830460-955ce3ccd263?ixlib=rb-=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=8&w=256&h=256&q=80"
                alt=""
              />
              <div>
                <div className="flex items-center">
                  <img
                    className="h-16 w-16 rounded-full sm:hidden"
                    src="https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2.6&w=256&h=256&q=80"
                    alt=""
                  />
                  <h1 className="ml-3 text-2xl font-bold leading-7 text-gray-900 sm:truncate sm:leading-9">
                    Привет, {userData?.email}
                  </h1>
                </div>
                <dl className="mt-6 flex flex-col sm:ml-3 sm:mt-1 sm:flex-row sm:flex-wrap">
                  <dd className="mt-3 flex items-center text-sm font-medium capitalize text-gray-500 sm:mr-6 sm:mt-0">
                    <CheckCircleIcon
                      className="mr-1.5 h-5 w-5 flex-shrink-0 text-green-400"
                      aria-hidden="true"
                    />
                    Подтвержденный аккаунт
                  </dd>
                </dl>
              </div>
            </div>
          </div>
          <div className="mt-6 flex space-x-3 md:mt-0 md:ml-4">
            <button
              type="button"
              className="inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-cyan-500 focus:ring-offset-2"
            >
              Пополнить счет
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

const Info = () => (
  <div className="mx-auto max-w-6xl px-4 sm:px-6 lg:px-8">
    <div className="mt-2 grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-3">
      <div className="overflow-hidden rounded-lg bg-white shadow">
        <div className="p-5">
          <div className="flex items-center">
            <div className="flex-shrink-0">
              <ScaleIcon className="h-6 w-6 text-gray-400" aria-hidden="true" />
            </div>
            <div className="ml-5 w-0 flex-1">
              <dl>
                <dt className="truncate text-sm font-medium text-gray-500">
                  Баланс
                </dt>
                <dd>
                  <div className="flex items-center text-lg font-medium text-gray-900">
                    123 344{" "}
                    <img
                      className="object-fit h-2.5 w-2"
                      src="https://upload.wikimedia.org/wikipedia/commons/f/f8/Tenge_symbol.svg"
                      alt=""
                    />
                  </div>
                </dd>
              </dl>
            </div>
          </div>
        </div>
        <div className="bg-gray-50 px-5 py-3">
          <div className="text-sm">
            <a className="font-medium text-cyan-700 hover:text-cyan-900"> </a>
          </div>
        </div>
      </div>
    </div>
  </div>
);

const formatTime = (time: string) =>
  new Intl.DateTimeFormat("ru", {
    year: "numeric",
    month: "numeric",
    day: "2-digit",
  }).format(new Date(time));

export const Tenge = () => (
  <img
    className="object-fit h-2.5 w-2"
    src="https://upload.wikimedia.org/wikipedia/commons/f/f8/Tenge_symbol.svg"
    alt=""
  />
);
export default Cabinet;
