// Code generated by go-bindata. DO NOT EDIT.
// sources:
// state.html.tmpl (5.057kB)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _stateHtmlTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x6d\x53\xe3\x38\xf2\x7f\xcf\xa7\xe8\xbf\x6b\xf7\x3f\x33\xc5\xd8\x4e\x42\x02\x03\xe7\xa4\x86\xe5\x61\x09\x3b\xc0\xf0\xb8\xc3\x5c\xdd\x0b\xc5\x6a\xc7\x4a\x64\xc9\x48\xed\x84\x90\xe2\xbb\x5f\xd9\x4e\x42\x08\x21\xcb\xec\x55\xdd\x5d\x5d\xad\x5e\x80\xad\x6e\xb5\xfa\xe1\xd7\xdd\x4e\x07\xff\xc7\x75\x48\xa3\x14\x21\xa6\x44\xb6\xd6\x82\xfc\x1f\x48\xa6\xba\x4d\x07\x95\xd3\x5a\x5b\x0b\x62\x64\xbc\xb5\x06\x00\x10\x90\x20\x89\x2d\x83\xa9\x76\xb3\x94\x33\x42\xe3\x5a\x62\x84\x81\x5f\x52\x4a\x2e\x29\x54\x1f\x0c\xca\xa6\x63\x69\x24\xd1\xc6\x88\xe4\x40\x6c\x30\x6a\x3a\x31\x51\x6a\x77\x7c\xdf\x12\x0b\xfb\x29\xa3\xd8\xeb\x68\x4d\x96\x0c\x4b\x43\xae\xbc\x50\x27\xfe\x6c\xc3\xaf\x7b\x0d\xaf\xe2\x87\xd6\x3e\xed\x79\x89\x50\x5e\x68\xad\x03\x42\x11\x76\x8d\xa0\x51\xd3\xb1\x31\xdb\xf8\x54\x77\xb7\x59\x9b\x6a\xea\x22\xdd\xab\xd6\xae\xfb\xdb\xdd\xcb\xed\x0e\xdb\x97\xf5\x6a\xf5\xf4\x7c\x37\x3d\x4c\xf6\x6a\x9b\x07\xc3\xdd\xb3\xa3\x4f\xbf\x77\xbf\xcb\xc6\xc9\xed\xed\xfd\x61\x14\xae\x9f\x86\x5f\x3b\x55\xfe\xdb\xaf\xbd\xad\xcb\xbe\x03\xa1\xd1\xd6\x6a\x23\xba\x42\x35\x1d\xa6\xb4\x1a\x25\x3a\xb3\xce\x0f\xd8\x95\x1b\xd1\xb3\x1c\xa5\x18\x18\x4f\x21\xf9\x2a\x4d\xfc\xcf\x91\x36\xc4\x86\x68\x75\x82\x7e\xa4\xd5\xf4\xd9\x8d\x0c\xe2\xe7\x86\x57\xad\x4f\xcc\x64\x52\xce\x0c\x9c\x5c\x5a\x5c\x55\x3e\xe7\xcb\xeb\x58\x97\xb4\x96\x24\x52\xb7\xa3\x89\x74\x02\xde\xf4\x5d\x28\x85\x06\xc6\x33\xde\x7c\x25\xec\xde\x1d\x0a\x4e\xf1\x0e\x54\x2b\x95\x9f\xff\x36\x23\x3e\x96\xe2\xfd\x89\xfc\xc0\x2f\xc3\xbc\x16\x74\x34\x1f\x4d\xee\xe6\x62\x00\xa1\x64\xd6\x36\x9d\x30\xd7\x5a\x28\x34\xce\x93\x2e\xe3\xf1\x4f\x36\x8c\x91\x67\x12\xcd\x7e\x96\xa4\xb0\xd3\x04\xa1\x38\xde\x83\x07\x95\xc7\xc7\x39\x3e\x11\x41\x97\xe0\xbd\x44\x05\xde\x07\xa8\xce\xd1\x16\xef\x61\x12\x0d\x41\xf1\xd7\x15\x2a\xd2\x90\x90\xdb\x80\xe4\xde\x65\x19\x69\x20\xbc\x27\x37\x44\x45\x68\x1c\xd0\x2a\x94\x22\xec\x37\x1d\xa9\x19\x3f\xbe\x3c\x3b\x7d\xff\xc1\x01\xa3\x25\x36\x9d\x4e\x46\xa4\xd5\x9c\xaa\xd3\x75\xa5\x61\x20\x70\x08\x8c\x73\x41\x42\x2b\x26\x81\x63\x27\xeb\x02\xcf\x92\xd4\x7e\x84\x54\x22\xb3\x08\x06\xef\x32\xb4\x04\x0c\x72\xb9\x60\xd0\xa6\x5a\x59\x7c\x21\x2e\x10\x53\xbd\x23\x66\x21\x62\x6e\x01\x91\x44\xba\x35\xa7\x15\xf8\xe2\xf9\xfd\x81\xcf\xc5\x60\xde\x7b\xa8\xf8\x9c\x27\xe6\xbd\x50\x18\x3d\x74\xb7\xfd\x6a\x6d\xc1\x88\x20\xae\xcf\x98\x3a\xee\x86\xd3\xba\x9c\x44\x20\xf0\xe3\xfa\x02\x6b\xba\xc4\xfe\x18\x61\x1a\x33\xd0\x11\x0c\x63\x54\x90\xe7\xb3\x15\xa4\x8d\x40\x0b\x5d\x24\x40\x75\x97\x61\x86\x3c\xcf\x31\x0d\x14\x23\x5c\x17\xd9\x0e\xe7\xf9\xb6\xb7\x60\xd5\xc2\x35\x01\xb1\x8e\xc4\xa9\x92\xe5\x4b\x11\x37\x89\x11\xe5\xe1\xac\x2f\x89\x4b\x40\x39\xfc\x66\x87\xf2\x17\x57\x8a\x6e\x4c\x4b\x79\xcd\xcb\xcd\x89\x10\x28\xe0\xdc\x74\x66\x90\xff\xd9\x69\xb5\xf7\x03\x9f\xe2\xb7\x1e\xa9\xe7\x47\x4e\x59\x82\x2b\x0f\x2d\x27\x14\x44\x9b\x32\xd5\x9a\xf8\xab\x9d\x23\x75\xc0\x64\xe0\x17\xbb\xaf\x1f\x5a\xc4\x51\x0e\x7d\x37\x14\x26\x94\x08\xc9\xa8\x04\x7f\x22\xdd\x0d\x07\x38\x23\xe6\x92\xee\x76\x73\x9d\x27\x99\xef\xbc\x2a\x18\x00\x8a\xca\xdc\x74\xf6\x98\x0c\x33\xc9\x08\x39\x74\x98\x45\x0e\x5a\x15\x91\x25\x91\x20\x50\xcc\x08\x62\x66\x01\x25\x4b\x73\xa2\x15\x2a\xc4\x82\x2e\x99\x25\x08\x75\x92\x08\xfa\x08\x5c\x0c\x04\xcf\x25\x8c\x80\x41\xa8\x95\x25\xa6\x08\x22\x16\x92\x36\x39\x9c\x6a\xde\x92\x78\xcd\x8c\x5c\x4c\x88\x27\xc2\x2a\x4f\x9f\xe2\x3d\x4d\xf0\xb7\x9c\x31\xf0\x97\x01\x22\xe7\x9d\x35\xae\xe7\x32\x9f\x2a\xdc\xfc\x1a\x8f\x0d\x53\x5d\x84\xe7\x35\xcd\x9b\xe6\xd7\x42\xcd\x7a\x92\xf6\x0a\x18\x4b\x22\x6f\x8d\xc7\xde\x05\xa6\xda\x6b\xef\x3f\x3e\x06\x3e\x2d\x51\x68\x9e\x7b\x55\x20\x21\x60\x93\x66\x33\x95\x79\x7d\xf1\xe5\xf1\x71\x65\xf0\xf3\xb5\x14\x31\xe5\x6e\x2a\x59\x88\x09\x2a\x6a\x3a\x65\x33\x71\xa6\x70\x79\x7e\xc5\x6a\xbd\xa0\xf0\x5e\xc9\x9f\x67\xce\x2b\xae\x9a\x99\xe1\xb3\x55\x20\xf9\x23\x17\x8d\xc7\x64\x32\x15\x32\xc2\xfd\xcc\xb0\xbc\x88\x83\x37\xcd\xb3\xb7\xb8\x78\x3c\xf6\xf6\x33\xf4\x0e\xb5\x49\x18\x81\x73\xa2\xd5\x47\xa8\xd4\xe0\x98\x29\xa8\x55\x2a\x9b\x50\x6d\xec\x54\xea\x3b\x95\x06\x9c\x5c\x5e\x39\xab\x04\x2e\xc7\xdd\x78\x8c\xd2\xfe\x59\xb4\x40\xa8\x65\x5e\x29\x9a\x4e\xb5\x52\x71\x66\xe5\x70\xae\xeb\xfd\x01\x42\xd2\x67\x7d\xd4\x69\x9d\xea\xe7\xd5\x5d\xa8\x59\xf9\x7f\x51\xb8\x17\x6c\xfb\x51\xab\x9f\x35\xb3\x39\xee\x97\xd9\x16\xf8\x45\x57\x78\xda\x5c\xe8\x8c\x7f\xb2\x13\xce\xf7\xa8\xb7\x75\xc3\x5d\x48\x8d\xd0\xf9\x37\x24\x14\x0d\x2f\x2f\x61\xcf\xdc\x45\x1a\xca\xef\x5c\x6f\xc9\xd9\xa1\x36\x7d\x34\x79\x19\x24\xa1\x32\x9d\x59\x39\x02\x8e\x85\x20\x9b\x97\xce\x04\x98\xe2\x60\x51\x71\x3b\x91\x52\x08\xec\x0a\xb2\x68\x06\x68\xfe\xea\xa3\x14\x97\x41\x13\xaa\xbb\x9a\xeb\xeb\x24\x4c\xab\xb9\x2e\x73\xe7\xab\xf0\x3f\xd2\x26\x4a\xf0\x15\xd8\xfb\xab\x53\xac\x5c\xff\xce\x4e\xe1\x4d\xe1\xf5\xc6\xc6\x30\xc5\xd9\x1b\xd9\x2f\xf1\xee\x7f\xb0\x41\xdc\x95\x05\xf4\xbf\xb1\x3b\xcc\x3f\xda\xd0\x88\x94\xc0\x9a\x70\xee\xd7\xb7\xe6\xe8\xf5\xee\x32\x34\xa3\x62\x94\x50\x3e\xba\x1b\x5e\xc3\xab\x7a\x56\x8a\xa4\xf8\x75\xdd\x5b\x3a\x3d\xd8\x8f\xbe\xf1\x87\x5a\x4c\x5f\x8f\x2a\xd2\x5e\x5e\xda\x86\xda\xbb\x4a\xb3\x9e\xff\x30\xaa\xef\xad\x9f\xfd\x9a\xb2\x44\x1f\xde\x8c\x36\x3e\x9d\xdc\xfc\xa2\x0e\xd6\xdb\x9d\xce\xcd\xed\x35\x0e\xd7\xcf\xcc\xde\x37\x76\xd1\x8f\x7a\xaf\x4f\x0f\x02\xbf\xd4\x75\x95\xe2\xcb\xc6\x06\xa9\x4e\x53\x34\x5e\xcf\x7e\xae\x7a\xd5\x4d\xaf\xe2\x73\x61\xc9\xcf\x12\x3e\xa5\xbc\x6e\xcc\xf9\xe6\xc1\xf6\xc5\xd1\xa0\xd3\x1e\x7d\x3f\x3c\xd6\x11\xad\xd7\x92\xe3\xce\x11\x3b\xf8\x5d\x72\x39\x68\x6f\xb7\xcf\x6e\x47\x0d\xb5\xf1\x70\xb3\xfd\xf0\x70\x45\x49\x7b\xe3\xba\x6f\xf9\xf9\xc5\xcd\x40\xdf\x9f\x44\x5a\xef\xea\x7f\xc9\x98\x1f\x98\xed\xf4\x16\x47\x3b\xcb\xcd\x39\xeb\xde\x5c\x0c\xb2\xdd\xab\xaf\xd5\x87\xad\xe3\xde\xd1\x97\x7e\x76\x76\xbd\xf5\x6d\xb8\x55\xa9\xaf\xc7\x9f\x36\x1a\x5f\xcc\xfa\xe6\xf9\x97\xed\xeb\xc1\x6d\xef\xfb\xc1\x46\x3b\xcd\x36\xaf\xd2\xad\x46\x6f\xeb\x97\xd8\xef\x5f\x54\x8e\x7f\x6b\xff\xa0\x39\x4f\xd8\xfb\xe9\xfd\xbb\xbf\x2f\xad\x8a\xff\x78\xf7\x61\x3a\x77\x79\xff\x61\xc6\x1e\x65\x2a\x2c\xbe\x48\x9f\x46\x12\x0b\xf3\x98\xa1\x50\x5c\x0f\x3d\xa9\xc3\xe2\xd3\xd5\xb3\xc8\x4c\x18\xc3\x7a\x13\xde\xfd\x7f\x54\x7c\x92\x36\x7b\x56\xab\x77\x2f\xe7\x34\x13\xcd\x02\xbf\x4c\x99\xc0\x2f\x07\x76\xff\x0c\x00\x00\xff\xff\xb4\xab\x24\xaf\xc1\x13\x00\x00")

func stateHtmlTmplBytes() ([]byte, error) {
	return bindataRead(
		_stateHtmlTmpl,
		"state.html.tmpl",
	)
}

func stateHtmlTmpl() (*asset, error) {
	bytes, err := stateHtmlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "state.html.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x88, 0xdd, 0x21, 0xbc, 0xdb, 0xd3, 0x99, 0xd5, 0x65, 0xfb, 0xc2, 0x5f, 0x7, 0xff, 0x31, 0xbb, 0xf, 0x21, 0xf0, 0x97, 0x37, 0x98, 0xb6, 0x9d, 0x20, 0x1f, 0x99, 0xde, 0xe5, 0xd1, 0x7c, 0x98}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"state.html.tmpl": stateHtmlTmpl,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"state.html.tmpl": {stateHtmlTmpl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}