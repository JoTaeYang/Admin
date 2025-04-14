import { useNavigate } from "react-router";
import { useEffect, useState } from "react";
import { List } from "~/components/Table/List";
import { TextField } from "~/components/Table/TextField";
import { Datagrid } from "~/components/Table/Datagrid";
import { useApi } from "~/components/api/ApiProvider";

const dummyAccounts = [
    { id: 1, email: "admin1@bigf.com", username: "adminMaster", role: "ADMIN" },
    { id: 2, email: "user2@bigf.com", username: "john_doe", role: "USER" },
    { id: 3, email: "manager3@bigf.com", username: "pm_lee", role: "MANAGER" },
    { id: 4, email: "qa4@bigf.com", username: "qa_test", role: "QA" },
    { id: 5, email: "guest5@bigf.com", username: "guest", role: "GUEST" },
    { id: 6, email: "ad2min1@bigf.com", username: "adminMaster", role: "ADMIN" },
    { id: 7, email: "us2er2@bigf.com", username: "john_doe", role: "USER" },
    { id: 8, email: "ma2nager3@bigf.com", username: "pm_lee", role: "MANAGER" },
    { id: 9, email: "qa24@bigf.com", username: "qa_test", role: "QA" },
    { id: 10, email: "gu2est5@bigf.com", username: "guest", role: "GUEST" },
];


export default function AccountListPage() {
    const [accounts, setAccounts] = useState<any[]>([]);
    const navigate = useNavigate();
    const headerArr = ["Email", "Username", "Role", "Action"]
    const columnWidth = `w-[${100 / headerArr.length}%]`;
    const { get } = useApi();


    useEffect(() => {        
        const response =  get("/management/account/list");
      
        fetch("/api/account/list")
          .then((res) => res.json())
          .then(setAccounts)
          .catch(console.error);
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
                                <TextField className={columnWidth} record={record} source="email" />
                                <TextField className={columnWidth} record={record} source="username" />
                                <TextField className={columnWidth} record={record} source="role" />                                
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

