<template>
  <div>
    <el-card style="margin: 0 auto">
      <!-- 搜索和添加 -->
      <el-row>
        <div class="search">
          <el-col :span="6">
            <el-input @keyup.enter.native="searchUserList" placeholder="请输入内容" size="mini" v-model="searchUsername"/>
          </el-col>
          <el-col :span="2">
            <el-button size="mini" type="primary" @click="searchUserList">搜索</el-button>
          </el-col>
        </div>
        <el-col :span="2">
          <el-button size="mini" type="primary" @click="centerDialogVisible=true">添加</el-button>
        </el-col>
      </el-row>

      <!-- 表单数据 -->
      <el-row style="margin-top: 10px">
        <el-table border style="width: 100%" :data="tableData" size="small">
          <el-table-column prop="ID" label="ID"></el-table-column>
          <el-table-column prop="username" label="用户名"></el-table-column>
          <el-table-column prop="role" label="角色" :formatter="roleToSstr"></el-table-column>
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

    <!-- 添加用户 -->
    <el-dialog title="添加用户" :visible.sync="centerDialogVisible" width="30%" center>
      <el-form size="small" :model="ruleForm" :rules="rules" ref="ruleForm" class="loginForm" :hide-required-asterisk="true">
        <el-form-item prop="username" label="用户名">
          <el-input size="mini" v-model="ruleForm.username" clearable></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input size="mini" clearable type="password" v-model="ruleForm.password"></el-input>
        </el-form-item>
        <el-form-item size="mini" label="管理员" prop="role">
          <el-select v-model="ruleForm.role" placeholder="请选择">
            <el-option label="是" value="1"></el-option>
            <el-option label="否" value="2"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button  type="primary" @click="addUser('ruleForm')">确认</el-button>
          <el-button type="warning" @click="resetForm('ruleForm')">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>

    <!-- 编辑用户 -->
    <el-dialog title="编辑用户" :visible.sync="editcenterDialogVisible" width="30%" center>
      <el-form size="small" :model="editruleForm" :rules="rules" ref="editruleForm" :hide-required-asterisk="true">
        <el-form-item label="用户ID">
          <el-input size="mini" :disabled="true" v-model="editruleForm.ID" clearable></el-input>
        </el-form-item>

        <el-form-item prop="username" label="用户名">
          <el-input size="mini" v-model="editruleForm.username" clearable></el-input>
        </el-form-item>

        <el-form-item prop="password" label="密码">
          <el-input size="mini" clearable type="password" v-model="editruleForm.password"></el-input>
        </el-form-item>

        <el-form-item label="管理员" prop="role">
          <el-select size="mini" v-model="editruleForm.role" placeholder="请选择">
            <el-option label="是" value="1"></el-option>
            <el-option label="否" value="2"></el-option>
          </el-select>
        </el-form-item>
        
        <el-form-item>
          <el-button size="mini" type="primary" @click="editUser('editruleForm')">确认</el-button>
          <el-button size="mini" type="warning" @click="resetForm('editruleForm')">取消</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
  export default {
    name: "UserList",
    data() {
      return {
        tableData: [],                       // 显示到表单的数据
        data_page_sizes: [5, 10, 20, 50],    // 每页显示多少条数据
        data_page_size: 5,                   // 每页默认显示多少条数据
        data_page_num: 1,                    // 初始页
        data_total: 0,                       // 总数据
        searchUsername: '',                  // 搜索用户
        centerDialogVisible: false,          // 添加用户提示框显示
        editcenterDialogVisible: false,      // 编辑用户提示框显示
        // 添加用户弹框的form表单数据
        ruleForm: {
          username: '',
          password: '',
          role: '2',
        },
        // 编辑用户弹框form表单数据
        editruleForm: {
          ID: '',
          username: '',
          password: '',
          role: '',
        },
        // 校验规则
        rules: {
          // 用户名
          username: [
            {required: true, message: '请输入账号', trigger: 'blur'},
            {min: 4, max: 10, message: '长度在4到10个字符', trigger: 'blur'}
          ],
          // 密码
          password: [
            {required: true, message: '请输入密码', trigger: 'blur'},
            {min: 6, max: 20, message: '长度在6到20个字符', trigger: 'blur'}
          ],
          // 权限
          role: [
            {required: true, message: '请选择权限', trigger: 'change'}
          ],
        },
      }
    },

    // 页面加载完成以后，自动加载数据到表格，vue生命周期
    created() {
      this.getUserList()
    },

    methods: {
      // 从服务器拉取用户数据
      getUserList() {
        this.$http.get('users', {
          params: {
            pagesize: this.data_page_size,
            pagenum: this.data_page_num,
            username: this.searchUsername,
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
        this.getUserList();
      },
      handleCurrentChange(val) {
        // console.log(`当前页: ${val}`);
        this.data_page_num = parseInt(`${val}`);
        this.getUserList();
      },

      // 格式转换 权限数字转化为文字字符
      roleToSstr(row) {
        return row.role == 1 ? '管理员' : '订阅者'
      },

      // 搜索用户
      searchUserList() {
        this.getUserList();
      },

      // 删除用户
      handleDelete(row) {
        if (window.confirm("确认删除用户 [" + row.username + '] 吗?')) {
          this.$http.delete("user/" + row.ID).then(res => {
            if (res.data.status != 200) return this.$message.error(res.data.message)
            this.$message.success('删除成功!')
            this.getUserList();
          })
        } else {
          this.$message.info('删除已取消!')
        }
      },

      // 添加用户
      // 确认
      addUser(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            // 通过验证
            this.$http.post('user/add', {
              username: this.ruleForm.username,
              password: this.ruleForm.password,
              role: parseInt(this.ruleForm.role), // 需要传入整数
            }).then(response => {
              if (response.data.status != 200) return this.$message.error(response.data.message)
              this.$message.success(response.data.message)
              this.$refs[formName].resetFields(); // 清空表单里的数据
              this.centerDialogVisible = false    // 关闭弹出框
              this.getUserList()                  // 重新获取一下用户列表数据
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

      // 编辑用户
      handleEdit(row) {
        this.editcenterDialogVisible = true // 打开弹窗
        // 填充数据
        this.editruleForm.ID = row.ID
        this.editruleForm.username = row.username
        this.editruleForm.role = (row.role).toString()
        this.$http.get(`user/${row.ID}`).then(response => {
          if (response.data.status != 200) return console.log(response.data.message)
          this.editruleForm.password = response.data.data.password
        }).catch(error => {
          console.log(error)
        })
      },
      // 确认按钮
      editUser(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            // 验证通过
            this.$http.put(`user/${this.editruleForm.ID}`, {
              username: this.editruleForm.username,
              role: parseInt(this.editruleForm.role), // 需要传入整数
              password: this.editruleForm.password,
            }).then(response => {
              if (response.data.status != 200) return this.$message.error(response.data.message)
              this.$message.success(response.data.message)
              this.editcenterDialogVisible = false // 取消编辑弹框
              this.getUserList()
            })
          } else {
            return false
          }
        })
      },
    }
  }
</script>
