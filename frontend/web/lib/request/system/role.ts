import { request } from "@/lib/request/fetch";
import type {
  ApiResponse,
  CrudSuccess,
  ListParams,
  ListResponse,
} from "@/lib/request/types";

/** 角色信息 */
export interface Role {
  id: string;
  title: string;
  code: string;
  description: string;
  status: string;
  create_by: string;
  create_time: string;
  update_by: string;
  update_time: string;
  delete_flag: number;
}

/** 创建/更新角色请求体 */
export interface RolePayload {
  title?: string;
  code?: string;
  description?: string;
}

/** GET /system/role/list 角色列表 */
export async function getRoleList(
  params: ListParams = {},
): Promise<ListResponse<Role>> {
  const query = new URLSearchParams();
  Object.entries(params).forEach(([k, v]) => {
    if (v !== undefined && v !== null) query.set(k, String(v));
  });
  const qs = query.toString();
  return request(`/system/role/list${qs ? `?${qs}` : ""}`, { method: "GET" });
}

/** GET /system/role/:id 角色详情 */
export async function getRoleDetail(id: string): Promise<ApiResponse<Role>> {
  return request(`/system/role/${id}`, { method: "GET" });
}

/** POST /system/role 创建角色 */
export async function createRole(
  payload: RolePayload,
): Promise<ApiResponse<CrudSuccess>> {
  return request("/system/role", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
}

/** PUT /system/role/:id 更新角色 */
export async function updateRole(
  id: string,
  payload: RolePayload,
): Promise<ApiResponse<CrudSuccess>> {
  return request(`/system/role/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
}

/** DELETE /system/role/:id 删除角色 */
export async function deleteRole(
  id: string,
): Promise<ApiResponse<CrudSuccess>> {
  return request(`/system/role/${id}`, { method: "DELETE" });
}

/** PUT /system/role/:id/status 更新角色状态 */
export async function updateRoleStatus(
  id: string,
  status: "enable" | "disable",
): Promise<ApiResponse<CrudSuccess>> {
  return request(`/system/role/${id}/status`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ status }),
  });
}
