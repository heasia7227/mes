import { request } from "@/lib/request/fetch";
import type { ApiResponse, CrudSuccess, ListParams, ListResponse } from "@/lib/request/types";
import type { SysDictionary } from "@/lib/types/system/dictionary";

/** GET /system/dictionary/list 字典列表 */
export async function getDictionaryList(params: ListParams = {}): Promise<ListResponse<SysDictionary>> {
    const query = new URLSearchParams();
    Object.entries(params).forEach(([k, v]) => {
        if (v !== undefined && v !== null) query.set(k, String(v));
    });
    const qs = query.toString();
    return request(`/system/dictionary/list${qs ? `?${qs}` : ""}`, {
        method: "GET",
    });
}

/** GET /system/dictionary/:id 字典详情 */
export async function getDictionaryDetail(id: string): Promise<ApiResponse<SysDictionary>> {
    return request(`/system/dictionary/${id}`, { method: "GET" });
}

/** POST /system/dictionary 创建字典 */
export async function createDictionary(payload: SysDictionary): Promise<ApiResponse<CrudSuccess>> {
    return request("/system/dictionary", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
    });
}

/** PUT /system/dictionary/:id 更新字典 */
export async function updateDictionary(id: string, payload: SysDictionary): Promise<ApiResponse<CrudSuccess>> {
    return request(`/system/dictionary/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
    });
}

/** DELETE /system/dictionary/:id 删除字典 */
export async function deleteDictionary(id: string): Promise<ApiResponse<CrudSuccess>> {
    return request(`/system/dictionary/${id}`, { method: "DELETE" });
}
