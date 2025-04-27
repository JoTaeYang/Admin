package model

import (
	"github.com/JoTaeYang/Admin/gpkg/pt"
)

type SliceCurrency []*Currency

type Currency struct {
	UserId       string `json:"user_id"`
	CurrencyType int64  `json:"currency_type"`
	Count        int64  `json:"count"`
	CreateAt     int64  `json:"create_at"`
	UpdateAt     int64  `json:"update_at"`
}

func (m *Currency) GetKey() string {
	return "currency"
}

func (m *Currency) GetEModel() EModel {
	return ECurrency
}

func (m *Currency) GetCreate() []interface{} {
	return []interface{}{
		m.UserId, m.CurrencyType, m.Count,
	}
}

func (m *Currency) ConvertGRPC() *pt.DataItem {
	return &pt.DataItem{
		Item: &pt.DataItem_Currency{
			Currency: &pt.Currency{
				UserId:       m.UserId,
				CurrencyType: m.CurrencyType,
				Count:        m.Count,
			},
		},
	}
}
