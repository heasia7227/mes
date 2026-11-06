"use client";

import { useState } from "react";
import { Card, Button, Space, Tag, Input, Table, Empty } from "antd";
import { BookOutlined, SearchOutlined, ReloadOutlined } from "@ant-design/icons";
import { SysDictionaryCatalog } from "@/lib/types/system/dictionary-catalog";
import DictionaryDataAdd from "./add";

interface IProps {
    selectedCatalog?: SysDictionaryCatalog;
}

const DictionaryDataList = ({ selectedCatalog }: IProps) => {
    const [keyword, setKeyword] = useState<string>("");
    const [loading, setLoading] = useState<boolean>(false);

    return (
        <>
            <Card size="small" className="main-card">
                <div
                    style={{
                        marginBottom: 16,
                    }}
                >
                    <div>
                        <Space size={8}>
                            <BookOutlined />
                            <span style={{ fontSize: 16, fontWeight: 500 }}>字典数据</span>
                            {selectedCatalog && <Tag color="blue">{selectedCatalog.catalog_name}</Tag>}
                        </Space>
                    </div>
                </div>

                {selectedCatalog ? (
                    <>
                        <div style={{ display: "flex", justifyContent: "space-between" }}>
                            <div>
                                <Input
                                    allowClear
                                    size="large"
                                    prefix={<SearchOutlined style={{ color: "#a0aab8" }} />}
                                    placeholder="搜索字典名称 / 字典值"
                                    value={keyword}
                                    onChange={(event) => setKeyword(event.target.value)}
                                    style={{ width: 260, marginBottom: 12 }}
                                />
                            </div>
                            <div>
                                <Space>
                                    <Button icon={<ReloadOutlined />} size="large">
                                        刷新
                                    </Button>
                                    <DictionaryDataAdd />
                                </Space>
                            </div>
                        </div>
                        <Table
                            rowKey="id"
                            dataSource={[]}
                            loading={loading}
                            scroll={{ x: "max-content" }}
                            pagination={{
                                pageSize: 10,
                                size: "large",
                                showSizeChanger: true,
                                showTotal: (total) => `共 ${total} 条`,
                            }}
                            columns={[
                                {
                                    title: "字典名称",
                                    dataIndex: "dictionary_name",
                                    width: 150,
                                },
                                {
                                    title: "字典值",
                                    dataIndex: "dictionary_value",
                                    width: 110,
                                },
                                { title: "目录编码", dataIndex: "catalog_code", width: 170 },
                                { title: "排序", dataIndex: "sort", width: 80 },
                                {
                                    title: "创建时间",
                                    dataIndex: "create_time",
                                    width: 170,
                                },
                                {
                                    title: "操作",
                                    key: "actions",
                                    width: 150,
                                    fixed: "right",
                                    render: (_, row) => (
                                        <Space size={0}>
                                            <DictionaryDataAdd isEdit={true} record={row} />
                                            <Button type="link" size="small" danger>
                                                删除
                                            </Button>
                                        </Space>
                                    ),
                                },
                            ]}
                        />
                    </>
                ) : (
                    <div style={{ height: "calc(100vh - 165px); position: relative" }}>
                        <Empty
                            style={{
                                position: "absolute",
                                top: "50%",
                                left: "50%",
                                transform: "translate(-50%, -50%)",
                            }}
                            description="请选择左侧字典目录查看数据"
                        />
                    </div>
                )}
            </Card>
        </>
    );
};

export default DictionaryDataList;
