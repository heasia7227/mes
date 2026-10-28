"use client";

import ResourcePage from "@/components/ResourcePage";
import { loginLogs } from "@/lib/mock-data";
import { LOG_STATUS_MAP } from "@/lib/status";

export default function LoginLogPage() {
  return (
    <ResourcePage
      title="登录日志"
      description="记录账号登录明细，对应数据表 t_sys_login_log。"
      data={loginLogs}
      readOnly
      fields={[]}
      displayKey="login_name"
      searchKeys={["login_name", "login_ip", "login_location", "browser", "os"]}
      searchPlaceholder="搜索账号 / IP / 登录位置 / 浏览器"
      statusField="login_status"
      statusOptions={[
        { label: "成功", value: "1" },
        { label: "失败", value: "0" },
      ]}
      columns={[
        { key: "login_name", title: "登录账号", width: 120 },
        { key: "login_ip", title: "登录 IP", width: 140 },
        { key: "login_location", title: "登录位置", width: 160 },
        { key: "browser", title: "浏览器", width: 130 },
        { key: "os", title: "操作系统", width: 130 },
        {
          key: "login_status",
          title: "状态",
          width: 90,
          valueType: "status",
          statusMap: LOG_STATUS_MAP,
        },
        { key: "message", title: "提示信息", ellipsis: true },
        { key: "login_time", title: "登录时间", width: 170, valueType: "date" },
      ]}
    />
  );
}
