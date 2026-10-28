import { AntdRegistry } from "@ant-design/nextjs-registry";
import type { Metadata } from "next";
import { Suspense } from "react";
import { AuthProvider } from "@/lib/auth-context";
import { ThemeProvider } from "@/lib/theme-context";
import "./globals.css";

export const metadata: Metadata = {
  title: "MES 制造执行系统",
  description: "MES 制造执行系统后台管理平台",
};

export default function RootLayout({ children }: LayoutProps<"/">) {
  return (
    <html lang="zh-CN">
      <body>
        <Suspense fallback={null}>
          <AntdRegistry>
            <ThemeProvider>
              <AuthProvider>{children}</AuthProvider>
            </ThemeProvider>
          </AntdRegistry>
        </Suspense>
      </body>
    </html>
  );
}
