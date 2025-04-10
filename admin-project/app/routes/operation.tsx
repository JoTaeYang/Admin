
import {  Outlet } from "react-router";

export default function OperationLayout() {
  return (
    <div>
      {/* 하위 탭 콘텐츠 영역 */}
      <div className="p-6 bg-white">
        <Outlet />
      </div>
    </div>
  );
}