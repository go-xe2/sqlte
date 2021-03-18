package sqlte

import (
	"embed"
	"github.com/go-xe2/x/os/xlog"
	"testing"

	//"github.com/hashicorp/hcl"
)

//go:embed sqlTemplate
var res embed.FS

type DboSqlMap struct {
	Select map[string]string `json:"select"`
	Execute map[string]string `json:"execute"`
}

func TestHclParse(t *testing.T)  {
	Template.Bootstrap(res)
	xlog.Info(Template.fileTemplates)
	//dir, err := res.ReadDir("sqls")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//for _, d := range dir {
	//	info, err := d.Info()
	//	fmt.Println( ", err:", err)
	//	fmt.Println("name:", info.Name())
	//	fmt.Println("isDir", info.IsDir())
	//	fmt.Println("isDir", info.Mode())
	//}
	//bits, err := res.ReadFile("sqls/demo.hcl")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fmt.Println("str:", string(bits))
	//var data map[string]*DboSqlMap
	//if err := hcl.Unmarshal(bits, &data); err != nil {
	//	t.Fatal(err)
	//}
	//xlog.Info("data:", data)
	//fmt.Println("===================")
	//xlog.Info("data.getBuyerOrderPage:", data["getBuyerOrderPage"], "====> type:", reflect.TypeOf(data["getBuyerOrderPage"]))
	//fmt.Println("========================")
	//fmt.Println("data.Select:", data["getBuyerOrderPage"].Select)
	//fmt.Println("========================")
	//fmt.Println("data.Execute:", data["getBuyerOrderPage"].Execute)
	////getBuyerOrderPage := data["getBuyerOrderPage"].([]map[string]interface{})
	////queryRows := getBuyerOrderPage[0]["queryRows"].(string)
	////fmt.Println("==>queryRows:", queryRows)
	////var envParams = map[string]interface{}{
	////	"status": 3,
	////	"start_date": 3232323,
	////	"buyer_id": "123456",
	////	"page": map[string]interface{}{
	////		"Offset": 1,
	////		"PageSize": 20,
	////	},
	////}
	////tpl, err := template.New("sql").Parse(queryRows)
	////if err != nil {
	////	t.Fatal(err)
	////}
	////bufs := bytes.Buffer{}
	////tpl.Execute(&bufs, envParams)
	////fmt.Println("sql:", bufs.String())
}
