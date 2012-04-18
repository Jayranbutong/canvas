/*
  Copyright (c) 2012 José Carlos Nieto, http://xiam.menteslibres.org/

  Permission is hereby granted, free of charge, to any person obtaining
  a copy of this software and associated documentation files (the
  "Software"), to deal in the Software without restriction, including
  without limitation the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the Software, and to
  permit persons to whom the Software is furnished to do so, subject to
  the following conditions:

  The above copyright notice and this permission notice shall be
  included in all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package main

/*
#cgo LDFLAGS: -lMagickWand -lMagickCore 
#cgo CFLAGS: -fopenmp -I/usr/include/ImageMagick  
#include <stdlib.h>
#include <wand/magick_wand.h>
*/
import "C"

/*
import (
  "errors"
  "fmt"
)

type Errno int

func (e Errno) Error() string {
  s := errText[e]
  if s == "" {
    return fmt.Sprintf("errno %d", int(e))
  }
  return s
}
*/

type Canvas struct {
  wand *C.MagickWand
  filename string
  width string
  height string
}

func (c Canvas) init() {
  C.MagickWandGenesis()
}

func (c Canvas) Open(filename string) (bool) {
  status := C.MagickReadImage(c.wand, C.CString(filename))
  if status == C.MagickFalse {
    return false
  }
  return true
}

func (c Canvas) Write(filename string) (bool) {
  status := C.MagickWriteImage(c.wand, C.CString(filename))
  if status == C.MagickFalse {
    return false
  }
  return true
}

func (c Canvas) Destroy() {
  if c.wand != nil {
    C.DestroyMagickWand(c.wand)
  }
  C.MagickWandTerminus()
}

func NewCanvas() *Canvas {
  c := &Canvas{}
  c.init()
  c.wand = C.NewMagickWand()
  return c
}

func main() {

  canvas := NewCanvas()

  opened := canvas.Open("example.jpg")

  if opened {
    canvas.Write("example-go.png")
  }

  canvas.Destroy()

}
