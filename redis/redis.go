package ormRedis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
)

// Node 单机节点配置
type Node struct {
	Addr     string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database int    `json:"database"`
}

// New 单机模式下创建客户端
//
//goland:noinspection ALL
func New(ctx context.Context, node *Node) (*redis.Client, error) {
	if node == nil {
		return nil, errors.New("not found redis node")
	}

	// 创建客户端
	client := redis.NewClient(&redis.Options{
		Addr:     node.Addr,
		Username: node.User,
		Password: node.Password,
		DB:       node.Database,
	})

	// 测试链接
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return client, nil
}

// ClusterNode 集群单节点地址
type ClusterNode struct {
	Addr     string `json:"addr"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// NewClusterByNode 根据集群内单个节点地址创建客户端
//
//goland:noinspection ALL
func NewClusterByNode(ctx context.Context, clusterNode *ClusterNode) (*redis.ClusterClient, error) {
	if clusterNode == nil {
		return nil, errors.New("not found redis node")
	}

	// 创建客户端
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{clusterNode.Addr},
		Username: clusterNode.User,
		Password: clusterNode.Password,
	})

	// 测试链接
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return client, nil
}

// ClusterNodes 集群多节点配置
type ClusterNodes struct {
	Addrs    []string `json:"addrs"`
	User     string   `json:"user"`
	Password string   `json:"password"`
}

// NewClusterByNodes 根据集群内多个节点地址创建客户端
//
//goland:noinspection ALL
func NewClusterByNodes(ctx context.Context, clusterNodes *ClusterNodes) (*redis.ClusterClient, error) {
	if clusterNodes == nil || len(clusterNodes.Addrs) == 0 {
		return nil, errors.New("not found redis node")
	}

	// 创建客户端
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    clusterNodes.Addrs,
		Username: clusterNodes.User,
		Password: clusterNodes.Password,
	})

	// 测试链接
	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return client, nil
}
