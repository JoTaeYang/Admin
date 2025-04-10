import { useEffect, useState } from "react";
import dropboxData from "./dropboxData.json";

/*
리터럴로 제한한 방법임
interface DropBoxProps {
  tab: "management" | "operation"; // 또는 string으로도 가능
}
*/


interface DropBoxProps {
  tab: string;
  title : string;
}

export default function DropBox({ tab, title = "default"  }: DropBoxProps) {
  const [items, setItems] = useState<string[]>([]);
  const [isOpen, setIsOpen] = useState(true);

  useEffect(() => {
    const data = dropboxData[tab as keyof typeof dropboxData] || [];
    setItems(data);
    setIsOpen(true); // 탭 바뀌면 자동 펼침
  }, [tab]);

  return (
    <div className="bg-white border border-gray-300 rounded-md shadow-md overflow-hidden transition-all duration-300">
      {/* 헤더 (Toggle 버튼 역할) */}
      <button
        onClick={() => setIsOpen((prev) => !prev)}
        className="w-full text-left px-4 py-2 font-bold text-gray-700 bg-gray-100 hover:bg-gray-200 transition"
      >
        {title} {isOpen ? "▲" : "▼"}
      </button>

      {/* 내용 영역 */}
      <div
        className={`px-4 py-2 text-sm text-gray-600 transition-all duration-300 ${
          isOpen ? "max-h-[500px] opacity-100" : "max-h-0 opacity-0"
        } overflow-hidden`}
      >
        <ul className="list-disc pl-5 space-y-1">
          {items.map((item, i) => (
            <li key={i}>{item}</li>
          ))}
        </ul>
      </div>
    </div>
  );
}