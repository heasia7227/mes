"use client";

import {
  DeleteOutlined,
  EditOutlined,
  PlusOutlined,
  ReloadOutlined,
  SearchOutlined,
} from "@ant-design/icons";
import {
  App as AntdApp,
  Button,
  Card,
  Form,
  Input,
  InputNumber,
  Modal,
  Popconfirm,
  Select,
  Space,
  Switch,
  Table,
  Tag,
  Typography,
  Row,
  Col,
} from "antd";
import type {
  ExpandableConfig,
  TablePaginationConfig,
} from "antd/es/table/interface";
import dayjs from "dayjs";
import { useMemo, useState, type ReactNode } from "react";
import type { StatusMap } from "@/lib/status";

export type RowData = Record<string, any>;

export interface ResourceColumn {
  key: string;
  title: string;
  width?: number;
  ellipsis?: boolean;
  valueType?: "text" | "date" | "status" | "number";
  statusMap?: StatusMap;
  render?: (value: any, record: RowData) => ReactNode;
}

export interface ResourceField {
  key: string;
  label: string;
  type?: "text" | "textarea" | "number" | "select" | "switch" | "date";
  options?: Array<{ label: string; value: string | number }>;
  required?: boolean;
  placeholder?: string;
  initialValue?: any;
  combineCol?: boolean;
}

export interface ResourcePageProps {
  sectionClassName?: string;
  title: string;
  description?: string;
  columns: ResourceColumn[];
  fields: ResourceField[];
  fieldCol?: number;
  data: RowData[];
  rowKey?: string;
  displayKey?: string;
  searchKeys?: string[];
  searchPlaceholder?: string;
  statusField?: string;
  statusOptions?: Array<{ label: string; value: string }>;
  readOnly?: boolean;
  expandable?: ExpandableConfig<RowData>;
  pagination?: false | TablePaginationConfig;
}

function isChecked(value: any) {
  return value === true || value === 1 || value === "1";
}

