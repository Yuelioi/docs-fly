<template>
  <template class="bg-base-100" v-if="width > 720">
    <PcHeader :filteredNavs="filteredNavs"></PcHeader>
  </template>
  <template v-else>
    <MobileHeader :filteredNavs="filteredNavs"></MobileHeader>
  </template>
</template>

<script setup lang="ts">
import { Nav } from '@/models/home'
import { useWindowSize } from '@vueuse/core'
const { width } = useWindowSize()

const navs = ref<Nav[]>([])
//
const filteredNavs = computed(() => {
  return navs.value.slice().sort((pre, next) => {
    return (pre.metadata.order = next.metadata.order)
  })
})
onMounted(async () => {
  const nav_data = await getDBNav()
  // 验证Nav数据库信息并排序
  if (nav_data) {
    navs.value = nav_data.data
  } else {
    const [ok, data] = await getNav()
    if (ok) {
      navs.value = data['data'].sort(
        (pre: Nav, next: Nav) => pre.metadata.order - next.metadata.order
      )
      await addDBNav(navs.value)
    } else {
      navs.value = [new Nav()]
    }
  }
})
</script>

<style scoped lang="css">
.cat-menu a.router-link-active.router-link-exact-active {
  border-left: #0088ff solid 6px;
}
</style>
