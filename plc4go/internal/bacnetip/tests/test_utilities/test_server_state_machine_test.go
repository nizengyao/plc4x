/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package test_utilities

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/apache/plc4x/plc4go/internal/bacnetip"
	"github.com/apache/plc4x/plc4go/internal/bacnetip/tests"
	"github.com/apache/plc4x/plc4go/spi/testutils"
)

func TestServerStateMachine(t *testing.T) {
	testingLogger := testutils.ProduceTestingLogger(t)
	// create a client state machine, trapped server, and bind them together
	client, err := tests.NewTrappedClient(testingLogger)
	require.NoError(t, err)
	server, err := tests.NewServerStateMachine(testingLogger)
	require.NoError(t, err)
	err = bacnetip.Bind(testingLogger, client, server)
	require.NoError(t, err)

	// make pdu object
	pdu := bacnetip.NewPDU(tests.NewDummyMessage())

	// make a send transition from start to success, run the machine
	server.GetStartState().Send(pdu, nil).Success("")

	// run the machine
	err = server.Run()
	assert.Nil(t, err)

	// check for success
	assert.False(t, server.IsRunning())
	assert.True(t, server.GetCurrentState().IsSuccessState())

	// make sure the pdu was sent
	assert.Equal(t, pdu, client.GetConfirmationReceived())

	// check the transaction log
	assert.Len(t, server.GetTransactionLog(), 1)
	assert.Contains(t, server.GetTransactionLog()[0], pdu.String())
}
