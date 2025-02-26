<template>
  <div class="info-container">
    <el-card shadow="hover" class="box-card">
      <template #header>
        <div class="card-header">
          <span>服务信息</span>
        </div>
      </template>
      <div class="info-list">
        <div class="info-item">
          <span class="label">运行目录：</span>
          <span>{{ osInfo.rootDir }}</span>
        </div>
        <div class="info-item">
          <span class="label">本机IP：</span>
          <span>{{ osInfo.ipAddr }}</span>
        </div>
      </div>
    </el-card>

    <el-card shadow="hover" class="box-card">
      <template #header>
        <div class="card-header">
          <span>系统信息</span>
        </div>
      </template>
      <div class="info-list">
        <div class="info-item">
          <span class="label">主机名：</span>
          <span>{{ infoStat?.hostname ?? '--' }}</span>
        </div>
        <div class="info-item">
          <span class="label">操作系统：</span>
          <span>{{ infoStat?.os ?? '--' }}【{{infoStat.platform}}】 </span>
        </div>
        <div class="info-item">
          <span class="label">发行版本：</span>
          <span>{{ infoStat?.platformFamily ?? '--' }}</span>
        </div>
        <div class="info-item">
          <span class="label">系统版本：</span>
          <span>{{ infoStat?.kernelVersion ?? '--' }}  {{ infoStat?.kernelArch ?? '--' }}</span>
        </div>
        <div class="info-item">
          <span class="label">版本号：</span>
          <span>{{ infoStat?.platformVersion ?? '--' }}</span>
        </div>
        <div class="info-item">
          <span class="label">进程数：</span>
          <span>{{ infoStat?.procs ?? '--' }}</span>
        </div>
        <div class="info-item">
          <span class="label">运行时间：</span>
          <span>{{ uptime.day }} 天 {{ uptime.hours }} 小时 {{ uptime.minutes }} 分钟
                  {{ uptime.seconds }} 秒</span>
        </div>
      </div>
    </el-card>

    <el-card shadow="hover" class="box-card">
      <template #header>
        <div class="card-header">
          <span>硬件信息</span>
        </div>
      </template>
      <div class="info-list">
        <div class="info-item">
          <span class="label">处理器：</span>
          <span>{{ osInfo.cpu }}</span>
        </div>
        <div class="info-item">
          <span class="label">显卡(GPU)：</span>
          <span>{{ osInfo.gpu }}</span>
        </div>
        <div class="info-item">
          <span class="label">内存：</span>
          <div style="display: flex;width: 50%;align-items: center;justify-content: flex-end">
            <el-progress class="disk-progress-cls"
                       :color="customColors"
                       type="line"
                       :text-inside="true"
                       :stroke-width="20"
                       :percentage="osInfo.memory?.percent??0">
              <span style="color: black">{{ osInfo.memory?.percent ?? 0 }}%</span>
            </el-progress>
            <span style="color: black">{{
                                osInfo.memory?.usedMemory
                              }} GB/ {{ osInfo.memory?.totalMemory }} GB</span>
          </div>
        </div>
        <div class="info-item">
          <span class="label">磁盘：</span>
          <div style="display: flex;width: 50%;align-items: center;justify-content: flex-end">
            <el-progress class="disk-progress-cls" type="line"
                       :color="customColorsDisk"
                       :text-inside="true"
                       :stroke-width="20"
                       :percentage="diskUsePercent">
          </el-progress>
            <span style="color:#000;"> {{ diskHasUse.size }} {{ diskHasUse.unit }}/ {{
                diskAllSize.size
              }} {{ diskAllSize.unit }}</span>
          </div>
        </div>
        <div class="info-item">
          <span class="label">BIOS：</span>
          <span>{{ osInfo.bios }}</span>
        </div>
        <div class="info-item">
          <span class="label">制造商：</span>
          <span>{{ osInfo.manufacturer }}</span>
        </div>
        <div class="info-item">
          <span class="label">产品名称：</span>
          <span>{{ osInfo.productName }}</span>
        </div>
        <div class="info-item">
          <span class="label">序列号：</span>
          <span>{{  }}</span>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import {onMounted, onUnmounted, ref} from 'vue';
import {OsInfo} from "../../../../wailsjs/go/api/SystemOs";
import {dto, host} from "../../../../wailsjs/go/models";
import {ElMessage} from 'element-plus'
import {formatBytes, formatDateDHMS} from "../../../utils/util";

const osInfo = ref(dto.SystemInfo.createFrom({}));
const infoStat = ref(host.InfoStat.createFrom({}))
const diskUsePercent = ref("");
const diskAllSize = ref({
  size: "0",
  unit: 'GB'
});
const diskHasUse = ref({
  size: "0",
  unit: 'GB'
});
const customColors = [
  {color: '#c5ffa3', percentage: 20},
  {color: '#12ab00', percentage: 40},
  {color: '#00f891', percentage: 60},
  {color: '#e8a961', percentage: 80},
  {color: '#ef5353', percentage: 90},
  {color: '#fa0000', percentage: 100},
]
const customColorsDisk = [
  {color: '#c3c7ff', percentage: 20},
  {color: '#84a1ff', percentage: 40},
  {color: '#4b8aff', percentage: 60},
  {color: '#177dff', percentage: 80},
  {color: '#0027ff', percentage: 90},
  {color: '#a30000', percentage: 100},
]
let intervalId: any = null;
// 初始化状态
const totalSeconds = ref(0);
const uptime = ref({day: 0, hours: 0, minutes: 0, seconds: 0})

async function getOsInfo() {
  OsInfo().then(res => {
      osInfo.value = res;
      infoStat.value = res.infoStat
      uptime.value = formatDateDHMS(res.infoStat.uptime)
      diskUsePercent.value = res.disk.usedPercent.toFixed(2)
      diskAllSize.value = formatBytes(res.disk.total);
      diskHasUse.value = formatBytes(res.disk.used);
  }).catch(err => {
    ElMessage.error(err)
  })
}

// 增加定时任务
function startTimer() {
  intervalId = setInterval(() => {
    getOsInfo()
  }, 5000);
}

// 清除定时任务
function stopTimer() {
  clearInterval(intervalId);
}

onMounted(() => {
  getOsInfo();
  startTimer();
})
onUnmounted(() => {
  stopTimer();
})
</script>
<style scoped>
.info-container {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.box-card {
  margin-top: 20px;
  width: 100%;
  max-width: 800px;
}

.card-header {
  font-size: 20px;
  color: #ad77ff;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.info-list {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.info-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 30px;


  .disk-progress-cls {
    width: 50%;
    align-items: normal;
    margin-bottom: 0;
  }
}

.label {
  font-size: 15px;
  color: #86abeb;
  font-weight: bold;
  text-align: right;
  width: 120px;
}

</style>
