
import Home from "../src/components/Home/Home"
import Stats from "../src/components/Stats/Stats"
import VueRouter from 'vue-router'
const routes = [
    { path: "/" , redirect: "/Home"},
    {
        path: "/Home",
        component: Home,
    },
    {
        path: "/Stats",
        component: Stats,
    }
]

const router = new VueRouter({
    // Removes # from URL
    mode: "history",
    routes: routes
});

export default router;