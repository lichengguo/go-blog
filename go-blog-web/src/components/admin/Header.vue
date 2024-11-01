<template>
  <el-button size="small" type="danger" @click="loginOut">退出</el-button>
</template>

<script>
    export default {
      name: "Header",
      methods: {
        loginOut(){
          // 删除浏览器存储在本地的token
          window.sessionStorage.clear('ginblog_token')
          // 跳转到登录页面
          this.$router.push('/login')
        }
      },
    }
</script>

<style scoped>

</style>
