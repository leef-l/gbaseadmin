<script setup lang="ts">
import { h } from 'vue';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, useVbenModal } from '@vben/common-ui';
import { Button, message, Modal, Tag, Tooltip } from 'ant-design-vue';
import { QuestionCircleOutlined } from '@ant-design/icons-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getBannerList, deleteBanner, batchDeleteBanner, exportBanner } from '#/api/play/banner';
import type { BannerItem } from '#/api/play/banner/types';
import FormModal from './modules/form.vue';
import DetailDrawer from './modules/detail-drawer.vue';

/** 跳转类型映射 */
const linkTypeMap: Record<number, string> = {
  1: '内页',
  2: '外链',
  3: '活动页',
  4: '商品页',
  5: '陪玩师页',
  6: '唤醒App',
};

/** 跳转类型标签颜色 */
const linkTypeColorMap: Record<number, string> = {
  1: 'blue',
  2: 'cyan',
  3: 'orange',
  4: 'purple',
  5: 'geekblue',
  6: 'magenta',
};

/** 状态映射 */
const statusMap: Record<number, string> = {
  0: '禁用',
  1: '启用',
};

/** 状态标签颜色 */
function getStatusColor(val: number): string {
  return val === 1 ? 'green' : 'red';
}

/** 渲染带 Tooltip 的列标题 */
function tooltipHeader(label: string, tip: string) {
  return () => h('span', {}, [
    label + ' ',
    h(Tooltip, { title: tip }, {
      default: () => h(QuestionCircleOutlined, { style: { color: '#999', marginLeft: '4px' } }),
    }),
  ]);
}

/** 表单弹窗 */
const [FormModalComp, formModalApi] = useVbenModal({
  connectedComponent: FormModal,
  destroyOnClose: true,
});

/** 详情抽屉 */
const [DetailDrawerComp, detailDrawerApi] = useVbenModal({
  connectedComponent: DetailDrawer,
  destroyOnClose: true,
});

/** 搜索表单配置 */
const formOptions: VbenFormProps = {
  collapsed: false,
  showCollapseButton: true,
  submitOnChange: false,
  submitOnEnter: true,
  schema: [
    {
      component: 'Input',
      componentProps: { placeholder: '请输入Banner标题', allowClear: true },
      fieldName: 'title',
      label: 'Banner标题',
    },
    {
      component: 'Input',
      componentProps: { placeholder: '请输入备注', allowClear: true },
      fieldName: 'remark',
      label: '备注',
    },
    {
      component: 'RangePicker',
      fieldName: 'timeRange',
      label: '创建时间',
      componentProps: {
        showTime: true,
        format: 'YYYY-MM-DD HH:mm:ss',
        valueFormat: 'YYYY-MM-DD HH:mm:ss',
        class: 'w-full',
      },
    },
  ],
};

/** 表格列配置 */
const gridOptions: VxeGridProps<BannerItem> = {
  columns: [
    { type: 'checkbox', width: 50 },
    { title: '序号', type: 'seq', width: 50 },
    { field: 'title', title: 'Banner标题' },
    { field: 'image', title: 'Banner图片', width: 80, slots: { default: 'image_cell' } },
    { field: 'linkType', title: '跳转类型', width: 110, slots: { default: 'linkType_cell' } },
    { field: 'linkValue', title: '跳转值', slots: { header: tooltipHeader('跳转值', '页面路径/URL/业务ID/App Scheme') } },
    { field: 'sort', title: '排序', width: 80, slots: { header: tooltipHeader('排序', '越大越前') } },
    { field: 'status', title: '状态', width: 90, slots: { default: 'status_cell' } },
    { field: 'startTime', title: '生效开始时间' },
    { field: 'endTime', title: '生效结束时间' },
    { field: 'remark', title: '备注' },
    { field: 'createdAt', title: '创建时间', width: 180, formatter: 'formatDateTime', sortable: true },
    { title: '操作', width: 240, fixed: 'right', slots: { default: 'action' } },
  ],
  height: 'auto',
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page, sorts }, formValues) => {
        const { timeRange, ...rest } = formValues;
        const params: Record<string, any> = {
          pageNum: page.currentPage,
          pageSize: page.pageSize,
          ...rest,
        };
        if (timeRange && timeRange.length === 2) {
          params.startTime = timeRange[0];
          params.endTime = timeRange[1];
        }
        if (sorts && sorts.length > 0) {
          const sort = sorts[0];
          if (sort && sort.field && sort.order) {
            params.orderBy = sort.field;
            params.orderDir = sort.order;
          }
        }
        const res = await getBannerList(params as any);
        return { items: res?.list ?? [], total: res?.total ?? 0 };
      },
    },
  },
  sortConfig: {
    remote: true,
    trigger: 'cell',
  },
  toolbarConfig: {
    custom: true,
    refresh: true,
    search: true,
  },
};

