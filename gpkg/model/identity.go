package model

type Identity struct {
	ID       int32  `json:"id"`
	UserId   string `json:"user_id"`
	ShardIdx int    `json:"shard_idx"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

func (m *Identity) GetTable() string {
	return "identity"
}

func (m *Identity) GetCreate() []interface{} {
	return []interface{}{
		m.ID, m.UserId, m.ShardIdx,
	}
}
