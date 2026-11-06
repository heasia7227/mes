"use client";

import { App as AntdApp, Button, Modal } from "antd";
import { DeleteOutlined, ExclamationCircleOutlined } from "@ant-design/icons";
import { SysUser } from "@/lib/types/system/user";
import { deleteUser } from "@/lib/request/system/user";

interface IProps {
    record: SysUser;
}

function DeleteUser({ record }: IProps) {
    const { message } = AntdApp.useApp();
    const [modal, contextHolder] = Modal.useModal();

    const confirm = () => {
        if (!record?.user_name) {
            return;
        }
        if (record?.login_name === "admin") {
            modal.error({
                title: "系统提示",
                content: "管理员账号不能删除",
            });
            return;
        }
        modal.confirm({
            title: "系统提示",
            icon: <ExclamationCircleOutlined />,
            content: `确认删除用户 ⌈${record?.user_name}⌋ 吗？`,
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
                // 删除用户
                const res = await deleteUser(record.id);
                if (res.data?.success) {
                    message.success("删除成功");
                } else {
                    message.error("删除失败");
                }
            },
        });
    };

    return (
        <>
            {contextHolder}
            <Button type="link" size="small" danger icon={<DeleteOutlined />} onClick={confirm}>
                删除
            </Button>
        </>
    );
}

export default DeleteUser;
