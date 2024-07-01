<template>
    <div class="relative flex flex-col justify-center">
        <div class="">
            <img
                class="w-auto h-10 mx-auto"
                src="https://cdn.yuelili.com/docs/web/assets/ydocs-256.png"
                alt="Your Company" />
            <h2 class="mt-4 text-lg font-bold leading-9 tracking-tight text-center">登录</h2>
        </div>

        <div class="mt-4">
            <form class="space-y-6" action="#" method="POST">
                <div>
                    <label for="username" class="block text-sm font-medium leading-6">用户名</label>
                    <div class="mt-2">
                        <input
                            id="username"
                            v-model="username"
                            type="username"
                            autocomplete="username"
                            required
                            class="px-4 block w-full rounded-md border-0 py-1.5 shadow-sm ring-1 ring-inset ring-gray-300 bg-theme-base placeholder:text-gray-400 focus:ring-2 focus:ring-inset !focus:ring-indigo-600" />
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
                            class="bg-theme-base px-4 block w-full rounded-md border-0 py-1.5 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-200" />
                    </div>
                </div>

                <div>
                    <button
                        type="submit"
                        @click.prevent="login"
                        class="flex w-full justify-center rounded-md bg-theme-primary-base px-3 py-1.5 text-theme-text-inverse text-sm font-semibold leading-6 shadow-sm hover:bg-blue-600 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">
                        登录
                    </button>
                </div>
            </form>
        </div>
    </div>
</template>

<script setup lang="ts">
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
