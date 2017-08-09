// Code generated by go-bindata.
// sources:
// bootstrap/README.md
// bootstrap/amazon_k8s_ubuntu_16.04_master.sh
// bootstrap/amazon_k8s_ubuntu_16.04_node.sh
// bootstrap/digitalocean_k8s_ubuntu_16.04_master.sh
// bootstrap/digitalocean_k8s_ubuntu_16.04_node.sh
// bootstrap/inject.go
// bootstrap/vpn/meshbirdMaster.sh
// bootstrap/vpn/meshbirdNode.sh
// bootstrap/vpn/openvpnMaster.sh
// bootstrap/vpn/openvpnNode.sh
// DO NOT EDIT!

package bootstrap

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

var _bootstrapReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x91\xb1\x6e\xdc\x30\x10\x44\x7b\x7e\xc5\x00\x57\xa4\x31\x94\x4f\x08\x90\x22\xc0\x15\x49\x65\x38\xad\x28\x71\x64\xae\x8f\xe2\x0a\xe4\xd2\x82\x10\xe4\xdf\x03\xea\x1c\x23\x81\xeb\xdd\xd9\x79\x3b\x73\xc1\x57\x55\xab\x56\xfc\xe6\xdc\x63\x64\x25\x7c\x21\x2c\x12\xd3\xdf\x01\xea\x5c\x64\xb3\x0a\x8b\xde\x50\xa3\x6c\xd8\xc5\xe2\xb9\x14\xb8\xf8\x96\x0c\xe3\xad\x4d\x32\x6b\xc9\x23\xb6\xa2\x8b\x24\xd6\xc1\xb9\x6f\x64\xc2\x52\x48\x98\xc2\x87\x80\x43\x5b\x81\xee\xf9\x01\x5a\xb0\x6a\x90\xe5\xe8\x67\xba\xab\xc1\xe7\x03\x26\x2b\x87\x93\xe4\xdd\xb5\xf3\x70\x59\x38\x9b\xbc\x32\x1d\xd8\x3b\xc5\x4e\xb4\xae\xaa\x18\x5b\x65\x41\xf0\xe6\xc7\xee\x22\x59\x0c\x1e\x4f\xdf\x9d\xbb\x5c\x2e\xb8\x22\x93\xa1\x0f\x8c\xeb\x96\xbc\x11\xda\x0c\x9a\x09\x5d\xde\xac\x3f\x3c\xea\xdc\x0f\xed\xa8\x08\x9a\x3f\xd9\x80\x9f\x45\x8c\x98\x7c\x8d\x48\x72\x23\x7c\xff\x71\xf8\xcf\x60\xd5\xc2\x13\x02\x92\xe1\x3f\x9c\xbc\x43\xd7\xa8\x2d\x05\x5c\x11\xf4\x8b\x73\xd7\xe5\xf4\x28\xf4\x29\x1d\x98\x7d\x86\xe6\x74\xe0\x99\x06\x31\x2c\x45\xd7\x7f\x53\xf5\x39\x20\xeb\x1e\xd9\xd3\x48\x95\x0f\xa7\xb8\xab\x7a\x0e\xbd\x8a\xf1\xc9\xa7\xc6\xfa\xeb\xf7\x88\x6a\xa5\xcd\xd6\xbf\x0e\x5c\x24\x13\x73\xab\xa6\x2b\x6e\x3c\x3e\xbf\xf6\x2d\x6c\x5e\xca\x5b\xa1\xbb\xa4\x84\x89\x90\xfc\xc2\xd9\x18\x20\xd9\xf4\xde\xd4\x1d\x7e\x70\x8f\x51\xea\xfb\x9e\xc7\xac\x81\x98\xa3\xcf\xcf\x7c\x38\xc1\xa4\x76\x11\x73\xb8\x47\x3d\x11\x2f\xad\xda\x79\x7e\xf8\x13\x00\x00\xff\xff\x8c\x9a\x9e\x97\x62\x02\x00\x00")

func bootstrapReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapReadmeMd,
		"bootstrap/README.md",
	)
}

func bootstrapReadmeMd() (*asset, error) {
	bytes, err := bootstrapReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/README.md", size: 610, mode: os.FileMode(436), modTime: time.Unix(1501967388, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapAmazon_k8s_ubuntu_1604_masterSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x54\x4d\x6b\x24\x37\x10\xbd\xeb\x57\xbc\xcc\x18\x1c\xc3\xaa\x65\x3b\x21\x01\x83\x03\xd9\x8d\x0f\xce\x86\x5d\x63\x9c\xe4\x92\x8b\x5a\xaa\xe9\x56\x46\x2d\x35\x52\x69\xec\xc1\x76\x7e\x7b\x50\x77\xfb\x0b\x96\x90\xdb\xce\xa5\xa7\x4a\xaf\x5e\xbd\x92\x9e\xb4\xfe\x46\x95\x9c\x54\xeb\x82\xa2\xb0\x43\xab\x73\x2f\x32\x31\x24\x09\x63\xf1\x8f\x10\x6b\xc8\xaf\xf4\x13\x6b\xdc\xf4\x94\x09\x3b\xed\x0b\x65\xe8\x44\x70\xe1\x6f\x32\x4c\x16\x2e\x70\x04\xf7\x84\x6c\x92\x1b\xb9\xc1\x9f\x34\x01\xe8\x6e\xf4\xce\x38\xf6\x7b\x84\xc8\x28\xd9\x85\x0e\x1a\x4c\xc3\xe8\x35\xd7\xc0\xeb\xd0\x15\xdd\x11\x38\x2e\x74\x13\xcf\xdc\x44\xac\xa1\x73\x5d\xa1\x60\x62\x49\x13\xac\x27\x94\x4c\xa9\x66\xbd\x1b\xdc\x04\x77\xa9\xe6\x10\x37\x6f\x98\x63\xe7\x0c\x5c\xa8\x80\x4c\xd8\x38\x4f\xb9\xc1\xfb\x3d\x2c\x65\xd7\x05\x68\xef\x5f\x06\x58\x86\xca\x7d\x2c\xde\x8a\x35\x5a\x82\x6e\xfd\xa4\xaa\x25\xd4\x23\xd0\x8c\x54\x02\xbb\x81\xde\x41\x07\x3b\x4f\xdb\x93\xf7\xcb\xcc\x48\xa4\x3d\x6e\x63\xda\x36\xb8\xdc\x60\x1f\x0b\x02\x91\x85\x89\xc1\x3a\x76\x31\x68\x3f\x4b\x7a\x87\xdb\xe4\x98\xe0\xb8\x8a\x9b\x8e\x78\x8d\x98\x30\xe8\x2d\x41\x87\xc8\x3d\xa5\x37\xcc\x8d\x58\x8b\xb5\xb8\xf9\xfc\xf1\xe2\xd3\xf9\xea\xf2\xd3\xaf\x17\x1f\x6e\x2e\x7e\x99\xc2\x95\xb8\xfa\x7c\x7d\xf3\x92\xac\xd1\xea\x6b\x7a\x44\x98\x92\x3c\x64\x46\xcf\x3c\xe6\x33\xa5\x46\x6d\xb6\xba\xa3\xdc\x18\x1f\x8b\x6d\xba\x18\x3b\x4f\x8d\x89\x83\xd2\x23\x2b\x1b\x4d\xfd\xca\x2d\xed\x9b\x6e\xec\xf0\x80\x5c\x6c\xc4\x92\x82\xb6\x16\x52\x70\x2c\xa6\x87\x22\x9e\xb0\x2a\xc7\x92\x0c\xe5\xc6\xbb\xcc\x8d\x55\xdb\xd2\x52\x0a\xc4\x4b\x46\xe4\x1e\xd2\xe0\x90\x4c\x1f\xb1\xb2\xd4\x4e\x4a\xce\x54\x2d\x6d\x5e\x61\x5d\x54\x78\x09\xe5\x1d\x05\xa7\x3d\x06\xed\xc2\x0a\x3f\xfd\xdf\x66\x87\x42\x54\xa9\x1d\x31\xca\x68\x35\x13\xe4\xfe\x39\xe3\x42\xe6\x6a\x31\xb9\xc7\x5f\x02\x00\x72\x34\x9a\x97\xff\xd4\x72\xb5\x57\x5e\x42\x1b\xcd\x96\x52\xe3\xe2\x12\x57\x0e\x4e\x3a\xe4\x31\x26\x96\xd3\x5e\x2e\x2b\x55\x81\x27\x7e\x15\x69\x3b\x9c\x9f\x34\x3f\x36\xc7\xf2\xf8\x78\x49\x4f\x7b\x2d\x0b\x3b\x9f\x85\x10\x79\x9f\x99\x06\xc3\x1e\x14\x26\x4f\xcf\xdd\x5e\xe5\x33\xeb\xc4\x4f\x69\x71\xf5\xfb\xfb\xdf\x2e\x3f\x5c\x5e\x9d\x1f\x7c\x4b\xe6\x74\x20\xd6\x56\xb3\x86\x94\x63\x69\xbd\x33\xd2\x8d\xbb\xef\xf1\x00\x53\x18\xd2\x62\x85\x15\xe4\x06\xa7\x47\xe2\xea\xfa\xf2\x8f\x9f\x6f\x2e\xa6\x42\x37\xd6\xc3\xab\x36\x8e\xb7\xb0\xb4\x03\x71\x7f\x8c\x07\xe8\xdb\x2d\x0e\x95\x0b\xc4\x50\xb8\x1f\x93\x0b\x8c\x83\xd3\xc7\xc3\x67\xbe\x95\xaa\x74\x27\x47\x42\x2c\xb3\x21\x51\x26\x7e\x8e\x5c\x70\x0c\x29\xf5\xe8\x32\xa5\x1d\x25\xd9\xba\x60\x65\xdd\x26\x1c\xdc\x57\xfb\x3f\x42\x4a\x8e\x5b\x0a\x38\xb8\x9f\xee\xc8\x23\xde\xe0\xb5\xdd\x51\x62\x97\x49\x56\x81\x94\x73\xad\x5b\x26\x7e\x7c\x83\x34\x94\x58\xd2\x1d\x27\x2d\xb3\x0e\x6f\x71\x07\xf7\xcf\xd3\x3e\x8a\xe9\x65\xd4\x61\x9b\xf1\x91\x7c\xa6\x3d\xce\x8e\x26\xbd\x75\x6b\xf5\x38\xfa\xd9\x00\x72\xf3\x64\x45\x1b\x4d\x6e\xc6\x14\xeb\xd3\x63\xb4\x77\x26\x36\x31\x75\x6a\x77\xda\x7c\xa7\x3a\xe2\xfa\x78\xc9\xe9\x48\xe8\xb5\xe3\xd4\x62\x28\x5d\xdf\x12\xd5\xc7\xfc\xb4\xac\xed\xa0\x4e\x9a\x1f\xd4\x42\xb5\xd7\x83\x9f\x1b\xca\x49\x44\x0c\x1b\xd7\xcd\x8e\x7e\x45\xa6\xed\xe0\x42\x53\x17\x85\x18\xb6\xd6\x25\xc8\x11\xaa\x8f\x03\xa9\xd2\x96\xc0\x45\x4d\xb7\x45\x98\xf1\x3f\x4a\xbf\x50\xa0\xe6\x7e\xc2\xf4\xf1\x36\x40\x5e\x63\x5e\x3c\x9b\x3f\x5f\xea\xf0\x6f\x00\x00\x00\xff\xff\x3e\xd7\xe8\x85\xf3\x06\x00\x00")

