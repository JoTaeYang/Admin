  import { createContext, useContext, useState } from "react";
import type { ReactNode } from "react"; 

type AlertType = "success" | "error";

interface Alert {
  message: string;
  type: AlertType;
}

const AlertContext = createContext<(alert: Alert) => void>(() => {});

export function AlertProvider({ children }: { children: ReactNode }) {
  const [alert, setAlert] = useState<Alert | null>(null);

  const showAlert = (newAlert: Alert) => {
    setAlert(newAlert);
    setTimeout(() => setAlert(null), 3000); // 3초 뒤 자동 사라짐
  };

  return (
    <AlertContext.Provider value={showAlert}>
      {children}
      {alert && (
        <div
          className={`fixed top-4 left-1/2 -translate-x-1/2 px-6 py-3 rounded shadow-lg text-white font-bold transition-all z-50 ${
            alert.type === "error" ? "bg-red-500" : "bg-blue-500"
          }`}
        >
          {alert.message}
        </div>
      )}
    </AlertContext.Provider>
  );
}

export function useAlert() {
  return useContext(AlertContext);
}