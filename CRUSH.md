# Catwalk - AI Provider Database

## Build/Test Commands

- `go run .` - Build and run the main HTTP server on :8080
- `go run ./cmd/{provider-name}` - Build and run a CLI to update the `{provider-name}.json` file
- `go test ./...` - Run all tests

## Code Style Guidelines

- Package comments: Start with "Package name provides/represents..."
- Imports: Standard library first, then third-party, then local packages
- Error handling: Use `fmt.Errorf("message: %w", err)` for wrapping
- Struct tags: Use json tags with omitempty for optional fields
- Constants: Group related constants with descriptive comments
- Types: Use custom types for IDs (e.g., `InferenceProvider`, `Type`)
- Naming: Use camelCase for unexported, PascalCase for exported
- Comments: Use `//nolint:directive` for linter exceptions
- HTTP: Always set timeouts, use context, defer close response bodies
- JSON: Use `json.MarshalIndent` for pretty output, validate unmarshaling
- File permissions: Use 0o600 for sensitive config files
- Always format code with `gofumpt`

## Adding more provider commands

- Create the `./cmd/{provider-name}/main.go` file
- Try to use the provider API to figure out the available models. If there's no
  endpoint for listing the models, look for some sort of structured text format
  (usually in the docs). If none of that exist, refuse to create the command,
  and add it to the `MANUAL_UPDATES.md` file.
- Add it to `.github/workflows/update.yml`

## Updating providers manually

### TypeSafe

`typesafe.json` is evaluation-only (Jev decision models, no chat models). TypeSafe's `GET /v1/models` requires an API key, so update the model list and pricing by hand from `https://docs.typesafe.ai/models`. Evaluation models bill input tokens only; `cost_per_1m_out` must stay `0`. The Vercel AI Gateway entry for Jev (`typesafe-ai/jev`) is populated automatically by `cmd/vercel` from models with `"type": "evaluation"`.

### Kev

`kev.json` lists the in-process Kev checkpoints served by `fantasy/providers/kev` (weights downloaded at runtime; no endpoint, no key, zero cost). Keep ids in sync with the `Checkpoint*` constants in that package.

### Zai

For `zai`, we'll need to grab the model list and capabilities from `https://docs.z.ai/guides/overview/overview`.

That page does not contain the exact `context_window` and `default_max_tokens` though. We can grab the exact value from `./internal/providers/configs/openrouter.json`.
