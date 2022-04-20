<template>
  <div class="login">
    <!-- Logo -->
    <!-- <img class="logo" alt="Vue logo" src="../assets/logo.png" /> -->
    <h1 class="logo">🐸</h1>
    <HelloWorld msg="欢迎使用蛙聊" />
    <!-- Bottom button -->
    <!-- 底部按钮 -->
    <el-form class="login-form" ref="form" :model="form">
      <el-form-item>
        <el-input v-model="form.username" placeholder="账号"></el-input>
      </el-form-item>
      <el-form-item>
        <el-input type="password" v-model="form.password" placeholder="密码" @keyup.native.enter="login"></el-input>
      </el-form-item>
    </el-form>
    <div class="button-list">
      <el-button type="primary" @click="login">登录</el-button>
      <el-button type="plain" @click="toRegister">注册</el-button>
    </div>
    
  </div>
</template>

<script>
import HelloWorld from "@/components/HelloWorld.vue";

export default {
  name: "Login",
  components: {
    HelloWorld,
  },
  inject: ["app"],
  data() {
    return {
      form: {
        username: "",
        password: "",
      },
    };
  },
  methods: {
    login() {
      if (!this.form.username) {
        this.app.showSystemNotification("请输入账号", "danger");
        return;
      }
      if (!this.form.password) {
        this.app.showSystemNotification("请输入密码", "danger");
        return;
      }
      go.main.App.Login(this.form.username, this.form.password).then((result) => {
        console.log(result)
      });
      // runtime.EventsEmit("login", this.form);

    },
    toRegister() {
      this.$router.push("/register");
    }
  }
};
</script>
<style lang="scss">
.login {
  height: 100%;
  padding-top: 60px;
  .logo {
    display: block;
    font-size: 64px;
    text-align: center;
  }
  .login-form {
    width: 300px;
    margin: 0 auto;
  }
  .button-list {
    width: 300px;
    margin: 0 auto;

    .el-button {
      display: block;
      width: 100%;
      margin: 10px 0;
    }
  }
}
</style>
