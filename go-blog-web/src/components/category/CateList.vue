<template>
  <div>
    <el-card style="margin: 0 auto">
      <!-- 添加文章分类 -->
      <el-row>
        <el-col :span="2">
          <el-button size="small" type="primary" @click="centerDialogVisible = true">添加</el-button>
        </el-col>
      </el-row>
      
      <!-- 表格数据 -->
      <el-row style="margin-top: 10px">
        <el-table border style="width: 100%" :data="tableData" size="small">
          <el-table-column prop="id" label="ID"></el-table-column>
          <el-table-column prop="name" label="分类名称"></el-table-column>
          <el-table-column prop="operation" label="操作">
            <template slot-scope="scope">
              <el-button size="mini" type="primary" @click="handleEdit(scope.row)">编辑</el-button>
              <el-button size="mini" type="danger" @click="handleDelete(scope.row)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-row>
      
      <!-- 分页 -->
      <el-row>
        <el-col :span="12" :offset="12">
          <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :page-sizes="data_page_sizes"
            :page-size="data_page_size"
            layout="total, sizes, prev, pager, next, jumper"
            :total="data_total">
          </el-pagination>
        </el-col>
      </el-row>
    </el-card>

    <!-- 添加分类 -->
    <el-dialog title="添加分类" :visible.sync="centerDialogVisible" width="30%" center>
      <el-form size="mini" @submit.native.prevent :model="ruleForm" :rules="rules" ref="ruleForm" :hide-required-asterisk="true">
        <el-form-item prop="name" label="分类名称">
          <el-input size="mini" v-model="ruleForm.name" clearable></el-input>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="addCate('ruleForm')">确认</el-button>
          <el-button size="mini" @click="resetForm('ruleForm')">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>

    <!-- 编辑分类 -->
    <el-dialog title="编辑分类" :visible.sync="editcenterDialogVisible" width="30%" center>
      <el-form size="mini" :model="editruleForm" :rules="rules" ref="editruleForm" :hide-required-asterisk="true">
        <el-form-item label="分类ID">
          <el-input size="mini" :disabled="true" v-model="editruleForm.ID" clearable></el-input>
        </el-form-item>
        <el-form-item prop="name" label="分类名称">
          <el-input size="mini" v-model="editruleForm.name" clearable></el-input>
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="editCate('editruleForm')">确认</el-button>
          <el-button size="mini" type="warning" @click="resetForm('editruleForm')">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
  export default {
    name: "CateList",
    data() {
      return {
        tableData: [],                       // 显示到表单的数据
        data_page_sizes: [5, 10, 20, 50],    // 每页显示多少条数据
        data_page_size: 5,                   // 每页默认显示多少条数据
        data_page_num: 1,                    // 初始页
        data_total: 0,                       // 总数据
        centerDialogVisible: false,          // 添加分类提示框显示
        editcenterDialogVisible: false,      // 编辑分类提示框显示
        // 添加分类弹框的form表单数据
        ruleForm: {
          name: '',
        },
        // 编辑分类弹框form表单数据
        editruleForm: {
          ID: '',
          name: '',
        },
        // 校验规则
        rules: {
          // 分类
          name: [
            {required: true, message: '请输入分类', trigger: 'blur'},
            {min: 1, max: 10, message: '长度在1到10个字符', trigger: 'blur'}
          ],
        },
      }
    },

    // 页面加载完成以后，自动加载数据到表格 vue生命周期
    created() {
      this.getCateList()
    },

    methods: {
      // 从服务器拉取分类数据
      getCateList() {
        this.$http.get('category', {
          params: {
            pagesize: this.data_page_size,
            pagenum: this.data_page_num,
          }
        }).then(response => {
          if (response.data.status != 200) return this.$message.error(response.data.message)
          this.tableData = response.data.data
          this.data_total = response.data.total
        }).catch(error => {
          console.log(error)
        })
      },

      // 分页
      handleSizeChange(val) {
        // console.log(`每页 ${val} 条`);
        this.data_page_size = parseInt(`${val}`);
        this.getCateList();
      },
      handleCurrentChange(val) {
        // console.log(`当前页: ${val}`);
        this.data_page_num = parseInt(`${val}`);
        this.getCateList();
      },

      // 删除分类
      handleDelete(row) {
        if (window.confirm("确认删除分类 [" + row.name + '] 吗?')) {
          this.$http.delete("category/" + row.id).then(res => {
            if (res.data.status != 200) return this.$message.error(res.data.message)
            this.$message.success('删除成功!')
            this.getCateList();
          })
        } else {
          this.$message.info('删除已取消!')
        }
      },

      // 添加分类
      // 确认
      addCate(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            // 通过验证
            this.$http.post('category/add', {
              name: this.ruleForm.name,
            }).then(response => {
              if (response.data.status != 200) return this.$message.error(response.data.message)
              this.$message.success(response.data.message)
              this.$refs[formName].resetFields(); // 清空表单里的数据
              this.centerDialogVisible = false    // 关闭弹出框
              this.getCateList()                  // 重新获取一下分类列表数据
            }).catch(error => {
              this.$message.error(error)
            })
          } else {
            return false
          }
        })
      },
      // 取消
      resetForm(formName) {
        this.$refs[formName].resetFields();  // 清空表单里的数据
        this.centerDialogVisible = false     // 取消添加弹框
        this.editcenterDialogVisible = false // 取消编辑弹框
      },

      // 编辑分类
      handleEdit(row) {
        this.editcenterDialogVisible = true // 打开弹窗
        // 填充数据
        this.editruleForm.ID = row.id
        this.editruleForm.name = row.name
      },
      // 确认按钮
      editCate(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            // 验证通过
            this.$http.put(`category/${this.editruleForm.ID}`, {
              name: this.editruleForm.name,
            }).then(response => {
              if (response.data.status != 200) return this.$message.error(response.data.message)
              this.$message.success(response.data.message)
              this.editcenterDialogVisible = false // 取消编辑弹框
              this.getCateList()
            })
          } else {
            return false
          }
        })
      },
    }
  }
</script>

