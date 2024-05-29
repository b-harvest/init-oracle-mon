package store

type GlobalStateType struct {
	Status *StatusType  `json:"status"`
	States []*StateType `json:"state"`
	size   int
	front  int
	rear   int
}

type StateType struct {
	Height           int64 `json:"height"`
	BlockSign        bool  `json:"block_sign"`
	OracleSign       bool  `json:"oracle_sign"`
	OracleDoubleSign bool  `json:"oracle_double_sign"`
}

type StatusType struct {
	Status        bool   `json:"status"`
	OracleMissed  string `json:"oracle_miss_cnt"`
	Uptime        string `json:"uptime"`
	WindowSize    uint64
	OracleMissCnt uint64
}