func bootstrapAmazon_k8s_ubuntu_1604_masterShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapAmazon_k8s_ubuntu_1604_masterSh,
		"bootstrap/amazon_k8s_ubuntu_16.04_master.sh",
	)
}

func bootstrapAmazon_k8s_ubuntu_1604_masterSh() (*asset, error) {
	bytes, err := bootstrapAmazon_k8s_ubuntu_1604_masterShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/amazon_k8s_ubuntu_16.04_master.sh", size: 1779, mode: os.FileMode(436), modTime: time.Unix(1502293337, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapAmazon_k8s_ubuntu_1604_nodeSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x52\x3d\x6f\xdc\x38\x14\xec\xf9\x2b\xe6\x56\x06\xdc\x58\x92\xaf\x3a\xc0\x80\x0f\xb8\x8f\x2d\x7c\x87\x38\x40\xb2\x40\x9a\x34\x14\xf9\x2c\xd1\x4b\xf1\x09\x7c\x8f\xb6\x17\x8e\xf3\xdb\x03\x69\x15\x7f\x74\xe9\xac\x46\x9a\xe1\x68\x38\x8f\x9c\xea\xb7\xb6\x48\x6e\xbb\x90\x5a\x4a\x77\xe8\xac\x0c\x46\x48\x51\x93\x71\x1e\xdf\x8d\xa9\x50\xbf\xd3\x63\x2a\xec\x06\x12\xc2\x9d\x8d\x85\x04\x36\x13\x42\xba\x25\xa7\xe4\x11\x92\x32\x74\x20\x88\xcb\x61\xd2\x06\x5f\x68\x11\xd0\xc3\x14\x83\x0b\x1a\x0f\x48\xac\x28\x12\x52\x0f\x0b\xa5\x71\x8a\x56\x67\x10\x6d\xea\x8b\xed\x09\xca\xab\xdd\xe2\x73\xdc\xc4\x54\xb0\x32\xaf\x50\x72\x5c\xf2\x22\x1b\x08\x45\x28\xcf\x6c\x0c\x63\x58\xe4\x21\xcf\x1c\xf8\xe6\x8d\x33\xf7\xc1\x21\xa4\x59\x20\x84\x9b\x10\x49\x1a\xfc\x7d\x80\x27\x09\x7d\x82\x8d\xf1\x65\x80\x75\x28\x19\xb8\x44\x6f\x2a\x74\x04\xdb\xc5\x25\x55\x47\x98\xaf\xc0\x2a\x72\x49\x1a\x46\x3a\x83\x4d\xfe\x38\xed\x40\x31\xae\x33\x23\x93\x8d\xb8\xe7\xbc\x6f\x70\x75\x83\x03\x17\x24\x22\x0f\xc7\xc9\x07\x0d\x9c\x6c\x3c\x46\x3a\xc3\x7d\x0e\x4a\x08\x3a\x87\x5b\xae\xb8\x02\x67\x8c\x76\x4f\xb0\x89\x75\xa0\xfc\xc6\xb9\x31\x95\xa9\xcc\xee\xe3\xff\xdb\xeb\xcb\xcd\xd5\xf5\x7f\xdb\x7f\x76\xdb\x7f\x17\xb8\x31\x1f\xfe\xfa\xbc\xdb\x7e\x7a\xa1\x8f\x78\xf3\x9e\x3d\x31\xae\xe4\x88\x5a\x30\xa8\x4e\x72\xd1\xb6\x93\x75\x7b\xdb\x93\x34\x2e\x72\xf1\x4d\xcf\xdc\x47\x6a\x1c\x8f\xad\x9d\xb4\xf5\xec\xe6\x77\xbd\xa7\x43\xd3\x4f\x3d\xbe\x41\x8a\x67\xac\x14\xac\xf7\xa8\x8d\x72\x71\x03\x5a\xd2\x45\xdb\x0a\x97\xec\x48\x9a\x18\x44\x1b\xdf\xee\x4b\x47\x39\x91\xae\x8c\x91\x01\xb5\xc3\x29\xb9\x81\xb1\xf1\xd4\x2d\x49\x2e\xda\xf9\xd7\xe6\x95\x36\x70\x8b\x17\x58\x3f\x50\x0a\x36\x62\xb4\x21\x6d\xf0\xe7\xaf\x6e\x76\x6a\xcc\x1c\xb5\x27\x45\x99\xbc\x55\x42\x7d\x78\x66\x42\x12\x9d\x6b\x56\x1f\xf0\xd5\x00\x80\xb0\xb3\xba\x7e\x53\xa7\x73\xc5\x64\x85\x9e\xdd\x9e\x72\x13\x78\xc5\xb3\x87\x66\x9b\x64\xe2\xac\xf5\x72\x96\xeb\xca\x9c\x20\x92\xbe\x42\xd6\x8f\x97\xbf\x37\x7f\x34\xe7\xf5\xf9\xb9\x31\x72\x10\xa5\xd1\x69\x04\xa5\xa5\xc4\x47\xeb\x57\xbc\xa8\xcd\xfa\x93\x36\xab\x03\x32\x09\xe9\x33\xba\xe5\x90\x50\xd7\xca\x7b\x4a\x38\x79\x5c\xfa\xf6\x84\x93\xc7\x63\xc3\x9e\x7e\x04\x00\x00\xff\xff\x23\x59\x14\x94\xaa\x04\x00\x00")

