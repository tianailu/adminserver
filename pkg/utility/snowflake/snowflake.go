package snowflake

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/tianailu/adminserver/pkg/db/redis"
)

var node *snowflake.Node

func GetNode() *snowflake.Node {
	return node
}

func CreateSnowflakeClient() error {
	ctx := context.Background()
	cmd := redis.GetRDB().Incr(ctx, redis.AdminServerNodeIdKey)

	err := cmd.Err()
	if err != nil {
		return err
	}

	nodeId := cmd.Val()
	snowflakeNode, err := snowflake.NewNode(nodeId)
	if err != nil {
		return err
	}

	node = snowflakeNode

	return nil
}
