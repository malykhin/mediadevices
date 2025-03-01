package video

import (
	"image"
	"reflect"
	"testing"
)

func TestScale(t *testing.T) {
	cases := map[string]struct {
		src           image.Image
		width, height int
		expected      image.Image
	}{
		"RGBA": {
			src: &image.RGBA{
				Pix: []uint8{
					// R     G     B     A |   R     G     B     A |   R     G     B     A |   R     G     B     A
					0x80, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF,
					0x80, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF,
					0x00, 0x40, 0x00, 0xFF, 0x00, 0x40, 0x00, 0xFF, 0x00, 0x00, 0x60, 0xFF, 0x00, 0x00, 0x60, 0xFF,
					0x00, 0x40, 0x00, 0xFF, 0x00, 0x40, 0x00, 0xFF, 0x00, 0x00, 0x60, 0xFF, 0x00, 0x00, 0x60, 0xFF,
				},
				Stride: 16,
				Rect:   image.Rect(0, 0, 4, 4),
			},
			width:  2,
			height: 2,
			expected: &image.RGBA{
				Pix: []uint8{
					// R     G     B     A |   R     G     B     A
					0x80, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF,
					0x00, 0x40, 0x00, 0xFF, 0x00, 0x00, 0x60, 0xFF,
				},
				Stride: 8,
				Rect:   image.Rect(0, 0, 2, 2),
			},
		},
		"RGBASameSize": {
			src: &image.RGBA{
				Pix: []uint8{
					// R     G     B     A |   R     G     B     A |   R     G     B     A |   R     G     B     A
					0x80, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF,
					0x80, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF,
					0x00, 0x40, 0x00, 0xFF, 0x00, 0x40, 0x00, 0xFF, 0x00, 0x00, 0x60, 0xFF, 0x00, 0x00, 0x60, 0xFF,
					0x00, 0x40, 0x00, 0xFF, 0x00, 0x40, 0x00, 0xFF, 0x00, 0x00, 0x60, 0xFF, 0x00, 0x00, 0x60, 0xFF,
				},
				Stride: 16,
				Rect:   image.Rect(0, 0, 4, 4),
			},
			width:  4,
			height: 4,
			expected: &image.RGBA{
				Pix: []uint8{
					// R     G     B     A |   R     G     B     A |   R     G     B     A |   R     G     B     A
					0x80, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF,
					0x80, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0xFF,
					0x00, 0x40, 0x00, 0xFF, 0x00, 0x40, 0x00, 0xFF, 0x00, 0x00, 0x60, 0xFF, 0x00, 0x00, 0x60, 0xFF,
					0x00, 0x40, 0x00, 0xFF, 0x00, 0x40, 0x00, 0xFF, 0x00, 0x00, 0x60, 0xFF, 0x00, 0x00, 0x60, 0xFF,
				},
				Stride: 16,
				Rect:   image.Rect(0, 0, 4, 4),
			},
		},
		"I444": {
			src: &image.YCbCr{
				SubsampleRatio: image.YCbCrSubsampleRatio444,
				Y: []uint8{
					0xF0, 0xF0, 0x00, 0x00, 0x00, 0x00,
					0xF0, 0xF0, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40,
					0x00, 0x00, 0x80, 0x80, 0x00, 0x00,
					0x00, 0x00, 0x80, 0x80, 0x00, 0x00,
				},
				Cb: []uint8{
					0x20, 0x20, 0x80, 0x80, 0x80, 0x80,
					0x20, 0x20, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0xC0, 0xC0,
					0x80, 0x80, 0x80, 0x80, 0xC0, 0xC0,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
				},
				Cr: []uint8{
					0xE0, 0xE0, 0x80, 0x80, 0x80, 0x80,
					0xE0, 0xE0, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x40, 0x40, 0x80, 0x80,
					0x80, 0x80, 0x40, 0x40, 0x80, 0x80,
				},
				YStride: 6,
				CStride: 6,
				Rect:    image.Rect(0, 0, 6, 6),
			},
			width:  3,
			height: 3,
			expected: &image.YCbCr{
				SubsampleRatio: image.YCbCrSubsampleRatio444,
				Y: []uint8{
					0xF0, 0x00, 0x00,
					0x00, 0x00, 0x40,
					0x00, 0x80, 0x00,
				},
				Cb: []uint8{
					0x20, 0x80, 0x80,
					0x80, 0x80, 0xC0,
					0x80, 0x80, 0x80,
				},
				Cr: []uint8{
					0xE0, 0x80, 0x80,
					0x80, 0x80, 0x80,
					0x80, 0x40, 0x80,
				},
				YStride: 3,
				CStride: 3,
				Rect:    image.Rect(0, 0, 3, 3),
			},
		},
		"I444SameSize": {
			src: &image.YCbCr{
				SubsampleRatio: image.YCbCrSubsampleRatio444,
				Y: []uint8{
					0xF0, 0xF0, 0x00, 0x00, 0x00, 0x00,
					0xF0, 0xF0, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40,
					0x00, 0x00, 0x80, 0x80, 0x00, 0x00,
					0x00, 0x00, 0x80, 0x80, 0x00, 0x00,
				},
				Cb: []uint8{
					0x20, 0x20, 0x80, 0x80, 0x80, 0x80,
					0x20, 0x20, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0xC0, 0xC0,
					0x80, 0x80, 0x80, 0x80, 0xC0, 0xC0,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
				},
				Cr: []uint8{
					0xE0, 0xE0, 0x80, 0x80, 0x80, 0x80,
					0xE0, 0xE0, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x40, 0x40, 0x80, 0x80,
					0x80, 0x80, 0x40, 0x40, 0x80, 0x80,
				},
				YStride: 6,
				CStride: 6,
				Rect:    image.Rect(0, 0, 6, 6),
			},
			width:  6,
			height: 6,
			expected: &image.YCbCr{
				SubsampleRatio: image.YCbCrSubsampleRatio444,
				Y: []uint8{
					0xF0, 0xF0, 0x00, 0x00, 0x00, 0x00,
					0xF0, 0xF0, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40,
					0x00, 0x00, 0x80, 0x80, 0x00, 0x00,
					0x00, 0x00, 0x80, 0x80, 0x00, 0x00,
				},
				Cb: []uint8{
					0x20, 0x20, 0x80, 0x80, 0x80, 0x80,
					0x20, 0x20, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0xC0, 0xC0,
					0x80, 0x80, 0x80, 0x80, 0xC0, 0xC0,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
				},
				Cr: []uint8{
					0xE0, 0xE0, 0x80, 0x80, 0x80, 0x80,
					0xE0, 0xE0, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x40, 0x40, 0x80, 0x80,
					0x80, 0x80, 0x40, 0x40, 0x80, 0x80,
				},
				YStride: 6,
				CStride: 6,
				Rect:    image.Rect(0, 0, 6, 6),
			},
		},
		"I422": {
			src: &image.YCbCr{
				SubsampleRatio: image.YCbCrSubsampleRatio422,
				Y: []uint8{
					0xF0, 0xF0, 0x10, 0x10, 0x00, 0x00, 0x00, 0x00,
					0xF0, 0xF0, 0x10, 0x10, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x80, 0x80, 0x30, 0x30, 0x00, 0x00,
					0x00, 0x00, 0x80, 0x80, 0x30, 0x30, 0x00, 0x00,
				},
				Cb: []uint8{
					0x20, 0x20, 0x80, 0x80,
					0x20, 0x20, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0xE0, 0xE0,
					0x80, 0x80, 0xE0, 0xE0,
				},
				Cr: []uint8{
					0xE0, 0xE0, 0x80, 0x80,
					0xE0, 0xE0, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0xF0, 0xF0, 0x40, 0x40,
					0xF0, 0xF0, 0x40, 0x40,
				},
				YStride: 8,
				CStride: 4,
				Rect:    image.Rect(0, 0, 8, 8),
			},
			width:  4,
			height: 4,
			expected: &image.YCbCr{
				SubsampleRatio: image.YCbCrSubsampleRatio422,
				Y: []uint8{
					0xF0, 0x10, 0x00, 0x00,
					0x00, 0x00, 0x40, 0x00,
					0x00, 0x00, 0x00, 0x00,
					0x00, 0x80, 0x30, 0x00,
				},
				Cb: []uint8{
					0x20, 0x80,
					0x80, 0x80,
					0x80, 0x80,
					0x80, 0xE0,
				},
				Cr: []uint8{
					0xE0, 0x80,
					0x80, 0x80,
					0x80, 0x80,
					0xF0, 0x40,
				},
				YStride: 4,
				CStride: 2,
				Rect:    image.Rect(0, 0, 4, 4),
			},
		},
		"I422SameSize": {
			src: &image.YCbCr{
				SubsampleRatio: image.YCbCrSubsampleRatio422,
				Y: []uint8{
					0xF0, 0xF0, 0x10, 0x10, 0x00, 0x00, 0x00, 0x00,
					0xF0, 0xF0, 0x10, 0x10, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x80, 0x80, 0x30, 0x30, 0x00, 0x00,
					0x00, 0x00, 0x80, 0x80, 0x30, 0x30, 0x00, 0x00,
				},
				Cb: []uint8{
					0x20, 0x20, 0x80, 0x80,
					0x20, 0x20, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0xE0, 0xE0,
					0x80, 0x80, 0xE0, 0xE0,
				},
				Cr: []uint8{
					0xE0, 0xE0, 0x80, 0x80,
					0xE0, 0xE0, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0xF0, 0xF0, 0x40, 0x40,
					0xF0, 0xF0, 0x40, 0x40,
				},
				YStride: 8,
				CStride: 4,
				Rect:    image.Rect(0, 0, 8, 8),
			},
			width:  8,
			height: 8,
			expected: &image.YCbCr{
				SubsampleRatio: image.YCbCrSubsampleRatio422,
				Y: []uint8{
					0xF0, 0xF0, 0x10, 0x10, 0x00, 0x00, 0x00, 0x00,
					0xF0, 0xF0, 0x10, 0x10, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x80, 0x80, 0x30, 0x30, 0x00, 0x00,
					0x00, 0x00, 0x80, 0x80, 0x30, 0x30, 0x00, 0x00,
				},
				Cb: []uint8{
					0x20, 0x20, 0x80, 0x80,
					0x20, 0x20, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0xE0, 0xE0,
					0x80, 0x80, 0xE0, 0xE0,
				},
				Cr: []uint8{
					0xE0, 0xE0, 0x80, 0x80,
					0xE0, 0xE0, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
					0xF0, 0xF0, 0x40, 0x40,
					0xF0, 0xF0, 0x40, 0x40,
				},
				YStride: 8,
				CStride: 4,
				Rect:    image.Rect(0, 0, 8, 8),
			},
		},
		"I420": {
			src: &image.YCbCr{
				SubsampleRatio: image.YCbCrSubsampleRatio420,
				Y: []uint8{
					0xF0, 0xF0, 0x10, 0x10, 0x00, 0x00, 0x00, 0x00,
					0xF0, 0xF0, 0x10, 0x10, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x80, 0x80, 0x30, 0x30, 0x00, 0x00,
					0x00, 0x00, 0x80, 0x80, 0x30, 0x30, 0x00, 0x00,
				},
				Cb: []uint8{
					0x20, 0x20, 0x80, 0x80,
					0x20, 0x20, 0x80, 0x80,
					0x80, 0x80, 0xE0, 0xE0,
					0x80, 0x80, 0xE0, 0xE0,
				},
				Cr: []uint8{
					0xE0, 0xE0, 0x80, 0x80,
					0xE0, 0xE0, 0x80, 0x80,
					0xF0, 0xF0, 0x40, 0x40,
					0xF0, 0xF0, 0x40, 0x40,
				},
				YStride: 8,
				CStride: 4,
				Rect:    image.Rect(0, 0, 8, 8),
			},
			width:  4,
			height: 4,
			expected: &image.YCbCr{
				SubsampleRatio: image.YCbCrSubsampleRatio420,
				Y: []uint8{
					0xF0, 0x10, 0x00, 0x00,
					0x00, 0x00, 0x40, 0x00,
					0x00, 0x00, 0x00, 0x00,
					0x00, 0x80, 0x30, 0x00,
				},
				Cb: []uint8{
					0x20, 0x80,
					0x80, 0xE0,
				},
				Cr: []uint8{
					0xE0, 0x80,
					0xF0, 0x40,
				},
				YStride: 4,
				CStride: 2,
				Rect:    image.Rect(0, 0, 4, 4),
			},
		},
		"I420SameSize": {
			src: &image.YCbCr{
				SubsampleRatio: image.YCbCrSubsampleRatio420,
				Y: []uint8{
					0xF0, 0xF0, 0x10, 0x10, 0x00, 0x00, 0x00, 0x00,
					0xF0, 0xF0, 0x10, 0x10, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x80, 0x80, 0x30, 0x30, 0x00, 0x00,
					0x00, 0x00, 0x80, 0x80, 0x30, 0x30, 0x00, 0x00,
				},
				Cb: []uint8{
					0x20, 0x20, 0x80, 0x80,
					0x20, 0x20, 0x80, 0x80,
					0x80, 0x80, 0xE0, 0xE0,
					0x80, 0x80, 0xE0, 0xE0,
				},
				Cr: []uint8{
					0xE0, 0xE0, 0x80, 0x80,
					0xE0, 0xE0, 0x80, 0x80,
					0xF0, 0xF0, 0x40, 0x40,
					0xF0, 0xF0, 0x40, 0x40,
				},
				YStride: 8,
				CStride: 4,
				Rect:    image.Rect(0, 0, 8, 8),
			},
			width:  8,
			height: 8,
			expected: &image.YCbCr{
				SubsampleRatio: image.YCbCrSubsampleRatio420,
				Y: []uint8{
					0xF0, 0xF0, 0x10, 0x10, 0x00, 0x00, 0x00, 0x00,
					0xF0, 0xF0, 0x10, 0x10, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x80, 0x80, 0x30, 0x30, 0x00, 0x00,
					0x00, 0x00, 0x80, 0x80, 0x30, 0x30, 0x00, 0x00,
				},
				Cb: []uint8{
					0x20, 0x20, 0x80, 0x80,
					0x20, 0x20, 0x80, 0x80,
					0x80, 0x80, 0xE0, 0xE0,
					0x80, 0x80, 0xE0, 0xE0,
				},
				Cr: []uint8{
					0xE0, 0xE0, 0x80, 0x80,
					0xE0, 0xE0, 0x80, 0x80,
					0xF0, 0xF0, 0x40, 0x40,
					0xF0, 0xF0, 0x40, 0x40,
				},
				YStride: 8,
				CStride: 4,
				Rect:    image.Rect(0, 0, 8, 8),
			},
		},
		"I420NonSquareImage": {
			src: &image.YCbCr{
				SubsampleRatio: image.YCbCrSubsampleRatio420,
				Y: []uint8{
					0xF0, 0xF0, 0x10, 0x10, 0x00, 0x00, 0x00, 0x00, 0xF0, 0xF0, 0x10, 0x10,
					0xF0, 0xF0, 0x10, 0x10, 0x00, 0x00, 0x00, 0x00, 0xF0, 0xF0, 0x10, 0x10,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x80, 0x80, 0x30, 0x30, 0x00, 0x00, 0x30, 0x30, 0x00, 0x00,
					0x00, 0x00, 0x80, 0x80, 0x30, 0x30, 0x00, 0x00, 0x30, 0x30, 0x00, 0x00,
				},
				Cb: []uint8{
					0x20, 0x20, 0x80, 0x80, 0x50, 0x50,
					0x20, 0x20, 0x80, 0x80, 0x50, 0x50,
					0x80, 0x80, 0xE0, 0xE0, 0x30, 0x30,
					0x80, 0x80, 0xE0, 0xE0, 0x30, 0x30,
				},
				Cr: []uint8{
					0xE0, 0xE0, 0x80, 0x80, 0xB0, 0xB0,
					0xE0, 0xE0, 0x80, 0x80, 0xB0, 0xB0,
					0xF0, 0xF0, 0x40, 0x40, 0xC0, 0xC0,
					0xF0, 0xF0, 0x40, 0x40, 0xC0, 0xC0,
				},
				YStride: 12,
				CStride: 6,
				Rect:    image.Rect(0, 0, 12, 8),
			},
			width:  6,
			height: 4,
			expected: &image.YCbCr{
				SubsampleRatio: image.YCbCrSubsampleRatio420,
				Y: []uint8{
					0xF0, 0x10, 0x00, 0x00, 0xF0, 0x10,
					0x00, 0x00, 0x40, 0x00, 0x40, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x80, 0x30, 0x00, 0x30, 0x00,
				},
				Cb: []uint8{
					0x20, 0x80, 0x50,
					0x80, 0xE0, 0x30,
				},
				Cr: []uint8{
					0xE0, 0x80, 0xB0,
					0xF0, 0x40, 0xC0,
				},
				YStride: 6,
				CStride: 3,
				Rect:    image.Rect(0, 0, 6, 4),
			},
		},
	}
	for name, algo := range scalerTestAlgos {
		algo := algo
		t.Run(name, func(t *testing.T) {
			for name, c := range cases {
				c := c
				t.Run(name, func(t *testing.T) {
					trans := Scale(c.width, c.height, algo)
					r := trans(ReaderFunc(func() (image.Image, func(), error) {
						return c.src, func() {}, nil
					}))
					for i := 0; i < 4; i++ {
						out, _, err := r.Read()
						if err != nil {
							t.Fatalf("Unexpected error: %v", err)
						}
						if !reflect.DeepEqual(c.expected, out) {
							t.Errorf("Expected output image:\n%v\ngot:\n%v\nrepeat: %d", c.expected, out, i)
						}
						// Destroy output contents
						switch v := out.(type) {
						case *image.RGBA:
							v.Stride = 10
							v.Pix = v.Pix[:1]
							v.Rect.Max.X = 1
						case *image.YCbCr:
							v.YStride = 10
							v.CStride = 100
							v.Y = v.Y[:1]
							v.Cb = v.Cb[:2]
							v.Cr = v.Cr[:1]
							v.Rect.Max.X = 1
						}
					}
				})
			}
		})
	}
}

