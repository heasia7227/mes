"use client";

import ResourcePage from "@/components/ResourcePage";
import { depts } from "@/lib/mock-data";
import { DEPT_TYPE_MAP, STATUS_MAP } from "@/lib/status";

const parentOptions = depts.map((item) => ({
  label: `${item.dept_name}（${item.dept_code}）`,
  value: item.id,
}));

const formStatusOptions = [
  { label: "启用", value: 1 },
  { label: "禁用", value: 0 },
];

const filterStatusOptions = [
  { label: "启用", value: "1" },
  { label: "禁用", value: "0" },
];

export default function DeptPage() {
  return (
    <ResourcePage
      title="部门管理"
      description="维护组织架构与层级关系，对应数据表 t_sys_dept。"
      data={depts}
      displayKey="dept_name"
      searchKeys={["dept_name", "dept_code", "dept_full_name"]}
      searchPlaceholder="搜索部门名称 / 编码 / 全称"
      statusField="status"
      statusOptions={filterStatusOptions}
      columns={[
        { key: "dept_name", title: "部门名称", width: 140 },
        { key: "dept_code", title: "部门编码", width: 130 },
        { key: "dept_full_name", title: "部门全称", ellipsis: true },
        {
          key: "dept_type",
          title: "类型",
          width: 90,
          valueType: "status",
          statusMap: DEPT_TYPE_MAP,
        },
        { key: "sort_no", title: "排序", width: 80, valueType: "number" },
        {
          key: "status",
          title: "状态",
          width: 90,
          valueType: "status",
          statusMap: STATUS_MAP,
        },
        {
          key: "create_time",
          title: "创建时间",
          width: 170,
          valueType: "date",
        },
      ]}
      fieldCol={2}
      fields={[
        { key: "dept_name", label: "部门名称", required: true },
        { key: "dept_code", label: "部门编码", required: true },
        { key: "dept_full_name", label: "部门全称" },
        {
          key: "parent_id",
          label: "上级部门",
          type: "select",
          options: parentOptions,
        },
        {
          key: "dept_type",
          label: "部门类型",
          type: "select",
          options: [
            { label: "公司", value: "公司" },
            { label: "部门", value: "部门" },
            { label: "班组", value: "班组" },
          ],
          initialValue: "部门",
        },
        { key: "sort_no", label: "排序", type: "number", initialValue: 1 },
        { key: "remark", label: "备注", type: "textarea", combineCol: true },
      ]}
    />
  );
}
