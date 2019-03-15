import Vue from 'vue'
import VueRouter from 'vue-router'

import Device from './components/Pages/Device'
import Config from './components/Pages/Config'
import Network from './components/Pages/Network'
import RecControl from './components/Pages/RecControl.vue'

Vue.use(VueRouter)

const routes = [
  { path: '/device/:id', component: Device },
  { path: '/rec/control', component: RecControl },
  { path: '/config', component: Config },
  { path: "/", component: Network }
]

const router = new VueRouter({
  routes,
})

export default router
