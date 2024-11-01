<template>
  <div>
    <el-card style="margin: 0 auto">
      <!-- 搜索和新增 -->
      <el-row>
        <!-- 搜索按钮 -->
        <div class="search">
          <el-col :span="6">
            <el-input @keyup.enter.native="searchArtList" placeholder="请输入文章标题" size="mini" suffix-icon="el-icon-search" v-model="searchArtname"/>
          </el-col>
          <el-col :span="2">
            <el-button size="mini" type="primary" @click="searchArtList">搜索</el-button>
          </el-col>
        </div>
        <!-- 新增按钮 -->
        <el-col :span="2">
          <el-button size="mini" type="primary" @click="$router.push('/admin/addart')">新增</el-button>
        </el-col>
        <!-- 分类搜索按钮 -->
        <el-col :span="3">
          <el-select size="mini" v-model="cateValue" placeholder="请选择分类" @change="cateAtrList" clearable @clear="clearCateSelect">
            <el-option v-for="item in cateList" :key="item.id" :label="item.name" :value="item.id"></el-option>
          </el-select>
        </el-col>
      </el-row>

      <!-- 表格数据 -->
      <el-row style="margin-top: 10px">
        <el-table border style="width: 100%" :data="tableData" size="small">
          <el-table-column min-width="15%" prop="ID" label="ID"></el-table-column>
          <el-table-column min-width="35%" prop="Category.name" label="分类"></el-table-column>
          <el-table-column prop="title" label="文章标题"></el-table-column>
          <el-table-column prop="desc" label="文章描述"></el-table-column>
          <el-table-column prop="img"  label="缩略图">
            <template slot-scope="scope">
              <img :src="scope.row.img" alt="图片不见了" style="height: 80px; width:100%">
            </template>
          </el-table-column>
          <el-table-column min-width="60%" prop="operation" label="操作">
            <template slot-scope="scope">
              <router-link :to="{path:'/admin/addart',query:{id:scope.row.ID}}">
                <el-button size="mini" type="primary">编辑</el-button>
              </router-link>
              <el-button size="mini" type="danger" @click="handleDelete(scope.row)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-row>

      <!-- 分页 -->
      <el-row>
        <el-col :span="12" :offset="12">
          <el-pagination
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :page-sizes="data_page_sizes"
            :page-size="data_page_size"
            layout="total, sizes, prev, pager, next, jumper"
            :total="data_total">
          </el-pagination>
        </el-col>
      </el-row>
    </el-card>
  </div>
</template>

<script>
  export default {
    name: "ArtList",
    data() {
      return {
        tableData: [],                       // 显示到表单的数据
        data_page_sizes: [5, 10, 20, 50],    // 每页显示多少条数据
        data_page_size: 5,                   // 每页默认显示多少条数据
        data_page_num: 1,                    // 初始页
        data_total: 0,                       // 总数据
        searchArtname: '',                   // 搜索文章
        // 文章分类
        cateList: [],                         // 所有的分类列表数据
        cateValue: '',                        // 分类ID
      }
    },

    // 页面加载完成以后，自动加载数据
    created() {
      this.getArtList()   // 加载文章数据到表格
      this.getCateList()  // 加载所有分类到分类搜索按钮
    },

    methods: {
      // 从服务端拉取所有文章列表数据
      getArtList() {
        this.$http.get('article', {
          params: {
            pagesize: this.data_page_size, // 每页多少条
            pagenum: this.data_page_num,   // 第几页
            title: this.searchArtname,     // 搜索文章
          }
        }).then(response => {
          if (response.data.status !== 200) return this.$message.error(response.data.message)
          this.tableData = response.data.data    // 文章数据列表
          this.data_total = response.data.total  // 文章数据总量
        }).catch(error => {
          console.log(error)
        })
      },

      // 分页
      handleSizeChange(val) {
        // 判断分类ID是否非空
        if (this.cateValue !== '') {
          // 分类搜索的时候 分页走这里
          this.data_page_size = parseInt(`${val}`); // 每页多少条数据
          this.cateAtrList()
          // 每次选择分类的时候，重置一下分页的默认值
          // 确保下一次选择其他分类搜索不出现没有数据的BUG
          this.data_page_size = 5 // 每页数据量
          return
        }
        this.data_page_size = parseInt(`${val}`); // 每页多少条数据
        this.getArtList();
      },
      handleCurrentChange(val) {
        // 判断分类ID是否非空
        if (this.cateValue !== '') {
          // 分类搜索的时候 分页走这里
          this.data_page_num = parseInt(`${val}`); // 当前页码
          // 获取该分类下的文章
          this.cateAtrList()
          // 每次选择分类的时候，重置一下分页的默认值
          // 确保下一次选择其他分类搜索不出现没有数据的BUG
          this.data_page_num = 1 // 页码
          return
        }
        this.data_page_num = parseInt(`${val}`); // 当前页码
        this.getArtList();
      },

      // 搜索文章
      searchArtList() {
        this.cateValue = "" // 重置分类搜索选择的内容
        this.getArtList()
      },

      // 删除文章
      handleDelete(row) {
        if (window.confirm("确认删除文章 [" + row.title + '] 吗?')) {
          this.$http.delete("article/" + row.ID).then(res => {
            if (res.data.status !== 200) return this.$message.error(res.data.message)
            this.$message.success('删除成功!')
            this.getArtList();
          })
        } else {
          this.$message.info('删除已取消!')
        }
      },

      // 获取所有的分类信息
      getCateList() {
        this.$http.get('category').then(response => {
          if (response.data.status !== 200) return this.$message.error(response.data.message)
          this.cateList = response.data.data
          this.data_total = response.data.total
        }).catch(error => {
          console.log(error)
        })
      },

      // 查询分类下所有文章
      cateAtrList() {
        this.searchArtname = ""  // 重置搜索框内容
        // 如果分类ID为空，不发送请求到服务端
        if (this.cateValue === "") return false
        this.$http.get(`article/list/${this.cateValue}`, {
          params: {
            pagesize: this.data_page_size,
            pagenum: this.data_page_num,
          }
        }).then(response => {
          if (response.data.status !== 200) return this.$message.error(response.data.message)
          this.tableData = response.data.data
          this.data_total = response.data.total
        }).catch(error => {
          console.log(error)
        })
      },
      // 清空分类选择 重新获取所有文章列表
      clearCateSelect() {
        this.cateValue = ''      // 重置分类ID 让分页的时候走其他分支
        this.searchArtname = ""  // 重置搜索框内容
        this.getArtList()
      },
    }
  }
</script>

