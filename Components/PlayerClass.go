package Components

import "github.com/gogf/gf/encoding/gjson"

/*
	玩家类
*/

type TPlayer struct {
	Storage TStorage //存储对象

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

	LuckyValue int64 //幸运值，幸运值越高，发生好事件的概率越高，以体现强者恒强，富者恒富的精神
	Died       bool  //账号是否已终止
	//Milestone  []interface{} //发生过的逆天改命的重大里程碑事件 TODO 这里后面要改成结构体类型，为避免对象结构太复杂，改成独立对象，用ID关联
}

/*
	事件对象
*/

type TEvent struct {
	EventId             int64  //事件ID
	EventType           int    //事件类型，分为幸运、倒霉、一般三类
	EventDescription    string //事件描述
	EffectWealth        int64  //影响到Player的Wealth变动绝对值，可以是负数
	EffectSalary        int64  //影响到Player的每天固定收入，为百分数
	EffectSalaryFloat   int64  //影响到Player的每天固定收入的浮动变量
	EffectExpenses      int64  //影响到Player每天的固定支出，为百分数
	EffectExpensesFloat int64  //影响到Player每天的固定支出的浮动变量
	EffectHealth        int64  //影响到Player的健康度
	EffectHealthBack    int64  //影响到Player每日自动回血的数值
	EffectHappiness     int64  //影响到Player的快乐指数
	EffectHappinessBack int64  //影响到Player的每日自动恢复快乐指数
	EffectLuckyValue    int64  //影响到Player的幸运值
}

/*
	事件管理器对象
*/

type TEvents struct {
	TotalEventsCount    int64 //全部事件数量
	GoodLuckEventsCount int64 //幸运事件数量
	BadLuckEventsCount  int64 //倒霉事件数量
	NormalEventsCount   int64 //一般事件数量

}

/*
	TPlayer对象的方法：
	01、LoadPlayer(AId string) 按ID查询玩家数据，填充当前对象属性
	02、SavePlayer()	保存玩家数据到数据库，KEY即ID
	3、InitPlayer() 按照初始化规则初始化玩家数据
	4、Die(AFlag interface) Game Over，结束该账号，参数是结束原因的结构
	5、NextDay() 下一天，即接受新的事件，修改本身属性，并产生新的输出给接口，以便返回给前端（TODO SSE或者其他方式再议）
	6、GetNewEvent() 获取新的事件，返回事件对象结构
	7、UpdateByEvent(AEvent) 利用事件更新自身属性数据，参数为事件对象
	8、Summary() 在当前事件节点生成人生总结内容
	9、Json() 将当前属性输出为JSON
	10、初始化方法，主要是初始化Storage对象
*/

// Init 初始化方法，主要是初始化Storage对象
func (p *TPlayer) Init(AName string) {
	p.Storage = TStorage{}
	p.Storage.Init(AName)
}

// LoadPlayer 按ID查询玩家数据，填充当前对象属性
func (p *TPlayer) LoadPlayer(AId string) {
	err, res := p.Storage.Get(AId)
	if err != nil {
		p.Id = AId
		p.Age = res.GetInt64("Age")               //年龄，在游戏中存活的天数，用天数表示
		p.RemainDays = res.GetInt64("RemainDays") //距离游戏结束的剩余天数
		p.TotalDays = res.GetInt64("TotalDays")   //总生命天数，即玩家在游戏中能够存活的总天数

		p.Wealth = res.GetInt64("Wealth")               //总财富
		p.Salary = res.GetInt64("Salary")               //每天固定收入
		p.SalaryFloat = res.GetInt64("SalaryFloat")     //每天固定收入浮动比例，百分比，使用时需要除100
		p.Expenses = res.GetInt64("Expenses")           //每天固定支出
		p.ExpensesFloat = res.GetInt64("ExpensesFloat") //每天固定支出浮动比例，百分比，使用时需要除100

		p.Health = res.GetInt64("Health")         //总健康度
		p.HealthBack = res.GetInt64("HealthBack") //每天自动恢复健康度数值，相当于回血，如果是负值就是每天掉血

		p.BadLucks = res.GetInt64("BadLucks")   //累计的倒霉事件次数
		p.GoodLucks = res.GetInt64("GoodLucks") //累计的幸运事件次数

		p.Happiness = res.GetInt64("Happiness")         //快乐指数，遇到BadLuck降低，否则每天自动回血
		p.HappinessBack = res.GetInt64("HappinessBack") //每天自动回血的快乐指数

		p.LuckyValue = res.GetInt64("LuckyValue") //幸运值，幸运值越高，发生好事件的概率越高，以体现强者恒强，富者恒富的精神
		p.Died = res.GetBool("Died")              //账号是否已终止
	}
}

// SavePlayer 保存玩家数据到数据库，KEY即ID
func (p *TPlayer) SavePlayer() {
	mJson := gjson.New(p)
	_ = p.Storage.Set(p.Id, mJson)
}

// InitPlayer 按照初始化规则初始化玩家数据
func (p *TPlayer) InitPlayer() {
	p.Age = 0             //在游戏中已存活的天数为0
	p.TotalDays = 2920    //总生命天数，TODO 这里需要改成常量，暂设为2920=8*365就是8年
	p.RemainDays = 2920   //剩余天数
	p.Wealth = 10000      //初始财务为10000元，没有小数
	p.Salary = 10000 / 30 //以月工资1万元30天平均算
	p.SalaryFloat = 5     //月收入5%浮动
	p.Expenses = 100      //日支出
	p.ExpensesFloat = 10  //日支出浮动10%
	p.Health = 100        //初始健康度为100个点
	p.HealthBack = 10     //健康度每天自动回血10个点
	p.BadLucks = 0        //倒霉次数初始为0
	p.GoodLucks = 0       //幸运次数初始为0
	p.Happiness = 100     //初始快乐值为100
	p.HappinessBack = 10  //快乐值每天自动回血100
	p.LuckyValue = 50     //幸运度初始为50，满分100
	p.Died = false        //账号未终止
}

// GetNewEvent 获取新的事件，返回事件对象结构
func (p *TPlayer) GetNewEvent() TEvent {

}
