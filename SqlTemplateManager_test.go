package sqlte

import (
	"embed"
	_ "embed"
	"fmt"
	"testing"
)

//go:embed sqlTemplate
var res embed.FS

func TestSqlTemplate(t *testing.T) {
	TemplateManager.Bootstrap(res)
	tpl := TemplateManager.GetTemplate("demo")
	sql, err := tpl.MakeSql("queryRows", map[string]interface{}{
		"buyer_id": "123456",
		"status": 2,
		"start_date": 332323,
		"end_date": 33333,
		"page": map[string]interface{}{
			"Offset": 1,
			"PageSize": 30,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("sql:", sql)
}


