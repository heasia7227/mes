import { request } from "@/lib/request/fetch";
import type {
  ApiResponse,
  CrudSuccess,
  ListParams,
  ListResponse,
} from "@/lib/request/types";

/** 部门信息 */
export interface Dept {
  id: string;
  parent_id: string | null;
  dept_name: string;
  dept_code: string;
  dept_full_name: string;
  dept_type: string;
  sort: number;
  remark: string;
  status: string;
  id_path: string;
  name_path: string;
  create_by: string;
  create_time: string;
  update_by: string;
  update_time: string;
  delete_flag: number;
}

/** 创建/更新部门请求体 */
export interface DeptPayload {
  parent_id?: string | null;
  dept_name?: string;
  dept_code?: string;
  dept_full_name?: string;
  dept_type?: string;
  sort?: number;
  remark?: string;
}

/** GET /system/dept/list 部门列表 */
export async function getDeptList(
  params: ListParams = {},
): Promise<ListResponse<Dept>> {
  const query = new URLSearchParams();
  Object.entries(params).forEach(([k, v]) => {
    if (v !== undefined && v !== null) query.set(k, String(v));
  });
  const qs = query.toString();
  return request(`/system/dept/list${qs ? `?${qs}` : ""}`, { method: "GET" });
}

/** GET /system/dept/:id 部门详情 */
export async function getDeptDetail(id: string): Promise<ApiResponse<Dept>> {
  return request(`/system/dept/${id}`, { method: "GET" });
}

/** POST /system/dept 创建部门 */
export async function createDept(
  payload: DeptPayload,
): Promise<ApiResponse<CrudSuccess>> {
  return request("/system/dept", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
}

/** PUT /system/dept/:id 更新部门 */
export async function updateDept(
  id: string,
  payload: DeptPayload,
): Promise<ApiResponse<CrudSuccess>> {
  return request(`/system/dept/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
}

/** DELETE /system/dept/:id 删除部门 */
export async function deleteDept(
  id: string,
): Promise<ApiResponse<CrudSuccess>> {
  return request(`/system/dept/${id}`, { method: "DELETE" });
}

/** PUT /system/dept/:id/status 更新部门状态 */
export async function updateDeptStatus(
  id: string,
  status: "enable" | "disable",
): Promise<ApiResponse<CrudSuccess>> {
  return request(`/system/dept/${id}/status`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ status }),
  });
}
