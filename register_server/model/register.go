package model

import (
	"Distributed-system/register_server/params"
	"fmt"
	"log"
	"time"
)

// 初始化注册服务表
func NewRegistry() *Registry {
	registry := &Registry{
		Apps: make(map[string]*Application),
	}
	return registry
}

func NewInstance(req params.RequestRegister) *Instance {
	now := time.Now().UnixNano()
	instance := &Instance{
		Env:             req.Env,
		AppID:           req.AppID,
		Hostname:        req.HostName,
		Addrs:           req.Addrs,
		Version:         req.Version,
		Status:          req.Status,
		RegTimestamp:    now,
		UpTimestamp:     now,
		RenewTimestamp:  now,
		DirtyTimestamp:  now,
		LatestTimestamp: now,
	}
	return instance
}

// 创建application
func NewApplication(appid string) *Application {
	return &Application{
		AppID:     appid,
		Instances: make(map[string]*Instance), // 一个业务application有多个实例instance
	}
}

// 注册实例
func (r *Registry) Register(instance *Instance, latestTimestamp int64) (*Application, error) {
	key := fmt.Sprintf("%s_%s", instance.AppID, instance.Env) // 构造唯一key
	r.lock.Lock()
	// 判断该实例所对应的applocation是否存在
	app, ok := r.Apps[key]
	r.lock.Unlock()
	if !ok {
		app = NewApplication(instance.AppID)
	}
	_, isNew := app.AddInstance(instance, latestTimestamp)
	if isNew {

	}
	r.lock.Lock()
	r.Apps[key] = app
	r.lock.Unlock()
	return app, nil
}

// 在applicatio里新增实例
func (a *Application) AddInstance(instance *Instance, latestTimstamp int64) (*Instance, bool) {
	a.lock.Lock()
	defer a.lock.Unlock()
	ins, ok := a.Instances[instance.Hostname] // 业务对应的实例是否存在
	if ok {                                   // exists
		instance.UpTimestamp = ins.UpTimestamp
		if instance.DirtyTimestamp < ins.DirtyTimestamp {
			log.Println("register exist dirty timestamp")
			instance = ins
		}
	}
	a.Instances[instance.Hostname] = instance
	a.latestTimestamp = latestTimstamp
	returnIns := new(Instance) // 创建一个instance的拷贝返回
	*returnIns = *instance
	return returnIns, !ok   // 当ok为false，即不存在实例，是新增的时候返回true
}

//
