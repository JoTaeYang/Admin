package model

type Manager struct {
	ID       string      `json:"id"`
	Grade    string      `json:"grade"`
	Password string      `json:"password"`
	Name     string      `json:"name"`
	CreateAt string      `json:"create_at"`
	UpdateAt string      `json:"update_at"`
	Ttl      interface{} `json:"ttl"`
}

func (m *Manager) GetTable() string {
	return "manager"
}

func (m *Manager) GetCreate() []interface{} {
	return []interface{}{
		m.ID, m.Grade, m.Name, m.Password,
	}
}
