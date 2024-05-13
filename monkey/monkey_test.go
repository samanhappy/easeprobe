/*
 * Copyright (c) 2022, MegaEase
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package monkey

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyStruct struct{}

func (s *MyStruct) Method() string {
	return "original"
}

func TestPatch(t *testing.T) {
	originalFunc := func() string { return "original" }
	replacementFunc := func() string { return "replacement" }

	patch := Patch(originalFunc, replacementFunc)
	defer patch.Reset()

	assert.Equal(t, "replacement", originalFunc())
}

func TestUnpatch(t *testing.T) {
	originalFunc := func() string { return "original" }
	replacementFunc := func() string { return "replacement" }

	Patch(originalFunc, replacementFunc)
	Unpatch(originalFunc)

	assert.Equal(t, "original", originalFunc())
}

func TestPatchInstanceMethod(t *testing.T) {
	PatchInstanceMethod(reflect.TypeOf(&MyStruct{}), "Method", func(*MyStruct) string { return "replacement" })
	defer UnpatchInstanceMethod(reflect.TypeOf(&MyStruct{}), "Method")

	assert.Equal(t, "replacement", (&MyStruct{}).Method())
}

func TestUnpatchInstanceMethod(t *testing.T) {
	PatchInstanceMethod(reflect.TypeOf(&MyStruct{}), "Method", func(*MyStruct) string { return "replacement" })
	UnpatchInstanceMethod(reflect.TypeOf(&MyStruct{}), "Method")

	assert.Equal(t, "original", (&MyStruct{}).Method())
}

func TestUnpatchAll(t *testing.T) {
	originalFunc := func() string { return "original" }
	replacementFunc := func() string { return "replacement" }

	Patch(originalFunc, replacementFunc)
	PatchInstanceMethod(reflect.TypeOf(&MyStruct{}), "Method", func(*MyStruct) string { return "replacement" })
	UnpatchAll()

	assert.Equal(t, "original", originalFunc())
	assert.Equal(t, "original", (&MyStruct{}).Method())
}
