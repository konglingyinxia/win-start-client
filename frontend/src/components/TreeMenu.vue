<template>
  <template v-for="(item, index) in menu" :key="index">
    <el-menu-item v-if="!item.children|| item.children.length===0"
                  :index="item.meta.id"
                  :key="item.meta.id"
                  @click="handleClick(item,item.meta.id)"
    >
      <el-icon size="20">
        <component :is="item.meta.icon"></component>
      </el-icon>
      <span>{{ item.meta.title }}</span>
    </el-menu-item>
    <el-sub-menu v-else :index="item.meta.id">
      <template #title>
        <el-icon size="20">
          <component :is="item.meta.icon"></component>
        </el-icon>
        <span>{{ item.meta.title }}</span>
      </template>
      <tree-menu :index="item.meta.id" :menuChildren="item"></tree-menu>
    </el-sub-menu>
  </template>
</template>
<script setup lang="ts">
import {useRouter} from "vue-router"
import {useStore} from "vuex"

const props = defineProps(['menuChildren', 'index'])
const menu = props.menuChildren.children
const router = useRouter()
const store = useStore()

const handleClick = (item: any, active: string) => {
  router.push(item.meta.path)
  store.commit('addMenu', item.meta)
}
</script>

<style scoped>

</style>
