// Package public provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package public

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"sr6HKTrLIc4XzAz95nbzvwEAAP//CMqAAaxpAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
