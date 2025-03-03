package images

type ImageConfig struct {
	Model          string      `json:"model"`
	Prompt         string      `json:"prompt"`
	NegativePrompt string      `json:"negative_prompt,omitempty"`
	StylePreset    string      `json:"style_preset,omitempty"`
	Height         int         `json:"height,omitempty"`
	Width          int         `json:"width,omitempty"`
	Steps          int         `json:"steps,omitempty"`
	CfgScale       float32     `json:"cfg_scale,omitempty"`
	Seed           int         `json:"seed,omitempty"`
	LoraStrength   int         `json:"lora_strength,omitempty"`
	SafeMode       bool        `json:"safe_mode,omitempty"`
	HideWatermark  bool        `json:"hide_watermark,omitempty"`
	Inpaint        interface{} `json:"inpaint,omitempty"`
}
