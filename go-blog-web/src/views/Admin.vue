<template>
  <el-container class="container">
    <!-- 左侧导航栏 -->
    <el-aside width="200px" class="nav"> <Nav></Nav> </el-aside>
      
    <!-- 右侧内容栏 -->
    <el-container>
      <!-- 头部 -->
      <el-header class="headerBtn"> <Header></Header> </el-header>
      <!-- 内容 -->
      <el-main> <router-view></router-view> </el-main>
      <!-- 底部 -->
      <el-footer class="footer"> <Footer></Footer> </el-footer>
    </el-container>
  </el-container>
</template>

<script>
  import Header from "../components/admin/Header";
  import Footer from "../components/admin/Footer";
  import Nav from "../components/admin/Nav";

  export default {
    components: {Nav, Footer, Header}
  }
</script>

<style scoped>
  .container {
    height: 100%;
  }

  .headerBtn {
    background-color: #011582;
    display: flex;
    justify-content: flex-end;
    align-items: center;
  }

  .nav {
    background-color: #011582;
  }

  .footer {
    height: 40px!important;
  }
</style>
