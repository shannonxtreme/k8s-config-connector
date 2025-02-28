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

//go:build integration
// +build integration

package embed

import (
	"net/http"
	"path"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"

	"github.com/shurcooL/vfsgen"
)

var (
	AssetsDir                  = http.Dir(repo.GetServiceMappingsPathOrLogFatal())
	Assets     http.FileSystem = AssetsDir
	outputPath                 = path.Join(repo.GetCallerPackagePathOrLogFatal(), "assets_vfsdata.go")
	// the options used by vfsgen (in assets_generate.go) when generating the data for Assets that will be embedded in the binary
	// the reason to keep the options in this file is then we have file without the build tag 'ignore' that has a dependency on
	// vfsgen. Without this dependency, 'go mod vendor' will not import the dependency into vendor.
	VfsgenOptions = vfsgen.Options{
		PackageName:  "embed",
		BuildTags:    "!integration",
		VariableName: "Assets",
		Filename:     outputPath,
	}
)
