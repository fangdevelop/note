package note

import (
	"fmt"
	"gonote/util"
	"sync"
)

//2.1 转义字符
func EscapedCharacters() {
	fmt.Println("\n1 双引号")
	fmt.Println("图灵祖师说:\"我不知道我是梦见自己是机器的图灵，还是梦见自己是图灵的机器\"")
	fmt.Println("\n2 反斜线")
	fmt.Println("\\\\电子邮件说\\项目已取消\\清理文档的时候\\我哭了")
	fmt.Println("\n3 警报声")
	fmt.Println("\a既使一个程序只有三行长，也总有一天需要去维护它\a")
	fmt.Println("\n4 退格")
	fmt.Println("三天不编程，，\b生活变得毫无意义")
	fmt.Println("\n5 换页")
	fmt.Println("一个程序员正在写软件\f他的手指在键盘上飞舞\f程序编译时没有一条错误信息\f运行起来就如同一阵微风")
	fmt.Println("\n6 回车")
	fmt.Println("我的感官很悠闲，我的精神自由地按照它自己的直觉前进\r我的程序如同是自己在写自己") //一般与\n配合使用：\n\r
	fmt.Println("\n7 制表符")
	fmt.Println("一一\t坨坨\t方块\n\r害羞的\t内向的\t腼腆的")
	fmt.Println("\n8 纵向制表符")
	fmt.Println("确实，有时候我会遇到难题。\v当我发现难题的时候，我会慢下来，安静地观察。\v然后我改变一行代码，困难就烟消云散")
}

//2.2 变量与常量
func VariablesAndConstants() {
	fmt.Println("\n1 变量")
	var v1 int
	var v2 int = 2
	var v3 = v2
	v1 = 1
	v4 := 4
	var (
		v5     = v1
		v6 int = 6
		v7 int
	)
	var v8, v9 int = v2, v3
	v10, v11 := v4, v5
	fmt.Printf("v6=%v,v7=%v,v8=%v,v9=%v,v10=%v,v11=%v\n", v6, v7, v8, v9, v10, v11)
	fmt.Println("\n2 常量")
	const (
		c1 = 8
		c2 = iota //当前行数，注意：行数是从0开始计算
		c3 = iota //iota可以当做正常数字运算
		c4        //默认值为上一行的值
		c5 = 12
		c6
	)
	fmt.Printf("c1=%v,c2=%v,c3=%v,c4=%v,c5=%v,c6=%v\n", c1, c2, c3, c4, c5, c6)
}

//2.3 基本数据类型
func BasicDataTypes() {
	fmt.Println("\n2.3.1 整数型")
	//int8, uint8, int16, uint16, int32, uint32, int64, uint64, int(默认), uint //int, uint为操作系统位数
	var (
		n1        = 0b0101 //0b/0B表示2进制
		n2 int8   = 0o77   //0o/0O/0表示8进制
		n3 uint16 = 0xAF   //0x/0X表示16进制
	)
	fmt.Printf("n1=%v,type is %T\n", n1, n1)
	fmt.Printf("n2=%v,type is %T\n", n2, n2)
	fmt.Printf("n3=%v,type is %T\n", n3, n3)

	fmt.Println("\n2.3.2 浮点型")
	//float32, float64(默认)
	var (
		f1         = 1.0
		f2 float32 = 1
	)
	fmt.Printf("f1=%v,type is %T\n", f1, f1)
	fmt.Printf("f2=%v,type is %T\n", f2, f2)

	fmt.Println("\n2.3.4 数值型数据类型转换")
	n2 = int8(n3)
	fmt.Printf("n2=%v\n", n2)

	fmt.Println("\n2.3.5 字符型")
	var (
		c1 byte = 78 //byte, uint8的别名, 通常表示一个ASCII码
		c2      = '0'
		c3 rune = 23454 //rune, int32的别名 通常表示一个UTF-8(ASCII的超集)码
	)
	fmt.Printf("c1的码值=%v,这个码值对应的字符是%c,type is %T\n", c1, c1, c1)
	fmt.Printf("c2的码值=%v,这个码值对应的字符是%c,type is %T\n", c2, c2, c2)
	fmt.Printf("c3的码值=%v,这个码值对应的字符是%c,type is %T\n", c3, c3, c3)
	c4 := 'A' - 'a'
	c5 := 'x'
	c6 := c5 + c4
	fmt.Printf("c6的码值=%v,这个码值对应的字符是%c,type is %T\n", c6, c6, c6)

	fmt.Println("\n2.3.6 布尔型")
	var bool1 bool = true //bool
	fmt.Printf("bool1=%v,type is %T\n", bool1, bool1)

	fmt.Println("\n2.3.7 字符串")
	var s1 = "hello" //string
	fmt.Println(s1, "world")
	fmt.Println(len(s1))
	s2 := `	var (
		c1 byte = 78
		c2      = '0'
		c3 rune = 23454
	)`
	fmt.Println(s2)
	s3 := "你好世界"
	fmt.Println("len([]rune(s3))=", len([]rune(s3)))
	s4 := s3[3:6]
	fmt.Println("s4=", s4)
}