func bootstrapAmazon_k8s_ubuntu_1604_nodeShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapAmazon_k8s_ubuntu_1604_nodeSh,
		"bootstrap/amazon_k8s_ubuntu_16.04_node.sh",
	)
}

func bootstrapAmazon_k8s_ubuntu_1604_nodeSh() (*asset, error) {
	bytes, err := bootstrapAmazon_k8s_ubuntu_1604_nodeShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/amazon_k8s_ubuntu_16.04_node.sh", size: 1194, mode: os.FileMode(436), modTime: time.Unix(1502272644, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapDigitalocean_k8s_ubuntu_1604_masterSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x55\x41\x6f\x1b\x37\x13\xbd\xf3\x57\xbc\x4f\x12\xe0\x2f\x40\xb8\xb4\x5d\xa0\x05\x0c\xb8\x40\x92\xfa\xe0\xb6\x48\x0c\xc3\x6d\x2f\xbd\x50\xe4\x68\x97\x15\x97\x5c\x90\xb3\xb2\x05\xdb\xf9\xed\x05\xa9\x95\x2c\x5d\x8a\x02\x3d\x44\x17\x69\x1e\x1f\x67\xde\x90\x6f\x28\x31\xff\x9f\x1a\x73\x52\x4b\x17\x14\x85\x0d\x96\x3a\x77\x22\x13\x43\x92\x30\x16\x5f\x85\x98\x43\x7e\xa3\x8f\x98\xe3\xa1\xa3\x4c\xd8\x68\x3f\x52\x86\x4e\x04\x17\xfe\x22\xc3\x64\xe1\x02\x47\x70\x47\xc8\x26\xb9\x81\x1b\xfc\x41\x95\x40\x4f\x83\x77\xc6\xb1\xdf\x22\x44\xc6\x98\x5d\x68\xa1\xc1\xd4\x0f\x5e\x73\x09\xbc\x0e\xed\xa8\x5b\x02\xc7\x29\x5d\xcd\xb3\x2b\x22\xe6\xd0\xb9\xac\x50\x30\x71\x4c\x95\xd6\x11\xc6\x4c\xa9\xa0\xde\xf5\xae\xd2\x5d\x2a\x18\xe2\xea\x24\x73\x6c\x9d\x81\x0b\x85\x90\x09\x2b\xe7\x29\x37\xf8\xb8\x85\xa5\xec\xda\x00\xed\xfd\x5b\x03\x53\x53\xb9\x8b\xa3\xb7\x62\x8e\x25\x41\x2f\x7d\x55\xb5\x24\x94\x2b\xd0\x8c\x34\x06\x76\x3d\xbd\x87\x0e\x76\xd7\x6d\x47\xde\x4f\x3d\x23\x91\xf6\x78\x8c\x69\xdd\xe0\x76\x85\x6d\x1c\x11\x88\x2c\x4c\x0c\xd6\xb1\x8b\x41\xfb\x9d\xa4\xf7\x78\x4c\x8e\x09\x8e\x8b\xb8\x7a\xc5\x73\xc4\x84\x5e\xaf\x09\x3a\x44\xee\x28\x9d\x64\x6e\xc4\x5c\xcc\xc5\xc3\x97\x5f\x6e\x3e\x5f\xcf\x6e\x3f\xff\x7c\xf3\xe9\xe1\xe6\xa7\x1a\xce\xc4\xdd\x97\xfb\x87\x37\xb0\x44\xb3\xca\xfe\x86\x3e\x11\x77\xf7\xb7\xbf\x7f\x78\xb8\xb9\xbd\xbb\x5e\xfc\xdf\xad\x4c\x0c\x2b\xd7\xe2\x05\x6d\xa2\x01\xf2\x03\x2e\x30\xe3\x31\x9c\xcf\xf6\x90\x0b\xc4\xc0\x0b\xcc\xc8\x90\x16\xb3\xab\x19\xe4\x0a\x97\x47\x08\x2a\x72\xf1\x4e\x90\xe9\x22\x16\x87\xfc\xf8\x11\x8a\xfb\x41\x35\x6e\x10\x77\xbf\x7d\xfc\xf5\xf6\x53\xad\x69\xc6\xe4\xb1\x2f\xdc\xf4\xf4\x4e\x88\x0a\xc9\x8c\x8e\x79\xc8\x57\x4a\x0d\xda\xac\x75\x4b\xb9\x31\x3e\x8e\xb6\x69\x63\x6c\x3d\x35\x26\xf6\x4a\x0f\xac\x6c\x34\xe5\x5b\xae\x69\xdb\xb4\x43\xd1\x9e\x47\x1b\x31\x41\xd0\xd6\x42\x0a\x8e\xa3\xe9\xa0\x88\x2b\x57\xe5\x38\x26\x43\xb9\xf1\x2e\x73\x63\xd5\x7a\x5c\x52\x0a\xc4\x13\x22\x72\x07\x69\x70\x56\x1b\x98\x59\x5a\x56\x25\x57\xaa\x6c\x6d\x8e\xb8\x2e\x2a\xbc\x85\xf2\x89\x82\xd3\x1e\xbd\x76\x61\x56\x9a\xfd\x77\xc5\xce\x84\x28\x52\x5b\x62\x8c\x83\xd5\x4c\x90\xdb\x03\xe2\x42\xe6\xe2\x7c\xb9\xc5\x9f\x02\x00\x72\x34\x9a\xa7\xdf\xb4\xe4\xe2\xfa\x3c\x85\x36\x9a\x35\xa5\xc6\xc5\x29\x2e\x39\x38\xe9\x90\x87\x98\x58\xd6\xb3\x9c\x56\x8a\x02\x4f\x7c\x14\x69\xdb\x5f\x5f\x34\x3f\x34\xe7\xf2\xfc\x7c\x82\xeb\x59\xcb\x91\x9d\xcf\x42\x88\xbc\xcd\x4c\xbd\x61\x0f\x0a\x75\xd4\x76\xd5\x8e\xf0\xcc\x3a\xf1\x1e\x3e\x75\xd5\x50\x2e\xa1\x4c\x49\x7c\x84\xa5\x0d\x8a\xa1\xf0\x02\xfd\xb8\xc6\x99\xaa\x86\x52\x78\x1e\x92\x0b\x8c\xc5\xe5\xeb\xd9\xc1\x4b\x33\x55\xac\xf4\x5f\x9c\x34\xf5\x86\x44\x99\xf8\x10\xb9\xe0\x18\x52\xea\xc1\x65\x4a\x1b\x4a\x72\xe9\x82\x95\xe5\x98\xb0\x78\x2e\x53\xf9\x0a\x29\x39\xae\x29\x60\xf1\x5c\x47\xf7\x15\x27\x7c\x6d\x37\x94\xd8\x65\x92\xa5\x31\xca\xb9\xec\x9b\x84\xbc\x9e\x30\x0d\x25\x96\xf4\xc4\x49\xcb\xac\xc3\x29\x6f\xf1\x7c\xe8\xe8\x55\xec\xb4\x96\x83\xd4\xc3\xe0\x77\xd7\x2d\x57\x7b\xe3\xd9\x68\x72\x33\xa4\x58\xde\x3f\xa3\xbd\x33\xb1\x89\xa9\x55\x9b\xcb\xe6\x3b\xd5\x12\x97\x17\x54\xd6\x0b\xa0\x63\x7f\xa9\xc9\x3e\xba\x3c\x68\xaa\x8b\x79\xbf\xac\x6d\xaf\x2e\x9a\xef\xd5\x94\x6a\xab\x7b\xbf\x2b\x28\xab\x88\xdd\x1b\x50\xfd\x7b\x94\x4c\xdb\xde\x85\xa6\x2c\x96\x3f\xb6\xfb\x18\x59\xf4\x6b\xeb\x12\xe4\x80\xaf\xaa\xce\x85\x30\xc3\x3f\x6c\xdb\xb3\xd4\xae\x80\x10\x7f\x07\x00\x00\xff\xff\x70\x8f\xcc\x07\x41\x07\x00\x00")

