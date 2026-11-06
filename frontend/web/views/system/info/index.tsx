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
  Typography,
} from "antd";
import { useEffect, useState } from "react";
import {
  getSystemInfo,
  updateSystemInfo,
  type InfoPayload,
} from "@/lib/request/system/info";

interface InfoFormValues {
  system_name: string;
  system_logo?: string;
  login_background?: string;
  footer_text?: string;
}

export default function SystemInfoView() {
  const { message } = AntdApp.useApp();
  const [form] = Form.useForm<InfoFormValues>();
  const [saving, setSaving] = useState(false);
  const [infoId, setInfoId] = useState<string>("");

  useEffect(() => {
    const load = async () => {
      try {
        const res = await getSystemInfo();
        if (res.code !== 200 || !res.data) {
          message.warning(res.message || "系统信息加载失败");
          return;
        }
        setInfoId(res.data.id);
        form.setFieldsValue({
          system_name: res.data.system_name,
          system_logo: res.data.system_logo,
          login_background: res.data.login_background,
          footer_text: res.data.footer_text,
        });
      } catch (error: any) {
        message.error(error?.message ?? "系统信息加载失败");
      }
    };
    void load();
  }, [form, message]);

  const handleFinish = async (values: InfoFormValues) => {
    if (!infoId) {
      message.warning("系统信息未加载，无法保存");
      return;
    }
    const payload: InfoPayload = {
      system_name: values.system_name,
      system_logo: values.system_logo,
      login_background: values.login_background,
      footer_text: values.footer_text,
    };
    setSaving(true);
    try {
      const res = await updateSystemInfo(infoId, payload);
      if (res.code !== 200) {
        message.error(res.message || "保存失败");
        return;
      }
      message.success("系统信息已保存");
    } catch (error: any) {
      message.error(error?.message ?? "保存失败");
    } finally {
      setSaving(false);
    }
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
