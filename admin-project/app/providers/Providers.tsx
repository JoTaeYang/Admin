import type { ReactNode } from "react";
import { AlertProvider } from "../components/alert/AlertProvider";
import { ApiProvider } from "~/components/api/ApiProvider";

export default function Providers({ children }: { children: ReactNode }) {
    return (
        <AlertProvider>
            <ApiProvider>
                {/* 다른 Provider들도 여기서 감싸기 */}
                {children}
            </ApiProvider>
        </AlertProvider>
    );
  }