export default function ResourcePage({
  sectionClassName,
  title,
  description,
  columns,
  fields,
  fieldCol = 1,
  data,
  rowKey = "id",
  displayKey,
  searchKeys,
  searchPlaceholder,
  statusField,
  statusOptions,
  readOnly = false,
  expandable,
  pagination,
}: ResourcePageProps) {
  const [rows, setRows] = useState<RowData[]>(data);
  const [keyword, setKeyword] = useState("");
  const [statusValue, setStatusValue] = useState<string | undefined>();
  const [modalOpen, setModalOpen] = useState(false);
  const [editing, setEditing] = useState<RowData | null>(null);
  const [form] = Form.useForm<RowData>();
  const { message } = AntdApp.useApp();

  const displayKeyName = displayKey ?? columns[0]?.key ?? rowKey;

  const filteredRows = useMemo(() => {
    const normalizedKeyword = keyword.trim().toLowerCase();
    const keys =
      searchKeys && searchKeys.length > 0
        ? searchKeys
        : columns.map((item) => item.key);

    return rows.filter((row) => {
      const matchKeyword =
        !normalizedKeyword ||
        keys.some((key) =>
          String(row[key] ?? "")
            .toLowerCase()
            .includes(normalizedKeyword),
        );
      const matchStatus =
        !statusValue ||
        !statusField ||
        String(row[statusField] ?? "") === statusValue;
      return matchKeyword && matchStatus;
    });
  }, [rows, keyword, statusValue, searchKeys, statusField, columns]);

  const openCreate = () => {
    setEditing(null);
    form.resetFields();
    const defaults: RowData = {};
    fields.forEach((field) => {
      if (field.initialValue !== undefined)
        defaults[field.key] = field.initialValue;
    });
    form.setFieldsValue(defaults);
    setModalOpen(true);
  };

  const openEdit = (record: RowData) => {
    setEditing(record);
    const values: RowData = {};
    fields.forEach((field) => {
      values[field.key] =
        field.type === "switch"
          ? isChecked(record[field.key])
          : (record[field.key] ?? "");
    });
    form.resetFields();
    form.setFieldsValue(values);
    setModalOpen(true);
  };

  const handleSubmit = async () => {
    const values = await form.validateFields();
    const payload: RowData = { ...values };
    fields.forEach((field) => {
      if (field.type === "switch") {
        payload[field.key] = values[field.key] ? 1 : 0;
      }
    });

    const now = dayjs().format("YYYY-MM-DD HH:mm:ss");

    if (editing) {
      setRows((prev) =>
        prev.map((row) =>
          String(row[rowKey]) === String(editing[rowKey])
            ? { ...row, ...payload, update_time: now }
            : row,
        ),
      );
      message.success("修改成功");
    } else {
      setRows((prev) => [
        {
          [rowKey]: `${Date.now()}`,
          create_time: now,
          update_time: now,
          delete_flag: 0,
          ...payload,
        },
        ...prev,
      ]);
      message.success("新增成功");
    }

    setModalOpen(false);
    setEditing(null);
    form.resetFields();
  };

  const handleDelete = (record: RowData) => {
    setRows((prev) =>
      prev.filter((row) => String(row[rowKey]) !== String(record[rowKey])),
    );
    message.success("删除成功");
  };

  const handleRefresh = () => {
    setRows(data);
    setKeyword("");
    setStatusValue(undefined);
    message.success("数据已刷新");
  };

  const tableColumns: any[] = [
    ...columns.map((column) => ({
      title: column.title,
      dataIndex: column.key,
      key: column.key,
      width: column.width,
      ellipsis: column.ellipsis,
      render: (value: any, record: RowData) => {
        if (column.render) return column.render(value, record);
        if (value === null || value === undefined || value === "") {
          return <Typography.Text type="secondary">-</Typography.Text>;
        }
        if (column.valueType === "status") {
          const meta = column.statusMap?.[String(value)];
          return meta ? (
            <Tag color={meta.color}>{meta.text}</Tag>
          ) : (
            String(value)
          );
        }
        if (column.valueType === "date") {
          const parsed = dayjs(value);
          return parsed.isValid()
            ? parsed.format("YYYY-MM-DD HH:mm")
            : String(value);
        }
        if (column.valueType === "number") {
          return <span>{Number(value)}</span>;
        }
        return String(value);
      },
    })),
  ];

  if (!readOnly) {
    tableColumns.push({
      title: "操作",
      key: "__actions",
      width: readOnly ? 90 : 150,
      fixed: "right" as const,
      render: (_: any, record: RowData) => (
        <Space size={0}>
          <Button
            type="link"
            size="small"
            icon={<EditOutlined />}
            onClick={() => openEdit(record)}
          >
            编辑
          </Button>
          <Popconfirm
            title="确认删除"
            description={`确定删除「${record[displayKeyName] ?? "该记录"}」吗？`}
            okText="删除"
            cancelText="取消"
            okButtonProps={{ danger: true }}
            onConfirm={() => handleDelete(record)}
          >
            <Button type="link" size="small" danger icon={<DeleteOutlined />}>
              删除
            </Button>
          </Popconfirm>
        </Space>
      ),
    });
  }

  return (
    <div>
      <div style={{ marginBottom: 16 }}>
        <Typography.Title level={4} style={{ margin: 0 }}>
          {title}
        </Typography.Title>
      </div>

      <Card className="main-card">
        <div
          style={{
            display: "flex",
            flexWrap: "wrap",
            gap: 12,
            justifyContent: "space-between",
            marginBottom: 16,
          }}
        >
          <Space wrap>
            <Input
              allowClear
              size="large"
              prefix={<SearchOutlined style={{ color: "#a0aab8" }} />}
              placeholder={searchPlaceholder ?? "输入关键字搜索"}
              value={keyword}
              onChange={(event) => setKeyword(event.target.value)}
              style={{ width: 260 }}
            />
            {statusField && statusOptions && (
              <Select
                allowClear
                size="large"
                placeholder="全部状态"
                style={{ width: 140 }}
                options={statusOptions}
                value={statusValue}
                onChange={(value) => setStatusValue(value)}
              />
            )}
          </Space>

          <Space>
            <Button
              size="large"
              icon={<ReloadOutlined />}
              onClick={handleRefresh}
            >
              刷新
            </Button>
            {!readOnly && (
              <Button
                type="primary"
                size="large"
                icon={<PlusOutlined />}
                onClick={openCreate}
              >
                新增
              </Button>
            )}
          </Space>
        </div>

        <Table
          rowKey={rowKey}
          columns={tableColumns}
          dataSource={filteredRows}
          scroll={{
            x: "max-content",
            y: "auto",
          }}
          classNames={{
            section: sectionClassName || "main-table-section",
          }}
          expandable={expandable}
          pagination={
            pagination ?? {
              pageSize: 10,
              size: "large",
              showSizeChanger: true,
              showTotal: (total) => `共 ${total} 条`,
            }
          }
        />
      </Card>

      {!readOnly && (
        <Modal
          title={editing ? `编辑${title}` : `新增${title}`}
          open={modalOpen}
          style={{ top: 20 }}
          width={600}
          okText="保存"
          cancelText="取消"
          forceRender
          okButtonProps={{ size: "large" }}
          onOk={handleSubmit}
          cancelButtonProps={{ size: "large" }}
          onCancel={() => {
            setModalOpen(false);
            setEditing(null);
          }}
        >
          <Form form={form} layout="vertical" style={{ marginTop: 16 }}>
            <Row gutter={12}>
              {fields.map((field) => (
                <Col
                  key={field.key}
                  span={field.combineCol ? 24 : 24 / fieldCol}
                >
                  <Form.Item
                    name={field.key}
                    label={field.label}
                    valuePropName={
                      field.type === "switch" ? "checked" : "value"
                    }
                    rules={
                      field.required
                        ? [{ required: true, message: `请输入${field.label}` }]
                        : undefined
                    }
                  >
                    {renderField(field)}
                  </Form.Item>
                </Col>
              ))}
            </Row>
          </Form>
        </Modal>
      )}
    </div>
  );
}

function renderField(field: ResourceField) {
  const placeholder = field.placeholder ?? `请输入${field.label}`;

  switch (field.type) {
    case "textarea":
      return <Input.TextArea size="large" rows={3} placeholder={placeholder} />;
    case "number":
      return (
        <InputNumber
          size="large"
          style={{ width: "100%" }}
          placeholder={placeholder}
        />
      );
    case "select":
      return (
        <Select
          allowClear
          size="large"
          placeholder={`请选择${field.label}`}
          options={field.options}
        />
      );
    case "switch":
      return <Switch />;
    default:
      return <Input size="large" placeholder={placeholder} />;
  }
}
