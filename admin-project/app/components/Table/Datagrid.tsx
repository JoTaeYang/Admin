interface DatagridProps {
    headers: string[];
    children: React.ReactNode;
  }
  
  export const Datagrid = ({ headers, children }: DatagridProps) => {
    return (
      <>
        <thead className="bg-gray-100 text-sm text-gray-700 text-left">
          <tr>
            {headers.map((header) => (
              <th key={header} className="px-4 py-2 border-b">{header}</th>
            ))}
          </tr>
        </thead>
        <tbody>
          {children}
        </tbody>
      </>
    );
  };
  