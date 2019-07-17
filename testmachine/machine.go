/*
 * Copyright 2019, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package testmachine

import (
	"bytes"
	"github.com/offchainlabs/arb-avm/goloader"
	"log"

	"github.com/offchainlabs/arb-util/machine"
	"github.com/offchainlabs/arb-util/protocol"
	"github.com/offchainlabs/arb-util/value"

	gomachine "github.com/offchainlabs/arb-avm/vm"

	"github.com/offchainlabs/arb-avm-cpp/cmachine"
)

type Machine struct {
	cppmachine *cmachine.Machine
	gomachine  *gomachine.Machine
}

func New(codeFile string, warnMode bool) (*Machine, error) {
	gm, err := goloader.LoadMachineFromFile(codeFile, warnMode)
	return &Machine{
		cmachine.New(codeFile),
		gm,
	}, err
}

func (m *Machine) Hash() [32]byte {
	h1 := m.cppmachine.Hash()
	h2 := m.gomachine.Hash()
	if h1 != h2 {
		log.Fatalln("Hash error at pc", m.gomachine.GetPC())
	}
	return h1
}

func (m *Machine) Clone() machine.Machine {
	return &Machine{m.cppmachine.Clone().(*cmachine.Machine), m.gomachine.Clone().(*gomachine.Machine)}
}

func (m *Machine) InboxHash() value.HashOnlyValue {
	h1 := m.cppmachine.InboxHash()
	h2 := m.gomachine.InboxHash()
	if h1.Equal(h2) {
		log.Fatalln("InboxHash error at pc", m.gomachine.GetPC())
	}
	return h1
}

func (m *Machine) HasPendingMessages() bool {
	h1 := m.cppmachine.HasPendingMessages()
	h2 := m.gomachine.HasPendingMessages()
	if h1 != h2 {
		log.Fatalln("HasPendingMessages error at pc", m.gomachine.GetPC())
	}
	return h1
}

func (m *Machine) SendOnchainMessage(msg protocol.Message) {
	m.cppmachine.SendOnchainMessage(msg)
	m.gomachine.SendOnchainMessage(msg)
}

func (m *Machine) DeliverOnchainMessage() {
	m.cppmachine.DeliverOnchainMessage()
	m.gomachine.DeliverOnchainMessage()
}

func (m *Machine) SendOffchainMessages(msgs []protocol.Message) {
	m.cppmachine.SendOffchainMessages(msgs)
	m.gomachine.SendOffchainMessages(msgs)
}

func (m *Machine) ExecuteAssertion(maxSteps int32, timeBounds protocol.TimeBounds) *protocol.Assertion {
	a := &protocol.Assertion{}
	for i := int32(0); i < maxSteps; i++ {
		a1 := m.cppmachine.ExecuteAssertion(1, timeBounds)
		a2 := m.gomachine.ExecuteAssertion(1, timeBounds)

		if !a1.Equals(a2) {
			log.Fatalln("ExecuteAssertion error at pc", m.gomachine.GetPC())
		}
		a.AfterHash = a1.AfterHash
		a.NumSteps++
		a.Logs = append(a.Logs, a1.Logs...)
		a.OutMsgs = append(a.OutMsgs, a1.OutMsgs...)
		if a1.NumSteps == 0 {
			break
		}
	}
	return a
}

func (m *Machine) MarshalForProof() ([]byte, error) {
	h1, err1 := m.cppmachine.MarshalForProof()
	h2, err2 := m.gomachine.MarshalForProof()

	if err1 != nil {
		return nil, err1
	}

	if err2 != nil {
		return nil, err2
	}

	if !bytes.Equal(h1, h2) {
		log.Fatalln("MarshalForProof error at pc", m.gomachine.GetPC())
	}
	return h1, nil
}
