// Copyright 2011 The Graphics-Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package graphics

import (
	"errors"
	"image"
	"image/draw"

	interp2 "gitlab.com/my-group322/pictures/pics-svc/pkg/imgediting/code.google.com/p/graphics-go/graphics/interp"
)

// RotateOptions are the rotation parameters.
// Angle is the angle, in radians, to rotate the image clockwise.
type RotateOptions struct {
	Angle float64
}

// Rotate produces a rotated version of src, drawn onto dst.
func Rotate(dst draw.Image, src image.Image, opt *RotateOptions) error {
	if dst == nil {
		return errors.New("graphics: dst is nil")
	}
	if src == nil {
		return errors.New("graphics: src is nil")
	}

	angle := 0.0
	if opt != nil {
		angle = opt.Angle
	}

	return I.Rotate(angle).TransformCenter(dst, src, interp2.Bilinear)
}