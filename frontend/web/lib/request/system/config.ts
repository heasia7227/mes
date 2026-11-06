import { request } from "@/lib/request/fetch";
import type {
  ApiResponse,
  CrudSuccess,
  ListParams,
  ListResponse,
} from "@/lib/request/types";

/** 系统配置信息 */
export interface Config {
  id: string;
  config_name: string;
  config_key: string;
  config_value: string;
  config_type: string;
  create_by: string;
  create_time: string;
  update_by: string;
  update_time: string;
  remark: string;
  delete_flag: number;
}

/** 创建/更新配置请求体 */
export interface ConfigPayload {
  config_name?: string;
  config_key?: string;
  config_value?: string;
  config_type?: string;
  remark?: string;
}

/** GET /system/config/list 配置列表 */
export async function getConfigList(
  params: ListParams = {},
): Promise<ListResponse<Config>> {
  const query = new URLSearchParams();
  Object.entries(params).forEach(([k, v]) => {
    if (v !== undefined && v !== null) query.set(k, String(v));
  });
  const qs = query.toString();
  return request(`/system/config/list${qs ? `?${qs}` : ""}`, { method: "GET" });
}

/** GET /system/config/:id 配置详情 */
export async function getConfigDetail(
  id: string,
): Promise<ApiResponse<Config>> {
  return request(`/system/config/${id}`, { method: "GET" });
}

/** POST /system/config 创建配置 */
export async function createConfig(
  payload: ConfigPayload,
): Promise<ApiResponse<CrudSuccess>> {
  return request("/system/config", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
}

/** PUT /system/config/:id 更新配置 */
export async function updateConfig(
  id: string,
  payload: ConfigPayload,
): Promise<ApiResponse<CrudSuccess>> {
  return request(`/system/config/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
}

/** DELETE /system/config/:id 删除配置 */
export async function deleteConfig(
  id: string,
): Promise<ApiResponse<CrudSuccess>> {
  return request(`/system/config/${id}`, { method: "DELETE" });
}
