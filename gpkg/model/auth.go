package model

type Auth struct {
	ID           int32       `json:"id"`
	UserId       string      `json:"user_id"`
	RefreshToken string      `json:"refresh_token"`
	Grade        string      `json:"grade"`
	BanEndAt     interface{} `json:"ban_end_at"`
	CreateAt     int64       `json:"create_at"`
	UpdateAt     int64       `json:"update_at"`
}

func (m *Auth) GetTable() string {
	return "auth"
}

func (m *Auth) GetCreate() []interface{} {
	return []interface{}{
		m.ID, m.Grade,
	}
}
