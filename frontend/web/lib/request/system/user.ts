import { request } from "@/lib/request/fetch";
import type { ApiResponse, CrudSuccess, ListParams, ListResponse } from "@/lib/request/types";
import type { SysUser, ChangePasswordPayload } from "@/lib/types/system/user";

/** GET /system/user/list 用户列表 */
export async function getUserList(params: ListParams = {}): Promise<ListResponse<SysUser>> {
    const query = new URLSearchParams();
    Object.entries(params).forEach(([k, v]) => {
        if (v !== undefined && v !== null) query.set(k, String(v));
    });
    const qs = query.toString();
    return request(`/system/user/list${qs ? `?${qs}` : ""}`, { method: "GET" });
}

/** GET /system/user/:id 用户详情 */
export async function getUserDetail(id: string): Promise<ApiResponse<SysUser>> {
    return request(`/system/user/${id}`, { method: "GET" });
}

/** POST /system/user 创建用户 */
export async function createUser(payload: SysUser): Promise<ApiResponse<CrudSuccess>> {
    return request("/system/user", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
    });
}

/** PUT /system/user/:id 更新用户 */
export async function updateUser(id: string, payload: SysUser): Promise<ApiResponse<CrudSuccess>> {
    return request(`/system/user/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
    });
}

/** DELETE /system/user/:id 删除用户 */
export async function deleteUser(id: string): Promise<ApiResponse<CrudSuccess>> {
    return request(`/system/user/${id}`, { method: "DELETE" });
}

/** PUT /system/user/:id/status 更新用户状态 */
export async function updateUserStatus(id: string, status: "enable" | "disable"): Promise<ApiResponse<CrudSuccess>> {
    return request(`/system/user/${id}/status`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ status }),
    });
}

/** POST /system/user/:id/change-password 修改密码 */
export async function changePassword(id: string, payload: ChangePasswordPayload): Promise<ApiResponse<CrudSuccess>> {
    return request(`/system/user/${id}/change-password`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
    });
}

/** PUT /system/user/:id/reset-password 重置密码 */
export async function resetPassword(id: string): Promise<ApiResponse<CrudSuccess>> {
    return request(`/system/user/${id}/reset-password`, { method: "PUT" });
}

/** PUT /system/user/:id/offline 用户下线 */
export async function userOffline(id: string): Promise<ApiResponse<CrudSuccess>> {
    return request(`/system/user/${id}/offline`, { method: "PUT" });
}
