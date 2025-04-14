import { useNavigate } from "react-router";

interface BackButtonProps {
  label?: string;
  className?: string;
}

export default function BackButton({ label = "돌아가기", className }: BackButtonProps) {
  const navigate = useNavigate();

  return (
    <button
      onClick={() => navigate(-1)}
      className={`text-gray-500 hover:text-blue-500 ${className ?? ""}`}
      title="돌아가기"
    >
      {label}
    </button>
  );
}