// MIT License
//
// Copyright (c) 2016-2018 GACHAIN
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"fmt"
	"path/filepath"

	"github.com/andlabs/ui"
)

// SimpleGui - using if main program run without arguments
func SimpleGui() {
	err := ui.Main(func() {
		packMsg := fmt.Sprintf("Pack: select any file in source dir")
		unpackMsg := fmt.Sprintf("Unpack: select source file")
		btnPack := ui.NewButton(packMsg)
		btnUnpack := ui.NewButton(unpackMsg)
		box := ui.NewHorizontalBox()
		box.Append(btnPack, true)
		box.Append(btnUnpack, true)
		window := ui.NewWindow(currentTitle, 300, 100, false)
		window.SetMargined(true)
		window.SetChild(box)
		btnPack.OnClicked(func(*ui.Button) {
			packMode = true
			wSelectFile := ui.NewWindow("select dir", 300, 100, false)
			inputName = ui.OpenFile(wSelectFile)

			if inputName != "" {
				absPath, _ := filepath.Abs(inputName)
				inputName = filepath.Dir(absPath) + separator
				checkOutput()
				ui.Quit()
			}
		})
		btnUnpack.OnClicked(func(*ui.Button) {
			unpackMode = true
			wSelectFile := ui.NewWindow("select file", 300, 100, false)
			inputName = ui.OpenFile(wSelectFile)

			if inputName != "" {
				checkOutput()
				ui.Quit()
			}
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
