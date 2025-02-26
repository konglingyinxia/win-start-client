<template>
  <div class="app-container-header flex flex-wrap gap-4">
    <el-button type="primary" @Click="handleAdd" plain>新增</el-button>
    <el-button type="success" @Click="handleReload" plain>刷新</el-button>
    <el-button-group>
      <el-button type="success" @Click="handleStartAll">启动</el-button>
      <el-button type="danger" @Click="handleStopAll">停止</el-button>
    </el-button-group>
    <div style="display: flex; align-items: center;">
      <label>应用根目录：
        <el-text type="warning"> {{ appHomePath.appHomeKey }}:</el-text>
        <el-text style="margin-left: 10px;" type="primary">{{ appHomePath.appHomePath }}</el-text>
        <el-button icon="FolderOpened" link @click="openFileExplorer(appHomePath.appHomePath)"></el-button>
      </label>
    </div>
  </div>
  <el-table :data="tableData"
            v-loading="tableDataLoading"
            ref="multipleTableRef"
            @selection-change="handleSelectionChange"
            style="width: 100%">
    <el-table-column type="selection" width="55"/>
    <el-table-column fixed="left" label="项目" min-width="150" prop="name">
      <template #default="scope">
        <div>
          <el-text type="primary">名称：
            <el-tag type="success">{{ scope.row.name }}</el-tag>
          </el-text>
        </div>
        <div>
          <el-text type="primary">备注：
            <el-tag type="info">{{ scope.row.remark }}</el-tag>
          </el-text>
        </div>
      </template>
    </el-table-column>
    <el-table-column label="启动属性" min-width="120" prop="startAttr">
      <template #default="scope">
        <el-row>
          <el-col :span="24">
            <el-text type="info">自启：
              <el-tag v-if="scope.row.autoStart === true" type="success">是</el-tag>
              <el-tag v-else-if="scope.row.autoStart === false" type="warning">否</el-tag>
            </el-text>
          </el-col>
          <el-col :span="24">
            <el-text type="info">顺序：
              <el-tag type="primary">{{ scope.row.startOrder }}</el-tag>
            </el-text>
          </el-col>
          <el-col :span="24">
            <el-text type="info">延时：
              <el-tag type="danger">{{ scope.row.startDelay }}</el-tag>
              秒
            </el-text>
          </el-col>
        </el-row>
      </template>
    </el-table-column>
    <el-table-column label="应用状态" min-width="150" prop="appRunStatus">
      <template #default="scope">
        <el-popover effect="light"
                    trigger="click"
                    placement="top-start"
                    width="400">
          <template #default>
            <el-row :gutter="20">
              <el-col :span="24" class="run-status-container-item-class">
                PID：
                <el-tag type="info" v-for="item in scope.row.runStatus.pids">
                  {{ item }}
                </el-tag>
              </el-col>
              <el-col :span="12" class="run-status-container-item-class">
                内存使用: <span v-if="scope.row.runMemory">{{ scope.row.runMemory.size }} {{
                  scope.row.runMemory.unit
                }}</span>
              </el-col>
              <el-col :span="12" class="run-status-container-item-class">
                内存占比: <span> {{ scope.row.runStatus.memPercent.toFixed(2) }} %</span>
              </el-col>
              <el-col :span="12" class="run-status-container-item-class">
                存储使用: {{ scope.row.diskSizeValue.size }} {{ scope.row.diskSizeValue.unit }}
              </el-col>
            </el-row>
          </template>
          <template #reference>
            <el-row
                v-loading="!scope.row.runStatusNoLoading"
            >
              <el-col :span="24">
                <el-space>
                  <el-text type="info">状态：
                  <el-tag v-if="scope.row.runStatus.status === 'running'" type="success">running</el-tag>
                  <el-tag v-if="scope.row.runStatus.status === 'stopped'" type="danger">stopped</el-tag>
                  <el-tag v-if="scope.row.runStatus.status === 'padding'" type="warning">padding</el-tag>
                </el-text>
                </el-space>

              </el-col>
              <el-col :span="24">
                <el-text type="info">CPU：{{ scope.row.runStatus.cpuPercent.toFixed(2) }}%</el-text>
              </el-col>
              <el-col :span="24">
                <el-text type="info">内存：{{ scope.row.runStatus.memPercent.toFixed(2) }}%</el-text>
              </el-col>
            </el-row>
          </template>
        </el-popover>
      </template>
    </el-table-column>
    <el-table-column label="端口" prop="ports">
      <template #default="scope">
        <el-button type="primary" plain link @click="handleRefreshPort(scope.row)">端口刷新</el-button>
        <div style="display: flex; margin-top: 10px; flex-wrap: wrap;justify-content: space-between">
          <el-tag v-for="item in scope.row.portsArr" type="info">{{ item }}</el-tag>
        </div>
      </template>
    </el-table-column>
    <el-table-column label="版本" prop="version"></el-table-column>
    <el-table-column label="目录" prop="appDir">
      <template #default="scope">
        <div>
          {{ scope.row.appDir }}
        </div>
        <el-button
            icon="FolderOpened" type="primary"
            link @click="openFileExplorer(appHomePath.appHomePath+appHomePath.separator+scope.row.appDir)"></el-button>
      </template>
    </el-table-column>
    <el-table-column label="创建时间" prop="createdAt"></el-table-column>
    <el-table-column fixed="right" label="操作" min-width="120">
      <template #default="scope">
        <div style="display: flex">
          <el-button type="primary" link @click="handleEdit(scope.row)">编辑</el-button>
          <el-button style="margin-left: 5px" type="text" @click="handleLogInfo(scope.row)">日志</el-button>
          <!-- 下拉命令        -->
          <el-dropdown style="margin-left: 5px">
            <el-button type="primary" link>
              命令
            </el-button>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item>
                  <el-button type="success" link @click="handleStart(scope.row)">启动</el-button>
                </el-dropdown-item>
                <el-dropdown-item>
                  <el-button type="danger" link @click="handleStop(scope.row)">停止</el-button>
                </el-dropdown-item>
                <el-dropdown-item>
                  <el-button type="warning" link @click="handleRestart(scope.row)">重启</el-button>
                </el-dropdown-item>
                <el-dropdown-item>
                  <el-button type="info" link @click="handleVisibleDelete(scope.row)">删除</el-button>
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>

        </div>
      </template>
    </el-table-column>
  </el-table>
  <!-- 删除确认弹窗 -->
  <el-dialog
      style="text-align: left"
      v-model="dialogDeleteVisible"
      title="删除"
      width="30%"
  >
    <div style="display: flex;align-items: center">
      <el-icon color="#909399" size="30px">
        <WarningFilled/>
      </el-icon>
      <span style="margin-left: 10px">确定要删除应用：
        <el-tag type="danger">{{ deleteRow?.name }}</el-tag> 吗？</span>
    </div>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="dialogDeleteVisible = false">取消</el-button>
        <el-button type="primary" @click="handleDelete(deleteRow)">确认</el-button>
      </div>
    </template>

  </el-dialog>
  <!--新增/编辑弹窗    -->
  <el-dialog
      :title="(form.id === ''|| form.id === undefined) ? '新增应用' : '编辑应用'"
      v-model="openForm"
      draggable
      width="65%"
  >
    <!--    环境根目录  -->
    <el-form :model="form"
             :rules="rules"
             status-icon
             label-width="auto"
             label-position="left">
      <el-form-item label="父级目录" prop="envHomePath">
        <el-input v-model="appHomePath.appHomeKey" disabled>
          <template #prepend>键名：</template>
        </el-input>
        <el-input v-model="appHomePath.appHomePath" disabled>
          <template #prepend>键值：</template>
          <template #suffix>
            <el-button icon="FolderOpened" type="primary" link
                       @click="openFileExplorer(appHomePath.appHomePath)"></el-button>
          </template>
        </el-input>
      </el-form-item>
      <el-form-item label="类型" prop="type">
        <!--          mysql,redis,nginx,custom-->
        <el-select v-model="form.type" placeholder="请选择类型">
          <el-option label="mysql" value="mysql"></el-option>
          <el-option label="redis" value="redis"></el-option>
          <el-option label="nginx" value="nginx"></el-option>
          <el-option label="自定义" value="custom" selected></el-option>
        </el-select>
      </el-form-item>
      <el-row :gutter="20">
        <el-col :span="8">
          <el-form-item label="名称" prop="name">
            <el-input v-model="form.name" placeholder="请输入名称"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="目录" prop="appDir">
            <el-input v-model="form.appDir" placeholder="请输入目录"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="版本" prop="version">
            <el-input v-model="form.version" placeholder="请输入版本"></el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item label="启动命令" prop="startCmd">
        <el-input v-model="form.startCmd" placeholder="请输入启动命令"></el-input>
      </el-form-item>
      <el-form-item label="停止命令" prop="stopCmd">
        <el-input v-model="form.stopCmd" placeholder="默认杀死进程"></el-input>
      </el-form-item>
      <el-form-item label="重启命令" prop="restartCmd">
        <el-input v-model="form.restartCmd" placeholder="默认杀死进程后启动"></el-input>
      </el-form-item>
      <el-row :gutter="20">
        <el-col :span="6">
          <el-form-item label="是否自启" prop="autoStart">
            <el-switch v-model="form.autoStart"/>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="启动顺序" prop="startOrder">
            <el-input-number v-model="form.startOrder" placeholder="请输入名称"
                             :min="1" :max="9999"
            ></el-input-number>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="启动延时" prop="startDelay">
            <el-input-number v-model="form.startDelay" placeholder="请输入版本"
                             :min="5" :max="3600"
            >
              <template #suffix>
                <span>秒</span>
              </template>
            </el-input-number>
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item label="备注" prop="remark">
        <el-input v-model="form.remark" placeholder="请输入备注"></el-input>
      </el-form-item>
      <el-form-item label="依赖环境" prop="envId">
        <el-select v-model="form.envId" placeholder="请选择依赖环境">
          <el-option label="无" value=""></el-option>
          <el-option v-for="(item, index) in externalEnvList"
                     :key="index"
                     :label="item.envName+'('+item.remark+')'"
                     :value="item.id"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="环境变量">
        <el-row align="top">
          <el-col :span="1">
            <el-button icon="plus" type="text" @click="addKeyValuePair"></el-button>
          </el-col>
          <el-col v-for="(item, index) in keyValuePairs" :key="index">
            <el-row :gutter="10">
              <el-col :span="7">
                <el-input v-model="item.key" placeholder="键"></el-input>
              </el-col>
              <el-col :span="16">
                <el-input v-model="item.value" placeholder="引用变量示例格式 ${KEY} / $KEY / %KEY%"></el-input>
              </el-col>
              <el-col :span="1">
                <el-button icon="minus" type="text" @click="removeKeyValuePair(index)"></el-button>
              </el-col>
            </el-row>
          </el-col>
        </el-row>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button type="primary" @click="handleSubmit">提交</el-button>
      <el-button @click="openForm = false">取消</el-button>
    </template>
  </el-dialog>

  <!--日志弹窗-->
  <AppLog ref="dialogAppLogRef"></AppLog>
