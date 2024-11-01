<template>
  <div class="container">
    <div class="loginBox">
      <el-form :model="loginFormData" :rules="rules" ref="loginFormData" class="loginFormData">
        <!-- 用户名框 -->
        <el-form-item prop="username">
          <el-input prefix-icon="el-icon-user" v-model="loginFormData.username" placeholder="请输入用户名" clearable></el-input>
        </el-form-item>
        <!-- 密码框 -->
        <el-form-item prop="password">
          <el-input @keyup.enter.native="submitForm('loginFormData')" type="password" prefix-icon="el-icon-position"
                    v-model="loginFormData.password" placeholder="请输入密码" clearable></el-input>
        </el-form-item>
        <!-- 按钮 -->
        <el-form-item class="loginBtn">
          <el-button type="primary" @click="submitForm('loginFormData')">登 录</el-button>
          <el-button type="warning" @click="resetForm('loginFormData')">取 消</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        // 登录表单数据
        loginFormData: {
          username: '',
          password: '',
        },
        // 表单数据校验规则
        rules: {
          username: [
            {required: true, message: '请输入账号', trigger: 'blur'},
            {min: 4, max: 10, message: '长度在4到10个字符', trigger: 'blur'}
          ],
          password: [
            {required: true, message: '请输入密码', trigger: 'blur'},
            {min: 6, max: 20, message: '长度在6到20个字符', trigger: 'blur'}
          ],
        }
      };
    },
    methods: {
      // 登录按钮
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            // 表单数据验证通过，则提交请求到后端服务
            this.$http.post('login', this.loginFormData).then(response => {
              if (response.data.status != 200) {
                this.$message.error(response.data.message)
                return
              }
              // 把token保存到浏览器
              window.sessionStorage.setItem('ginblog_token', response.data.token)
              // 跳转到amdin/index页面
              this.$router.push('admin/index') 
            }).catch(error => {
              this.$message.error(error)
            })
          } else {
            return false;
          }
        });
      },
      // 取消按钮
      resetForm(formName) {
        this.$refs[formName].resetFields();
      },
    }
  }
</script>

<style scoped>
  .container {
    height: 100%;
    background-color: #282c34;
  }

  .loginBox {
    width: 450px;
    height: 280px;
    background-color: white;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -80%);
    border-radius: 10px;
  }

  .loginFormData {
    width: 100%;
    position: absolute;
    bottom: 10%;
    padding: 0 20px;
    box-sizing: border-box;
  }

  .loginBtn {
    display: flex;
    justify-content: flex-end;
  }
</style>
