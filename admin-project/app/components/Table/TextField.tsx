export const TextField = ({ record, source, className }: { record: any; source: string; className: string }) => (
    <td className={`${className} px-4 py-2 border-b`}>{record[source]}</td>
  );
  