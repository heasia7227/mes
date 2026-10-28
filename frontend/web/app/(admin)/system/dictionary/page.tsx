"use client";

import {
  BookOutlined,
  DeleteOutlined,
  EditOutlined,
  FolderAddOutlined,
  PlusOutlined,
  ReloadOutlined,
  SearchOutlined,
} from "@ant-design/icons";
import {
  App as AntdApp,
  Button,
  Card,
  Col,
  Empty,
  Form,
  Input,
  InputNumber,
  Modal,
  Popconfirm,
  Row,
  Select,
  Space,
  Table,
  Tag,
  Tree,
  Typography,
} from "antd";
import dayjs from "dayjs";
import { useMemo, useState } from "react";
import { dictionaries, dictionaryCatalogs } from "@/lib/mock-data";
import type { SysDictionary, SysDictionaryCatalog } from "@/lib/types";

type CatalogRow = SysDictionaryCatalog & Record<string, any>;
type DictRow = SysDictionary & Record<string, any>;

interface TreeNode {
  key: string;
  title: string;
  children?: TreeNode[];
}

function buildTree(list: CatalogRow[], parentId = ""): TreeNode[] {
  return list
    .filter((item) => (item.parent_id ?? "") === parentId)
    .map((item) => ({
      key: item.id,
      title: item.catalog_name,
      children: buildTree(list, item.id),
    }))
    .filter((node) => node.children?.length || true);
}

function now() {
  return dayjs().format("YYYY-MM-DD HH:mm:ss");
}

