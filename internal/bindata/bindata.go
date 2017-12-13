// Code generated by go-bindata.
// sources:
// internal/runtime/VERSION
// internal/runtime/runtime.go
// macro/raw.go
// macro/rawfile.go
// macro/rewrite.go
// macro/runtime.go
// DO NOT EDIT!

package bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _internalRuntimeVersion = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xca\x4d\x2c\x2e\x49\x2d\x02\x04\x00\x00\xff\xff\xd6\xa3\x09\x2d\x06\x00\x00\x00")

func internalRuntimeVersionBytes() ([]byte, error) {
	return bindataRead(
		_internalRuntimeVersion,
		"internal/runtime/VERSION",
	)
}

func internalRuntimeVersion() (*asset, error) {
	bytes, err := internalRuntimeVersionBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/runtime/VERSION", size: 6, mode: os.FileMode(420), modTime: time.Unix(1513164412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalRuntimeRuntimeGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x56\xcd\x6e\xe3\x36\x10\x3e\x8b\x4f\x31\xdd\x43\x97\x5a\x18\xf2\xf6\x6a\xd4\xbd\x74\x5b\xa0\x05\xb6\x29\x36\xbd\x14\x41\xd0\xd0\xf4\x28\xa2\x23\x91\x02\x45\x39\x30\x02\xbf\x7b\x31\x43\xea\xc7\xb1\xd7\xd9\xbd\x51\xf3\xfb\xcd\xcc\x37\xa4\x5a\xa5\x9f\xd4\x23\x82\xef\x6d\x30\x0d\x0a\x61\x9a\xd6\xf9\x00\xef\x1e\x4d\xa8\xfa\x4d\xa1\x5d\xb3\x6c\x54\x08\x15\x3e\x37\x3d\xd6\x35\xfa\xe5\xce\x1d\x96\x8d\xd2\xde\xbd\x13\x62\xb9\x84\x5f\x2b\x65\xa1\x0b\xbe\xd7\x41\x84\x43\x8b\x73\x01\xbc\x88\x4c\xab\x56\x69\x13\x0e\x60\x6c\x10\xd9\x5e\xd5\x3d\x76\x00\x70\x77\x6f\x6c\x40\x5f\x2a\x8d\x2f\x47\x91\x75\x68\xb7\x24\x7e\x2d\xf7\xa8\xf7\x97\xe4\xba\x76\x1d\x6e\x01\x60\xe3\x5c\x2d\x8e\x22\xa6\xa6\x28\xb3\xd4\x9c\x0c\x00\x4e\x3c\x5b\xef\x1a\xd3\xe1\x89\xf0\xc8\x95\x7c\xc2\x12\xbd\xc7\x2d\x94\x96\x3e\xff\xb9\xf9\x74\xb3\x02\x5d\xa3\xb2\x7d\x2b\xca\xde\xea\xd1\x42\xe6\x73\x77\x4a\xc5\x0d\x29\xbe\xa8\x67\xf9\x20\x4c\x09\xf2\x07\x19\x2a\xd3\x81\xb1\x5d\x50\x56\xa3\x2b\x47\xdf\x3c\x07\x8f\xa1\xf7\x16\x2c\x3e\xcf\x22\x8a\xbd\xf2\xd0\x61\x5d\xc2\x1a\xc8\x57\xb0\xa0\x85\x35\xdb\xfd\x1d\x51\x4b\x82\x11\x8c\xb3\xd2\x63\xe7\xea\x3d\x2e\xc0\xe3\x0e\x75\xc8\x09\x04\x79\x17\x49\x01\x6b\x48\xa7\x51\x4e\x76\x2c\xa6\x83\x38\xe6\x42\xb0\x22\x54\x68\x61\x0d\x2d\x1f\x8a\x8d\xb1\x5b\xd9\xe6\x51\xa5\x55\xd0\x15\xeb\xf8\x34\x2a\x1f\x72\x9a\x0c\x17\x31\x2b\x9c\x5c\x1e\xf2\xd4\x4c\x62\x81\xc5\x1a\x06\xc0\xb1\x81\x49\x2a\xe7\xa4\xc8\xe1\x03\x53\xe6\x65\x8c\xf9\x23\x7d\xbf\x88\x6c\xe4\xce\x0a\x86\xd3\x42\x64\xc3\xb8\x6e\x69\xd8\xce\x42\xa8\x10\x74\x8c\x1b\x93\x48\x1d\x23\xe6\x6c\x22\x23\x0b\x66\xf3\x3a\x1b\x9e\x29\x41\x17\x89\x51\x94\xf6\xbc\xb4\xd4\xfe\xd4\x44\x49\x13\xf9\xcd\x7b\xe7\xe5\xfb\x2e\xa1\x50\x90\x02\x24\x28\xef\xf3\x9c\xba\x74\x14\x22\x5b\x2e\x81\x68\x0c\x2d\xda\xad\xb1\x8f\x9c\xaf\x46\x2b\x75\xc1\xec\xce\xe1\x17\xf8\x98\xf2\xea\x3d\xac\xd6\x90\x14\x77\x1f\xef\x45\x96\xfd\xc7\x23\xd3\x7b\x6a\x47\x94\xc3\x64\xf1\xd3\x8a\x4c\x66\x40\x49\x3c\x50\x20\x56\xce\x30\xae\x96\x14\x8d\x4f\xe0\x3a\xd7\x80\xb1\xb0\xe9\xcb\x12\xfd\x0c\x70\x5c\xdf\x1c\x7e\xa6\x86\x0d\x33\xe4\x49\x25\x15\xac\x41\xb5\x54\xe8\x68\xbc\x80\x08\xe3\xbb\x41\x58\x77\xd2\xb6\x05\x6c\x6a\xa7\x9f\x58\x15\x57\xd3\x9b\xc7\x2a\x80\x75\xcf\x10\x77\xad\xe3\x4d\xa1\x81\xc8\xa2\x28\x72\xb6\xec\x2a\xd7\xd7\x5b\xd8\x20\x94\xc6\x22\x94\xce\x47\x63\xad\x3a\x5c\xc0\xae\xef\x02\x28\x0b\xbf\xff\xfb\xc7\x74\x31\xac\xd6\xf3\xad\xcc\x74\x11\xef\xa6\x59\x61\x2c\x58\x70\x26\xaa\x9d\xeb\x5b\x01\xc4\x42\x17\x22\x1b\x42\xad\x20\x1d\x88\xb5\xd3\xca\x24\x61\xa2\xf1\x17\xaa\x51\x45\xdf\xd7\xfc\x25\xdd\xf9\x5d\xb3\x5c\x42\x6a\xf6\xb5\x11\x0d\xa4\x8a\xf4\x67\x56\x45\xd5\x44\xab\x98\xf3\x64\x78\xa3\x51\x64\xd6\xdb\x13\x9b\x38\x96\xc6\xd6\x5b\x9e\xd3\x30\x36\xee\x52\x37\xc3\xc7\xdf\x13\x3c\xde\x1f\x46\xc7\x8a\x09\x1c\x7d\x32\xb6\xa1\xfd\x83\x05\x21\x13\xd9\xd9\xd2\x66\x27\x37\x91\xdd\x16\xed\x9b\x3b\x7b\x71\x63\xbf\x99\xa6\x5c\xf0\x95\xac\x73\xd3\xb7\x43\xb2\xf3\x59\x33\x23\xc2\xef\xbb\xa1\x2e\x6e\xd2\xb0\x8b\x69\x87\xbe\xca\xf5\xe1\x82\x19\xb9\xce\x82\xc5\xc0\xd9\xaf\x91\xf8\x33\xd9\xfb\x89\xa8\xf1\x39\x7e\x2d\x25\xe8\xb7\x18\xe4\x13\x1e\xe8\x95\xe6\xa5\xbe\x70\x3b\x47\x37\x91\xfd\x79\x7b\xf3\x97\xcc\xa1\x51\xed\x5d\xb4\xbe\x7f\xf5\x66\x73\x92\x26\x26\x99\x5e\x7d\xb7\xd9\x5d\xf1\x89\x60\x41\x7b\x54\x01\x3b\x50\x7c\x67\x34\x2a\x3d\xf0\x9f\x55\x2b\x07\x00\xf3\x07\x29\x26\x89\x11\xe2\x8e\x36\xf0\x21\x0a\xe9\x95\xf9\xc6\x92\xf8\x47\xa1\x70\x9b\xdd\xdd\x13\x1e\xee\xa7\x05\x1c\xa6\x79\x31\xfa\xd5\x2e\xcc\x30\x72\x60\x71\x14\xff\x07\x00\x00\xff\xff\x12\x46\x68\x69\xd2\x09\x00\x00")

