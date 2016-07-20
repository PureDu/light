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

var _templateTxt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x56\x4d\x73\xdb\x36\x10\x3d\x53\xbf\x62\xc3\x49\x33\x64\x43\xc3\x75\x7b\x73\xc7\x07\x37\x76\xa6\x9a\x69\xec\xd4\x76\xda\x83\xc7\xd3\x40\x14\x24\xb1\xa5\x00\x06\x00\x6d\x79\x34\xfc\xef\x5d\x00\xa4\x40\xca\xa4\xfc\x91\xe6\x50\x5d\x04\x2c\x16\x8b\x7d\x8b\xf7\xc0\xdd\xdf\x87\x93\x73\x38\x3b\xbf\x82\xd3\x93\xf1\x15\x5c\xfd\x3a\xbe\x84\xf7\xe3\xdf\x4e\xe1\xd5\x08\x97\xc6\x1a\x32\x05\x73\xc6\x99\xa4\x9a\x4d\x61\x72\x0f\x73\x31\xa1\x1a\x8d\x5a\x88\x3c\x01\x25\x4a\x99\x32\x98\x49\xb1\x84\xf5\x9a\x5c\xda\x69\x55\x91\x51\x41\xd3\x7f\xe8\x9c\x19\xe3\x47\x37\xac\xaa\xd1\x28\x5b\x16\x42\x6a\x88\x46\x41\x38\xb9\xd7\x4c\x85\x38\x60\x3c\x15\xd3\x8c\xcf\xf7\xff\x56\x82\x1b\xc3\x6c\xa9\xcd\xdf\x3c\xd3\x8b\x72\x42\x52\xb1\xdc\xbf\x13\x7a\x61\x3c\x72\x31\xc7\x95\xf5\x7a\x0f\x24\xe5\x18\x9c\x8c\x6d\x3c\x85\xa1\xd1\x4a\xdc\xdf\x1e\x30\x3e\x05\x1c\xc7\xa3\x91\xbe\x2f\x6c\x0a\x67\x74\x89\xe7\x83\xd2\xb2\x4c\xf5\x1a\x13\x69\xc5\xf8\xc0\xf4\x42\x4c\x95\xd9\x31\x2b\x79\x0a\xd1\xf7\xeb\xf5\xeb\x7a\x47\xec\x37\x47\x16\x89\xa4\x4b\x75\xa9\x25\x26\x83\x8b\xc6\x74\xc1\x54\x99\x6b\x6f\x83\xf5\x28\xb8\xa5\xd2\x40\x0c\x94\x5e\x6a\x73\x26\xae\xe0\x6c\x52\xce\x00\x2c\x6a\xf2\x4b\x39\x9b\x31\x89\x36\x2a\xe7\x0a\xae\x6f\x32\xae\x99\x9c\xd1\x94\x61\x6a\x01\xa6\xdd\x86\xf8\x5e\xd2\xf9\x92\x71\x07\x12\x17\xb2\x19\x90\x77\xc2\x22\x04\x1c\x63\x0a\x66\x86\x93\x35\x8e\x99\x19\x35\x27\x1f\xc1\x67\x73\x25\x38\xc4\xd5\xcf\x76\xf3\x26\xea\x31\x1e\x6c\x3d\xad\xd1\xc4\x1c\xab\x31\x77\x96\x66\xb7\xcb\x5c\x21\xc4\x22\xc7\xdc\x22\x63\x4e\x20\x7c\xbd\x0e\xdf\x86\x18\xf8\x0f\x2a\xab\x0a\x87\x55\x98\xd8\x5d\x41\xcb\x9f\x51\x1d\x85\xc9\x77\x2a\x4c\x20\x67\x3c\x6a\xbc\xe3\xf8\xfa\xe0\xf0\x26\x81\xbd\x83\xd8\x6e\x99\x09\x09\x7f\x21\x89\xe0\xf0\xa8\x4e\xac\xf1\x34\x65\x34\x3f\x5b\xa0\x23\xa0\x45\x81\xd0\x22\x33\x4b\x00\xab\xf5\xd3\x8f\x91\x8a\x5d\x8c\x0d\x08\x96\x2b\x66\x91\x5c\xe1\xa5\x93\x8f\x32\x5b\x66\x3a\xbb\x65\x35\xa6\xbe\x40\x9b\xb4\xfa\x43\x1c\xe7\x19\x55\xbb\xb7\xb7\x4e\x69\x81\x6c\x87\xab\xf7\x1b\x4a\xa0\xc3\x27\x3e\x65\x32\xcf\x38\x73\x20\xaf\x6f\x0c\x1f\xac\xc3\x83\xc5\x04\x98\x94\x78\xa4\x51\x05\xf9\x40\xa5\x5a\xd0\x3c\xea\x64\x1c\x60\xa6\xc6\xe7\xd5\x11\xf0\x2c\x6f\x2a\x86\x1a\x21\xa7\x52\x0a\x19\xe1\x9a\xf3\x0b\x24\xd3\xa5\xe4\xbe\x5c\x43\x70\xba\x19\x78\x1c\x35\xab\xda\x63\x64\x33\xf9\x53\x66\x9a\x39\xe6\x5b\x72\xc4\x0f\x08\x5a\x79\x4e\xfa\xcd\x4e\x20\xc5\x62\x9b\xf8\x86\x0d\x99\xa7\x82\xcd\xd1\x80\x42\xcf\x4d\xaa\xc5\xc2\x30\x30\x7c\x8b\x5c\x4b\x05\xbf\x25\x63\x2d\x68\x94\xbd\x3d\x30\x45\x37\x91\xbf\x94\x4c\xde\x9b\x18\xf8\x88\x90\xcb\x02\x53\xd3\xb3\xc8\xe4\x5a\xa7\x19\x27\x78\x30\x21\xc4\x88\xcc\x54\xea\x84\x4d\xca\x79\x64\x77\xc5\x6d\x8b\x39\xbc\x76\xb3\x88\xd8\x17\x20\xe7\x85\xa1\x05\x84\x19\x57\x4c\xea\xb0\xa5\xc7\x13\xa6\xb4\xbb\x68\x83\x6c\x8a\xb3\x6d\x6c\xc6\xb3\x96\x9e\x77\xae\x77\xf7\xf2\x35\xb0\x51\x36\xb8\xcd\x2c\x81\x37\xdd\xfb\xc7\xda\x7a\x86\x3d\x4e\xb1\x27\x86\xac\x2f\xb8\x3d\x74\x4c\x44\xcf\xab\x15\xbe\xed\xbf\x9b\x6a\x5d\x88\x3b\x57\xb6\x04\x9a\x5a\x91\xcb\x94\x72\x1b\xd6\x56\x6e\x80\xa0\x3d\xfc\xf4\xa6\xe6\x22\x3a\xc6\xcd\x5d\x74\xa9\xbc\xbb\xaa\x5c\xe8\xa1\xca\x9a\x94\x0e\x6b\x61\x7d\xe2\x4b\x2f\xad\x6d\x01\x76\x4b\xd3\x2f\xb7\x5e\xbd\xb5\x05\x17\x54\x3b\x0a\xeb\xfc\x4c\xbc\x51\xf7\x3a\x25\x53\xee\x01\x38\xf4\x75\x3f\x5d\xb1\x74\xbb\xe6\xdf\xb4\xca\xf0\x83\xcd\x61\x53\xec\xda\x8a\xb9\x11\xbc\x7e\x75\x8c\x1f\xb1\x14\xbb\x82\xc8\xe9\xbe\x11\xb7\x83\x01\x5b\xa2\x29\x8b\x29\x76\x10\x4e\x34\xcf\x00\xd7\x83\xed\x21\xb4\x3e\x64\xfd\xc0\xb6\x71\xb9\x6c\x86\x50\x0d\x61\x99\xb2\x9c\xfd\x1f\xb1\xf8\x0f\x5c\x1b\xcd\x9c\xd5\x6f\xd9\xe0\xdb\x55\xb7\x08\x56\x52\x17\xf6\x84\x5e\x65\xd9\x6e\xc8\xae\xba\xa7\xc7\x16\xa4\xed\x7f\xc6\xee\x4e\x57\x05\x26\xa7\x32\xc1\x1f\x7c\x51\x5a\xdd\x89\x17\x73\xd3\x9d\xf4\x2a\xf9\xf1\xf7\xec\xb9\x5f\xe1\xa1\x88\xcf\xf8\x30\xbe\xf0\xad\xfc\x26\xe4\xa8\x59\xb1\xab\xb2\xc3\xef\xe4\xa6\x15\x73\x68\xfa\x2f\xb0\xc1\xfb\xf8\x63\xfa\xe6\x3f\xea\x5d\x86\xea\x5e\xfb\x0d\xf3\x3c\xcf\x54\x4d\x74\x89\xd2\xd8\x88\x56\xaf\xdc\x25\xbd\x54\xb0\xfe\x09\x6f\x0b\x71\xca\xb0\xc1\x07\x73\x10\x79\x97\x0b\xc5\xac\x02\x9d\xc2\xa8\xa6\x5b\xc2\xb0\x49\x99\xee\xc7\xfa\x9f\xb1\x95\x8e\x62\x77\xdc\xaa\x7e\x56\x76\xab\x28\xb0\x31\x3d\x6d\x71\x96\xc0\xca\x9e\xb8\xb3\x21\xe9\x61\xc5\x6e\xc1\x3d\xad\x83\xe8\x48\xee\x2b\xda\x92\x1e\xd1\x75\x6e\xbc\x3b\x71\x3c\xb4\x15\x7c\x69\x1f\xf2\xf0\x2a\x83\xc7\x2b\xb5\xa3\xd1\x78\xa2\x82\x5e\x2a\xa1\xaf\x6d\x48\x06\x2a\x59\x79\xe6\xd7\xf5\xc4\x68\x51\xfc\xf3\x0b\xb5\xe0\xbf\x4a\x8e\x97\xb6\xd7\xf1\xe7\x55\x23\x3f\xfe\x37\x00\x00\xff\xff\x67\x8b\xc6\xe3\x8c\x10\x00\x00")

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

	info := bindataFileInfo{name: "template.txt", size: 4236, mode: os.FileMode(420), modTime: time.Unix(1468998116, 0)}
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

