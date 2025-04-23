package gen

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
)

var (
	node *snowflake.Node
)

func UUID() string {
	return uuid.NewString()
}

func SnowFlake() int64 {
	if node == nil {
		node, _ = snowflake.NewNode(1)
	}

	return node.Generate().Int64()
}
