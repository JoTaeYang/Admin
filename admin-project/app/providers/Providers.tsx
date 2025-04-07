import type { ReactNode } from "react";
import { AlertProvider } from "../components/alert/AlertProvider";

export default function Providers({ children }: { children: ReactNode }) {
    return (
        <AlertProvider>
            {/* 다른 Provider들도 여기서 감싸기 */}
            {children}
        </AlertProvider>
    );
  }