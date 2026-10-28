"use client";

import {
  BgColorsOutlined,
  CheckOutlined,
  CompressOutlined,
  LaptopOutlined,
  MoonOutlined,
  SunOutlined,
} from "@ant-design/icons";
import {
  Button,
  Divider,
  Drawer,
  Segmented,
  Slider,
  Space,
  Switch,
  Tooltip,
  Typography,
  theme,
} from "antd";
import {
  THEME_COLORS,
  useThemeSettings,
  type ThemeMode,
} from "@/lib/theme-context";

interface ThemeSettingsDrawerProps {
  open: boolean;
  onClose: () => void;
}

export default function ThemeSettingsDrawer({
  open,
  onClose,
}: ThemeSettingsDrawerProps) {
  const { settings, updateSettings, resetSettings } = useThemeSettings();
  const { token } = theme.useToken();

  return (
    <Drawer
      title="主题设置"
      size={320}
      open={open}
      onClose={onClose}
      extra={
        <Button type="link" size="small" onClick={resetSettings}>
          恢复默认
        </Button>
      }
    >
      <Space orientation="vertical" size={24} style={{ width: "100%" }}>
        <section>
          <Typography.Text strong>主题模式</Typography.Text>
          <Typography.Paragraph
            type="secondary"
            style={{ fontSize: 12, margin: "4px 0 12px" }}
          >
            切换后台整体明暗风格，跟随系统时自动响应系统设置。
          </Typography.Paragraph>
          <Segmented
            block
            value={settings.mode}
            onChange={(value) => updateSettings({ mode: value as ThemeMode })}
            options={[
              { label: "亮色", value: "light", icon: <SunOutlined /> },
              { label: "暗色", value: "dark", icon: <MoonOutlined /> },
              { label: "跟随", value: "system", icon: <LaptopOutlined /> },
            ]}
          />
        </section>

        <Divider style={{ margin: 0 }} />

        <section>
          <Typography.Text strong>
            <BgColorsOutlined style={{ marginRight: 6 }} />
            主题色
          </Typography.Text>
          <Typography.Paragraph
            type="secondary"
            style={{ fontSize: 12, margin: "4px 0 12px" }}
          >
            选择品牌主色，按钮、菜单选中态等会同步更新。
          </Typography.Paragraph>
          <div style={{ display: "flex", flexWrap: "wrap", gap: 12 }}>
            {THEME_COLORS.map((color) => {
              const selected = settings.primaryColor === color.value;
              return (
                <Tooltip title={color.label} key={color.value}>
                  <button
                    type="button"
                    aria-label={color.label}
                    onClick={() =>
                      updateSettings({ primaryColor: color.value })
                    }
                    style={{
                      width: 30,
                      height: 30,
                      padding: 0,
                      borderRadius: "50%",
                      background: color.value,
                      border: selected
                        ? `2px solid ${token.colorText}`
                        : "2px solid transparent",
                      cursor: "pointer",
                      display: "inline-flex",
                      alignItems: "center",
                      justifyContent: "center",
                    }}
                  >
                    {selected ? (
                      <CheckOutlined style={{ color: "#fff", fontSize: 13 }} />
                    ) : null}
                  </button>
                </Tooltip>
              );
            })}
          </div>
        </section>

        <Divider style={{ margin: 0 }} />

        <section>
          <Space style={{ width: "100%", justifyContent: "space-between" }}>
            <Typography.Text strong>
              <CompressOutlined style={{ marginRight: 6 }} />
              紧凑模式
            </Typography.Text>
            <Switch
              checked={settings.compact}
              onChange={(checked) => updateSettings({ compact: checked })}
            />
          </Space>
          <Typography.Paragraph
            type="secondary"
            style={{ fontSize: 12, margin: "4px 0 0" }}
          >
            缩小组件内边距，适合信息密度较高的场景。
          </Typography.Paragraph>
        </section>

        <section>
          <Typography.Text strong>组件圆角</Typography.Text>
          <Typography.Paragraph
            type="secondary"
            style={{ fontSize: 12, margin: "4px 0 0" }}
          >
            当前圆角：{settings.borderRadius}px
          </Typography.Paragraph>
          <Slider
            min={0}
            max={16}
            step={2}
            value={settings.borderRadius}
            onChange={(value) => updateSettings({ borderRadius: value })}
            marks={{ 0: "0", 6: "6", 12: "12", 16: "16" }}
          />
        </section>

        <Typography.Text type="secondary" style={{ fontSize: 12 }}>
          主题设置会自动保存在当前浏览器中，下次访问继续生效。
        </Typography.Text>
      </Space>
    </Drawer>
  );
}
