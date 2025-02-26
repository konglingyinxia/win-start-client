<template>
  <div class="app-container-header flex flex-wrap gap-4">
    <el-button type="primary" @Click="handleAdd" plain>新增</el-button>
    <el-button type="success" @Click="handleReload" plain>刷新</el-button>
    <div style="display: flex; align-items: center;">
      <label>环境根目录：
        <el-text type="warning"> {{ envHomePath.envHomeKey }}:</el-text>
        <el-text style="margin-left: 10px;" type="primary">{{ envHomePath.envHomePath }}</el-text>
        <el-button icon="FolderOpened" link
                   @click="openFileExplorer(envHomePath.envHomePath)"></el-button>

      </label>
    </div>
  </div>
    <el-table :data="tableData"
              style="width: 100%">
      <el-table-column label="序号" width="100" type="index" >
      </el-table-column>
      <el-table-column label="环境名称" prop="envName"></el-table-column>
      <el-table-column label="环境目录" prop="envDir">
        <template #default="scope">
          <div>
            <el-text>{{ scope.row.envDir }}</el-text>
          </div>
          <el-button icon="FolderOpened" type="primary" link
                     @click="openFileExplorer(envHomePath.envHomePath+envHomePath.separator+scope.row.envDir)"></el-button>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" prop="createdAt"></el-table-column>
      <el-table-column label="备注" prop="remark"></el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button type="text" @click="handleEdit(scope.row)">编辑</el-button>
          <el-popconfirm
              title="您确定要删除此内容吗?"
              @Confirm="handleDelete(scope.row)"
              cancel-button-text="否"
              confirm-button-text="是"
          >
            <template #reference>
             <el-button type="text">删除</el-button>
            </template>
          </el-popconfirm>
          <el-button type="warning" @click="handleCHeckEnv(scope.row)" link>验证</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!--新增环境弹窗    -->
    <el-dialog :title="(form.id === ''|| form.id === undefined) ? '新增环境' : '编辑环境'"
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
          <el-input v-model="envHomePath.envHomeKey" disabled>
            <template #prepend>键名：</template>
          </el-input>
          <el-input v-model="envHomePath.envHomePath" disabled>
            <template #prepend>键值：</template>
            <template #suffix>
              <el-button icon="FolderOpened" type="primary" link
                         @click="openFileExplorer(envHomePath.envHomePath)"></el-button>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="环境名称" prop="envName">
          <el-input v-model="form.envName"></el-input>
        </el-form-item>
        <el-form-item label="环境目录" prop="envDir">
          <el-input v-model="form.envDir"></el-input>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="form.remark"></el-input>
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
    <!--验证环境弹窗    -->
    <el-dialog title="验证环境"
               v-model="checkEnvFormVisible"
               draggable
               width="65%"
    >
      <el-form :model="checkEnvForm"
               :rules="checkEnvFormRules"
      >
        <el-form-item label="父级目录" prop="envHomePath">
          <el-input v-model="envHomePath.envHomeKey" disabled>
            <template #prepend>键名：</template>
          </el-input>
          <el-input v-model="envHomePath.envHomePath" disabled>
            <template #prepend>键值：</template>
          </el-input>
        </el-form-item>
        <el-form-item label="环境名称" prop="envName">
          <el-input v-model="form.envName" disabled></el-input>
        </el-form-item>
        <el-form-item label="环境ID" prop="id" hidden>
          <el-input v-model="form.id" disabled></el-input>
        </el-form-item>
        <el-form-item label="环境目录" prop="envDir">
          <el-input v-model="form.envDir" disabled></el-input>
        </el-form-item>
        <el-form-item label="执行命令" prop="command">
          <el-input v-model="checkEnvForm.command">
            <template #append>
              <el-button type="primary" @click="handleCheckExecute">执行</el-button>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item label="执行结果">
          <el-input type="textarea"
                    v-model="handleCHackResult"
                    rows="5"
                    disabled></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="checkEnvFormVisible = false">取消</el-button>
      </template>
    </el-dialog>
