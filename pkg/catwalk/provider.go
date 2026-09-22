package catwalk

// Type represents the type of AI provider.
type Type string

// All the supported AI provider types.
const (
	TypeOpenAI       Type = "openai"
	TypeOpenAICompat Type = "openai-compat"
	TypeOpenRouter   Type = "openrouter"
	TypeVercel       Type = "vercel"
	TypeAnthropic    Type = "anthropic"
	TypeGoogle       Type = "google"
	TypeAzure        Type = "azure"
	TypeBedrock      Type = "bedrock"
	TypeVertexAI     Type = "google-vertex"
	TypeTypeSafe     Type = "typesafe"
	// TypeKev runs Kev decision models in-process; no endpoint or API key.
	TypeKev Type = "kev"
)

// InferenceProvider represents the inference provider identifier.
type InferenceProvider string

// All the inference providers supported by the system.
const (
	InferenceProviderOpenAI           InferenceProvider = "openai"
	InferenceProviderAnthropic        InferenceProvider = "anthropic"
	InferenceProviderSynthetic        InferenceProvider = "synthetic"
	InferenceProviderGemini           InferenceProvider = "gemini"
	InferenceProviderAzure            InferenceProvider = "azure"
	InferenceProviderBedrock          InferenceProvider = "bedrock"
	InferenceProviderBedrockEurope    InferenceProvider = "bedrock-europe"
	InferenceProviderBedrockMantle    InferenceProvider = "bedrock-mantle"
	InferenceProviderVertexAI         InferenceProvider = "vertexai"
	InferenceProviderXAI              InferenceProvider = "xai"
	InferenceProviderGrok             InferenceProvider = "grok"
	InferenceProviderZAI              InferenceProvider = "zai"
	InferenceProviderDeepSeek         InferenceProvider = "deepseek"
	InferenceProviderZhipu            InferenceProvider = "zhipu"
	InferenceProviderZhipuCoding      InferenceProvider = "zhipu-coding"
	InferenceProviderGROQ             InferenceProvider = "groq"
	InferenceProviderOpenRouter       InferenceProvider = "openrouter"
	InferenceProviderCerebras         InferenceProvider = "cerebras"
	InferenceProviderVenice           InferenceProvider = "venice"
	InferenceProviderChutes           InferenceProvider = "chutes"
	InferenceProviderHuggingFace      InferenceProvider = "huggingface"
	InferenceAIHubMix                 InferenceProvider = "aihubmix"
	InferenceKimiCoding               InferenceProvider = "kimi-coding"
	InferenceProviderCopilot          InferenceProvider = "copilot"
	InferenceProviderCortecs          InferenceProvider = "cortecs"
	InferenceProviderVercel           InferenceProvider = "vercel"
	InferenceProviderMiniMax          InferenceProvider = "minimax"
	InferenceProviderMiniMaxChina     InferenceProvider = "minimax-china"
	InferenceProviderIoNet            InferenceProvider = "ionet"
	InferenceProviderQiniuCloud       InferenceProvider = "qiniucloud"
	InferenceProviderAvian            InferenceProvider = "avian"
	InferenceProviderNebius           InferenceProvider = "nebius"
	InferenceProviderNeuralwatt       InferenceProvider = "neuralwatt"
	InferenceProviderOpenCodeZen      InferenceProvider = "opencode-zen"
	InferenceProviderOpenCodeGo       InferenceProvider = "opencode-go"
	InferenceProviderAlibabaSingapore InferenceProvider = "alibaba-singapore"
	InferenceProviderTypeSafe         InferenceProvider = "typesafe"
	InferenceProviderKev              InferenceProvider = "kev"
)

// Provider represents an AI provider configuration.
type Provider struct {
	Name                    string            `json:"name"`
	ID                      InferenceProvider `json:"id"`
	APIKey                  string            `json:"api_key,omitempty"`
	APIEndpoint             string            `json:"api_endpoint,omitempty"`
	Type                    Type              `json:"type,omitempty"`
	DefaultLargeModelID     string            `json:"default_large_model_id,omitempty"`
	DefaultSmallModelID     string            `json:"default_small_model_id,omitempty"`
	DefaultEmbeddingModelID string            `json:"default_embedding_model_id,omitempty"`
	// DefaultEvaluationModelID names the preferred entry in EvaluationModels.
	DefaultEvaluationModelID string  `json:"default_evaluation_model_id,omitempty"`
	Models                   []Model `json:"models,omitempty"`
	// EmbeddingModels lists the provider's text-embedding models. It is
	// kept separate from Models so consumers iterating chat models are
	// unaffected by embedding entries.
	EmbeddingModels []Model `json:"embedding_models,omitempty"`
	// EvaluationModels lists the provider's decision/evaluation models (e.g.
	// TypeSafe Jev), which answer typed questions with probabilities instead
	// of generating text. Kept separate from Models for the same reason as
	// EmbeddingModels.
	EvaluationModels []Model           `json:"evaluation_models,omitempty"`
	DefaultHeaders   map[string]string `json:"default_headers,omitempty"`
}

