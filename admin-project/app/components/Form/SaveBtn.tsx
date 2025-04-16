import { useFormContext } from "./FormContext";
import { useApi } from "../api/ApiProvider";

export  function SaveBtn({ endpoint, onSuccess }: { endpoint: string; onSuccess?: () => void }) {
    const { values } = useFormContext();
    const { post } = useApi();

    const handleClick = async () => {
      const res = await post(endpoint, values);
        
      if (res.message === "success" && onSuccess) {    
        console.log("check")
        onSuccess();
      }
      
    };
  
    return (
      <button className="w-full bg-blue-500 text-white px-4 py-2 rounded" onClick={handleClick}>
        Save
      </button>
    );
  }
  