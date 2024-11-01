package model

import (
	"encoding/base64"
	"fmt"
	"go-blog-server/utils/errmsg"

	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

// 用户

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=10" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=1" label:"角色码"`
}

// 如果用户表user中没有数据，默认创建一个管理员
func DefaultUserAdd() {
	var count int64
	var user User
	db.Table("user").Count(&count) // 查询数据库user表是否存在数据
	if count == 0 {
		// 数据库没有记录
		// 创建一条默认的记录
		user = User{
			Username: "admin",
			Password: "123456",
			Role:     1,
		}
		code := CreateUser(&user)
		if code != errmsg.SUCCSE {
			panic("初始化管理用户失败")
		}
	}
}

// 查询用户是否存在
func CheckUser(name string) int {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	// ID>0 证明数据库有记录
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED // 1001
	}
	return errmsg.SUCCSE
}

// 根据ID查询单个用户
func GetUser(id int) (User, int) {
	var user User
	err := db.Where("ID = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCSE
}

// 新增用户
func CreateUser(data *User) int {
	// 验证客户端传过来的数据的合法性
	if data.Username == "" || data.Password == "" {
		return errmsg.ERROR
	}
	//data.Password = ScryptPw(data.Password) // 密码加密
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// 查询用户列表
func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	// 可以按照名字搜索用户
	if username != "" {
		db.Where("username LIKE ?", "%"+username+"%").Find(&users).Count(&total)
		err := db.Where("username LIKE ?", "%"+username+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, 0
		}
	} else {
		db.Find(&users).Count(&total)
		err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error // 分页 limit offset 为-1时不做分页
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, 0
		}
	}
	return users, total
}

// 编辑用户时，检测用户名是否存在
func EditCheckUser(id int, name string) int {
	// 编辑用户的时候，可以不修改用户名称，只修改权限
	// 所以要先排除自己ID本身的用户名同名的情况
	var users []User
	var total int64
	err := db.Where("username = ? AND id != ?", name, id).Find(&users).Count(&total).Error
	if err != nil {
		return errmsg.ERROR_USERNAME_USED
	}
	if total < 1 {
		return errmsg.SUCCSE
	}
	return errmsg.ERROR_USERNAME_USED
}

// 编辑用户
func EditUser(id int, data *User) int {
	var user User
	var pwd User
	// 这里使用map的方式去更新数据库的内容
	// struct 不会更新零值
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	db.Select("password").Where("id = ?", id).First(&pwd)
	if pwd.Password != data.Password {
		// 不相等 证明前端改动密码了，所以重新加密一下
		data.Password = ScryptPw(data.Password)
	}
	maps["password"] = data.Password // 更新密码
	err := db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除
func DeleteUser(id int) int {
	var user User
	// gorm如果使用了gorm.model，那么会在数据库软删除
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 钩子函数
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return nil
}

// 密码加密
func ScryptPw(password string) string {
	// key加密后的长度
	const KeyLen = 8
	// 加盐
	salt := make([]byte, 8)
	salt = []byte{'a', 'l', 'a', 't', 'o', 'm', 1, 2}
	// 调用加密算法
	// func Key(password, salt []byte, N, r, p, keyLen int) ([]byte, error)
	// 建议参数为N=32768、r=8和p=1
	// 参数N、r和p应该随着内存延迟和CPU并行度的增加而增加
	HashPw, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, KeyLen)
	if err != nil {
		fmt.Println("密码加密失败", err)
		return ""
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// 登录验证
func CheckLogin(username, password string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	// 密码验证
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	// 角色验证
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCSE
}