// ModelOptions stores extra options for models.
type ModelOptions struct {
	Temperature      *float64       `json:"temperature,omitempty"`
	TopP             *float64       `json:"top_p,omitempty"`
	TopK             *int64         `json:"top_k,omitempty"`
	FrequencyPenalty *float64       `json:"frequency_penalty,omitempty"`
	PresencePenalty  *float64       `json:"presence_penalty,omitempty"`
	ProviderOptions  map[string]any `json:"provider_options,omitempty"`
}

// Model represents an AI model configuration.
type Model struct {
	ID                     string       `json:"id"`
	Name                   string       `json:"name"`
	CostPer1MIn            float64      `json:"cost_per_1m_in"`
	CostPer1MOut           float64      `json:"cost_per_1m_out"`
	CostPer1MInCached      float64      `json:"cost_per_1m_in_cached"`
	CostPer1MOutCached     float64      `json:"cost_per_1m_out_cached"`
	ContextWindow          int64        `json:"context_window"`
	DefaultMaxTokens       int64        `json:"default_max_tokens"`
	CanReason              bool         `json:"can_reason"`
	ReasoningLevels        []string     `json:"reasoning_levels,omitempty"`
	DefaultReasoningEffort string       `json:"default_reasoning_effort,omitempty"`
	SupportsImages         bool         `json:"supports_attachments"`
	Options                ModelOptions `json:"options,omitzero"`
	// Image holds per-model image-input constraints. When zero, callers
	// should resolve the provider-family defaults via
	// [Model.ResolveImageLimits]; non-zero fields override those
	// defaults (e.g. newer Claude models that allow a larger long edge).
	Image ImageLimits `json:"image,omitzero"`
	// Supports1MContext indicates the model supports the 1M context window beta feature.
	Supports1MContext bool `json:"supports_1m_context,omitempty"`
	// Long context pricing (when >200K input tokens with 1M context enabled).
	// These are the premium rates; standard rates are in the CostPer1M* fields above.
	LongContextCostPer1MIn       float64 `json:"long_context_cost_per_1m_in,omitempty"`
	LongContextCostPer1MOut      float64 `json:"long_context_cost_per_1m_out,omitempty"`
	LongContextCostPer1MInCached float64 `json:"long_context_cost_per_1m_in_cached,omitempty"`
	// Dimensions is the size of the output vector for embedding models.
	// For embedding models that support multiple output sizes this is the
	// default/recommended dimensionality. It is zero for chat models.
	Dimensions int64 `json:"dimensions,omitempty"`
}

// IsEmbedding reports whether the model is a text-embedding model,
// identified by a non-zero output Dimensions.
func (m Model) IsEmbedding() bool {
	return m.Dimensions > 0
}

// ImageLimits describes a model's documented image-input constraints.
// All values are the provider's published maximums; consumers are
// expected to apply their own safety margin before enforcing them.
//
// Two classes of limit are modeled:
//
//   - Per-image: every single image must independently satisfy
//     MaxLongEdge (longest edge, pixels) and MaxBytesPerImage (encoded
//     size).
//   - Aggregate: summed across every image in a request (the full
//     replayed thread plus the new message) the totals must satisfy
//     MaxAggregatePixels (sum of width*height) and MaxAggregateBytes
//     (total encoded bytes); MaxImages caps the count.
//
// A zero field means "unknown" and should be resolved against the
// provider-family defaults from [DefaultImageLimits].
type ImageLimits struct {
	MaxLongEdge        int64 `json:"max_long_edge,omitempty"`
	MaxBytesPerImage   int64 `json:"max_bytes_per_image,omitempty"`
	MaxAggregatePixels int64 `json:"max_aggregate_pixels,omitempty"`
	MaxAggregateBytes  int64 `json:"max_aggregate_bytes,omitempty"`
	MaxImages          int64 `json:"max_images,omitempty"`
}

// IsZero reports whether no image limit has been set.
func (l ImageLimits) IsZero() bool {
	return l == ImageLimits{}
}

// merge returns l with any zero field filled from def.
func (l ImageLimits) merge(def ImageLimits) ImageLimits {
	if l.MaxLongEdge == 0 {
		l.MaxLongEdge = def.MaxLongEdge
	}
	if l.MaxBytesPerImage == 0 {
		l.MaxBytesPerImage = def.MaxBytesPerImage
	}
	if l.MaxAggregatePixels == 0 {
		l.MaxAggregatePixels = def.MaxAggregatePixels
	}
	if l.MaxAggregateBytes == 0 {
		l.MaxAggregateBytes = def.MaxAggregateBytes
	}
	if l.MaxImages == 0 {
		l.MaxImages = def.MaxImages
	}
	return l
}

