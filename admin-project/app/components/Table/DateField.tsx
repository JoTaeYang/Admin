export const DateField = ({ record, source }: { record: any; source: string }) => (
    <td className="px-4 py-2 border-b">{new Date(record[source]).toLocaleDateString()}</td>
  );
  