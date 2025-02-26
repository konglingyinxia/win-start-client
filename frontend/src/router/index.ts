import {createMemoryHistory, createRouter, RouteRecordRaw} from 'vue-router'

import IndexView from "../views/Main.vue";
import ApplicationView from "../views/application/app/Index.vue"
import DashboardView from "../views/dashboard/Index.vue"
import envView from "../views/application/env/Index.vue"

export const DashboardRouterItem =
    {
        path: 'dashboard',
        component: DashboardView,
        meta: {
            id: "dashboard", title: '首页', icon: 'HomeFilled', path: '/dashboard'
        },
    }

const routes: RouteRecordRaw[] = [
    {
        path: '/',
        component: IndexView,
        name: "main",
        children: [
            DashboardRouterItem,
            {
                path: 'application',
                meta: {
                    id: 'application', title: '应用中心', icon: 'Grid', path: '/application'
                },
                children: [
                    {
                        path: 'app',
                        component: ApplicationView,
                        meta: {
                            id: 'app', title: '应用管理', icon: 'Menu', path: '/application/app'
                        },
                    },
                    {
                        path: 'env',
                        component: envView,
                        meta: {
                            id: 'env', title: '运行环境', icon: 'Compass', path: '/application/env'
                        },
                    },
                    // //文件管理
                    // {
                    //     path: 'file',
                    //     component: () => import("../views/application/file/Index.vue"),
                    //     meta: {
                    //         id: 'file', title: '文件管理', icon: 'Files', path: '/application/file'
                    //     },
                    // },
                ]
            },
            //系统管理
            {
                path: 'system',
                meta: {
                    id: 'system', title: '系统信息', icon: 'DataAnalysis', path: '/system'
                },
                children: [
                    //监控
     /*               {
                        path: 'monitor',
                        component: () => import("../views/system/monitor/Index.vue"),
                        meta: {
                            id: 'monitor', title: '系统监控', icon: 'Odometer', path: '/system/monitor'
                        },
                    },*/
                    //进程管理
/*                    {
                        path: 'process',
                        component: () => import("../views/system/process/Index.vue"),
                        meta: {
                            id: 'process', title: '进程管理', icon: 'Monitor', path: '/system/process'
                        },
                    },*/
                    //关于系统
                    {
                        path: 'about',
                        component: () => import("../views/system/about/Index.vue"),
                        meta: {
                            id: 'about', title: '关于系统', icon: 'InfoFilled', path: '/system/about'
                        },
                    },

                ]
            },
        ]
    },
]

export const router = createRouter({
    history: createMemoryHistory(),
    routes: routes,
})
export default router
