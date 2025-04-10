import DropDown from "~/components/dropdown/Dropdown";

export default function UserPage() {

    const HandleChange = (value : string) => {
        console.log("select item : ", value);
    };

    return (
      <div>
        <h1 className="text-2xl font-bold mb-4">👤 User Management</h1>
        <p>여기에 계정 테이블을 넣으면 돼.</p>
        <DropDown tab="users.search.type" placeholder="검색 타입" onChange={HandleChange}/>
      </div>
    );
  }
  