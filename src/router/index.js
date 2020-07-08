import Vue from "vue"
import Router from "vue-router"
import Dashboard from "../components/dashboard/index"
import Setting from "../components/setting/index"

Vue.use(Router)

const routes = [
    { path: '/', component: Dashboard},
    { path: '/setting', component: Setting},
];

export default new Router({
    // mode: 'hash',
    // base: process.env.BASE_URL,
    routes: routes,
})