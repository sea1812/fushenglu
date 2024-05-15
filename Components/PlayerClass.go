package Components

/*
	玩家类
*/

type TPlayer struct {
	Id string //ID，用邮箱地址代替

	Age        int64 //年龄，在游戏中存活的天数，用天数表示
	RemainDays int64 //距离游戏结束的剩余天数
	TotalDays  int64 //总生命天数，即玩家在游戏中能够存活的总天数

	Wealth        int64 //总财富
	Salary        int64 //每天固定收入
	SalaryFloat   int64 //每天固定收入浮动比例，百分比，使用时需要除100
	Expenses      int64 //每天固定支出
	ExpensesFloat int64 //每天固定支出浮动比例，百分比，使用时需要除100

	Health     int64 //总健康度
	HealthBack int64 //每天自动恢复健康度数值，相当于回血，如果是负值就是每天掉血

	BadLucks  int64 //累计的倒霉事件次数
	GoodLucks int64 //累计的幸运事件次数

	Happiness     int64 //快乐指数，遇到BadLuck降低，否则每天自动回血
	HappinessBack int64 //每天自动回血的快乐指数

	LuckyValue int64         //幸运值，幸运值越高，发生好事件的概率越高，以体现强者恒强，富者恒富的精神
	Milestone  []interface{} //发生过的逆天改命的重大里程碑事件 TODO 这里后面要改成结构体类型
}

/*
	TPlayer对象的方法：
	1、LoadPlayer(AId string) 按ID查询玩家数据，填充当前对象属性
	2、SavePlayer()	保存玩家数据到数据库
	3、InitPlayer() 按照初始化规则初始化玩家数据
	4、Die(AFlag interface) Game Over，结束该账号，参数是结束原因的结构
*/
