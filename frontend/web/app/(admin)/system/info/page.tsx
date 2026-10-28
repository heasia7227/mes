"use client";

import {
  App as AntdApp,
  Alert,
  Button,
  Card,
  Col,
  Form,
  Input,
  Row,
  Switch,
  Typography,
} from "antd";
import { useState } from "react";
import { systemInfo } from "@/lib/mock-data";

interface InfoFormValues {
  system_name: string;
  system_logo?: string;
  login_background?: string;
  footer_text?: string;
  is_enabled: boolean;
}

export default function SystemInfoPage() {
  const { message } = AntdApp.useApp();
  const [form] = Form.useForm<InfoFormValues>();
  const [saving, setSaving] = useState(false);

  const handleFinish = async () => {
    setSaving(true);
    await new Promise((resolve) => setTimeout(resolve, 400));
    setSaving(false);
    message.success("系统信息已保存");
  };

  return (
    <div>
      <div style={{ marginBottom: 16 }}>
        <Typography.Title level={4} style={{ margin: 0 }}>
          系统信息
        </Typography.Title>
      </div>

      <Card style={{ height: "calc(100vh - 142px)" }}>
        <Alert
          type="info"
          showIcon
          style={{ marginBottom: 20 }}
          title="系统名称与页脚信息会同步展示在登录页和后台顶栏，Logo 地址支持填写图片 URL。"
        />

        <Form<InfoFormValues>
          form={form}
          layout="vertical"
          onFinish={handleFinish}
          initialValues={{
            system_name: systemInfo.system_name,
            system_logo: systemInfo.system_logo,
            login_background: systemInfo.login_background,
            footer_text: systemInfo.footer_text,
            is_enabled: systemInfo.is_enabled === 1,
          }}
        >
          <Row gutter={16}>
            <Col xs={24} md={12}>
              <Form.Item
                name="system_name"
                label="系统名称"
                rules={[{ required: true, message: "请输入系统名称" }]}
              >
                <Input size="large" placeholder="请输入系统名称" />
              </Form.Item>
            </Col>
            <Col xs={24} md={12}>
              <Form.Item name="footer_text" label="页脚文案">
                <Input size="large" placeholder="请输入页脚版权信息" />
              </Form.Item>
            </Col>
            <Col xs={24} md={12}>
              <Form.Item name="system_logo" label="系统 Logo 地址">
                <Input
                  size="large"
                  placeholder="https://example.com/logo.png"
                />
              </Form.Item>
            </Col>
            <Col xs={24} md={12}>
              <Form.Item name="login_background" label="登录背景图地址">
                <Input
                  size="large"
                  placeholder="https://example.com/login-bg.jpg"
                />
              </Form.Item>
            </Col>
            <Col xs={24}>
              <Form.Item
                name="is_enabled"
                label="是否启用"
                valuePropName="checked"
              >
                <Switch checkedChildren="启用" unCheckedChildren="停用" />
              </Form.Item>
            </Col>
          </Row>

          <Button
            type="primary"
            htmlType="submit"
            size="large"
            loading={saving}
          >
            保存配置
          </Button>
        </Form>
      </Card>
    </div>
  );
}
