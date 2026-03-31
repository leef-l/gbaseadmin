<script setup lang="ts">
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { onMounted, ref, watch } from 'vue';
import { Page, useVbenModal } from '@vben/common-ui';
import {
  Button,
  Card,
  Col,
  Empty,
  Image as AImage,
  message,
  Modal,
  Pagination,
  Row,
  Tag,
  Tooltip,
  Tree,
  Upload,
} from 'ant-design-vue';
import {
  AppstoreOutlined,
  CloudUploadOutlined,
  DeleteOutlined,
  EditOutlined,
  FileOutlined,
  FilePdfOutlined,
  FileWordOutlined,
  FileExcelOutlined,
  FileZipOutlined,
  UnorderedListOutlined,
  FolderOutlined,
} from '@ant-design/icons-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getFileList, deleteFile } from '#/api/upload/file';
import { getDirTree } from '#/api/upload/dir';
import { uploadFile } from '#/api/upload/uploader';
import type { FileItem } from '#/api/upload/file/types';
import type { DirItem } from '#/api/upload/dir/types';
import FormModal from './modules/form.vue';

/** 标签颜色池 */
const TAG_COLORS = ['green', 'red', 'blue', 'orange', 'cyan', 'purple', 'geekblue', 'magenta'];

/** 存储类型选项 */
const storageOptions = [
  { label: '本地', value: 1 },
  { label: '阿里云OSS', value: 2 },
  { label: '腾讯云COS', value: 3 },
];
const storageMap: Record<number, string> = { 1: '本地', 2: '阿里云OSS', 3: '腾讯云COS' };
function getStorageColor(val: number): string {
  const idx = [1, 2, 3].indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** 是否图片 */
const isImageOptions = [
  { label: '否', value: 0 },
  { label: '是', value: 1 },
];
const isImageMap: Record<number, string> = { 0: '否', 1: '是' };
function getIsImageColor(val: number): string {
  const idx = [0, 1].indexOf(val);
  return TAG_COLORS[idx >= 0 ? idx % TAG_COLORS.length : 0] ?? 'default';
}

/** 视图模式 */
const VIEW_MODE_KEY = 'upload_file_view_mode';
const viewMode = ref<'grid' | 'list'>(
  (localStorage.getItem(VIEW_MODE_KEY) as 'grid' | 'list') || 'list',
);
watch(viewMode, (v) => localStorage.setItem(VIEW_MODE_KEY, v));

/** 目录树 */
const dirTree = ref<DirItem[]>([]);
const selectedDirId = ref<string>('');
const treeLoading = ref(false);
const sideCollapsed = ref(false);

async function loadDirTree() {
  treeLoading.value = true;
  try {
    const list = await getDirTree();
    dirTree.value = [{ id: '', name: '全部文件', children: list } as DirItem];
  } catch {
    dirTree.value = [];
  } finally {
    treeLoading.value = false;
  }
}

function onDirSelect(keys: any[]) {
  selectedDirId.value = keys[0] ?? '';
  if (viewMode.value === 'list') {
    gridApi.reload();
  } else {
    gridPageNum.value = 1;
    loadGridData();
  }
}

/** 表单弹窗 */
const [FormModalComp, formModalApi] = useVbenModal({
  connectedComponent: FormModal,
  destroyOnClose: true,
});

/** 搜索关键词 */
const searchKeyword = ref('');

/** 上传中状态 */
const uploading = ref(false);

/** 自定义上传 */
async function handleCustomUpload(options: any) {
  uploading.value = true;
  try {
    await uploadFile(options.file, selectedDirId.value || undefined);
    message.success('上传成功');
    options.onSuccess?.();
    refreshCurrentView();
  } catch (e: any) {
    message.error(e?.message || '上传失败');
    options.onError?.(e);
  } finally {
    uploading.value = false;
  }
}

/** 搜索表单配置 */
const formOptions: VbenFormProps = {
  collapsed: false,
  showCollapseButton: true,
  submitOnChange: false,
  submitOnEnter: true,
  schema: [
    {
      component: 'Select',
      componentProps: { allowClear: true, options: storageOptions, placeholder: '请选择存储类型', class: 'w-full' },
      fieldName: 'storage',
      label: '存储类型',
    },
    {
      component: 'Select',
      componentProps: { allowClear: true, options: isImageOptions, placeholder: '请选择是否图片', class: 'w-full' },
      fieldName: 'isImage',
      label: '是否图片',
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<FileItem> = {
  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { field: 'dirName', title: '所属目录' },
    { field: 'name', title: '文件名称' },
    { field: 'url', title: '文件地址', minWidth: 200 },
    { field: 'ext', title: '扩展名', width: 80 },
    { field: 'size', title: '文件大小', width: 100 },
    { field: 'storage', title: '存储类型', width: 120, slots: { default: 'storage_cell' } },
    { field: 'isImage', title: '是否图片', width: 100, slots: { default: 'isImage_cell' } },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime' },
    { title: '操作', width: 160, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        const res = await getFileList({
          pageNum: page.currentPage,
          pageSize: page.pageSize,
          dirID: selectedDirId.value || undefined,
          name: searchKeyword.value || undefined,
          ...formValues,
        });
        return { items: res?.list ?? [], total: res?.total ?? 0 };
      },
    },
  },
  toolbarConfig: { custom: true, refresh: true, search: true },
};

const [Grid, gridApi] = useVbenVxeGrid({ formOptions, gridOptions });

/** Grid 视图数据 (区块模式) */
const gridDataList = ref<FileItem[]>([]);
const gridTotal = ref(0);
const gridPageNum = ref(1);
const gridPageSize = ref(20);
const gridLoading = ref(false);

async function loadGridData() {
  gridLoading.value = true;
  try {
    const res = await getFileList({
      pageNum: gridPageNum.value,
      pageSize: gridPageSize.value,
      dirID: selectedDirId.value || undefined,
      name: searchKeyword.value || undefined,
    });
    gridDataList.value = res?.list ?? [];
    gridTotal.value = res?.total ?? 0;
  } catch {
    gridDataList.value = [];
    gridTotal.value = 0;
  } finally {
    gridLoading.value = false;
  }
}

function onGridPageChange(page: number, pageSize: number) {
  gridPageNum.value = page;
  gridPageSize.value = pageSize;
  loadGridData();
}

/** 刷新当前视图 */
function refreshCurrentView() {
  if (viewMode.value === 'list') {
    gridApi.reload();
  } else {
    loadGridData();
  }
}

/** 搜索 */
function handleSearch() {
  if (viewMode.value === 'list') {
    gridApi.reload();
  } else {
    gridPageNum.value = 1;
    loadGridData();
  }
}

/** 切换视图 */
function switchView(mode: 'grid' | 'list') {
  viewMode.value = mode;
  if (mode === 'grid') {
    gridPageNum.value = 1;
    loadGridData();
  }
}

/** 新建 */
function handleCreate() {
  formModalApi.setData(null).open();
}

/** 编辑 */
function handleEdit(row: FileItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: FileItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该文件记录吗？',
    okType: 'danger',
    async onOk() {
      await deleteFile(row.id);
      message.success('删除成功');
      refreshCurrentView();
    },
  });
}

