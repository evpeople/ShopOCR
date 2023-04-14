import {createRouter, createWebHistory} from 'vue-router';

import Home from './components/Home.vue';
import Login from './components/Login.vue';
import Register from './components/Register.vue';

// lazy-loaded
const Profile = () => import('./components/Profile.vue')
const BoardAdmin = () => import('./components/BoardAdmin.vue')
const BoardModerator = () => import('./components/BoardModerator.vue')
const BoardUser = () => import('./components/BoardUser.vue')
const Upload = () => import('./components/Upload.vue')
const Map = () => import('./components/Baidu.vue')
const Map2 = () => import('./components/Baidu2.vue')
const Map3 = () => import('./components/Baidu3.vue')

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home,
  },
  {
    path: '/home',
    component: Home,
  },
  {
    path: '/login',
    component: Login,
  },
  {
    path: '/upload',
    name: 'upload',
    // lazy-loaded
    component: Upload,
  },
  {
    path: '/register',
    component: Register,
  },
  {
    path: '/profile',
    name: 'profile',
    // lazy-loaded
    component: Profile,
  },
  {
    path: '/admin',
    name: 'admin',
    // lazy-loaded
    component: BoardAdmin,
  },
  {
    path: '/mod',
    name: 'moderator',
    // lazy-loaded
    component: BoardModerator,
  },
  {
    path: '/user',
    name: 'user',
    // lazy-loaded
    component: BoardUser,
  },
  {
    path: '/oldmap',
    name: 'name',
    // lazy-loaded
    component: Map,
  },
  {
    path: '/map2',
    name: 'name2',
    // lazy-loaded
    component: Map2,
  },
  {
    path: '/map3',
    name: 'name3',
    // lazy-loaded
    component: Map3,
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;