func TestScaleFastBoxSampling(t *testing.T) {
	if !hasCGOConvert {
		t.Skip("Skip: nocgo implementation is not supported for FastBoxSampling")
	}
	cases := map[string]struct {
		src           image.Image
		width, height int
		expected      image.Image
	}{
		"I420NonSquareImage": {
			src: &image.YCbCr{
				SubsampleRatio: image.YCbCrSubsampleRatio420,
				Y: []uint8{
					0xF0, 0xF0, 0x10, 0x10, 0x00, 0x00, 0x00, 0x00, 0xF0, 0xF0, 0x10, 0x10,
					0xF0, 0xF0, 0x10, 0x10, 0x00, 0x00, 0x00, 0x00, 0xF0, 0xF0, 0x10, 0x10,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00, 0x40, 0x40, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x80, 0x80, 0x30, 0x30, 0x00, 0x00, 0x30, 0x30, 0x00, 0x00,
					0x00, 0x00, 0x80, 0x80, 0x30, 0x30, 0x00, 0x00, 0x30, 0x30, 0x00, 0x00,
				},
				Cb: []uint8{
					0x20, 0x20, 0x80, 0x80, 0x50, 0x50,
					0x20, 0x20, 0x80, 0x80, 0x50, 0x50,
					0x80, 0x80, 0xE0, 0xE0, 0x30, 0x30,
					0x80, 0x80, 0xE0, 0xE0, 0x30, 0x30,
				},
				Cr: []uint8{
					0xE0, 0xE0, 0x80, 0x80, 0xB0, 0xB0,
					0xE0, 0xE0, 0x80, 0x80, 0xB0, 0xB0,
					0xF0, 0xF0, 0x40, 0x40, 0xC0, 0xC0,
					0xF0, 0xF0, 0x40, 0x40, 0xC0, 0xC0,
				},
				YStride: 12,
				CStride: 6,
				Rect:    image.Rect(0, 0, 12, 8),
			},
			width:  6,
			height: 4,
			expected: &image.YCbCr{
				SubsampleRatio: image.YCbCrSubsampleRatio420,
				Y: []uint8{
					0xF0, 0x80, 0x08, 0x00, 0x78, 0x80,
					0x08, 0x00, 0x20, 0x20, 0x20, 0x20,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x40, 0x58, 0x18, 0x18, 0x18,
				},
				Cb: []uint8{
					0x20, 0x50, 0x68,
					0x68, 0xB0, 0x88,
				},
				Cr: []uint8{
					0xE0, 0xB0, 0x98,
					0xD0, 0x98, 0x80,
				},
				YStride: 6,
				CStride: 3,
				Rect:    image.Rect(0, 0, 6, 4),
			},
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			trans := Scale(c.width, c.height, ScalerFastBoxSampling)
			r := trans(ReaderFunc(func() (image.Image, func(), error) {
				return c.src, func() {}, nil
			}))
			for i := 0; i < 4; i++ {
				out, _, err := r.Read()
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if !reflect.DeepEqual(c.expected, out) {
					t.Errorf("Expected output image:\n%v\ngot:\n%v\nrepeat: %d", c.expected, out, i)
				}
				// Destroy output contents
				switch v := out.(type) {
				case *image.RGBA:
					v.Stride = 10
					v.Pix = v.Pix[:1]
					v.Rect.Max.X = 1
				case *image.YCbCr:
					v.YStride = 10
					v.CStride = 100
					v.Y = v.Y[:1]
					v.Cb = v.Cb[:2]
					v.Cr = v.Cr[:1]
					v.Rect.Max.X = 1
				}
			}
		})
	}
}

func BenchmarkScale(b *testing.B) {
	for name, algo := range scalerBenchAlgos {
		algo := algo
		b.Run(name, func(b *testing.B) {
			for name, sz := range imageSizes {
				cases := map[string]image.Image{
					"RGBA": image.NewRGBA(image.Rect(0, 0, sz[0], sz[1])),
					"I444": image.NewYCbCr(image.Rect(0, 0, sz[0], sz[1]), image.YCbCrSubsampleRatio444),
				}
				b.Run(name, func(b *testing.B) {
					for name, img := range cases {
						img := img
						b.Run(name, func(b *testing.B) {
							trans := Scale(640, 360, algo)
							r := trans(ReaderFunc(func() (image.Image, func(), error) {
								return img, func() {}, nil
							}))

							for i := 0; i < b.N; i++ {
								_, _, err := r.Read()
								if err != nil {
									b.Fatalf("Unexpected error: %v", err)
								}
							}
						})
					}
				})
			}
		})
	}
}
