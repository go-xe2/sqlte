package sqlte

import (
	"bytes"
	"fmt"
)

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

// MakeSql 根据参数解析模板生成sql
func (it *SqlTemplate) MakeSql(tplName string, args map[string]interface{}) (string, error)  {
	buf := bytes.Buffer{}
	CompileName := fmt.Sprintf("%s.%s", it.DboName, tplName)
	if err := it.Manager.Compile.ExecuteTemplate(&buf, CompileName, args); err != nil {
		return "", err
	}
	return buf.String(), nil
}

