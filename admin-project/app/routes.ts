import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
    index("routes/home.tsx"),                                    //default path
    route("", "components/layout/TopTabsLayout.tsx", [
        route("dashboard", "routes/dashboard.tsx", [
            index("routes/dashboard/index.tsx"),                     // /dashboard
        ]),     // /dashboard add
        route("operation", "routes/operation.tsx", [
            index("routes/operation/index.tsx"),                     // /dashboard
        ]),    // /operation
        route("management", "routes/management.tsx", [
            index("routes/management/index.tsx"),                     // /dashboard
            route("accout", "routes/management/account.tsx"),
        ])   // /dashboard/management
    ])
] satisfies RouteConfig;

