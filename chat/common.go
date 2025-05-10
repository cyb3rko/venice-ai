package chat

type ChatConfig struct {
	Model               string                   `json:"model"`
	Messages            []map[string]interface{} `json:"messages"`
	FrequencyPenalty    float32                  `json:"frequency_penalty,omitempty"`
	Logprobs            bool                     `json:"logprobs,omitempty"`
	MaxCompletionTokens int                      `json:"max_completion_tokens,omitempty"`
	MaxTemp             float32                  `json:"max_temp,omitempty"`
	MaxTokens           int                      `json:"max_tokens,omitempty"`
	MinP                float32                  `json:"min_p,omitempty"`
	MinTemp             float32                  `json:"min_temp,omitempty"`
	N                   int                      `json:"n,omitempty"`
	PresencePenalty     float32                  `json:"presence_penalty,omitempty"`
	RepetitionPenalty   float32                  `json:"repetition_penalty,omitempty"`
	Seed                int                      `json:"seed,omitempty"`
	Stop                []string                 `json:"stop,omitempty"`
	StopTokenIds        []int                    `json:"stop_token_ids,omitempty"`
	Stream              bool                     `json:"stream,omitempty"`
	StreamOptions       StreamOptions            `json:"stream_options,omitempty"`
	Temperature         float32                  `json:"temperature,omitempty"`
	TopK                int                      `json:"top_k,omitempty"`
	TopP                float32                  `json:"top_p,omitempty"`
	VeniceParameters    []any                    `json:"venice_parameters,omitempty"`
	ParallelToolCalls   bool                     `json:"parallel_tool_calls,omitempty"`
	ResponseFormat      interface{}              `json:"response_format,omitempty"`
	ToolChoice          interface{}              `json:"tool_choice,omitempty"`
	Tools               []interface{}            `json:"tools,omitempty"`
}

type UserMessage struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}

type StreamOptions struct {
	IncludeUsage bool `json:"include_usage,omitempty"`
}

type VeniceParameters struct {
	CharacterSlug             bool `json:"character_slug,omitempty"`
	EnableWebSearch           bool `json:"enable_web_search,omitempty"`
	IncludeVeniceSystemPrompt bool `json:"include_venice_system_prompt,omitempty"`
}
