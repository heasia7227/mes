import { request } from "@/lib/request/fetch";
import type { ApiResponse, CrudSuccess } from "@/lib/request/types";

/** 系统信息 */
export interface SystemInfo {
  id: string;
  system_name: string;
  system_logo: string;
  login_background: string;
  footer_text: string;
  create_by: string;
  create_time: string;
  update_by: string;
  update_time: string;
  delete_flag: number;
}

/** 更新系统信息请求体 */
export interface InfoPayload {
  system_name?: string;
  system_logo?: string;
  login_background?: string;
  footer_text?: string;
}

/** GET /system/info 获取系统信息 */
export async function getSystemInfo(): Promise<ApiResponse<SystemInfo>> {
  return request("/system/info", { method: "GET" });
}

/** PUT /system/info/:id 更新系统信息 */
export async function updateSystemInfo(
  id: string,
  payload: InfoPayload,
): Promise<ApiResponse<CrudSuccess>> {
  return request(`/system/info/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
}
