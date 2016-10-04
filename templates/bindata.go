// Code generated by go-bindata.
// sources:
// golang.funcs.tmpl
// golang.header.tmpl
// postgres.tmpl
// DO NOT EDIT!

package templates

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

var _golangFuncsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x59\x5f\x73\x9b\xba\x12\x7f\x86\x4f\xb1\x97\xe9\x74\x20\xd7\x26\xed\x6b\xef\x70\xef\xb8\x89\x9b\x9b\x33\xae\xf3\xc7\xce\x9c\x87\x4c\xe6\x44\x35\xc2\xa5\xc1\x82\x08\xd1\x24\x87\xf0\xdd\xcf\x48\xe2\x3f\x01\x83\xed\xf6\x25\x13\xa4\xdf\xae\x56\xbb\xbf\x5d\xad\xe4\x38\x3e\x3e\x82\xb3\x8b\xd9\x64\x7e\x06\x5f\x6e\xe6\x27\xcb\xf3\x8b\x39\x2c\xa7\x5f\x2f\x67\x93\xe5\x74\x01\x47\xc7\x49\xa2\xaa\x71\x3c\x06\x1b\x3b\x2e\xc1\xa0\x05\x88\xa2\x4d\xa8\xc1\x38\x49\xc4\x38\x45\x64\x8d\xe1\x9d\x3b\x82\x77\x88\xae\xe1\x93\x05\xa6\x98\x53\xe2\x18\x5c\x07\xde\xb9\x90\x24\x23\x88\x63\x4c\xec\x24\x89\x63\x81\x32\xe7\x68\x83\x21\x49\x20\xfb\x5e\xbe\x04\xfc\x5b\x28\xc4\xc4\xce\x95\x67\xff\x57\x2c\x40\x74\xbd\xc3\xfa\x42\x57\xc3\x82\x7e\x2b\x12\xd7\x1b\xef\xb5\x2a\x71\xbd\x7e\x2b\x39\x11\x59\x8d\xbf\x63\x64\x63\xaa\x71\xf3\x94\x95\x4f\x42\x06\x21\xdb\x30\xeb\x3e\x8e\xc1\x5c\x5c\xcd\x20\x49\xee\x55\x25\xa0\x38\x40\x14\xdb\x23\xa0\x38\xf0\xd0\x0a\x8f\x00\x53\xca\x2d\xf1\xbf\xfd\x30\xd3\x59\x9d\x0b\x1a\xaa\xe2\x3a\x62\xf2\x5f\x16\x10\xd7\x83\x58\x55\x14\xfe\x69\xc1\x9f\x14\x05\x53\x4a\x75\x4c\xa9\xa1\x2a\x0a\xc5\x2c\xa2\x44\x55\x12\x55\xb1\xb1\x83\x69\xa6\x5a\x37\x54\x85\x6b\xf5\xfc\xf5\x82\x6d\x98\xd0\x2a\x76\xc7\xf0\x26\xf0\x10\xcb\x83\x62\x4e\xe8\x3a\x84\x24\x31\xd4\x74\xeb\xe9\x06\x8f\x8f\x60\x31\x9d\x4d\x4f\x96\xb0\x58\x4e\x96\xd3\xaf\xd3\xf9\x52\x52\x2b\x8e\xf3\xad\x87\xd8\xc3\x2b\x26\x76\xcd\xbd\x00\xba\xff\xed\x07\x1c\x9d\x7e\x36\xaa\x0b\x49\xd8\x98\x43\x34\x30\x39\xba\x0c\x5f\x3e\x6f\x85\xe7\x41\xa9\x7a\xbe\x02\xe4\x66\x9f\x61\x26\x1c\xce\x68\xb4\x62\x92\x39\xe6\x97\x88\xac\x16\x91\xe3\xb8\xcf\x90\x24\xba\x90\x2f\x96\xca\x32\x23\x77\x02\xe8\x14\x87\x91\xc7\xe0\xa8\xa2\x48\x06\x0a\x53\xea\x53\x03\xe2\x9a\x92\x0a\x01\x84\xbd\x4a\xaa\xc4\x82\xf7\x55\x7b\x12\x35\x8d\x62\x46\x05\xf3\x2a\xc2\xf4\xe5\xda\x7f\x7a\xd6\xe3\xb8\x35\x34\xa9\x86\xc5\x0a\x91\xd4\xbe\x82\x20\x96\x05\xe1\xa3\x67\x4e\x29\x9d\xfb\xd7\xfe\x53\x28\xa8\x22\x59\xc1\x99\x33\xe2\x7f\x04\x3d\x9a\x84\x2a\xa3\xaa\xb4\x12\x5b\x10\x93\x72\x39\xa9\xa5\x91\x04\x05\x47\x26\xb3\x59\x9d\x27\x4d\xa2\x8c\x91\xe7\xf5\x25\x0b\xf2\xbc\xa1\x84\xa9\x89\x74\x93\xa6\x00\x17\xc4\x09\xbc\x88\x22\xcf\xfd\x1b\x1f\x86\x42\xb7\x77\xfb\x91\xa8\xc6\x94\x85\x30\x5c\x7f\x9f\x45\xa4\x83\x2f\xbf\x38\xd8\x97\x93\xb3\xe9\x69\x8f\x70\x07\x68\x8d\xed\xbe\x01\x17\xe0\xa1\x21\x6f\x08\x75\x07\xbd\x0c\xe7\xbb\xba\xe4\xdf\x83\x63\xef\x3a\xb9\xa7\xbb\xa9\x90\x9f\x26\x7c\x2d\x00\x80\x15\xf3\x1f\x30\x81\x90\x51\x97\xac\x47\xe0\xb9\x1b\x97\x81\x4b\x58\x17\x67\xa4\x8c\x1f\xb1\x5c\x6c\x10\x8b\xc4\xba\xae\x93\x2d\x6d\x59\xa0\x69\x10\x8b\xd1\x92\x45\x16\x68\x1f\x34\x31\x98\xa8\x3d\x88\x57\x75\x41\xeb\x91\x92\xf7\x10\xe9\x32\xe9\x8e\xb7\xd2\x53\xd3\x9a\x14\x15\x32\x1e\xce\xeb\x1f\xdf\x89\xf4\x1f\x97\x2e\x9c\x64\x81\xb3\x61\xe6\x22\xa0\x2e\x61\x29\xf6\x56\xe0\xc6\x1f\xef\x4c\xee\x5a\x11\xf3\x0b\x22\xd3\x84\xeb\xad\x71\x3f\x57\xd5\x9e\x06\x27\x17\x37\xf3\xc6\xb1\x58\xe6\xff\xca\x8f\xc8\xf6\x63\x51\xa0\x7a\x33\xbe\x13\xdd\xd0\xfd\x1d\x85\xbd\x35\xd7\xb0\x2d\x29\x54\x5e\x9f\xbb\xe2\x84\x7f\x1f\xea\xb4\x75\x09\xdb\xb3\x3a\xf6\x3c\x47\xf9\x09\xfa\xbe\x7e\x84\x36\x69\xf8\x61\xc7\x1a\x59\xb8\xab\x70\x2a\x9f\xfa\x3f\x0a\x0f\xe5\xaa\x6f\xbe\xef\xd5\x7c\xa5\x88\xd8\xe4\xcd\xa4\xe8\x26\xfb\x84\x67\xb7\x13\xc4\x41\x5e\x28\x3b\xd7\xb2\x5b\x84\x09\xf0\x5f\xee\xb9\xb6\xb4\x39\x9d\xce\xa6\xcb\x69\x47\x3b\x69\x63\x0f\x33\xbc\x35\x6f\x24\xac\x37\xbd\x9b\xf0\x16\x86\x57\x80\xdc\xec\x53\x31\xb0\x77\xe0\xa4\x5e\xfb\xad\xc8\xf5\x69\x24\xf3\xb8\xe6\x5c\x9f\x3e\xe3\xd5\x7e\xc1\x6b\x90\xbb\xca\x20\x8a\x43\x93\xf7\x92\x13\xc7\xc1\x2b\x86\x6d\x7d\x17\x95\xfd\x89\xd1\xaf\x87\x4c\xc3\xd3\xa7\x87\x2c\xa0\x43\x59\xd2\xaf\x87\x6c\x80\x2b\x6c\x39\x54\x1b\x29\x3d\xb7\x43\x69\x3c\x34\x69\xde\xaa\x86\x25\xc2\x0c\xe4\x4b\x47\x6d\x4d\x95\xb6\xf1\xe4\x7c\xbe\x98\x5e\x77\xdd\x47\x5d\x12\x62\xba\xfd\xe0\x95\xb0\xde\xd4\x68\xc2\x5b\x68\x51\x01\x8a\x33\x92\x62\x54\x2b\x20\xbf\xe4\x06\x9a\x2e\x8c\x22\xe6\x8f\x5d\xe2\xb2\xdc\xcc\x2d\xbd\x61\xc7\x3d\x35\x6d\x74\x89\xcf\xc0\xbc\x16\xd1\xf9\xfc\x32\xf8\xdc\x3d\x17\x76\x85\x5b\x6e\xb0\x7b\x5c\x52\x04\x47\xbc\x10\x0f\xa7\x7d\x61\xda\x60\x3b\x82\x87\x4a\xa1\x9c\xa1\x90\x49\x6d\xe7\xdd\xc4\xef\xdc\x15\x3f\xb3\xeb\xcf\x17\x9f\x5f\xf8\x67\xc9\xfb\x7a\xf0\x60\x54\xf2\xa2\x91\x25\x75\x2e\x12\x7f\x2c\x17\xe8\x9b\x14\xb9\xc0\xd0\xf4\x78\x53\xb0\x3b\x51\xea\x22\xbb\xa7\xcc\xaf\xc8\x8d\xbf\xb2\xea\x36\x88\x4b\x69\x38\x2b\x61\xee\x68\x13\x9b\xd6\x65\x6f\x8d\xfc\x8a\x35\xc7\xd8\x0e\xe7\xfe\x93\xa0\x37\xf1\x9f\x38\xe5\xe6\xfe\x93\x6e\x34\x5e\x26\xe5\xe3\xa6\x39\x89\x98\x1f\x42\xfa\xa8\x99\x3f\xd9\x7e\xb2\x78\xc4\xcc\x73\xe2\xb2\x2d\xef\xa7\xc7\x47\x70\x73\x79\x3a\xe9\xec\xd2\xa2\xc0\x46\x3d\xba\x34\x09\xeb\xcd\xa2\x26\xbc\x85\x3b\x15\x20\x37\xfb\x46\x0c\xec\xd7\xa5\x8d\x40\xaa\x0d\xc1\x34\xcd\x8a\x2a\xa9\x7d\xa7\xaa\x5c\xb2\xb4\xc4\xad\x71\x76\x29\x1f\x5e\x7b\xb9\x54\x6b\xf9\xe5\x5c\x34\x4d\xf3\xcd\x4a\x9b\xbe\x01\x54\xca\x52\xf6\x02\xd0\x5a\x9c\xe4\x6b\x40\x09\xd2\x56\x78\xd3\xcd\xb4\xd4\xde\xcc\xae\x7d\x8d\xe8\xae\xba\x87\xd9\xe1\xa1\x8b\x70\xca\x80\xfe\x45\xb8\x2e\x30\x34\x7d\x86\x14\xe1\x16\x91\xdf\x98\x52\x3b\xe5\x0e\x0f\x57\x4b\x69\xae\x50\x6d\x60\x21\x7e\x6b\x41\xf1\xeb\xd0\x4f\x44\xe1\x31\x84\xdb\x3b\xf9\xfc\x26\x07\x7e\x22\x2f\xc2\x7c\xd0\x25\x0c\x53\x07\xad\x70\x9c\xa8\x4a\x88\x31\xe1\xfc\xdc\xa0\xe0\x56\xa2\xf9\x3c\x9f\x41\xb6\x2d\x44\xf8\x2c\x5f\x42\x27\xbc\x34\x67\x0f\x7a\x72\xaa\xa4\x4a\x5c\xea\x79\x23\xe1\xda\xcf\x23\xf0\x1f\xb8\x18\x57\x7e\xcb\xc5\xee\xfe\x03\xfe\x03\x9f\x57\xa4\x15\xb7\xae\xfd\x7c\x07\x96\xd4\xa2\x2a\xbc\xa7\x50\x0a\x30\x58\xe2\xd5\x4c\x42\x8d\x3c\x27\x1e\x43\xb0\x00\x05\x01\x26\xb6\xfe\x18\x8e\x40\xd8\xf3\x6f\xd0\xc0\x82\xff\x69\x86\x9a\xe9\x2e\x40\xf2\x3b\xb5\x55\x74\x2e\x07\x38\xa4\x32\xb7\xe8\x1a\x27\xc8\x89\xef\x45\x1b\x02\x49\xa2\x8d\xca\xe7\x55\xa1\x87\x8b\x38\x3e\xe5\xe1\x97\xe1\x12\xc5\x40\x28\xcd\xb8\xc6\xfd\x92\x6b\x95\x83\xe6\x4a\xe8\xe5\x87\xa1\x6e\x64\x92\xa6\x44\x18\xf2\x0d\x90\x7b\x24\x64\x1b\x26\xfc\x2c\x82\x12\x9a\xd7\xe9\xef\x68\xf7\x71\x6c\x2e\xae\x66\x49\x72\x3f\xca\xbd\xa7\xbd\xbe\xbe\x5e\x2d\x5e\x5f\x5f\xb5\x62\x2c\x93\xfb\xc3\x77\x89\x70\xa9\x36\x02\xcd\x18\xc1\x47\xe3\xf7\xff\xf0\xc7\x13\x81\x6b\xaf\xc5\xae\x9a\x62\xe5\xfb\xdf\x38\x49\x0c\xe5\xad\x1f\x0c\xf3\x94\x2a\xbd\x64\xff\x13\x00\x00\xff\xff\xe5\x99\xa5\x36\x80\x1e\x00\x00")

func golangFuncsTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangFuncsTmpl,
		"golang.funcs.tmpl",
	)
}

func golangFuncsTmpl() (*asset, error) {
	bytes, err := golangFuncsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.funcs.tmpl", size: 7808, mode: os.FileMode(420), modTime: time.Unix(1475618841, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _golangHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x57\xdf\x8f\xe3\xb6\x11\x7e\x16\xff\x8a\x89\xd1\x3b\x48\xbb\x5e\xfa\x9e\x9d\x6c\x8b\xf8\xbc\x0d\x0a\x5c\x77\x0f\x1b\x1f\x8a\x22\x08\x0e\x94\x34\xf6\xb2\xa6\x48\x1f\x45\xd9\x72\x5c\xff\xef\xc5\x90\x94\x2c\xaf\x9d\xf4\x9a\xc6\x0f\x2b\x89\x9c\x1f\xdf\x7c\xfc\x38\xe4\x4e\x26\xf0\xfd\xa7\xc5\xd3\x0f\x0f\x8f\x0f\xcf\xdf\x2f\x1e\xe6\x30\xfb\x27\xac\xa4\x7b\x69\x72\x5e\x98\x6a\x52\x6f\x44\x81\x95\xd1\x6b\xdc\xaf\xcc\xa4\xcc\x5b\x36\x99\xc0\xfc\x09\x1e\x9f\x16\xf0\x30\xff\xdb\x82\x33\xb6\x11\xc5\x5a\xac\x10\x0e\x07\xe0\x1f\xe3\xfb\xf1\xc8\x98\xac\x36\xc6\x3a\x48\x59\x32\x2a\x85\x13\xb9\xa8\x71\x52\x7f\x51\x23\x96\x8c\x96\x95\xa3\x47\xed\xac\xd4\xab\xda\xbf\xee\x75\x41\x4f\x27\x2b\x1c\x31\x96\x8c\x06\x18\xfe\x55\x19\x69\x8d\x26\xe7\x76\xc4\x32\xc6\xb6\xc2\xc2\x3f\xac\xd8\x3c\x58\x0b\xf7\xb0\x6c\x74\x91\xa2\xb5\x80\xd6\x1a\x9b\x85\x07\x1c\x2c\xba\xc6\x6a\xfa\x3a\x7a\x87\x0f\x66\xb5\x42\x1b\xac\x97\xc6\x56\xc2\x41\xc8\x3f\x06\x61\x57\x35\x70\xce\xa5\x76\x68\x97\xa2\xc0\xc3\x31\xf3\x3e\x8f\x66\x07\xf7\x40\x98\xf8\xa3\xd9\x31\x46\xce\xa0\xcc\x0a\x87\xf9\x0e\x2c\x91\xcb\x2e\xfc\x37\xf7\xa0\xa5\x82\xb7\x6f\x69\xb6\xfb\x3a\xb0\x24\x09\xf3\xe9\x28\xa0\x2b\x94\xa9\xa5\x5e\x41\xed\x84\xc3\x0a\xb5\x9b\xc2\x9b\xed\x68\x4c\x4e\x19\x4b\x8e\xec\xc8\xd8\xe1\x00\x56\xe8\x15\x02\xff\xd1\xd9\xa6\x70\x35\x91\xea\xf6\x9b\x40\xf4\xa3\xa8\x88\x65\x2a\xa1\x29\x1c\x1c\xd8\xe1\x70\xd7\xd9\xff\x55\xa2\x2a\xbd\x79\x32\x34\xa5\xf7\x05\xf9\x77\xef\x62\x45\x36\xe4\x88\xba\xa4\x57\x9f\xf5\x0e\xe4\x12\xf8\xa7\x8d\x5f\x33\x85\xc1\x04\xfe\x14\x13\x4d\xef\xfb\x80\x3d\x98\x6e\xee\x78\xf4\x5e\x08\x3d\x8f\x54\x79\x61\x54\x53\x69\xf2\x49\xb3\xc8\x38\x4b\xb6\x42\x35\xf4\x3d\x60\x9c\x25\x1b\x2b\xb7\xc2\xe1\x59\xc4\x34\x63\xc7\x01\x15\x3d\xac\x53\x8d\x64\xde\x6c\x4e\xf0\x36\x56\x6a\xb7\x84\xd1\x9b\xfa\x4d\x3d\x82\xb4\x10\x15\xaa\xd2\xec\x74\x17\x34\xbb\x52\x40\xef\x7f\x4e\x53\x58\xef\x33\x3c\x03\x42\xd3\xed\xd0\x36\xbb\xca\xc4\x01\xa2\x0c\x5f\x65\x49\xb7\x19\xc4\xf0\x69\xf3\x7a\x32\x83\x2b\x9c\x9d\x42\x8d\x28\xeb\x7b\x6f\x01\xc7\xe3\xe8\xb7\xe2\x5c\xa1\xf9\x0c\x52\x07\x3e\x6d\x7e\x13\xce\xf5\x85\x81\xc3\x49\x3c\x77\xc7\xeb\xef\xd4\x07\x3c\xcb\xf3\xd9\x49\xaa\xc9\x0d\xed\x64\x3e\x9f\x31\x96\x28\xdb\xc0\x8d\xb2\x0d\x69\xcf\xe7\x7f\xda\xa0\x4e\x6b\xd3\xd8\x02\xfb\xfd\xa9\x6c\xf3\xb9\x96\xbf\x78\x61\x65\x90\xde\xcc\x67\xe3\xc1\xe6\x2b\x73\xff\x45\x6b\xef\xe3\xbe\x37\x5a\x63\xe1\x52\xaa\x6f\x2e\x85\xc2\xc2\xc1\xbf\x4f\xb2\xf8\x32\x82\xe3\x71\x0c\x21\x45\xe6\xf7\xee\xab\xad\x1a\xf9\xd1\x52\x8d\xbb\x36\x93\x76\x1b\x93\x75\xb3\x6f\xe7\x33\xb2\x9d\xcf\xa6\x50\xe6\x63\x96\x50\x25\x53\xd0\xb8\xfb\xf0\xfc\x29\x25\x48\x1d\xe8\x6c\xcc\x92\xe3\x98\xa2\xf5\x35\xa6\x65\x0e\x37\xf3\x59\x06\xef\x95\xa9\x69\x7d\x5e\xb5\x93\x32\xe7\xca\x36\xfc\x63\x63\x57\x98\x66\x7d\xca\x0e\x4b\x99\xf3\xf9\x8c\x47\xdf\xec\x32\xea\xc6\xe2\x46\x58\x4c\x6b\x57\x75\x3d\x2e\x83\x34\x8e\x96\x10\xd8\xff\xd1\x55\x6e\x0c\x16\x37\x8a\xf6\xa9\xef\x8a\x59\xe0\xf1\x84\x03\xb5\xb3\xfb\x9e\xdc\x88\xea\x07\x74\x3e\xf2\x7f\xa5\xce\xff\x41\x6b\x89\xb6\x6e\xc2\x47\x8c\xb9\x43\x4e\x38\x74\x81\x9f\x03\x96\xd4\xdb\x64\xf0\x6b\x9c\x29\xb3\x22\xff\x61\x75\xa1\x83\x73\xce\x87\x42\xf7\x15\x4c\x26\xb0\x78\x9a\x3f\x4d\xc1\xa2\x2e\xd1\x82\xcf\xf0\x62\x54\x89\xb6\xbe\xd2\xb6\x87\x8d\xba\xfe\xa2\xa6\x30\x82\x5b\xf0\x89\x6e\x61\xe4\x93\xdc\xfb\x36\x4d\x6f\x5d\x9f\xf6\xf2\x5e\xb4\x17\xf2\x5e\xb4\xb4\x90\x84\xf9\xb2\x06\x2f\x72\x52\xf2\xa2\x1d\x2a\xd9\xb5\x43\xb2\xe7\x33\x3e\xc3\x95\xd4\x6d\xfa\x7f\xa8\x74\xd1\x92\xed\xa2\x9d\x82\x6b\x49\xa5\x65\x1e\xf5\xfa\x9a\x5e\xd7\xc2\xcd\xa2\xcd\xe0\xbd\xa9\x2a\xe9\x2e\x35\xf9\x4a\x83\x8e\x0a\xe4\x9d\x71\x76\x19\xe7\xd9\x28\x95\x8b\x62\xfd\x95\x91\x4e\xe6\x57\x62\xfd\x61\x82\xee\x5c\x7a\xbb\x60\x70\x0f\xae\xe5\x65\xce\x87\x79\x7e\x97\xbc\x5d\x80\xd1\xf6\xd8\xb2\x41\xa6\xab\x74\xff\x4f\x6a\x0e\x28\x87\x2e\xbd\x6d\xd6\x4b\xd1\x6f\x9f\x81\x1a\xc9\x76\x40\x11\x63\x89\xc6\xd6\x8d\x89\xd4\x2d\xdc\x78\x6b\x96\xac\x71\x0f\xf1\xd7\x1d\xd1\x16\x97\x75\x1c\x92\xda\xf5\xe1\xa9\x61\x9f\x82\x57\x4d\xe7\xb5\xd7\x05\xff\x7b\xe3\x30\xa8\x3e\xfc\xfa\x3e\x9f\x14\x62\x23\x0a\xe9\xf6\x3e\x54\x62\x8d\x71\xde\x20\x66\x97\xba\xc4\x96\xbe\x2b\xb1\xf9\x29\xe4\xff\x39\x22\x3b\x32\xba\x70\x86\xb6\x1a\x0f\xae\x1a\x84\x06\xa9\xa5\x93\x42\xc9\x5f\xb0\x24\x48\x3c\xd0\xda\xb7\xdf\x3e\xf5\x18\x86\xa9\x33\x7f\xd6\x10\x70\x45\x9b\xec\xad\xb2\xcd\x21\x6e\x8a\xf0\x0b\xad\xbc\x73\x99\xf6\xce\x34\xea\x51\x4e\x03\xcc\x35\xa6\x17\x58\x7d\xa3\x67\x89\xe2\x54\x1e\x27\x92\x81\x52\xf8\xcf\x7e\xd8\xb3\x3e\x18\x8e\xc2\x51\xb1\x50\xa9\x6b\xb4\x2e\x3e\x6a\x40\x10\x4b\x87\x16\x84\x1b\x83\xd4\x85\xf5\x57\xc5\x1a\x14\x57\xa8\x63\xc9\xa9\xf2\x35\x65\xd1\x27\xc5\x31\x08\x17\xd7\xd5\x8b\x46\x53\xa5\x22\x00\x62\x89\xe8\x91\x21\x4b\xb0\x83\x23\x1c\x7d\xc4\x09\xcd\x12\xdd\x4d\x60\xc4\x65\xb1\x32\x5b\x8c\x0f\xc2\xb5\xb4\xa6\x02\x49\x58\x64\xed\xc6\x50\xe2\x39\xb8\x57\xd8\x82\x5f\x8a\x43\x5c\x21\x79\x8f\x26\xe2\x0b\xcf\x3e\xbd\x7f\x19\x62\x93\x0a\x26\x13\x10\x5b\x23\x4b\xa8\xb0\x32\x76\x0f\x0a\xc5\xba\x1e\x14\xf3\xeb\x36\xfd\xee\xeb\x70\xe1\x56\x16\x2e\x3d\x57\x48\xb8\xce\xf7\x63\xdf\xc1\xbb\xc1\xd6\x0f\xcd\xd5\x99\xcf\xde\x93\x98\x55\xa8\x53\xc5\xbd\x36\x32\xb8\xeb\xfd\x58\x52\x34\xd6\xa2\x76\x41\x67\x71\xb5\x97\xc6\x82\xa4\x91\x77\xdf\x82\x84\xef\xa0\x0b\xf4\x2d\xc8\xdb\x5b\x9f\xa6\xf3\xba\x87\xf8\x16\x19\xf0\x98\xba\xb9\x3e\xa2\x77\x49\x36\x42\xcb\x22\x1d\x85\x5d\x64\x1a\x07\x66\x09\xbb\x17\x51\xac\x39\x58\xea\x85\x85\xd1\xa5\x74\xd2\xe8\xfa\x2f\xa3\x8c\x25\x24\xd2\x41\x38\xee\x37\xfa\x9f\x63\x99\x49\x61\xb4\x93\xba\xc1\x68\x37\x99\xc0\xee\xc5\x08\xc8\xcd\x1e\x64\x0d\xee\x45\xd6\x50\x37\xb9\x53\xc8\x41\xf1\x28\x8b\x5c\x99\x5d\x0d\x62\x27\xf6\x7d\x50\xbf\x5e\x42\x97\x21\xc4\xb0\x16\x0e\x79\xe3\xce\xec\xc2\xc2\xed\xa4\x52\x90\x23\x54\x0d\xfd\x4f\x54\x82\x33\x90\xe3\xa5\x3b\x10\x87\xee\x05\xc1\x67\x90\x0e\xad\xa0\xd2\xbe\x61\x89\xef\x6b\x9f\x07\xb4\x0f\x73\x30\xba\xa1\x45\xbc\x69\x9c\x20\x2e\x4a\x54\xe8\xb0\x5b\xc1\x71\xef\xb3\xc6\x3d\x4d\xfb\xff\xf5\xba\x31\x6a\x9f\xa7\x6b\xd7\xd9\x62\x0d\x53\xc7\x7b\xc1\xb9\xd2\xe8\xd6\x44\x4d\xb6\x3f\xbd\x6e\x4e\x97\xab\x78\x38\x29\x5e\x35\xfc\x83\xa1\x33\x90\x25\x25\x2e\xd1\x82\x1f\xfa\xa4\x55\x18\x64\x09\x8e\xc1\xac\xbd\xea\x02\xde\x9f\xd6\xb8\xff\xd9\xeb\xd5\xac\xfd\xf2\xf5\x25\xa2\x47\xcf\x4f\x4d\x21\x2a\x86\x86\xd1\x2f\xf9\xed\x2d\x3b\x9d\x68\xf1\x78\xf2\xea\x0e\xe7\x4a\xbc\x87\x28\x3a\x71\x3e\x86\xd3\xac\x4d\xfd\xd7\x33\xe6\x52\x97\x54\x4d\xf6\x3b\x2e\x26\x48\xdd\xcf\xd7\x4e\xb6\x44\xe9\xd4\xdf\xb0\xa8\xc3\xae\x71\x4f\xfd\x75\x8d\xfb\xae\x95\x5e\xc3\x3f\xac\xbd\xeb\x63\x7d\x41\xe7\xf5\x5c\x2c\x43\x7f\xc7\x1c\x76\xa2\xaf\x21\xde\x67\xb8\xbb\x23\xe3\xd0\x34\x14\xef\xb6\x7a\x76\x99\x26\xde\xdc\xbf\x32\x78\x17\xf2\x9d\x8f\xc4\xfe\x13\x00\x00\xff\xff\xbe\x64\x69\x58\x92\x11\x00\x00")

func golangHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_golangHeaderTmpl,
		"golang.header.tmpl",
	)
}

func golangHeaderTmpl() (*asset, error) {
	bytes, err := golangHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "golang.header.tmpl", size: 4498, mode: os.FileMode(420), modTime: time.Unix(1475619075, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _postgresTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x56\x4d\x8f\xda\x38\x18\x3e\xc3\xaf\x78\x55\x71\x60\xd0\x90\xde\x77\x55\x8d\x32\xe0\xe9\xa6\x0d\x0e\x93\x98\xad\xe6\x98\x12\x53\xbc\x4a\x1c\x36\x18\x8d\x90\xf1\x7f\x5f\xd9\xce\x27\x4d\xd8\xe9\x6a\x0e\x7b\x73\xec\xc7\xcf\xfb\xbc\x9f\x8e\x94\x1f\x67\xb0\x0e\x22\xf2\x39\x44\xd1\xb3\x0f\x11\x71\x09\x5a\x21\x4c\x80\xa0\xd5\xda\x77\x09\x8a\x00\x66\x1f\x95\x1a\x8f\xa5\x9c\x43\x42\x77\x8c\x53\xf8\x70\xdc\xee\x69\x16\x7f\x80\xb9\x52\x66\xbf\x88\xf9\x0f\x0a\x0e\x89\xbf\xa7\xf4\x08\x4a\x8d\x17\x21\x72\x09\x02\xe2\x3e\xfa\x08\xa4\x04\x07\xc7\x19\x05\xa5\x60\xda\xc2\x4f\xd8\x3d\x4c\xb6\x79\x0a\xbf\x7d\x02\x67\x91\xa7\xa7\x8c\x9b\xcb\x52\xb2\x1d\x4c\x98\x52\xf7\x52\x52\x9e\x28\x35\x92\x12\x5e\x99\xd8\x5b\xb4\x36\x3a\x6a\x73\xea\x35\x39\x1f\x68\x75\x32\x07\xb6\x03\x07\xe7\x02\x9f\xd2\x54\x03\x70\x40\x00\x6f\x7c\x1f\xf4\x19\xe5\x49\x17\x18\xd2\x1d\x2d\x28\xdf\x1a\xae\x10\x3d\xa1\x10\xe1\x05\x8a\x0c\x6d\x7d\x66\x7d\x03\xa5\xa6\xdd\x6d\x2b\x1b\x94\xba\x83\x00\xc3\x12\xf9\x88\xa0\xab\x9b\x01\x5f\xd2\x94\x0a\x7d\xb9\x6d\x7f\x78\x6d\x54\xad\x0b\x96\xc5\xc5\xf9\x2b\x3d\xeb\x90\xdc\x8f\xd6\xa1\xb7\x72\xc3\x17\xf8\x8a\x5e\x60\x3a\x10\xc3\xce\x9d\x26\x8a\x50\x86\x51\x4a\x0b\x35\x2b\x6d\x4e\xa9\xbb\x6b\xd3\x65\x26\x37\x9c\xfd\x7d\xa2\xd6\xf2\x06\x7b\xcf\x1b\x34\x68\xf4\x97\x4d\x29\x35\xbe\xfb\x7d\x5c\xef\xb7\x56\x9d\x1a\x4b\x72\xc1\xe3\x8c\x56\x45\x06\x75\x06\x1c\xfd\x51\xc7\xbd\xe3\x40\x87\x20\xa5\x3b\xf1\x57\xce\xb8\x65\xf0\xd1\x13\x81\x2f\x81\x87\x6d\x76\xd8\x8f\xbd\xa8\x19\x75\xea\xa4\x04\x41\xb3\x43\x1a\x8b\xb6\x6d\xc7\xa7\x3b\xa1\x11\x9f\x86\x00\x86\xe9\xa6\x8e\x6d\xce\x13\x26\x58\xce\x9b\x7e\xd1\x09\xb6\x1e\x2c\xb2\x43\x5d\xd0\x7d\xf4\x35\xaa\x56\xd2\x78\xaf\x77\x83\x03\x2d\x62\x91\x17\xfa\xe4\xe1\xa7\x9a\xea\xd8\x29\x43\xf6\x26\x6b\x76\xd1\x6f\xb3\x3c\x6b\x5b\x7e\x23\x5d\x6f\xac\xba\x3a\xbd\xb7\x08\xf4\x1a\x65\x1e\xd6\x85\xd9\xc2\x1e\x69\x4a\xb7\xa2\x0d\xf5\x78\x4f\xa1\x0f\xe4\xea\x75\x4f\x0b\x7a\x3d\xd7\xca\x72\xe7\x89\xad\xf7\x4a\x1f\xdb\x01\xcf\x05\x4c\x18\x28\xf5\xed\x0f\x14\x9a\xbe\xa7\xe9\xd1\x54\x94\x8b\x97\x50\x17\x76\x47\x60\xab\x1c\x2c\xe9\xad\x61\xd0\x9d\xba\xc6\xb5\xf9\xeb\x3e\x16\x75\x57\x4c\x84\xa9\x61\x2d\xac\xaa\xe6\x2a\xa2\x29\xe5\xe0\x7c\xdb\xc7\xa2\x56\xdc\x72\x47\x93\x98\x5b\x06\x60\x24\x9a\x1e\x06\xdb\xc4\x8d\xf0\xd2\x80\x6d\x3b\x7b\xad\xd5\xd4\x56\xac\xf6\xb9\xb2\xd1\xe0\x67\xbf\xe2\x95\x75\x28\x42\x3e\x5a\x10\xe8\x49\x68\xe9\xb5\xa3\x14\x3c\x85\xc1\x0a\xda\x03\xc1\x8e\xf2\x72\x72\xe9\xc2\xf8\x92\x33\xfb\x92\x8c\xba\x75\xd4\x4c\x04\xa7\xbe\xf6\xd3\x73\xb0\xa8\xf2\xd3\xc7\x50\xd6\xc7\x35\xa8\x97\x28\x28\x12\x5a\x3c\x9e\x2d\x4b\x10\x2e\x51\x08\x8f\x2f\x43\xad\x52\x82\x9b\xb9\x06\x6e\xb4\x18\x20\xf6\x59\xc6\x84\xa5\xf5\xbd\x95\x47\xe0\xa1\x0b\xbc\x31\x89\x4e\xbc\x1b\x69\xb3\x33\x9d\xdd\xfd\xdf\x83\xfa\x6f\x3d\xf1\x3d\x3e\xb6\xda\xb6\x17\x99\x98\x47\xd8\x82\xca\x77\x7a\xc0\xe9\x7e\xc1\xef\xa3\x97\xf1\x23\x2d\xca\x1c\x78\x38\x42\x21\x01\x0f\x93\x60\x50\x85\xfd\x29\x32\xb1\x1c\x0d\xbd\xc0\xcd\xaf\xd3\x70\x17\x5f\x3f\xc6\xf0\xa7\xeb\x6f\x50\x34\x35\xae\xfd\x27\xd6\x87\xd6\xcb\x5e\xfb\x6e\xfe\xa0\xc8\x26\xc4\x1e\xfe\x0c\xb3\xe1\x30\x9c\x0e\x49\x5c\x25\x63\xb3\x5e\xba\xe5\x4f\x53\xfd\x20\x47\x88\xc0\xe5\x72\x79\x8e\x2e\x97\xcb\xbb\x24\xe5\x86\xb0\x7f\x02\x00\x00\xff\xff\x5f\xa3\x67\x2b\x04\x0b\x00\x00")

func postgresTmplBytes() ([]byte, error) {
	return bindataRead(
		_postgresTmpl,
		"postgres.tmpl",
	)
}

func postgresTmpl() (*asset, error) {
	bytes, err := postgresTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "postgres.tmpl", size: 2820, mode: os.FileMode(420), modTime: time.Unix(1475618810, 0)}
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
	"golang.funcs.tmpl": golangFuncsTmpl,
	"golang.header.tmpl": golangHeaderTmpl,
	"postgres.tmpl": postgresTmpl,
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
	"golang.funcs.tmpl": &bintree{golangFuncsTmpl, map[string]*bintree{}},
	"golang.header.tmpl": &bintree{golangHeaderTmpl, map[string]*bintree{}},
	"postgres.tmpl": &bintree{postgresTmpl, map[string]*bintree{}},
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

