import { request } from "@/lib/request/fetch";
import type {
  ApiResponse,
  CrudSuccess,
  ListParams,
  ListResponse,
} from "@/lib/request/types";

/** 菜单信息 */
export interface Menu {
  id: string;
  title: string;
  icon: string;
  parent_id: string;
  ranking: number;
  menu_type: string;
  design_type: string;
  menu_path: string;
  permission: string;
  page_type: string;
  status: string;
  create_by: string;
  create_time: string;
  update_by: string;
  update_time: string;
  delete_flag: number;
}

/** 创建/更新菜单请求体 */
export interface MenuPayload {
  title?: string;
  icon?: string;
  parent_id?: string;
  ranking?: number;
  menu_type?: string;
  design_type?: string;
  menu_path?: string;
  permission?: string;
  page_type?: string;
}

/** GET /system/menu/list 菜单列表 */
export async function getMenuList(
  params: ListParams = {},
): Promise<ListResponse<Menu>> {
  const query = new URLSearchParams();
  Object.entries(params).forEach(([k, v]) => {
    if (v !== undefined && v !== null) query.set(k, String(v));
  });
  const qs = query.toString();
  return request(`/system/menu/list${qs ? `?${qs}` : ""}`, { method: "GET" });
}

/** GET /system/menu/:id 菜单详情 */
export async function getMenuDetail(id: string): Promise<ApiResponse<Menu>> {
  return request(`/system/menu/${id}`, { method: "GET" });
}

/** POST /system/menu 创建菜单 */
export async function createMenu(
  payload: MenuPayload,
): Promise<ApiResponse<CrudSuccess>> {
  return request("/system/menu", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
}

/** PUT /system/menu/:id 更新菜单 */
export async function updateMenu(
  id: string,
  payload: MenuPayload,
): Promise<ApiResponse<CrudSuccess>> {
  return request(`/system/menu/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
}

/** DELETE /system/menu/:id 删除菜单 */
export async function deleteMenu(
  id: string,
): Promise<ApiResponse<CrudSuccess>> {
  return request(`/system/menu/${id}`, { method: "DELETE" });
}

/** PUT /system/menu/:id/status 更新菜单状态 */
export async function updateMenuStatus(
  id: string,
  status: "enable" | "disable",
): Promise<ApiResponse<CrudSuccess>> {
  return request(`/system/menu/${id}/status`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ status }),
  });
}
