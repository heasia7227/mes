"use client";

import ResourcePage from "@/components/ResourcePage";
import { menus } from "@/lib/mock-data";
import { MENU_TYPE_MAP, STATUS_MAP } from "@/lib/status";
import { useMemo, useState } from "react";
import type { SysMenu } from "@/lib/types";

type MenuRow = SysMenu & Record<string, any>;
type MenuTreeNode = MenuRow & { key: string; children?: MenuTreeNode[] };

function buildTree(list: MenuRow[], parentId = ""): MenuTreeNode[] {
  return list
    .filter((item) => (item.parent_id ?? "") === parentId)
    .sort((a, b) => (a.ranking ?? 0) - (b.ranking ?? 0))
    .map((item) => {
      const children = buildTree(list, item.id);
      return {
        ...item,
        key: item.id,
        children: children.length ? children : undefined,
      };
    });
}

export default function LoginLogPage() {
  const [expandedKeys, setExpandedKeys] = useState<string[]>([]);

  const treeData = useMemo(() => {
    return buildTree(menus as MenuRow[]);
  }, []);

  return (
    <ResourcePage
      sectionClassName="main-tree-table-section"
      title="菜单管理"
      description="管理系统菜单，对应数据表 t_sys_menu。"
      data={treeData}
      pagination={false}
      expandable={{
        expandedRowKeys: expandedKeys,
        onExpandedRowsChange: (keys) => setExpandedKeys(keys.map(String)),
      }}
      fields={[
        { key: "title", label: "菜单名称", required: true },
        { key: "icon", label: "图标", required: true },
        { key: "menu_type", label: "类型", required: true },
        { key: "ranking", label: "排序", required: true },
        { key: "menu_path", label: "路由地址", required: true },
        { key: "permission", label: "权限标识", required: true },
      ]}
      displayKey="title"
      searchKeys={["title"]}
      searchPlaceholder="搜索菜单名称"
      statusField="status"
      statusOptions={[
        { label: "成功", value: "1" },
        { label: "失败", value: "0" },
      ]}
      columns={[
        { title: "菜单名称", key: "title", width: 220 },
        { title: "图标", key: "icon", width: 140 },
        {
          title: "类型",
          key: "menu_type",
          width: 90,
          valueType: "status",
          statusMap: MENU_TYPE_MAP,
        },
        { title: "排序", key: "ranking", width: 70 },
        { title: "路由地址", key: "menu_path", width: 220 },
        { title: "权限标识", key: "permission", width: 220 },
        {
          title: "状态",
          key: "status",
          width: 90,
          valueType: "status",
          statusMap: STATUS_MAP,
        },
        {
          title: "创建时间",
          key: "create_time",
          width: 170,
          valueType: "date",
        },
      ]}
    />
  );
}