export default function DictionaryPage() {
  const { message } = AntdApp.useApp();

  const [catalogs, setCatalogs] = useState<CatalogRow[]>(
    dictionaryCatalogs as CatalogRow[],
  );
  const [items, setItems] = useState<DictRow[]>(dictionaries as DictRow[]);
  const [selectedCatalogId, setSelectedCatalogId] = useState<
    string | undefined
  >();
  const [keyword, setKeyword] = useState("");

  const [catalogModalOpen, setCatalogModalOpen] = useState(false);
  const [editingCatalog, setEditingCatalog] = useState<CatalogRow | null>(null);
  const [catalogForm] = Form.useForm();

  const [dictModalOpen, setDictModalOpen] = useState(false);
  const [editingDict, setEditingDict] = useState<DictRow | null>(null);
  const [dictForm] = Form.useForm();

  const treeData = useMemo(() => buildTree(catalogs), [catalogs]);

  const selectedCatalog = useMemo(
    () => catalogs.find((item) => item.id === selectedCatalogId),
    [catalogs, selectedCatalogId],
  );

  const filteredItems = useMemo(() => {
    const kw = keyword.trim().toLowerCase();
    return items.filter((row) => {
      const matchCatalog =
        !selectedCatalogId || row.catalog_id === selectedCatalogId;
      const matchKeyword =
        !kw ||
        [row.dictionary_name, row.dictionary_value, row.catalog_code].some(
          (value) =>
            String(value ?? "")
              .toLowerCase()
              .includes(kw),
        );
      return matchCatalog && matchKeyword;
    });
  }, [items, selectedCatalogId, keyword]);

  const catalogOptions = useMemo(
    () =>
      catalogs.map((item) => ({ label: item.catalog_name, value: item.id })),
    [catalogs],
  );

  /* ---------- 目录操作 ---------- */
  const openCatalogModal = (catalog?: CatalogRow) => {
    setEditingCatalog(catalog ?? null);
    catalogForm.resetFields();
    catalogForm.setFieldsValue(
      catalog
        ? {
            catalog_name: catalog.catalog_name,
            catalog_code: catalog.catalog_code,
            parent_id: catalog.parent_id || undefined,
          }
        : { parent_id: selectedCatalogId },
    );
    setCatalogModalOpen(true);
  };

  const handleCatalogSubmit = async () => {
    const values = await catalogForm.validateFields();
    if (editingCatalog) {
      setCatalogs((prev) =>
        prev.map((item) =>
          item.id === editingCatalog.id
            ? {
                ...item,
                ...values,
                parent_id: values.parent_id ?? "",
                update_time: now(),
              }
            : item,
        ),
      );
      message.success("目录已更新");
    } else {
      const id = `C${Date.now()}`;
      setCatalogs((prev) => [
        ...prev,
        {
          id,
          catalog_name: values.catalog_name,
          catalog_code: values.catalog_code,
          parent_id: values.parent_id ?? "",
          create_time: now(),
          delete_flag: 0,
        } as CatalogRow,
      ]);
      setSelectedCatalogId(id);
      message.success("目录已新增");
    }
    setCatalogModalOpen(false);
  };

  const handleCatalogDelete = () => {
    if (!selectedCatalog) return;
    setCatalogs((prev) =>
      prev.filter((item) => item.id !== selectedCatalog.id),
    );
    setItems((prev) =>
      prev.filter((item) => item.catalog_id !== selectedCatalog.id),
    );
    setSelectedCatalogId(undefined);
    message.success("目录已删除");
  };

  /* ---------- 字典数据操作 ---------- */
  const openDictModal = (row?: DictRow) => {
    if (!selectedCatalog) {
      message.warning("请先在左侧选择字典目录");
      return;
    }
    setEditingDict(row ?? null);
    dictForm.resetFields();
    dictForm.setFieldsValue(
      row
        ? {
            dictionary_name: row.dictionary_name,
            dictionary_value: row.dictionary_value,
            sort: row.sort,
            association_dict: row.association_dict,
            association_dict_name: row.association_dict_name,
            association_dict_value: row.association_dict_value,
          }
        : { sort: 1 },
    );
    setDictModalOpen(true);
  };

  const handleDictSubmit = async () => {
    const values = await dictForm.validateFields();
    if (!selectedCatalog) return;
    if (editingDict) {
      setItems((prev) =>
        prev.map((item) =>
          item.id === editingDict.id
            ? { ...item, ...values, update_time: now() }
            : item,
        ),
      );
      message.success("字典项已更新");
    } else {
      setItems((prev) => [
        {
          id: `DC${Date.now()}`,
          catalog_id: selectedCatalog.id,
          catalog_code: selectedCatalog.catalog_code,
          create_time: now(),
          delete_flag: 0,
          ...values,
        } as DictRow,
        ...prev,
      ]);
      message.success("字典项已新增");
    }
    setDictModalOpen(false);
  };

  const handleDictDelete = (row: DictRow) => {
    setItems((prev) => prev.filter((item) => item.id !== row.id));
    message.success("字典项已删除");
  };

  return (
    <div>
      <div style={{ marginBottom: 16 }}>
        <Typography.Title level={4} style={{ margin: 0 }}>
          字典管理
        </Typography.Title>
      </div>

      <Row gutter={16}>
        <Col xs={24} md={8} lg={7} xl={6}>
          <Card size="small" style={{ height: "calc(100vh - 142px)" }}>
            <div
              style={{
                display: "flex",
                justifyContent: "space-between",
                marginBottom: 16,
              }}
            >
              <div style={{ fontSize: 16, fontWeight: 500 }}>字典目录</div>
              <div>
                <Button
                  type="link"
                  size="small"
                  icon={<FolderAddOutlined />}
                  onClick={() => openCatalogModal()}
                >
                  新增
                </Button>
              </div>
            </div>
            <Tree
              blockNode
              defaultExpandAll
              showLine
              treeData={treeData}
              selectedKeys={selectedCatalogId ? [selectedCatalogId] : []}
              onSelect={(keys) =>
                setSelectedCatalogId((keys[0] as string) ?? undefined)
              }
              style={{ height: "calc(100vh - 250px)" }}
            />

            <div style={{ marginTop: 16, display: "flex", gap: 8 }}>
              <Button
                size="small"
                icon={<EditOutlined />}
                disabled={!selectedCatalog}
                onClick={() =>
                  selectedCatalog && openCatalogModal(selectedCatalog)
                }
                block
              >
                编辑
              </Button>
              <Popconfirm
                title="确认删除"
                description="删除目录将同时移除其下字典数据"
                okText="删除"
                cancelText="取消"
                okButtonProps={{ danger: true }}
                onConfirm={handleCatalogDelete}
              >
                <Button
                  size="small"
                  danger
                  icon={<DeleteOutlined />}
                  disabled={!selectedCatalog}
                  block
                >
                  删除
                </Button>
              </Popconfirm>
            </div>
          </Card>
        </Col>

        <Col xs={24} md={16} lg={17} xl={18}>
          <Card style={{ height: "calc(100vh - 142px)" }} size="small">
            <div
              style={{
                marginBottom: 16,
              }}
            >
              <div>
                <Space size={8}>
                  <BookOutlined />
                  <span style={{ fontSize: 16, fontWeight: 500 }}>
                    字典数据
                  </span>
                  {selectedCatalog && (
                    <Tag color="blue">{selectedCatalog.catalog_name}</Tag>
                  )}
                </Space>
              </div>
            </div>

            {selectedCatalogId && (
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
                    <Button
                      icon={<ReloadOutlined />}
                      size="large"
                      onClick={() => {
                        setItems(dictionaries as DictRow[]);
                        setCatalogs(dictionaryCatalogs as CatalogRow[]);
                        setKeyword("");
                        message.success("数据已刷新");
                      }}
                    >
                      刷新
                    </Button>
                    <Button
                      type="primary"
                      size="large"
                      icon={<PlusOutlined />}
                      onClick={() => openDictModal()}
                    >
                      新增
                    </Button>
                  </Space>
                </div>
              </div>
            )}

            {selectedCatalogId ? (
              <Table<DictRow>
                rowKey="id"
                dataSource={filteredItems}
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
                    render: (value: string) =>
                      value ? dayjs(value).format("YYYY-MM-DD HH:mm") : "-",
                  },
                  {
                    title: "操作",
                    key: "actions",
                    width: 150,
                    fixed: "right",
                    render: (_, row) => (
                      <Space size={0}>
                        <Button
                          type="link"
                          size="small"
                          icon={<EditOutlined />}
                          onClick={() => openDictModal(row)}
                        >
                          编辑
                        </Button>
                        <Popconfirm
                          title="确认删除"
                          description={`确定删除「${row.dictionary_name}」吗？`}
                          okText="删除"
                          cancelText="取消"
                          okButtonProps={{ danger: true }}
                          onConfirm={() => handleDictDelete(row)}
                        >
                          <Button
                            type="link"
                            size="small"
                            danger
                            icon={<DeleteOutlined />}
                          >
                            删除
                          </Button>
                        </Popconfirm>
                      </Space>
                    ),
                  },
                ]}
              />
            ) : (
              <Empty
                style={{ padding: "48px 0" }}
                description="请选择左侧字典目录查看数据"
              />
            )}
          </Card>
        </Col>
      </Row>

      <Modal
        title={editingCatalog ? "编辑字典目录" : "新增字典目录"}
        open={catalogModalOpen}
        okText="保存"
        cancelText="取消"
        forceRender
        onOk={handleCatalogSubmit}
        onCancel={() => setCatalogModalOpen(false)}
      >
        <Form form={catalogForm} layout="vertical" style={{ marginTop: 8 }}>
          <Form.Item
            name="catalog_name"
            label="目录名称"
            rules={[{ required: true, message: "请输入目录名称" }]}
          >
            <Input placeholder="请输入目录名称" />
          </Form.Item>
          <Form.Item
            name="catalog_code"
            label="目录编码"
            rules={[{ required: true, message: "请输入目录编码" }]}
          >
            <Input placeholder="例如 sys_user_sex" />
          </Form.Item>
          <Form.Item name="parent_id" label="上级目录">
            <Select
              allowClear
              placeholder="不选择则为顶级目录"
              options={catalogOptions.filter(
                (item) => item.value !== editingCatalog?.id,
              )}
            />
          </Form.Item>
        </Form>
      </Modal>

      <Modal
        title={editingDict ? "编辑字典项" : "新增字典项"}
        open={dictModalOpen}
        okText="保存"
        cancelText="取消"
        forceRender
        width={560}
        onOk={handleDictSubmit}
        onCancel={() => setDictModalOpen(false)}
      >
        <Form form={dictForm} layout="vertical" style={{ marginTop: 8 }}>
          <Form.Item
            name="dictionary_name"
            label="字典名称"
            rules={[{ required: true, message: "请输入字典名称" }]}
          >
            <Input placeholder="请输入字典名称" />
          </Form.Item>
          <Form.Item
            name="dictionary_value"
            label="字典值"
            rules={[{ required: true, message: "请输入字典值" }]}
          >
            <Input placeholder="请输入字典值" />
          </Form.Item>
          <Form.Item name="sort" label="排序">
            <InputNumber
              style={{ width: "100%" }}
              min={0}
              placeholder="请输入排序"
            />
          </Form.Item>
          <Form.Item name="association_dict" label="关联字典">
            <Input placeholder="选填" />
          </Form.Item>
          <Form.Item name="association_dict_name" label="关联字典名称">
            <Input placeholder="选填" />
          </Form.Item>
          <Form.Item name="association_dict_value" label="关联字典值">
            <Input placeholder="选填" />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
}
