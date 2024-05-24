<template>
    <div class="fixed h-full w-full left-0 py-64 px-8 z-50">
        <div class="rounded-lg w-[576px] mx-auto top-36 bg-white-base dark:bg-dark-light">
            <div class="flex flex-col justify-center py-12 relative">
                <div class="absolute right-4 top-4">
                    <i class="pi pi-times" @click="showLoginWindow = false" :size="24"></i>
                </div>

                <div class="sm:mx-auto sm:w-full sm:max-w-sm">
                    <img
                        class="mx-auto h-10 w-auto"
                        src="https://cdn.yuelili.com/web/assets/logo.webp"
                        alt="Your Company" />
                    <h2 class="mt-10 text-center text-2xl font-bold leading-9 tracking-tight">
                        请登录管理员账号
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
                                    class="px-4 block w-full rounded-md border-0 py-1.5 shadow-sm ring-1 ring-inset ring-gray-300 dark:bg-slate-800 placeholder:text-gray-400 focus:ring-2 focus:ring-inset !focus:ring-indigo-600 sm:text-sm sm:leading-6" />
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
                                    class="dark:bg-slate-800 px-4 block w-full rounded-md border-0 py-1.5 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-200 sm:text-sm sm:leading-6" />
                            </div>
                        </div>

                        <div>
                            <button
                                type="submit"
                                @click.prevent="login"
                                class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">
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
import { ref } from 'vue'
import { fetchAuthLogin } from '@/handlers'
import { storeToRefs } from 'pinia'

import { Message } from '@/plugins/message'

import { basicStore } from '@/stores/index'
const basic = basicStore()
let { isAdmin, token } = storeToRefs(basic)

const showLoginWindow = defineModel('showLoginWindow')

const username = ref('')
const password = ref('')

async function login() {
    if (username.value && password.value) {
        const [ok, data] = await fetchAuthLogin(username.value, password.value)
        if (ok) {
            localStorage.setItem('token', data)
            showLoginWindow.value = false
            isAdmin.value = true
            token.value = data
            Message('登录成功', 'success')
        } else {
            Message('登录失败', 'error')
        }
    }
}
</script>
