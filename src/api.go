package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func src_api_file_wren() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x4a, 0xce,
		0x49, 0x2c, 0x2e, 0x56, 0x70, 0xcb, 0xcc, 0x49, 0x55, 0xc8, 0x2c, 0x56,
		0x70, 0x4a, 0x2c, 0x4e, 0x05, 0xb3, 0xab, 0xb9, 0x14, 0x80, 0x20, 0x2f,
		0xb5, 0x1c, 0xc8, 0x02, 0x33, 0x6b, 0x21, 0x54, 0x50, 0x6a, 0x62, 0x8a,
		0x46, 0x1a, 0x50, 0x45, 0x5e, 0x62, 0x6e, 0xaa, 0x26, 0x54, 0x19, 0x08,
		0x14, 0xa5, 0x96, 0x94, 0x16, 0xe5, 0x29, 0x14, 0x97, 0x16, 0xa4, 0x16,
		0xe9, 0xa1, 0xaa, 0x42, 0xd6, 0x1f, 0x5e, 0x94, 0x59, 0x92, 0x0a, 0x97,
		0xd2, 0x51, 0x48, 0x49, 0x2c, 0x49, 0xd4, 0x51, 0x00, 0xea, 0xc9, 0xc5,
		0x69, 0x18, 0x1e, 0x2d, 0x50, 0x93, 0x6b, 0xb9, 0x00, 0x01, 0x00, 0x00,
		0xff, 0xff, 0xb3, 0xbf, 0x6e, 0x2c, 0xc6, 0x00, 0x00, 0x00,
	},
		"src/api/file.wren",
	)
}

func src_api_main_wren() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xe2, 0xca,
		0x4d, 0xcc, 0xcc, 0xd3, 0x4b, 0x4e, 0xcc, 0xc9, 0xe1, 0x02, 0x04, 0x00,
		0x00, 0xff, 0xff, 0x2b, 0x79, 0xfa, 0xfa, 0x0b, 0x00, 0x00, 0x00,
	},
		"src/api/main.wren",
	)
}

func src_api_markdown_wren() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x64, 0x8e,
		0x31, 0xca, 0xc3, 0x30, 0x0c, 0x85, 0xf7, 0x9c, 0x42, 0xff, 0xe6, 0x2c,
		0x7f, 0x4e, 0xe0, 0xa5, 0x43, 0xa1, 0x43, 0xa1, 0xd0, 0x03, 0x14, 0x93,
		0x88, 0xc6, 0xd4, 0x76, 0x8c, 0xa4, 0x90, 0xa1, 0xe4, 0xee, 0xb5, 0x83,
		0x93, 0x26, 0x54, 0xd3, 0x93, 0xf4, 0xde, 0xc7, 0x6b, 0x9d, 0x61, 0x86,
		0xab, 0xa1, 0x57, 0x37, 0x4c, 0x01, 0x2c, 0xc3, 0xc9, 0x30, 0x6e, 0xfb,
		0xbb, 0x82, 0x34, 0x7e, 0x5b, 0xe1, 0xb1, 0xe9, 0x79, 0x79, 0x45, 0x43,
		0x8c, 0x5d, 0x12, 0xe9, 0x55, 0xf4, 0x5c, 0x1d, 0x42, 0x5a, 0xb5, 0x43,
		0x10, 0x0c, 0x52, 0x1f, 0xe2, 0x1a, 0xca, 0xf9, 0x00, 0xd2, 0xaa, 0x17,
		0xef, 0xea, 0x1d, 0x4d, 0x43, 0xbe, 0xac, 0xd0, 0x80, 0x53, 0xea, 0xb4,
		0xc8, 0x72, 0xb9, 0x65, 0xdb, 0x5d, 0xc8, 0x86, 0xa7, 0x5a, 0xe1, 0x75,
		0xe9, 0x9d, 0x87, 0x50, 0x46, 0x0a, 0xc0, 0x63, 0x44, 0xfa, 0x5f, 0xcc,
		0x5f, 0xdb, 0x0f, 0xe7, 0x6c, 0x1d, 0xaa, 0x68, 0xa4, 0xdf, 0x13, 0x9a,
		0x06, 0x2e, 0x3e, 0x3a, 0xf4, 0xb9, 0xac, 0xf4, 0x96, 0xff, 0x4a, 0x6e,
		0xae, 0x3e, 0x01, 0x00, 0x00, 0xff, 0xff, 0x82, 0xe6, 0x00, 0x27, 0x3e,
		0x01, 0x00, 0x00,
	},
		"src/api/markdown.wren",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"src/api/file.wren": src_api_file_wren,
	"src/api/main.wren": src_api_main_wren,
	"src/api/markdown.wren": src_api_markdown_wren,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
func AssetDir(name string) ([]string, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	pathList := strings.Split(cannonicalName, "/")
	node := _bintree
	for _, p := range pathList {
		node = node.Children[p]
		if node == nil {
			return nil, fmt.Errorf("Asset %s not found", name)
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"src": &_bintree_t{nil, map[string]*_bintree_t{
		"api": &_bintree_t{nil, map[string]*_bintree_t{
			"file.wren": &_bintree_t{src_api_file_wren, map[string]*_bintree_t{
			}},
			"main.wren": &_bintree_t{src_api_main_wren, map[string]*_bintree_t{
			}},
			"markdown.wren": &_bintree_t{src_api_markdown_wren, map[string]*_bintree_t{
			}},
		}},
	}},
}}
