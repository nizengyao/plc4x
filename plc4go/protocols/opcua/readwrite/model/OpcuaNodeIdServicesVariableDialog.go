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

package model

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaNodeIdServicesVariableDialog is an enum
type OpcuaNodeIdServicesVariableDialog int32

type IOpcuaNodeIdServicesVariableDialog interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_Prompt                               OpcuaNodeIdServicesVariableDialog = 2831
	OpcuaNodeIdServicesVariableDialog_DialogResponseMethodType_InputArguments                  OpcuaNodeIdServicesVariableDialog = 9032
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState                         OpcuaNodeIdServicesVariableDialog = 9035
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_Id                      OpcuaNodeIdServicesVariableDialog = 9036
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_Name                    OpcuaNodeIdServicesVariableDialog = 9037
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_Number                  OpcuaNodeIdServicesVariableDialog = 9038
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_EffectiveDisplayName    OpcuaNodeIdServicesVariableDialog = 9039
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_TransitionTime          OpcuaNodeIdServicesVariableDialog = 9040
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_EffectiveTransitionTime OpcuaNodeIdServicesVariableDialog = 9041
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_TrueState               OpcuaNodeIdServicesVariableDialog = 9042
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_FalseState              OpcuaNodeIdServicesVariableDialog = 9043
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState                          OpcuaNodeIdServicesVariableDialog = 9055
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_Id                       OpcuaNodeIdServicesVariableDialog = 9056
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_Name                     OpcuaNodeIdServicesVariableDialog = 9057
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_Number                   OpcuaNodeIdServicesVariableDialog = 9058
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_EffectiveDisplayName     OpcuaNodeIdServicesVariableDialog = 9059
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_TransitionTime           OpcuaNodeIdServicesVariableDialog = 9060
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_EffectiveTransitionTime  OpcuaNodeIdServicesVariableDialog = 9061
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_TrueState                OpcuaNodeIdServicesVariableDialog = 9062
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_FalseState               OpcuaNodeIdServicesVariableDialog = 9063
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_ResponseOptionSet                    OpcuaNodeIdServicesVariableDialog = 9064
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_DefaultResponse                      OpcuaNodeIdServicesVariableDialog = 9065
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_OkResponse                           OpcuaNodeIdServicesVariableDialog = 9066
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_CancelResponse                       OpcuaNodeIdServicesVariableDialog = 9067
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_LastResponse                         OpcuaNodeIdServicesVariableDialog = 9068
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_Respond_InputArguments               OpcuaNodeIdServicesVariableDialog = 9070
	OpcuaNodeIdServicesVariableDialog_DialogConditionType_Respond2_InputArguments              OpcuaNodeIdServicesVariableDialog = 24313
	OpcuaNodeIdServicesVariableDialog_DialogResponse2MethodType_InputArguments                 OpcuaNodeIdServicesVariableDialog = 24315
)

var OpcuaNodeIdServicesVariableDialogValues []OpcuaNodeIdServicesVariableDialog

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableDialogValues = []OpcuaNodeIdServicesVariableDialog{
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_Prompt,
		OpcuaNodeIdServicesVariableDialog_DialogResponseMethodType_InputArguments,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_Id,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_Name,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_Number,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_EffectiveDisplayName,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_TransitionTime,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_EffectiveTransitionTime,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_TrueState,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_FalseState,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_Id,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_Name,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_Number,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_EffectiveDisplayName,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_TransitionTime,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_EffectiveTransitionTime,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_TrueState,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_FalseState,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_ResponseOptionSet,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_DefaultResponse,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_OkResponse,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_CancelResponse,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_LastResponse,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_Respond_InputArguments,
		OpcuaNodeIdServicesVariableDialog_DialogConditionType_Respond2_InputArguments,
		OpcuaNodeIdServicesVariableDialog_DialogResponse2MethodType_InputArguments,
	}
}

