<template>
  <div>
    <div class="logo">
      <span>My Blog</span>
    </div>

    <div>
      <el-row class="tac">
        <el-menu default-active="1" class="el-menu-vertical-demo" background-color="#011582" text-color="#ffffff">
          <el-menu-item index="index" @click="goToPage">
            <i class="el-icon-discover"></i>
            <span slot="title">仪表盘</span>
          </el-menu-item>
          
          <el-submenu index="2" >
            <template slot="title">
              <i class="el-icon-document"></i>
              <span>文章管理</span>
            </template>
            <el-menu-item-group>
              <el-menu-item index="addart" @click="goToPage">
                <i class="el-icon-edit-outline"></i>
                <span>编写文章</span>
              </el-menu-item>
              <el-menu-item index="artlist" @click="goToPage">
                <i class="el-icon-document-copy"></i>
                <span>文章列表</span>
              </el-menu-item>
            </el-menu-item-group>
          </el-submenu>
          
          <el-menu-item index="catelist" @click="goToPage">
            <i class="el-icon-tickets"></i>
            <span slot="title">分类列表</span>
          </el-menu-item>
          
          <el-menu-item index="userlist" @click="goToPage">
            <i class="el-icon-user"></i>
            <span slot="title">用户列表</span>
          </el-menu-item>
        </el-menu>
      </el-row>
    </div>
  </div>
</template>

<script>
  export default {
    name: "Nav",
    methods:{
      goToPage(item) {
        this.$router.push('/admin/' + item.index).catch((err) => err) // 防止多次点击一个菜单报错
      }
    },
  }
</script>

<style scoped>
  .logo {
    height: 32px;
    margin: 16px;
    background-color: #eeeeee;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 18px;
  }
</style>
