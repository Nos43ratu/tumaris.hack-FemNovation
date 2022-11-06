import React from "react";

const Home = () => {
  return (
    <>
      <Header />
      <Hero />
    </>
  );
};

const Hero = () => {
  return (
    <div className="relative mx-auto max-w-7xl px-4 pt-20 pb-16 text-center sm:px-6 lg:px-8 lg:pt-32">
      <h1 className="font-display mx-auto max-w-4xl text-5xl font-medium tracking-tight text-slate-900 sm:text-7xl">
        Помогаем найти{" "}
        <span className="relative whitespace-nowrap text-blue-600">
          <svg
            aria-hidden="true"
            viewBox="0 0 418 42"
            className="absolute top-2/3 left-0 h-[0.58em] w-full fill-blue-300/70"
            preserveAspectRatio="none"
          >
            <path d="M203.371.916c-26.013-2.078-76.686 1.963-124.73 9.946L67.3 12.749C35.421 18.062 18.2 21.766 6.004 25.934 1.244 27.561.828 27.778.874 28.61c.07 1.214.828 1.121 9.595-1.176 9.072-2.377 17.15-3.92 39.246-7.496C123.565 7.986 157.869 4.492 195.942 5.046c7.461.108 19.25 1.696 19.17 2.582-.107 1.183-7.874 4.31-25.75 10.366-21.992 7.45-35.43 12.534-36.701 13.884-2.173 2.308-.202 4.407 4.442 4.734 2.654.187 3.263.157 15.593-.78 35.401-2.686 57.944-3.488 88.365-3.143 46.327.526 75.721 2.23 130.788 7.584 19.787 1.924 20.814 1.98 24.557 1.332l.066-.011c1.201-.203 1.53-1.825.399-2.335-2.911-1.31-4.893-1.604-22.048-3.261-57.509-5.556-87.871-7.36-132.059-7.842-23.239-.254-33.617-.116-50.627.674-11.629.54-42.371 2.494-46.696 2.967-2.359.259 8.133-3.625 26.504-9.81 23.239-7.825 27.934-10.149 28.304-14.005.417-4.348-3.529-6-16.878-7.066Z"></path>
          </svg>
          <span className="relative">друг друга</span>
        </span>{" "}
      </h1>
      <p className="mx-auto mt-6 max-w-2xl text-lg tracking-tight text-slate-700">
        Kustoma - сервис, который помогает объединять людей, которые ищут и
        создают handmade товар.
      </p>
      <div className="mt-10 flex justify-center gap-x-6">
        <a
          className="group inline-flex items-center justify-center rounded-full bg-slate-900 py-2 px-4 text-sm font-semibold text-white hover:bg-slate-700 hover:text-slate-100 focus:outline-none focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-slate-900 active:bg-slate-800 active:text-slate-300"
          href="/sign-in"
        >
          Найти услугу
        </a>
        <a
          className="group inline-flex items-center justify-center rounded-full py-2 px-4 text-sm text-slate-700 ring-1 ring-slate-200 hover:text-slate-900 hover:ring-slate-300 focus:outline-none focus-visible:outline-blue-600 focus-visible:ring-slate-300 active:bg-slate-100 active:text-slate-600"
          href="/sign-in"
        >
          <span className="">Стать исполнителем</span>
        </a>
      </div>
    </div>
  );
};

const Header = () => {
  return (
    <header className="py-10">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <nav className="relative z-50 flex justify-between">
          <div className="flex items-center md:gap-x-12">
            <a aria-label="Home" href="/#">
              <svg
                width="200"
                height="26"
                viewBox="0 0 556 74"
                fill="#2563eb"
                className="h-[20px] w-[150px] md:h-[26px] md:w-[200px]"
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
            <div className="hidden md:flex md:gap-x-6">
              <a
                className="inline-block rounded-lg py-1 px-2 text-sm text-slate-700 hover:bg-slate-100 hover:text-slate-900"
                href="/#features"
              >
                Как это работает?
              </a>
              <a
                className="inline-block rounded-lg py-1 px-2 text-sm text-slate-700 hover:bg-slate-100 hover:text-slate-900"
                href="/#testimonials"
              >
                Возможности
              </a>
              <a
                className="inline-block rounded-lg py-1 px-2 text-sm text-slate-700 hover:bg-slate-100 hover:text-slate-900"
                href="/#pricing"
              >
                О нас
              </a>
            </div>
          </div>
          <div className="flex items-center gap-x-5 md:gap-x-8">
            <div className="hidden md:block">
              <a
                className="inline-block rounded-lg py-1 px-2 text-sm text-slate-700 hover:bg-slate-100 hover:text-slate-900"
                href="/sign-in"
              >
                Войти
              </a>
            </div>
            <a
              className="group inline-flex items-center justify-center rounded-full bg-blue-600 py-2 px-4 text-sm font-semibold text-white hover:bg-blue-500 hover:text-slate-100 focus:outline-none focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-blue-600 active:bg-blue-800 active:text-blue-100"
              href="/sign-up"
            >
              <span>
                Начни <span className="hidden lg:inline">сейчас</span>
              </span>
            </a>
          </div>
        </nav>
      </div>
    </header>
  );
};

export default Home;