</template>
<script setup lang="ts">
import {onMounted, onUnmounted, ref} from "vue";
import {api as comm, req} from "../../../../wailsjs/go/models"
import * as api from "../../../../wailsjs/go/api/AppManager"
import * as envApi from "../../../../wailsjs/go/api/EnvManager"
import * as openApi from "../../../../wailsjs/go/api/CommManager"
import type {FormRules, TableInstance} from 'element-plus'
import {ElMessage} from 'element-plus'
import {formatBytes} from "../../../utils/util"
import AppLog from "@/views/application/app/log/Index.vue"

const openForm = ref(false);
const form = ref(req.AppReq.createFrom())
const appHomePath = ref(comm.AppManager.createFrom())
const keyValuePairs = ref([])
// 验证环境弹窗
const checkEnvFormVisible = ref(false)
const checkEnvForm = ref(req.EnvCheckReq.createFrom())
const handleCHackResult = ref('')
const externalEnvList = ref([])

const multipleTableRef = ref<TableInstance>()
const multipleSelection = ref([])
const dialogAppLogRef = ref();

const rules = ref<FormRules<req.AppReq>>({
  name: [
    {required: true, message: '请输入环境名称', trigger: 'blur'},
    {min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur'}
  ],
  appDir: [
    {required: true, message: '请输入环境目录', trigger: 'blur'},
    {min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur'}
  ],
  startCmd: [
    {required: true, message: '请输入启动命令', trigger: 'blur'},
    {min: 2, max: 200, message: '长度在 2 到 200 个字符', trigger: 'blur'}
  ],
})
const checkEnvFormRules = ref<FormRules<req.EnvCheckReq>>({
  command: [
    {required: true, message: '请输入执行命令', trigger: 'blur'},
  ],
})

const dialogDeleteVisible = ref(false)
const deleteRow = ref()
onMounted(() => {
  api.GetAppHomePath().then(res => {
    appHomePath.value = res;
  });
  listData();
  startTimer();
})
onUnmounted(() => {
  stopTimer();
})

const tableData = ref([])
const tableDataLoading = ref(false)
// 刷新环境列表
const handleReload = () => {
  listData();
  ElMessage.success("刷新成功！");
}

// 获取应用列表
const listData = () => {
  api.List().then(res => {
    if (res.code === 200) {
      tableData.value = res.data;
      if (res.data != null) {
        tableData.value.forEach(item => {
          item.diskSizeValue = formatBytes(item.diskSize)
          if (item.ports) {
            item.portsArr = JSON.parse(item.ports)
          } else {
            item.portsArr = []
          }
          item.appRunStatusVisible = false
        })
        getRunStatus()
      }
    } else {
      ElMessage.error(res.msg);
    }
  }).catch(error => {
    ElMessage.error("获取应用列表失败：" + error);
  })
}
//获取应用状态
const getRunStatus = () => {
  if (tableData.value === undefined || tableData.value === null || tableData.value.length === 0) {
    return;
  }
  tableData.value.forEach(item => {
    api.Status(item.id).then(res => {
      if (res.code === 200) {
        item.runStatus = res.data;
        if (res.data.memSize) {
          item.runMemory = formatBytes(res.data.memSize)
        }
        item.runStatusNoLoading = true;
      }
    }).catch(err => {
      ElMessage.error("获取进程状态失败：" + err);
    })
  })
}

//获取外部环境列表
const getExternalEnvList = () => {
  envApi.List().then(res => {
    externalEnvList.value = res.data;
  })
}

const resetForm = () => {
  form.value = req.AppReq.createFrom()
  form.value.type = "custom"
  form.value.startDelay = 5
  form.value.startOrder = 1
  form.value.autoStart = true
  keyValuePairs.value = []
}

const handleAdd = () => {
  resetForm()
  //打开新增弹窗
  openForm.value = true;
  //获取外部环境列表
  getExternalEnvList();
}
// 添加新的键值对
const addKeyValuePair = () => {
  keyValuePairs.value.push({key: '', value: ''});
};
// 删除特定的键值对
const removeKeyValuePair = (index: number) => {
  keyValuePairs.value.splice(index, 1);
};
//删除环境
const handleVisibleDelete = (row: any) => {
  dialogDeleteVisible.value = true
  deleteRow.value = row
}
const handleDelete = (row: any) => {
  api.Delete(row.id).then(res => {
    if (res.code === 200) {
      ElMessage.success("删除应用成功！");
      dialogDeleteVisible.value = false
      deleteRow.value = {}
      listData();
    } else {
      ElMessage.error(res.msg);
    }
  }).catch(err => {
    ElMessage.error("删除应用失败：" + err);
  })
}
const handleEdit = (row: any) => {
  resetForm()
  openForm.value = true;
  keyValuePairs.value = []
  api.GetById(row.id).then(res => {
    if (res.code === 200) {
      form.value = res.data;
      const varsMap = JSON.parse(res.data.envVars);
      for (const [key, value] of Object.entries(varsMap)) {
        keyValuePairs.value.push({key, value});
      }
      //获取外部环境列表
      getExternalEnvList();
    } else {
      ElMessage.error(res.msg);
    }
  }).catch(err => {
    ElMessage.error("获取环境详情失败：" + err);
  })
}

const handleSubmit = () => {
  if (!form.value.name) {
    ElMessage.error("名称不能为空！");
    return;
  }
  if (!form.value.appDir) {
    ElMessage.error("目录不能为空！");
    return;
  }
  if (!form.value.startCmd) {
    ElMessage.error("启动命令不能为空！");
    return;
  }
  //数组转map
  const varMap = new Map<string, string>();
  for (const item of keyValuePairs.value) {
    varMap.set(item.key, item.value);
    if (item.key === '' || item.value === '') {
      ElMessage.error("键或值存在空值，请检查！");
      return;
    }
  }
  form.value.envVars = JSON.stringify(Object.fromEntries(varMap))
  console.log(form.value)
  if (form.value.id === "" || form.value.id === undefined) {
    api.Add(form.value).then(res => {
      if (res.code === 200) {
        ElMessage.success("新增应用成功！");
        tableData.value = res.data;
        listData();
        openForm.value = false;
      } else {
        ElMessage.error(res.msg);
      }
    }).catch(err => {
      ElMessage.error("新增应用失败：" + err);
    })
  } else {
    console.log(form.value)
    api.Update(form.value).then(res => {
      if (res.code === 200) {
        ElMessage.success("编辑应用成功！");
        openForm.value = false;
        tableData.value = res.data;
        listData();
      } else {
        ElMessage.error(res.msg);
      }
    }).catch(err => {
      ElMessage.error("编辑应用失败：" + err);
    })
  }
};
const handleCHeckEnv = (row: any) => {
  checkEnvForm.value = req.EnvCheckReq.createFrom()
  checkEnvFormVisible.value = true;
  handleCHackResult.value = "";
  form.value = row
};
const handleCheckExecute = () => {
  checkEnvForm.value.id = form.value.id

};

const handleSelectionChange = (val: any[]) => {
  multipleSelection.value = val
}

//指令命令-------------------
const handleLogInfo = (row: any) => {
  dialogAppLogRef.value!.acceptParams(row)
}
const handleStart = (row: any) => {
  row.runStatusNoLoading = false
  api.Start(row.id).then(res => {
    if (res.code === 200) {
      ElMessage.success("启动成功！");
    } else {
      ElMessage.error(res.msg);
    }
  }).catch(err => {
    ElMessage.error("启动失败：" + err);
  })

}
const handleStop = (row: any) => {
  row.runStatusNoLoading = false

  async function stop() {
    api.Stop(row.id).then(res => {
      if (res.code === 200) {
        ElMessage.success("停止成功！");
      } else {
        ElMessage.error(res.msg);
      }
    }).catch(err => {
      ElMessage.error("停止失败：" + err);
    })
  }

  stop()
}
const handleRestart = (row: any) => {
  row.runStatusNoLoading = false
  api.Restart(row.id).then(res => {
    if (res.code === 200) {
      ElMessage.success("重启成功！");
    } else {
      ElMessage.error(res.msg);
    }
    row.runStatusNoLoading = true
  }).catch(err => {
    ElMessage.error("重启失败：" + err);
    row.runStatusNoLoading = true
  })
}
const handleStartAll = () => {
  if (multipleSelection.value.length === 0) {
    ElMessage.error("请选择要启动的应用！");
    return;
  }
  tableDataLoading.value = true;
  const ids = multipleSelection.value.map(item => item.id);
  api.StartAll(ids).then(res => {
    if (res.code === 200) {
      ElMessage.success("启动成功！");
    } else {
      ElMessage.error(res.msg);
    }
    tableDataLoading.value = false;
  }).catch(err => {
    ElMessage.error("启动失败：" + err);
    tableDataLoading.value = false;
  })
}
const handleStopAll = () => {
  if (multipleSelection.value.length === 0) {
    ElMessage.error("请选择要停止的应用！");
    return;
  }
  tableDataLoading.value = true;
  const ids = multipleSelection.value.map(item => item.id);
  api.StopAll(ids).then(res => {
    if (res.code === 200) {
      ElMessage.success("停止成功！");
    } else {
      ElMessage.error(res.msg);
    }
    tableDataLoading.value = false;
  }).catch(err => {
    ElMessage.error("停止失败：" + err);
    tableDataLoading.value = false;
  })
}
let intervalId: any = null;

// 增加定时任务
function startTimer() {
  intervalId = setInterval(() => {
    getRunStatus();
  }, 5000);
}

// 清除定时任务
function stopTimer() {
  clearInterval(intervalId);
}

//打开文件夹
const openFileExplorer = (dir: any) => {
  openApi.OpenFileExplorer(dir).then(res => {
    if (res.code === 200) {
      ElMessage.success("打开目录成功！");
    } else {
      ElMessage.error(res.msg);
    }
  }).catch(err => {
    ElMessage.error("打开目录失败：" + err);
  })
}

//刷新端口
const handleRefreshPort = (row: any) => {
  api.RefreshPorts(row.id).then(res => {
    if (res.code === 200) {
      row.ports = res.data;
      row.portsArr = JSON.parse(row.ports)
      ElMessage.success("刷新端口成功！");
    } else {
      ElMessage.error(res.msg);
    }
  }).catch(err => {
    ElMessage.error("刷新端口失败：", err)
  })
}


</script>

<style scoped>
.run-status-container-item-class {
  margin-top: 15px;
}

</style>
