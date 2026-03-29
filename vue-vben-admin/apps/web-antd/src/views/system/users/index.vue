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
import { getUsersList, deleteUsers } from '#/api/system/users';
import type { UsersItem } from '#/api/system/users/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

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
const dataList = ref<UsersItem[]>([]);
const total = ref(0);
const formRef = ref();

const queryParams = reactive({
  pageNum: 1,
  pageSize: 10,
  status: undefined as number | undefined,
});

/** 列定义 */
const columns = [
  { title: '登录用户名', dataIndex: 'username', key: 'username' },
  { title: '昵称/显示名', dataIndex: 'nickname', key: 'nickname' },
  { title: '邮箱地址', dataIndex: 'email', key: 'email' },
  { title: '头像图片 URL', dataIndex: 'avatar', key: 'avatar' },
  { title: '状态', dataIndex: 'status', key: 'status', width: 120 },
  { title: '创建时间', dataIndex: 'createdAt', key: 'createdAt', width: 180 },
  { title: '操作', key: 'action', width: 200, fixed: 'right' as const },
];

/** 加载数据 */
async function loadData() {
  loading.value = true;
  try {
    const res = await getUsersList({
      pageNum: queryParams.pageNum,
      pageSize: queryParams.pageSize,
      status: queryParams.status,
    });
    dataList.value = res?.list ?? [];
    total.value = res?.total ?? 0;
  } finally {
    loading.value = false;
  }
}

/** 搜索 */
function handleSearch() {
  queryParams.pageNum = 1;
  loadData();
}

/** 重置 */
function handleReset() {
  queryParams.pageNum = 1;
  queryParams.status = undefined;
  loadData();
}

/** 新建 */
function handleCreate() {
  formRef.value?.open();
}

/** 编辑 */
function handleEdit(record: UsersItem) {
  formRef.value?.open(record.id);
}

/** 删除 */
function handleDelete(record: UsersItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该用户表吗？',
    okType: 'danger',
    async onOk() {
      await deleteUsers(record.id);
      message.success('删除成功');
      loadData();
    },
  });
}

/** 分页变化 */
function handlePageChange(page: number, pageSize: number) {
  queryParams.pageNum = page;
  queryParams.pageSize = pageSize;
  loadData();
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
      :pagination="{
        current: queryParams.pageNum,
        pageSize: queryParams.pageSize,
        total: total,
        showSizeChanger: true,
        showQuickJumper: true,
        showTotal: (t: number) => `共 ${t} 条`,
        onChange: handlePageChange,
      }"
      :scroll="{ x: 'max-content' }"
    >
      <template #bodyCell="{ column, record }">
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
