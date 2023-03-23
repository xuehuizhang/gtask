package task

import (
	"errors"
	"sync"
)

type Rule struct {
	Name   string `json:"name" yaml:"name"`     //Job名称
	Spec   string `json:"spec" yaml:"spec"`     //job执行周期
	Url    string `json:"url" yaml:"url"`       //http url
	Method string `json:"method" yaml:"method"` //http method
	Params string `json:"params" yaml:"params"` //http json params
	Cate   int    `json:"cate" yaml:"cate"`     //0 http 1 bash
}

type RulePool struct {
	m    map[string]*Rule `json:"m"`
	lock sync.RWMutex
}

func InitRulePool() *RulePool {
	return &RulePool{
		m:    make(map[string]*Rule),
		lock: sync.RWMutex{},
	}
}

func (p *RulePool) Del(name string) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	delete(p.m, name)
	return nil
}

func (p *RulePool) SaveOrUpdate(r *Rule) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.m[r.Name] = r
}

func (p *RulePool) Get(name string) (*Rule, error) {
	p.lock.RLock()
	defer p.lock.RUnlock()
	rule, ok := p.m[name]
	if ok {
		return rule, nil
	}
	return nil, errors.New("un found")
}
