// 所有路由，其中name=admin和左边菜单栏一致
export const routes = [
  {
    path: "/login",
    component: () => import("@/views/login/index"),
  },
  {
    path: "/",
    name: "admin",
    component: () => import("@/views/admin/index"),
    children: [
      {
        path: "/",
        component: () => import("@/views/admin/welcome/index"),
        text: "首页"
      },
      {
        path: "/setting",
        component: () => import("@/views/admin/setting/index"),
        text: "设置"
      },
    ],
  },
];

// 获取admin的菜单
export function adminMenu() {
  return routes.filter(x => x.name == "admin").shift()['children'];
}