//2.4 指针
func Pointer() {
	// 取址符：&(获取当前变量的地址)
	// 取值符：*(访问地址指向的值)
	// 数据类型：*指向的类型
	var increase = func(n *int) {
		*n++ //n = n + 1
		fmt.Printf("\nincrease结束时n=%v\nn的内存地址为%v\nn指向的值为%v\n", n, &n, *n)
	}
	var src = 2022
	increase(&src)
	fmt.Printf("\n调用increase(ptr)之后，src=%v\nsrc的内存地址为%v\n", src, &src)
	var ptr = new(int)
	fmt.Printf("\nptr=%v\nptr的内存地址为%v\nptr指向的值为%v\n", ptr, &ptr, *ptr)
	//注意：引用类型的默认值为nil(空)，需要分配内存空间（引用已有值类型，或通过内建函数new()/make()来分配）
}

//2.5 fmt格式字符
func FmtVerbs() {
	fmt.Println("\n2.5.1 通用")
	fmt.Printf("%%\n") //%%代表%
	//%v代表value, 值
	//%T代表Type, 数据类型

	fmt.Println("\n2.5.2 整数")
	i := 123
	//%d代表decimal, 十进制
	//%b代表binary, 二进制(没有前缀)
	//%o代表octal, 八进制(没有前缀)
	//%x代表hexadecimal, 十六进制a-f(没有前缀)
	//%X代表hexadecimal, 十六进制A-F(没有前缀)
	fmt.Printf("%U\n", i) //%U代表Unicode, U+四位16进制int32
	fmt.Printf("%c\n", i) //%c代表character, Unicode码值所对应的字符
	fmt.Printf("%q\n", i) //%q代表quoted, 带单引号的Unicode码值所对应的字符

	fmt.Println("\n2.5.3 浮点数")
	f := 123.456
	fmt.Printf("%f\n", f)   //%f或%F代表float, 小数
	fmt.Printf("%.2f\n", f) //%.2f代表保留2位小数的%f(%.f为保留0位)
	fmt.Printf("%20f\n", f) //%5f代表最小宽度为5的%f
	//可以结合使用, 如%5.2f
	fmt.Printf("%b\n", f) //%b表示指数为2的幂的无小数科学记数法
	//%e表示使用小写e的科学计数法
	fmt.Printf("%E\n", f) //%E表示使用大写E的科学计数法
	//%g表示自动对宽度较大的数采用%e
	//%G表示自动对宽度较大的数采用%E
	fmt.Printf("%X\n", f) //%x表示hexadecimal, 0x十六进制科学计数法
	//%X表示hexadecimal, 0X十六进制科学计数法

	fmt.Println("\n2.5.4 布尔")
	fmt.Printf("%t\n", f == 123.456) //%t表示true or false, true或false的单词

	fmt.Println("\n2.5.5 字符串或byte切片")
	//%s表示string, 按字符串输出
	s := "hello world"
	fmt.Printf("%q\n", s) //%q表示quoted, 带双引号的按字符串输出
	fmt.Printf("%x\n", s) //%x表示hexadecimal, 每个byte按两位小写十六进制码值输出
	//%X表示hexadecimal, 每个byte按两位大写十六进制码值输出

	fmt.Println("\n2.5.6 指针")
	p := &s
	fmt.Printf("%p\n", p) //%p表示pointer, 0x开头的16进制地址
	//所有适用于整数的格式字符也适用于指针
}

