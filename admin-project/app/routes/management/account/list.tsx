import { useNavigate } from "react-router";
import { useEffect, useState } from "react";
import { List } from "~/components/Table/List";
import { TextField } from "~/components/Table/TextField";
import { Datagrid } from "~/components/Table/Datagrid";
import { useApi } from "~/components/api/ApiProvider";
import { type Manager } from "~/types/Manager";
import { ManagerAPI } from "./endpoint";

export default function AccountListPage() {
    const navigate = useNavigate();
    const headerArr = ["id", "grade", "name", "create_at", "update_at","Action"]
    const columnWidth = `w-[${100 / headerArr.length}%]`;
    const { get } = useApi();
    const [accounts, setAccounts] = useState<Manager[]>([]);

    useEffect(() => {        
        async function fetchList() {
            try {
                const response = await get(ManagerAPI.list);
                console.log(response.data)
                setAccounts(response.data);
            }
            catch(err) {
                alert(`invalid request`);
            }            
        }        
        fetchList();
    }, []);

    return (
        <div className="p-6">
            <div className="flex justify-between items-center mb-4">
                <h2 className="text-xl font-bold">Account 목록</h2>
                <button
                    className="bg-blue-500 text-white px-4 py-2 rounded"
                    onClick={() => navigate("create")}
                >
                    Create
                </button>
            </div>

            <List data={accounts} perPage={5}>
                {(rows) => (
                    <Datagrid headers={headerArr}>
                        {rows.map((record) => (
                            <tr key={record.id}>
                                <TextField className={columnWidth} record={record} source="id" />
                                <TextField className={columnWidth} record={record} source="grade" />
                                <TextField className={columnWidth} record={record} source="name" />                                
                                <TextField className={columnWidth} record={record} source="create_at" />    
                                <TextField className={columnWidth} record={record} source="update_at" />    
                                <td className="px-4 py-2 border-b">
                                    <button
                                        onClick={() =>
                                            navigate(`edit/${record.id}`, {
                                                state: record, // ✅ 현재 row 데이터를 함께 전달
                                            })
                                        }
                                        className="text-blue-600 hover:underline"
                                    >
                                        Edit
                                    </button>
                                </td>
                            </tr>
                        ))}
                    </Datagrid>
                )}
            </List>
        </div>
    );
}

