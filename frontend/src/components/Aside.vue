<template>
  <el-menu
      :style="{width:!isCollapse?'149px':'64px'}"
      default-active="2"
      class="aside-container"
      @open="handleOpen"
      @close="handleClose"
      :collapse="isCollapse"
  >
    <p class="company-name" v-if="!isCollapse">{{ SysConfig.CompanyName }}</p>
    <p class="company-name" v-else>
      <el-icon>
        <img :src="avatar" alt=""/>
      </el-icon>
    </p>
    <TreeMenu class="tree-menu-class" :menuChildren="menuChildren"/>
  </el-menu>
</template>
<script setup lang="ts">
import {SysConfig} from "../config/SysConfig";
import TreeMenu from "./TreeMenu.vue";
import {useRouter} from 'vue-router'
import {computed, reactive, ref} from "vue"
import {useStore} from "vuex"
import AvatarSrc from "@/assets/images/appicon.png";

const avatar = ref(AvatarSrc)

const router = useRouter()
const menuChildren = reactive(router.options.routes[0])

const store = useStore()
const isCollapse = computed(() => store.state.menu.isCollapseTag)

const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
</script>

<style scoped>
.aside-container {
  height: 100%;
  .company-name{
    height: 40px;
    font-size: 20px;
    padding-top: 7px;
  }
}

</style>
