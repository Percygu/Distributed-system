package model

import (
	"sync"
)

// 服务实例结构
type Instance struct {
	Env      string   `json:"env"`      // 环境
	AppID    string   `json:"appid"`    // 业务id，应用服务的唯一标识
	Hostname string   `json:"hostname"` // 服务实例的唯一标识
	Addrs    []string `json:"addrs"`    // 服务实例的地址，可以是 http 或 rpc 地址，多个地址可以维护数组
	Version  string   `json:"version"`  // 服务实例版本
	Status   uint32   `json:"status"`   // 服务实例状态，用于控制上下线

	RegTimestamp    int64 `json:"reg_timestamp"`    // 注册时间戳
	UpTimestamp     int64 `json:"up_timestamp"`     // 上线时间戳
	RenewTimestamp  int64 `json:"renew_timestamp"`  // 最近续约时间戳
	DirtyTimestamp  int64 `json:"dirty_timestamp"`  // 脏时间戳
	LatestTimestamp int64 `json:"latest_timestamp"` //
}

type Application struct {
	AppID           string // 业务id
	Instances       map[string]*Instance
	latestTimestamp int64
	lock            sync.RWMutex
}

// 注册表数据结构
type Registry struct {
	Apps map[string]*Application
	lock sync.RWMutex
}