//2.6 运算符
func Operator() {
	fmt.Println("\n2.6.1 算数运算符")
	//+，-，*，/，%
	fmt.Printf("8%%3=%d\n", 8%3)
	i := 123
	//++， --
	i++ //i=i+1
	fmt.Printf("i=%d\n", i)

	fmt.Println("\n2.6.2 位运算符")
	var b uint8 = 0b00111100
	fmt.Printf("b>>2=%b\n", b>>2)
	fmt.Printf("b<<2=%b\n", b<<2)
	var b1 uint8 = 0b00111100
	var b2 uint8 = 0b11001111
	fmt.Printf("b1&b2=%b\n", b1&b2) //按位与,都为1则为1
	fmt.Printf("b1|b2=%b\n", b1|b2) //按位或,一个为1则为1
	fmt.Printf("b1^b2=%b\n", b1^b2) //按位异或, 不同则为1

	fmt.Println("\n2.6.3 赋值运算符")
	//=，+=，-=，*=，/=，%=，>>=，<<=，&=，|=，^=
	b += 3 //b=b+3
	fmt.Printf("b=%d\n", b)

	fmt.Println("\n2.6.4 关系运算符")
	//>，>=，<，<=，==，!=
	fmt.Printf("b1==b2?%t\n", b1 == b2)

	fmt.Println("\n2.6.5 逻辑运算符")
	//&&，||，!
	bool1 := true
	bool2 := false
	fmt.Printf("bool1&&bool2?%t\n", bool1 && bool2) //与
	fmt.Printf("bool1||bool2?%t\n", bool1 || bool2) //或
	fmt.Printf("!bool2?%t\n", !bool2)               //非
}

//3.1 if…else
func IfElse() {
	var age uint8
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age)
	if age < 13 {
		fmt.Println("小朋友不要学编程哦")
	} else if age < 25 {
		fmt.Println("大朋友不要学编程哦")
	} else {
		fmt.Println("老朋友不要学编程哦")
	}

	fmt.Println("3.1.2 if 简短语句")
	if i := 3; i > 0 {
		fmt.Println("i=", i)
	} else if i > 3 {
		fmt.Println("i=", i)
	}
}

//3.2 switch…case
func SwitchCase() {
	var weekday uint8
	fmt.Println("请输入星期（数字）")
	fmt.Scanln(&weekday)
	switch weekday {
	case 1: //case结尾会自动break，如果需要继续匹配下一项可以加入fallthrough
		fmt.Println("酱油炒饭")
	case 2:
		fmt.Println("酱油炒面")
	default: //default可以省略
		fmt.Println("输入有误")
	}
}

//3.3 for循环
func For() {
	//break//结束
	//continue//结束本次继续下一次

	fmt.Println("\n3.3.1 无限循环")
	i := 1
	for {
		fmt.Print(i, "\t")
		i++
		if i == 11 {
			fmt.Println()
			break
		}
	}

	fmt.Println("\n3.3.2 条件循环")
	i = 1
	for i < 11 {
		fmt.Print(i, "\t")
		i++
	}
	fmt.Println()

	fmt.Println("\n3.3.3 标准for循环")
	for j := 1; j < 11; j++ {
		fmt.Print(j, "\t")
	}
	fmt.Println()
}

//3.4 label与goto
func LabelAndGoto() {
	fmt.Println("\n 3.4.1 label")
outside:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print("+ ")
			if i == 9 && j == 4 {
				break outside
			}
		}
		fmt.Println()
	}
	fmt.Println("\n 3.4.2 goto")
	fmt.Print("1 ")
	if i := 1; i != 1 {
		goto four //不推荐
	}
	fmt.Print("2 ")
	fmt.Print("3 ")
four:
	fmt.Print("4 ")
	fmt.Print("5 ")
}

//3.5 函数
func Function() {
	res1, res2 := func(n1, n2 int) (sum, difference int) {
		sum = n1 + n2
		difference = n1 - n2
		return
	}(2, 3)
	fmt.Println("res1=", res1, ", res2=", res2)
	//fmt.Printf("getRes=%v,Type of getRes=%T\n", getRes, getRes)
}

