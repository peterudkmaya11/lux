package lux

import (
	"sync"
)

//AgentManager is a struct to manage and synchronize all the Agents
type AgentManager struct {
	numAgents int
	ticker    chan chan struct{}
	wg        sync.WaitGroup
	lock      sync.Mutex
}

//NewAgentManager create an AgentManager and initialize all required values
func NewAgentManager() *AgentManager {
	out := AgentManager{
		numAgents: 0,
		ticker:    make(chan chan struct{}),
		wg:        sync.WaitGroup{},
		lock:      sync.Mutex{},
	}
	return &out
}

//Tick will notify every agent that they need to execute their callback.
func (am *AgentManager) Tick() {
	tick := make(chan struct{})
	am.lock.Lock()
	am.wg.Add(am.numAgents)
	for x := 0; x < am.numAgents; x++ {
		am.ticker <- tick
	}
	am.lock.Unlock()
	close(tick)
	am.wg.Wait()
}

//NewAgent starts a goroutine that will run callback every frame until it returns false, it will then die.
func (am *AgentManager) NewAgent(callback func() bool) {
	am.lock.Lock()
	am.numAgents++
	am.lock.Unlock()
	go func() {
		for {
			tick := <-am.ticker
			<-tick
			docontinue := callback()
			am.wg.Done()
			if docontinue {
				continue
			}
			return
		}
	}()
}
