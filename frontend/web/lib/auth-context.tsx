"use client";

import {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useMemo,
  useState,
  type ReactNode,
} from "react";
import { loginRequest } from "@/lib/request/system/login";

const AUTH_KEY = "mes_auth";

export interface AuthUser {
  id: string;
  user_name: string;
  login_name: string;
  avatar_url?: string;
  introduction?: string;
}

interface AuthSession {
  token: string;
  user: AuthUser;
}

interface LoginResult {
  ok: boolean;
  message?: string;
}

export interface LoginParams {
  loginName: string;
  password: string;
  captcha: string;
  captchaKey: string;
  remember: boolean;
}

interface AuthContextValue {
  user: AuthUser | null;
  token: string | null;
  loading: boolean;
  login: (params: LoginParams) => Promise<LoginResult>;
  logout: () => void;
}

const AuthContext = createContext<AuthContextValue | null>(null);

function readSession(): AuthSession | null {
  if (typeof window === "undefined") return null;
  const raw =
    window.localStorage.getItem(AUTH_KEY) ??
    window.sessionStorage.getItem(AUTH_KEY);
  if (!raw) return null;
  try {
    return JSON.parse(raw) as AuthSession;
  } catch {
    window.localStorage.removeItem(AUTH_KEY);
    window.sessionStorage.removeItem(AUTH_KEY);
    return null;
  }
}

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<AuthUser | null>(null);
  const [token, setToken] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const session = readSession();
    setUser(session?.user ?? null);
    setToken(session?.token ?? null);
    setLoading(false);
  }, []);

  const login = useCallback(
    async ({
      loginName,
      password,
      captcha,
      captchaKey,
      remember,
    }: LoginParams): Promise<LoginResult> => {
      const result = await loginRequest({
        login_name: loginName.trim(),
        password,
        captcha: captcha.trim(),
        captcha_key: captchaKey,
      });

      if (result.code !== 200 || !result.data?.token) {
        return { ok: false, message: result.message || "登录失败" };
      }

      const { token: authToken, user: profile } = result.data;
      const session: AuthSession = {
        token: authToken,
        user: {
          id: profile.id,
          user_name: profile.user_name,
          login_name: profile.login_name,
          avatar_url: profile.avatar ?? "",
        },
      };

      if (remember) {
        window.localStorage.setItem(AUTH_KEY, JSON.stringify(session));
        window.sessionStorage.removeItem(AUTH_KEY);
      } else {
        window.sessionStorage.setItem(AUTH_KEY, JSON.stringify(session));
        window.localStorage.removeItem(AUTH_KEY);
      }

      setUser(session.user);
      setToken(session.token);
      return { ok: true };
    },
    [],
  );

  const logout = useCallback(() => {
    window.localStorage.removeItem(AUTH_KEY);
    window.sessionStorage.removeItem(AUTH_KEY);
    setUser(null);
    setToken(null);
  }, []);

  const value = useMemo(
    () => ({ user, token, loading, login, logout }),
    [user, token, loading, login, logout],
  );

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth 必须在 AuthProvider 内使用");
  }
  return context;
}
