/**
 * 与 backend/common/models/system 下的 Go 模型保持一致
 */

export interface BaseEntity {
  id: string;
  create_by?: string;
  create_time?: string;
  update_by?: string;
  update_time?: string;
  delete_flag?: number;
}

/** t_sys_user */
export interface SysUser extends BaseEntity {
  user_name: string;
  login_name: string;
  password?: string;
  post_code?: string;
  sex?: string;
  phone?: string;
  email?: string;
  avatar_url?: string;
  introduction?: string;
  status?: string;
}

/** t_sys_role */
export interface SysRole extends BaseEntity {
  title: string;
  code: string;
  description?: string;
  status?: string;
}

/** t_sys_menu */
export interface SysMenu extends BaseEntity {
  title: string;
  icon?: string;
  parent_id?: string;
  ranking?: number;
  menu_type?: string;
  design_type?: string;
  menu_path?: string;
  permission?: string;
  page_type?: string;
  status?: string;
}

/** t_sys_dept */
export interface SysDept extends BaseEntity {
  parent_id?: string | null;
  dept_name: string;
  dept_code: string;
  dept_full_name?: string;
  dept_type?: string;
  sort_no?: number;
  remark?: string;
  status?: number;
  id_path?: string;
  name_path?: string;
}

/** t_sys_post */
export interface SysPost extends BaseEntity {
  post_name: string;
  post_code: string;
  post_status?: number;
  sort_no?: number;
  remark?: string;
}

/** t_sys_rank */
export interface SysRank extends BaseEntity {
  rank_name: string;
  rank_status?: number;
  sort_no?: number;
  remark?: string;
}

/** t_sys_dictionary_catalog */
export interface SysDictionaryCatalog extends BaseEntity {
  catalog_name: string;
  catalog_code: string;
  parent_id?: string;
}

/** t_sys_dictionary */
export interface SysDictionary extends BaseEntity {
  catalog_id: string;
  catalog_code: string;
  dictionary_value: string;
  dictionary_name: string;
  sort?: number;
  association_dict?: string;
  association_dict_name?: string;
  association_dict_value?: string;
}

/** t_sys_config */
export interface SysConfig {
  config_id: number;
  config_name: string;
  config_key: string;
  config_value: string;
  config_type?: string;
  create_by?: string;
  create_time?: string;
  update_by?: string;
  update_time?: string;
  remark?: string;
  delete_flag?: number;
}

/** t_sys_info */
export interface SysInfo {
  id: string;
  system_name: string;
  system_logo?: string;
  login_background?: string;
  footer_text?: string;
  is_enabled?: number;
  create_by?: string;
  create_time?: string;
  update_by?: string;
  update_time?: string;
  delete_flag?: number;
}

/** t_sys_log */
export interface SysLog {
  id: string;
  ip?: string;
  log_content?: string;
  is_exception?: string;
  operate_type?: string;
  operate_module?: string;
  operate_by?: string;
  operate_time?: string;
  status?: string;
  duration?: number;
  description?: string;
}

/** t_sys_login_log */
export interface SysLoginLog extends BaseEntity {
  login_name: string;
  login_ip?: string;
  login_location?: string;
  browser?: string;
  os?: string;
  login_status?: number;
  message?: string;
  login_time?: string;
}
