import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
    index("routes/home.tsx"),                                    //default path
    route("", "components/layout/TopTabsLayout.tsx", [
        route("dashboard", "routes/dashboard.tsx", [
            index("routes/dashboard/index.tsx"),                     
        ]),    
        route("operation", "routes/operation.tsx", [
            index("routes/operation/index.tsx"),                     
            route("user", "routes/operation/user.tsx"),
        ]),    
        route("management", "routes/management.tsx", [
            index("routes/management/index.tsx"),                     
            route("account", "routes/management/account.tsx"),
        ])  
    ])
] satisfies RouteConfig;

