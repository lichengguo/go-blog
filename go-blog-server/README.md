# myBlogGolang

### 介绍
Go 博客项目



### 技术栈

#### 前端
+ [Vue](https://cn.vuejs.org/v2/guide/)
+ [Element](https://element.eleme.cn/#/zh-CN/component/installation)
+ [Quill富文本](https://www.kancloud.cn/liuwave/quill/1434140)

#### 后端

+ [gin框架](https://www.kancloud.cn/shuangdeyu/gin_book/949415)

+ [gorm-v1](https://v1.gorm.io/zh_CN/docs/)
+ [gorm-v2](https://gorm.io/zh_CN/docs/)
+ [ini配置文件解析](https://ini.unknwon.io/docs/intro/getting_started)
+ [日志文件](https://github.com/sirupsen/logrus) 
+ [日志文件分割](https://pkg.go.dev/github.com/lestrrat-go/file-rotatelogs) 
+ [日志HOOK](https://pkg.go.dev/github.com/rifflock/lfshook) 

+ [跨域](https://github.com/gin-contrib/cors)
+ [jwt鉴权](https://godoc.org/github.com/dgrijalva/jwt-go#example-NewWithClaims--StandardClaims)
+ [scrypt加密](https://godoc.org/golang.org/x/crypto/scrypt)

+ [字段校验validator](https://github.com/go-playground/validator)
+ [mysql5.6]()
  

### 软件架构

软件架构说明



### 安装教程

#### 前端安装
使用vue-cli开发前端
vue-cli项目初始化
```shell script
lichengguo@lichengguodeMacBook-Pro web % pwd
/Users/lichengguo/go/src/my-blog-golang/web
lichengguo@lichengguodeMacBook-Pro web % vue init webpack admin
? Project name admin                      # 回车
? Project description A Vue.js project    # 回车
? Author lichengguo <1029612787@qq.com>   # 回车
? Vue build standalone
? Install vue-router? Yes
? Use ESLint to lint your code? No
? Set up unit tests No
? Setup e2e tests with Nightwatch? No
? Should we run `npm install` for you after the project has been created? (recommended) npm
```
需要安装的插件如下
+ Element # npm i element-ui -S
+ Quill   # npm install vue-quill-editor --save

运行项目命令

```shell
# 在该项目目录下
npm run dev
```

打包命令如下

```shell
npm run build
```



#### 后端安装

略.



### 使用说明

+ 后端程序第一次启动会在数据库创建一个默认的管理用户

  账号`admin`

  密码`123456` 





### 


