package utils

import (
	"fmt"
	"hackaton/log"
	"html/template"
	"path/filepath"
)

var TemplateCache map[string]*template.Template

func NewTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl.html"))
	if err != nil {
		log.ErrorLogger.Println(err)
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println(name)
		ts, err := template.ParseFiles(page)
		if err != nil {
			log.ErrorLogger.Println(err)
			return nil, err
		}
		//ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl.html"))
		//if err != nil {
		//	return nil, err
		//}
		//
		//ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl.html"))
		//if err != nil {
		//	return nil, err
		//}
		cache[name] = ts
	}
	return cache, nil
}
