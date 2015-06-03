package lux

import (
	"sync"
)

type Agent struct {
	death   chan struct{}
	manager *AgentManager
}

func (agent Agent) Seppuku() {
	agent.manager.lock.Lock()
	close(agent.death)
	agent.manager.numAgents--
	agent.manager.lock.Unlock()
}

//AgentManager is a struct to manage and synchronize all the Agents.
type AgentManager struct {
	numAgents int
	ticker    chan chan struct{}
	wg        sync.WaitGroup
	lock      sync.Mutex
}

//NewAgentManager create an AgentManager and initialize all required values.
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
	D("4 ", am.numAgents)
	for x := 0; x < am.numAgents; x++ {
		am.ticker <- tick
	}
	D("5")
	am.lock.Unlock()
	close(tick)
	am.wg.Wait()
}

//NewAgent starts a goroutine that will run callback every frame until it returns false, it will then die.
func (am *AgentManager) NewAgent(callback func() bool) Agent {
	am.lock.Lock()
	am.numAgents++
	am.lock.Unlock()
	agent := Agent{
		death:   make(chan struct{}),
		manager: am,
	}
	go func() {
		for {
			select {
			case tick := <-am.ticker:
				<-tick //wait for signal to start working
				docontinue := callback()
				am.wg.Done()
				if docontinue {
					continue
				}
				return
			case <-agent.death:
				return
			}
		}
	}()
	return agent
}
