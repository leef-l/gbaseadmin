<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue';
import { message, Modal } from 'ant-design-vue';
import {
  PlusOutlined,
  EditOutlined,
  DeleteOutlined,
  SearchOutlined,
  ReloadOutlined,
} from '@ant-design/icons-vue';
import { getRoleTree, deleteRole } from '#/api/system/role';
import type { RoleItem } from '#/api/system/role/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 数据范围选项 */
const dataScopeOptions = [
  { label: '全部', value: 1 },
  { label: '本部门及以下', value: 2 },
  { label: '本部门', value: 3 },
  { label: '仅本人', value: 4 },
  { label: '自定义', value: 5 },
];

/** 数据范围映射 */
const dataScopeMap: Record<number, string> = {
  1: '全部',
  2: '本部门及以下',
  3: '本部门',
  4: '仅本人',
  5: '自定义',
};

/** 数据范围颜色 */
function getDataScopeColor(val: number): string {
  const keys = [1, 2, 3, 4, 5];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** 状态选项 */
const statusOptions = [
  { label: '关闭', value: 0 },
  { label: '开启', value: 1 },
];

/** 状态映射 */
const statusMap: Record<number, string> = {
  0: '关闭',
  1: '开启',
};

/** 状态颜色 */
function getStatusColor(val: number): string {
  const keys = [0, 1];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

const loading = ref(false);
const dataList = ref<RoleItem[]>([]);
const formRef = ref();

const queryParams = reactive({
  dataScope: undefined as number | undefined,
  status: undefined as number | undefined,
});

/** 列定义 */
const columns = [
  { title: '角色名称', dataIndex: 'title', key: 'title' },
  { title: '数据范围', dataIndex: 'dataScope', key: 'dataScope', width: 120 },
  { title: '排序（升序）', dataIndex: 'sort', key: 'sort' },
  { title: '状态', dataIndex: 'status', key: 'status', width: 120 },
  { title: '创建时间', dataIndex: 'createdAt', key: 'createdAt', width: 180 },
  { title: '操作', key: 'action', width: 200, fixed: 'right' as const },
];

/** 加载数据 */
async function loadData() {
  loading.value = true;
  try {
    const params: Record<string, any> = {};
    if (queryParams.dataScope !== undefined) {
      params.dataScope = queryParams.dataScope;
    }
    if (queryParams.status !== undefined) {
      params.status = queryParams.status;
    }
    const res = await getRoleTree(params);
    dataList.value = res ?? [];
  } finally {
    loading.value = false;
  }
}

/** 搜索 */
function handleSearch() {
  loadData();
}

/** 重置 */
function handleReset() {
  queryParams.dataScope = undefined;
  queryParams.status = undefined;
  loadData();
}

/** 新建 */
function handleCreate() {
  formRef.value?.open();
}

/** 编辑 */
function handleEdit(record: RoleItem) {
  formRef.value?.open(record.id);
}

/** 删除 */
function handleDelete(record: RoleItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该角色表吗？',
    okType: 'danger',
    async onOk() {
      await deleteRole(record.id);
      message.success('删除成功');
      loadData();
    },
  });
}

onMounted(() => {
  loadData();
});
</script>

<template>
  <div class="p-4">
    <!-- 搜索栏 -->
    <div class="mb-4 flex items-center gap-3">
      <a-select
        v-model:value="queryParams.dataScope"
        :options="dataScopeOptions"
        placeholder="数据范围"
        allow-clear
        style="width: 160px"
      />
      <a-select
        v-model:value="queryParams.status"
        :options="statusOptions"
        placeholder="状态"
        allow-clear
        style="width: 160px"
      />
      <a-button type="primary" @click="handleSearch">
        <template #icon><SearchOutlined /></template>
        搜索
      </a-button>
      <a-button @click="handleReset">
        <template #icon><ReloadOutlined /></template>
        重置
      </a-button>
      <div class="flex-1" />
      <a-button type="primary" @click="handleCreate">
        <template #icon><PlusOutlined /></template>
        新建
      </a-button>
    </div>

    <!-- 数据表格 -->
    <a-table
      :columns="columns"
      :data-source="dataList"
      :loading="loading"
      row-key="id"
      :children-column-name="'children'"
      :pagination="false"
      default-expand-all-rows
      :scroll="{ x: 'max-content' }"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'dataScope'">
          <a-tag :color="getDataScopeColor(record.dataScope)">
            {{ dataScopeMap[record.dataScope] || record.dataScope }}
          </a-tag>
        </template>
        <template v-if="column.key === 'status'">
          <a-tag :color="getStatusColor(record.status)">
            {{ statusMap[record.status] || record.status }}
          </a-tag>
        </template>
        <template v-if="column.key === 'action'">
          <div class="flex gap-2">
            <a-button type="link" size="small" @click="handleEdit(record)">
              <template #icon><EditOutlined /></template>
              编辑
            </a-button>
            <a-button type="link" danger size="small" @click="handleDelete(record)">
              <template #icon><DeleteOutlined /></template>
              删除
            </a-button>
          </div>
        </template>
      </template>
    </a-table>

    <!-- 表单弹窗 -->
    <FormModal ref="formRef" @success="loadData" />
  </div>
</template>
