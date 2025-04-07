import apiUrls from "./config/env.json";

export type EnvType = keyof typeof apiUrls;

// 환경별 API 주소 정의
export const API_BASE_URLS: Record<EnvType, string> = apiUrls;