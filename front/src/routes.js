import Vue from 'vue'
import VueRouter from 'vue-router'

import Device from './components/Pages/Device'
import Config from './components/Pages/Config'
import Network from './components/Pages/Network'
import RecControl from './components/Pages/RecControl.vue'
import CamControl from './components/Pages/CamControl.vue'
import Cockpit from './components/Pages/Cockpit.vue'
import CCU from './components/CCU/Main'

Vue.use(VueRouter)

const routes = [
  { path: '/device/:id', component: Device },
  { path: '/ccu', component: CCU },
  { path: '/cam/control', component: CamControl },
  { path: '/rec/control', component: RecControl },
  { path: '/cockpit', component: Cockpit },
  { path: '/config', component: Config, meta: { offline: true }},
  { path: "/", component: Network }
]

const router = new VueRouter({
  routes,
})

export default router
