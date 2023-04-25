import Vue from 'vue'
import App from './App'
import { myHttp } from '@/util/api.js'
Vue.prototype.$myHttp = myHttp

Vue.config.productionTip = false

App.mpType = 'app'

const app = new Vue({
    ...App
})
app.$mount()
