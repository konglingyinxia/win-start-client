<template>
  <el-drawer
      v-model="logVisible"
      :title="titleData"
      :before-close="handleClose"
      direction="ltr"
      size="65%"
      :destroy-on-close="true"
      :close-on-click-modal="false"
      :close-on-press-escape="false"
  >
    <div style="text-align: left; margin-top: 20px"
         v-for="item in tableData"
         :key="item.id"
    >
      <el-row>
        <el-col :span="6">
          <label>
            <el-tag>{{ item.startOrder}}</el-tag>. {{ item.name }}
            <el-tag type="info" v-if="item.remark">{{ item.remark }}</el-tag>
          </label>
        </el-col>
        <el-col :span="2">
          <el-tag type="success" v-if="item.autoStart === true">自启</el-tag>
          <el-tag type="warning" v-else>非自启</el-tag>
        </el-col>
        <el-col :span="14">
          <el-tag type="primary" v-if="item.status === 'started'">启动中...</el-tag>
          <el-tag type="success" v-else-if="item.status === 'running'">运行中</el-tag>
          <span v-else-if="item.status === 'failed'">
              <el-tag type="danger">自启失败</el-tag>
              <el-text style="margin-left: 10px" type="info" v-if="item.error">{{ item.error }}</el-text>
          </span>
          <span v-else>
            <el-tag v-if="item.autoStart === true" type="warning">等待中</el-tag>
            <el-tag v-if="item.autoStart === false" type="info">不启动</el-tag>
          </span>
        </el-col>
      </el-row>
    </div>
    <div style="text-align: left; margin-top: 20px" v-if="startEnd">
      <el-text type="success" v-if="startEndError === ''">全部启动成功...</el-text>
      <el-text type="danger" v-else><span style="color: #72767b">启动失败：</span>{{ startEndError }}</el-text>
    </div>
  </el-drawer>
</template>
<script setup lang="ts">
import {defineExpose, onMounted, ref} from "vue";
import {ElMessage} from "element-plus";
import * as api from "../../wailsjs/go/api/AppManager"

const titleData = ref("应用启动");
const logVisible = ref(false);
const tableData = ref([])
const startEnd = ref(false)
const startEndError = ref("")


const listenAppLaunchStop = (res: any) => {
  //应用程序初次启动事件
  if (res.id != "" &&
      res.eventType === "start" &&
      res.eventStatus === "running") {
    tableData.value.forEach(item => {
      if (item.id === res.id) {
        item.status = res.status;
        item.error = res.error
      }
    })
  }
  //应用程序初次启动事件
  if (res.id === "" &&
      res.eventType === "start" &&
      res.eventStatus === "end") {
    startEnd.value = true;
    startEndError.value = res.error;
  }
}
const listData = () => {
  api.List().then(res => {
    if (res.code === 200) {
      tableData.value = res.data;
    } else {
      ElMessage.error(res.msg);
    }
  }).catch(error => {
    ElMessage.error("获取应用列表失败：" + error);
  })
}

const handleClose = () => {
  logVisible.value = false;
}

const acceptParams = (res: any) => {
  logVisible.value = true;
  if (res.id === '' && res.eventStatus === 'running') {
  } else {
    listenAppLaunchStop(res);
  }
}

defineExpose({
  acceptParams,
});
onMounted(() => {
  listData()
})
</script>

<style scoped>

</style>
