import { NavLink, Outlet } from "react-router";

const menuItems = [
  { name: "ACCOUNT", path: "account" },
];

export default function ManagementLayout() {
  return (
    <div className="flex h-full">
      {/* 좌측 사이드바 */}
      <div className="w-48 bg-blue-200 p-4 space-y-2 text-sm font-bold">
        {menuItems.map((item) => (
          <NavLink
            key={item.name}
            to={item.path}
            className={({ isActive }) =>
              `block px-3 py-2 rounded hover:bg-blue-300 ${
                isActive ? "bg-blue-500 text-white" : ""
              }`
            }
          >
            {item.name}
          </NavLink>
        ))}
      </div>

      {/* 우측 콘텐츠 영역 */}
      <div className="flex-1 bg-white p-6 overflow-auto">
        <Outlet />
      </div>
    </div>
  );
}
