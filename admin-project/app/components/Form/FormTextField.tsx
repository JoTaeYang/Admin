import { useFormContext } from "./FormContext";

interface FormTextFieldProps {
  name: string;
  placeholder?: string;
}

export function FormTextField({ name, placeholder }: FormTextFieldProps) {
  const { values } = useFormContext();  
  return (
    <div className="flex flex-col gap-1">
      {placeholder && (
        <label className="text-sm text-gray-500">{placeholder}</label>
      )}
      <div className="px-2 py-1 rounded border bg-gray-50 text-gray-800">
        {values[name] || "-"}
      </div>
    </div>
  );
}