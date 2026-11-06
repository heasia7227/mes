import { request } from "@/lib/request/fetch";
import type {
  ApiResponse,
  CrudSuccess,
  ListParams,
  ListResponse,
} from "@/lib/request/types";

/** 职级信息 */
export interface Rank {
  id: string;
  rank_name: string;
  status: string;
  sort: number;
  remark: string;
  create_by: string;
  create_time: string;
  update_by: string;
  update_time: string;
  delete_flag: number;
}

/** 创建/更新职级请求体 */
export interface RankPayload {
  rank_name?: string;
  sort?: number;
  remark?: string;
}

/** GET /system/rank/list 职级列表 */
export async function getRankList(
  params: ListParams = {},
): Promise<ListResponse<Rank>> {
  const query = new URLSearchParams();
  Object.entries(params).forEach(([k, v]) => {
    if (v !== undefined && v !== null) query.set(k, String(v));
  });
  const qs = query.toString();
  return request(`/system/rank/list${qs ? `?${qs}` : ""}`, { method: "GET" });
}

/** GET /system/rank/:id 职级详情 */
export async function getRankDetail(id: string): Promise<ApiResponse<Rank>> {
  return request(`/system/rank/${id}`, { method: "GET" });
}

/** POST /system/rank 创建职级 */
export async function createRank(
  payload: RankPayload,
): Promise<ApiResponse<CrudSuccess>> {
  return request("/system/rank", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
}

/** PUT /system/rank/:id 更新职级 */
export async function updateRank(
  id: string,
  payload: RankPayload,
): Promise<ApiResponse<CrudSuccess>> {
  return request(`/system/rank/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
}

/** DELETE /system/rank/:id 删除职级 */
export async function deleteRank(
  id: string,
): Promise<ApiResponse<CrudSuccess>> {
  return request(`/system/rank/${id}`, { method: "DELETE" });
}

/** PUT /system/rank/:id/status 更新职级状态 */
export async function updateRankStatus(
  id: string,
  status: "enable" | "disable",
): Promise<ApiResponse<CrudSuccess>> {
  return request(`/system/rank/${id}/status`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ status }),
  });
}
