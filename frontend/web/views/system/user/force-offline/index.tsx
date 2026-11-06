"use client";

import { App as AntdApp, Button, Modal } from "antd";
import { MinusCircleOutlined, ExclamationCircleOutlined } from "@ant-design/icons";
import { SysUser } from "@/lib/types/system/user";
import { userOffline } from "@/lib/request/system/user";

interface IProps {
    record: SysUser;
}

function ForceOffline({ record }: IProps) {
    const { message } = AntdApp.useApp();
    const [modal, contextHolder] = Modal.useModal();

    const confirm = () => {
        if (!record?.user_name) {
            return;
        }
        modal.confirm({
            title: "系统提示",
            icon: <ExclamationCircleOutlined />,
            content: `确认强制下线用户 ⌈${record?.user_name}⌋ 吗？`,
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
                // 强制下线用户
                const res = await userOffline(record.id);
                if (res.data?.success) {
                    message.success("强制下线成功");
                } else {
                    message.error("强制下线失败");
                }
            },
        });
    };

    return (
        <>
            {contextHolder}
            <Button type="link" size="small" danger icon={<MinusCircleOutlined />} onClick={confirm}>
                强制下线
            </Button>
        </>
    );
}
export default ForceOffline;
