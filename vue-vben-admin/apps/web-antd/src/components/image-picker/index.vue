<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue';
import {
  Row,
  Col,
  Card,
  Checkbox,
  Input,
  Tree,
  Pagination,
  Button,
  Empty,
  Spin,
} from 'ant-design-vue';
import { getFileList } from '#/api/upload/file';
import { getDirTree } from '#/api/upload/dir';
import type { FileItem } from '#/api/upload/file/types';
import type { DirItem } from '#/api/upload/dir/types';

interface Props {
  multiple?: boolean;
  maxCount?: number;
  accept?: string;
}

const props = withDefaults(defineProps<Props>(), {
  multiple: false,
  maxCount: 1,
  accept: 'image',
});

const emit = defineEmits<{
  confirm: [images: { id: string; url: string; name: string; size: number; ext: string }[]];
}>();

/** State */
const loading = ref(false);
const treeLoading = ref(false);
const keyword = ref('');
const selectedDirId = ref<string | undefined>(undefined);
const fileList = ref<FileItem[]>([]);
const selectedIds = ref<Set<string>>(new Set());
const dirTreeData = ref<DirItem[]>([]);
const pagination = reactive({ current: 1, pageSize: 20, total: 0 });

const effectiveMax = computed(() => (props.multiple ? props.maxCount : 1));

const treeNodes = computed(() => {
  const allNode: any = { id: undefined, name: '全部文件', children: dirTreeData.value };
  return [allNode];
});

/** Load directory tree */
async function loadDirTree() {
  treeLoading.value = true;
  try {
    dirTreeData.value = await getDirTree();
  } finally {
    treeLoading.value = false;
  }
}

/** Load file list */
async function loadFileList() {
  loading.value = true;
  try {
    const res = await getFileList({
      pageNum: pagination.current,
      pageSize: pagination.pageSize,
      isImage: 1,
      dirID: selectedDirId.value,
      name: keyword.value || undefined,
    });
    fileList.value = res?.list ?? [];
    pagination.total = res?.total ?? 0;
  } finally {
    loading.value = false;
  }
}

/** Toggle image selection */
function toggleSelect(file: FileItem) {
  const ids = selectedIds.value;
  if (ids.has(file.id)) {
    ids.delete(file.id);
  } else {
    if (!props.multiple) {
      ids.clear();
    }
    if (ids.size < effectiveMax.value) {
      ids.add(file.id);
    }
  }
}

/** Directory select */
function onDirSelect(keys: (string | undefined)[]) {
  selectedDirId.value = keys[0] ?? undefined;
  pagination.current = 1;
  loadFileList();
}

/** Search */
function handleSearch() {
  pagination.current = 1;
  loadFileList();
}

/** Pagination */
function onPageChange(page: number) {
  pagination.current = page;
  loadFileList();
}

function onPageSizeChange(_current: number, size: number) {
  pagination.pageSize = size;
  pagination.current = 1;
  loadFileList();
}

/** Confirm selection */
function handleConfirm() {
  const selected = fileList.value
    .filter((f) => selectedIds.value.has(f.id))
    .map((f) => ({
      id: f.id,
      url: f.url,
      name: f.name,
      size: Number(f.size) || 0,
      ext: f.ext ?? '',
    }));
  emit('confirm', selected);
}

/** Reset state (called from parent) */
function reset() {
  selectedIds.value.clear();
  keyword.value = '';
  selectedDirId.value = undefined;
  pagination.current = 1;
}

defineExpose({ reset, loadFileList });

onMounted(() => {
  loadDirTree();
  loadFileList();
});
</script>

<template>
  <div class="image-picker">
    <!-- Top bar -->
    <div class="image-picker__header">
      <Input.Search
        v-model:value="keyword"
        placeholder="搜索文件名"
        style="width: 240px"
        allow-clear
        @search="handleSearch"
      />
      <span class="image-picker__count">
        已选 {{ selectedIds.size }} / {{ effectiveMax }} 项
      </span>
      <Button type="primary" :disabled="selectedIds.size === 0" @click="handleConfirm">
        确认选择
      </Button>
    </div>

    <div class="image-picker__body">
      <!-- Left sidebar: directory tree -->
      <div class="image-picker__sidebar">
        <Spin :spinning="treeLoading">
          <Tree
            :tree-data="treeNodes"
            :field-names="{ title: 'name', key: 'id', children: 'children' }"
            default-expand-all
            block-node
            @select="onDirSelect"
          />
        </Spin>
      </div>

      <!-- Right area: image grid -->
      <div class="image-picker__content">
        <Spin :spinning="loading">
          <template v-if="fileList.length > 0">
            <Row :gutter="[12, 12]">
              <Col v-for="file in fileList" :key="file.id" :span="4">
                <Card
                  hoverable
                  class="image-picker__card"
                  :class="{ 'image-picker__card--selected': selectedIds.has(file.id) }"
                  :body-style="{ padding: '8px' }"
                  @click="toggleSelect(file)"
                >
                  <div class="image-picker__checkbox">
                    <Checkbox :checked="selectedIds.has(file.id)" />
                  </div>
                  <div class="image-picker__thumb">
                    <img :src="file.url" :alt="file.name" />
                  </div>
                  <div class="image-picker__name" :title="file.name">{{ file.name }}</div>
                </Card>
              </Col>
            </Row>
          </template>
          <Empty v-else description="暂无图片" />
        </Spin>

        <div v-if="pagination.total > 0" class="image-picker__pagination">
          <Pagination
            v-model:current="pagination.current"
            :page-size="pagination.pageSize"
            :total="pagination.total"
            size="small"
            show-size-changer
            :page-size-options="['20', '40', '60']"
            @change="onPageChange"
            @show-size-change="onPageSizeChange"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.image-picker {
  display: flex;
  flex-direction: column;
  height: 520px;
}

.image-picker__header {
  display: flex;
  align-items: center;
  gap: 12px;
  padding-bottom: 12px;
  border-bottom: 1px solid #f0f0f0;
}

.image-picker__count {
  flex: 1;
  text-align: right;
  color: #666;
  font-size: 13px;
}

.image-picker__body {
  display: flex;
  flex: 1;
  gap: 12px;
  padding-top: 12px;
  overflow: hidden;
}

.image-picker__sidebar {
  width: 200px;
  min-width: 200px;
  overflow-y: auto;
  border-right: 1px solid #f0f0f0;
  padding-right: 12px;
}

.image-picker__content {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow-y: auto;
}

.image-picker__card {
  position: relative;
  cursor: pointer;
}

.image-picker__card--selected {
  border-color: #1677ff;
  box-shadow: 0 0 0 2px rgba(22, 119, 255, 0.2);
}

.image-picker__checkbox {
  position: absolute;
  top: 4px;
  left: 4px;
  z-index: 1;
}

.image-picker__thumb {
  width: 100%;
  aspect-ratio: 1;
  overflow: hidden;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #fafafa;
}

.image-picker__thumb img {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
}

.image-picker__name {
  margin-top: 4px;
  font-size: 12px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  text-align: center;
}

.image-picker__pagination {
  display: flex;
  justify-content: flex-end;
  padding-top: 12px;
  margin-top: auto;
}
</style>
