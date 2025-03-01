// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stream

import (
	"fmt"
	"io"

	"github.com/ghodss/yaml"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	yamlSeperator = []byte("---\n")
	// YAML streams are terminated with "..." which signifies the end of transmission: https://yaml.org/spec/1.2/spec.html
	yamlTransmissionTerminator = []byte("...")
)

type UnstructuredStream interface {
	Next() (*unstructured.Unstructured, error)
}

type YAMLStream struct {
	unstructuredStream UnstructuredStream
	nextBytes          []byte
	nextUnstructured   *unstructured.Unstructured
	nextErr            error
	// true if the consumer of the stream has read at least one non-error result
	returnedAtLeastOneNonErrorResult bool
	// true if the end of the stream was reached and the tranmission terminator was returned
	returnedTransmissionTerminator bool
}

func NewYAMLStream(unstructuredStream UnstructuredStream) *YAMLStream {
	yamlStream := YAMLStream{
		unstructuredStream: unstructuredStream,
	}
	return &yamlStream
}

func (y *YAMLStream) Next() ([]byte, *unstructured.Unstructured, error) {
	if y.nextErr == nil && y.nextBytes == nil {
		// this occurs on the first call to Next() or AFTER an error, while putting a fillNext(...) in the
		// NewYAMLStream(...) would result in cleaner code it would mean that NewYAMLStream(...) could take a "long time"
		// while contacting GCP to get the first unstructured which could result in some undesireable user experiences
		y.fillNext()
	}
	// if this is EOF and we have not YET returned the terminator AND we wrote at least one result, return "...", otherwise, return EOF
	bytes, unstructured, err := y.nextBytes, y.nextUnstructured, y.nextErr
	if err != nil {
		// if the error is EOF and we have not yet returned the YAML terminator, '...', then return it
		if err == io.EOF {
			if !y.returnedTransmissionTerminator {
				if y.returnedAtLeastOneNonErrorResult {
					y.returnedTransmissionTerminator = true
					return yamlTransmissionTerminator, nil, nil
				}
			}
		}
		y.nextBytes = nil
		y.nextUnstructured = nil
		y.nextErr = nil
		return nil, nil, err
	}
	bytes = append(yamlSeperator, bytes...)
	y.fillNext()
	y.returnedAtLeastOneNonErrorResult = true
	return bytes, unstructured, nil
}

func (y *YAMLStream) fillNext() {
	y.nextBytes, y.nextUnstructured, y.nextErr = y.getNext()
}

func (y *YAMLStream) getNext() ([]byte, *unstructured.Unstructured, error) {
	unstructured, err := y.unstructuredStream.Next()
	if err != nil {
		if err != io.EOF {
			err = fmt.Errorf("error getting unstructured: %v", err)
		}
		return nil, unstructured, err
	}
	// the status field does not make sense for output as these YAMLs should be used in a git workflow
	delete(unstructured.Object, "status")
	bytes, err := yaml.Marshal(unstructured.Object)
	if err != nil {
		return nil, unstructured, fmt.Errorf("error marshalling unstructured to YAML: %v", err)
	}
	return bytes, unstructured, nil
}
