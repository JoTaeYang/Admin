import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
    index("routes/home.tsx"),                       //default path
    route("dashboard", "routes/dashboard.tsx"),     // /dashboard add
] satisfies RouteConfig;
