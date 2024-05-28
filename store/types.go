package store

type GlobalStateType struct {
	State []*StateType `json:"state"`
	size  int
	front int
	rear  int
}

type StateType struct {
	Status           bool   `json:"status"`
	Height           int64  `json:"height"`
	OracleMissCnt    uint64 `json:"oracle_miss_cnt"`
	BlockSign        bool   `json:"block_sign"`
	OracleSign       bool   `json:"oracle_sign"`
	OracleDoubleSign bool   `json:"oracle_double_sign"`
}