//3.6 defer
func deferUtil() func(int) int {
	i := 0
	return func(n int) int { //n是形参：形式参数，定义函数时使用的参数
		fmt.Printf("本次调用接收到n=%v\n", n)
		i++
		fmt.Printf("匿名工具函数被第%v次调用\n", i)
		return i
	}
}
func Defer() int {
	f := deferUtil()
	defer f(1) //1是实参：调用时传递给函数的实际参数
	defer f(2)
	defer f(3)
	return f(4)
}
func DeferRecover() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	n := 0
	fmt.Println(3 / n)
}

//4.1 数组
func Array() {
	//4.1.1 声明
	var a [3]int = [...]int{ //长度是数组类型的一部分, 长度不能留空, 留空是切片类型//等号右侧的长度可以简写为[…]自动判断
		1,
		456,
		789,
	}
	a[0] = 123
	fmt.Println("for 遍历")
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%v]=%v\n", i, a[i])
	}
	fmt.Println("\n4.1.2 for…range 遍历")
	for i, v := range a {
		fmt.Printf("a[%v]=%v\n", i, v)
	}
	fmt.Println("\n4.1.3 多维数组")
	var twoDimensionalArray [3][4]int = [3][4]int{
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{3, 4, 5, 6},
	}
	for i, v := range twoDimensionalArray {
		for i2, v2 := range v {
			fmt.Printf("a[%v][%v]=%v\t", i, i2, v2)
		}
		fmt.Println()
	}
}

//4.2 切片
func Slice() {
	//切片是对数组的引用
	//切片本身并不存储任何数据，它只是描述了底层数组中的一段
	array := [5]int{1, 2, 3, 4, 5}
	var s1 []int = array[1:4] //[开始引用的index:结束引用的index+1)//[0:len(array)]等效于[:]
	s1[0] = 0
	fmt.Println("array=", array)
	s2 := s1[1:]
	s2[0] = 0
	fmt.Println("array=", array)
	var s3 []int
	fmt.Println("does s3==nil?", s3 == nil)
	s3 = make([]int, 3) //make([]Type, len, cap)//如果不写cap,默认与len相同
	fmt.Printf("len(s3)=%v,cap(s3)=%v", len(s3), cap(s3))
	s4 := []int{1, 2, 3} //由系统自动创建底层数组
	fmt.Println("s4=", s4)

	fmt.Println("\n4.2.3 追加元素")
	s1 = append(s1, 6, 7, 8) //底层创建了新的数组，不再引用原数组
	s1[4] = 0
	fmt.Println("array=", array)
	fmt.Println("s1=", s1)
	s5 := append(s1, s2...)
	fmt.Println("s5=", s5)

	fmt.Println("\n4.2.4 复制数组")
	s6 := []int{1, 1}
	copy(s5, s6) //容量能接收多少，就接收多少
	fmt.Println("s5", s5)

	fmt.Println("\n4.2.5 string与[]byte")
	str := "hello 世界"
	fmt.Printf("[]byte(str)=%v\n[]byte(str)=%s\n", []byte(str), []byte(str))
	for i, v := range str {
		fmt.Printf("str[%d]=%c\n", i, v)
	}
	key := util.SelectByKey("注册", "登录", "退出")
	fmt.Println("接收到key=", key)
}

//4.3 map
func Map() {
	var m1 map[string]string
	fmt.Println("m1 == nil ?", m1 == nil)
	m1 = make(map[string]string) //make(Type, 初始size)//初始size可省略，默认为1
	m1["早上"] = "敲代码"
	m1["中午"] = "送外卖"
	m1["晚上"] = "开滴滴"
	fmt.Println("m1 =", m1)
	m2 := map[string]string{
		"下午": "改bug",
		"凌晨": "卖早餐",
	}
	fmt.Println("m2 =", m2)
	v, ok := m2["早上"]
	if ok {
		fmt.Println("v =", v)
	} else {
		fmt.Println("key不存在")
	}
	delete(m1, "晚上")
	fmt.Println("m1 =", m1)
	//m1 = nil
	m2 = make(map[string]string)
	fmt.Println("m2 =", m2)
	for key, value := range m1 {
		fmt.Printf("m1[%v]=%v\n", key, value)
	}
}

