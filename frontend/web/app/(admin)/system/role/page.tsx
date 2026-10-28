"use client";

import ResourcePage from "@/components/ResourcePage";
import { roles } from "@/lib/mock-data";
import { STATUS_MAP } from "@/lib/status";

const statusOptions = [
  { label: "启用", value: "1" },
  { label: "禁用", value: "0" },
];

export default function RolePage() {
  return (
    <ResourcePage
      title="角色管理"
      description="维护角色编码与权限归属，对应数据表 t_sys_role。"
      data={roles}
      displayKey="title"
      searchKeys={["title", "code", "description"]}
      searchPlaceholder="搜索角色名称 / 编码 / 描述"
      statusField="status"
      statusOptions={statusOptions}
      columns={[
        { key: "title", title: "角色名称", width: 150 },
        { key: "code", title: "角色编码", width: 190 },
        { key: "description", title: "描述", ellipsis: true },
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
        { key: "title", label: "角色名称", required: true },
        { key: "code", label: "角色编码", required: true },
        {
          key: "description",
          label: "角色描述",
          type: "textarea",
          combineCol: true,
        },
      ]}
    />
  );
}
