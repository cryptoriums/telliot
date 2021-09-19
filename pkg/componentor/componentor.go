// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package componentor

type Componentor interface {
	Status() bool
	ToggleStatus()
	ID() string
}

type MultiComponentor struct {
	componentors []Componentor
}

func (self *MultiComponentor) Add(componentor Componentor) {
	self.componentors = append(self.componentors, componentor)
}

func (self *MultiComponentor) ToggleStatus() {
	for _, c := range self.componentors {
		c.ToggleStatus()
	}
}

func (self *MultiComponentor) Status() bool {
	return self.componentors[0].Status()
}

func (self *MultiComponentor) ID() string {
	id := ""
	for _, c := range self.componentors {
		id += c.ID() + " "
	}
	return id
}
