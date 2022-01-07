package homework

import "fmt"

// 注册
func Register() {
	for {
		var username, pwd, cpwd string
		fmt.Println("欢迎注册")
		fmt.Print("用户名:")
		fmt.Scanf("%s\n", &username)
		fmt.Print("密码:")
		fmt.Scanf("%s\n", &pwd)
		fmt.Print("确认密码:")
		fmt.Scanf("%s\n", &cpwd)
		if username == "" || pwd == "" || cpwd == "" {
			fmt.Println("输入不得为空")
			continue
		}
		if pwd != cpwd {
			fmt.Println("两次输入的秘密不一致")
			continue
		}
		fmt.Println("注册成功")
		break
	}
}
