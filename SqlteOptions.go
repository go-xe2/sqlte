package sqlte

type SqlteOptions struct {
	TemplateDirName string `json:"template_dir_name"` // sql模块目录名称
	TemplateExt string `json:"template_ext"`
}

var DefaultSqlteOptions = SqlteOptions{
	TemplateDirName: "sqlTemplate",
	TemplateExt: ".hcl", // 模块文件名后缀
}
