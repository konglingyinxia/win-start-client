<template>
  <div class="common-layout">
    <el-container class="container-me">
      <Aside />
      <el-container>
        <el-header>
          <NavHeader/>
        </el-header>
        <el-main>
            <RouterView/>
        </el-main>
      </el-container>
    </el-container>
  </div>
  <AppLaunching ref="appLaunchingRef"></AppLaunching>
  <AppStopping ref="appStoppingRef"></AppStopping>
</template>
<script setup lang="ts">
import Aside from "../components/Aside.vue";
import NavHeader from "../components/NavHeader.vue"
import {onBeforeMount, onMounted, onUnmounted, ref} from "vue";
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import {DashboardRouterItem} from "../router"
import AppLaunching from "./AppLaunching.vue"
import AppStopping from "./AppStopping.vue"
import {EventsOff, EventsOn} from "../../wailsjs/runtime/runtime"

const appLaunchingRef = ref()
const appStoppingRef = ref()

const store = useStore()

const router = useRouter()

const listenAppLaunchStop = () => {
  EventsOn("app_start_stop_event", (res: any) => {
    //应用程序关闭事件
    if (res.eventType === "stop") {
      appStoppingRef.value!.acceptParams(res)
    }
    //应用程序初次启动事件
    if (res.eventType === "start") {
      appLaunchingRef.value!.acceptParams(res)
    }
  })
}
onBeforeMount(() => {
  listenAppLaunchStop()
})
onMounted(() => {
  router.push(DashboardRouterItem.meta.path)
  store.commit('addMenu', DashboardRouterItem.meta)
})
onUnmounted(() => {
  EventsOff("app_stop_stop_event")
})
</script>

<style scoped>
.common-layout {
  @apply h-full;
}

.container-me {
  @apply h-full;
}

</style>