func internalRuntimeRuntimeGoBytes() ([]byte, error) {
	return bindataRead(
		_internalRuntimeRuntimeGo,
		"internal/runtime/runtime.go",
	)
}

func internalRuntimeRuntimeGo() (*asset, error) {
	bytes, err := internalRuntimeRuntimeGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/runtime/runtime.go", size: 2514, mode: os.FileMode(420), modTime: time.Unix(1513164412, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _macroRawGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\xcb\x41\x0a\x83\x40\x0c\x46\xe1\x75\x73\x8a\x1f\x57\x2d\xc8\xb8\xef\x11\xba\xb4\x27\x08\x43\x94\xb4\x9a\x91\x64\xac\x05\xf1\xee\xc5\x5d\x97\x8f\xc7\xb7\x70\x7e\xf3\x28\x98\x39\x7b\x21\xea\x3a\xf4\xbc\x41\x6d\x52\x93\x00\xe3\xf1\x84\x7c\x17\x97\x08\x2d\x76\xee\x57\xdc\x9b\x9e\xb7\xb6\xcc\x5a\x1b\x1a\x56\xcb\xa7\xb8\x86\x67\x44\x75\xb5\xb1\xc5\x87\x3d\x90\x52\x52\xab\xe2\x03\x67\xd9\x8f\x1b\xfe\x02\x3b\x5d\x5c\xea\xea\x06\xd3\x89\x0e\xfa\x05\x00\x00\xff\xff\xf6\x6c\xce\x61\x85\x00\x00\x00")

