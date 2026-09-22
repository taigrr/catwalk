package providers

import (
	"slices"
	"testing"

	"github.com/taigrr/catwalk/pkg/catwalk"
)

func TestValidDefaultModels(t *testing.T) {
	for _, p := range GetAll() {
		t.Run(p.Name, func(t *testing.T) {
			evaluationOnly := len(p.Models) == 0 && len(p.EvaluationModels) > 0
			var modelIds []string
			for _, m := range p.Models {
				modelIds = append(modelIds, m.ID)
			}
			if !evaluationOnly && !slices.Contains(modelIds, p.DefaultLargeModelID) {
				t.Errorf("Default large model %q not found in provider %q", p.DefaultLargeModelID, p.Name)
			}
			if !evaluationOnly && !slices.Contains(modelIds, p.DefaultSmallModelID) {
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
			if len(p.EvaluationModels) > 0 && p.DefaultEvaluationModelID == "" {
				t.Errorf("Provider %q has evaluation models but no default_evaluation_model_id", p.Name)
			}
			if p.DefaultEvaluationModelID != "" {
				var evaluationIDs []string
				for _, m := range p.EvaluationModels {
					evaluationIDs = append(evaluationIDs, m.ID)
					if m.CostPer1MOut != 0 {
						t.Errorf("Evaluation model %q in %q has non-zero output cost; decision models bill input only", m.ID, p.Name)
					}
				}
				if !slices.Contains(evaluationIDs, p.DefaultEvaluationModelID) {
					t.Errorf("Default evaluation model %q not found in provider %q", p.DefaultEvaluationModelID, p.Name)
				}
			}
		})
	}
}

func TestEvaluationProviders(t *testing.T) {
	var vercel, typesafe *catwalk.Provider
	for _, p := range GetAll() {
		switch p.ID {
		case catwalk.InferenceProviderVercel:
			vercel = &p
		case catwalk.InferenceProviderTypeSafe:
			typesafe = &p
		}
	}
	if vercel == nil || typesafe == nil {
		t.Fatal("expected both vercel and typesafe providers")
	}
	if vercel.DefaultEvaluationModelID != "typesafe-ai/jev" {
		t.Errorf("vercel default evaluation model = %q", vercel.DefaultEvaluationModelID)
	}
	if typesafe.Type != catwalk.TypeTypeSafe || len(typesafe.Models) != 0 {
		t.Errorf("typesafe provider should be evaluation-only, got %+v", typesafe)
	}
	if !slices.Contains(catwalk.KnownProviders(), catwalk.InferenceProviderTypeSafe) {
		t.Error("typesafe missing from KnownProviders")
	}
	if !slices.Contains(catwalk.KnownProviderTypes(), catwalk.TypeTypeSafe) {
		t.Error("typesafe missing from KnownProviderTypes")
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
