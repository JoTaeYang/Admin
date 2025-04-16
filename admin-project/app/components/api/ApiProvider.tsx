import { createContext, useContext } from "react";
import { API_BASE_URLS } from "~/config";

const defaultHeaders = {
  "Content-Type": "application/json",
  // 예: Authorization: `Bearer ${token}` 도 여기 추가 가능
};

const request = async (
  method: "GET" | "POST" | "PUT" | "DELETE",
  path: string,
  body?: any
) => {
  const stored = localStorage.getItem("env");
  const env = (stored === "Live" || stored === "QA" || stored === "Dev") ? stored : "Live";
  
  const res = await fetch(`${API_BASE_URLS[env]}${path}`, {
    method,
    credentials: "include",
    headers: defaultHeaders,
    body: body ? JSON.stringify(body) : undefined,
  });


  if (!res.ok) {
    const error = await res.json().catch(() => ({}));
    console.log(res)

    throw new Error(error.message || `API Error: ${res.status}`);
  }

  return res.json().catch(() => ({})); // 응답이 비어있어도 처리
};

const ApiContext = createContext({
  get: (path: string) => request("GET", path),
  post: (path: string, body: any) => request("POST", path, body),
  put: (path: string, body: any) => request("PUT", path, body),
  del: (path: string) => request("DELETE", path),
});

export const ApiProvider = ({ children }: { children: React.ReactNode }) => {
  return (
    <ApiContext.Provider value={{
      get: (path) => request("GET", path),
      post: (path, body) => request("POST", path, body),
      put: (path, body) => request("PUT", path, body),
      del: (path) => request("DELETE", path),
    }}>
      {children}
    </ApiContext.Provider>
  );
};

export const useApi = () => useContext(ApiContext);
