import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import Admin from '../views/Admin.vue'
// 页面子路由
import Index from "../components/admin/Index";
import AddArt from "../components/article/AddArt";
import ArtList from "../components/article/ArtList";
import CateList from "../components/category/CateList";
import UserList from "../components/user/UserList";

Vue.use(VueRouter)

const originalReplace = VueRouter.prototype.replace;
VueRouter.prototype.replace = function replace(location) {
  return originalReplace.call(this, location).catch(err => err);
};

const routes = [
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/admin',
    name: 'admin',
    component: Admin,
    children: [
      {path: 'index', component: Index},
      {path: 'addart', component: AddArt},
      {path: 'addart/:id', component: AddArt},
      {path: 'artlist', component: ArtList},
      {path: 'catelist', component: CateList},
      {path: 'userlist', component: UserList},
    ],
  }
]

const router = new VueRouter({
  mode: "history", // 去掉URL中的 # 号
  routes
})

// 路由导航守卫
router.beforeEach((to, from, next) => {
  const token = window.sessionStorage.getItem('ginblog_token')
  if (to.path === '/login') return next()
  if (!token) {
    next('/login')
  } else {
    next()
  }
})


export default router
