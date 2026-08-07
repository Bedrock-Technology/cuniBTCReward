package main

import "testing"

func TestConvert(t *testing.T) {
	document := map[string]any{
		"swagger":  "2.0",
		"basePath": "/api",
		"consumes": []any{"application/json"},
		"produces": []any{"application/json"},
		"info":     map[string]any{"title": "test", "version": "1.0"},
		"definitions": map[string]any{
			"Request": map[string]any{"type": "object"},
		},
		"securityDefinitions": map[string]any{
			"bearerAuth": map[string]any{"type": "apiKey", "in": "header", "name": "Authorization"},
		},
		"paths": map[string]any{
			"/things": map[string]any{
				"post": map[string]any{
					"parameters": []any{map[string]any{
						"in": "body", "name": "body", "required": true,
						"schema": map[string]any{"$ref": "#/definitions/Request"},
					}},
					"responses": map[string]any{
						"200": map[string]any{"description": "", "schema": map[string]any{"type": "object"}},
					},
				},
			},
		},
	}

	if err := convert(document); err != nil {
		t.Fatal(err)
	}
	if document["openapi"] != "3.0.3" {
		t.Fatalf("openapi = %v", document["openapi"])
	}
	for _, removed := range []string{"swagger", "definitions", "securityDefinitions", "consumes", "produces"} {
		if _, ok := document[removed]; ok {
			t.Errorf("legacy field %q was not removed", removed)
		}
	}

	operation := document["paths"].(map[string]any)["/things"].(map[string]any)["post"].(map[string]any)
	if _, ok := operation["parameters"]; ok {
		t.Error("body parameter was not removed")
	}
	requestBody := operation["requestBody"].(map[string]any)
	schema := requestBody["content"].(map[string]any)["application/json"].(map[string]any)["schema"].(map[string]any)
	if schema["$ref"] != "#/components/schemas/Request" {
		t.Errorf("request schema ref = %v", schema["$ref"])
	}
	response := operation["responses"].(map[string]any)["200"].(map[string]any)
	if response["description"] == "" || response["content"] == nil {
		t.Errorf("response was not converted: %#v", response)
	}
	components := document["components"].(map[string]any)
	if components["schemas"] == nil || components["securitySchemes"] == nil {
		t.Errorf("components were not converted: %#v", components)
	}
}

func TestConvertArrayRefAndQueryParameter(t *testing.T) {
	document := map[string]any{
		"swagger": "2.0",
		"paths": map[string]any{"/things": map[string]any{"get": map[string]any{
			"parameters": []any{map[string]any{
				"name": "ids", "in": "query", "type": "array",
				"items": map[string]any{"type": "integer"}, "collectionFormat": "multi",
			}},
			"responses": map[string]any{"200": map[string]any{
				"description": "ok", "schema": map[string]any{
					"$ref": "#/definitions/Thing", "items": map[string]any{"type": "object"},
				},
			}},
		}}},
		"definitions": map[string]any{"Thing": map[string]any{"type": "object"}},
	}

	if err := convert(document); err != nil {
		t.Fatal(err)
	}
	operation := document["paths"].(map[string]any)["/things"].(map[string]any)["get"].(map[string]any)
	parameter := operation["parameters"].([]any)[0].(map[string]any)
	if parameter["type"] != nil || parameter["schema"] == nil || parameter["explode"] != true {
		t.Errorf("query parameter was not converted: %#v", parameter)
	}
	response := operation["responses"].(map[string]any)["200"].(map[string]any)
	schema := response["content"].(map[string]any)["application/json"].(map[string]any)["schema"].(map[string]any)
	if schema["type"] != "array" || schema["$ref"] != nil {
		t.Fatalf("array response was not normalized: %#v", schema)
	}
	items := schema["items"].(map[string]any)
	if items["$ref"] != "#/components/schemas/Thing" {
		t.Errorf("array item ref = %v", items["$ref"])
	}
}

func TestConvertIsIdempotentForOpenAPI3(t *testing.T) {
	document := map[string]any{"openapi": "3.0.3"}
	if err := convert(document); err != nil {
		t.Fatal(err)
	}
}
