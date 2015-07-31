package lux

import (
	"sync"
)

// Agent is the handle to control a tickable agent.
type Agent struct {
	death, wake, sleep chan struct{}
	manager            *AgentManager
	sleeping           bool
}

// Seppuku kills this agent.
func (a *Agent) Seppuku() {
	a.manager.lock.Lock()
	a.death <- struct{}{}
	a.manager.numAgents--
	a.manager.lock.Unlock()
}

// IsSleeping return wether this agent is sleeping or not
func (a *Agent) IsSleeping() bool {
	return a.sleeping
}

// Sleep puts this agent to sleep, meaning he won't react to Ticks
func (a *Agent) Sleep() {
	a.sleeping = true
	a.sleep <- struct{}{}
}

// Awake wakes this agent from it's slumber. He will receive the next tick
func (a *Agent) Awake() {
	a.wake <- struct{}{}
	a.sleeping = false
}

// AgentManager is a struct to manage and synchronize all the Agents.
type AgentManager struct {
	numAgents int
	ticker    chan chan struct{}
	wg        sync.WaitGroup
	lock      sync.Mutex
}

// NewAgentManager create an AgentManager and initialize all required values.
func NewAgentManager() *AgentManager {
	return &AgentManager{
		ticker: make(chan chan struct{}),
	}
}

// Tick will notify every agent that they need to execute their callback.
func (am *AgentManager) Tick() {
	if am.numAgents == 0 {
		return
	}
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

// NewAgent starts a goroutine that will run callback every frame until it returns false, it will then die.
func (am *AgentManager) NewAgent(callback func() bool) *Agent {
	am.lock.Lock()

	agent := Agent{
		death:   make(chan struct{}),
		wake:    make(chan struct{}),
		sleep:   make(chan struct{}),
		manager: am,
	}

	go func() {
		for {
			select {
			case tick := <-am.ticker:
				<-tick // Wait for signal to start working
				docontinue := callback()
				if !docontinue {
					am.lock.Lock()
					am.numAgents--
					am.lock.Unlock()
					am.wg.Done()
					return
				}
			case <-agent.sleep:
				am.lock.Lock()
				am.numAgents--
				am.lock.Unlock()
				<-agent.wake
				am.lock.Lock()
				am.numAgents++
				am.lock.Unlock()
			case <-agent.death:
				return
			}
			am.wg.Done()
		}
	}()

	am.numAgents++

	am.lock.Unlock()

	return &agent
}

// AgentCount return the number of active agent.
func (am *AgentManager) AgentCount() int {
	return am.numAgents
}
