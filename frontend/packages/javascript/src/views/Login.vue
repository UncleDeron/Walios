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
        <el-input type="password" v-model="form.password" placeholder="密码"></el-input>
      </el-form-item>
    </el-form>
    <div class="button-list">
      <el-button type="primary" @click="login">登录</el-button>
      <el-button type="plain">注册</el-button>
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
        this.$message.error("请输入账号");
        return;
      }
      if (!this.form.password) {
        this.$message.error("请输入密码");
        return;
      }
      go.main.App.Login(this.form.username, this.form.password).then((result) => {
        console.log(result)
      });
    }
  }
};
</script>
<style lang="scss">
.login {
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
