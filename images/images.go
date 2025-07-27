package images

import "github.com/hajimehoshi/ebiten/v2"

func Mirror(img *ebiten.Image) *ebiten.Image {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()
	mirrored := ebiten.NewImage(w, h)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(-1, 1)
	opts.GeoM.Translate(float64(w), 0)
	mirrored.DrawImage(img, opts)

	return mirrored
}
