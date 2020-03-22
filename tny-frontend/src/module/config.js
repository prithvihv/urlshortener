
export const configStore = {
    state: {
        app_server_host: "localhost",
        app_server_port: 9000
    },
    getters: {
        getHost: state => state.app_server_host,
        getAppServerURL: state => "http://www." + state.app_server_host + ":" + state.app_server_port + "/"
    }
}