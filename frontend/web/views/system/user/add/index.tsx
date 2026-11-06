"use client";

import { useState } from "react";
import { Button, Modal, Form, Input, Row, Col } from "antd";
import { PlusOutlined, EditOutlined } from "@ant-design/icons";
import { SysUser } from "@/lib/types/system/user";
import DeptSelect from "@/components/DeptSelect";
import PostSelect from "@/components/PostSelect";
import RankSelect from "@/components/RankSelect";

interface IProps {
    isEdit?: boolean;
    record?: SysUser;
    refresh?: () => void;
}

const AddUser = ({ isEdit = false, refresh }: IProps) => {
    const [visible, setVisible] = useState<boolean>(false);

    const handleOpen = () => {
        setVisible(true);
    };

    const handleCancel = () => {
        setVisible(false);
    };

    const handleSubmit = () => {
        setVisible(false);
    };

    return (
        <>
            {isEdit ? (
                <Button type="link" size="small" icon={<EditOutlined />} onClick={handleOpen}>
                    编辑
                </Button>
            ) : (
                <Button type="primary" size="large" icon={<PlusOutlined />} onClick={handleOpen}>
                    新增
                </Button>
            )}

            <Modal
                title={isEdit ? "编辑用户" : "新增用户"}
                open={visible}
                onCancel={() => setVisible(false)}
                width={700}
                okButtonProps={{
                    size: "large",
                }}
                cancelButtonProps={{
                    size: "large",
                }}
                style={{ top: 10 }}
                mask={{
                    closable: false,
                }}
            >
                <Form layout="vertical" size="large">
                    <Row gutter={12}>
                        <Col span={12}>
                            <Form.Item
                                label="用户名"
                                name="user_name"
                                rules={[{ required: true, message: "请输入用户名" }]}
                            >
                                <Input placeholder="请输入用户名" />
                            </Form.Item>
                        </Col>
                        <Col span={12}>
                            <Form.Item
                                label="部门"
                                name="dept_id"
                                rules={[{ required: true, message: "请选择所属部门" }]}
                            >
                                <DeptSelect />
                            </Form.Item>
                        </Col>
                        <Col span={12}>
                            <Form.Item
                                label="岗位"
                                name="post_id"
                                rules={[{ required: true, message: "请选择所属岗位" }]}
                            >
                                <PostSelect />
                            </Form.Item>
                        </Col>
                        <Col span={12}>
                            <Form.Item
                                label="职级"
                                name="rank_id"
                                rules={[{ required: true, message: "请选择所属职级" }]}
                            >
                                <RankSelect />
                            </Form.Item>
                        </Col>
                        <Col span={12}>
                            <Form.Item
                                label="登录账号"
                                name="login_name"
                                rules={[{ required: true, message: "请输入登录账号" }]}
                            >
                                <Input placeholder="请输入登录账号" />
                            </Form.Item>
                        </Col>
                        <Col span={12}>
                            <Form.Item
                                label="登录密码"
                                name="password"
                                rules={[{ required: true, message: "请输入登录密码" }]}
                            >
                                <Input.Password placeholder="请输入登录密码" disabled={isEdit} />
                            </Form.Item>
                        </Col>
                        <Col span={12}>
                            <Form.Item
                                label="工号"
                                name="post_code"
                                rules={[{ required: true, message: "请输入工号" }]}
                            >
                                <Input placeholder="请输入工号" />
                            </Form.Item>
                        </Col>
                        <Col span={12}>
                            <Form.Item label="性别" name="sex" rules={[{ required: true, message: "请选择性别" }]}>
                                <Input placeholder="请选择性别" />
                            </Form.Item>
                        </Col>
                        <Col span={12}>
                            <Form.Item
                                label="手机号"
                                name="phone"
                                rules={[
                                    { required: true, message: "请输入手机号" },
                                    { pattern: /^1[3-9]\d{9}$/, message: "请输入正确的手机号" },
                                ]}
                            >
                                <Input placeholder="请输入手机号" />
                            </Form.Item>
                        </Col>
                        <Col span={12}>
                            <Form.Item
                                label="邮箱"
                                name="email"
                                rules={[
                                    { required: true, message: "请输入邮箱" },
                                    { type: "email", message: "请输入正确的邮箱" },
                                ]}
                            >
                                <Input placeholder="请输入邮箱" />
                            </Form.Item>
                        </Col>
                        <Col span={24}>
                            <Form.Item
                                label="个人简介"
                                name="introduction"
                                rules={[{ required: true, message: "请输入个人简介" }]}
                            >
                                <Input.TextArea placeholder="请输入个人简介" />
                            </Form.Item>
                        </Col>
                    </Row>
                </Form>
            </Modal>
        </>
    );
};

export default AddUser;
