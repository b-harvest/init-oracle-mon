package store

type GlobalStateType struct {
	State []*StateType `json:"state"`
	size  int
	front int
	rear  int
}

type StateType struct {
	Status     bool  `json:"status"`
	Height     int64 `json:"height"`
	BlockSign  bool  `json:"block_sign"`
	OracleSign bool  `json:"oracle_sign"`
	DoubleSign bool  `json:"double_sign"`
}