func bootstrapDigitalocean_k8s_ubuntu_1604_masterShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapDigitalocean_k8s_ubuntu_1604_masterSh,
		"bootstrap/digitalocean_k8s_ubuntu_16.04_master.sh",
	)
}

func bootstrapDigitalocean_k8s_ubuntu_1604_masterSh() (*asset, error) {
	bytes, err := bootstrapDigitalocean_k8s_ubuntu_1604_masterShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/digitalocean_k8s_ubuntu_16.04_master.sh", size: 1857, mode: os.FileMode(436), modTime: time.Unix(1502294245, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapDigitalocean_k8s_ubuntu_1604_nodeSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x52\x4d\x4f\xdc\x4a\x10\xbc\xfb\x57\xd4\x5b\x23\x71\xc1\x36\xef\xf4\x24\x24\x9e\x94\x8f\x3d\x90\x28\x44\x4a\x56\xca\x25\x97\xf1\x4c\x63\x0f\x3b\x9e\xb6\xa6\x7b\x80\x15\x21\xbf\x3d\xf2\x07\x02\x92\x4b\x6e\xf8\x62\x57\x75\x77\x4d\xb5\xa7\x8a\xf2\x9f\x26\x4b\x6a\x5a\x1f\x1b\x8a\x37\x68\x8d\xf4\x85\x90\xa2\xa2\xc2\x3a\xfc\x2c\x8a\x12\xd5\x2b\x3d\x45\x89\x5d\x4f\x42\xb8\x31\x21\x93\xc0\x24\x82\x8f\xd7\x64\x95\x1c\x7c\x54\x86\xf6\x04\xb1\xc9\x8f\x5a\xe3\x1b\xcd\x0d\x74\x37\x06\x6f\xbd\x86\x03\x22\x2b\xb2\xf8\xd8\xc1\x40\x69\x18\x83\xd1\x09\x04\x13\xbb\x6c\x3a\x82\xf2\x2a\x37\xeb\x2c\x87\x14\x25\x8c\x4c\x15\x8a\x96\x73\x9a\xdb\x7a\x42\x16\x4a\x13\x1b\xfc\xe0\xe7\x76\x9f\x26\x0e\x7c\xf5\x42\x99\x3b\x6f\xe1\xe3\xd4\x20\x84\x2b\x1f\x48\x6a\xbc\x3d\xc0\x91\xf8\x2e\xc2\x84\xf0\xb4\xc0\xba\x94\xf4\x9c\x83\x2b\x4a\xb4\x04\xd3\x86\xd9\x55\x4b\x98\xae\xc0\x28\x52\x8e\xea\x07\x3a\x81\x89\x6e\xd9\xb6\xa7\x10\xd6\x9d\x91\xc8\x04\xdc\x72\xda\xd7\xb8\xb8\xc2\x81\x33\x22\x91\x83\xe5\xe8\xbc\x7a\x8e\x26\x2c\x96\x4e\x70\x9b\xbc\x12\xbc\x4e\xe6\xe6\x2b\x2e\xc1\x09\x83\xd9\x13\x4c\x64\xed\x29\xbd\x50\xae\x8b\xb2\x28\x8b\xdd\xe7\x8f\xdb\xcb\xf3\xcd\xc5\xe5\x87\xed\xbb\xdd\xf6\xfd\x0c\x37\xc5\xa7\x37\x5f\x77\xdb\x2f\x4f\xf4\x82\x37\xaf\x99\x93\x42\xb2\x63\xd8\x9c\x02\x2a\x41\xaf\x3a\xca\x59\xd3\x8c\xc6\xee\x4d\x47\x52\xdb\xc0\xd9\xd5\x1d\x73\x17\xa8\xb6\x3c\x34\x66\xd4\xc6\xb1\x9d\xde\xd5\x9e\x0e\x75\x37\x76\xf8\x81\x59\x64\xa5\x60\x9c\x43\xb5\xe8\x2a\x67\xdb\xa3\x21\x9d\x07\x1a\xe1\x9c\x2c\x49\x1d\xbc\x68\xed\x9a\x7d\x6e\x29\x45\xd2\x95\x59\x46\xa4\x47\x65\x71\x4c\xb6\x67\x6c\x1c\xb5\xb3\xa7\xb3\x66\x9a\xaf\x9f\x0d\x78\x6e\xf0\x04\xab\x3b\x8a\xde\x04\x0c\xc6\xc7\x0d\xfe\xff\xdb\x13\x8f\xd7\xf5\x27\xe7\x1d\x29\xf2\xe8\x8c\x12\xaa\xc3\x4b\xda\x47\xd1\x29\x7f\xd5\x01\xdf\x0b\x00\x10\xb6\x46\xd7\x6f\x6a\x75\xca\x9e\xac\xd0\xb1\xdd\x53\xaa\x3d\xaf\x78\xd2\xd0\x64\xa2\x8c\x9c\xb4\x9a\xff\xef\x5a\x99\xbc\x04\xd2\x67\xc8\xb8\xe1\xfc\xdf\xfa\xbf\xfa\xb4\x3a\x3d\x5d\x9d\xc9\x41\x94\x06\xab\x01\x14\xe7\x88\x2f\xfa\xbf\x17\x45\x4d\xd2\xc7\xda\x52\xac\xb6\x8f\x9a\x48\x24\xa4\x7f\xb0\xd7\xec\x23\xaa\x4a\x79\x4f\x11\x47\xf7\x73\x44\x1f\x70\x74\xbf\x84\xf2\xa1\xf8\x15\x00\x00\xff\xff\x4c\xd3\xf9\xfb\xdf\x04\x00\x00")

func bootstrapDigitalocean_k8s_ubuntu_1604_nodeShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapDigitalocean_k8s_ubuntu_1604_nodeSh,
		"bootstrap/digitalocean_k8s_ubuntu_16.04_node.sh",
	)
}

