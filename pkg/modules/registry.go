package modules

import "sync"

type ModulesRegistry struct {
	modules map[string]Module

	modulesMu sync.Mutex
}

func NewModulesRegistry() *ModulesRegistry {
	return &ModulesRegistry{}
}

func (mr *ModulesRegistry) AddModule(m Module) {
	mr.modulesMu.Lock()
	defer mr.modulesMu.Unlock()

	mr.modules[m.Name()] = m
}
