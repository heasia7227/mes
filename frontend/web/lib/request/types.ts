/** 后端统一响应结构：{ code, message, data } */
export interface ApiResponse<T = unknown> {
  code: number;
  message: string;
  data?: T;
}

/** 分页信息 */
export interface Pagination {
  total: number | string;
  page_num: string;
  page_size: string;
}

/** 列表响应 data 结构 */
export interface ListData<T> {
  list: T[];
  pagination: Pagination;
}

/** 列表接口响应：ApiResponse<ListData<T>> */
export type ListResponse<T> = ApiResponse<ListData<T>>;

/** 通用 CRUD 成功响应 */
export interface CrudSuccess {
  success: boolean;
  id?: string;
}

/** 列表查询参数 */
export interface ListParams {
  page_num?: number;
  page_size?: number;
  [key: string]: unknown;
}
