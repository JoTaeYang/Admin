import { useState } from "react";

interface ListProps {
  data: any[];
  perPage?: number;
  children: (visibleData: any[]) => React.ReactNode;
}

export const List = ({ data, perPage = 10, children }: ListProps) => {
  const [page, setPage] = useState(1);
  const pageCount = Math.ceil(data.length / perPage);

  const start = (page - 1) * perPage;
  const end = start + perPage;
  const visible = data.slice(start, end);

  return (
    <div className="space-y-4 mt-6">
      {/* 테이블 */}
      <div className="border rounded-md overflow-auto">
        <table className="w-full table-fixed border-collapse">
          {children(visible)}
        </table>
      </div>

      {/* 페이지네이션 */}
      <div className="flex justify-center gap-2">
        <button
          disabled={page <= 1}
          onClick={() => setPage((p) => Math.max(1, p - 1))}
          className="px-3 py-1 border rounded bg-gray-100 disabled:opacity-40"
        >
          ◀ 이전
        </button>

        <span className="text-sm text-gray-600 mt-1">
          페이지 {page} / {pageCount}
        </span>

        <button
          disabled={page >= pageCount}
          onClick={() => setPage((p) => Math.min(pageCount, p + 1))}
          className="px-3 py-1 border rounded bg-gray-100 disabled:opacity-40"
        >
          다음 ▶
        </button>
      </div>
    </div>
  );
};
