import { createContext, useContext, useState } from "react";
import type { ReactNode } from "react";

type FormValues = { [key: string]: string };

interface FormContextProps {
  values: FormValues;
  setValue: (name: string, value: string) => void;
}

const FormContext = createContext<FormContextProps | null>(null);

export const useFormContext = () => {
  const ctx = useContext(FormContext);
  if (!ctx) throw new Error("useFormContext must be used inside <FormLayout>");
  return ctx;
};

export function FormLayout({ children }: { children: ReactNode }) {
  const [values, setValues] = useState<FormValues>({});

  const setValue = (name: string, value: string) => {
    setValues((prev) => ({ ...prev, [name]: value }));
  };

  return (
    <FormContext.Provider value={{ values, setValue }}>
      <div className="space-y-4">{children}</div>
    </FormContext.Provider>
  );
}