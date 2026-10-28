"use client";

import type { MenuProps } from "antd";
import {
  ApartmentOutlined,
  BookOutlined,
  DashboardOutlined,
  FileTextOutlined,
  IdcardOutlined,
  InfoCircleOutlined,
  LoginOutlined,
  MenuOutlined,
  SettingOutlined,
  TeamOutlined,
  ToolOutlined,
  TrophyOutlined,
  UserOutlined,
} from "@ant-design/icons";
import type { ReactNode } from "react";

export interface NavItem {
  key: string;
  label: string;
  icon?: ReactNode;
  path?: string;
  children?: NavItem[];
}

export const navigation: NavItem[] = [
  {
    key: "/dashboard",
    label: "工作台",
    icon: <DashboardOutlined />,
    path: "/dashboard",
  },
  {
    key: "/system",
    label: "系统管理",
    icon: <SettingOutlined />,
    children: [
      {
        key: "/system/user",
        label: "用户管理",
        icon: <UserOutlined />,
        path: "/system/user",
      },
      {
        key: "/system/role",
        label: "角色管理",
        icon: <TeamOutlined />,
        path: "/system/role",
      },
      {
        key: "/system/menu",
        label: "菜单管理",
        icon: <MenuOutlined />,
        path: "/system/menu",
      },
      {
        key: "/system/dept",
        label: "部门管理",
        icon: <ApartmentOutlined />,
        path: "/system/dept",
      },
      {
        key: "/system/post",
        label: "岗位管理",
        icon: <IdcardOutlined />,
        path: "/system/post",
      },
      {
        key: "/system/rank",
        label: "职级管理",
        icon: <TrophyOutlined />,
        path: "/system/rank",
      },
      {
        key: "/system/dictionary",
        label: "字典管理",
        icon: <BookOutlined />,
        path: "/system/dictionary",
      },
      {
        key: "/system/config",
        label: "参数配置",
        icon: <ToolOutlined />,
        path: "/system/config",
      },
      {
        key: "/system/info",
        label: "系统信息",
        icon: <InfoCircleOutlined />,
        path: "/system/info",
      },
      {
        key: "/system/log",
        label: "操作日志",
        icon: <FileTextOutlined />,
        path: "/system/log",
      },
      {
        key: "/system/login-log",
        label: "登录日志",
        icon: <LoginOutlined />,
        path: "/system/login-log",
      },
    ],
  },
];

export function toMenuItems(items: NavItem[]): NonNullable<MenuProps["items"]> {
  return items.map((item) => ({
    key: item.key,
    icon: item.icon,
    label: item.label,
    children: item.children ? toMenuItems(item.children) : undefined,
  }));
}

export function findTrail(
  pathname: string,
  items: NavItem[] = navigation,
): NavItem[] {
  for (const item of items) {
    if (item.path === pathname) return [item];
    if (item.children) {
      const childTrail = findTrail(pathname, item.children);
      if (childTrail.length) return [item, ...childTrail];
    }
  }
  return [];
}

export function getSelectedKey(
  pathname: string,
  items: NavItem[] = navigation,
): string {
  const leaves: string[] = [];
  const walk = (list: NavItem[]) => {
    list.forEach((item) => {
      if (item.path) leaves.push(item.path);
      if (item.children) walk(item.children);
    });
  };
  walk(items);

  return (
    leaves
      .filter((path) => pathname === path || pathname.startsWith(`${path}/`))
      .sort((a, b) => b.length - a.length)[0] ?? pathname
  );
}

export function getOpenKeys(
  pathname: string,
  items: NavItem[] = navigation,
): string[] {
  return findTrail(pathname, items)
    .filter((item) => item.children)
    .map((item) => item.key);
}