func macroRawGoBytes() ([]byte, error) {
	return bindataRead(
		_macroRawGo,
		"macro/raw.go",
	)
}

func macroRawGo() (*asset, error) {
	bytes, err := macroRawGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "macro/raw.go", size: 133, mode: os.FileMode(420), modTime: time.Unix(1513164284, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _macroRawfileGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x1c\xca\xb1\x0d\xc2\x40\x0c\x05\xd0\x1a\x4f\xf1\x95\x0a\x24\xa4\xf4\x0c\x40\x41\xcb\x04\xd6\xc9\x17\x0c\x8e\x13\xf9\x9c\x2a\xca\xee\xe8\x52\xbf\xb7\x72\xf9\xf1\x24\x98\xb9\xc4\x42\x34\x8e\x78\xaa\x09\xd4\x4d\x5d\x1a\x18\xaf\x37\xaa\x9a\x74\xf9\xb6\xc7\xd0\xf5\xbe\xcc\x9a\x03\xd5\xcd\xcb\xb9\xaf\x3d\xac\x9c\x1f\xb4\x0c\xf5\xe9\x06\xf5\x94\xa8\x5c\x64\x3f\xb0\xd3\x25\x24\xb7\x70\xb8\x1a\x1d\xf4\x0f\x00\x00\xff\xff\x95\xdc\x11\x78\x72\x00\x00\x00")

func macroRawfileGoBytes() ([]byte, error) {
	return bindataRead(
		_macroRawfileGo,
		"macro/rawfile.go",
	)
}

func macroRawfileGo() (*asset, error) {
	bytes, err := macroRawfileGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "macro/rawfile.go", size: 114, mode: os.FileMode(420), modTime: time.Unix(1513164284, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _macroRewriteGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\xcc\x41\xae\x82\x40\x10\x45\xd1\xf1\xef\x55\xbc\xf4\xe8\x9b\x10\x98\xbb\x0c\x77\x50\x36\x0f\x2c\x85\x82\x54\x37\x6a\x42\xd8\xbb\x21\xd1\x84\xe1\x1d\xdc\x33\x4b\x7a\x48\x4f\x8c\x92\x7c\x0a\xa1\x69\x70\xe1\xcb\xb5\x10\x62\xe0\x7b\x76\xe6\xac\x93\xc1\x39\x0f\x92\xd4\x7a\x94\x1b\xf1\x14\x57\xb9\x0e\xcc\x90\x0c\x23\x5b\xb6\xfb\x7a\xcf\xe7\xf8\xdd\xab\x69\xd4\x12\x43\xb7\x58\xfa\x89\xff\x3b\x87\x5c\x5c\xad\xaf\x0e\x44\x5d\xd7\x6a\x85\xde\x49\xe2\xba\x9d\x70\x08\xac\xe1\xcf\x59\x16\x37\xc4\x18\xb6\xf0\x09\x00\x00\xff\xff\xf8\x63\xd0\xc5\xae\x00\x00\x00")

func macroRewriteGoBytes() ([]byte, error) {
	return bindataRead(
		_macroRewriteGo,
		"macro/rewrite.go",
	)
}

func macroRewriteGo() (*asset, error) {
	bytes, err := macroRewriteGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "macro/rewrite.go", size: 174, mode: os.FileMode(420), modTime: time.Unix(1513164284, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _macroRuntimeGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x34\xca\x41\x0a\x02\x31\x0c\x40\xd1\xb5\x39\x45\x98\x95\x82\x74\xf6\x1e\xc3\x1b\x84\x98\x0c\x51\x9b\x91\x34\x5d\x95\xde\x5d\x04\x67\xf9\x3f\xef\x43\xfc\xa2\x4d\xb0\x12\xc7\x0e\xb0\xae\x78\xef\x9e\x56\x05\xd5\x7f\xf5\x6c\xb7\xe5\x7f\xae\x7b\xb5\x5c\x40\xbb\xf3\x81\xce\x0f\xd1\x86\xa5\x94\x96\x61\xbe\x5d\xd0\x3c\x25\x94\x58\xc6\xc4\x01\xa7\x90\xec\xe1\xe8\xf6\x86\x09\xdf\x00\x00\x00\xff\xff\x36\x56\xba\xa4\x6b\x00\x00\x00")

func macroRuntimeGoBytes() ([]byte, error) {
	return bindataRead(
		_macroRuntimeGo,
		"macro/runtime.go",
	)
}

func macroRuntimeGo() (*asset, error) {
	bytes, err := macroRuntimeGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "macro/runtime.go", size: 107, mode: os.FileMode(420), modTime: time.Unix(1513164284, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
var _bindata = map[string]func() (*asset, error){
	"internal/runtime/VERSION":    internalRuntimeVersion,
	"internal/runtime/runtime.go": internalRuntimeRuntimeGo,
	"macro/raw.go":                macroRawGo,
	"macro/rawfile.go":            macroRawfileGo,
	"macro/rewrite.go":            macroRewriteGo,
	"macro/runtime.go":            macroRuntimeGo,
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
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"internal": &bintree{nil, map[string]*bintree{
		"runtime": &bintree{nil, map[string]*bintree{
			"VERSION":    &bintree{internalRuntimeVersion, map[string]*bintree{}},
			"runtime.go": &bintree{internalRuntimeRuntimeGo, map[string]*bintree{}},
		}},
	}},
	"macro": &bintree{nil, map[string]*bintree{
		"raw.go":     &bintree{macroRawGo, map[string]*bintree{}},
		"rawfile.go": &bintree{macroRawfileGo, map[string]*bintree{}},
		"rewrite.go": &bintree{macroRewriteGo, map[string]*bintree{}},
		"runtime.go": &bintree{macroRuntimeGo, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
