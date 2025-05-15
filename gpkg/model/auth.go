package model

type Auth struct {
	ID           int64       `json:"id"`
	UserId       string      `json:"user_id"`
	RefreshToken string      `json:"refresh_token"`
	Grade        string      `json:"grade"`
	ShardIdx     int64       `json:"shard_idx"`
	BanEndAt     interface{} `json:"ban_end_at"`
	CreateAt     int64       `json:"create_at"`
	UpdateAt     int64       `json:"update_at"`
}

func (m *Auth) GetKey() string {
	return "auth"
}

func (m *Auth) GetEModel() EModel {
	return EAuth
}

func (m *Auth) GetCreate() []interface{} {
	return []interface{}{
		m.UserId, m.Grade, m.ShardIdx,
	}
}

func (m *ModelHub) GetAuth() (*Auth, bool) {
	return GetFromContext[*Auth](m.DataCtx, EAuth)
}
