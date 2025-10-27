package main

import (
	"fmt"
	"sync"
)

type Proposal struct {
	Number int
	Value  string
}

type Acceptor struct {
	mu         sync.Mutex
	promisedID int
	accepted   *Proposal
}

func (a *Acceptor) Prepare(n int) (bool, *Proposal) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if n > a.promisedID {
		a.promisedID = n
		return true, a.accepted
	}
	return false, nil
}

func (a *Acceptor) Accept(p Proposal) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if p.Number >= a.promisedID {
		a.promisedID = p.Number
		a.accepted = &p
		return true
	}
	return false
}

func Propose(value string, acceptors []*Acceptor) (string, bool) {
	proposalNum := 1
	var acceptedValue string

	// Prepare phase
	promises := 0
	var highest *Proposal
	for _, a := range acceptors {
		ok, prev := a.Prepare(proposalNum)
		if ok {
			promises++
			if prev != nil && (highest == nil || prev.Number > highest.Number) {
				highest = prev
			}
		}
	}
	if promises <= len(acceptors)/2 {
		return "", false
	}

	if highest != nil {
		value = highest.Value
	}

	// Accept phase
	accepts := 0
	for _, a := range acceptors {
		if a.Accept(Proposal{Number: proposalNum, Value: value}) {
			accepts++
		}
	}
	if accepts > len(acceptors)/2 {
		acceptedValue = value
		return acceptedValue, true
	}
	return "", false
}

func main() {
	// 3 acceptors
	acceptors := []*Acceptor{{}, {}, {}}

	value, ok := Propose("Hello Paxos", acceptors)
	if ok {
		fmt.Println("Consensus reached on:", value)
	} else {
		fmt.Println("Consensus failed")
	}
}
