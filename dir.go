// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package ach

import (
	"fmt"
	"os"
	"path/filepath"
)

// ReadDir will attempt to parse all ACH files in the given directory. Only files which
// parse successfully will be returned.
func ReadDir(dir string) ([]*File, error) {
	readACH := func(path string) (*File, error) {
		fd, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("opening %s failed: %v", path, err)
		}
		defer fd.Close()

		f, err := NewReader(fd).Read()
		if err != nil {
			return nil, fmt.Errorf("reading %s failed: %v", path, err)
		}
		return &f, nil
	}

	readJSON := func(path string) (*File, error) {
		bs, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("opening %s failed: %v", path, err)
		}
		return FileFromJSON(bs)
	}

	infos, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	out := make([]*File, 0, len(infos))
	for i := range infos {
		path := filepath.Join(dir, infos[i].Name())

		stat, err := os.Stat(path)
		if err != nil {
			return nil, fmt.Errorf("stat of %s failed: %v", path, err)
		}
		if stat.IsDir() {
			continue
		}

		f, err1 := readACH(path)
		if f != nil {
			out = append(out, f)
			continue
		}
		f, err2 := readJSON(path)
		if f != nil {
			out = append(out, f)
			continue
		}

		if err1 != nil && err2 != nil {
			return out, fmt.Errorf("%s failed to parse: %v | %v", path, err1, err2)
		}
	}
	return out, nil
}
