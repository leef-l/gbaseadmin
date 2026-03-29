<script setup lang="ts">
import { ref, reactive, nextTick } from 'vue';
import { message } from 'ant-design-vue';
import type { FormInstance } from 'ant-design-vue';
import {
  getMenuDetail,
  createMenu,
  updateMenu,
} from '#/api/system/menu';
import { getMenuTree } from '#/api/system/menu';
import type { MenuItem } from '#/api/system/menu/types';

const emit = defineEmits<{ success: [] }>();

const visible = ref(false);
const confirmLoading = ref(false);
const isEdit = ref(false);
const editId = ref('');
const formRef = ref<FormInstance>();
const treeData = ref<MenuItem[]>([]);

/** 表单初始值 */
const initialForm = {
  parentID: '' as string,
  title: '',
  type: 1 as number | undefined,
  path: '',
  component: '',
  permission: '',
  icon: '',
  sort: 0,
  isShow: 1,
  isCache: 0,
  linkURL: '',
  status: 1,
};

const formData = reactive({ ...initialForm });

/** 校验规则 */
const rules = {
  title: [
    {
      required: true,
      message: '请输入菜单名称',
      trigger: 'blur',
    },
  ],
};

/** 类型选项 */
const typeOptions = [
  { label: '目录', value: 1 },
  { label: '菜单', value: 2 },
  { label: '按钮', value: 3 },
  { label: '外链', value: 4 },
  { label: '内链', value: 5 },
];

/** 打开弹窗 */
async function open(id?: string) {
  visible.value = true;
  isEdit.value = !!id;
  // 重置表单
  Object.assign(formData, { ...initialForm });
  await nextTick();
  formRef.value?.clearValidate();

  // 加载树形数据
  try {
    const res = await getMenuTree();
    treeData.value = [
      { id: '0', name: '顶级节点', children: res ?? [] } as any,
    ];
  } catch {
    // ignore
  }

  if (id) {
    editId.value = id;
    try {
      const detail = await getMenuDetail(id);
      if (detail) {
        formData.parentID = detail.parentID ?? initialForm.parentID;
        formData.title = detail.title ?? initialForm.title;
        formData.type = detail.type ?? initialForm.type;
        formData.path = detail.path ?? initialForm.path;
        formData.component = detail.component ?? initialForm.component;
        formData.permission = detail.permission ?? initialForm.permission;
        formData.icon = detail.icon ?? initialForm.icon;
        formData.sort = detail.sort ?? initialForm.sort;
        formData.isShow = detail.isShow ?? initialForm.isShow;
        formData.isCache = detail.isCache ?? initialForm.isCache;
        formData.linkURL = detail.linkURL ?? initialForm.linkURL;
        formData.status = detail.status ?? initialForm.status;
      }
    } catch {
      message.error('获取详情失败');
    }
  }
}

/** 提交 */
async function handleOk() {
  try {
    await formRef.value?.validate();
  } catch {
    return;
  }

  confirmLoading.value = true;
  try {
    if (isEdit.value) {
      await updateMenu({
        id: editId.value,
        parentID: formData.parentID,
        title: formData.title,
        type: formData.type,
        path: formData.path,
        component: formData.component,
        permission: formData.permission,
        icon: formData.icon,
        sort: formData.sort,
        isShow: formData.isShow,
        isCache: formData.isCache,
        linkURL: formData.linkURL,
        status: formData.status,
      });
      message.success('更新成功');
    } else {
      await createMenu({
        parentID: formData.parentID,
        title: formData.title,
        type: formData.type,
        path: formData.path,
        component: formData.component,
        permission: formData.permission,
        icon: formData.icon,
        sort: formData.sort,
        isShow: formData.isShow,
        isCache: formData.isCache,
        linkURL: formData.linkURL,
        status: formData.status,
      });
      message.success('创建成功');
    }
    visible.value = false;
    emit('success');
  } finally {
    confirmLoading.value = false;
  }
}

defineExpose({ open });
</script>

<template>
  <a-modal
    v-model:open="visible"
    :title="isEdit ? '编辑菜单表' : '新建菜单表'"
    :confirm-loading="confirmLoading"
    :destroy-on-close="false"
    width="600px"
    @ok="handleOk"
  >
    <a-form
      ref="formRef"
      :model="formData"
      :rules="rules"
      :label-col="{ span: 5 }"
      :wrapper-col="{ span: 17 }"
      class="mt-4"
    >
      <a-form-item label="上级菜单ID，0 表示顶级菜单" name="parentID">
        <a-tree-select
          v-model:value="formData.parentID"
          :tree-data="treeData"
          :field-names="{ label: 'name', value: 'id', children: 'children' }"
          placeholder="请选择上级菜单ID，0 表示顶级菜单"
          allow-clear
          tree-default-expand-all
        />
      </a-form-item>
      <a-form-item label="菜单名称" name="title">
        <a-input v-model:value="formData.title" placeholder="请输入菜单名称" :maxlength="50" />
      </a-form-item>
      <a-form-item label="类型" name="type">
        <a-select v-model:value="formData.type" :options="typeOptions" placeholder="请选择类型" allow-clear />
      </a-form-item>
      <a-form-item label="前端路由路径" name="path">
        <a-input v-model:value="formData.path" placeholder="请输入前端路由路径" :maxlength="200" />
      </a-form-item>
      <a-form-item label="前端组件路径" name="component">
        <a-input v-model:value="formData.component" placeholder="请输入前端组件路径" :maxlength="200" />
      </a-form-item>
      <a-form-item label="权限标识（如 system" name="permission">
        <a-input v-model:value="formData.permission" placeholder="请输入权限标识（如 system" :maxlength="100" />
      </a-form-item>
      <a-form-item label="菜单图标（图标名称）" name="icon">
        <a-input v-model:value="formData.icon" placeholder="请输入图标名称" />
      </a-form-item>
      <a-form-item label="排序（升序）" name="sort">
        <a-input-number v-model:value="formData.sort" placeholder="请输入排序（升序）" style="width: 100%" />
      </a-form-item>
      <a-form-item label="是否显示" name="isShow">
        <a-switch v-model:checked="formData.isShow" :checked-value="1" :un-checked-value="0" />
      </a-form-item>
      <a-form-item label="是否缓存" name="isCache">
        <a-switch v-model:checked="formData.isCache" :checked-value="1" :un-checked-value="0" />
      </a-form-item>
      <a-form-item label="外链/内链地址（type=4或5时有效）" name="linkURL">
        <a-input v-model:value="formData.linkURL" placeholder="请输入外链/内链地址（type=4或5时有效）" addon-before="https://" />
      </a-form-item>
      <a-form-item label="状态" name="status">
        <a-switch v-model:checked="formData.status" :checked-value="1" :un-checked-value="0" />
      </a-form-item>
    </a-form>
  </a-modal>
</template>
