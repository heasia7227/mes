"use client";

import ResourcePage from "@/components/ResourcePage";
import { posts, users } from "@/lib/mock-data";
import { SEX_MAP, STATUS_MAP } from "@/lib/status";

const postOptions = posts.map((item) => ({
  label: `${item.post_name}（${item.post_code}）`,
  value: item.post_code,
}));

const statusOptions = [
  { label: "启用", value: "1" },
  { label: "禁用", value: "0" },
];

export default function UserPage() {
  return (
    <ResourcePage
      title="用户管理"
      description="维护系统账号、岗位与状态，对应数据表 t_sys_user。"
      data={users}
      displayKey="user_name"
      searchKeys={["user_name", "login_name", "phone", "email"]}
      searchPlaceholder="搜索姓名 / 账号 / 手机号 / 邮箱"
      statusField="status"
      statusOptions={statusOptions}
      columns={[
        { key: "user_name", title: "姓名", width: 110 },
        { key: "login_name", title: "登录账号", width: 130 },
        {
          key: "sex",
          title: "性别",
          width: 80,
          valueType: "status",
          statusMap: SEX_MAP,
        },
        { key: "phone", title: "手机号", width: 140 },
        { key: "email", title: "邮箱", width: 200, ellipsis: true },
        { key: "post_code", title: "岗位编码", width: 120 },
        {
          key: "status",
          title: "状态",
          width: 90,
          valueType: "status",
          statusMap: STATUS_MAP,
        },
        {
          key: "create_time",
          title: "创建时间",
          width: 170,
          valueType: "date",
        },
      ]}
      fieldCol={2}
      fields={[
        { key: "user_name", label: "姓名", required: true },
        {
          key: "dept_id",
          label: "所属部门",
          type: "select",
          options: postOptions,
        },
        {
          key: "post_id",
          label: "所属岗位",
          type: "select",
          options: postOptions,
        },
        {
          key: "rank_id",
          label: "所属职级",
          type: "select",
          options: postOptions,
        },
        { key: "login_name", label: "登录账号", required: true },
        {
          key: "password",
          label: "密码",
          required: true,
          placeholder: "新增必填，编辑时不修改可留空",
        },
        { key: "post_code", label: "岗位编码" },
        {
          key: "sex",
          label: "性别",
          type: "select",
          options: [
            { label: "男", value: "0" },
            { label: "女", value: "1" },
          ],
        },
        { key: "phone", label: "手机号" },
        { key: "email", label: "邮箱" },
        {
          key: "introduction",
          label: "个人简介",
          type: "textarea",
          combineCol: true,
        },
      ]}
    />
  );
}
