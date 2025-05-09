package changer

import (
	"github.com/JoTaeYang/Admin/gpkg/pt"
)

type Currency struct {
	*Changer
}

func newCurrency(changer *Changer) *Currency {
	return &Currency{
		Changer: changer,
	}
}

func (u *Currency) Use(t pt.Currency_T, cnt int64) bool {
	currency, ok := u.Changer.dataCtx.GetCurrency(t)
	if !ok {
		return false
	}

	// 재화 소모 체크
	if currency.Count < cnt {
		return false
	}

	// DB Update
	u.updater.AddUpsert(currency)

	// Meta Item Log 추가

	return true
}
