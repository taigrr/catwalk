package providers

import (
	"slices"
	"testing"
)

func TestValidDefaultModels(t *testing.T) {
	for _, p := range GetAll() {
		t.Run(p.Name, func(t *testing.T) {
			var modelIds []string
			for _, m := range p.Models {
				modelIds = append(modelIds, m.ID)
			}
			if !slices.Contains(modelIds, p.DefaultLargeModelID) {
				t.Errorf("Default large model %q not found in provider %q", p.DefaultLargeModelID, p.Name)
			}
			if !slices.Contains(modelIds, p.DefaultSmallModelID) {
				t.Errorf("Default small model %q not found in provider %q", p.DefaultSmallModelID, p.Name)
			}
			if p.DefaultEmbeddingModelID != "" {
				var embeddingIDs []string
				for _, m := range p.EmbeddingModels {
					embeddingIDs = append(embeddingIDs, m.ID)
				}
				if !slices.Contains(embeddingIDs, p.DefaultEmbeddingModelID) {
					t.Errorf("Default embedding model %q not found in provider %q", p.DefaultEmbeddingModelID, p.Name)
				}
			}
		})
	}
}

func TestAnthropicNewModelsImageLongEdgeOverride(t *testing.T) {
	want := map[string]int64{
		"claude-fable-5":    2576,
		"claude-mythos-5":   2576,
		"claude-opus-5":     2576,
		"claude-opus-4-8":   2576,
		"claude-opus-4-7":   2576,
		"claude-sonnet-4-6": 1568, // no override -> provider default
	}
	for _, p := range GetAll() {
		if p.ID != "anthropic" {
			continue
		}
		for _, m := range p.Models {
			exp, ok := want[m.ID]
			if !ok {
				continue
			}
			if got := m.ResolveImageLimits(p.Type).MaxLongEdge; got != exp {
				t.Errorf("%s MaxLongEdge = %d, want %d", m.ID, got, exp)
			}
		}
	}
}
