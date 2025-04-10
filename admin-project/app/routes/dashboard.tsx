
import ProtectedRoute from "~/components/protected/ProtectedRoute";

function DashboardContent() {
    return (
      <div className="p-6">
        <h1 className="text-3xl font-bold mb-4">📊 Dashboard</h1>
        <p>관리자 전용 대시보드입니다.</p>
      </div>
    );
  }
  
export default function Dashboard() {
    return (
      <ProtectedRoute>
        <DashboardContent />
      </ProtectedRoute>
    );
  }