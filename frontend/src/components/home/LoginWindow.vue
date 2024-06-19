<template>
    <div ref="refDom" class="fixed left-0 z-50 w-full h-full px-8 py-64">
        <div
            class="rounded-lg w-[576px] mx-auto shadow-2xl top-36 bg-theme-card dark:bg-dark-light">
            <div class="relative flex flex-col justify-center py-12">
                <div class="absolute right-4 top-4 text-[1.25rem]">
                    <BIconX @click="showLoginWindow = false"></BIconX>
                </div>

                <div class="sm:mx-auto sm:w-full sm:max-w-sm">
                    <img
                        class="w-auto h-10 mx-auto"
                        src="https://cdn.yuelili.com/web/assets/logo.webp"
                        alt="Your Company" />
                    <h2 class="mt-10 text-2xl font-bold leading-9 tracking-tight text-center">
                        请登录账号
                    </h2>
                </div>

                <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
                    <form class="space-y-6" action="#" method="POST">
                        <div>
                            <label for="username" class="block text-sm font-medium leading-6"
                                >用户名</label
                            >
                            <div class="mt-2">
                                <input
                                    id="username"
                                    v-model="username"
                                    type="username"
                                    autocomplete="username"
                                    required
                                    class="px-4 block w-full rounded-md border-0 py-1.5 shadow-sm ring-1 ring-inset ring-gray-300 bg-theme-base placeholder:text-gray-400 focus:ring-2 focus:ring-inset !focus:ring-indigo-600 sm:text-sm sm:leading-6" />
                            </div>
                        </div>

                        <div>
                            <div class="flex items-center justify-between">
                                <label for="password" class="block text-sm font-medium leading-6"
                                    >密码</label
                                >
                            </div>
                            <div class="mt-2">
                                <input
                                    id="password"
                                    v-model="password"
                                    type="password"
                                    autocomplete="current-password"
                                    required
                                    class="bg-theme-base px-4 block w-full rounded-md border-0 py-1.5 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-200 sm:text-sm sm:leading-6" />
                            </div>
                        </div>

                        <div>
                            <button
                                type="submit"
                                @click.prevent="login"
                                class="flex w-full justify-center rounded-md bg-theme-primary-base px-3 py-1.5 text-sm font-semibold leading-6 shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">
                                登录
                            </button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { BIconX } from 'bootstrap-icons-vue'

const refDom = ref<HTMLElement>()

const basic = basicStore()
let { isAdmin } = storeToRefs(basic)

const showLoginWindow = defineModel('showLoginWindow')

const username = ref('')
const password = ref('')

async function login() {
    if (username.value && password.value) {
        const [ok, data] = await fetchAuthLogin(username.value, password.value)
        if (ok) {
            localStorage.setItem('token', data['data'])
            showLoginWindow.value = false
            isAdmin.value = true
            await Message({ message: '登录成功', type: 'success' })
        } else {
            await Message({ message: '登录失败', type: 'error' })
        }
    }
}
</script>
