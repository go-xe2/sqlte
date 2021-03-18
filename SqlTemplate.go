package sqlte

type SqlTemplate struct {
	DboName string `json:"dbo_name"`
	Select map[string]string `json:"select"`
	Execute map[string]string `json:"execute"`
	Manager *SqlTemplateManager
}

func (it SqlTemplate) New(manager *SqlTemplateManager) *SqlTemplate {
	it.Select = make(map[string]string)
	it.Execute = make(map[string]string)
	it.Manager = manager
	return &it
}
