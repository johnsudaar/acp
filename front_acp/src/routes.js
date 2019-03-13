import Vue from 'vue'
import VueRouter from 'vue-router'

import Device from './components/Pages/Device'
import Config from './components/Pages/Config'
import Network from './components/Pages/Network'

Vue.use(VueRouter)

const routes = [
  { path: '/device/:id', component: Device },
  { path: '/config', component: Config },
  { path: "/", component: Network }
]

const router = new VueRouter({
  routes,
})

export default router
