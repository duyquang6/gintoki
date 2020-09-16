package dependency

import (
	"gintoki/domain/tool"

	tool2 "gintoki/infrastructure/tool"
)

func NewIdGenerator() tool.IdGenerator {
	return tool2.NewUuidGenerator()
}