//4.4 自定义数据类型&类型别名
func TypeDefintionAndTypeAlias() {
	fmt.Println("\n4.4.1 自定义数据")
	type mesType uint16
	var u1000 uint16 = 1000
	var textMes mesType = mesType(u1000)
	fmt.Printf("textMes=%v, Type of textMes=%T\n", textMes, textMes)

	fmt.Println("\n4.4.2 类型别名")
	type myUint16 = uint16
	var myu16 myUint16 = u1000
	fmt.Printf("myu16=%v, Type of myu16=%T\n", myu16, myu16)
}

//4.5 结构体
type User struct {
	Name string `json:"name"`
	Id   uint32
}
type Account struct {
	User
	password string
}
type Contact struct {
	*User
	Remark string
}

func Struct() {
	var u1 User = User{
		Name: "张三",
	}
	u1.Id = 10000
	var u2 *User = &User{
		Name: "李四",
	}
	u2.Id = 10001 //(*u2).Id=10001
	var a1 = Account{
		User: User{
			Name: u1.Name,
		},
		password: "666",
	}
	var c1 *Contact = &Contact{
		User: &User{
			Id: u2.Id,
		},
		Remark: "王麻子",
	}
	c1.Name = "王五" //c1.User.Name="王五"//没有重复字段时可以简写
	fmt.Println("a1=", a1)
	fmt.Println("c1=", c1)
	fmt.Println("c1.User=", *((*c1).User))
}

//5.1 方法
func (u User) printName() {
	fmt.Println("u.Name=", u.Name)
}
func (u *User) setId() {
	(*u).Id = 10000 //注意“.”优先级高于“&”/“*”, 使用时可以简写(隐式间接引用)
}
func Method() {
	u := &User{ //可以使用“&”前缀快速声明结构体指针
		Name: "小方块",
	}
	u.printName()
	u.setId()
	fmt.Println("u=", u)
}

//5.2 接口
type textMes struct {
	Type string
	Text string
}

func (tm *textMes) setText() {
	tm.Text = "hello"
}

type imgMes struct {
	Type string
	Img  string
}

func (im *imgMes) setImg() {
	im.Img = "清明上河图"
}

type Mes interface {
	setType()
}

func (tm *textMes) setType() {
	tm.Type = "文字消息"
}
func (im *imgMes) setType() {
	im.Type = "图片消息"
}
func SendMes(m Mes) {
	m.setType()
	switch mptr := m.(type) {
	case *textMes:
		mptr.setText()
	case *imgMes:
		mptr.setImg()
	}
	fmt.Println("m=", m)
}
func Interface() {
	tm := textMes{}
	SendMes(&tm)
	im := imgMes{}
	SendMes(&im)
	var n1 int = 1
	n1interface := interface{}(n1)
	n2, ok := n1interface.(int)
	if ok {
		fmt.Println("n2=", n2)
	} else {
		fmt.Println("类型断言失败")
	}
}

//5.3 协程
var (
	c    int
	lock sync.Mutex
)

func PrimeNum(n int) {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return
		}
	}
	fmt.Printf("%v\t", n)
	lock.Lock()
	c++
	lock.Unlock()
}
func Goroutine() {
	for i := 2; i < 100001; i++ {
		go PrimeNum(i)
	}
	var key string
	fmt.Scanln(&key)
	fmt.Printf("\n共找到%v个素数\n", c)
}

//5.4 channel
func pushNum(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
func pushPrimeNum(n int, c chan int) {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return
		}
	}
	c <- n
}
func Channel() {
	var c1 chan int = make(chan int)
	go pushNum(c1)
	// for {
	// 	v, ok := <-c1
	// 	if ok {
	// 		fmt.Printf("%v\t", v)
	// 	} else {
	// 		break
	// 	}
	// }
	for v := range c1 {
		fmt.Printf("%v\t", v)
	}
	var c2 chan int = make(chan int, 100)
	for i := 2; i < 100001; i++ {
		go pushPrimeNum(i, c2)
	}
Print:
	for {
		select {
		case v := <-c2:
			fmt.Printf("%v\t", v)
		default:
			fmt.Println("所有素数已经找到")
			break Print
		}
	}
}
