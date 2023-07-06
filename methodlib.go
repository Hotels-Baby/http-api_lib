package http_api_toolbox

import "strings"

type MethodLib struct {
	Methods []string
}

func (ml *MethodLib) AddMethod(method string) {
	ml.Methods = append(ml.Methods, method)
}

func (ml *MethodLib) IsMethodAllowed(method string) bool {
	for _, m := range ml.Methods {
		if strings.EqualFold(m, method) {
			return true
		}
	}

	return false
}
