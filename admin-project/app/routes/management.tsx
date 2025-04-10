import { NavLink, Outlet } from "react-router";


export default function ManagementLayout() {
  return (
    <div className="flex h-full">

      {/* 우측 콘텐츠 영역 */}
      <div className="flex-1 bg-white  overflow-auto">
        <Outlet />
      </div>
    </div>
  );
}
