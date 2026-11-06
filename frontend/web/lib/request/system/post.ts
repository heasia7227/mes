import { request } from "@/lib/request/fetch";
import type {
  ApiResponse,
  CrudSuccess,
  ListParams,
  ListResponse,
} from "@/lib/request/types";

/** 岗位信息 */
export interface Post {
  id: string;
  post_name: string;
  post_code: string;
  status: string;
  sort: number;
  remark: string;
  create_by: string;
  create_time: string;
  update_by: string;
  update_time: string;
  delete_flag: number;
}

/** 创建/更新岗位请求体 */
export interface PostPayload {
  post_name?: string;
  post_code?: string;
  sort?: number;
  remark?: string;
}

/** GET /system/post/list 岗位列表 */
export async function getPostList(
  params: ListParams = {},
): Promise<ListResponse<Post>> {
  const query = new URLSearchParams();
  Object.entries(params).forEach(([k, v]) => {
    if (v !== undefined && v !== null) query.set(k, String(v));
  });
  const qs = query.toString();
  return request(`/system/post/list${qs ? `?${qs}` : ""}`, { method: "GET" });
}

/** GET /system/post/:id 岗位详情 */
export async function getPostDetail(id: string): Promise<ApiResponse<Post>> {
  return request(`/system/post/${id}`, { method: "GET" });
}

/** POST /system/post 创建岗位 */
export async function createPost(
  payload: PostPayload,
): Promise<ApiResponse<CrudSuccess>> {
  return request("/system/post", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
}

/** PUT /system/post/:id 更新岗位 */
export async function updatePost(
  id: string,
  payload: PostPayload,
): Promise<ApiResponse<CrudSuccess>> {
  return request(`/system/post/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
}

/** DELETE /system/post/:id 删除岗位 */
export async function deletePost(
  id: string,
): Promise<ApiResponse<CrudSuccess>> {
  return request(`/system/post/${id}`, { method: "DELETE" });
}

/** PUT /system/post/:id/status 更新岗位状态 */
export async function updatePostStatus(
  id: string,
  status: "enable" | "disable",
): Promise<ApiResponse<CrudSuccess>> {
  return request(`/system/post/${id}/status`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ status }),
  });
}
