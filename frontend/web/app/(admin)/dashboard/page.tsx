"use client";

import {
  ApartmentOutlined,
  ArrowRightOutlined,
  DatabaseOutlined,
  FieldTimeOutlined,
  SafetyCertificateOutlined,
  SettingOutlined,
  TeamOutlined,
  UserOutlined,
} from "@ant-design/icons";
import {
  Button,
  Card,
  Col,
  Descriptions,
  Progress,
  Row,
  Space,
  Statistic,
  Table,
  Tag,
  Typography,
  theme,
} from "antd";
import dayjs from "dayjs";
import { useRouter } from "next/navigation";
import { useAuth } from "@/lib/auth-context";
import { depts, loginLogs, roles, systemInfo, users } from "@/lib/mock-data";
import { navigation } from "@/lib/navigation";
import { LOG_STATUS_MAP } from "@/lib/status";

const quickLinks =
  navigation.find((item) => item.key === "/system")?.children ?? [];

export default function DashboardPage() {
  const { user } = useAuth();
  const { token } = theme.useToken();
  const router = useRouter();

  const todayLogins = loginLogs.filter(
    (item) =>
      dayjs(item.login_time).isSame(dayjs(), "day") && item.login_status === 1,
  ).length;

  const stats = [
    {
      title: "系统用户",
      value: users.length,
      icon: <UserOutlined />,
      color: token.colorPrimary,
    },
    {
      title: "组织部门",
      value: depts.length,
      icon: <ApartmentOutlined />,
      color: token.colorSuccess,
    },
    {
      title: "角色数量",
      value: roles.length,
      icon: <TeamOutlined />,
      color: token.colorWarning,
    },
    {
      title: "今日登录",
      value: todayLogins,
      icon: <FieldTimeOutlined />,
      color: token.colorInfo,
    },
  ];

  return (
    <div>
      <Card
        style={{ marginBottom: 16, border: "none", overflow: "hidden" }}
        styles={{
          body: {
            background: `linear-gradient(120deg, ${token.colorPrimary} 0%, #13c2c2 100%)`,
            color: "#fff",
            padding: "26px 28px",
          },
        }}
      >
        <div
          style={{
            display: "flex",
            flexWrap: "wrap",
            gap: 16,
            alignItems: "center",
            justifyContent: "space-between",
          }}
        >
          <div>
            <Typography.Title level={3} style={{ color: "#fff", margin: 0 }}>
              {dayjs().hour() < 12
                ? "上午好"
                : dayjs().hour() < 18
                  ? "下午好"
                  : "晚上好"}
              ，{user?.user_name}
            </Typography.Title>
            <div
              style={{
                marginTop: 8,
                color: "rgba(255,255,255,0.82)",
                fontSize: 13,
              }}
            >
              今天是 {dayjs().format("YYYY年M月D日 dddd")}，欢迎使用{" "}
              {systemInfo.system_name}。
            </div>
          </div>
          <Space>
            <Button
              ghost
              icon={<SettingOutlined />}
              onClick={() => router.push("/system/user")}
            >
              账号管理
            </Button>
            <Button
              style={{
                background: "#fff",
                borderColor: "#fff",
                color: token.colorPrimary,
              }}
              icon={<ArrowRightOutlined />}
              onClick={() => router.push("/system/log")}
            >
              查看操作日志
            </Button>
          </Space>
        </div>
      </Card>

      <Row gutter={[16, 16]} style={{ marginBottom: 16 }}>
        {stats.map((item) => (
          <Col key={item.title} xs={24} sm={12} xl={6}>
            <Card>
              <div style={{ display: "flex", alignItems: "center", gap: 16 }}>
                <span
                  style={{
                    width: 48,
                    height: 48,
                    flex: "0 0 48px",
                    borderRadius: 12,
                    display: "inline-flex",
                    alignItems: "center",
                    justifyContent: "center",
                    fontSize: 22,
                    color: item.color,
                    background: `${item.color}1a`,
                  }}
                >
                  {item.icon}
                </span>
                <Statistic title={item.title} value={item.value} />
              </div>
            </Card>
          </Col>
        ))}
      </Row>

      <Row gutter={[16, 16]} style={{ marginBottom: 16 }}>
        <Col xs={24} xl={16}>
          <Card
            title="快捷入口"
            extra={<Typography.Text type="secondary">系统管理</Typography.Text>}
          >
            <Row gutter={[12, 12]}>
              {quickLinks.map((item) => (
                <Col key={item.key} xs={12} sm={8} md={6}>
                  <button
                    type="button"
                    onClick={() => item.path && router.push(item.path)}
                    style={{
                      width: "100%",
                      padding: "16px 12px",
                      borderRadius: token.borderRadiusLG,
                      border: `1px solid ${token.colorBorderSecondary}`,
                      background: token.colorFillQuaternary,
                      cursor: "pointer",
                      display: "flex",
                      flexDirection: "column",
                      alignItems: "center",
                      gap: 8,
                      color: token.colorText,
                      transition: "all 0.2s",
                    }}
                  >
                    <span style={{ fontSize: 20, color: token.colorPrimary }}>
                      {item.icon}
                    </span>
                    <span style={{ fontSize: 13 }}>{item.label}</span>
                  </button>
                </Col>
              ))}
            </Row>
          </Card>
        </Col>

        <Col xs={24} xl={8}>
          <Card title="今日生产概览" style={{ height: "100%" }}>
            <Space orientation="vertical" size={18} style={{ width: "100%" }}>
              <div>
                <div
                  style={{
                    display: "flex",
                    justifyContent: "space-between",
                    fontSize: 13,
                  }}
                >
                  <span>计划达成率</span>
                  <span style={{ fontWeight: 600 }}>92%</span>
                </div>
                <Progress
                  percent={92}
                  showInfo={false}
                  strokeColor={token.colorPrimary}
                />
              </div>
              <div>
                <div
                  style={{
                    display: "flex",
                    justifyContent: "space-between",
                    fontSize: 13,
                  }}
                >
                  <span>设备开动率</span>
                  <span style={{ fontWeight: 600 }}>88%</span>
                </div>
                <Progress
                  percent={88}
                  showInfo={false}
                  strokeColor={token.colorSuccess}
                />
              </div>
              <div>
                <div
                  style={{
                    display: "flex",
                    justifyContent: "space-between",
                    fontSize: 13,
                  }}
                >
                  <span>一次检验合格率</span>
                  <span style={{ fontWeight: 600 }}>98.6%</span>
                </div>
                <Progress
                  percent={98.6}
                  showInfo={false}
                  strokeColor={token.colorInfo}
                />
              </div>
            </Space>
          </Card>
        </Col>
      </Row>

      <Row gutter={[16, 16]}>
        <Col xs={24} xl={16}>
          <Card
            title="最近登录"
            extra={
              <Button
                type="link"
                size="small"
                onClick={() => router.push("/system/login-log")}
              >
                查看更多
              </Button>
            }
          >
            <Table
              rowKey="id"
              size="small"
              pagination={false}
              dataSource={loginLogs.slice(0, 5)}
              scroll={{ x: "max-content" }}
              columns={[
                {
                  title: "登录账号",
                  dataIndex: "login_name",
                  key: "login_name",
                  width: 120,
                },
                {
                  title: "登录 IP",
                  dataIndex: "login_ip",
                  key: "login_ip",
                  width: 140,
                },
                {
                  title: "登录位置",
                  dataIndex: "login_location",
                  key: "login_location",
                },
                {
                  title: "登录时间",
                  dataIndex: "login_time",
                  key: "login_time",
                  width: 170,
                  render: (value: string) =>
                    dayjs(value).format("YYYY-MM-DD HH:mm"),
                },
                {
                  title: "状态",
                  dataIndex: "login_status",
                  key: "login_status",
                  width: 90,
                  render: (value: number) => {
                    const meta = LOG_STATUS_MAP[String(value)];
                    return <Tag color={meta.color}>{meta.text}</Tag>;
                  },
                },
              ]}
            />
          </Card>
        </Col>

        <Col xs={24} xl={8}>
          <Card title="系统信息" style={{ height: "100%" }}>
            <Descriptions column={1} size="small" colon={false}>
              <Descriptions.Item label="系统名称">
                {systemInfo.system_name}
              </Descriptions.Item>
              <Descriptions.Item label="系统版本">v1.0.0</Descriptions.Item>
              <Descriptions.Item label="后端服务">
                <Space size={4}>
                  <DatabaseOutlined />
                  Beego API Service
                </Space>
              </Descriptions.Item>
              <Descriptions.Item label="数据存储">MySQL</Descriptions.Item>
              <Descriptions.Item label="运行状态">
                <Tag color="success">正常</Tag>
              </Descriptions.Item>
              <Descriptions.Item label="质量保障">
                <Space size={4}>
                  <SafetyCertificateOutlined />
                  全流程可追溯
                </Space>
              </Descriptions.Item>
            </Descriptions>
          </Card>
        </Col>
      </Row>
    </div>
  );
}
