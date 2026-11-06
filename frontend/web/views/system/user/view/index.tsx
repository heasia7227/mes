"use client";

import { useState } from "react";
import { Button, Modal, Descriptions } from "antd";
import { EyeOutlined } from "@ant-design/icons";
import { SysUser } from "@/lib/types/system/user";

interface IProps {
    record: SysUser;
}

function UserView({ record }: IProps) {
    const [visible, setVisible] = useState<boolean>(false);

    return (
        <>
            <Button
                type="link"
                size="small"
                icon={<EyeOutlined />}
                onClick={() => {
                    console.log("查看用户详情", record);
                    setVisible(true);
                }}
            >
                查看
            </Button>

            <Modal
                title="用户详情"
                open={visible}
                onCancel={() => setVisible(false)}
                footer={null}
                width={700}
                style={{ top: 10 }}
                mask={{
                    closable: false,
                }}
            >
                <Descriptions
                    bordered
                    column={2}
                    styles={{
                        label: {
                            textAlign: "center",
                        },
                    }}
                >
                    <Descriptions.Item label="姓名">{record.user_name}</Descriptions.Item>
                    <Descriptions.Item label="部门">{record.dept_name}</Descriptions.Item>
                    <Descriptions.Item label="岗位">{record.post_name}</Descriptions.Item>
                    <Descriptions.Item label="职级">{record.rank_name}</Descriptions.Item>
                    <Descriptions.Item label="登录账号">{record.login_name}</Descriptions.Item>
                    <Descriptions.Item label="性别">{record.sex}</Descriptions.Item>
                    <Descriptions.Item label="岗位编码">{record.post_code}</Descriptions.Item>
                    <Descriptions.Item label="手机号">{record.phone}</Descriptions.Item>
                    <Descriptions.Item label="邮箱">{record.email}</Descriptions.Item>
                    <Descriptions.Item label="状态">{record.status}</Descriptions.Item>
                    <Descriptions.Item label="简介" span={2}>
                        {record.introduction}
                    </Descriptions.Item>
                    <Descriptions.Item label="创建人">{record.create_by}</Descriptions.Item>
                    <Descriptions.Item label="创建时间">{record.create_time}</Descriptions.Item>
                    <Descriptions.Item label="更新人">{record.update_by}</Descriptions.Item>
                    <Descriptions.Item label="更新时间">{record.update_time}</Descriptions.Item>
                </Descriptions>
            </Modal>
        </>
    );
}

export default UserView;
