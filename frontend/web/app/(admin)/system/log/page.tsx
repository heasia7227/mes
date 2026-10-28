"use client";

import ResourcePage from "@/components/ResourcePage";
import { logs } from "@/lib/mock-data";
import { LOG_STATUS_MAP, YES_NO_MAP } from "@/lib/status";

export default function LogPage() {
  return (
    <ResourcePage
      title="操作日志"
      description="记录后台关键操作与异常信息，对应数据表 t_sys_log。"
      data={logs}
      readOnly
      fields={[]}
      displayKey="description"
      searchKeys={["operate_module", "operate_type", "operate_by", "log_content", "ip"]}
      searchPlaceholder="搜索模块 / 操作类型 / 操作人 / IP"
      statusField="status"
      statusOptions={[
        { label: "成功", value: "1" },
        { label: "失败", value: "0" },
      ]}
      columns={[
        { key: "operate_module", title: "操作模块", width: 130 },
        { key: "operate_type", title: "操作类型", width: 100 },
        { key: "operate_by", title: "操作人", width: 110 },
        { key: "ip", title: "IP 地址", width: 140 },
        { key: "log_content", title: "操作内容", ellipsis: true },
        { key: "is_exception", title: "是否异常", width: 100, valueType: "status", statusMap: YES_NO_MAP },
        { key: "status", title: "结果", width: 90, valueType: "status", statusMap: LOG_STATUS_MAP },
        { key: "duration", title: "耗时(ms)", width: 100, valueType: "number" },
        { key: "operate_time", title: "操作时间", width: 170, valueType: "date" },
      ]}
    />
  );
}
