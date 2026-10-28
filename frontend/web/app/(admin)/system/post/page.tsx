"use client";

import ResourcePage from "@/components/ResourcePage";
import { posts } from "@/lib/mock-data";
import { NORMAL_STATUS_MAP } from "@/lib/status";

const formStatusOptions = [
  { label: "正常", value: 1 },
  { label: "停用", value: 0 },
];

const filterStatusOptions = [
  { label: "正常", value: "1" },
  { label: "停用", value: "0" },
];

export default function PostPage() {
  return (
    <ResourcePage
      title="岗位管理"
      description="维护岗位信息与编码，对应数据表 t_sys_post。"
      data={posts}
      displayKey="post_name"
      searchKeys={["post_name", "post_code", "remark"]}
      searchPlaceholder="搜索岗位名称 / 编码 / 备注"
      statusField="post_status"
      statusOptions={filterStatusOptions}
      columns={[
        { key: "post_name", title: "岗位名称", width: 160 },
        { key: "post_code", title: "岗位编码", width: 150 },
        {
          key: "post_status",
          title: "状态",
          width: 90,
          valueType: "status",
          statusMap: NORMAL_STATUS_MAP,
        },
        { key: "sort_no", title: "排序", width: 80, valueType: "number" },
        { key: "remark", title: "备注", ellipsis: true },
        {
          key: "create_time",
          title: "创建时间",
          width: 170,
          valueType: "date",
        },
      ]}
      fields={[
        { key: "post_name", label: "岗位名称", required: true },
        { key: "post_code", label: "岗位编码", required: true },
        { key: "sort_no", label: "排序", type: "number", initialValue: 1 },
        { key: "remark", label: "备注", type: "textarea" },
      ]}
    />
  );
}
