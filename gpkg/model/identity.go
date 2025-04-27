package model

type Identity struct {
	ID       int64  `json:"id"`
	UserId   string `json:"user_id"`
	ShardIdx int64  `json:"shard_idx"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

func (m *Identity) GetKey() string {
	return "identity"
}

func (m *Identity) GetEModel() EModel {
	return EIdentity
}

func (m *Identity) GetCreate() []interface{} {
	return []interface{}{
		m.ID, m.UserId, m.ShardIdx,
	}
}

/*
model 마다 어떤 DB를 읽어오는지

*/
