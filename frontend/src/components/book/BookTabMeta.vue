<template>
    <div class="flex pb-4 toolbar">
        <button type="button" class="px-3 py-1 ml-auto btn hover:" @click="updateMeta">更新</button>
        <button type="button" class="px-3 py-1 ml-3 btn hover:" @click="saveMeta">保存</button>
    </div>
    <div>
        <table class="w-full border-collapse table-auto">
            <thead>
                <tr class="*:px-4 *:py-2 *:border">
                    <th class="border">ID</th>
                    <th>{{ translate('title') }}</th>
                    <th>{{ translate('order') }}</th>
                    <th>{{ translate('status') }}</th>
                </tr>
            </thead>
            <tbody>
                <tr class="*:px-4 *:py-2" v-for="meta in metas.categorys" :key="meta.url">
                    <td>
                        <input type="text" class="bg-transparent" disabled v-model="meta.name" />
                    </td>
                    <td>
                        <input
                            type="text"
                            class="w-full p-2 bg-transparent border rounded-sm"
                            v-model="meta.title" />
                    </td>
                    <td>
                        <input
                            type="text"
                            class="w-full p-2 bg-transparent border rounded-sm"
                            v-model.number="meta.order" />
                    </td>
                    <td class="text-center">
                        <input
                            type="checkbox"
                            id="checkbox"
                            v-model="meta.status"
                            :true-value="false"
                            :false-value="true" />
                    </td>
                </tr>
            </tbody>

            <tbody>
                <tr class="*:px-4 *:py-2" v-for="meta in metas.documents" :key="meta.url">
                    <td>
                        <input type="text" class="bg-transparent" disabled v-model="meta.name" />
                    </td>
                    <td>
                        <input
                            type="text"
                            class="w-full p-2 bg-transparent border rounded-sm"
                            v-model="meta.title" />
                    </td>
                    <td>
                        <input
                            type="text"
                            class="w-full p-2 bg-transparent border rounded-sm"
                            v-model.number="meta.order" />
                    </td>
                    <td class="text-center">
                        <input
                            type="checkbox"
                            id="checkbox"
                            v-model="meta.status"
                            :true-value="false"
                            :false-value="true" />
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script setup lang="ts">
import { LocalMetaDatas } from '@/models/base'

const basic = basicStore()
const route = useRoute()
const locale = computed(() => basic.locale)
const translate = basic.translate

const metas = ref<LocalMetaDatas>(new LocalMetaDatas())

async function saveMeta() {
    const result = await saveBookMeta(
        (route.params['bookPath'] as string[]).join('/'),
        locale.value,
        metas.value
    )
    if (result[0]) {
        Message({ message: '保存成功' })
    } else {
        Message({ message: '保存失败', type: 'warn' })
    }
}

async function updateMeta() {
    await updateBookMeta()
}

onMounted(async () => {
    let [ok, data] = await getBookMeta(
        (route.params['bookPath'] as string[]).join('/'),
        locale.value
    )

    if (ok) {
        metas.value = data['data']
    }
})
</script>
