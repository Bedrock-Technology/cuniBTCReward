package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strconv"
	"strings"
)

type goField struct {
	jsonName string
	format   string
	embedded string
}

func formatsFromGoFile(path string) (map[string]map[string]string, error) {
	parsed, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
	if err != nil {
		return nil, fmt.Errorf("parse Go types from %s: %w", path, err)
	}

	structs := map[string][]goField{}
	for _, declaration := range parsed.Decls {
		general, ok := declaration.(*ast.GenDecl)
		if !ok || general.Tok != token.TYPE {
			continue
		}
		for _, specification := range general.Specs {
			typeSpec, ok := specification.(*ast.TypeSpec)
			if !ok {
				continue
			}
			structure, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}
			for _, field := range structure.Fields.List {
				if len(field.Names) == 0 {
					if name := namedType(field.Type); name != "" {
						structs[typeSpec.Name.Name] = append(structs[typeSpec.Name.Name], goField{embedded: name})
					}
					continue
				}
				jsonName := jsonFieldName(field)
				if jsonName == "" || jsonName == "-" {
					continue
				}
				if format := numericFormat(field.Type); format != "" {
					structs[typeSpec.Name.Name] = append(structs[typeSpec.Name.Name], goField{jsonName: jsonName, format: format})
				}
			}
		}
	}

	result := make(map[string]map[string]string, len(structs))
	for name := range structs {
		result[name] = collectFormats(name, structs, map[string]bool{})
	}
	return result, nil
}

func collectFormats(name string, structs map[string][]goField, visiting map[string]bool) map[string]string {
	result := map[string]string{}
	if visiting[name] {
		return result
	}
	visiting[name] = true
	defer delete(visiting, name)

	for _, field := range structs[name] {
		if field.embedded != "" {
			for jsonName, format := range collectFormats(field.embedded, structs, visiting) {
				result[jsonName] = format
			}
			continue
		}
		result[field.jsonName] = field.format
	}
	return result
}

func applyFormats(document map[string]any, formats map[string]map[string]string) {
	components, _ := document["components"].(map[string]any)
	schemas, _ := components["schemas"].(map[string]any)
	for schemaName, fields := range formats {
		schema, _ := schemas[schemaName].(map[string]any)
		properties, _ := schema["properties"].(map[string]any)
		for propertyName, format := range fields {
			property, _ := properties[propertyName].(map[string]any)
			if property != nil {
				property["format"] = format
			}
		}
	}
}

func jsonFieldName(field *ast.Field) string {
	if field.Tag == nil {
		return ""
	}
	tag, err := strconv.Unquote(field.Tag.Value)
	if err != nil {
		return ""
	}
	return strings.Split(reflect.StructTag(tag).Get("json"), ",")[0]
}

func namedType(expression ast.Expr) string {
	switch expression := expression.(type) {
	case *ast.Ident:
		return expression.Name
	case *ast.StarExpr:
		return namedType(expression.X)
	default:
		return ""
	}
}

func numericFormat(expression ast.Expr) string {
	switch expression := expression.(type) {
	case *ast.Ident:
		switch expression.Name {
		case "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
			return expression.Name
		case "float32":
			return "float"
		case "float64":
			return "double"
		}
	case *ast.StarExpr:
		return numericFormat(expression.X)
	}
	return ""
}
