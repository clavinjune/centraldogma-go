// Copyright 2018 LINE Corporation
//
// LINE Corporation licenses this file to you under the Apache License,
// version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at:
//
//   https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"flag"
	"reflect"
	"testing"

	"github.com/urfave/cli"
)

func newParentContext(connectURL string) *cli.Context {
	parentFlags := flag.NewFlagSet("test", 0)
	parentFlags.String("connect", connectURL, "")
	return cli.NewContext(nil, parentFlags, nil)
}
func newContext(flagArguments []string, connectURL, revision string) *cli.Context {
	parent := newParentContext(connectURL)

	flags := flag.FlagSet{}
	flags.Parse(flagArguments)
	flags.String("revision", revision, "")
	return cli.NewContext(nil, &flags, parent)
}

func newGetCmdContext(flagArguments []string, connectURL, revision string, isRecursive bool) *cli.Context {
	parent := newParentContext(connectURL)

	flags := flag.FlagSet{}
	flags.Parse(flagArguments)
	flags.String("revision", revision, "")
	flags.Bool("recursive", isRecursive, "")
	return cli.NewContext(nil, &flags, parent)
}

func TestSplitPath(t *testing.T) {
	var tests = []struct {
		path string
		want []string
	}{
		{"", nil},
		{"/", nil},
		{"/foo/bar", []string{"foo", "bar"}},
		{"/foo/bar/", []string{"foo", "bar"}},
		{"foo/bar/a.txt", []string{"foo", "bar", "/a.txt"}},
		{"/foo/bar/b.txt", []string{"foo", "bar", "/b.txt"}},
		{"//foo//bar//c.txt", []string{"foo", "bar", "/c.txt"}},
		{"/foo/bar/a/", []string{"foo", "bar", "/a/"}},
		{"/foo/bar/a/d.txt", []string{"foo", "bar", "/a/d.txt"}},
		{"/foo/bar/a/b/e.txt", []string{"foo", "bar", "/a/b/e.txt"}},
		{"/foo/bar/a/b//f.txt", []string{"foo", "bar", "/a/b/f.txt"}},
	}
	for _, test := range tests {
		if got := splitPath(test.path); !reflect.DeepEqual(got, test.want) {
			t.Errorf("splitPath(%q) = %q, want:%q", test.path, got, test.want)
		}
	}
}
