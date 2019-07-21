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
	"encoding/json"
	"io"

	"github.com/kubernetes/pkg/printers"
	"github.com/kubernetes/staging/src/k8s.io/apiserver/pkg/util/flag"
	"github.com/kubernetes/staging/src/k8s.io/cli-runtime/pkg/genericclioptions"
	// "k8s.io/apimachinery/pkg/runtime"
	// "sigs.k8s.io/yaml"
)

// type Object interface {
// 	GetObjectKind() schema.ObjectKind
// 	DeepCopyObject() Object
// }

// JSONPrinter is an implementation of ResourcePrinter which outputs an object as JSON.
type JSONPrinter struct{}

// PrintObj is an implementation of ResourcePrinter.PrintObj which simply writes the object to the Writer.
func (p *JSONPrinter) PrintObj(obj interface{}, w io.Writer) error {
	// we use reflect.Indirect here in order to obtain the actual value from a pointer.
	// we need an actual value in order to retrieve the package path for an object.
	// using reflect.Indirect indiscriminately is valid here, as all runtime.Objects are supposed to be pointers.
	// if InternalObjectPreventer.IsForbidden(reflect.Indirect(reflect.ValueOf(obj)).Type().PkgPath()) {
	// 	return fmt.Errorf(InternalObjectPrinterErr)
	// }

	// switch obj := obj.(type) {
	// case *runtime.Unknown:
	// 	var buf bytes.Buffer
	// 	err := json.Indent(&buf, obj.Raw, "", "    ")
	// 	if err != nil {
	// 		return err
	// 	}
	// 	buf.WriteRune('\n')
	// 	_, err = buf.WriteTo(w)
	// 	return err
	// }

	// if obj.GetObjectKind().GroupVersionKind().Empty() {
	// 	return fmt.Errorf("missing apiVersion or kind; try GetObjectKind().SetGroupVersionKind() if you know the type")
	// }

	data, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	_, err = w.Write(data)
	return err
}

// YAMLPrinter is an implementation of ResourcePrinter which outputs an object as YAML.
// The input object is assumed to be in the internal version of an API and is converted
// to the given version first.
type YAMLPrinter struct{}

// ViewOptions holds the command-line options for 'config view' sub command
type ViewOptions struct {
	PrintFlags  *genericclioptions.PrintFlags
	PrintObject printers.ResourcePrinterFunc

	// ConfigAccess clientcmd.ConfigAccess
	Merge       flag.Tristate
	Flatten     bool
	Minify      bool
	RawByteData bool

	Context      string
	OutputFormat string

	genericclioptions.IOStreams
}
