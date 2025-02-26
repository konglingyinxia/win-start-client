<template>
  <el-drawer
      v-model="logVisible"
      :before-close="handleClose"
      @opened="addEventListenerLog"
      :size="elDrawerSize"
  >
    <template #title>
      <div class="log-drawer-header-cls">
        <el-text>{{ titleData }}</el-text>
        <el-button link icon="FullScreen" type="info" style="padding-right: 10px" @click="handleFullScreen">
        </el-button>
      </div>
    </template>
    <codemirror
        :autofocus="true"
        :indent-with-tab="true"
        :tabSize="4"
        style="margin-top: 20px; height: calc(100vh - 230px);text-align: left;"
        :lineWrapping="true"
        :matchBrackets="true"
        theme="cobalt"
        :extensions="extensions"
        v-model="logInfo"
        :styleActiveLine="true"
        @ready="handleReady"
        :disabled="true"
    />
    <template #footer>
      <el-button @click="handleClose">关闭</el-button>
    </template>
  </el-drawer>
</template>

<script setup lang="ts">

import {defineExpose, ref, shallowRef} from 'vue';
import {Codemirror} from "vue-codemirror"
import {javascript} from '@codemirror/lang-javascript';
import {oneDark} from '@codemirror/theme-one-dark';
import {EventsOff, EventsOn} from "../../../../wailsjs/runtime/runtime"
import * as logApi from '../../../../wailsjs/go/api/LogManager'
import {ElMessage} from 'element-plus';

const data = ref();
const titleData = ref();
const logInfo = ref<string>('');
const extensions = [javascript(), oneDark];
const logVisible = ref(false);
const EventsOnName = ref();
const elDrawerSize = ref("60%");
const fullScreenSet = ref(false)

const view = shallowRef();
const handleReady = (payload) => {
  view.value = payload.view;
};
const handleClose = () => {
  logApi.CloseSystemLog().then((res: any) => {
    if (res.code === 200) {

    } else {
      ElMessage.error(res.msg);
    }
  })
  //移除监听
  EventsOff(EventsOnName.value);
  logVisible.value = false;
  logInfo.value = '';
  data.value = null;
}
const addEventListenerLog = () => {
  //读取日志
  logApi.SystemLog().then((res: any) => {
    if (res.code === 200) {
      EventsOnName.value = "system-log";
      EventsOn(EventsOnName.value, function (params: any) {
        logInfo.value += params + '\n';
        const state = view.value.state;
        view.value.dispatch({
          selection: {anchor: state.doc.length, head: state.doc.length},
          scrollIntoView: true,
        });
      })
    } else {
      ElMessage.error(res.msg);
    }
  }).catch((err: any) => {
    ElMessage.error(err.msg);
  });

}

const acceptParams = () => {
  logVisible.value = true;
  elDrawerSize.value = "60%";
  titleData.value = "系统日志";
}
const handleFullScreen = () => {
  if (fullScreenSet.value) {
    elDrawerSize.value = "60%";
    fullScreenSet.value = false;
  } else {
    elDrawerSize.value = "100%";
    fullScreenSet.value = true;
  }
}

defineExpose({
  acceptParams,
});
</script>

<style scoped>

</style>
