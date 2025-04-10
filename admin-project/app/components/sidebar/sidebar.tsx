import { NavLink, useLocation } from "react-router";
import type { FC } from "react";

interface SidebarProps {
  basePath: string;
  items: { name: string; path: string }[];
}

const Sidebar: FC<SidebarProps> = ({ basePath, items }) => {
  const location = useLocation();

  return (
    <aside className="w-48 bg-blue-200 min-h-screen p-4 flex flex-col justify-start">
      {items.map((item) => {
        const fullPath = `${basePath}/${item.path}`;
        const isActive = location.pathname.startsWith(fullPath);
        return (
          <NavLink
            key={item.name}
            to={fullPath}
            className={`block px-4 py-2 rounded hover:bg-blue-300 ${
              isActive ? "bg-blue-500 text-white" : ""
            }`}
          >
            {item.name}
          </NavLink>
        );
      })}
    </aside>
  );
};

export default Sidebar;
