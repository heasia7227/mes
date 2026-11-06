"use client";

import { App as AntdApp, Button, Modal } from "antd";
import { PauseCircleOutlined, PlayCircleOutlined, ExclamationCircleOutlined } from "@ant-design/icons";
import { SysUser } from "@/lib/types/system/user";
import { updateUserStatus } from "@/lib/request/system/user";

interface IProps {
    record: SysUser;
    refresh: () => void;
}

function ChangeStatus({ record, refresh }: IProps) {
    const { message } = AntdApp.useApp();
    const [modal, contextHolder] = Modal.useModal();

    const confirm = () => {
        if (!record?.user_name) {
            return;
        }
        modal.confirm({
            title: "系统提示",
            icon: <ExclamationCircleOutlined />,
            content: `确认${record?.status !== "enable" ? "启用" : "禁用"}用户 ⌈${record?.user_name}⌋ 吗？`,
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
                // 更新用户状态
                const res = await updateUserStatus(record.id, record?.status === "enable" ? "disable" : "enable");
                if (res.data?.success) {
                    message.success("操作成功");
                    refresh();
                } else {
                    message.error("操作失败");
                }
            },
        });
    };

    return (
        <>
            {contextHolder}
            <Button
                type="link"
                size="small"
                danger={record?.status === "enable"}
                icon={record?.status !== "enable" ? <PlayCircleOutlined /> : <PauseCircleOutlined />}
                onClick={confirm}
            >
                {record?.status !== "enable" ? "启用" : "禁用"}
            </Button>
        </>
    );
}
export default ChangeStatus;
