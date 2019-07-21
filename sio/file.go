// Copyright © 2019 Alexey Stolpovskikh <stolpovskikh@hotmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sio

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/alexstov/sling/slog"
	"github.com/alexstov/sling/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Sfile - Filer interface implementation.
type Sfile struct {
	Filer  Filer
	logger slog.Logger
}

// NewFiler creates a new Filer instance.
func NewFiler(logger slog.Logger) (filer Filer, err error) {
	f := &Sfile{
		logger: logger,
	}
	f.Filer = f

	return f.Filer, nil
}

// ContentType request type.
type ContentType int

const (
	// UnknownType unarchived content
	UnknownType ContentType = iota
	// ZipType application/zip
	ZipType
	// GzipType application/x-gzip
	GzipType
	// XrarType application/x-rar-compressed
	XrarType
	// WebmType video/webm
	WebmType
	// TtcfType font/collection
	TtcfType
	// OtfType font/otf
	OtfType
	// WoffType font/woff
	WoffType
	// Woff2Type font/woff2
	Woff2Type
	// WasmType application/wasm
	WasmType
	// TtfType font/ttf
	TtfType
)

func (s ContentType) String() string {
	return [...]string{"UknknownType", "ZipType", "GzipType", "XrarType", "WebmType", "TtcfType", "OtfType", "WoffType", "Woff2Type", "WasmType", "TtfType"}[s]
}

var log *logrus.Logger

// SetLogger sets Emul logger.
func SetLogger(logger *logrus.Logger) {
	log = logger
}

// SaveToFile saves bytes to file.
func SaveToFile(bytes []byte, dir string) error {
	var err error

	if dirExists, err := exists(dir); !dirExists || err != nil {
		log.Error(err)
	}

	return err
}

// Filter files by a name pattern.
func Filter(path, pattern string, filter func([]os.FileInfo, string) []os.FileInfo) ([]os.FileInfo, error) {
	entries, err := find(path)
	if err != nil {
		return nil, err
	}
	filtered := filter(entries, pattern)

	return filtered, nil
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func find(path string) ([]os.FileInfo, error) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

// MatchPattern matches filename to pattern
func MatchPattern(es []os.FileInfo, pattern string) []os.FileInfo {
	var buf []os.FileInfo
	for _, e := range es {

		matched, err := filepath.Match(pattern, e.Name())
		if err != nil {
			log.Error(err)
		}

		if matched {
			buf = append(buf, e)
		}
	}
	return buf
}

// DetermineContentType determines file type.
func (fi *Sfile) DetermineContentType(filePath string) (contentType ContentType, err error) {
	var file *os.File

	if filePath == "" {
		fi.logger.Out(logrus.WarnLevel, nil, "Empty file path.")
		return UnknownType, err
	}

	if file, err = os.Open(filePath); err != nil {
		// TODO: add Logger interface methods.
		// log.Warning("Invalid file path.", err)
		// fi.logger.Out(logrus.WarnLevel, "Invalid file path.", err)
		fi.logger.Out(logrus.WarnLevel, nil, "Invalid file path.")
		return UnknownType, err
	}

	defer file.Close()

	buff := make([]byte, 512)

	if _, err = file.Read(buff); err != nil {
		fi.logger.Out(logrus.WarnLevel, logrus.Fields{"filePath": filePath, "error": err}, "Cannot read the file.")
		return UnknownType, err
	}

	fileType := http.DetectContentType(buff)
	fi.logger.Out(logrus.DebugLevel, logrus.Fields{"fileType": fileType}, "Determined file content type.")

	switch fileType {
	case "application/zip":
		contentType = ZipType
	case "application/x-gzip":
		contentType = GzipType
	case "application/x-rar-compressed":
		contentType = XrarType
	case "video/webm":
		contentType = WebmType
	case "font/collection":
		contentType = TtcfType
	case "font/otf":
		contentType = OtfType
	case "font/woff":
		contentType = WoffType
	case "font/woff2":
		contentType = Woff2Type
	case "application/wasm":
		contentType = WasmType
	case "font/ttf":
		contentType = TtfType
	default:
		contentType = UnknownType
	}

	return contentType, nil
}

// BuildFilePath builds file path using directory and filename.
func (fi *Sfile) BuildFilePath(dir string, filename string) (filePath string, err error) {
	var file *os.File
	filePath = filepath.Join(dir, filename)
	_, err = os.Stat(filePath)

	if os.IsNotExist(err) {
		if file, err = os.Create(filePath); err != nil {
			err = errors.Wrap(err, "os.Create(filePath)")
			return
		}
		defer os.Remove(filePath)
		defer file.Close()
	}

	if err != nil {
		err = errors.Wrap(err, "os.Stat(filePath)")
		return
	}

	return filePath, nil
}

// CleanDir removes all subdirectories and files from the directory.
func (fi *Sfile) CleanDir(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

// ReadFile calls io.ReadFile to read the file named by filename.
func (fi *Sfile) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// ReadArchive calls io.ReadArchive to read the file named by filename.
func (fi *Sfile) ReadArchive(filename string) ([]byte, error) {
	return util.ReadArchive(filename)
}

// Mkdir delegates the call to create file directory.
func (fi *Sfile) Mkdir(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}

// CreateFile calls os.Create to create the file.
func (fi *Sfile) CreateFile(name string) (*os.File, error) {
	return os.Create(name)
}

// CloseFile delegages the call to close the file.
func (fi *Sfile) CloseFile(f *os.File) error {
	return f.Close()
}

// WriteFile delegates the call to write the file.
func (fi *Sfile) WriteFile(f *os.File, buf []byte) (n int, err error) {
	return f.Write(buf)
}

// RemoveAll removes path and any children it contains using os.RemoveAll
func (fi *Sfile) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

// MkdirAll creates a directory named path using os.MkdirAll
func (fi *Sfile) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}
