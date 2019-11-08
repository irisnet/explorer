package vo

type ResultVo struct {
	Block    SimpleBlockVo    `json:"block,omitempty"`
	Proposal SimpleProposalVo `json:"proposal,omitempty"`
}

type SimpleBlockVo struct {
	Height    int64  `json:"height,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Hash      string `json:"hash,omitempty"`
}

type SimpleProposalVo struct {
	ProposalId uint64 `json:"proposal_id,omitempty"`
	Title      string `json:"title,omitempty"`
	Type       string `json:"type,omitempty"`
	Status     string `json:"status,omitempty"`
	SubmitTime string `json:"submit_time,omitempty"`
}

type EnvConfig struct {
	CurEnv  string     `json:"cur_env"`
	ChainId string     `json:"chain_id"`
	Configs []ConfigVo `json:"configs"`
}

type ConfigVo struct {
	EnvNm      string `json:"env_nm,omitempty"`
	Env        string `json:"env,omitempty"`
	Host       string `json:"host,omitempty"`
	ChainId    string `json:"chain_id,omitempty"`
	ShowFaucet int    `json:"show_faucet,omitempty"`
}