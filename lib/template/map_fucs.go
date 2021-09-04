package template

import (
	"html/template"

	"github.com/jeffprestes/test-bbolt/lib/contx"
)

// FuncMaps to view
func FuncMaps() []template.FuncMap {
	return []template.FuncMap{
		map[string]interface{}{
			"Tr": contx.I18n,
		}}
}
