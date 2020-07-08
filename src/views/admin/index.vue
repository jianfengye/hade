<template>
  <v-app>
    <v-app-bar
      app
      clipped-left
    >
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
      <span class="title ml-3 mr-5">运营管理后台</span>
      <v-spacer></v-spacer>
    </v-app-bar>
    <v-navigation-drawer
        v-model="drawer"
        app
        clipped
        color="grey lighten-4"
      >
      <v-list
        >
          <v-item-group
          >
            <v-list-item
              v-for="(item,i) in items"
              :key="i"
              :to="item.path"
            >
            <v-list-item-avatar>
              <v-icon
                :class="[item.iconClass]"
                v-text="item.icon"
              ></v-icon>
            </v-list-item-avatar>
              <v-list-item-content>
                <v-list-item-title class="grey--text">
                  {{ item.text }}
                </v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-item-group>
        </v-list>
    </v-navigation-drawer>
    <v-main>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-main>
    <v-footer app>
    </v-footer>
  </v-app>
</template>

<script>
import {adminMenu} from "@/router.js";
let menu = adminMenu();

export default {
  name: 'Admin',

  components: {
  },

  data: () => ({
    drawer: null,
    items: menu,
  }),
};
</script>

<style scoped>
#keep .v-navigation-drawer__border {
  display: none
}
</style>