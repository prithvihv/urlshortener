import Vue from "vue"
import axios from "axios"
export const urlStore = {
    state: {
        tnyStore: {}
    },
    getters: {
        tnyStore: state => state.tnyStore,
    },
    mutations: {
        newtnyStore(state, { tny, full }) {
            Vue.set(state.tnyStore, full, tny)
        }
    },
    actions: {
        async fetchShotURL({ rootGetters, commit }, url) {
            let appServerURL = rootGetters.getAppServerURL
            let fullURL = appServerURL
            console.log(fullURL)
            // let res = await axios.post(fullURL)
            let res = await axios.request({
                url: appServerURL,
                data: url,
                method: "POST"
            })
            console.log(res)
            commit("newtnyStore", { "tny": res.data.TinyURLuid, "full": res.data.FullURL })
        }
    }
}