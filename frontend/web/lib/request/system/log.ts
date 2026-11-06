import { request } from "@/lib/request/fetch";
import type { ListParams, ListResponse } from "@/lib/request/types";

/** 操作日志信息 */
export interface Log {
  id: string;
  ip: string;
  log_content: string;
  is_exception: string;
  action: string;
  module: string;
  operate_by: string;
  operate_time: string;
  status: string;
  duration: number;
  description: string;
}

/** GET /system/log/list 操作日志列表 */
export async function getLogList(
  params: ListParams = {},
): Promise<ListResponse<Log>> {
  const query = new URLSearchParams();
  Object.entries(params).forEach(([k, v]) => {
    if (v !== undefined && v !== null) query.set(k, String(v));
  });
  const qs = query.toString();
  return request(`/system/log/list${qs ? `?${qs}` : ""}`, { method: "GET" });
}
