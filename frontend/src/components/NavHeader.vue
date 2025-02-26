<template>
  <div class="header-container">
    <div class="header-left flex-box">
      <el-icon @click="store.commit('collapseTag')" size="20" class="icon flex-box">
        <Fold/>
      </el-icon>
      <el-scrollbar ref="scrollbarRef" @wheel.prevent="handleScroll" >
        <div ref="innerRef" class="scrollbar-li-item">
          <p
            v-for="(item, index) in selectMenu"
            :key="item.path"
            class="tab flex-box scrollbar-li-item"
            :class="{selected:item.path===nowRoute.path}"
        >
          <el-icon size="12" class=" flex-box">
            <component :is="item.icon"/>
          </el-icon>
          <router-link class="text flex-box" :to="{path:item.path}">
            {{ item.title }}
          </router-link>
          <el-icon size="12" class="close flex-box"
                   @click="closeTab(item,index)"
          >
            <Close/>
          </el-icon>
          </p>
        </div>
      </el-scrollbar>
    </div>
    <div class="header-right">
      <el-switch
          v-model="themeLight"
          inline-prompt
          :active-icon="Moon"
          :inactive-icon="Sunny"
          @click="toggleDark()"
      >
      </el-switch>
    </div>

  </div>
</template>
<script setup lang="ts">

import {Close, Fold, Moon, Sunny} from "@element-plus/icons-vue";
import {computed, ref} from "vue";
import AvatarSrc from "@/assets/images/appicon.png";
import {useStore} from "vuex"
import {useRoute, useRouter} from "vue-router";
import {ElScrollbar} from 'element-plus'
import {DashboardRouterItem} from "../router";
import {useColorMode, useDark, useToggle} from '@vueuse/core'
//当前路由
const nowRoute = useRoute()
const router = useRouter()

const avatar = ref(AvatarSrc)
const innerRef = ref<HTMLDivElement>()
const store = useStore()
const selectMenu = computed(() => store.state.menu.selectMenu)
const scrollbarRef = ref()

const themeLight = ref(true)
const system = useColorMode()
system.value = 'light'
const isDark = useDark()
const toggleDark = useToggle(isDark)
//关闭标签页
const closeTab = (item: any, index: number) => {
  store.commit("removeTab", item)
  if (nowRoute.path !== item.path) {
    return
  }
  if (nowRoute.path === item.path) {
    if (selectMenu.value.length === index) {
      // 如果当前路由是最后一个标签页，则跳转到首页
      if (selectMenu.value.length === 0) {
        router.push('/dashboard')
        store.commit("addMenu", DashboardRouterItem.meta)
      } else {
        router.push(selectMenu.value[index - 1].path)
      }
    } else {
      router.push(selectMenu.value[index].path)
    }
  }
}

function handleScroll(e: any) {
  const wheelDelta = e.wheelDelta || -e.deltaY * 40
  scrollbarRef.value.setScrollLeft(scrollbarRef.value.wrapRef.scrollLeft - wheelDelta)
}
</script>

<style scoped>
.flex-box{
  display: flex;
  align-items: center;
  height: 100%;
}
.header-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 100%;
  padding-right: 10px;
  .header-left{
    height: 100%;
    .icon{
      width: 40px;
      height: 100%;
    }
    .icon:hover{
      cursor: pointer;
    }

    .tab {
      padding: 0 10px;
      height: 100%;

      .text {
        margin: 0 5px;
      }

      .close {
        visibility: hidden;
      }

      &.selected {

        a {
          color: #409eff;
        }

        i {
          color: #409eff;
        }
      }
    }

    .tab:hover {

      .close {
        visibility: inherit;
        cursor: pointer;
      }
    }
  }
  .header-right{
    border: none;
    margin: 0;
    .user-name{
      margin-left: 10px;
    }
    .right-avatar{
      width: 30px;
      height: 30px;
    }
  }

  a {
    height: 100%;
    font-size: 15px;
  }
}

.scrollbar-li-item {
  height: 40px;
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
}
</style>
