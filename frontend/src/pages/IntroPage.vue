<template>
    <!--container divide space -->
    <div class="container w-20 space-y-0 divide-y-8 divide-yellow-400">
        <div class="bg-slate-600 size-24"></div>
        <div class="bg-slate-600 size-24"></div>
        <div class="bg-slate-600 size-24"></div>
        <div class="bg-slate-600 size-24"></div>
        <div class="bg-slate-600 size-24"></div>
        <div class="bg-slate-600 size-24"></div>
    </div>

    <!-- 文字隔断  line-clamp限制行数 truncate限制一行-->
    <div class="w-48 truncate line-clamp-2">
        Lorem ipsum dolor sit amet consectetur, adipisicing elit. Iure debitis illum soluta
        quibusdam quam officia dolores in, suscipit, sapiente deserunt provident pariatur ipsum quis
        nobis numquam praesentium. Amet, velit minus?
    </div>
    <!-- 过渡渐变 -->
    <div
        class="w-full h-48 bg-gradient-to-r from-20% from-orange-500 to-slate-400 via-sky-400"></div>

    <!-- ring 使用box-shadow的伪边框 -->
    <button class="w-24 h-12 bg-sky-600 ring-4 ring-red-600">Tex</button>

    <AsyncComp></AsyncComp>

    <nav class="bg-white border-gray-200 dark:bg-gray-900">
        <div class="flex flex-wrap items-center justify-between max-w-screen-xl p-4 mx-auto">
            <a href="https://flowbite.com" class="flex items-center space-x-3 rtl:space-x-reverse">
                <img
                    src="https://flowbite.com/docs/images/logo.svg"
                    class="h-8"
                    alt="Flowbite Logo" />
                <span class="self-center text-2xl font-semibold whitespace-nowrap dark:text-white"
                    >Flowbite</span
                >
            </a>
            <div class="flex items-center space-x-1 md:order-2 md:space-x-2 rtl:space-x-reverse">
                <a
                    href="#"
                    class="text-gray-800 dark:text-white hover:bg-gray-50 focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm px-4 py-2 md:px-5 md:py-2.5 dark:hover:bg-gray-700 focus:outline-none dark:focus:ring-gray-800"
                    >Login</a
                >
                <a
                    href="#"
                    class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 md:px-5 md:py-2.5 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800"
                    >Sign up</a
                >
                <button
                    data-collapse-toggle="mega-menu"
                    type="button"
                    class="inline-flex items-center justify-center w-10 h-10 p-2 text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
                    aria-controls="mega-menu"
                    aria-expanded="false">
                    <span class="sr-only">Open main menu</span>
                    <svg
                        class="w-5 h-5"
                        aria-hidden="true"
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 17 14">
                        <path
                            stroke="currentColor"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M1 1h15M1 7h15M1 13h15" />
                    </svg>
                </button>
            </div>
            <div
                id="mega-menu"
                class="items-center justify-between hidden w-full md:flex md:w-auto md:order-1">
                <ul
                    class="flex flex-col mt-4 font-medium md:flex-row md:mt-0 md:space-x-8 rtl:space-x-reverse">
                    <li>
                        <a
                            href="#"
                            class="block px-3 py-2 text-blue-600 border-b border-gray-100 hover:bg-gray-50 md:hover:bg-transparent md:border-0 md:hover:text-blue-600 md:p-0 dark:text-blue-500 md:dark:hover:text-blue-500 dark:hover:bg-gray-700 dark:hover:text-blue-500 md:dark:hover:bg-transparent dark:border-gray-700"
                            aria-current="page"
                            >Home</a
                        >
                    </li>
                    <li>
                        <button
                            id="mega-menu-dropdown-button"
                            data-dropdown-toggle="mega-menu-dropdown"
                            class="flex items-center justify-between w-full px-3 py-2 font-medium text-gray-900 border-b border-gray-100 md:w-auto hover:bg-gray-50 md:hover:bg-transparent md:border-0 md:hover:text-blue-600 md:p-0 dark:text-white md:dark:hover:text-blue-500 dark:hover:bg-gray-700 dark:hover:text-blue-500 md:dark:hover:bg-transparent dark:border-gray-700">
                            Company
                            <svg
                                class="w-2.5 h-2.5 ms-3"
                                aria-hidden="true"
                                xmlns="http://www.w3.org/2000/svg"
                                fill="none"
                                viewBox="0 0 10 6">
                                <path
                                    stroke="currentColor"
                                    stroke-linecap="round"
                                    stroke-linejoin="round"
                                    stroke-width="2"
                                    d="m1 1 4 4 4-4" />
                            </svg>
                        </button>
                        <div
                            id="mega-menu-dropdown"
                            class="absolute z-10 grid hidden w-auto grid-cols-2 text-sm bg-white border border-gray-100 rounded-lg shadow-md dark:border-gray-700 md:grid-cols-3 dark:bg-gray-700">
                            <div class="p-4 pb-0 text-gray-900 md:pb-4 dark:text-white">
                                <ul class="space-y-4" aria-labelledby="mega-menu-dropdown-button">
                                    <li>
                                        <a
                                            href="#"
                                            class="text-gray-500 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-500">
                                            About Us
                                        </a>
                                    </li>
                                    <li>
                                        <a
                                            href="#"
                                            class="text-gray-500 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-500">
                                            Library
                                        </a>
                                    </li>
                                    <li>
                                        <a
                                            href="#"
                                            class="text-gray-500 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-500">
                                            Resources
                                        </a>
                                    </li>
                                    <li>
                                        <a
                                            href="#"
                                            class="text-gray-500 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-500">
                                            Pro Version
                                        </a>
                                    </li>
                                </ul>
                            </div>
                            <div class="p-4 pb-0 text-gray-900 md:pb-4 dark:text-white">
                                <ul class="space-y-4">
                                    <li>
                                        <a
                                            href="#"
                                            class="text-gray-500 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-500">
                                            Blog
                                        </a>
                                    </li>
                                    <li>
                                        <a
                                            href="#"
                                            class="text-gray-500 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-500">
                                            Newsletter
                                        </a>
                                    </li>
                                    <li>
                                        <a
                                            href="#"
                                            class="text-gray-500 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-500">
                                            Playground
                                        </a>
                                    </li>
                                    <li>
                                        <a
                                            href="#"
                                            class="text-gray-500 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-500">
                                            License
                                        </a>
                                    </li>
                                </ul>
                            </div>
                            <div class="p-4">
                                <ul class="space-y-4">
                                    <li>
                                        <a
                                            href="#"
                                            class="text-gray-500 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-500">
                                            Contact Us
                                        </a>
                                    </li>
                                    <li>
                                        <a
                                            href="#"
                                            class="text-gray-500 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-500">
                                            Support Center
                                        </a>
                                    </li>
                                    <li>
                                        <a
                                            href="#"
                                            class="text-gray-500 dark:text-gray-400 hover:text-blue-600 dark:hover:text-blue-500">
                                            Terms
                                        </a>
                                    </li>
                                </ul>
                            </div>
                        </div>
                    </li>
                    <li>
                        <a
                            href="#"
                            class="block px-3 py-2 text-gray-900 border-b border-gray-100 hover:bg-gray-50 md:hover:bg-transparent md:border-0 md:hover:text-blue-600 md:p-0 dark:text-white md:dark:hover:text-blue-500 dark:hover:bg-gray-700 dark:hover:text-blue-500 md:dark:hover:bg-transparent dark:border-gray-700"
                            >Team</a
                        >
                    </li>
                    <li>
                        <a
                            href="#"
                            class="block px-3 py-2 text-gray-900 border-b border-gray-100 hover:bg-gray-50 md:hover:bg-transparent md:border-0 md:hover:text-blue-600 md:p-0 dark:text-white md:dark:hover:text-blue-500 dark:hover:bg-gray-700 dark:hover:text-blue-500 md:dark:hover:bg-transparent dark:border-gray-700"
                            >Contact</a
                        >
                    </li>
                </ul>
            </div>
        </div>
    </nav>
</template>

<script setup lang="ts">
const AsyncComp = defineAsyncComponent({
    // the loader function
    loader: () => new Promise((res, rej) => h('div', ['<span>Loading</span>'])),

    // A component to use while the async component is loading
    loadingComponent: h('div', ['<span>Loading</span>']),
    // Delay before showing the loading component. Default: 200ms.
    delay: 200,

    // A component to use if the load fails
    errorComponent: h('div', ['<span>Loading</span>']),
    // The error component will be displayed if a timeout is
    // provided and exceeded. Default: Infinity.
    timeout: 3000
})
</script>
