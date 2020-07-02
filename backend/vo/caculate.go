package vo

type ExStaticMonthDataRespond struct {
	Total int         `json:"total"`
	Type  string      `json:"type"`
	Data  interface{} `json:"data"`
}

type ExStaticDelegatorMonthVo struct {
	Date                   string  `json:"统计月份"`
	Address                string  `json:"委托地址"`
	Rewards                float64 `json:"期初待领取奖励"`
	TerminalRewards        float64 `json:"期末待领取奖励"`
	PeriodWithdrawRewards  float64 `json:"期间领取到账"`
	PeriodIncrementRewards float64 `json:"期间奖励净增量"`
	TerminalDelegation     float64 `json:"期末在托量"`
	IncrementDelegation    float64 `json:"期间委托净增量"`
	PeriodDelegationTimes  int     `json:"期间委托交易次数"`
	AnnualizedRate         string  `json:"年化"`
	CaculateDate           string  `json:"统计更新日期"`
	//CreateAt               int64   `json:"create_at"`
	//UpdateAt               int64   `json:"update_at"`
}

type ExStaticValidatorMonthVo struct {
	Date                    string  `json:"统计月份"`
	Address                 string  `json:"自委托地址"`
	CreateValidatorHeight   int64   `json:"初始创建块高"`
	OperatorAddress         string  `json:"验证节点运营地址"`
	Status                  string  `json:"状态"`
	TerminalCommission      float64 `json:"期末待领取佣金"`
	PeriodCommission        float64 `json:"期间领取到账"`
	IncrementCommission     float64 `json:"期间佣金净增量"`
	TerminalDelegation      float64 `json:"期末受托数量"`
	IncrementDelegation     float64 `json:"期间受托净增量"`
	Tokens                  float64 `json:"期末权重"` //权重排名用
	TerminalDelegatorN      int     `json:"期末受托地址数"`
	IncrementDelegatorN     int     `json:"期间受托地址净增量"`
	TerminalSelfBond        float64 `json:"期末自托数量"`
	IncrementSelfBond       float64 `json:"期间自托净增量"`
	CommissionRateMax       float64 `json:"期间最高佣金率"`
	CommissionRateMin       float64 `json:"期间最低佣金率"`
	FoundationDelegateT     float64 `json:"期末受基金会委托量"`
	FoundationDelegateIncre float64 `json:"期间受基金会委托净增量"`
	Moniker                 string  `json:"别名"`
	CaculateDate            string  `json:"统计更新日期"`
	//CreateAt                int64   `json:"create_at"`
	//UpdateAt                int64   `json:"update_at"`
}
