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
{{- if .HasParentID}}
import { get{{.ModelName}}Tree, delete{{.ModelName}} } from '#/api/system/{{.ModuleName}}';
{{- else}}
import { get{{.ModelName}}List, delete{{.ModelName}} } from '#/api/system/{{.ModuleName}}';
{{- end}}
import type { {{.ModelName}}Item } from '#/api/system/{{.ModuleName}}/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];
{{range .Fields}}
{{- if and (not .IsHidden) (.IsEnum)}}
/** {{.Label}}选项 */
const {{.NameLower}}Options = [
{{- range .EnumValues}}
  { label: '{{.Label}}', value: {{.Value}} },
{{- end}}
];

/** {{.Label}}映射 */
const {{.NameLower}}Map: Record<number, string> = {
{{- range .EnumValues}}
  {{.Value}}: '{{.Label}}',
{{- end}}
};

/** {{.Label}}颜色 */
function get{{.NameCamel}}Color(val: number): string {
  const keys = [{{range $i, $v := .EnumValues}}{{if $i}}, {{end}}{{$v.Value}}{{end}}];
  const idx = keys.indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}
{{end}}
{{- end}}
const loading = ref(false);
const dataList = ref<{{.ModelName}}Item[]>([]);
{{- if not .HasParentID}}
const total = ref(0);
{{- end}}
const formRef = ref();

const queryParams = reactive({
{{- if not .HasParentID}}
  pageNum: 1,
  pageSize: 10,
{{- end}}
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID) (.IsEnum)}}
  {{.NameLower}}: undefined as number | undefined,
{{- end}}
{{- end}}
});

/** 列定义 */
const columns = [
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID) (not .IsParentID) (not .IsTimeField) (not .IsMultiFK)}}
  { title: '{{.Label}}', dataIndex: '{{.NameLower}}', key: '{{.NameLower}}'{{if .IsEnum}}, width: 120{{end}} },
{{- end}}
{{- end}}
{{- range .Fields}}
{{- if and (not .IsHidden) (.IsTimeField)}}
  { title: '{{.Label}}', dataIndex: '{{.NameLower}}', key: '{{.NameLower}}', width: 180 },
{{- end}}
{{- end}}
  { title: '创建时间', dataIndex: 'createdAt', key: 'createdAt', width: 180 },
  { title: '操作', key: 'action', width: 200, fixed: 'right' as const },
];

/** 加载数据 */
async function loadData() {
  loading.value = true;
  try {
{{- if .HasParentID}}
    const params: Record<string, any> = {};
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID) (not .IsParentID) (.IsEnum)}}
    if (queryParams.{{.NameLower}} !== undefined) {
      params.{{.NameLower}} = queryParams.{{.NameLower}};
    }
{{- end}}
{{- end}}
    const res = await get{{.ModelName}}Tree(params);
    dataList.value = res ?? [];
{{- else}}
    const res = await get{{.ModelName}}List({
      pageNum: queryParams.pageNum,
      pageSize: queryParams.pageSize,
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID) (.IsEnum)}}
      {{.NameLower}}: queryParams.{{.NameLower}},
{{- end}}
{{- end}}
    });
    dataList.value = res?.list ?? [];
    total.value = res?.total ?? 0;
{{- end}}
  } finally {
    loading.value = false;
  }
}

/** 搜索 */
function handleSearch() {
{{- if not .HasParentID}}
  queryParams.pageNum = 1;
{{- end}}
  loadData();
}

/** 重置 */
function handleReset() {
{{- if not .HasParentID}}
  queryParams.pageNum = 1;
{{- end}}
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID) (.IsEnum)}}
  queryParams.{{.NameLower}} = undefined;
{{- end}}
{{- end}}
  loadData();
}

/** 新建 */
function handleCreate() {
  formRef.value?.open();
}

/** 编辑 */
function handleEdit(record: {{.ModelName}}Item) {
  formRef.value?.open(record.id);
}

/** 删除 */
function handleDelete(record: {{.ModelName}}Item) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该{{.Comment}}吗？',
    okType: 'danger',
    async onOk() {
      await delete{{.ModelName}}(record.id);
      message.success('删除成功');
      loadData();
    },
  });
}
{{- if not .HasParentID}}

/** 分页变化 */
function handlePageChange(page: number, pageSize: number) {
  queryParams.pageNum = page;
  queryParams.pageSize = pageSize;
  loadData();
}
{{- end}}

onMounted(() => {
  loadData();
});
</script>

<template>
  <div class="p-4">
    <!-- 搜索栏 -->
    <div class="mb-4 flex items-center gap-3">
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID) (.IsEnum)}}
      <a-select
        v-model:value="queryParams.{{.NameLower}}"
        :options="{{.NameLower}}Options"
        placeholder="{{.Label}}"
        allow-clear
        style="width: 160px"
      />
{{- end}}
{{- end}}
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
{{- if .HasParentID}}
      :children-column-name="'children'"
      :pagination="false"
      default-expand-all-rows
{{- else}}
      :pagination="{
        current: queryParams.pageNum,
        pageSize: queryParams.pageSize,
        total: total,
        showSizeChanger: true,
        showQuickJumper: true,
        showTotal: (t: number) => `共 ${t} 条`,
        onChange: handlePageChange,
      }"
{{- end}}
      :scroll="{ x: 'max-content' }"
    >
      <template #bodyCell="{ column, record }">
{{- range .Fields}}
{{- if and (not .IsHidden) (not .IsID) (not .IsParentID) (not .IsTimeField) (not .IsMultiFK) (.IsEnum)}}
        <template v-if="column.key === '{{.NameLower}}'">
          <a-tag :color="get{{.NameCamel}}Color(record.{{.NameLower}})">
            {{"{{"}} {{.NameLower}}Map[record.{{.NameLower}}] || record.{{.NameLower}} {{"}}"}}
          </a-tag>
        </template>
{{- end}}
{{- end}}
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
