"use client";

import ResourcePage from "@/components/ResourcePage";
import { ranks } from "@/lib/mock-data";
import { NORMAL_STATUS_MAP } from "@/lib/status";

const formStatusOptions = [
  { label: "启用", value: 1 },
  { label: "停用", value: 0 },
];

const filterStatusOptions = [
  { label: "启用", value: "1" },
  { label: "停用", value: "0" },
];

export default function RankPage() {
  return (
    <ResourcePage
      title="职级管理"
      description="维护员工技能职级，对应数据表 t_sys_rank。"
      data={ranks}
      displayKey="rank_name"
      searchKeys={["rank_name", "remark"]}
      searchPlaceholder="搜索职级名称 / 备注"
      statusField="rank_status"
      statusOptions={filterStatusOptions}
      columns={[
        { key: "rank_name", title: "职级名称", width: 180 },
        {
          key: "rank_status",
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
        { key: "rank_name", label: "职级名称", required: true },
        { key: "sort_no", label: "排序", type: "number", initialValue: 1 },
        { key: "remark", label: "备注", type: "textarea" },
      ]}
    />
  );
}
