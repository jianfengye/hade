import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import {routes} from "./router.js";
import Router from "vue-router";

Vue.use(Router);
Vue.config.productionTip = false

let router = new Router({
  routes: routes,
});

new Vue({
  vuetify,
  router,
  render: h => h(App)
}).$mount('#app')
