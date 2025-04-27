package model

import "github.com/JoTaeYang/Admin/gpkg/pt"

type Profile struct {
	UserId          string `json:"user_id"`
	Name            string `json:"name"`
	NameChangeAt    string `json:"name_change_at"`
	NameChangeCount int64  `json:"name_change_count"`
	CreateAt        string `json:"create_at"`
	UpdateAt        string `json:"update_at"`
}

func (m *Profile) GetKey() string {
	return "profile"
}

func (m *Profile) GetEModel() EModel {
	return EProfile
}

func (m *Profile) GetCreate() []interface{} {
	return []interface{}{
		m.UserId, m.Name,
	}
}

func (m *Profile) ConvertGRPC() *pt.DataItem {
	return &pt.DataItem{
		Item: &pt.DataItem_Profile{
			Profile: &pt.Profile{
				UserId:          m.UserId,
				Name:            m.Name,
				NameChangeAt:    0,
				NameChangeCount: m.NameChangeCount,
			},
		},
	}
}