func bootstrapDigitalocean_k8s_ubuntu_1604_nodeSh() (*asset, error) {
	bytes, err := bootstrapDigitalocean_k8s_ubuntu_1604_nodeShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/digitalocean_k8s_ubuntu_16.04_node.sh", size: 1247, mode: os.FileMode(436), modTime: time.Unix(1502294245, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapInjectGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x91\xd1\x6a\xdb\x30\x18\x85\xaf\xa3\xa7\x38\xe4\x2a\x01\xcf\x6e\x7b\x33\xe8\xd8\x85\xd7\x76\xcc\xb4\x24\x10\xa7\x2b\xa5\xf4\x42\x96\xff\xd8\x5a\x1d\x49\x93\x7e\xd7\x0d\xa5\x0f\xb4\xd7\xd8\x93\x0d\x25\xce\x58\x99\x6f\x0c\xfa\xcf\xf9\x74\xce\xaf\x2c\xc3\x85\x75\x3b\xaf\x9b\x96\xf1\xfb\x17\xce\x4e\x4e\x3f\x62\xdd\x12\xae\xfb\x4a\x2b\xeb\x0d\xf2\x9e\x5b\xeb\x83\xc8\x32\x91\x65\xb8\xd1\x8a\x4c\xa0\x1a\xbd\xa9\xc9\x83\x5b\x42\xee\xa4\x6a\xe9\x38\x49\xf0\x9d\x7c\xd0\xd6\xe0\x2c\x3d\xc1\x2c\x0a\xa6\xe3\x68\x3a\xff\x14\x11\x3b\xdb\x63\x2b\x77\x30\x96\xd1\x07\x02\xb7\x3a\x60\xa3\x3b\x02\xbd\x28\x72\x0c\x6d\xa0\xec\xd6\x75\x5a\x1a\x45\x18\x34\xb7\xfb\x7b\x46\x4a\x1a\x19\xf7\x23\xc3\x56\x2c\xb5\x81\x84\xb2\x6e\x07\xbb\xf9\x57\x08\xc9\x63\xe8\xf8\xb5\xcc\xee\x3c\xcb\x86\x61\x48\xe5\x3e\x70\x6a\x7d\x93\x75\x07\x69\xc8\x6e\x8a\x8b\xab\x45\x79\xf5\xe1\x2c\x3d\x19\x4d\xb7\xa6\xa3\x10\xe0\xe9\x67\xaf\x3d\xd5\xa8\x76\x90\xce\x75\x5a\xc9\xaa\x23\x74\x72\x80\xf5\x90\x8d\x27\xaa\xc1\x36\x86\x1e\xbc\x66\x6d\x9a\x04\xc1\x6e\x78\x90\x9e\x22\xa6\xd6\x81\xbd\xae\x7a\x7e\xb7\xb3\x63\x44\x1d\xde\x09\xac\x81\x34\x98\xe6\x25\x8a\x72\x8a\x2f\x79\x59\x94\x49\x84\xdc\x15\xeb\x6f\xcb\xdb\x35\xee\xf2\xd5\x2a\x5f\xac\x8b\xab\x12\xcb\x15\x2e\x96\x8b\xcb\x62\x5d\x2c\x17\x25\x96\x5f\x91\x2f\xee\x71\x5d\x2c\x2e\x13\x90\xe6\x96\x3c\xe8\xc5\xf9\xd8\xc0\x7a\xe8\xb8\x4d\xaa\xf7\xab\x2b\x89\xde\x45\xd8\xd8\x43\xa4\xe0\x48\xe9\x8d\x56\xe8\xa4\x69\x7a\xd9\x10\x1a\xfb\x4c\xde\x68\xd3\xc0\x91\xdf\xea\x10\x5f\x35\x40\x9a\x3a\x62\x3a\xbd\xd5\x2c\x79\x7f\xf4\x5f\xaf\x54\x08\x27\xd5\x53\x84\x54\xd6\x72\x60\x2f\x9d\x10\x7a\xeb\xac\x67\x4c\x63\x5f\xd3\x84\xa9\x10\x9b\xde\x28\x14\xe6\x07\x29\x9e\xd5\x92\x25\x1e\x1e\xab\x1d\x53\x82\x67\xd9\xf5\x14\xb0\x95\xee\xe1\xa0\x7e\x3c\xfc\xe6\x98\x1d\x25\xe4\xbd\xf5\x73\xbc\x8a\x49\x60\x7f\x19\xcd\xe7\x9f\x71\x50\xed\x59\x73\x31\x89\xd5\x9e\x12\x3c\xc7\x89\x97\xa6\xa1\x23\xf7\x55\x4c\xfe\xba\x8e\xa6\x90\xae\xc8\x75\x52\xd1\x6c\x9c\x24\x7b\x6f\x82\xd3\xb9\x98\xbc\x89\x89\x27\xee\xbd\x19\x13\x1e\x35\xf3\x04\x46\x77\xe2\x4d\xfc\x09\x00\x00\xff\xff\xbc\x34\x80\x62\x48\x03\x00\x00")

func bootstrapInjectGoBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapInjectGo,
		"bootstrap/inject.go",
	)
}

func bootstrapInjectGo() (*asset, error) {
	bytes, err := bootstrapInjectGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/inject.go", size: 840, mode: os.FileMode(436), modTime: time.Unix(1501967388, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapVpnMeshbirdmasterSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xce\x3d\x4b\xc4\x40\x10\xc6\xf1\x7e\x3e\xc5\xe3\xdd\x71\x68\x91\x4c\x65\x23\x68\x21\x1e\x28\x12\x11\x05\xc1\x4a\x92\xcd\xe8\xae\x6e\x76\xc3\xce\xc4\x17\x08\x7e\x76\x11\x14\x63\x39\xfc\x99\x1f\xcf\x7a\x8f\x27\x2d\xdc\x85\xc4\x92\x5e\xd1\xb5\xea\x49\xc5\x50\x09\xb9\x1e\x9f\x44\x6b\xdc\x5d\x5f\xa1\x11\xf5\xe4\xa6\x12\xe1\xcd\xc6\x23\xe6\x41\xd4\x77\xa1\xf4\xb5\xcb\x03\x87\xa4\xd6\xc6\x58\xab\xc7\x0c\xf5\xf4\x1b\x91\xe4\x0d\xdb\x13\xb0\x0d\x23\xc7\xfc\xf4\x22\x1f\xd4\xec\x6e\xcf\x4f\x2f\x6e\xce\x1e\x2e\x77\xf7\xc7\x9b\x7d\xd7\xda\x32\x63\x86\x9b\x0c\x55\x8f\x15\x56\xa8\x1e\x71\x78\x40\xe2\x7c\xc6\x66\xf9\x87\x1f\xb2\xfe\x06\xe5\x7d\xcc\xc5\xf0\xdf\x5d\x5e\x7f\x73\x9e\x73\x48\xd8\xd2\x57\x00\x00\x00\xff\xff\xda\xd3\x92\x18\xf4\x00\x00\x00")

func bootstrapVpnMeshbirdmasterShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapVpnMeshbirdmasterSh,
		"bootstrap/vpn/meshbirdMaster.sh",
	)
}

