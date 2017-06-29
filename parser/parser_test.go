package parser

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestNew(t *testing.T) {
	New()
}

func TestParser_ParseGeneralApiInfo(t *testing.T) {
	var expected = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample server Petstore server.",
        "title": "Swagger Example API",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "petstore.swagger.io",
    "basePath": "/v2",
    "paths": {}
}`
	gopath := os.Getenv("GOPATH")
	assert.NotNil(t, gopath)
	p := New()
	p.ParseGeneralApiInfo(path.Join(gopath, "src", "github.com/swag-gonic/swag/example/main.go"))

	b, _ := json.MarshalIndent(p.swagger, "", "    ")
	assert.Equal(t, expected, string(b))
}

func TestGetAllGoFileInfo(t *testing.T) {
	searchDir := "../example"

	p := New()
	p.GetAllGoFileInfo(searchDir)

	assert.NotEmpty(t, p.files["../example/main.go"])
	assert.NotEmpty(t, p.files["../example/web/handler.go"])
	assert.Equal(t, 4, len(p.files))
}

func TestParser_ParseType(t *testing.T) {
	searchDir := "../example"

	p := New()
	p.GetAllGoFileInfo(searchDir)

	for _, file := range p.files {
		p.ParseType(file)
	}

	assert.NotNil(t, p.TypeDefinitions["api"]["Pet3"])
	assert.NotNil(t, p.TypeDefinitions["web"]["Pet"])
	assert.NotNil(t, p.TypeDefinitions["web"]["Pet2"])
	assert.NotNil(t, p.TypeDefinitions["main"])
}
