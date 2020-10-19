// Package datagen Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// ../data/k8s-install/app.yaml
package datagen

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _k8sInstallAppYaml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x54\xcb\x6e\xdb\x3a\x10\xdd\xeb\x2b\x06\x46\x96\x57\x89\x6d\xe0\xb6\x28\x81\x2c\xda\x2c\xd2\x02\x49\x6a\xd4\x41\xbb\x28\xba\x18\x53\x13\x9b\x08\x5f\x25\x47\x6a\x0d\xc3\xff\x5e\x50\x96\x65\xca\x8f\x30\x40\x60\xf2\xcc\x9c\x33\x4f\x95\x65\x59\xa0\x57\xdf\x29\x44\xe5\xac\x80\x66\x52\xbc\x2a\x5b\x09\x98\x53\x68\x94\xa4\xc2\x10\x63\x85\x8c\xa2\x00\xb0\x68\x48\xc0\x66\x03\xd7\x1d\xfa\x84\x86\x60\xbb\xed\xa0\xe8\x51\x76\xf8\xd3\xfe\xba\x43\x37\x1b\x50\x2f\xbd\xd7\x47\x6b\x1d\x23\x2b\x67\x23\x94\x2d\x8e\x87\x17\xd1\x9a\x97\x10\xd0\x2e\x09\xae\x5e\x69\xfd\x1f\x5c\x35\xa8\x6b\x02\x71\x7b\x96\xa2\x65\x68\x35\x92\x35\x6c\xb7\x02\x46\xe9\xb2\x73\xda\x6e\x47\x1d\x23\xd9\x6a\x1f\x4d\x7f\x89\x9e\x64\x52\xf4\x2e\x70\x2b\x5d\xb6\x3f\x05\x7c\x78\x3f\x9e\xb4\xb4\xbb\x9c\x97\xc1\xcb\x21\x3a\xcd\xd0\x15\xb3\x2f\x00\x22\x69\x92\xec\x82\x68\x21\xf4\x5e\xc0\x1f\x5c\x7b\xa7\x2c\x97\x91\x42\x43\xa1\x00\xe0\xb5\x27\x01\x0f\x0e\xab\x4f\xa8\xd1\x4a\x0a\xc5\x71\x0f\xd0\xfb\x78\x73\x68\x04\x23\xd3\x4b\xad\xe7\xc4\x67\x9a\x71\x2a\xf0\x76\x27\x34\x2e\x48\xc7\x37\x22\xdc\x57\x64\x98\x8c\x41\x96\xab\x87\xcc\xf7\x62\x7e\xf1\x30\x19\x67\xd3\x27\xe3\x35\x32\x75\xb4\x59\x3e\xe9\xe8\x81\xc2\x45\x0d\x80\x7d\x94\x5d\xe3\xd3\x70\x7d\x31\xb8\xa4\x59\xad\xf5\x9c\x64\x20\xee\x26\x2b\x1d\x35\x44\x7a\xfe\x32\x1b\xe8\x63\xef\xde\x79\x30\x39\xe9\x48\x67\x19\x95\xa5\x70\xc2\x93\xc5\xd7\xcb\x1e\xd6\x85\x42\xab\x71\x20\xca\x02\x9b\x39\xad\xe4\xfa\x28\x92\xdd\x63\x6e\x2f\x9d\x31\x68\xab\x43\x7d\x4a\x18\xed\xab\x33\x3a\x14\x2d\x2c\x63\x6e\x72\x14\x56\x09\xa1\xb6\xd9\xad\x44\x29\xc9\x73\xc9\x2e\xe6\xaf\x4d\xd3\xe4\xd7\x6a\x71\x7b\x93\x3a\xd5\xfe\xbb\xae\x16\x39\xa6\x55\x64\xb2\x65\xda\x91\xdb\xf1\x75\xfb\x27\xfa\xfd\x19\xda\xa4\x4d\xc9\x6d\xa6\xbd\x4d\xbf\x81\x7b\x97\xbe\xce\xb3\xe1\x42\xee\xce\x60\x2d\x2f\xbb\x4c\x4f\x5c\xba\x5d\xed\xe6\x4d\x35\x64\x29\xc6\x59\x70\x0b\x12\x99\x2d\x4b\x3f\x77\xf2\x95\x38\x7f\x84\x6e\xfb\x07\xaa\x17\x29\x92\xd0\xfd\x09\x01\xf2\x4a\xc0\xcd\x19\xd2\x41\x5c\xe9\x44\xb9\xa2\x14\xf0\xe7\xe7\xe7\xd9\xbc\x47\x02\x45\x57\x07\x49\x31\xe7\x0d\xf4\xbb\xa6\xc8\x71\xa8\x25\x7d\x2d\x60\x32\x1e\x9b\xc1\xab\x21\xe3\xc2\x5a\xc0\xf4\xff\x77\x8f\xaa\x47\x1a\xa7\x6b\x43\x8f\xae\xb6\xc3\x26\xec\x6a\x96\x7a\x9e\x91\x98\x64\x36\xdb\x65\x92\x41\x91\x64\x1d\x14\xaf\xef\x9c\x65\xfa\x9b\x25\xfe\x12\xef\x83\xab\x7d\x1b\xcc\xb8\xd8\x8b\xdd\x69\x54\xe6\xb9\xfb\x1c\x74\xdf\xde\xe3\xef\xc1\x91\x7c\xbe\xf5\x69\x6a\x63\x7c\x74\x15\x45\x01\x3f\x61\xf4\x8d\xb0\xfa\x11\x14\xd3\x57\x2b\x69\x04\xbf\x8a\x8b\xf5\x3a\x57\xad\xc8\x2e\xb4\xdb\x3a\xb9\x57\xc5\xbf\x00\x00\x00\xff\xff\xc2\xa7\x1b\xde\x14\x07\x00\x00"

func k8sInstallAppYamlBytes() ([]byte, error) {
	return bindataRead(
		_k8sInstallAppYaml,
		"k8s-install/app.yaml",
	)
}

func k8sInstallAppYaml() (*asset, error) {
	bytes, err := k8sInstallAppYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "k8s-install/app.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"k8s-install/app.yaml": k8sInstallAppYaml,
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
	"k8s-install": {nil, map[string]*bintree{
		"app.yaml": {k8sInstallAppYaml, map[string]*bintree{}},
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
