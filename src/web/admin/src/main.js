import Vue from 'vue'
import App from './App.vue'
import router from './router'

import './plugin/http'
import './plugin/antui'
import './assets/css/style.css'
import * as echarts from 'echarts'

Vue.prototype.$echarts = echarts
Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
