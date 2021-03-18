package sqlte

import (
	"fmt"
	"github.com/go-xe2/x/os/xlog"
	"github.com/hashicorp/hcl"
	"path/filepath"
	"strings"
	"sync"
	"text/template"
)

type SqlTemplateManager struct {
	Complete *template.Template
	fileTemplates map[string]*SqlTemplate
	rw sync.RWMutex
}

var Template = SqlTemplateManager{}.New()

func (it SqlTemplateManager) New() SqlTemplateManager  {
	it.Complete = template.New("main")
	it.fileTemplates = make(map[string]*SqlTemplate)
	return it
}

func (it *SqlTemplateManager) GetTemplate(dboName string) *SqlTemplate {
	it.rw.Lock()
	defer it.rw.Unlock()
	if tpl, ok := it.fileTemplates[dboName]; ok {
		return tpl
	}
	// 加载sql
	return nil
}

func enumDirFiles(loader SqlLoader, fileList *[]string, path string) error {
	fmt.Println("enumDirFiles path:", path)
	dirs, err := loader.ReadDir(path)
	if err != nil {
		return fmt.Errorf("读取目录%s出错:%s", DefaultSqlteOptions.TemplateDirName, err)
	}
	for _, dir := range dirs {
		fmt.Println("====>>", dir.Name(), " is dir:", dir.IsDir())
		if dir.IsDir() {
			if dir.Name() == "." {
				continue
			}
			if err := enumDirFiles(loader, fileList, filepath.Join(path, dir.Name())); err != nil {
				return err
			}
			continue
		}
		if !strings.EqualFold(filepath.Ext(dir.Name()), DefaultSqlteOptions.TemplateExt) {
			continue
		}
		*fileList = append(*fileList, filepath.Join(path, dir.Name()))
	}
	return nil
}

func (it *SqlTemplateManager) loadSqlTemplateFiles(loader SqlLoader) {
	if loader == nil {
		panic(ErrSqlLoaderNotInit)
	}
	var fileList = make([]string, 0, 100)
	if err := enumDirFiles(loader, &fileList, DefaultSqlteOptions.TemplateDirName); err != nil {
		panic(err)
	}
	xlog.Info("fileList:", fileList)
	it.rw.Lock()
	defer it.rw.Unlock()
	for _, fileName := range fileList {
		fileData, err := loader.ReadFile(fileName)
		if err != nil {
			panic(fmt.Errorf("读取模板文件%s出错%s", fileName, err.Error()))
		}
		var tpl SqlTemplate
		if err := hcl.Unmarshal(fileData, &tpl); err != nil {
			panic(fmt.Errorf("文件模板%s格式错误:%s", fileName, err.Error()))
		}
		var dboName = tpl.DboName
		if dboName == "" {
			dboName = fileName
		}
		// 处理模板
		var existsTpl *SqlTemplate
		if tmp, ok := it.fileTemplates[dboName]; ok {
			existsTpl = tmp
		} else {
			existsTpl = SqlTemplate{}.New(it)
		}
		// 合并
		for k, v := range tpl.Select {
			existsTpl.Select[k] = v
			_, err := it.Complete.New(k).Parse(v)
			if err != nil {
				panic(err)
			}
		}
		for k, v := range tpl.Execute {
			existsTpl.Execute[k] = v
			_, err := it.Complete.New(k).Parse(v)
			if err != nil {
				panic(err)
			}
		}
		it.fileTemplates[dboName] = existsTpl
	}
}

func (it SqlTemplateManager) Bootstrap(loader SqlLoader) *SqlTemplateManager {
	it.loadSqlTemplateFiles(loader)
	return &it
}

