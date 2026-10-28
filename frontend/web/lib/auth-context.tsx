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

interface AuthContextValue {
  user: AuthUser | null;
  loading: boolean;
  login: (loginName: string, password: string, remember: boolean) => Promise<LoginResult>;
  logout: () => void;
}

const DEMO_ACCOUNTS: Array<{ loginName: string; password: string; user: AuthUser }> = [
  {
    loginName: "admin",
    password: "123456",
    user: {
      id: "U001",
      user_name: "张伟",
      login_name: "admin",
      avatar_url: "",
      introduction: "系统管理员",
    },
  },
  {
    loginName: "operator",
    password: "123456",
    user: {
      id: "U002",
      user_name: "李娜",
      login_name: "operator",
      avatar_url: "",
      introduction: "生产主管",
    },
  },
];

const AuthContext = createContext<AuthContextValue | null>(null);

function readSession(): AuthSession | null {
  if (typeof window === "undefined") return null;
  const raw = window.localStorage.getItem(AUTH_KEY) ?? window.sessionStorage.getItem(AUTH_KEY);
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
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const session = readSession();
    setUser(session?.user ?? null);
    setLoading(false);
  }, []);

  const login = useCallback(
    async (loginName: string, password: string, remember: boolean): Promise<LoginResult> => {
      const account = DEMO_ACCOUNTS.find(
        (item) => item.loginName.toLowerCase() === loginName.trim().toLowerCase(),
      );

      if (!account || account.password !== password) {
        return { ok: false, message: "账号或密码错误，请使用演示账号 admin / 123456" };
      }

      const session: AuthSession = {
        token: `mock-token-${account.user.id}-${Date.now()}`,
        user: account.user,
      };

      if (remember) {
        window.localStorage.setItem(AUTH_KEY, JSON.stringify(session));
        window.sessionStorage.removeItem(AUTH_KEY);
      } else {
        window.sessionStorage.setItem(AUTH_KEY, JSON.stringify(session));
        window.localStorage.removeItem(AUTH_KEY);
      }

      setUser(session.user);
      return { ok: true };
    },
    [],
  );

  const logout = useCallback(() => {
    window.localStorage.removeItem(AUTH_KEY);
    window.sessionStorage.removeItem(AUTH_KEY);
    setUser(null);
  }, []);

  const value = useMemo(
    () => ({ user, loading, login, logout }),
    [user, loading, login, logout],
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
