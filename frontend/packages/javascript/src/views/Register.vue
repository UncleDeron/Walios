<template>
  <div class="register">
    <!-- Logo -->
    <!-- <img class="logo" alt="Vue logo" src="../assets/logo.png" /> -->
    <HelloWorld msg="注册蛙聊" />
    <!-- Bottom button -->
    <!-- 底部按钮 -->
    <el-form class="register-form" ref="form" :show-message="false" :model="form" :rules="rules" @validate="validateItem">
      <el-form-item prop="userAccount">
        <el-input v-model="form.userAccount" placeholder="账号"></el-input>
      </el-form-item>
      <el-form-item prop="nickName">
        <el-input v-model="form.nickName" placeholder="昵称"></el-input>
      </el-form-item>
      <el-form-item prop="email">
        <el-input v-model="form.email" placeholder="邮箱（用于重置密码）"></el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input type="password" v-model="form.password" show-password placeholder="密码（至少8位且必须包含字母、数字和特殊字符）"></el-input>
      </el-form-item>
      <el-form-item prop="confirmPwd">
        <el-input type="password" v-model="form.confirmPwd" show-password placeholder="确认密码"></el-input>
      </el-form-item>
    </el-form>
    <div class="button-list">
      <el-button type="primary" @click="register">注册</el-button>
      <el-button type="text">取消</el-button>
    </div>
    
  </div>
</template>

<script>
import HelloWorld from "@/components/HelloWorld.vue";
export default {
  name: "Register",
  components: {
    HelloWorld,
  },
  inject: ["app"],
  data() {
    return {
      form: {
        userAccount: "",
        nickName: "",
        email: "",
        password: "",
        confirmPwd: "",
      },
      rules: {
        userAccount: [
          { required: true, message: "请输入账号", trigger: "blur" },
          { min: 4, max: 12, message: "长度在 4 到 12 个字符", trigger: "blur" },
          { pattern: /^[a-zA-Z0-9_]+$/, message: "只能是数字、字母、下划线", trigger: "blur" },
        ],
        nickName: [
          { required: true, message: "请输入昵称", trigger: "blur" },
        ],
        email: [
          { required: true, message: "请输入邮箱", trigger: "blur" },
          { type: "email", message: "请输入正确的邮箱地址", trigger: "blur" },
        ],
        password: [
          { required: true, message: "请输入密码", trigger: "blur" },
          { min: 6, max: 12, message: "长度在 6 到 12 个字符", trigger: "blur" },
          { pattern: /^\S*(?=\S{8,})(?=\S*\d)(?=\S*[A-Za-z])(?=\S*[!@#$%^&*? ])\S*$/, message: "至少8位且必须包含字母、数字和特殊字符", trigger: "blur" },
        ],
        confirmPwd: [
          { required: true, message: "请输入确认密码", trigger: "blur" },
          { min: 6, max: 12, message: "长度在 6 到 12 个字符", trigger: "blur" },
          { validator: this.validateConfirmPwd, trigger: "blur" },
        ]
      },
    };
  },
  methods: {
    register() {
      this.$refs.form.validate((valid, errors) => {
        if (valid) {
          go.main.App.Register(this.form.username, this.form.password).then((result) => {
            console.log(result)
          });
        } else {
          console.log("error submit!!");
          this.app.showSystemNotification("注册信息不完整", "danger");
        }
      });
      
    },
    validateConfirmPwd(rule, value, callback) {
      if (value === this.form.password) {
        callback();
      } else {
        callback(new Error("两次输入密码不一致"));
      }
    },
    validateItem(prop, valid, msg) {
      if (!valid) {
        this.app.showSystemNotification(msg, "danger");
      }
    }
  },
  watch: {
    
  },
};
</script>
<style lang="scss">
.register {
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  .register-form {
    width: 300px;
    margin: 0 auto;
  }
  .button-list {
    width: 300px;
    margin: 0 auto 60px;

    .el-button {
      display: block;
      width: 100%;
      margin: 10px 0;
    }
  }
}
</style>
