import Vue from 'vue'
import VueRouter from 'vue-router'
import Vuex from 'vuex'
import App from './App.vue'
import router from "./router"
import store from "./module/index"

Vue.config.productionTip = false

Vue.use(Vuex)
Vue.use(VueRouter)

new Vue({
  render: h => h(App),
  router: router,
  store: store,
}).$mount('#app')
