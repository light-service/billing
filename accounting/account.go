package accounting

import "time"

// https://blog.csdn.net/qq1115787476/article/details/104262013

// 账户表
type Account struct {
	ID         int
	AccountKey string // 账号ID
	Type       string // 普通个人，普通企业，内部个人，内部企业
	TypeRemark string // 类型备注
	UserKey    string // 用户ID
	Balance    int    // 余额
	CreditLine int    //授信额度
	LockFund   int    // 冻结额度
	Status     int    // 风控冻结、系统冻结
	Direct     int    // 账户方向
	Version    int    // 版本号
	UpdatedAt  time.Time
	CreatedAt  time.Time
	DeletedAt  *time.Time
}

// 账务流水表
type AccountFlow struct {
	ID                 int
	TransactionID      string // 交易ID，业务系统传入
	AccountKey         int    // 记账方
	OpponentAccountKey string // 对手方
	Amount             int    // 交易金额
	InitAmount         int    // 交易初始值
	ResultAmount       int    // 交易结果值
	Direct             int    // 划款方向
	ReverseStatus      int    // 反记账方向
	Version            int    // 版本号
	UpdatedAt          time.Time
	CreatedAt          time.Time
	DeletedAt          *time.Time
}

// 冻结流水表
type AccountLockFlow struct {
	ID            int
	TransactionID string // 交易ID，业务系统传入
	AccountKey    int    // 记账方
	Amount        int    // 交易金额
	InitAmount    int    // 交易初始值
	ResultAmount  int    // 交易结果值
	Version       int    // 版本号
	UpdatedAt     time.Time
	CreatedAt     time.Time
	DeletedAt     *time.Time
}

// 账户变更记录
type AccountChangeFlow struct {
	ID           int
	AccountKey   int // 记账方
	InitAmount   int // 交易初始值
	ResultAmount int // 交易结果值
	Remark       string
	Version      int // 版本号
	UpdatedAt    time.Time
	CreatedAt    time.Time
	DeletedAt    *time.Time
}