/** 图片预览 */
const previewVisible = ref(false);
const previewUrl = ref('');
function handlePreview(url: string) {
  previewUrl.value = url;
  previewVisible.value = true;
}

/** 文件图标 */
const EXT_ICON_MAP: Record<string, any> = {
  pdf: FilePdfOutlined,
  doc: FileWordOutlined,
  docx: FileWordOutlined,
  xls: FileExcelOutlined,
  xlsx: FileExcelOutlined,
  zip: FileZipOutlined,
  rar: FileZipOutlined,
  '7z': FileZipOutlined,
};
const IMAGE_EXTS = new Set(['jpg', 'jpeg', 'png', 'gif', 'bmp', 'webp', 'svg']);

function isImage(file: FileItem): boolean {
  if (file.isImage === 1) return true;
  return IMAGE_EXTS.has((file.ext || '').toLowerCase().replace('.', ''));
}

function getFileIcon(file: FileItem) {
  const ext = (file.ext || '').toLowerCase().replace('.', '');
  return EXT_ICON_MAP[ext] || FileOutlined;
}

/** 格式化文件大小 */
function formatSize(size?: string): string {
  if (!size) return '-';
  const n = Number(size);
  if (Number.isNaN(n)) return size;
  if (n < 1024) return `${n} B`;
  if (n < 1024 * 1024) return `${(n / 1024).toFixed(1)} KB`;
  return `${(n / (1024 * 1024)).toFixed(1)} MB`;
}

onMounted(() => {
  loadDirTree();
  if (viewMode.value === 'grid') {
    loadGridData();
  }
});
</script>

