"use client";

import ResourcePage from "@/components/ResourcePage";
import { configs } from "@/lib/mock-data";
import { CONFIG_TYPE_MAP } from "@/lib/status";

export default function ConfigPage() {
  return (
    <ResourcePage
      title="参数配置"
      description="维护系统运行参数，对应数据表 t_sys_config。"
      data={configs}
      rowKey="config_id"
      displayKey="config_name"
      searchKeys={["config_name", "config_key", "config_value"]}
      searchPlaceholder="搜索参数名称 / 参数键 / 参数值"
      columns={[
        { key: "config_name", title: "参数名称", width: 170 },
        { key: "config_key", title: "参数键名", width: 220 },
        { key: "config_value", title: "参数值", width: 220, ellipsis: true },
        {
          key: "config_type",
          title: "类型",
          width: 110,
          valueType: "status",
          statusMap: CONFIG_TYPE_MAP,
        },
        { key: "remark", title: "备注", ellipsis: true },
        { key: "update_time", title: "更新时间", width: 170, valueType: "date" },
      ]}
      fields={[
        { key: "config_name", label: "参数名称", required: true },
        { key: "config_key", label: "参数键名", required: true },
        { key: "config_value", label: "参数值", required: true },
        {
          key: "config_type",
          label: "参数类型",
          type: "select",
          options: [
            { label: "系统内置", value: "Y" },
            { label: "自定义", value: "N" },
          ],
          initialValue: "N",
        },
        { key: "remark", label: "备注", type: "textarea" },
      ]}
    />
  );
}
