<template>
  <v-app id="inspire">
    <v-app id="inspire">
      <v-main>
        <v-container
          class="fill-height"
          fluid
        >
          <v-row
            align="center"
            justify="center"
          >
            <v-col
              cols="12"
              sm="8"
              md="4"
            >
              <v-card class="elevation-12">
                <v-toolbar
                  color="primary"
                  dark
                  flat
                >
                  <v-toolbar-title>登陆</v-toolbar-title>
                  <v-spacer></v-spacer>
                
                </v-toolbar>
                <v-card-text>
                  <v-form>
                    <v-text-field
                      label="用户名"
                      name="login"
                      v-model="name"
                      prepend-icon="mdi-account"
                      type="text"
                    ></v-text-field>
  
                    <v-text-field
                      id="password"
                      label="密码"
                      v-model="password"
                      name="password"
                      prepend-icon="mdi-lock"
                      type="password"
                    ></v-text-field>
                  </v-form>
                </v-card-text>
                <v-card-actions>
                  <v-spacer></v-spacer>
                  <v-btn @click="submit" color="primary">登陆</v-btn>
                  <v-btn @click="clear">clear</v-btn>
                </v-card-actions>
              </v-card>
            </v-col>
          </v-row>
        </v-container>
      </v-main>
    </v-app>
  </v-app>
</template>

<script>
import {login} from "@/api/user.js"
export default {
  name: "Login",

  data: () => ({
    name : '',
    password: '',
  }),

  methods: {
    submit() {
      login({
        name: this.name, 
        password: this.password
        }).then(resData => {
          // 成功跳转到首页
          console.log(resData)
          this.$router.push("/")
        }).catch(err => {
          console.log(err)
        })
    },
    clear(){

    }
  }
}
</script>