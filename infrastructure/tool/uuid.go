package tool

import (
	"gintoki/domain/tool"

	"github.com/google/uuid"
)

type uuidGenerator struct {
}

func (g *uuidGenerator) Generate() string {
	return uuid.New().String()
}

func NewUuidGenerator() tool.IdGenerator {
	return &uuidGenerator{}
}
