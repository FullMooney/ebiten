// Copyright 2016 Hajime Hoshi
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

// +build darwin
// +build !js
// +build !ios

package ui

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
//
// #import <AppKit/AppKit.h>
//
// static void currentMonitorPos(int* x, int* y) {
//   NSDictionary* screenDictionary = [[NSScreen mainScreen] deviceDescription];
//   NSNumber* screenID = [screenDictionary objectForKey:@"NSScreenNumber"];
//   CGDirectDisplayID aID = [screenID unsignedIntValue];
//   const CGRect bounds = CGDisplayBounds(aID);
//   *x = bounds.origin.x;
//   *y = bounds.origin.y;
// }
import "C"

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

func glfwScale() float64 {
	return 1
}

func adjustWindowPosition(x, y int) (int, int) {
	return x, y
}

func currentMonitor() *glfw.Monitor {
	x := C.int(0)
	y := C.int(0)
	C.currentMonitorPos(&x, &y)
	for _, m := range glfw.GetMonitors() {
		mx, my := m.GetPos()
		if int(x) == mx && int(y) == my {
			return m
		}
	}
	return nil
}
