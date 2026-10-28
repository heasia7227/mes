"use client";

import { App as AntdApp, ConfigProvider, theme as antdTheme } from "antd";
import zhCN from "antd/locale/zh_CN";
import dayjs from "dayjs";
import "dayjs/locale/zh-cn";
import {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useMemo,
  useState,
  type ReactNode,
} from "react";

dayjs.locale("zh-cn");

export type ThemeMode = "light" | "dark" | "system";

export interface ThemeSettings {
  mode: ThemeMode;
  primaryColor: string;
  compact: boolean;
  borderRadius: number;
}

export const DEFAULT_THEME_SETTINGS: ThemeSettings = {
  mode: "light",
  primaryColor: "#1677ff",
  compact: false,
  borderRadius: 4,
};

export const THEME_COLORS: Array<{ label: string; value: string }> = [
  { label: "拂晓蓝", value: "#1677ff" },
  { label: "薄暮紫", value: "#722ed1" },
  { label: "火山橙", value: "#fa541c" },
  { label: "日暮金", value: "#faad14" },
  { label: "极光青", value: "#13c2c2" },
  { label: "极光绿", value: "#52c41a" },
  { label: "酱紫红", value: "#eb2f96" },
];

const THEME_KEY = "mes_theme_settings";

interface ThemeContextValue {
  settings: ThemeSettings;
  isDark: boolean;
  updateSettings: (patch: Partial<ThemeSettings>) => void;
  resetSettings: () => void;
}

const ThemeContext = createContext<ThemeContextValue | null>(null);

function readSettings(): ThemeSettings {
  if (typeof window === "undefined") return DEFAULT_THEME_SETTINGS;
  const raw = window.localStorage.getItem(THEME_KEY);
  if (!raw) return DEFAULT_THEME_SETTINGS;
  try {
    return { ...DEFAULT_THEME_SETTINGS, ...(JSON.parse(raw) as ThemeSettings) };
  } catch {
    return DEFAULT_THEME_SETTINGS;
  }
}

export function ThemeProvider({ children }: { children: ReactNode }) {
  const [settings, setSettings] = useState<ThemeSettings>(
    DEFAULT_THEME_SETTINGS,
  );
  const [systemDark, setSystemDark] = useState(false);

  useEffect(() => {
    setSettings(readSettings());

    const media = window.matchMedia("(prefers-color-scheme: dark)");
    setSystemDark(media.matches);
    const onChange = (event: MediaQueryListEvent) =>
      setSystemDark(event.matches);
    media.addEventListener("change", onChange);
    return () => media.removeEventListener("change", onChange);
  }, []);

  const updateSettings = useCallback((patch: Partial<ThemeSettings>) => {
    setSettings((prev) => {
      const next = { ...prev, ...patch };
      window.localStorage.setItem(THEME_KEY, JSON.stringify(next));
      return next;
    });
  }, []);

  const resetSettings = useCallback(() => {
    setSettings(DEFAULT_THEME_SETTINGS);
    window.localStorage.removeItem(THEME_KEY);
  }, []);

  const isDark =
    settings.mode === "dark" || (settings.mode === "system" && systemDark);

  useEffect(() => {
    document.documentElement.style.colorScheme = isDark ? "dark" : "light";
  }, [isDark]);

  const algorithm = useMemo(() => {
    const list = [
      isDark ? antdTheme.darkAlgorithm : antdTheme.defaultAlgorithm,
    ];
    if (settings.compact) list.push(antdTheme.compactAlgorithm);
    return list;
  }, [isDark, settings.compact]);

  const contextValue = useMemo(
    () => ({ settings, isDark, updateSettings, resetSettings }),
    [settings, isDark, updateSettings, resetSettings],
  );

  return (
    <ThemeContext.Provider value={contextValue}>
      <ConfigProvider
        locale={zhCN}
        theme={{
          algorithm,
          token: {
            colorPrimary: settings.primaryColor,
            borderRadius: settings.borderRadius,
          },
        }}
      >
        <AntdApp style={{ minHeight: "100vh" }}>{children}</AntdApp>
      </ConfigProvider>
    </ThemeContext.Provider>
  );
}

export function useThemeSettings() {
  const context = useContext(ThemeContext);
  if (!context) {
    throw new Error("useThemeSettings 必须在 ThemeProvider 内使用");
  }
  return context;
}
