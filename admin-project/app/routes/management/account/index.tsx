import { Outlet } from "react-router";
import AccountListPage from "./list";

export default function AccountIndexPage() {
  return (
    <>
      <Outlet />
    </>
  );
}