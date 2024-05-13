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

	"github.com/agiledragon/gomonkey/v2"
)

var patchesMap = make(map[string]*gomonkey.Patches)

// Patch replaces a function with another
func Patch(target, replacement interface{}) *gomonkey.Patches {
	key := reflect.TypeOf(target).String()
	existingPatches, ok := patchesMap[key]
	if ok {
		existingPatches.Reset()
	}
	patches := gomonkey.ApplyFunc(target, replacement)
	patchesMap[key] = patches
	return patches
}

// Unpatch unpatch a patch
func Unpatch(target interface{}) bool {
	patches, ok := patchesMap[reflect.TypeOf(target).String()]
	if !ok {
		return false
	}
	patches.Reset()
	delete(patchesMap, reflect.TypeOf(target).String())
	return true
}

// PatchInstanceMethod replaces an instance method methodName for the type target with replacement
func PatchInstanceMethod(target reflect.Type, methodName string, replacement interface{}) *gomonkey.Patches {
	key := target.String() + methodName
	existingPatches, ok := patchesMap[key]
	if ok {
		existingPatches.Reset()
	}
	patches := gomonkey.ApplyMethod(target, methodName, replacement)
	patchesMap[key] = patches
	return patches
}

// UnpatchInstanceMethod unpatch a patch
func UnpatchInstanceMethod(target reflect.Type, methodName string) bool {
	patches, ok := patchesMap[target.String()+methodName]
	if !ok {
		return false
	}
	patches.Reset()
	delete(patchesMap, target.String()+methodName)
	return true
}

// UnpatchAll unpatch all patches
func UnpatchAll() {
	for _, patches := range patchesMap {
		patches.Reset()
	}
	patchesMap = make(map[string]*gomonkey.Patches)
}
