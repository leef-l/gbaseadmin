<script setup lang="ts">
import { ref } from 'vue';
import { message } from 'ant-design-vue';
import { getMenuTree } from '#/api/system/menu';
import { getRoleMenuIds, grantRoleMenu } from '#/api/system/role';
import type { MenuItem } from '#/api/system/menu/types';

const emit = defineEmits<{ success: [] }>();

const visible = ref(false);
const confirmLoading = ref(false);
const roleId = ref('');
const roleName = ref('');
const checkedKeys = ref<string[]>([]);
const treeData = ref<MenuItem[]>([]);
const expandedKeys = ref<string[]>([]);

/** 递归收集所有节点 key */
function collectKeys(nodes: MenuItem[]): string[] {
  const keys: string[] = [];
  for (const node of nodes) {
    keys.push(node.id);
    if (node.children?.length) {
      keys.push(...collectKeys(node.children));
    }
  }
  return keys;
}

/** 打开弹窗 */
async function open(id: string, name: string) {
  visible.value = true;
  roleId.value = id;
  roleName.value = name;
  checkedKeys.value = [];

  try {
    // 加载菜单树
    const res = await getMenuTree();
    treeData.value = (res ?? []) as MenuTreeItem[];
    expandedKeys.value = collectKeys(treeData.value);

    // 加载已授权菜单
    const menuRes = await getRoleMenuIds(id);
    checkedKeys.value = menuRes?.menuIds ?? [];
  } catch {
    message.error('加载数据失败');
  }
}

/** 提交 */
async function handleOk() {
  confirmLoading.value = true;
  try {
    await grantRoleMenu({
      id: roleId.value,
      menuIds: checkedKeys.value,
    });
    message.success('授权成功');
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
    :title="`授权菜单 - ${roleName}`"
    :confirm-loading="confirmLoading"
    width="500px"
    @ok="handleOk"
  >
    <div class="max-h-[400px] overflow-auto py-2">
      <a-tree
        v-model:checkedKeys="checkedKeys"
        v-model:expandedKeys="expandedKeys"
        :tree-data="treeData"
        :field-names="{ title: 'title', key: 'id', children: 'children' }"
        checkable
        check-strictly
      />
    </div>
  </a-modal>
</template>
