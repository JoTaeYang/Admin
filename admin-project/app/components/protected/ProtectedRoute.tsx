import { useAuth } from "~/hooks/useAuth";
import { Navigate } from "react-router";
import type { ReactNode } from "react";

export default function ProtectedRoute({ children }: { children: ReactNode }) {
  const { isLoggedIn } = useAuth();

  if (isLoggedIn === null) {
    return <div className="p-4">인증 확인 중...</div>; // 로딩 중
  }

  if (!isLoggedIn) {
    return <Navigate to="/" replace />; // 로그인 안 했으면 홈으로 강제 이동
  }

  return children;
}
