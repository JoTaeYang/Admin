import { createContext, useContext } from "react";

const BASE_URL = import.meta.env.VITE_API_BASE_URL || "http://localhost:3000";

const defaultHeaders = {
  "Content-Type": "application/json",
  // 예: Authorization: `Bearer ${token}` 도 여기 추가 가능
};

const request = async (
  method: "GET" | "POST" | "PUT" | "DELETE",
  path: string,
  body?: any
) => {
  const res = await fetch(`${BASE_URL}${path}`, {
    method,
    headers: defaultHeaders,
    body: body ? JSON.stringify(body) : undefined,
  });

  if (!res.ok) {
    const error = await res.json().catch(() => ({}));
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
