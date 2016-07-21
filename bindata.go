// Code generated by go-bindata.
// sources:
// template.txt
// DO NOT EDIT!

package main

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

var _templateTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x57\x4d\x73\xdb\x36\x10\x3d\x93\xbf\x62\xc3\x49\x33\x64\xc3\xd0\xf5\xd5\x1d\x1f\x9c\x58\x99\x6a\xa6\xb1\x53\xdb\x69\x0f\x1e\x4f\x03\x51\x90\xc4\x96\x5f\x01\x40\x5b\x1e\x0d\xff\x7b\x77\x01\x4a\x20\x25\x52\xb6\xd5\xe6\xd0\x93\xf0\xb1\x58\xbc\x7d\xbb\x0f\x5c\x1d\x1d\xc1\xf9\x25\x5c\x5c\xde\xc0\xe8\x7c\x7c\x03\x37\xbf\x8c\xaf\xe1\xe3\xf8\xd7\x11\xbc\x72\x71\x6b\xac\x20\x91\x30\xe7\x39\x17\x4c\xf1\x29\x4c\x1e\x61\x5e\x4c\x98\xc2\x45\x55\x14\x69\x08\xb2\xa8\x44\xcc\x61\x26\x8a\x0c\x56\xab\xe8\x5a\x4f\xeb\x3a\x72\x4b\x16\xff\xcd\xe6\x9c\x16\x3f\x9b\x61\x5d\xbb\x6e\x92\x95\x85\x50\xe0\xbb\x8e\x37\x79\x54\x5c\x7a\x38\xe0\x79\x5c\x4c\x93\x7c\x7e\xf4\x97\x2c\x72\x5a\x98\x65\x8a\x7e\xa4\x12\x71\x91\xdf\xd3\x70\x9e\xa8\x45\x35\x89\xe2\x22\x3b\x7a\x28\xd4\x82\x8c\xd3\x62\x8e\x3b\xab\xd5\x3b\x10\x2c\xc7\x7b\xa2\xb1\x76\x2d\xf1\x16\x5c\x8d\xcc\xcf\x3b\xe0\xf9\x14\x70\x1c\xb8\xae\x7a\x2c\x35\x9a\x0b\x96\x21\x14\x40\xef\x55\xac\x56\x88\xa9\xe5\xe3\x13\x57\x8b\x62\x2a\xe9\xc4\xac\xca\x63\xf0\x7f\x5c\xad\x5e\x37\x27\x02\x7b\xd8\xd7\x41\x09\x96\xc9\x6b\x25\x10\x0c\x6e\xd2\xd2\x15\x97\x55\xaa\xec\x1a\xac\x5c\xe7\x9e\x09\x8a\xd6\x91\x2a\x53\x74\x27\xee\xe0\x6c\x52\xcd\x00\x34\x01\xd1\xfb\x6a\x36\xe3\x02\xd7\x98\x98\x4b\xb8\xbd\x4b\x72\xc5\xc5\x8c\xc5\x1c\xa1\x39\x08\xbb\x1d\xe2\x47\xc1\xe6\x19\xcf\x4d\x90\xb8\x91\xcc\x20\xfa\x50\xe8\x08\x01\xc7\x08\x81\x66\x38\x59\xe1\x98\xd3\x68\x7d\xf3\x29\x7c\xa5\xec\xe0\x10\x77\xbf\xea\xc3\x1b\xaf\x67\x78\xb1\xb6\xd4\x8b\xe4\x73\x2c\xc7\xb9\x59\x59\x9f\x36\xc8\x25\x86\x58\xa6\x88\xcd\xa7\xe5\x10\xbc\xd7\x2b\xef\xad\x87\x8e\x7f\x67\xa2\xae\x71\x58\x7b\xa1\x3e\xe5\xb4\xec\x39\x53\xbe\x17\xfe\x20\xbd\x10\x52\x9e\xfb\x6b\xeb\x20\xb8\x3d\x3e\xb9\x0b\xe1\xdd\x71\xa0\x8f\xcc\x0a\x01\x7f\x62\x3d\xc1\xc9\x69\x03\x6c\x6d\x49\x34\x3a\x2d\x74\x37\x98\xc8\xe8\x2c\x4d\x58\x83\xda\x69\xb8\x3b\x05\x56\x96\x18\xb5\x4f\xb3\x90\x8e\x6b\xcb\xcf\x22\xc9\x12\x95\xdc\x53\xde\x64\x10\x58\x5f\x3c\x95\x7c\xaf\x07\xd9\x36\x6e\xd8\x74\x9c\x0d\x53\x74\xbe\x1f\xd0\x73\xe1\x58\x2e\xfa\x7d\xbe\x67\x32\x89\xf7\xfb\x34\xe7\xdd\xdd\x90\xa8\xf0\x70\xff\x4b\x3e\xe5\x22\x4d\x72\x6e\xa8\xbc\xbd\xa3\xaa\x73\x4d\x54\x5b\x9b\x21\x70\x21\xf0\x0a\x92\x61\xf4\x89\x09\xb9\x60\xa9\xdf\xb9\xc2\x41\x64\x64\xf3\xea\x14\xf2\x24\x5d\xe7\x05\x95\x18\x8d\x84\x28\x84\x8f\x7b\x0d\x65\x82\xab\x4a\xe4\x96\xaf\x21\xf8\x5d\x04\x36\x8e\x86\xed\xf6\x18\x35\x13\xfd\x21\x12\xc5\x8d\xbe\x74\x09\x06\x3b\x32\xa8\x6d\xe5\xdb\xc3\x46\x86\xe5\x62\x5b\x5e\x54\x73\x89\x2d\x38\x8d\x91\x82\x42\xcb\x0d\xd4\x72\x41\x75\xee\xbd\x6d\x5e\xa3\x68\xac\x0a\xe6\x27\x6f\x8f\x29\x67\xe4\xf9\x5b\xc5\xc5\x23\xf9\xc0\x57\x2b\xba\x2e\x11\x9a\x9a\xf9\x84\xb5\x81\x19\x84\x78\x71\x14\x45\x24\x65\x62\xea\x9c\x4f\xaa\xb9\xaf\x4f\x05\xed\x15\xba\xbc\x31\xd3\x11\xf1\x6f\x10\x5d\x96\x54\x06\xe0\x25\xb9\xe4\x42\x79\x2d\xd5\x9f\x73\xa9\x4c\xa2\x29\xb2\x29\xce\xb6\x63\x23\xcb\x46\xe0\xd6\xb8\x39\xbd\x53\x5c\x8e\xf6\xb0\x89\x99\x66\x21\xbc\xe9\xe6\x1e\x79\x6d\x09\xe6\xc9\xf2\x7a\xa6\xcb\xb5\xac\x5a\x43\x53\x85\xa4\x97\x25\x7e\x48\x7e\x23\xa6\xae\x8a\x07\x43\x59\x08\x6b\x9e\xa2\xeb\x98\xe5\xda\xad\x66\x6d\xa0\x38\x7b\x6a\xd3\x2e\xad\x93\xd0\x59\xdc\xe4\xa1\x5b\xc6\xfb\x19\xcd\x0b\xd5\xc7\x2a\xc1\x39\x69\x04\xf5\x25\xcf\xac\xa4\xb6\x85\xd7\xa5\xa5\x5f\x66\xbd\x3a\x6b\x0b\xcd\xa9\xf7\x90\x6a\xec\xc8\x9f\xdb\x4d\xa5\xe0\xd2\x08\xff\xc4\x72\x3e\x5a\xf2\x78\x9b\xef\xef\xca\x30\xfc\xa4\x31\x6c\x88\x6e\x56\x11\x5b\x84\xa9\x97\x67\xf8\x89\x8c\xb1\xfd\xf0\x8d\xde\xd7\xa2\x36\x61\xc0\x96\x58\xaa\x72\x8a\xad\x8a\x11\xcb\x0b\x82\xeb\x89\x6d\x37\xb4\xbe\xc8\xfa\x03\xdb\x8e\xcb\xa0\x19\x8a\x6a\x28\x96\x29\x4f\xf9\xff\x31\x16\xfb\x21\x6b\x47\x33\xe7\xcd\x1b\x36\xf8\x66\x35\x9f\x78\x2d\xa7\x2b\x7d\xc3\x8e\xaa\x74\x9f\xa5\x77\xcc\x93\xa3\xc9\x68\xdb\x5e\xf0\x87\xd1\xb2\x44\x60\x32\x29\xf2\x9d\xaf\x48\xab\xef\xb1\x22\xee\x74\x16\x6d\x05\x3f\xfd\x86\xbd\xf4\xab\x3b\xe4\xf1\x05\x1f\xc2\x03\xdf\xc7\xef\x52\x14\x4d\x35\xec\x63\xb5\xff\x6d\xdc\x34\x77\xa7\x9b\xce\x68\x37\x71\xeb\x58\x9f\x7e\x40\xdf\xfc\x47\x7d\xca\x10\xe7\x8d\xdd\x70\x6d\xa7\x89\x6c\x8a\x5b\xa0\x1c\x76\x84\xaa\xb3\x74\xa8\x52\xed\xdb\xdd\x56\xe0\x94\xe3\xff\x06\xa0\xdb\xa2\x0f\x69\x21\xb9\x96\x9e\x91\x16\x53\x6c\x4b\x15\x1a\x19\xb5\x3b\xda\xfe\x82\x2f\x95\x1f\x98\xeb\x96\x0d\xcc\xb6\x84\x46\x29\xcf\xfa\xd2\xa1\x1d\xdb\xe2\xc5\x59\x08\x4b\x7d\xed\xde\x36\xa4\xa7\x36\x86\x25\xf7\xbc\xbe\xa1\xdb\xbd\x1f\xde\x8c\xf4\xc8\xae\xdb\xee\x77\x26\xa6\x1a\x35\x85\x87\x76\x1f\xbb\xb9\x74\x9e\x66\x69\xa0\xbd\x78\xa6\x86\x0e\x15\xd1\xbf\x6d\x43\x06\x58\xac\x6d\xd9\x37\x5c\xa2\x37\x3f\xf8\xf9\x40\x21\xd8\x6f\x91\xa9\x47\xdd\xe1\xd8\xfb\x6a\xd7\x8e\xff\x09\x00\x00\xff\xff\x9c\xba\x9e\x4a\xeb\x10\x00\x00")

func templateTxtBytes() ([]byte, error) {
	return bindataRead(
		_templateTxt,
		"template.txt",
	)
}

func templateTxt() (*asset, error) {
	bytes, err := templateTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template.txt", size: 4331, mode: os.FileMode(420), modTime: time.Unix(1469075094, 0)}
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
	"template.txt": templateTxt,
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
	"template.txt": &bintree{templateTxt, map[string]*bintree{}},
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

