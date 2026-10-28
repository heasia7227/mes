export interface StatusMeta {
  text: string;
  color: string;
}

export type StatusMap = Record<string, StatusMeta>;

/** 通用启用 / 禁用 */
export const STATUS_MAP: StatusMap = {
  "1": { text: "启用", color: "success" },
  "0": { text: "禁用", color: "default" },
};

/** 正常 / 停用 */
export const NORMAL_STATUS_MAP: StatusMap = {
  "1": { text: "正常", color: "success" },
  "0": { text: "停用", color: "default" },
};

export const SEX_MAP: StatusMap = {
  "0": { text: "男", color: "blue" },
  "1": { text: "女", color: "magenta" },
};

export const MENU_TYPE_MAP: StatusMap = {
  M: { text: "目录", color: "blue" },
  C: { text: "菜单", color: "green" },
  F: { text: "按钮", color: "orange" },
};

export const DEPT_TYPE_MAP: StatusMap = {
  公司: { text: "公司", color: "geekblue" },
  部门: { text: "部门", color: "cyan" },
  班组: { text: "班组", color: "purple" },
};

export const YES_NO_MAP: StatusMap = {
  "1": { text: "是", color: "red" },
  "0": { text: "否", color: "green" },
};

export const LOG_STATUS_MAP: StatusMap = {
  "1": { text: "成功", color: "success" },
  "0": { text: "失败", color: "error" },
};

export const CONFIG_TYPE_MAP: StatusMap = {
  Y: { text: "系统内置", color: "blue" },
  N: { text: "自定义", color: "default" },
};
