import AdminLayout from "@/components/AdminLayout";

export const instant = false;

export default function AdminGroupLayout({ children }: LayoutProps<"/">) {
  return <AdminLayout>{children}</AdminLayout>;
}
