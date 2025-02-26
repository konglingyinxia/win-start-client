<template>
  <el-row :gutter="20">
    <el-col :span="16">
      <el-card>
        <template #header>
          <div class="setting-title" style="text-align: left">
            <el-text size="large" style="font-weight: bold">概览</el-text>
            <el-button type="primary" style="float: right;margin-top: -5px;" plain
                       @click="handleOverviewClickRefresh"
            >刷新
            </el-button>
          </div>
        </template>
        <el-row>
          <el-col :span="8">
            <el-text size="large">应用</el-text>
            <div>
              <el-space size="large" direction="vertical">
                <el-text type="success">运行中:
                  <el-text size="large" type="primary" style="font-size: 20px;">
                    {{ overviewData.runningNum }}
                  </el-text>
                </el-text>
                <el-text type="warning">已停止:
                  <el-text size="large" type="primary" style="font-size: 20px;">
                    {{ overviewData.stoppedNum }}
                  </el-text>
                </el-text>
              </el-space>
            </div>
          </el-col>
          <el-col :span="8">
            <el-text size="large">环境</el-text>
            <div>
              <el-text size="large" type="primary" style="font-weight: bold;font-size: 25px;">
                {{ overviewData.envNum }}
              </el-text>
            </div>
          </el-col>
          <el-col :span="8">
            <el-text size="large">存储</el-text>
            <div>
              <el-space size="medium" direction="vertical" alignment="flex-start">
                <el-text type="success">应用:
                  <el-text size="large" type="primary" style="font-size: 18px;">
                    {{ overviewDataApp.size }}
                  </el-text>
                  <span style="color: #909399;font-size: 12px;"> {{ overviewDataApp.unit }}</span>
                </el-text>
                <el-text type="warning">环境:
                  <el-text size="large" type="primary" style="font-size: 18px;">
                    {{ overviewDataEnv.size }}
                  </el-text>
                  <span style="color: #909399;font-size: 12px;">{{ overviewDataEnv.unit }}</span>
                </el-text>
                <el-text type="info">日志:
                  <el-text size="large" type="primary" style="font-size: 18px;">
                    {{ overviewDataLog.size }}
                  </el-text>
                  {{ overviewDataLog.unit }}
                </el-text>
              </el-space>
            </div>
          </el-col>
        </el-row>
      </el-card>
      <!--     状态 -->
      <el-card style="margin-top: 10px">
        <template #header>
          <div class="setting-title" style="text-align: left">
            <el-text size="large" style="font-weight: bold">状态</el-text>
            <el-button type="primary" style="float: right;margin-top: -5px;" plain
                       @click="handleClickLog"
            >系统日志
            </el-button>
          </div>
        </template>


      </el-card>
    </el-col>
    <el-col :span="8">
      <!--  全局设置     -->
      <el-card>
        <template #header>
          <div class="setting-title" style="text-align: left">
            <el-text size="large" style="font-weight: bold">全局设置</el-text>
          </div>
        </template>
        <div>
          <div class="setting-title" style="text-align: left;">
            <el-row>
              <!--              <el-col :span="12">-->
              <!--                开机自启：-->
              <!--                <el-switch v-model="autoStart" @change="changeAutoStart" active-color="#13ce66"-->
              <!--                           inactive-color="#ff4949"/>-->
              <!--              </el-col>-->
              <el-col :span="12">
                <el-button type="primary" @click="handleDefaultGlobal">设置</el-button>
              </el-col>
            </el-row>
          </div>
        </div>
      </el-card>
      <!--  系统信息     -->
      <el-card style="margin-top: 20px">
        <template #header>
          <div class="setting-title" style="text-align: left">
            <el-text size="medium" style="font-weight: bold">系统信息</el-text>
          </div>
        </template>
        <div style="text-align: left;">
          <el-space size="medium" direction="vertical" alignment="flex-start" wrap>
            <div>
              <el-space size="large">
                <el-text size="medium" style="width: 56px">主机名称</el-text>
                <el-text size="small" type="primary" style="font-weight: bold;">{{ homeOsInfo.infoStat?.hostname }}
                </el-text>
              </el-space>
            </div>
            <div>
              <el-space size="large">
                <el-text size="medium">发行版本</el-text>
                <el-text size="small" type="primary" style="font-weight: bold;">
                  {{ homeOsInfo.infoStat?.platform }} {{ homeOsInfo.infoStat?.platformFamily }}
                  {{ homeOsInfo.infoStat?.platformVersion }}
                </el-text>
              </el-space>
            </div>
            <div>
              <el-space size="large">
                <el-text size="medium">内核版本</el-text>
                <el-text size="small" type="primary" style="font-weight: bold;">
                  {{ homeOsInfo.infoStat?.kernelVersion }}
                </el-text>
              </el-space>
            </div>
            <div>
              <el-space size="large">
                <el-text size="medium">系统类型</el-text>
                <el-text size="small" type="primary" style="font-weight: bold;">
                  {{ homeOsInfo.infoStat?.os }}_{{ homeOsInfo.infoStat?.kernelArch }}
                </el-text>
              </el-space>
            </div>
            <div>
              <el-space size="large">
                <el-text size="medium">主机地址</el-text>
                <el-text size="small" type="primary" style="font-weight: bold;">
                  {{ homeOsInfo.ipAddr }}
                </el-text>
              </el-space>
            </div>
            <div>
              <el-space size="large">
                <el-text size="medium">启动时间</el-text>
                <el-text size="small" type="primary" style="font-weight: bold;">
                  {{ formatDate(homeOsInfo.infoStat?.bootTime * 1000) }}
                </el-text>
              </el-space>
            </div>
            <div>
              <el-space size="large">
                <el-text size="medium">运行时间</el-text>
                <el-text size="small" type="primary" style="font-weight: bold;">
                  {{ homeOsUptime.day }} 天 {{ homeOsUptime.hours }} 小时 {{ homeOsUptime.minutes }} 分钟
                  {{ homeOsUptime.seconds }} 秒
                </el-text>
              </el-space>
            </div>
          </el-space>
        </div>
      </el-card>
    </el-col>
  </el-row>
  <!--日志弹窗-->
  <SystemLog ref="dialogSystemLogRef"></SystemLog>
  <!-- 默认网页设置 -->
  <el-dialog title=全局设置
             draggable
             v-model="defaultWebVisible"
             width="50%"
             :before-close="handleDefaultWebClose">
    <el-form :model="defaultWebForm" label-width="100px">
      <el-form-item label="默认网页">
        <el-input v-model="defaultWebForm.defaultWeb" placeholder="跳转网页(默认: http://127.0.0.1:8005)">
        </el-input>
      </el-form-item>
      <el-form-item label="默认打开" prop="openDefaultWeb">
        <el-switch v-model="defaultWebForm.openDefaultWeb"/>
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button type="primary" @click="handleDefaultWebSave">保存</el-button>
      <el-button @click="defaultWebVisible = false">取消</el-button>
    </div>
  </el-dialog>
