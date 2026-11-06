"use client";

import { API_BASE_URL, request } from "@/lib/request/fetch";
import { md5 } from "@/lib/utils/md5";

/** 后端统一响应结构：{ code, message, data } */
export interface ApiResponse<T = unknown> {
  code: number;
  message: string;
  data?: T;
}

/** POST /auth/login 请求体 */
export interface LoginPayload {
  login_name: string;
  password: string;
  captcha: string;
  captcha_key: string;
}

/** POST /auth/login 响应数据 */
export interface LoginData {
  token: string;
  user: {
    id: string;
    login_name: string;
    user_name: string;
    avatar?: string;
  };
}

/** 生成验证码标识（由前端生成，后端据此关联验证码图片与校验值） */
export function createCaptchaKey(): string {
  if (
    typeof crypto !== "undefined" &&
    typeof crypto.randomUUID === "function"
  ) {
    return crypto.randomUUID();
  }
  return `ck-${Date.now()}-${Math.random().toString(36).slice(2)}`;
}

/** GET /auth/captcha?captcha_key=xxx，返回 PNG 图片 */
export function getCaptchaUrl(captchaKey: string): string {
  const query = new URLSearchParams({
    captcha_key: captchaKey,
    t: String(Date.now()),
  });
  return `${API_BASE_URL}/auth/captcha?${query.toString()}`;
}

/** POST /auth/login 登录 */
export async function loginRequest(
  payload: LoginPayload,
): Promise<ApiResponse<LoginData>> {
  payload.password = md5(payload.password);
  const data = await request("/auth/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });

  console.log("data: ", data);

  return data;
}
