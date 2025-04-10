import { NavLink, Outlet } from "react-router";

const topTabs = [
  { name: "Dashboard", path: "/dashboard" },
  { name: "Operation", path: "/operation" },
  { name: "Management", path: "/management" },
];

export default function TopTabsLayout() {
  return (
    <div className="h-screen flex flex-col">
      {/* 상단 탭 */}
      <div className="flex bg-blue-400 px-6 py-3 text-white shadow">
        <div className="flex items-center gap-2 mr-6">
          <img src="/logo.svg" alt="logo" className="w-6 h-6" />
          <span className="font-bold text-lg">BIGF</span>
          <span className="text-xs ml-1">Ver 1.0.0</span>
        </div>

        <div className="flex gap-2">
          {topTabs.map((tab) => (
            <NavLink
              key={tab.name}
              to={tab.path}
              className={({ isActive }) =>
                `px-4 py-2 font-semibold rounded-t-md transition-all ${
                  isActive ? "bg-blue-100 text-gray-800 shadow-inner" : "hover:bg-blue-300"
                }`
              }
            >
              {tab.name}
            </NavLink>
          ))}
        </div>
      </div>

      {/* 하위 콘텐츠 (탭별 라우팅) */}
      <div className="flex-1 overflow-auto bg-gray-50">
        <Outlet />
      </div>
    </div>
  );
}