<template>
  <Page auto-content-height>
    <FormModalComp @success="refreshCurrentView" />
    <div class="flex h-full gap-3">
      <!-- 左侧目录树 -->
      <div
        v-show="!sideCollapsed"
        class="flex-shrink-0 w-[240px] border border-solid border-gray-200 rounded-md bg-white dark:bg-gray-800 dark:border-gray-700 overflow-auto"
      >
        <div class="flex items-center justify-between px-3 py-2 border-b border-solid border-gray-200 dark:border-gray-700">
          <span class="font-medium text-sm">目录</span>
          <Button type="text" size="small" @click="sideCollapsed = true">
            收起
          </Button>
        </div>
        <Tree
          :tree-data="dirTree"
          :field-names="{ title: 'name', key: 'id', children: 'children' }"
          default-expand-all
          :selected-keys="[selectedDirId]"
          @select="onDirSelect"
          class="px-1 py-2"
        />
      </div>
      <Button
        v-if="sideCollapsed"
        type="text"
        class="flex-shrink-0 self-start mt-1"
        @click="sideCollapsed = false"
      >
        <FolderOutlined />
      </Button>

      <!-- 右侧内容区 -->
      <div class="flex-1 flex flex-col min-w-0 overflow-hidden">
        <!-- 工具栏 -->
        <div class="flex items-center gap-2 mb-3 flex-shrink-0">
          <Upload
            :custom-request="handleCustomUpload"
            :show-upload-list="false"
            :multiple="true"
          >
            <Button type="primary" :loading="uploading">
              <CloudUploadOutlined />
              上传文件
            </Button>
          </Upload>
          <Button @click="handleCreate">新建</Button>
          <input
            v-model="searchKeyword"
            class="ant-input px-2 py-1 w-[200px]"
            placeholder="搜索文件名..."
            @keyup.enter="handleSearch"
          />
          <Button @click="handleSearch">搜索</Button>
          <div class="ml-auto flex gap-1">
            <Tooltip title="列表视图">
              <Button
                :type="viewMode === 'list' ? 'primary' : 'default'"
                @click="switchView('list')"
              >
                <UnorderedListOutlined />
              </Button>
            </Tooltip>
            <Tooltip title="区块视图">
              <Button
                :type="viewMode === 'grid' ? 'primary' : 'default'"
                @click="switchView('grid')"
              >
                <AppstoreOutlined />
              </Button>
            </Tooltip>
          </div>
        </div>

        <!-- 列表视图 -->
        <div v-show="viewMode === 'list'" class="flex-1 overflow-hidden">
          <Grid>
            <template #toolbar-actions>
              <span />
            </template>
            <template #storage_cell="{ row }">
              <Tag :color="getStorageColor(row.storage)">
                {{ storageMap[row.storage] || row.storage }}
              </Tag>
            </template>
            <template #isImage_cell="{ row }">
              <Tag :color="getIsImageColor(row.isImage)">
                {{ isImageMap[row.isImage] || row.isImage }}
              </Tag>
            </template>
            <template #action="{ row }">
              <Button type="link" size="small" @click="handleEdit(row)">编辑</Button>
              <Button type="link" danger size="small" @click="handleDelete(row)">删除</Button>
            </template>
          </Grid>
        </div>

        <!-- 区块视图 -->
        <div v-show="viewMode === 'grid'" class="flex-1 overflow-auto">
          <div v-if="gridLoading" class="flex items-center justify-center py-20">
            加载中...
          </div>
          <Empty v-else-if="gridDataList.length === 0" description="暂无文件" class="py-20" />
          <Row v-else :gutter="[16, 16]">
            <Col
              v-for="file in gridDataList"
              :key="file.id"
              :xs="12" :sm="8" :md="6" :lg="4" :xl="4"
            >
              <Card
                hoverable
                size="small"
                class="h-full"
                :body-style="{ padding: '8px' }"
              >
                <!-- 缩略图 / 图标 -->
                <div
                  class="flex items-center justify-center h-[120px] bg-gray-50 dark:bg-gray-700 rounded mb-2 overflow-hidden cursor-pointer"
                  @click="isImage(file) ? handlePreview(file.url) : undefined"
                >
                  <img
                    v-if="isImage(file)"
                    :src="file.url"
                    :alt="file.name"
                    class="max-h-full max-w-full object-contain"
                  />
                  <component
                    v-else
                    :is="getFileIcon(file)"
                    class="text-4xl text-gray-400"
                  />
                </div>
                <!-- 文件信息 -->
                <Tooltip :title="file.name">
                  <div class="text-sm truncate font-medium">{{ file.name }}</div>
                </Tooltip>
                <div class="text-xs text-gray-400 mt-1 flex justify-between">
                  <span>{{ formatSize(file.size) }}</span>
                  <span>{{ file.createdAt?.slice(0, 10) || '' }}</span>
                </div>
                <!-- 操作 -->
                <div class="flex justify-end gap-1 mt-2">
                  <Button type="text" size="small" @click="handleEdit(file)">
                    <EditOutlined />
                  </Button>
                  <Button type="text" danger size="small" @click="handleDelete(file)">
                    <DeleteOutlined />
                  </Button>
                </div>
              </Card>
            </Col>
          </Row>
          <!-- 分页 -->
          <div v-if="gridTotal > 0" class="flex justify-end mt-4">
            <Pagination
              :current="gridPageNum"
              :page-size="gridPageSize"
              :total="gridTotal"
              show-size-changer
              show-quick-jumper
              :show-total="(total: number) => `共 ${total} 条`"
              @change="onGridPageChange"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- 图片预览 -->
    <AImage
      :style="{ display: 'none' }"
      :preview="{
        visible: previewVisible,
        onVisibleChange: (v: boolean) => { previewVisible = v; },
      }"
      :src="previewUrl"
    />
  </Page>
</template>
