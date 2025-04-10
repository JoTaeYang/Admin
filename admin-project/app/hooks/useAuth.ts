import { useEffect, useState } from "react";
import { API_BASE_URLS } from "~/config";
import { useNavigate } from "react-router";


export const useAuth = () => {
  const [isLoggedIn, setIsLoggedIn] = useState<boolean | null>(null); // null: 아직 판단 안 됨
  const stored = localStorage.getItem("env");
  const navigate = useNavigate();

  if (stored === null)
  {
    navigate("")
  }

  const env = (stored === "Live" || stored === "QA" || stored === "Dev") ? stored : "Live";

  useEffect(() => {
    fetch(`${API_BASE_URLS[env]}/me`, {
      method: "GET",
      credentials: "include",
    })
      .then((res) => setIsLoggedIn(res.ok))
      .catch(() => setIsLoggedIn(false));
  }, []);

  return { isLoggedIn };
};