package catwalk

import "testing"

func TestResolveImageLimits_OverrideAndDefaults(t *testing.T) {
	// A per-model override fills only MaxLongEdge; the rest come from
	// the provider-family default.
	m := Model{Image: ImageLimits{MaxLongEdge: 2576}}
	got := m.ResolveImageLimits(TypeAnthropic)

	def := DefaultImageLimits(TypeAnthropic)
	if got.MaxLongEdge != 2576 {
		t.Errorf("MaxLongEdge = %d, want override 2576", got.MaxLongEdge)
	}
	if got.MaxBytesPerImage != def.MaxBytesPerImage {
		t.Errorf("MaxBytesPerImage = %d, want default %d", got.MaxBytesPerImage, def.MaxBytesPerImage)
	}
	if got.MaxImages != def.MaxImages {
		t.Errorf("MaxImages = %d, want default %d", got.MaxImages, def.MaxImages)
	}
}

func TestResolveImageLimits_NoOverrideUsesDefault(t *testing.T) {
	m := Model{}
	if got, want := m.ResolveImageLimits(TypeOpenAI), DefaultImageLimits(TypeOpenAI); got != want {
		t.Errorf("ResolveImageLimits = %+v, want %+v", got, want)
	}
}

func TestDefaultImageLimits_UnknownTypeIsConservative(t *testing.T) {
	if got := DefaultImageLimits(Type("not-a-real-type")); got != conservativeImageLimits {
		t.Errorf("unknown type = %+v, want conservative %+v", got, conservativeImageLimits)
	}
}

func TestImageLimits_IsZero(t *testing.T) {
	if !(ImageLimits{}).IsZero() {
		t.Error("empty ImageLimits should be zero")
	}
	if (ImageLimits{MaxImages: 1}).IsZero() {
		t.Error("non-empty ImageLimits should not be zero")
	}
}
