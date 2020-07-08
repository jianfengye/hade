<template>
  <v-app>
    <v-app-bar
      app
      clipped-left
      color="amber"
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
          light
          rounded
        >
          <v-item-group
          >
            <v-subheader>选项</v-subheader>
            <v-list-item
              v-for="(item,i) in items"
              :key="i"
              :to="item.to"
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
export default {
  name: 'App',

  components: {
  },

  data: () => ({
    drawer: null,
    items: [
      { icon: 'folder', iconClass: 'grey lighten-1 white--text',  text: '管理面板', to: "/" },
      { icon: 'folder', iconClass: 'grey lighten-1 white--text',  text: '设置' ,to: "/setting"},
    ],
  }),
};
</script>

<style scoped>
#keep .v-navigation-drawer__border {
  display: none
}
</style>