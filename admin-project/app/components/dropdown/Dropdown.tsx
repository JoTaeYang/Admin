import { useState, useEffect } from "react";
import dropdownData from "./dropdownData.json";

interface DropDownProps {
    tab: string;    
    placeholder?: string;
    onChange? : (value : string) => void;
}
  

export default function DropDown({ tab, placeholder = "select", onChange } : DropDownProps) {
  const [selected, setSelected] = useState<string>("UID");
  const [open, setOpen] = useState<boolean>(false);
  const [items, setItems] = useState<string[]>([]);

  useEffect(() => {
    const data = dropdownData[tab as keyof typeof dropdownData] || [];
    setItems(data);
    setOpen(true); // 탭 바뀌면 자동 펼침
  }, [tab]);

  const handleSelect = (value: string) => {
    setSelected(value);
    setOpen(false);

    if (onChange) onChange(value)
  };

  return (
    <div className="inline-block text-left w-52 relative">
      <label className="font-semibold block mb-1">Custom:</label>

      {/* 선택된 항목 */}
      <button
        onClick={() => setOpen(!open)}
        className="w-full bg-blue-500 text-white font-semibold px-4 py-2 rounded shadow"
      >
        {selected ?? placeholder} 
        <span className="float-right">{open ? "▲" : "▼"}</span>
      </button>

      {/* 목록 */}
      {open && (
        <ul className="absolute z-10 w-full bg-blue-500 text-white mt-1 rounded shadow overflow-auto max-h-96">
          {items.map((item) => (
            <li
              key={item}
              onClick={() => handleSelect(item)}
              className={`px-4 py-2 hover:bg-blue-600 cursor-pointer ${
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