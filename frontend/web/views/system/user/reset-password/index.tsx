"use client";

import { App as AntdApp, Button, Modal } from "antd";
import { LockOutlined, ExclamationCircleOutlined } from "@ant-design/icons";
import { SysUser } from "@/lib/types/system/user";
import { resetPassword } from "@/lib/request/system/user";

interface IProps {
    record: SysUser;
}

function ResetPassword({ record }: IProps) {
    const { message } = AntdApp.useApp();
    const [modal, contextHolder] = Modal.useModal();

    const confirm = () => {
        if (!record?.user_name) {
            return;
        }
        modal.confirm({
            title: "系统提示",
            icon: <ExclamationCircleOutlined />,
            content: `确认重置用户 ⌈${record?.user_name}⌋ 的密码吗？`,
            okText: "确认",
            okButtonProps: {
                danger: true,
                size: "large",
            },
            cancelText: "取消",
            cancelButtonProps: {
                size: "large",
            },
            onOk: async () => {
                // 重置用户密码
                const res = await resetPassword(record.id);
                if (res.data?.success) {
                    message.success("重置密码成功");
                } else {
                    message.error("重置密码失败");
                }
            },
        });
    };

    return (
        <>
            {contextHolder}
            <Button type="link" size="small" danger icon={<LockOutlined />} onClick={confirm}>
                重置密码
            </Button>
        </>
    );
}
export default ResetPassword;
