"use client";

import {
  LogoutOutlined,
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  SettingOutlined,
  ThunderboltFilled,
} from "@ant-design/icons";
import {
  App as AntdApp,
  Avatar,
  Breadcrumb,
  Button,
  Divider,
  Grid,
  Layout,
  Menu,
  Space,
  Spin,
  Tooltip,
  Typography,
  theme,
} from "antd";
import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import { useEffect, useState, type ReactNode } from "react";
import ThemeSettingsDrawer from "@/components/ThemeSettingsDrawer";
import { useAuth } from "@/lib/auth-context";
import { systemInfo } from "@/lib/mock-data";
import {
  findTrail,
  getOpenKeys,
  getSelectedKey,
  navigation,
  toMenuItems,
} from "@/lib/navigation";
import { useThemeSettings } from "@/lib/theme-context";

const { Sider, Header, Content } = Layout;

function FullscreenSpin() {
  return (
    <div
      style={{
        minHeight: "100vh",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
      }}
    >
      <Spin size="large" description="加载中..." />
    </div>
  );
}

export default function AdminLayout({ children }: { children: ReactNode }) {
  const router = useRouter();
  const pathname = usePathname();
  const screens = Grid.useBreakpoint();
  const { user, loading, logout } = useAuth();
  const { isDark } = useThemeSettings();
  const { token } = theme.useToken();
  const { modal, message } = AntdApp.useApp();

  const [collapsed, setCollapsed] = useState(false);
  const [settingsOpen, setSettingsOpen] = useState(false);
  const [openKeys, setOpenKeys] = useState<string[]>([]);

  useEffect(() => {
    if (!loading && !user) {
      router.replace("/login");
    }
  }, [loading, user, router]);

  useEffect(() => {
    setOpenKeys(getOpenKeys(pathname));
  }, [pathname]);

  if (loading || !user) {
    return <FullscreenSpin />;
  }

  const trail = findTrail(pathname);
  const breadcrumbItems =
    trail.length === 0
      ? [{ title: "首页" }]
      : [
          ...(trail[0].key === "/dashboard"
            ? []
            : [{ title: <Link href="/dashboard">首页</Link> }]),
          ...trail.map((item) => ({ title: item.label })),
        ];

  const handleLogout = () => {
    modal.confirm({
      title: "退出登录",
      content: "确定要退出当前账号吗？",
      okText: "退出",
      cancelText: "取消",
      okButtonProps: { danger: true },
      onOk: () => {
        logout();
        message.success("已退出登录");
        router.replace("/login");
      },
    });
  };

  return (
    <Layout style={{ minHeight: "100vh" }}>
      <Sider
        theme={isDark ? "dark" : "light"}
        width={232}
        collapsedWidth={72}
        collapsible
        collapsed={collapsed}
        trigger={null}
        breakpoint="lg"
        onCollapse={(value) => setCollapsed(value)}
        style={{
          position: "sticky",
          top: 0,
          height: "100vh",
          overflow: "auto",
          borderRight: `1px solid ${token.colorBorderSecondary}`,
          background: token.colorBgContainer,
        }}
      >
        <Link
          href="/dashboard"
          style={{
            height: 64,
            display: "flex",
            alignItems: "center",
            gap: 10,
            padding: collapsed ? "0 21px" : "0 18px",
            borderBottom: `1px solid ${token.colorBorderSecondary}`,
            textDecoration: "none",
            overflow: "hidden",
          }}
        >
          <span
            style={{
              width: 34,
              height: 34,
              flex: "0 0 34px",
              borderRadius: 10,
              display: "inline-flex",
              alignItems: "center",
              justifyContent: "center",
              fontSize: 18,
              color: "#fff",
              background: `linear-gradient(135deg, ${token.colorPrimary} 0%, #13c2c2 100%)`,
              boxShadow: `0 6px 14px ${token.colorPrimary}55`,
            }}
          >
            <ThunderboltFilled />
          </span>
          {!collapsed && (
            <span style={{ minWidth: 0 }}>
              <span
                style={{
                  display: "block",
                  fontSize: 14,
                  fontWeight: 700,
                  color: token.colorText,
                  whiteSpace: "nowrap",
                }}
              >
                {systemInfo.system_name}
              </span>
              <span
                style={{
                  display: "block",
                  fontSize: 10,
                  letterSpacing: 1,
                  color: token.colorTextTertiary,
                  whiteSpace: "nowrap",
                }}
              >
                MANUFACTURING EXECUTION
              </span>
            </span>
          )}
        </Link>

        <Menu
          mode="inline"
          theme={isDark ? "dark" : "light"}
          items={toMenuItems(navigation)}
          selectedKeys={[getSelectedKey(pathname)]}
          openKeys={openKeys}
          onOpenChange={setOpenKeys}
          onClick={({ key }) => {
            if (key.startsWith("/")) router.push(key);
          }}
          style={{
            borderInlineEnd: "none",
            paddingTop: 8,
            background: "transparent",
          }}
        />

        {!collapsed && (
          <div
            style={{
              position: "absolute",
              bottom: 0,
              left: 0,
              right: 0,
              padding: "14px 18px",
              borderTop: `1px solid ${token.colorBorderSecondary}`,
              display: "flex",
              alignItems: "center",
              gap: 8,
              fontSize: 11,
              color: token.colorTextTertiary,
            }}
          >
            <span
              style={{
                width: 6,
                height: 6,
                borderRadius: "50%",
                background: token.colorSuccess,
                boxShadow: `0 0 0 3px ${token.colorSuccessBg}`,
              }}
            />
            系统运行正常 · v1.0.0
          </div>
        )}
      </Sider>

      <Layout style={{ background: token.colorBgLayout }}>
        <Header
          style={{
            height: 64,
            lineHeight: "64px",
            padding: "0 20px",
            background: token.colorBgContainer,
            borderBottom: `1px solid ${token.colorBorderSecondary}`,
            display: "flex",
            alignItems: "center",
            justifyContent: "space-between",
            position: "sticky",
            top: 0,
            zIndex: 20,
          }}
        >
          <Space size={12} align="center">
            <Button
              type="text"
              icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
              onClick={() => setCollapsed((prev) => !prev)}
              aria-label={collapsed ? "展开菜单" : "收起菜单"}
            />
            {screens.md && (
              <Breadcrumb
                items={breadcrumbItems}
                style={{ lineHeight: "64px", fontSize: 13 }}
              />
            )}
          </Space>

          <Space size={6} align="center">
            <Avatar
              size={30}
              style={{ background: token.colorPrimary, flex: "0 0 30px" }}
            >
              {user.user_name.slice(-2)}
            </Avatar>
            {screens.md && (
              <div style={{ lineHeight: 1.25, marginRight: 4 }}>
                <div
                  style={{
                    fontSize: 13,
                    fontWeight: 600,
                    color: token.colorText,
                  }}
                >
                  {user.user_name}
                </div>
                <div style={{ fontSize: 11, color: token.colorTextSecondary }}>
                  {user.introduction}
                </div>
              </div>
            )}
            <Divider
              orientation="vertical"
              style={{ height: 22, margin: "0 2px" }}
            />
            <Tooltip title="主题设置">
              <Button
                type="text"
                icon={<SettingOutlined />}
                onClick={() => setSettingsOpen(true)}
              >
                {screens.lg ? "设置" : null}
              </Button>
            </Tooltip>
            <Button
              type="text"
              danger
              icon={<LogoutOutlined />}
              onClick={handleLogout}
            >
              {screens.lg ? "退出" : null}
            </Button>
          </Space>
        </Header>

        <Content style={{ padding: screens.lg ? "16px" : "12px" }}>
          <div style={{ maxWidth: 1600, margin: "0 auto" }}>{children}</div>
        </Content>
      </Layout>

      <ThemeSettingsDrawer
        open={settingsOpen}
        onClose={() => setSettingsOpen(false)}
      />
    </Layout>
  );
}