func bootstrapVpnMeshbirdmasterSh() (*asset, error) {
	bytes, err := bootstrapVpnMeshbirdmasterShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/vpn/meshbirdMaster.sh", size: 244, mode: os.FileMode(436), modTime: time.Unix(1502293337, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapVpnMeshbirdnodeSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x92\x41\x6f\xd4\x30\x10\x85\xef\xfe\x15\x8f\x06\x71\x6a\x93\x3b\x12\x97\xd2\x95\x58\xaa\x56\x08\x2a\x50\x4f\xc8\x71\x66\xe3\x69\x9d\x71\xe4\x19\xb7\x5d\x09\xf1\xdb\x51\xb2\x8b\x60\x7f\x41\x73\x9b\x97\xe7\xe7\xf7\x69\xdc\xbc\xe9\xaa\x96\xae\x67\xe9\x48\x9e\xd0\x7b\x8d\x4e\xc9\x70\x41\x2e\x0c\xf8\xed\x5c\x83\x8b\x57\xfa\x5c\x83\xbb\x48\x4a\x78\xf2\xa9\x92\xc2\x17\x02\xcb\x03\x05\xa3\x01\x2c\x96\x61\x91\xa0\xa1\xf0\x6c\x2d\x7e\xd0\x6a\xa0\x97\x39\x71\x60\x4b\x7b\x48\x36\x54\x65\x19\xe1\x61\x34\xcd\xc9\xdb\x32\x24\x2f\x63\xf5\x23\xc1\xf2\x31\x6e\xcd\x39\x5c\xe2\x1a\x78\x5d\xfe\x90\x84\x5c\xcb\x6a\x8b\x84\xaa\x54\x16\x35\xf1\xc4\xab\x9d\xcb\xa2\x21\xef\x4e\x92\xf3\xc8\x01\x2c\x8b\x41\x09\x3b\x4e\xa4\x2d\x2e\xf7\x18\x48\x79\x14\xf8\x94\xfe\x01\x1c\xa1\x34\xe6\x9a\x06\xd7\xa0\x27\xf8\x3e\xad\xad\x7a\xc2\xb2\x02\x6f\x28\x55\x8c\x27\x3a\x87\x97\xe1\x40\x1b\x29\xa5\x23\x33\x0a\xf9\x84\xe7\x5c\x1e\x5b\x6c\x77\xd8\xe7\x0a\x21\x1a\x10\xb2\x0c\x6c\x9c\xc5\xa7\x43\xa5\x73\x3c\x17\x36\x02\xdb\x52\x6e\x5d\x71\x83\x5c\x30\xf9\x47\x82\x97\x6c\x91\xca\x49\x72\xeb\x1a\xd7\xb8\x9b\xcd\xb7\x4f\x97\xdb\xaf\x57\x3f\xaf\x37\xf7\x1f\xce\xb6\xb7\x9f\x37\x1f\xef\x36\x57\x8b\x7a\xbd\xb9\x3f\x7b\xcd\x87\xe1\x1a\x7c\xff\x72\x8b\x1b\xd2\xe8\x42\x2d\x09\xd1\x6c\x7e\xdf\x75\x13\x69\xec\xb9\x0c\x6d\xc8\x53\xc7\xa2\xe6\x53\x6a\x35\xe2\x17\x34\x3a\x7a\x99\x73\x31\x9c\x50\xbd\xfd\x7f\x72\x7f\x8f\xe3\x21\xb3\xe0\x9d\xfb\x13\x00\x00\xff\xff\x85\xc7\x28\xe5\x1c\x03\x00\x00")

func bootstrapVpnMeshbirdnodeShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapVpnMeshbirdnodeSh,
		"bootstrap/vpn/meshbirdNode.sh",
	)
}

