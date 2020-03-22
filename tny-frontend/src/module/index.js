import Vue from 'vue'
import Vuex from 'vuex'
import { urlStore } from './urlStore'
import { configStore } from './config'
Vue.use(Vuex)
export default new Vuex.Store({
    modules: {
        urlStore: urlStore,
        configStore: configStore
    }
})
