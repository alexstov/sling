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

import "os"

// Filer sends requests to the endpoint.
type Filer interface {
	CreateFile(name string) (*os.File, error)
	CloseFile(f *os.File) error
	WriteFile(f *os.File, b []byte) (n int, err error)
	ReadFile(filename string) ([]byte, error)
	ReadArchive(filename string) ([]byte, error)
	DetermineContentType(filePath string) (contentType ContentType, err error)
	BuildFilePath(dir string, filename string) (filePath string, err error)
	Mkdir(name string, perm os.FileMode) error
	CleanDir(dir string) error
	RemoveAll(path string) error
	MkdirAll(path string, perm os.FileMode) error
}
