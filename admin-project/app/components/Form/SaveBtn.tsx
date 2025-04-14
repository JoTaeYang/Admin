import { useFormContext } from "./FormContext";

export function SaveBtn({ endpoint }: { endpoint: string }) {
  const { values } = useFormContext();

  const handleClick = async () => {
    console.log(values)
    // const res = await fetch(endpoint, {
    //   method: "POST",
    //   headers: { "Content-Type": "application/json" },
    //   body: JSON.stringify(values),
    // });
    // const result = await res.json();
    // alert("응답: " + JSON.stringify(result));
  };

  return (
    <button
      className="bg-blue-500 text-white px-4 py-2 rounded shadow"
      onClick={handleClick}
    >
      Save
    </button>
  );
}
