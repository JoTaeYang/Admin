import { useFormContext } from "./FormContext";

interface FormInputProps {
    name: string;
    placeholder?: string;
  }

export function FormInput({ name, placeholder }: FormInputProps) {
  const { values, setValue } = useFormContext();

  return (
    <div className="flex flex-col gap-1">
      <label htmlFor={name} className="text-sm text-gray-500">
        {placeholder}
      </label>
      <input
        id={name}
        className="border px-2 py-1 rounded w-full"
        name={name}
        value={values[name] || ""}
        onChange={(e) => setValue(name, e.target.value)}        
      />
    </div>
  );
}
