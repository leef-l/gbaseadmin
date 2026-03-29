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
import { getMenuTree, deleteMenu } from '#/api/system/menu';
import type { MenuItem } from '#/api/system/menu/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 类型选项 */
const typeOptions = [
  { label: '目录', value: 1 },
  { label: '菜单', value: 2 },
  { label: '按钮', value: 3 },
  { label: '外链', value: 4 },
  { label: '内链', value: 5 },
];

/** 类型映射 */
const typeMap: Record<number, string> = {
  1: '目录',
  2: '菜单',
  3: '按钮',
  4: '外链',
  5: '内链',
};

/** 类型颜色 */
function getTypeColor(val: number): string {
  const keys = [1, 2, 3, 4, 5];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** 是否显示选项 */
const isShowOptions = [
  { label: '隐藏', value: 0 },
  { label: '显示', value: 1 },
];

/** 是否显示映射 */
const isShowMap: Record<number, string> = {
  0: '隐藏',
  1: '显示',
};

/** 是否显示颜色 */
function getIsShowColor(val: number): string {
  const keys = [0, 1];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** 是否缓存选项 */
const isCacheOptions = [
  { label: '不缓存', value: 0 },
  { label: '缓存', value: 1 },
];

/** 是否缓存映射 */
const isCacheMap: Record<number, string> = {
  0: '不缓存',
  1: '缓存',
};

/** 是否缓存颜色 */
function getIsCacheColor(val: number): string {
  const keys = [0, 1];
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
const dataList = ref<MenuItem[]>([]);
const formRef = ref();

const queryParams = reactive({
  type: undefined as number | undefined,
  isShow: undefined as number | undefined,
  isCache: undefined as number | undefined,
  status: undefined as number | undefined,
});

/** 列定义 */
const columns = [
  { title: '菜单名称', dataIndex: 'title', key: 'title' },
  { title: '类型', dataIndex: 'type', key: 'type', width: 120 },
  { title: '前端路由路径', dataIndex: 'path', key: 'path' },
  { title: '前端组件路径', dataIndex: 'component', key: 'component' },
  { title: '权限标识（如 system', dataIndex: 'permission', key: 'permission' },
  { title: '菜单图标（图标名称）', dataIndex: 'icon', key: 'icon' },
  { title: '排序（升序）', dataIndex: 'sort', key: 'sort' },
  { title: '是否显示', dataIndex: 'isShow', key: 'isShow', width: 120 },
  { title: '是否缓存', dataIndex: 'isCache', key: 'isCache', width: 120 },
  { title: '外链/内链地址（type=4或5时有效）', dataIndex: 'linkURL', key: 'linkURL' },
  { title: '状态', dataIndex: 'status', key: 'status', width: 120 },
  { title: '创建时间', dataIndex: 'createdAt', key: 'createdAt', width: 180 },
  { title: '操作', key: 'action', width: 200, fixed: 'right' as const },
];

/** 加载数据 */
async function loadData() {
  loading.value = true;
  try {
    const params: Record<string, any> = {};
    if (queryParams.type !== undefined) {
      params.type = queryParams.type;
    }
    if (queryParams.isShow !== undefined) {
      params.isShow = queryParams.isShow;
    }
    if (queryParams.isCache !== undefined) {
      params.isCache = queryParams.isCache;
    }
    if (queryParams.status !== undefined) {
      params.status = queryParams.status;
    }
    const res = await getMenuTree(params);
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
  queryParams.type = undefined;
  queryParams.isShow = undefined;
  queryParams.isCache = undefined;
  queryParams.status = undefined;
  loadData();
}

/** 新建 */
function handleCreate() {
  formRef.value?.open();
}

/** 编辑 */
function handleEdit(record: MenuItem) {
  formRef.value?.open(record.id);
}

/** 删除 */
function handleDelete(record: MenuItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该菜单表吗？',
    okType: 'danger',
    async onOk() {
      await deleteMenu(record.id);
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
        v-model:value="queryParams.type"
        :options="typeOptions"
        placeholder="类型"
        allow-clear
        style="width: 160px"
      />
      <a-select
        v-model:value="queryParams.isShow"
        :options="isShowOptions"
        placeholder="是否显示"
        allow-clear
        style="width: 160px"
      />
      <a-select
        v-model:value="queryParams.isCache"
        :options="isCacheOptions"
        placeholder="是否缓存"
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
        <template v-if="column.key === 'type'">
          <a-tag :color="getTypeColor(record.type)">
            {{ typeMap[record.type] || record.type }}
          </a-tag>
        </template>
        <template v-if="column.key === 'isShow'">
          <a-tag :color="getIsShowColor(record.isShow)">
            {{ isShowMap[record.isShow] || record.isShow }}
          </a-tag>
        </template>
        <template v-if="column.key === 'isCache'">
          <a-tag :color="getIsCacheColor(record.isCache)">
            {{ isCacheMap[record.isCache] || record.isCache }}
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