func OpcuaNodeIdServicesVariableDialogByValue(value int32) (enum OpcuaNodeIdServicesVariableDialog, ok bool) {
	switch value {
	case 24313:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_Respond2_InputArguments, true
	case 24315:
		return OpcuaNodeIdServicesVariableDialog_DialogResponse2MethodType_InputArguments, true
	case 2831:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_Prompt, true
	case 9032:
		return OpcuaNodeIdServicesVariableDialog_DialogResponseMethodType_InputArguments, true
	case 9035:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState, true
	case 9036:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_Id, true
	case 9037:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_Name, true
	case 9038:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_Number, true
	case 9039:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_EffectiveDisplayName, true
	case 9040:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_TransitionTime, true
	case 9041:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_EffectiveTransitionTime, true
	case 9042:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_TrueState, true
	case 9043:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_FalseState, true
	case 9055:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState, true
	case 9056:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_Id, true
	case 9057:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_Name, true
	case 9058:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_Number, true
	case 9059:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_EffectiveDisplayName, true
	case 9060:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_TransitionTime, true
	case 9061:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_EffectiveTransitionTime, true
	case 9062:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_TrueState, true
	case 9063:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_FalseState, true
	case 9064:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_ResponseOptionSet, true
	case 9065:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DefaultResponse, true
	case 9066:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_OkResponse, true
	case 9067:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_CancelResponse, true
	case 9068:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_LastResponse, true
	case 9070:
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_Respond_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableDialogByName(value string) (enum OpcuaNodeIdServicesVariableDialog, ok bool) {
	switch value {
	case "DialogConditionType_Respond2_InputArguments":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_Respond2_InputArguments, true
	case "DialogResponse2MethodType_InputArguments":
		return OpcuaNodeIdServicesVariableDialog_DialogResponse2MethodType_InputArguments, true
	case "DialogConditionType_Prompt":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_Prompt, true
	case "DialogResponseMethodType_InputArguments":
		return OpcuaNodeIdServicesVariableDialog_DialogResponseMethodType_InputArguments, true
	case "DialogConditionType_EnabledState":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState, true
	case "DialogConditionType_EnabledState_Id":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_Id, true
	case "DialogConditionType_EnabledState_Name":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_Name, true
	case "DialogConditionType_EnabledState_Number":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_Number, true
	case "DialogConditionType_EnabledState_EffectiveDisplayName":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_EffectiveDisplayName, true
	case "DialogConditionType_EnabledState_TransitionTime":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_TransitionTime, true
	case "DialogConditionType_EnabledState_EffectiveTransitionTime":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_EffectiveTransitionTime, true
	case "DialogConditionType_EnabledState_TrueState":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_TrueState, true
	case "DialogConditionType_EnabledState_FalseState":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_FalseState, true
	case "DialogConditionType_DialogState":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState, true
	case "DialogConditionType_DialogState_Id":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_Id, true
	case "DialogConditionType_DialogState_Name":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_Name, true
	case "DialogConditionType_DialogState_Number":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_Number, true
	case "DialogConditionType_DialogState_EffectiveDisplayName":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_EffectiveDisplayName, true
	case "DialogConditionType_DialogState_TransitionTime":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_TransitionTime, true
	case "DialogConditionType_DialogState_EffectiveTransitionTime":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_EffectiveTransitionTime, true
	case "DialogConditionType_DialogState_TrueState":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_TrueState, true
	case "DialogConditionType_DialogState_FalseState":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_FalseState, true
	case "DialogConditionType_ResponseOptionSet":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_ResponseOptionSet, true
	case "DialogConditionType_DefaultResponse":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_DefaultResponse, true
	case "DialogConditionType_OkResponse":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_OkResponse, true
	case "DialogConditionType_CancelResponse":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_CancelResponse, true
	case "DialogConditionType_LastResponse":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_LastResponse, true
	case "DialogConditionType_Respond_InputArguments":
		return OpcuaNodeIdServicesVariableDialog_DialogConditionType_Respond_InputArguments, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableDialogKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableDialogValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableDialog(structType any) OpcuaNodeIdServicesVariableDialog {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableDialog {
		if sOpcuaNodeIdServicesVariableDialog, ok := typ.(OpcuaNodeIdServicesVariableDialog); ok {
			return sOpcuaNodeIdServicesVariableDialog
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableDialog) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableDialog) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableDialogParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableDialog, error) {
	return OpcuaNodeIdServicesVariableDialogParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableDialogParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableDialog, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableDialog", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableDialog")
	}
	if enum, ok := OpcuaNodeIdServicesVariableDialogByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableDialog")
		return OpcuaNodeIdServicesVariableDialog(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableDialog) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableDialog) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableDialog", 32, int32(int32(e)), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableDialog) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_Respond2_InputArguments:
		return "DialogConditionType_Respond2_InputArguments"
	case OpcuaNodeIdServicesVariableDialog_DialogResponse2MethodType_InputArguments:
		return "DialogResponse2MethodType_InputArguments"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_Prompt:
		return "DialogConditionType_Prompt"
	case OpcuaNodeIdServicesVariableDialog_DialogResponseMethodType_InputArguments:
		return "DialogResponseMethodType_InputArguments"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState:
		return "DialogConditionType_EnabledState"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_Id:
		return "DialogConditionType_EnabledState_Id"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_Name:
		return "DialogConditionType_EnabledState_Name"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_Number:
		return "DialogConditionType_EnabledState_Number"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_EffectiveDisplayName:
		return "DialogConditionType_EnabledState_EffectiveDisplayName"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_TransitionTime:
		return "DialogConditionType_EnabledState_TransitionTime"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_EffectiveTransitionTime:
		return "DialogConditionType_EnabledState_EffectiveTransitionTime"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_TrueState:
		return "DialogConditionType_EnabledState_TrueState"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_EnabledState_FalseState:
		return "DialogConditionType_EnabledState_FalseState"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState:
		return "DialogConditionType_DialogState"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_Id:
		return "DialogConditionType_DialogState_Id"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_Name:
		return "DialogConditionType_DialogState_Name"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_Number:
		return "DialogConditionType_DialogState_Number"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_EffectiveDisplayName:
		return "DialogConditionType_DialogState_EffectiveDisplayName"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_TransitionTime:
		return "DialogConditionType_DialogState_TransitionTime"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_EffectiveTransitionTime:
		return "DialogConditionType_DialogState_EffectiveTransitionTime"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_TrueState:
		return "DialogConditionType_DialogState_TrueState"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_DialogState_FalseState:
		return "DialogConditionType_DialogState_FalseState"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_ResponseOptionSet:
		return "DialogConditionType_ResponseOptionSet"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_DefaultResponse:
		return "DialogConditionType_DefaultResponse"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_OkResponse:
		return "DialogConditionType_OkResponse"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_CancelResponse:
		return "DialogConditionType_CancelResponse"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_LastResponse:
		return "DialogConditionType_LastResponse"
	case OpcuaNodeIdServicesVariableDialog_DialogConditionType_Respond_InputArguments:
		return "DialogConditionType_Respond_InputArguments"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableDialog) String() string {
	return e.PLC4XEnumName()
}
