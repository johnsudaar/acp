package models

import "github.com/Scalingo/go-utils/mongo/document"

const (
	SceneCollection = "scene"
)

type Scene struct {
	document.Base `bson:",inline"`
	Name          string      `json:"name" bson:"name"`
	Elements      interface{} `json:"elements" bson:"elements"`
}