// conservativeImageLimits is the fallback for provider types without a
// documented entry. It is intentionally strict so an unknown model is
// never sent images that blow an undocumented ceiling.
var conservativeImageLimits = ImageLimits{
	MaxLongEdge:        1568,
	MaxBytesPerImage:   4_000_000,
	MaxAggregatePixels: 30_000_000,
	MaxAggregateBytes:  20_000_000,
	MaxImages:          50,
}

// defaultImageLimitsByType maps a provider API family to its documented
// image limits. Providers that proxy a known family (Bedrock/Vertex
// hosting Claude) use the stricter platform values (e.g. 5MB/image).
var defaultImageLimitsByType = map[Type]ImageLimits{
	TypeAnthropic: {
		MaxLongEdge:        1568,
		MaxBytesPerImage:   5_000_000,
		MaxAggregatePixels: 40_000_000,
		MaxAggregateBytes:  32_000_000,
		MaxImages:          100,
	},
	TypeBedrock: {
		MaxLongEdge:        1568,
		MaxBytesPerImage:   5_000_000,
		MaxAggregatePixels: 40_000_000,
		MaxAggregateBytes:  32_000_000,
		MaxImages:          100,
	},
	TypeVertexAI: {
		MaxLongEdge:        1568,
		MaxBytesPerImage:   5_000_000,
		MaxAggregatePixels: 40_000_000,
		MaxAggregateBytes:  32_000_000,
		MaxImages:          100,
	},
	TypeOpenAI: {
		MaxLongEdge:        2048,
		MaxBytesPerImage:   20_000_000,
		MaxAggregatePixels: 50_000_000,
		MaxAggregateBytes:  50_000_000,
		MaxImages:          100,
	},
	TypeGoogle: {
		MaxLongEdge:        3072,
		MaxBytesPerImage:   7_000_000,
		MaxAggregatePixels: 60_000_000,
		MaxAggregateBytes:  20_000_000,
		MaxImages:          3000,
	},
}

// DefaultImageLimits returns the documented image limits for a provider
// API family, or a conservative fallback when the family is unknown.
func DefaultImageLimits(t Type) ImageLimits {
	if l, ok := defaultImageLimitsByType[t]; ok {
		return l
	}
	return conservativeImageLimits
}

// ResolveImageLimits returns the model's effective image limits: any
// per-model override fields set on Image take precedence, with the rest
// filled from the provider-family defaults for t. This is the single
// entry point consumers should use to obtain a model's image limits.
func (m Model) ResolveImageLimits(t Type) ImageLimits {
	return m.Image.merge(DefaultImageLimits(t))
}

// KnownProviders returns all the known inference providers.
func KnownProviders() []InferenceProvider {
	return []InferenceProvider{
		InferenceProviderOpenAI,
		InferenceProviderSynthetic,
		InferenceProviderAnthropic,
		InferenceProviderGemini,
		InferenceProviderAzure,
		InferenceProviderBedrock,
		InferenceProviderBedrockEurope,
		InferenceProviderBedrockMantle,
		InferenceProviderVertexAI,
		InferenceProviderXAI,
		InferenceProviderZAI,
		InferenceProviderZhipu,
		InferenceProviderZhipuCoding,
		InferenceProviderGROQ,
		InferenceProviderOpenRouter,
		InferenceProviderCerebras,
		InferenceProviderVenice,
		InferenceProviderChutes,
		InferenceProviderHuggingFace,
		InferenceAIHubMix,
		InferenceKimiCoding,
		InferenceProviderCopilot,
		InferenceProviderCortecs,
		InferenceProviderVercel,
		InferenceProviderMiniMax,
		InferenceProviderMiniMaxChina,
		InferenceProviderQiniuCloud,
		InferenceProviderAvian,
		InferenceProviderNebius,
		InferenceProviderNeuralwatt,
		InferenceProviderOpenCodeZen,
		InferenceProviderOpenCodeGo,
		InferenceProviderTypeSafe,
		InferenceProviderKev,
	}
}

// KnownProviderTypes returns all the known inference providers types.
func KnownProviderTypes() []Type {
	return []Type{
		TypeOpenAI,
		TypeOpenAICompat,
		TypeOpenRouter,
		TypeVercel,
		TypeAnthropic,
		TypeGoogle,
		TypeAzure,
		TypeBedrock,
		TypeVertexAI,
		TypeTypeSafe,
		TypeKev,
	}
}
