import { request } from "@/lib/request/fetch";
import type { ApiResponse, CrudSuccess, ListParams, ListResponse } from "@/lib/request/types";
import type { SysDictionaryCatalog } from "@/lib/types/system/dictionary-catalog";

/** GET /system/dictionary-catalog/list 字典目录列表 */
export async function getDictionaryCatalogList(params: ListParams = {}): Promise<ListResponse<SysDictionaryCatalog>> {
    const query = new URLSearchParams();
    Object.entries(params).forEach(([k, v]) => {
        if (v !== undefined && v !== null) query.set(k, String(v));
    });
    const qs = query.toString();
    return request(`/system/dictionary-catalog/list${qs ? `?${qs}` : ""}`, {
        method: "GET",
    });
}

/** GET /system/dictionary-catalog/:id 字典目录详情 */
export async function getDictionaryCatalogDetail(id: string): Promise<ApiResponse<SysDictionaryCatalog>> {
    return request(`/system/dictionary-catalog/${id}`, { method: "GET" });
}

/** POST /system/dictionary-catalog 创建字典目录 */
export async function createDictionaryCatalog(payload: SysDictionaryCatalog): Promise<ApiResponse<CrudSuccess>> {
    return request("/system/dictionary-catalog", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
    });
}

/** PUT /system/dictionary-catalog/:id 更新字典目录 */
export async function updateDictionaryCatalog(
    id: string,
    payload: SysDictionaryCatalog,
): Promise<ApiResponse<CrudSuccess>> {
    return request(`/system/dictionary-catalog/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
    });
}

/** DELETE /system/dictionary-catalog/:id 删除字典目录 */
export async function deleteDictionaryCatalog(id: string): Promise<ApiResponse<CrudSuccess>> {
    return request(`/system/dictionary-catalog/${id}`, { method: "DELETE" });
}
