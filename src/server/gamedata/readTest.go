package gamedata

// 确保 bin/gamedata 目录中存在 Test.txt 文件
// 文件名必须和此结构体名称相同（大小写敏感）
// 结构体的一个实例映射 recordfile 中的一行
type Test struct {
	// 将第一列按 int 类型解析
	// "index" 表明在此列上建立唯一索引
	Id int "index"
	// 将第二列解析为长度为 4 的整型数组
	Arr [4]int
	// 将第三列解析为字符串
	Str string
}

// 读取 recordfile Test.txt 到内存中
// RfTest 即为 Test.txt 的内存镜像
var RfTest = readRf(Test{})

func init() {

}