</template>

<script setup lang="ts">
import {onMounted, onUnmounted, ref} from 'vue'
import SystemLog from "@/views/dashboard/log/Index.vue";
import {ElMessage} from "element-plus";
import {dto, model, res} from "../../../wailsjs/go/models"
import * as globalApi from "../../../wailsjs/go/api/GlobalSettingManager"
import * as systemApi from "../../../wailsjs/go/api/SystemOs";
import * as dashboardApi from "../../../wailsjs/go/api/DashboardManager";
import {formatBytes, formatDate, formatDateDHMS} from "../../utils/util"

const dialogSystemLogRef = ref()
const autoStart = ref(false)
const homeOsInfo = ref(dto.HomeOsInfo.createFrom())
const homeOsUptime = ref({day: 0, hours: 0, minutes: 0, seconds: 0})
const overviewData = ref(res.OverviewData.createFrom())
const overviewDataApp = ref({
  size: "0",
  unit: 'B'
});
const overviewDataEnv = ref({
  size: "0",
  unit: 'B'
});
const overviewDataLog = ref({
  size: "0",
  unit: 'B'
});
const changeAutoStart = (val: boolean) => {
  autoStart.value = val
}

onMounted(() => {
  getHomeOsInfo();
  handleOverviewClickRefresh();
  startTimer();
})
onUnmounted(() => {
  stopTimer();
})

const defaultWebVisible = ref(false)
const defaultWebForm = ref(model.GlobalSetting.createFrom(
    {
      defaultWeb: "http://127.0.0.1:8005",
      openDefaultWeb: false,
    }
))
const handleDefaultGlobal = () => {
  defaultWebVisible.value = true
  globalApi.GetGlobalSetting().then((res) => {
    if (res.code === 200) {
      if (res.data != null) {
        defaultWebForm.value = res.data
      }
    } else {
      ElMessage.error("获取设置失败:" + res.msg)
    }
  }).catch((err) => {
    ElMessage.error("获取设置失败:" + err)
  })
}
const handleDefaultWebSave = () => {
  globalApi.SetGlobalSetting(defaultWebForm.value).then((res) => {
    if (res.code === 200) {
      defaultWebVisible.value = false
      defaultWebForm.value = model.GlobalSetting.createFrom()
      ElMessage.success("保存成功")
    } else {
      ElMessage.error("保存失败:" + res.msg)
    }
  }).catch((err) => {
    ElMessage.error("保存失败:" + err)
  })
}
const handleDefaultWebClose = () => {
  defaultWebVisible.value = false
}
const handleClickLog = () => {
  dialogSystemLogRef.value!.acceptParams()
}
//概览刷新
const handleOverviewClickRefresh = () => {
  dashboardApi.GetDashboard().then(res => {
    overviewData.value = res
    overviewDataApp.value = formatBytes(res.storageApp)
    overviewDataEnv.value = formatBytes(res.storageEnv)
    overviewDataLog.value = formatBytes(res.storageLog)
  }).catch((err) => {
    ElMessage.error(err)
  })
}

//系统信息

let intervalId: any = null;
const getHomeOsInfo = () => {
  systemApi.HomeOsInfo().then(res => {
    homeOsInfo.value = res
    homeOsUptime.value = formatDateDHMS(res.infoStat.uptime)
  }).catch(err => {
    ElMessage.error("获取系统信息失败:" + err)
  })
}


// 增加定时任务
function startTimer() {
  intervalId = setInterval(() => {
    getHomeOsInfo()
  }, 5000);
}

// 清除定时任务
function stopTimer() {
  clearInterval(intervalId);
}

</script>

<style scoped>

</style>
