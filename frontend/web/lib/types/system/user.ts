import { BaseEntity } from "@/lib/types/types";

/** t_sys_user */
export interface SysUser extends BaseEntity {
    dept_id?: string;
    dept_name?: string;
    post_id?: string;
    post_name?: string;
    rank_id?: string;
    rank_name?: string;
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

/** 修改密码请求体 */
export interface ChangePasswordPayload {
    old_password: string;
    new_password: string;
}