func bootstrapVpnMeshbirdnodeSh() (*asset, error) {
	bytes, err := bootstrapVpnMeshbirdnodeShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/vpn/meshbirdNode.sh", size: 796, mode: os.FileMode(436), modTime: time.Unix(1502293337, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapVpnOpenvpnmasterSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x57\x5f\x6f\xda\xc8\x16\x7f\xf7\xa7\x38\xd7\x44\xda\x36\x5a\x7b\x42\x94\xdd\xbb\xdb\x0d\xd1\x52\x4a\x73\x51\xdb\x80\x12\xc8\x55\x24\xa4\x68\x18\x1f\xf0\x14\x33\x63\xcd\x8c\x49\xdc\xdc\xf4\xb3\x5f\xcd\xd8\x10\x0c\x24\x81\xee\x43\xf3\x42\x7c\xfe\xfc\xce\x39\x33\xe7\xdf\xd4\xfe\x45\x32\xad\xc8\x88\x0b\x82\x62\x0e\x23\xaa\x63\x4f\xa3\x81\x00\x3d\x16\xc1\x77\xcf\xab\x41\xf0\x93\xfe\xbc\x1a\xf4\x63\xd4\x08\x73\x9a\x64\xa8\x81\x2a\x04\x2e\xbe\x22\x33\x18\x01\x17\x46\x82\x89\x11\x34\x53\x3c\x35\x21\xfc\x17\x9d\x00\xde\xa7\x09\x67\xdc\x24\x39\x08\x69\x20\xd3\x5c\x4c\x80\x82\xc1\x59\x9a\x50\x63\x3f\x12\x2a\x26\x19\x9d\x20\x18\x59\xc2\x39\x9c\xc2\x88\x57\x03\xaa\x2d\x07\x05\x93\x99\x72\x62\x31\x42\xa6\x51\x59\x6a\xc2\x67\xdc\x89\x73\x65\x69\x20\xc7\x15\x64\x39\xe1\x0c\xb8\xb0\x02\x1a\x61\xcc\x13\xd4\x21\xbc\xcf\x21\x42\xcd\x27\x02\x68\x92\x3c\x05\x50\x06\xa5\x63\x99\x25\x91\x57\x83\x11\x02\x1d\x25\xce\xab\x11\x82\xbd\x02\x6a\x40\x65\xc2\xf0\x19\xfe\x0a\x54\x44\x45\xb4\x31\x26\x49\x19\x33\x28\xa4\x09\xdc\x49\x35\x0d\xa1\x33\x86\x5c\x66\x20\x10\x23\x60\x52\x44\xdc\x70\x29\x68\x52\xb8\xf4\x2b\xdc\x29\x6e\x10\xb8\xb1\xce\xb9\x2b\xae\x81\x54\x30\xa3\x53\x04\x2a\xa4\x89\x51\x55\x90\x43\xaf\xe6\xd5\xbc\x6e\xaf\x7d\x71\xdd\xbb\xb8\xfd\xd4\xbe\x69\x75\x07\x17\xfd\xcb\x9b\x86\x3f\xb8\xf2\x57\xe9\xbd\xcb\xee\x75\xe7\xa2\xd5\x6e\xf8\xad\x66\x85\xd1\xea\xf4\x6f\x1a\xfe\x15\x15\x1f\x15\x15\x8c\x6b\x26\x2b\xec\xee\xe5\x79\xc3\xff\x94\x8d\x38\x93\x4a\x54\x38\xed\x2f\xcd\xce\xe7\x86\xaf\xa4\x34\x7f\x27\x92\xd1\x24\x96\xda\x54\x75\x07\xcf\xa8\x5e\x34\xbf\xb4\x1b\xbe\x46\x35\x47\xe5\xff\xcc\xbc\xf5\x7a\x97\x9d\xeb\x66\xbf\x7d\xdb\xe9\x35\x0e\xde\xb0\x4c\x25\x10\x1b\x93\xbe\x23\xa4\xfe\xfb\x9f\xe1\xf1\x6f\x27\x61\xf9\x4b\x66\x68\x68\x44\x0d\x25\xf3\x3a\xe1\xc2\xa0\x1a\x53\x86\x9a\xa4\x8a\xcf\xa9\x41\x72\x44\x78\x3a\x3f\x21\x34\x8a\x14\x6a\xfd\xd6\x16\x63\x37\x45\x71\xdd\xbb\xf0\x3c\x9a\x9a\x60\x82\x06\xb2\x34\xa2\x06\x97\x9f\x5c\x68\x63\xf3\x2c\xc8\x41\xa6\x28\xe6\xa9\x00\xa4\x3a\x0f\x94\xa6\x9e\x67\x2f\x3c\x60\x34\xe2\x0a\xbe\x93\x92\x1d\x30\xea\x79\x1a\x23\x08\x38\x04\x08\xbe\x26\x78\x9f\x4a\x65\xe0\x53\xfb\xe6\xb6\xbc\xf7\xf0\x70\x0b\xb1\x31\xf4\x0f\x1e\x36\x73\xe4\x71\xe8\x13\xbf\x02\x4f\xe6\x54\xe9\x67\x4d\x2c\x52\xa8\x6a\x63\x99\x58\x6b\x46\x16\xf4\x7d\xad\xd8\x7c\x5c\x8b\xc2\x66\xe8\x7a\x08\x9d\xfe\xde\xfe\x77\x2f\xcf\xab\xc0\x36\xb7\xd7\x70\xbb\x97\xe7\xfb\xc2\xba\x3a\xa8\x02\x17\xa5\xb1\x06\xed\x88\x7b\xfb\x3c\x58\x73\x79\xb0\xe1\xf1\x60\x5f\x4c\x5b\x7d\x55\x54\x57\x8f\x6b\xb8\x96\xf6\x1c\xb2\x57\xab\xc1\x39\x0a\x54\xd4\xd8\x06\x68\xcb\x18\x18\x2a\xc3\xc7\x9c\x51\x83\xda\x0d\xa4\xd5\xb4\xd5\x32\x53\xcc\x76\x6e\xa5\xbd\x90\xb0\x04\xa9\x08\x68\x92\x78\x21\x19\x65\x3c\x89\x02\x46\x21\x08\x46\xd4\xb0\x78\x49\x9a\x62\x1e\x94\xd0\x25\x0b\x36\xfd\x5b\x4a\x47\xb1\xb7\x28\xa2\x20\x98\xa0\x98\x62\x0e\x41\xa0\x91\x29\x34\x30\xc5\x5c\x13\x43\xc3\x29\xe6\x55\xd7\x59\xc2\x51\x98\xaa\xeb\x2b\xf6\x97\x86\x0b\xb9\xb5\xb8\xcb\xfa\xb6\x4d\x7c\xcc\x27\x99\xa2\xb6\x91\x7b\x2c\xad\x9e\x97\xb3\xcd\x68\xc8\x94\xd9\xce\xb0\x76\x36\x19\x9b\xa1\x3a\x84\xa1\x07\x00\xbb\x89\x6f\xc7\x35\xcf\x19\x8c\xe2\xe3\xa3\x93\x3f\xc2\x14\x67\x40\xd0\xb0\x05\xdf\x9b\x64\xe2\x1b\x4f\x21\x60\xe0\xb6\x0f\x1d\x53\x85\x24\x92\x4b\x01\x82\xf7\x74\x96\x26\xa8\x89\x76\xbf\x41\x71\x1c\x81\x1b\xa9\xa4\xb8\xc0\xd0\xd2\xc2\xc9\x37\xf8\x1f\x18\xc4\x0a\xfe\xd6\x40\xa5\x18\xdb\xa3\xae\x41\x33\xfa\x9a\x69\x03\xfd\xcf\x57\x6b\xa7\x5c\x49\xed\xe1\x5f\x26\xd1\x01\xcd\x4c\x0c\x65\x78\x47\xe1\x21\xd9\xa0\x11\x7f\x37\xcb\x2b\xd8\x5b\x40\xa8\xcd\xa6\x20\xe2\x0a\x99\xf5\x04\x8e\x76\x44\x75\xf1\xb4\x85\x5b\x1c\x9a\xed\xab\xa0\x7e\xfc\x47\xd0\x7a\xdf\x02\x16\xf3\x14\xd7\x8b\x75\xf8\x17\xe3\xa9\x1d\xf5\x2b\x92\xe1\x21\xd9\x24\xfe\x48\x4c\x5b\x50\x28\xb8\x20\xaf\xfe\xd3\x3c\xfe\xed\xf7\x7d\x02\xba\xb2\x63\xcd\x6e\x5c\x76\xeb\x99\x28\x99\xa5\xeb\x81\x38\xae\x90\x23\x19\xe5\xe1\x21\x59\xf9\xfa\x01\xd7\x2d\x9e\x33\x02\x42\xba\xdf\xf0\x90\x54\xbe\x77\xc5\x74\xbe\xf7\xbb\x1f\xba\x6f\xee\x67\x59\xa4\x38\x7f\xfb\x0e\xc6\x5c\x44\x70\x47\x73\xbb\xd6\x4d\x16\x35\x2e\xf0\xce\x75\x07\x18\x4b\x05\x38\x47\x95\x97\xbd\x60\xf5\x3a\xa3\xcc\x2e\xb2\xae\x9f\xac\x36\x92\x35\xc7\x97\x52\x01\x13\xe1\x21\x59\xfd\xdc\xc3\xed\x85\xcd\x4e\xcf\xba\x74\x47\x55\xc4\xc5\x64\xcd\x54\x4d\xa0\x09\xed\x3a\x12\xf2\xf4\xb6\x94\x0a\x0f\xc9\x16\x6a\xa3\xbe\x30\xad\x73\xcd\x4c\x52\x9a\xd1\xb9\x36\x38\x63\x26\x01\x6d\xa8\x32\x8b\xf5\xe4\xef\x2d\xfd\xf7\x49\x14\x0b\xc7\x5e\x90\xdd\xda\x7a\x2b\x65\xed\x4e\xb5\xa5\xd0\x0a\xd8\x1d\xba\x28\x33\xa9\x72\xd0\x46\x65\xcc\x64\x0a\x5d\xa6\x69\x64\xf6\x5f\x6e\xbc\xd9\xd4\xae\x48\x81\x6d\xba\x05\x62\xd9\x7f\x34\x71\x0d\xc8\x63\xf1\x4c\x46\xf0\xef\xa3\xa3\xe7\x04\x9c\xc9\x27\xaf\x1c\x0f\xc6\x4a\xce\x60\xd1\xd8\x6c\x4b\xff\x91\xde\x57\x98\x73\x47\xba\x69\x7c\x44\x35\x56\x1a\x5d\x04\x0a\x67\xd2\xb8\x9b\x2d\xb7\xc8\xea\xb5\x96\xec\xd9\x62\x2a\x06\x75\xa8\xd7\xff\x3c\x59\xd0\x0f\x1e\x9e\x56\xd9\xc7\x82\xe3\xbf\x6a\xd6\x96\x6f\x51\x8d\xef\xca\x0a\xb2\xd9\xaf\x32\xb1\x5c\x5e\xf7\xa8\xe6\x17\x8c\xed\x59\xc3\xaf\xb9\xfd\xc5\xbe\x86\xb8\x29\x9f\x8d\x08\x11\x8e\x69\x96\x54\xe6\x78\xd5\x24\xa3\x50\x8c\x60\x32\xac\x3d\xfd\xbf\xbb\xcb\xae\x05\x2c\xee\xb3\x40\x59\xa3\xec\x8e\x65\xc7\x48\xa9\x38\xc5\x9c\x0c\x6b\x6b\x84\xd7\xc3\x6f\x95\x35\x83\xaf\x0c\x8f\xfb\xe7\x46\xc6\x6e\xae\xee\x30\x28\x76\x48\xeb\xc5\xf3\x56\xa3\xb1\x2f\x6f\xed\x21\x8b\x25\xf8\xd5\x21\x5a\xf7\xe1\xec\xec\x25\xb4\x42\xa9\x78\xf2\x06\xae\xfa\xb9\xc9\xe1\x78\x37\xb5\x2c\xad\xb6\xd7\xe2\x1d\x16\x28\xd4\x32\x99\x3b\xad\xdd\x70\x22\x79\x27\xfe\x29\x52\xa5\x07\xda\xc5\xcb\xb3\xfb\xf7\x87\xce\x65\x63\x63\x25\xf3\xba\x83\x7e\x6f\xd0\x77\x4c\x62\x66\xa9\xf7\xbe\x79\xd5\xbe\x6d\x75\x2f\x3e\x76\xce\x1b\x2f\x99\x60\xd4\xc0\xc1\xc3\x8a\xf4\x63\xb9\x37\x9e\xbe\x71\x81\x04\x08\xbf\x9c\x32\x7a\xf6\xcb\xdb\x92\x7e\xf0\x50\x3a\xf1\xb8\xd8\x54\x37\xe5\x09\xa3\x67\x43\x71\x6a\xf3\x7e\xbb\x62\xb1\x22\x3f\xab\x6d\xf5\x86\xe2\x74\x8a\xf9\x8b\xea\xb6\x18\xb6\xa8\x5b\xb5\xa1\x38\x5d\xac\x63\x5b\x21\xca\x15\x6d\x8b\xf6\xa6\xda\x99\x7d\x4b\x2c\x8f\x77\xc5\x7b\x77\x80\xff\x0f\x00\x00\xff\xff\xbe\xfd\x4d\x1a\x6a\x13\x00\x00")

func bootstrapVpnOpenvpnmasterShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapVpnOpenvpnmasterSh,
		"bootstrap/vpn/openvpnMaster.sh",
	)
}

func bootstrapVpnOpenvpnmasterSh() (*asset, error) {
	bytes, err := bootstrapVpnOpenvpnmasterShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/vpn/openvpnMaster.sh", size: 4970, mode: os.FileMode(436), modTime: time.Unix(1502296602, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapVpnOpenvpnnodeSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x92\xc1\x6e\x13\x31\x10\x86\xef\x7e\x8a\x9f\x6e\x0f\x20\x35\x6b\x8a\x0a\x12\x48\x45\x40\x09\x52\x38\xa4\x11\xaa\xca\xb1\x9a\x78\x27\xd9\xa1\xce\x78\x65\xcf\xa6\x44\x08\x9e\x1d\xed\x26\x40\xcb\x0b\xd4\x17\x6b\x66\x3e\xfd\x9e\xdf\x33\xd5\x13\xdf\x97\xec\x97\xa2\x9e\x75\x8b\x25\x95\xd6\x15\x36\x4c\xd8\x85\x06\xbf\x9c\xab\x30\x79\xa4\xe3\x2a\x5c\xb5\x5c\x18\x5b\x8a\x3d\x17\x50\x66\x88\x7e\xe3\x60\xdc\x40\xd4\x12\xac\x65\x94\x90\xa5\xb3\x1a\x5f\x79\x04\xf8\x7b\x17\x25\x88\xc5\x1d\x34\x19\xfa\x22\xba\x06\xc1\x78\xd3\x45\xb2\x21\x88\xa4\xeb\x9e\xd6\x0c\x4b\x07\xb9\x51\x67\xff\x88\xab\x40\x65\xa8\xb0\x86\xd4\xe7\x11\x6b\x19\x7d\xe1\x3c\x64\xa3\x6c\x64\xc4\x25\x0f\x39\xa4\xd5\x03\xe5\xb4\x96\x00\xd1\x01\x28\x8c\x95\x44\x2e\x35\x3e\xec\xd0\x70\x91\xb5\x82\x62\xfc\x67\xe0\x60\xaa\xb4\xa9\x8f\x8d\xab\xb0\x64\xd0\x32\x8e\x5d\x2d\x19\xc3\x08\xc8\x90\x7b\x35\xd9\xf0\x09\x48\x9b\xbd\xdb\x96\x63\x3c\x78\x46\x66\x8a\xb8\x4b\xf9\xb6\xc6\x6c\x85\x5d\xea\xa1\xcc\x0d\x42\xd2\x46\x4c\x92\x52\xdc\xb7\x74\x82\xbb\x2c\xc6\x10\x1b\x9a\x1b\x47\x5c\x21\x65\x6c\xe8\x96\x41\x9a\xac\xe5\xfc\x40\xb9\x76\x95\xab\xdc\xe5\x62\x3a\xbf\x5e\xcc\x6f\x2e\x2e\xe7\x9f\xce\x8f\x66\xf3\xcf\xd3\x8b\xab\xe9\xc7\x21\x3a\x7a\xcc\xad\x70\x8b\x2f\xb3\xeb\xf7\x57\xd3\x9b\xd9\xe2\xfc\xf8\x69\xe8\x73\x44\x6b\xd6\xbd\xf1\xfe\xf4\xd5\xeb\xfa\xc5\xcb\xb3\xfa\x70\xfb\x0d\x1b\x35\x64\xe4\xb7\xa7\x5e\xd4\x38\xaf\x28\x70\xf1\x5d\x96\x2d\x19\xfb\xe7\x5e\xba\xed\x99\xa7\xa6\xc9\x5c\xca\xb3\x61\xd5\x2f\x3b\xd6\xeb\xc5\xdc\x39\xea\x6c\xb2\x66\x43\xdf\x35\x64\xfc\x37\x14\x2d\x36\x4c\x71\xb2\x43\xea\x58\xb7\x9d\x3a\xc7\xa1\x4d\x98\x30\x8e\x7f\xdc\xff\xb0\x9f\x78\x0b\xcf\x16\xfc\x81\xf3\x21\x0a\xab\x95\x3a\x24\x5d\x39\x57\x76\xc5\x78\x13\x2c\xa2\x18\x65\xfb\xa3\xf6\xee\x40\xdd\xab\xb3\x8e\x6b\xf1\x3f\xf0\x3b\x00\x00\xff\xff\x9c\xb9\xeb\xe7\xb8\x03\x00\x00")

func bootstrapVpnOpenvpnnodeShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapVpnOpenvpnnodeSh,
		"bootstrap/vpn/openvpnNode.sh",
	)
}

func bootstrapVpnOpenvpnnodeSh() (*asset, error) {
	bytes, err := bootstrapVpnOpenvpnnodeShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/vpn/openvpnNode.sh", size: 952, mode: os.FileMode(436), modTime: time.Unix(1502294295, 0)}
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
	"bootstrap/README.md":                               bootstrapReadmeMd,
	"bootstrap/amazon_k8s_ubuntu_16.04_master.sh":       bootstrapAmazon_k8s_ubuntu_1604_masterSh,
	"bootstrap/amazon_k8s_ubuntu_16.04_node.sh":         bootstrapAmazon_k8s_ubuntu_1604_nodeSh,
	"bootstrap/digitalocean_k8s_ubuntu_16.04_master.sh": bootstrapDigitalocean_k8s_ubuntu_1604_masterSh,
	"bootstrap/digitalocean_k8s_ubuntu_16.04_node.sh":   bootstrapDigitalocean_k8s_ubuntu_1604_nodeSh,
	"bootstrap/inject.go":                               bootstrapInjectGo,
	"bootstrap/vpn/meshbirdMaster.sh":                   bootstrapVpnMeshbirdmasterSh,
	"bootstrap/vpn/meshbirdNode.sh":                     bootstrapVpnMeshbirdnodeSh,
	"bootstrap/vpn/openvpnMaster.sh":                    bootstrapVpnOpenvpnmasterSh,
	"bootstrap/vpn/openvpnNode.sh":                      bootstrapVpnOpenvpnnodeSh,
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
	"bootstrap": &bintree{nil, map[string]*bintree{
		"README.md":                               &bintree{bootstrapReadmeMd, map[string]*bintree{}},
		"amazon_k8s_ubuntu_16.04_master.sh":       &bintree{bootstrapAmazon_k8s_ubuntu_1604_masterSh, map[string]*bintree{}},
		"amazon_k8s_ubuntu_16.04_node.sh":         &bintree{bootstrapAmazon_k8s_ubuntu_1604_nodeSh, map[string]*bintree{}},
		"digitalocean_k8s_ubuntu_16.04_master.sh": &bintree{bootstrapDigitalocean_k8s_ubuntu_1604_masterSh, map[string]*bintree{}},
		"digitalocean_k8s_ubuntu_16.04_node.sh":   &bintree{bootstrapDigitalocean_k8s_ubuntu_1604_nodeSh, map[string]*bintree{}},
		"inject.go":                               &bintree{bootstrapInjectGo, map[string]*bintree{}},
		"vpn": &bintree{nil, map[string]*bintree{
			"meshbirdMaster.sh": &bintree{bootstrapVpnMeshbirdmasterSh, map[string]*bintree{}},
			"meshbirdNode.sh":   &bintree{bootstrapVpnMeshbirdnodeSh, map[string]*bintree{}},
			"openvpnMaster.sh":  &bintree{bootstrapVpnOpenvpnmasterSh, map[string]*bintree{}},
			"openvpnNode.sh":    &bintree{bootstrapVpnOpenvpnnodeSh, map[string]*bintree{}},
		}},
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
