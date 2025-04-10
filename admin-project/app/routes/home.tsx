import { useState, useEffect } from "react";
import { useApi } from "~/components/api/ApiProvider";
import { API_BASE_URLS, type EnvType } from "../config";
import { useNavigate } from "react-router";

export default function Home() {
  const ENV_LIST = ["Live", "QA", "Dev"] as const;
  type EnvType = (typeof ENV_LIST)[number];  
  const [env, setEnv] = useState<EnvType>("Live");
  const [id, setId] = useState(""); 
  const [pw, setPw] = useState("");
  const { post } = useApi();
  const navigate = useNavigate();
  
  useEffect(() => {
    const stored = localStorage.getItem("env");
    if (stored === "Live" || stored === "QA" || stored === "Dev") {
      setEnv(stored);
    }
  }, []);

  const handleLogin = async () => {
    try {
      const response = await post("/login", {
        id: id,
        password: pw,
      });
  
      console.log("로그인 성공!", response);
      navigate("/dashboard")
    } catch (err) {
      alert(`로그인 시도:\nEnv: ${env}\nID: ${id}\nPW: ${pw}`);
      console.error("로그인 실패:", err);
    }
  };

  // 선택 시 저장
  const handleEnvChange = (env: EnvType) => {
    localStorage.setItem("env", env);
    setEnv(env);
    
    console.log(`환경 선태됨: ${env}`)
  };

  return (
    <div className="min-h-screen bg-gray-100 flex flex-col items-center justify-center">
      {/* 환경 선택 버튼 */}
      <div className="flex gap-3 mb-6">
        {ENV_LIST.map((e) => (
          <button
            key={e}
            onClick={() => handleEnvChange(e)}
            className={`px-6 py-2 rounded shadow font-semibold border ${
              env === e
                ? "bg-blue-500 text-white border-blue-500"
                : "bg-gray-200 text-gray-700 border-gray-300"
            }`}
          >
            {e}
          </button>
        ))}
      </div>

      {/* 로그인 폼 */}
      <div className="w-80 space-y-4">
        <div>
          <label className="block text-sm font-medium mb-1">ID</label>
          <input
            type="text"
            value={id}
            onChange={(e) => setId(e.target.value)}
            className="w-full px-4 py-2 rounded border bg-gray-200 focus:outline-none"
          />
        </div>
        <div>
          <label className="block text-sm font-medium mb-1">PW</label>
          <input
            type="password"
            value={pw}
            onChange={(e) => setPw(e.target.value)}
            className="w-full px-4 py-2 rounded border bg-gray-200 focus:outline-none"
          />
        </div>
        <button
          onClick={handleLogin}
          className="w-full mt-4 px-4 py-2 bg-green-400 text-white font-bold rounded shadow hover:bg-green-500"
        >
          LOGIN
        </button>
      </div>
    </div>
  );
}