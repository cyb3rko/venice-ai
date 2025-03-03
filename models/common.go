package models

type ModelType string

const (
	ModelTypeDefault ModelType = ModelTypeText
	ModelTypeImage   ModelType = "image"
	ModelTypeText    ModelType = "text"
)