const [Grid, gridApi] = useVbenVxeGrid({
  formOptions,
  gridOptions,
});

/** 新建 */
function handleCreate() {
  formModalApi.setData(null).open();
}

/** 查看 */
function handleView(row: BannerItem) {
  detailDrawerApi.setData({ id: row.id }).open();
}

/** 编辑 */
function handleEdit(row: BannerItem) {
  formModalApi.setData({ id: row.id }).open();
}

/** 删除 */
function handleDelete(row: BannerItem) {
  Modal.confirm({
    title: '确认删除',
    content: '确定要删除该首页Banner轮播吗？',
    okType: 'danger',
    async onOk() {
      await deleteBanner(row.id);
      message.success('删除成功');
      gridApi.reload();
    },
  });
}
/** 批量删除 */
function handleBatchDelete() {
  const rows = gridApi.grid.getCheckboxRecords();
  if (rows.length === 0) {
    message.warning('请先选择要删除的数据');
    return;
  }
  Modal.confirm({
    title: '确认批量删除',
    content: `确定要删除选中的 ${rows.length} 条首页Banner轮播吗？`,
    okType: 'danger',
    async onOk() {
      await batchDeleteBanner(rows.map((r: BannerItem) => r.id));
      message.success('批量删除成功');
      gridApi.reload();
    },
  });
}

/** 导出 */
async function handleExport() {
  try {
    const formValues = await gridApi.formApi.getValues();
    const params: Record<string, any> = { ...formValues };
    if (params.timeRange && params.timeRange.length === 2) {
      params.startTime = params.timeRange[0];
      params.endTime = params.timeRange[1];
      delete params.timeRange;
    }
    const blob = await exportBanner(params);
    const url = URL.createObjectURL(blob as any);
    const a = document.createElement('a');
    a.href = url;
    a.download = '首页Banner轮播.csv';
    a.click();
    URL.revokeObjectURL(url);
    message.success('导出成功');
  } catch {
    message.error('导出失败');
  }
}
</script>

<template>
  <Page auto-content-height>
    <FormModalComp @success="() => gridApi.reload()" />
    <DetailDrawerComp />
    <Grid>
      <template #toolbar-actions>
        <Button type="primary" @click="handleCreate">新建</Button>
        <Button danger class="ml-2" @click="handleBatchDelete">批量删除</Button>
        <Button class="ml-2" @click="handleExport">导出</Button>
      </template>
      <template #image_cell="{ row }">
        <img v-if="row.image" :src="row.image" style="height:40px;width:60px;object-fit:cover;border-radius:4px;" />
        <span v-else>-</span>
      </template>
      <template #linkType_cell="{ row }">
        <Tag :color="linkTypeColorMap[row.linkType]">
          {{ linkTypeMap[row.linkType] || row.linkType }}
        </Tag>
      </template>
      <template #status_cell="{ row }">
        <Tag :color="getStatusColor(row.status)">
          {{ statusMap[row.status] ?? row.status }}
        </Tag>
      </template>
      <template #action="{ row }">
        <Button type="link" size="small" @click="handleView(row)">查看</Button>
        <Button type="link" size="small" @click="handleEdit(row)">编辑</Button>
        <Button type="link" danger size="small" @click="handleDelete(row)">删除</Button>
      </template>
    </Grid>
  </Page>
</template>
