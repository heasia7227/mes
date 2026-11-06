"use client";

import { useState, useMemo, useEffect } from "react";
import { Card, Space, Input, Button, Select, Table, Form, Dropdown } from "antd";
import { SearchOutlined, SyncOutlined, MoreOutlined } from "@ant-design/icons";
import { ColumnsType } from "antd/es/table";
import { getUserList } from "@/lib/request/system/user";
import { SysUser } from "@/lib/types/system/user";
import UserView from "./view";
import AddUser from "./add";
import DeleteUser from "./delete";
import ChangeStatus from "./change-status";
import ResetPassword from "./reset-password";
import ForceOffline from "./force-offline";

const UserList = () => {
    const [loading, setLoading] = useState<boolean>(false);
    const [dataSource, setDataSource] = useState<SysUser[]>([]);
    const [queryParams, setQueryParams] = useState<any>({ page_num: 1, page_size: 10 });
    const [form] = Form.useForm();

    useEffect(() => {
        fetchUserList(queryParams);
    }, []);

    const refresh = () => fetchUserList(queryParams);

    const onSearch = async () => {
        const values = await form.getFieldsValue();
        fetchUserList({ ...queryParams, ...values });
    };

    const onResetSearch = () => {
        form.resetFields();
        fetchUserList({ page_num: 1, page_size: 10 });
    };

    const onPageChange = (page: number, pageSize: number) => {
        const query = {
            ...queryParams,
            page_num: page,
            page_size: pageSize,
        };
        fetchUserList(query);
    };

    const fetchUserList = async (query: any) => {
        setLoading(true);
        const res = await getUserList(query);
        setDataSource(res.data?.list || []);
        if (res.data?.pagination) {
            setQueryParams({
                ...query,
                ...res.data.pagination,
            });
        }
        setLoading(false);
    };

    const columns: ColumnsType<any> = useMemo(
        () => [
            {
                title: "姓名",
                dataIndex: "user_name",
                key: "user_name",
                width: 110,
                align: "center",
                fixed: "left",
            },
            {
                title: "登录账号",
                dataIndex: "login_name",
                key: "login_name",
                width: 150,
                align: "center",
            },
            {
                title: "性别",
                dataIndex: "sex",
                key: "sex",
                width: 80,
                align: "center",
            },
            {
                title: "手机号",
                dataIndex: "phone",
                key: "phone",
                width: 100,
                align: "center",
            },
            {
                title: "邮箱",
                dataIndex: "email",
                key: "email",
                width: 150,
                align: "center",
            },
            {
                title: "岗位编码",
                dataIndex: "post_code",
                key: "post_code",
                width: 150,
                align: "center",
            },
            {
                title: "状态",
                dataIndex: "status",
                key: "status",
                width: 100,
                align: "center",
            },
            {
                title: "创建人",
                dataIndex: "create_by",
                key: "create_by",
                width: 100,
                align: "center",
            },
            {
                title: "创建时间",
                dataIndex: "create_time",
                key: "create_time",
                width: 180,
                align: "center",
            },
            {
                title: "修改人",
                dataIndex: "update_by",
                key: "update_by",
                width: 100,
                align: "center",
            },
            {
                title: "修改时间",
                dataIndex: "update_time",
                key: "update_time",
                width: 180,
                align: "center",
            },
            {
                title: "操作",
                key: "__actions",
                fixed: "right",
                render: (_, record: SysUser) => (
                    <Space size={0}>
                        <UserView record={record} />
                        <AddUser isEdit={true} record={record} refresh={refresh} />
                        <Dropdown
                            placement="bottomRight"
                            menu={{
                                items: [
                                    {
                                        label: <ResetPassword record={record} />,
                                        key: "reset-password",
                                    },
                                    {
                                        label: <ForceOffline record={record} />,
                                        key: "force-offline",
                                    },
                                    {
                                        label: <ChangeStatus record={record} refresh={refresh} />,
                                        key: "change-status",
                                    },
                                    {
                                        label: <DeleteUser record={record} />,
                                        key: "delete",
                                    },
                                ],
                            }}
                        >
                            <a onClick={(e) => e.preventDefault()}>
                                <Space>
                                    <MoreOutlined />
                                    更多
                                </Space>
                            </a>
                        </Dropdown>
                    </Space>
                ),
            },
        ],
        [],
    );

    return (
        <>
            <Card className="main-card">
                <div className="main-search-section">
                    <Form form={form}>
                        <Space>
                            <Form.Item name="keyword" noStyle>
                                <Input allowClear size="large" placeholder="模糊搜索" />
                            </Form.Item>
                            <Form.Item name="status" noStyle>
                                <Select allowClear size="large" placeholder="全部状态" />
                            </Form.Item>
                            <Button type="primary" size="large" icon={<SearchOutlined />} onClick={onSearch}>
                                查询
                            </Button>
                            <Button size="large" icon={<SyncOutlined />} onClick={onResetSearch}>
                                重置
                            </Button>
                        </Space>
                    </Form>
                    <Space>
                        <AddUser refresh={refresh} />
                    </Space>
                </div>

                <Table
                    className="main-table-section"
                    rowKey={"id"}
                    columns={columns}
                    dataSource={dataSource}
                    loading={loading}
                    bordered
                    scroll={{ x: "max-content", y: "auto" }}
                    pagination={{
                        current: queryParams.page_num,
                        pageSize: queryParams.page_size,
                        total: queryParams.total,
                        size: "large",
                        showSizeChanger: true,
                        showTotal: (t) => `共 ${t} 条`,
                        onChange: onPageChange,
                    }}
                />
            </Card>
        </>
    );
};

export default UserList;
