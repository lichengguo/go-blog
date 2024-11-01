<template>
  <el-card>
    <h3>{{id? '编辑文章':'添加文章'}}</h3>
    
    <el-form size="mini" :model="ruleForm" :rules="rules" ref="ruleForm" label-position="top" hide-required-asterisk>
      <el-form-item label="文章标题" prop="title" style="width: 30%">
        <el-input v-model="ruleForm.title"></el-input>
      </el-form-item>

      <el-form-item label="文章分类" prop="cid">
        <el-select v-model="ruleForm.cid" placeholder="请选择分类" clearable>
          <el-option v-for="item in ruleForm.cateList" :key="item.id" :label="item.name" :value="item.id"></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="文章描述" prop="desc" style="width: 80%">
        <el-input v-model="ruleForm.desc"></el-input>
      </el-form-item>

      <el-form-item label="文章缩略图" prop="img">
        <el-upload :action="upUrl" :headers="importHeaders" :on-success="uploadPic" :limit="1" :on-exceed="handleExceed" multiple>
          <el-button size="small" type="primary">点击上传</el-button>
        </el-upload>
        <template v-if="id">
          <img :src="ruleForm.img" alt="图片不见了" style="height: 80px; width:100%"/>
        </template>
      </el-form-item>

      <el-form-item label="文章内容" prop="content">
        <quill-editor ref="text" v-model="ruleForm.content" class="myQuillEditor" :options="editorOption" />
      </el-form-item>

      <el-form-item>
        <el-button size="mini" type="primary" @click="submitForm('ruleForm')">{{id? '更新':'提交'}}</el-button>
        <el-button size="mini" @click="resetForm('ruleForm')">取消</el-button>
      </el-form-item>
    </el-form>
  </el-card>
</template>
<script>
  import {Url} from "../../plugin/http";
  export default {
    created() {
      this.getCateList()
      // 获取从文章列表页面点击编辑按钮传递过来的ID
      this.id=this.$route.query.id; 
      // 如果this.id不为空 那么就是编辑文章
      if (this.id) {
        this.getArtInfo(this.id)
      }
    },
    data() {
      return {
        id:'',                // 文章列表页面编辑按钮传递ID
        // 富文本编辑器
        editorOption: {
          placeholder: "请在这里输入",
        } ,
        upUrl:Url + 'upload', // 上传缩略图的URL
        // 获取token，添加到请求头部
        importHeaders:{
          Authorization: `Bearer ${window.sessionStorage.getItem('ginblog_token')}`,
        },
        // 表单数据
        ruleForm: {
          title: '',                // 文章标题
          cid:'',                   // 文章分类的ID
          desc:'',                  // 文章描述
          content:'',               // 文章内容
          img:'',                   // 缩略图链接
          // 文章分类按钮
          cateList: [],             // 所有的分类列表数据
        },
        // 数据验证
        rules: {
          title: [
            {required: true, message: '文章标题不能为空', trigger: 'blur'}
          ],
          desc: [
            {required: true, message: '文章描述不能为空', trigger: 'blur'}
          ],
          content: [
            {required: true, message: '文章内容不能为空', trigger: 'blur'}
          ],
          cid: [
            {required: true, message: '请选择分类', trigger: 'change'}
          ],
        }
      };
    },
    methods: {
      // 提交&&更新按钮
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            // 新增提交文章
            if (!this.id) {
              this.$http.post('article/add', this.ruleForm).then(response =>{
                if (response.data.status !== 200) return this.$message.error(response.data.message)
                this.$router.push('/admin/artlist')
                this.$message.success('添加文章成功')
              }).catch(error =>{
                console.log(error)
              })
            } else {
              // 编辑文章
              this.$http.put(`article/${this.id}`, this.ruleForm).then(response =>{
                if (response.data.status !== 200) return this.$message.error(response.data.message)
                this.$router.push('/admin/artlist')
                this.$message.success('更新文章成功')
              }).catch(error =>{
                console.log(error)
              })
            }
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },

      // 取消
      resetForm(formName) {
        this.$refs[formName].resetFields();
        this.$router.push('/admin/artlist')
      },

      // 获取所有的分类信息
      getCateList() {
        this.$http.get('category').then(response => {
          if (response.data.status !== 200) return this.$message.error(response.data.message)
          this.ruleForm.cateList = response.data.data
        }).catch(error => {
          console.log(error)
        })
      },

      // 只允许上传一张缩略图
      handleExceed(files, fileList) {
        return this.$message.warning(`当前限制选择 1 个文件，本次选择了 ${files.length} 个文件，
        共选择了 ${files.length + fileList.length} 个文件`);
      },

      // 获取上传文件成功以后服务器返回的url链接
      uploadPic(response){
        if (response.status !== 200) {
          console.log("==="+response.message)
          return this.$message.error('文件上传失败！！！')
        }
        this.$message.success('OK')
        this.ruleForm.img = response.url  // 从服务器拿取到上传图片的链接地址，赋值到表单数据
      },

      // 获取一篇文章的详情，加载到页面
      getArtInfo(id){
        this.$http.get(`article/info/${id}`).then(response =>{
          if (response.data.status !== 200) return this.$message.error(response.data.message)
          this.ruleForm.title = response.data.data.title
          this.ruleForm.cid = response.data.data.Category.id
          this.ruleForm.desc = response.data.data.desc
          this.ruleForm.img = response.data.data.img
          this.ruleForm.content = response.data.data.content
        }).catch(error =>{
          console.log(error)
        })
      },
    }
  }
</script>
