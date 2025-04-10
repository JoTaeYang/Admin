import Sidebar from "./sidebar";
import { sidebarItems } from "./sidebarItems";
import { Outlet } from "react-router";

interface SidebarLayoutProps {
  tab: "management" | "operation";
}

export default function SidebarLayout({ tab }: SidebarLayoutProps) {
  return (
    <div className="flex min-h-screen">
      <Sidebar basePath={`/${tab}`} items={sidebarItems[tab]} />
      <main className="flex-1 bg-white min-h-screen">
        <Outlet />
      </main>
    </div>
  );
}
