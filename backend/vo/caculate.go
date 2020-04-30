package vo

type ExStaticMonthDataRespond struct {
	Total int         `json:"total"`
	Type  string      `json:"type"`
	Data  interface{} `json:"data"`
}

type ExStaticDelegatorMonthVo struct {
	Address                string  `json:"address"`
	Date                   string  `json:"date"`
	CaculateDate           string  `json:"caculate_date"`
	TerminalRewards        float64 `json:"terminal_rewards"`
	PeriodWithdrawRewards  float64 `json:"period_withdraw_rewards"`
	PeriodIncrementRewards float64 `json:"period_increment_rewards"`
	TerminalDelegation     float64 `json:"terminal_delegation"`
	IncrementDelegation    float64 `json:"increment_delegation"`
	PeriodDelegationTimes  int     `json:"period_delegation_times"`
	//CreateAt               int64   `json:"create_at"`
	//UpdateAt               int64   `json:"update_at"`
}

type ExStaticValidatorMonthVo struct {
	Address                 string  `json:"address"`
	OperatorAddress         string  `json:"operator_address"`
	CreateValidatorHeight   int64   `json:"create_validator_height"`
	Date                    string  `json:"date"`
	CaculateDate            string  `json:"caculate_date"`
	TerminalCommission      float64 `json:"terminal_commission"`
	PeriodCommission        float64 `json:"period_commission"`
	IncrementCommission     float64 `json:"increment_commission"`
	TerminalDelegation      float64 `json:"terminal_delegation"`
	IncrementDelegation     float64 `json:"increment_delegation"`
	Tokens                  float64 `json:"tokens"` //权重排名用
	TerminalDelegatorN      int     `json:"terminal_delegator_n"`
	IncrementDelegatorN     int     `json:"increment_delegator_n"`
	TerminalSelfBond        float64 `json:"terminal_self_bond"`
	IncrementSelfBond       float64 `json:"increment_self_bond"`
	CommissionRateMax       float64 `json:"commission_rate_max"`
	CommissionRateMin       float64 `json:"commission_rate_min"`
	FoundationDelegateT     float64 `json:"foundation_delegate_t"`
	FoundationDelegateIncre float64 `json:"foundation_delegate_incre"`
	Moniker                 string  `json:"moniker"`
	//CreateAt                int64   `json:"create_at"`
	//UpdateAt                int64   `json:"update_at"`
}
