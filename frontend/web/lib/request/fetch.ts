/**
 * 后端接口基础访问层
 * 默认后端网关地址：http://localhost:8080/api
 * 可通过环境变量 NEXT_PUBLIC_API_BASE_URL 覆盖
 */
export const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:8080/api";

export const request = async (url: string, options: RequestInit) => {
    const response = await fetch(`${API_BASE_URL}${url}`, options);
    const data = await response.json();
    return data;
};
