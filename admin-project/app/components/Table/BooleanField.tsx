export const BooleanField = ({ record, source }: { record: any; source: string }) => (
    <td className="px-4 py-2 border-b">{record[source] ? "✅" : "❌"}</td>
  );
  