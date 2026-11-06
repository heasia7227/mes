import { request } from "@/lib/request/fetch";
import type { ListParams, ListResponse } from "@/lib/request/types";

/** 登录日志信息 */
export interface LoginLog {
  id: string;
  login_name: string;
  login_ip: string;
  browser: string;
  os: string;
  status: string;
  message: string;
  login_time: string;
}

/** GET /system/login-log/list 登录日志列表 */
export async function getLoginLogList(
  params: ListParams = {},
): Promise<ListResponse<LoginLog>> {
  const query = new URLSearchParams();
  Object.entries(params).forEach(([k, v]) => {
    if (v !== undefined && v !== null) query.set(k, String(v));
  });
  const qs = query.toString();
  return request(`/system/login-log/list${qs ? `?${qs}` : ""}`, {
    method: "GET",
  });
}
