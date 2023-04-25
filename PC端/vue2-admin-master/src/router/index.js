import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/home',
    children: [{
      path: 'home',
      name: 'home',
      component: () => import('@/views/home/index'),
      meta: { title: '首页', icon: 'el-icon-s-home' }
    }]
  },
  {
    path:'/acl',
    name:'acl',
    component:Layout,
    redirect: '/acl/user/list',
    meta: { title: '权限管理', icon: 'el-icon-setting' },
    children: [
      {
        name: "user",
        path: "/acl/user/list",
        component: () => import("@/views/acl/user/index.vue"),
        meta: {
          title: "用户管理",
        },
      },
      {
        name: "role",
        path: "/acl/role/list",
        component: () => import("@/views/acl/role/index.vue"),
        meta: {
          title: "角色管理",
        },
      },
    ]
  },
  {
    path: '/product',
    component: Layout,
    name: 'product',
    redirect: '/product/trademark/list',
    meta: { title: '商品管理', icon: 'el-icon-box' },
    children: [
      {
        path: 'trademark/list',
        name: 'trademark',
        component: () => import('@/views/trademark/index'),
        meta: { title: '品牌管理'}
      },
      {
        path: 'attr/list',
        name: 'attr',
        component: () => import('@/views/attr/index'),
        meta: { title: '平台属性管理'}
      }
    ]
  },

  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