</template>
<script setup lang="ts">
import {onMounted, ref} from "vue";
import {api as comm, req} from "../../../../wailsjs/go/models"
import * as api from "../../../../wailsjs/go/api/EnvManager"
import type {FormRules} from 'element-plus'
import {ElMessage} from 'element-plus'
import * as openApi from "../../../../wailsjs/go/api/CommManager";

const openForm = ref(false);
const form = ref(req.EnvReqAdd.createFrom())
const envHomePath = ref(comm.EnvManager.createFrom())
const keyValuePairs = ref([])
// 验证环境弹窗
const checkEnvFormVisible = ref(false)
const checkEnvForm = ref(req.EnvCheckReq.createFrom())
const handleCHackResult = ref('')


const rules = ref<FormRules<req.EnvReqAdd>>({
  envName: [
    {required: true, message: '请输入环境名称', trigger: 'blur'},
    {min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur'}
  ],
  envDir: [
    {required: true, message: '请输入环境目录', trigger: 'blur'},
    {min: 2, max: 50, message: '长度在 2 到 50 个字符', trigger: 'blur'}
  ],
})
const checkEnvFormRules = ref<FormRules<req.EnvCheckReq>>({
  command: [
    {required: true, message: '请输入执行命令', trigger: 'blur'},
  ],
})

onMounted(() => {
      api.GetEnvHomePath().then(res => {
        envHomePath.value = res;
      });
  listData();
    }
)
const tableData = ref()
// 刷新环境列表
const handleReload = () => {
  listData();
  ElMessage.success("刷新成功！");
}

// 获取环境列表
const listData = () => {
  api.List().then(res => {
    tableData.value = res.data;
  })
}

const resetForm = () => {
  form.value = req.EnvReqAdd.createFrom()
  keyValuePairs.value = []
}

const handleAdd = () => {
  resetForm()
  //打开新增弹窗
  openForm.value = true;
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
const handleDelete = (row: any) => {
  api.Delete(row.id).then(res => {
    if (res.code === 200) {
      ElMessage.success("删除环境成功！");
      tableData.value = res.data;
      listData();
    } else {
      ElMessage.error(res.msg);
    }
  }).catch(err => {
    ElMessage.error("删除环境失败：" + err);
  })
}
const handleEdit = (row: any) => {
  resetForm()
  openForm.value = true;
  keyValuePairs.value = []
  api.Detail(row.id).then(res => {
    if (res.code === 200) {
      form.value = res.data;
      const varsMap = JSON.parse(res.data.envVars);
      for (const [key, value] of Object.entries(varsMap)) {
        keyValuePairs.value.push({key, value});
      }
    } else {
      ElMessage.error(res.msg);
    }
  }).catch(err => {
    ElMessage.error("获取环境详情失败：" + err);
  })
}

const handleSubmit = () => {
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
  if (form.value.id === "" || form.value.id === undefined) {
  api.Add(form.value).then(res => {
    if (res.code === 200) {
      ElMessage.success("新增环境成功！");
      openForm.value = false;
      tableData.value = res.data;
      listData();
    } else {
      ElMessage.error(res.msg);
    }
  }).catch(err => {
    ElMessage.error("新增环境失败：" + err);
  })
  } else {
    api.Update(form.value).then(res => {
      if (res.code === 200) {
        ElMessage.success("编辑环境成功！");
        openForm.value = false;
        tableData.value = res.data;
        listData();
      } else {
        ElMessage.error(res.msg);
      }
    }).catch(err => {
      ElMessage.error("编辑环境失败：" + err);
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
  api.CheckEnvVars(checkEnvForm.value).then(res => {
    if (res.code === 200) {
      handleCHackResult.value = res.data;
      ElMessage.success("验证环境成功！");
    } else {
      ElMessage.error(res.msg);
    }
  }).catch(err => {
    ElMessage.error("验证环境失败：" + err);
  })
};


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
</script>

<style scoped>

</style>
