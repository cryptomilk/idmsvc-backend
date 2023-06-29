// Package metrics provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package metrics

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

	"H4sIAAAAAAAC/3RSQYvcPAz9K0bfd2iZNJ7t9lC87KGwpeyhtFB6WuZgHCUxjWVjKcMsQ/57sWcyS6fd",
	"S5AVPenpPR3BxZAiIQmDOUJGTpEY6+MrSvauhi6SIEkJBQ+i02Q93ak9ZvaR7rfttv1wp9xoM6Pcz9K/",
	"+1hq2Y0YbIn+z9iDgf/0yzx9+sv6wYqFZVka6JBd9kl8JDDrfLWSglJyBpWeFWeOr6BOhcqyemEMDchz",
	"QjDAkj0Np6me+vh3n88HG9KEKvbq0/dH1SH7gVTvM4uyKeVo3agsdWpmT4Ma4mRpUCzW/VJv9j5h3qAb",
	"42aIOWxcpH6akWTjJrRksxu9oJM549tCCnPgb/0PzHvvCr1RJLHR2kXiOGGbsRuttC4GbZPXY2DfBb2/",
	"KVgvU4E8EkueA5LYskEhDQ2cHQIDxaMbWBqICckmDwZu2217Cw0kK2NVVJfPgPIPNahL0ZMoiQoPKTIq",
	"GVGFs9h9zPXtr0hc3IY6N9fsYwcGvqCs99X8eXbvt9vXLuZSp1dsPYk5BJufq2fXzIpAdmAwT7BmdtV0",
	"xly0AfN0vepPxrrPA+5xiimc2M95AgMalt2l4zXwenk8awYNkA3Fo5XCslt+BwAA//8RNhOOegMAAA==",
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
