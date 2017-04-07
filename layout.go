package ui

import (
	"html/template"
	"path"
	"github.com/gobly/core"
)

var funcMap = template.FuncMap{
	"App" : func() core.Application { return core.App },
}

func LoadSingle(mainTemplate string) *template.Template {
	mainTemplate = path.Join(core.CurrentModule(), mainTemplate)
	footerTemplate := path.Join(core.CurrentFolder(), "html", "footer.html")
	singleTemplate := path.Join(core.CurrentFolder(), "html", "single.html")

	return template.Must(template.New("single.html").Funcs(funcMap).ParseFiles(singleTemplate, mainTemplate, footerTemplate))
}

func LoadDouble(mainTemplate string, sidebarTemplate string) *template.Template {
	mainTemplate = path.Join(core.CurrentModule(), mainTemplate)
	sidebarTemplate = path.Join(core.CurrentModule(), sidebarTemplate)
	footerTemplate := path.Join(core.CurrentFolder(), "html", "footer.html")
	doubleTemplate := path.Join(core.CurrentFolder(), "html", "double.html")

	return template.Must(template.New("double.html").Funcs(funcMap).ParseFiles(doubleTemplate, mainTemplate, sidebarTemplate, footerTemplate))
}
