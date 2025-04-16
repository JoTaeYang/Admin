import { useState, useEffect } from "react";
import { useFormContext } from "./FormContext";

import dropdownData from "./dropdownData.json";

interface FormDropDownProps {
    tab: string;    
    name? : string;
    placeholder?: string;
    onChange? : (value : string) => void;
}
  


export  function FormDropDown({ tab, name, placeholder = "select", onChange } : FormDropDownProps) {
  const { values, setValue } = useFormContext();  
  const [selected, setSelected] = useState<string>("select");
  const [open, setOpen] = useState<boolean>(false);
  const [items, setItems] = useState<string[]>([]);
  const dynamicMargin = open ? `${items.length * 3}rem` : "1rem";

  

  useEffect(() => {
    const data = dropdownData[tab as keyof typeof dropdownData] || [];
    setItems(data);          
    if (values && items.length > 0 && name != null) {      
      const value = values[name];
      if (value && items.includes(value)) {
        setSelected(value);
      }
    }
    setOpen(false); // 탭 바뀌면 자동 펼침
  }, [tab, values]);

  const handleSelect = (value: string) => {
    setSelected(value);
    if (name != null)
        setValue(name, value)    
    setOpen(false);
    if (onChange) onChange(value)
  };

  return (
    <div style={{marginBottom : dynamicMargin}}className={`inline-block text-left w-52 relative`}>
      <label className="font-semibold block mb-1">{name}</label>

      {/* 선택된 항목 */}
      <button
        onClick={() => setOpen(!open)}
        className="w-full bg-white-500 text-black font-semibold px-4 py-2 rounded shadow"
      >
        {selected ?? placeholder} 
        <span className="float-right">{open ? "▲" : "▼"}</span>
      </button>

      {/* 목록 */}
      {open && (
        <ul className="absolute top-full z-10 w-full bg-white-500 text-black mt-1 rounded shadow overflow-auto max-h-96">
          {items.map((item) => (
            <li
              key={item}
              onClick={() => handleSelect(item)}
              className={`px-4 py-2 hover:bg-white-600 cursor-pointer ${
                selected === item ? "bg-blue-600" : ""
              }`}
            >
              {item}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}