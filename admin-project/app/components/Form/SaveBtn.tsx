import { useFormContext } from "./FormContext";

export  function SaveBtn({ endpoint, onSuccess }: { endpoint: string; onSuccess?: () => void }) {
    const { values } = useFormContext();
  
    const handleClick = async () => {
      const res = await fetch(endpoint, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(values),
      });
  
      if (res.ok && onSuccess) {
        onSuccess();
      }
    };
  
    return (
      <button className="w-full bg-blue-500 text-white px-4 py-2 rounded" onClick={handleClick}>
        Save
      </button>
    );
  }
  