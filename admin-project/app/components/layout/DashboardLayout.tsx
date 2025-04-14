// components/Layout/DashboardLayout.tsx
export default function DashboardLayout() {
    return (
      <div className="p-6 grid grid-cols-2 gap-4">
        <div className="p-4 bg-white rounded shadow">📈 방문자 통계</div>
        <div className="p-4 bg-white rounded shadow">📋 최근 작업 로그</div>
        <div className="p-4 bg-white rounded shadow col-span-2">🧠 공지사항 영역</div>
      </div>
    );
  }
  