package flake

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func InitFlake() error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	root := r.Int63n(1024)
	n, err := snowflake.NewNode(root)
	if err != nil {
		return err
	}
	node = n
	return nil
}
func GetInt64() int64 {
	return node.Generate().Int64()
}
