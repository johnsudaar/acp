import Vue from 'vue'
import VueRouter from 'vue-router'

import Device from './components/Pages/Device'
import Config from './components/Pages/Config'
import Network from './components/Pages/Network'
import RecControl from './components/Pages/RecControl.vue'
import CamControl from './components/Pages/CamControl.vue'
import Cockpit from './components/Pages/Cockpit.vue'
import CCU from './components/CCU/Main'
import Scenes from './components/Pages/Scenes'
import SceneEdit from './components/Pages/SceneEdit'
import ScenePreview from './components/Pages/ScenePreview'
import SceneActive from './components/Pages/SceneActive'
import Program from './components/Pages/Program'
import ProgramView from './components/Pages/ProgramView'
import Timers from './components/Pages/Timers'

Vue.use(VueRouter)

const routes = [
  { path: '/device/:id', component: Device },
  { path: '/ccu', component: CCU },
  { path: '/cam/control', component: CamControl },
  { path: '/rec/control', component: RecControl },
  { path: '/cockpit', component: Cockpit, meta: { doubleMenu: true }},
  { path: '/scenes', component: Scenes },
  { path: '/scenes/_active', component: SceneActive, name: 'scene_active', meta: {fullscreen: true}},
  { path: '/scenes/:id', component: SceneEdit, name: 'scene_edit'},
  { path: '/scenes/:id/preview', component: ScenePreview, name: 'scene_preview', meta: { fullscreen: true }},
  { path: '/programs', component: Program },
  { path: '/programs/:id', component: ProgramView },
  { path: '/timers', component: Timers },
  { path: '/config', component: Config, meta: { offline: true }},
  { path: "/", component: Network }
]

const router = new VueRouter({
  routes,
})

export default router
