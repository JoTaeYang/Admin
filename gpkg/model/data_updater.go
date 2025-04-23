package model

/*
Selector 에 전달한 key 데이터와
LoadTx로 나온 results 결과를 어떻게 잘 저장한 후에
updater에 들어온 데이터가 results에 있다? => update
없다? => create
이런식으로 하면 어떨까 고민이 된다.
*/
type Updater struct {
	updater map[EModel]interface{}
	creater map[EModel]interface{}
	deleter map[EModel]interface{}
}

func (u *Updater) AddUpdate(key EModel, data interface{}) {
	u.updater[key] = data
}

func (u *Updater) AddPut(key EModel, data interface{}) {
	u.creater[key] = data
}

func (u *Updater) AddDelete(key EModel, data interface{}) {
	u.deleter[key] = data
}

func (*Updater) Execute() {

}

/*
History
Prev, After 관리
기존 DataContext 어떻게 좀 하면,,
*/
