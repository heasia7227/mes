"use client";

import {
  ApartmentOutlined,
  CheckCircleOutlined,
  LockOutlined,
  ReloadOutlined,
  SafetyCertificateOutlined,
  ThunderboltFilled,
  UserOutlined,
} from "@ant-design/icons";
import {
  App as AntdApp,
  Alert,
  Button,
  Checkbox,
  Form,
  Input,
  Tooltip,
  Typography,
} from "antd";
import { useRouter } from "next/navigation";
import { useCallback, useEffect, useState } from "react";
import { createCaptchaKey, getCaptchaUrl } from "@/lib/request/system/login";
import { useAuth } from "@/lib/auth-context";
import { systemInfo } from "@/lib/mock-data";

interface LoginFormValues {
  login_name: string;
  password: string;
  captcha: string;
  captcha_key: string;
  remember: boolean;
}

export default function LoginView() {
  const router = useRouter();
  const { user, login } = useAuth();
  const { message } = AntdApp.useApp();
  const [form] = Form.useForm<LoginFormValues>();
  const [loading, setLoading] = useState(false);
  const [captchaKey, setCaptchaKey] = useState("");

  useEffect(() => {
    if (user) {
      router.replace("/dashboard");
    }
  }, [user, router]);

  // 刷新验证码：重新生成 captcha_key，并清空已输入的验证码
  const refreshCaptcha = useCallback(() => {
    const key = createCaptchaKey();
    setCaptchaKey(key);
    form.setFieldsValue({ captcha_key: key, captcha: "" });
  }, [form]);

  // 挂载后再生成验证码，避免服务端与客户端随机值不一致导致 hydration 警告
  useEffect(() => {
    const timer = window.setTimeout(refreshCaptcha, 0);
    return () => window.clearTimeout(timer);
  }, [refreshCaptcha]);

  const handleSubmit = async (values: LoginFormValues) => {
    setLoading(true);
    try {
      const result = await login({
        loginName: values.login_name,
        password: values.password,
        captcha: values.captcha,
        captchaKey: values.captcha_key,
        remember: values.remember,
      });
      if (!result.ok) {
        message.error(result.message ?? "登录失败");
        // 验证码一次性校验，无论成功失败都需要刷新
        refreshCaptcha();
        return;
      }
      message.success("登录成功，欢迎回来");
      router.replace("/dashboard");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="login-page">
      <section className="login-brand">
        <div className="login-brand-header">
          <span
            style={{
              width: 44,
              height: 44,
              borderRadius: 12,
              display: "inline-flex",
              alignItems: "center",
              justifyContent: "center",
              fontSize: 22,
              color: "#fff",
              background: "linear-gradient(135deg, #1677ff 0%, #13c2c2 100%)",
              boxShadow: "0 8px 20px rgba(22, 119, 255, 0.35)",
            }}
          >
            <ThunderboltFilled />
          </span>
          <div>
            <div style={{ fontSize: 18, fontWeight: 600, letterSpacing: 0.5 }}>
              {systemInfo.system_name}
            </div>
            <div style={{ fontSize: 12, color: "rgba(255,255,255,0.55)" }}>
              Manufacturing Execution System
            </div>
          </div>
        </div>

        <div className="login-brand-main">
          <h1 className="login-brand-title">
            让制造现场
            <br />
            更透明、更高效
          </h1>
          <p className="login-brand-desc">
            覆盖生产计划、工单执行、质量追溯与设备管理的全流程数字化平台，
            实时掌握每一道工序的执行状态。
          </p>
          <div className="login-brand-features">
            <div className="login-brand-feature">
              <span className="login-brand-feature-icon">
                <ApartmentOutlined />
              </span>
              生产执行过程透明化管理
            </div>
            <div className="login-brand-feature">
              <span className="login-brand-feature-icon">
                <SafetyCertificateOutlined />
              </span>
              物料与质量全链路追溯
            </div>
            <div className="login-brand-feature">
              <span className="login-brand-feature-icon">
                <CheckCircleOutlined />
              </span>
              多角色协同的统一工作台
            </div>
          </div>
        </div>

        <div className="login-brand-footer">{systemInfo.footer_text}</div>
      </section>

      <section className="login-form-panel">
        <div className="login-card">
          <Typography.Title level={3} className="login-card-title">
            欢迎登录
          </Typography.Title>
          <p className="login-card-subtitle">请使用系统账号登录后台管理平台</p>

          <Alert
            type="info"
            showIcon
            style={{ marginBottom: 22 }}
            title={
              <span style={{ fontSize: 13 }}>
                请输入账号、密码及图片中的 4 位数字验证码
              </span>
            }
          />

          <Form<LoginFormValues>
            form={form}
            name="login"
            layout="vertical"
            initialValues={{
              login_name: "",
              password: "",
              remember: true,
            }}
            onFinish={handleSubmit}
            size="large"
            requiredMark={false}
          >
            <Form.Item
              name="login_name"
              label="账号"
              rules={[{ required: true, message: "请输入账号" }]}
            >
              <Input
                prefix={<UserOutlined style={{ color: "#a0aab8" }} />}
                placeholder="请输入账号"
                autoComplete="username"
              />
            </Form.Item>

            <Form.Item
              name="password"
              label="密码"
              rules={[{ required: true, message: "请输入密码" }]}
            >
              <Input.Password
                prefix={<LockOutlined style={{ color: "#a0aab8" }} />}
                placeholder="请输入密码"
                autoComplete="current-password"
              />
            </Form.Item>

            <Form.Item name="captcha_key" hidden>
              <Input />
            </Form.Item>

            <Form.Item label="验证码" required>
              <div
                style={{ display: "flex", gap: 12, alignItems: "flex-start" }}
              >
                <Form.Item
                  name="captcha"
                  noStyle
                  rules={[{ required: true, message: "请输入验证码" }]}
                >
                  <Input
                    prefix={
                      <SafetyCertificateOutlined style={{ color: "#a0aab8" }} />
                    }
                    placeholder="请输入验证码"
                    maxLength={4}
                    autoComplete="off"
                    style={{ flex: 1 }}
                  />
                </Form.Item>
                <Tooltip title="看不清？点击图片刷新">
                  <div
                    className="login-captcha"
                    onClick={refreshCaptcha}
                    role="button"
                    tabIndex={0}
                    aria-label="刷新验证码"
                    onKeyDown={(event) => {
                      if (event.key === "Enter" || event.key === " ") {
                        event.preventDefault();
                        refreshCaptcha();
                      }
                    }}
                  >
                    {captchaKey ? (
                      // eslint-disable-next-line @next/next/no-img-element
                      <img src={getCaptchaUrl(captchaKey)} alt="验证码" />
                    ) : (
                      <ReloadOutlined />
                    )}
                  </div>
                </Tooltip>
              </div>
            </Form.Item>

            <Form.Item
              name="remember"
              valuePropName="checked"
              style={{ marginBottom: 20 }}
            >
              <Checkbox>记住登录状态</Checkbox>
            </Form.Item>

            <Form.Item style={{ marginBottom: 0 }}>
              <Button type="primary" htmlType="submit" block loading={loading}>
                登 录
              </Button>
            </Form.Item>
          </Form>

          <div className="login-footer-text">{systemInfo.footer_text}</div>
        </div>
      </section>
    </div>
  );
}
