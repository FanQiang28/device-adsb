// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2019 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package mock

import (
	"context"
	"encoding/json"

	contract "github.com/xinminsu/go-mod-core-contracts/models"
)

var (
	DeviceProfileRandomBoolGenerator           = contract.DeviceProfile{}
	DeviceProfileRandomIntegerGenerator        = contract.DeviceProfile{}
	DeviceProfileRandomUnsignedGenerator       = contract.DeviceProfile{}
	DeviceProfileRandomFloatGenerator          = contract.DeviceProfile{}
	DuplicateDeviceProfileRandomFloatGenerator = contract.DeviceProfile{}
	NewDeviceProfile                           = contract.DeviceProfile{}
)

type DeviceProfileClientMock struct{}

func (DeviceProfileClientMock) Add(dp *contract.DeviceProfile, ctx context.Context) (string, error) {
	panic("implement me")
}

func (DeviceProfileClientMock) Delete(id string, ctx context.Context) error {
	panic("implement me")
}

func (DeviceProfileClientMock) DeleteByName(name string, ctx context.Context) error {
	panic("implement me")
}

func (DeviceProfileClientMock) DeviceProfile(id string, ctx context.Context) (contract.DeviceProfile, error) {
	panic("implement me")
}

func (DeviceProfileClientMock) DeviceProfiles(ctx context.Context) ([]contract.DeviceProfile, error) {
	populateDeviceProfileMock()
	return []contract.DeviceProfile{
		DeviceProfileRandomBoolGenerator,
		DeviceProfileRandomIntegerGenerator,
		DeviceProfileRandomUnsignedGenerator,
		DeviceProfileRandomFloatGenerator,
	}, nil
}

func (DeviceProfileClientMock) DeviceProfileForName(name string, ctx context.Context) (contract.DeviceProfile, error) {
	panic("implement me")
}

func (DeviceProfileClientMock) Update(dp contract.DeviceProfile, ctx context.Context) error {
	panic("implement me")
}

func (DeviceProfileClientMock) Upload(yamlString string, ctx context.Context) (string, error) {
	panic("implement me")
}

func (DeviceProfileClientMock) UploadFile(yamlFilePath string, ctx context.Context) (string, error) {
	panic("implement me")
}

func populateDeviceProfileMock() {
	deviceProfileDataRandomBoolGenerator := `{"created":1551711642676,"modified":1551711642676,"origin":0,"description":"Example of Device-Virtual","id":"b06e124f-3e46-483d-b18b-fc1f93f835c6","name":"Random-Boolean-Generator","manufacturer":"IOTech","model":"Device-Virtual-01","labels":["device-virtual-example"],"objects":null,"deviceResources":[{"description":"used to decide whether to re-generate a random value","name":"Enable_Randomization","tag":null,"properties":{"value":{"type":"Bool","readWrite":"W","minimum":null,"maximum":null,"defaultValue":"true","size":null,"word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":null},"units":{"type":"String","readWrite":"R","defaultValue":"Random"}},"attributes":null},{"description":"Generate random boolean value","name":"RandomValue_Bool","tag":null,"properties":{"value":{"type":"Bool","readWrite":"R","minimum":"false","maximum":"true","defaultValue":"true","size":null,"word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":null},"units":{"type":"String","readWrite":"R","defaultValue":"random bool value"}},"attributes":null}],"deviceCommands":[{"name":"RandomValue_Bool","get":[{"index":null,"operation":"get","object":"RandomValue_Bool","parameter":"RandomValue_Bool","resource":null,"secondary":[],"mappings":{}}],"set":[{"index":null,"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"Enable_Randomization","secondary":[],"mappings":{}},{"index":null,"operation":"set","object":"RandomValue_Bool","parameter":"RandomValue_Bool","resource":"RandomValue_Bool","secondary":[],"mappings":{}}]}],"coreCommands":[{"created":1551711642676,"modified":0,"origin":0,"id":"ed0d88b8-e202-4bdd-a941-606b5c7f40d4","name":"RandomValue_Bool","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Bool","responses":[{"code":"200","description":null,"expectedValues":["RandomValue_Bool"]},{"code":"503","description":"service unavailable","expectedValues":[]}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Bool","responses":[{"code":"200","description":null,"expectedValues":[]},{"code":"503","description":"service unavailable","expectedValues":[]}],"parameterNames":["RandomValue_Bool","Enable_Randomization"]}}]}`
	deviceProfileDataRandomIntegerGenerator := `{"created":1553721583525,"modified":1553721583525,"description":"Example of Device-Virtual","id":"1da96459-4fef-4860-9980-a653c28c5ff6","name":"Random-Integer-Generator","manufacturer":"IOTech","model":"Device-Virtual-01","labels":["device-virtual-example"],"deviceResources":[{"description":"used to decide whether to re-generate a random value","name":"Enable_Randomization","properties":{"value":{"type":"Bool","readWrite":"W","defaultValue":"true"},"units":{"type":"String","readWrite":"R","defaultValue":"Random"}}},{"description":"Generate random int8 value","name":"RandomValue_Int8","properties":{"value":{"type":"Int8","readWrite":"R","defaultValue":"0"},"units":{"type":"String","readWrite":"R","defaultValue":"random int8 value"}}},{"description":"Generate random int16 value","name":"RandomValue_Int16","properties":{"value":{"type":"Int16","readWrite":"R","defaultValue":"0"},"units":{"type":"String","readWrite":"R","defaultValue":"random int16 value"}}},{"description":"Generate random int32 value","name":"RandomValue_Int32","properties":{"value":{"type":"Int32","readWrite":"R","defaultValue":"0"},"units":{"type":"String","readWrite":"R","defaultValue":"random int32 value"}}},{"description":"Generate random int64 value","name":"RandomValue_Int64","properties":{"value":{"type":"Int64","readWrite":"R","defaultValue":"0"},"units":{"type":"String","readWrite":"R","defaultValue":"random int64 value"}}},{"description":"ResourceTestTransform_Pass","name":"ResourceTestTransform_Pass","properties":{"value":{"type":"Int8","readWrite":"RW","defaultValue":"0","offset":"1.0"},"units":{"type":"String","readWrite":"R"}}},{"description":"ResourceTestTransform_Fail","name":"ResourceTestTransform_Fail","properties":{"value":{"type":"Int8","readWrite":"RW","defaultValue":"0","offset":"error"},"units":{"type":"String","readWrite":"R"}}},{"description":"ResourceTestAssertion_Pass","name":"ResourceTestAssertion_Pass","properties":{"value":{"type":"Int8","readWrite":"RW","defaultValue":"0","assertion":"123"},"units":{"type":"String","readWrite":"R"}}},{"description":"ResourceTestAssertion_Fail","name":"ResourceTestAssertion_Fail","properties":{"value":{"type":"Int8","readWrite":"RW","defaultValue":"0","assertion":"12"},"units":{"type":"String","readWrite":"R"}}},{"description":"ResourceTestMapping_Pass","name":"ResourceTestMapping_Pass","properties":{"value":{"type":"Int8","readWrite":"RW","defaultValue":"0"},"units":{"type":"String","readWrite":"R"}}},{"description":"ResourceTestMapping_Fail","name":"ResourceTestMapping_Fail","properties":{"value":{"type":"Int8","readWrite":"RW","defaultValue":"0"},"units":{"type":"String","readWrite":"R"}}},{"description":"NoDeviceResourceForResult","name":"NoDeviceResourceForResult","properties":{"value":{"type":"String","readWrite":"RW"},"units":{"type":"String","readWrite":"R"}}},{"description":"Error","name":"Error","properties":{"value":{"type":"String","readWrite":"RW"},"units":{"type":"String","readWrite":"R"}}}],"deviceCommands":[{"name":"RandomValue_Int8","get":[{"operation":"get","object":"RandomValue_Int8","parameter":"RandomValue_Int8"}],"set":[{"operation":"set","object":"RandomValue_Int8","parameter":"RandomValue_Int8","resource":"RandomValue_Int8"},{"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Int8"}]},{"name":"RandomValue_Int16","get":[{"operation":"get","object":"RandomValue_Int16","parameter":"RandomValue_Int16"}],"set":[{"operation":"set","object":"RandomValue_Int16","parameter":"RandomValue_Int16","resource":"RandomValue_Int16"},{"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Int16"}]},{"name":"RandomValue_Int32","get":[{"operation":"get","object":"RandomValue_Int32","parameter":"RandomValue_Int32"}],"set":[{"operation":"set","object":"RandomValue_Int32","parameter":"RandomValue_Int32","resource":"RandomValue_Int32"},{"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Int32"}]},{"name":"RandomValue_Int64","get":[{"operation":"get","object":"RandomValue_Int64","parameter":"RandomValue_Int64"}],"set":[{"operation":"set","object":"RandomValue_Int64","parameter":"RandomValue_Int64","resource":"RandomValue_Int64"},{"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Int64"}]},{"name":"ResourceTestTransform_Fail","get":[{"operation":"get","object":"ResourceTestTransform_Fail","parameter":"ResourceTestTransform_Fail"}],"set":[{"operation":"set","object":"ResourceTestTransform_Fail","parameter":"ResourceTestTransform_Fail"}]},{"name":"ResourceTestAssertion_Pass","get":[{"operation":"get","object":"ResourceTestAssertion_Pass","parameter":"ResourceTestAssertion_Pass"}],"set":[{"operation":"set","object":"ResourceTestAssertion_Pass","parameter":"ResourceTestAssertion_Pass"}]},{"name":"ResourceTestAssertion_Fail","get":[{"operation":"get","object":"ResourceTestAssertion_Fail","parameter":"ResourceTestAssertion_Fail"}],"set":[{"operation":"set","object":"ResourceTestAssertion_Fail","parameter":"ResourceTestAssertion_Fail"}]},{"name":"ResourceTestMapping_Pass","get":[{"operation":"get","object":"ResourceTestMapping_Pass","parameter":"ResourceTestMapping_Pass","mappings":{"123":"Pass"}}],"set":[{"operation":"set","object":"ResourceTestMapping_Pass","parameter":"ResourceTestMapping_Pass","mappings":{"Pass":"123"}}]},{"name":"ResourceTestMapping_Fail","get":[{"operation":"get","object":"ResourceTestMapping_Fail","parameter":"ResourceTestMapping_Fail","mappings":{"12":"Pass"}}],"set":[{"operation":"set","object":"ResourceTestMapping_Fail","parameter":"ResourceTestMapping_Fail","mappings":{"Pass":"12"}}]},{"name":"NoDeviceResourceForOperation","get":[{"operation":"get","object":"ResourceNotFound","parameter":"Error"}],"set":[{"operation":"set","object":"ResourceNotFound","parameter":"Error"}]},{"name":"NoDeviceResourceForResult","get":[{"operation":"get","object":"NoDeviceResourceForResult","parameter":"NoDeviceResourceForResult"}]},{"name":"Error","get":[{"index":"1","operation":"get","object":"Error","parameter":"Error"},{"index":"2","operation":"get","object":"Error","parameter":"Error"}],"set":[{"index":"1","operation":"set","object":"Error","parameter":"Error"},{"index":"2","operation":"set","object":"Error","parameter":"Error"}]}],"coreCommands":[{"created":1553721583522,"modified":1553721583522,"id":"6b2c4f1d-9537-4cf6-a25b-9f0bf7761be6","name":"RandomValue_Int8","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Int8","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Int8","parameterNames":["RandomValue_Int8","Enable_Randomization"]}},{"created":1553721583523,"modified":1553721583523,"id":"b5b3d990-7e86-4222-a92a-cacde97a23a2","name":"RandomValue_Int16","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Int16","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Int16","parameterNames":["RandomValue_Int16","Enable_Randomization"]}},{"created":1553721583523,"modified":1553721583523,"id":"d0ce104e-0582-4651-9c8a-86e51c4b1180","name":"RandomValue_Int32","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Int32","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Int32","parameterNames":["RandomValue_Int32","Enable_Randomization"]}},{"created":1553721583523,"modified":1553721583523,"id":"8105e902-194a-49b1-a71f-296a0057dfff","name":"RandomValue_Int64","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Int64","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Int64","parameterNames":["RandomValue_Int64","Enable_Randomization"]}},{"created":1553721583523,"modified":1553721583523,"id":"86f7074f-4350-4b2c-913b-69288abeec79","name":"ResourceTestTransform_Fail","get":{"path":"/api/v1/device/{deviceId}/ResourceTestTransform_Fail","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/ResourceTestTransform_Fail","parameterNames":["ResourceTestTransform_Fail"]}},{"created":1553721583523,"modified":1553721583523,"id":"02819de2-9011-4991-9ce6-f4f7f1990a79","name":"ResourceTestAssertion_Fail","get":{"path":"/api/v1/device/{deviceId}/ResourceTestAssertion_Fail","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/ResourceTestAssertion_Fail","parameterNames":["ResourceTestAssertion_Fail"]}},{"created":1553721583523,"modified":1553721583523,"id":"cc625b93-7655-4e5b-920c-77387a4ec3a2","name":"ResourceTestMapping_Pass","get":{"path":"/api/v1/device/{deviceId}/ResourceTestMapping_Pass","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/ResourceTestMapping_Pass","parameterNames":["ResourceTestMapping_Pass"]}},{"created":1553721583523,"modified":1553721583523,"id":"b0b4d373-430d-46f7-88cc-cd392aaa5c90","name":"ResourceTestMapping_Fail","get":{"path":"/api/v1/device/{deviceId}/ResourceTestMapping_Fail","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/ResourceTestMapping_Fail","parameterNames":["ResourceTestMapping_Fail"]}},{"created":1553721583523,"modified":1553721583523,"id":"ffa2c711-49ea-4e8a-84ad-0e00ef69de3f","name":"NoDeviceResourceForOperation","get":{"path":"/api/v1/device/{deviceId}/NoDeviceResourceForOperation","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/NoDeviceResourceForOperation","parameterNames":["Error"]}},{"created":1553721583524,"modified":1553721583524,"id":"32841000-5fa1-4a27-975f-18456f564149","name":"NoDeviceResourceForResult","get":{"path":"/api/v1/device/{deviceId}/NoDeviceResourceForResult","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/NoDeviceResourceForResult","parameterNames":["NoDeviceResourceForResult"]}},{"created":1553721583524,"modified":1553721583524,"id":"2f2cc0cc-7cd1-4a18-9202-df0d049b94de","name":"Error","get":{"path":"/api/v1/device/{deviceId}/Error","responses":[{"code":"503","description":"service unavailable"}]},"put":{"path":"/api/v1/device/{deviceId}/Error","parameterNames":["Error","Error"]}}]}`
	deviceProfileDataRandomUnsignedIntegerGenerator := `{"created":1551711642733,"modified":1551711642733,"origin":0,"description":"Example of Device-Virtual","id":"e436c8e8-0ba4-4257-84b0-1f718cadd33e","name":"Random-UnsignedInteger-Generator","manufacturer":"IOTech","model":"Device-Virtual-01","labels":["device-virtual-example"],"objects":null,"deviceResources":[{"description":"used to decide whether to re-generate a random value","name":"Enable_Randomization","tag":null,"properties":{"value":{"type":"Bool","readWrite":"W","minimum":null,"maximum":null,"defaultValue":"true","size":null,"word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":null},"units":{"type":"String","readWrite":"R","defaultValue":"Random"}},"attributes":null},{"description":"Generate random uint8 value","name":"RandomValue_Uint8","tag":null,"properties":{"value":{"type":"Uint8","readWrite":"R","minimum":"0","maximum":"100","defaultValue":"0","size":null,"word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":null},"units":{"type":"String","readWrite":"R","defaultValue":"random uint8 value"}},"attributes":null},{"description":"Generate random uint16 value","name":"RandomValue_Uint16","tag":null,"properties":{"value":{"type":"Uint16","readWrite":"R","minimum":"0","maximum":"100","defaultValue":"0","size":null,"word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":null},"units":{"type":"String","readWrite":"R","defaultValue":"random uint16 value"}},"attributes":null},{"description":"Generate random uint32 value","name":"RandomValue_Uint32","tag":null,"properties":{"value":{"type":"Uint32","readWrite":"R","minimum":"0","maximum":"100","defaultValue":"0","size":null,"word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":null},"units":{"type":"String","readWrite":"R","defaultValue":"random uint32 value"}},"attributes":null},{"description":"Generate random uint64 value","name":"RandomValue_Uint64","tag":null,"properties":{"value":{"type":"Uint64","readWrite":"R","minimum":"0","maximum":"100","defaultValue":"0","size":null,"word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":null},"units":{"type":"String","readWrite":"R","defaultValue":"random uint64 value"}},"attributes":null}],"deviceCommands":[{"name":"RandomValue_Uint8","get":[{"index":null,"operation":"get","object":"RandomValue_Uint8","parameter":"RandomValue_Uint8","resource":null,"secondary":[],"mappings":{}}],"set":[{"index":null,"operation":"set","object":"RandomValue_Uint8","parameter":"RandomValue_Uint8","resource":"RandomValue_Uint8","secondary":[],"mappings":{}},{"index":null,"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Uint8","secondary":[],"mappings":{}}]},{"name":"RandomValue_Uint16","get":[{"index":null,"operation":"get","object":"RandomValue_Uint16","parameter":"RandomValue_Uint16","resource":null,"secondary":[],"mappings":{}}],"set":[{"index":null,"operation":"set","object":"RandomValue_Uint16","parameter":"RandomValue_Uint16","resource":"RandomValue_Uint16","secondary":[],"mappings":{}},{"index":null,"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Uint16","secondary":[],"mappings":{}}]},{"name":"RandomValue_Uint32","get":[{"index":null,"operation":"get","object":"RandomValue_Uint32","parameter":"RandomValue_Uint32","resource":null,"secondary":[],"mappings":{}}],"set":[{"index":null,"operation":"set","object":"RandomValue_Uint32","parameter":"RandomValue_Uint32","resource":"RandomValue_Uint32","secondary":[],"mappings":{}},{"index":null,"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Uint32","secondary":[],"mappings":{}}]},{"name":"RandomValue_Uint64","get":[{"index":null,"operation":"get","object":"RandomValue_Uint64","parameter":"RandomValue_Uint64","resource":null,"secondary":[],"mappings":{}}],"set":[{"index":null,"operation":"set","object":"RandomValue_Uint64","parameter":"RandomValue_Uint64","resource":"RandomValue_Uint64","secondary":[],"mappings":{}},{"index":null,"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Uint64","secondary":[],"mappings":{}}]}],"coreCommands":[{"created":1551711642730,"modified":0,"origin":0,"id":"ced85b30-3ef0-45fa-a57e-97108b89ca67","name":"RandomValue_Uint8","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Uint8","responses":[{"code":"200","description":null,"expectedValues":["RandomValue_Uint8"]},{"code":"503","description":"service unavailable","expectedValues":[]}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Uint8","responses":[{"code":"200","description":null,"expectedValues":[]},{"code":"503","description":"service unavailable","expectedValues":[]}],"parameterNames":["RandomValue_Uint8","Enable_Randomization"]}},{"created":1551711642730,"modified":0,"origin":0,"id":"fbc90059-b10d-482d-8cf5-012b35758789","name":"RandomValue_Uint16","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Uint16","responses":[{"code":"200","description":null,"expectedValues":["RandomValue_Uint16"]},{"code":"503","description":"service unavailable","expectedValues":[]}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Uint16","responses":[{"code":"200","description":null,"expectedValues":[]},{"code":"503","description":"service unavailable","expectedValues":[]}],"parameterNames":["RandomValue_Uint16","Enable_Randomization"]}},{"created":1551711642733,"modified":0,"origin":0,"id":"27ac1912-5799-4141-916e-4030c59cc56b","name":"RandomValue_Uint32","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Uint32","responses":[{"code":"200","description":null,"expectedValues":["RandomValue_Uint32"]},{"code":"503","description":"service unavailable","expectedValues":[]}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Uint32","responses":[{"code":"200","description":null,"expectedValues":[]},{"code":"503","description":"service unavailable","expectedValues":[]}],"parameterNames":["RandomValue_Uint32","Enable_Randomization"]}},{"created":1551711642733,"modified":0,"origin":0,"id":"c7a53c56-f563-471e-840c-4e48c59dc362","name":"RandomValue_Uint64","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Uint64","responses":[{"code":"200","description":null,"expectedValues":["RandomValue_Uint64"]},{"code":"503","description":"service unavailable","expectedValues":[]}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Uint64","responses":[{"code":"200","description":null,"expectedValues":[]},{"code":"503","description":"service unavailable","expectedValues":[]}],"parameterNames":["RandomValue_Uint64","Enable_Randomization"]}}]}`
	deviceProfileDataRandomFloatGenerator := `{"created":1551711642693,"modified":1551711642693,"origin":0,"description":"Example of Device-Virtual","id":"08dc1ab9-a76f-4b61-b177-6fa1c44e6424","name":"Random-Float-Generator","manufacturer":"IOTech","model":"Device-Virtual-01","labels":["device-virtual-example"],"objects":null,"deviceResources":[{"description":"used to decide whether to re-generate a random value","name":"Enable_Randomization","tag":null,"properties":{"value":{"type":"Bool","readWrite":"W","minimum":null,"maximum":null,"defaultValue":"true","size":null,"word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":null},"units":{"type":"String","readWrite":"R","defaultValue":"Random"}},"attributes":null},{"description":"Generate random float32 value","name":"RandomValue_Float32","tag":null,"properties":{"value":{"type":"Float32","readWrite":"R","minimum":"0","maximum":"100","defaultValue":"3.14159","size":null,"word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":null},"units":{"type":"String","readWrite":"R","defaultValue":"random float32 value"}},"attributes":null},{"description":"Generate random float64 value","name":"RandomValue_Float64","tag":null,"properties":{"value":{"type":"Float64","readWrite":"R","minimum":"0","maximum":"100","defaultValue":"3.14159265359","size":null,"word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":null},"units":{"type":"String","readWrite":"R","defaultValue":"random float64 value"}},"attributes":null}],"deviceCommands":[{"name":"RandomValue_Float32","get":[{"index":null,"operation":"get","object":"RandomValue_Float32","parameter":"RandomValue_Float32","resource":null,"secondary":[],"mappings":{}}],"set":[{"index":null,"operation":"set","object":"RandomValue_Float32","parameter":"RandomValue_Float32","resource":"RandomValue_Float32","secondary":[],"mappings":{}},{"index":null,"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Float32","secondary":[],"mappings":{}}]},{"name":"RandomValue_Float64","get":[{"index":null,"operation":"get","object":"RandomValue_Float64","parameter":"RandomValue_Float64","resource":null,"secondary":[],"mappings":{}}],"set":[{"index":null,"operation":"set","object":"RandomValue_Float64","parameter":"RandomValue_Float64","resource":"RandomValue_Float64","secondary":[],"mappings":{}},{"index":null,"operation":"set","object":"Enable_Randomization","parameter":"Enable_Randomization","resource":"RandomValue_Float64","secondary":[],"mappings":{}}]}],"coreCommands":[{"created":1551711642693,"modified":0,"origin":0,"id":"e724174f-8226-4ad2-a5b6-826ff9bfc851","name":"RandomValue_Float32","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Float32","responses":[{"code":"200","description":null,"expectedValues":["RandomValue_Float32"]},{"code":"503","description":"service unavailable","expectedValues":[]}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Float32","responses":[{"code":"200","description":null,"expectedValues":[]},{"code":"503","description":"service unavailable","expectedValues":[]}],"parameterNames":["RandomValue_Float32","Enable_Randomization"]}},{"created":1551711642693,"modified":0,"origin":0,"id":"fa2904d9-e7eb-4304-b55c-c66d80f878c1","name":"RandomValue_Float64","get":{"path":"/api/v1/device/{deviceId}/RandomValue_Float64","responses":[{"code":"200","description":null,"expectedValues":["RandomValue_Float64"]},{"code":"503","description":"service unavailable","expectedValues":[]}]},"put":{"path":"/api/v1/device/{deviceId}/RandomValue_Float64","responses":[{"code":"200","description":null,"expectedValues":[]},{"code":"503","description":"service unavailable","expectedValues":[]}],"parameterNames":["RandomValue_Float64","Enable_Randomization"]}}]}`
	newDeviceProfileData := `{"created":1552129341568,"modified":1552129341568,"origin":1552129341469,"description":"GS1 AC Drive","id":"cce09f07-ce9a-40c8-aff4-9b023f09adae","name":"GS1-AC-Drive","manufacturer":"Automation Direct","model":"GS1-10P5","labels":["modbus","industrial"],"objects":null,"deviceResources":[{"description":"Get the OutputVoltage.","name":"HoldingRegister_8455","tag":null,"properties":{"value":{"type":"Float","readWrite":"R","minimum":"0.00","maximum":"300","defaultValue":"0.00","size":"4","word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":"3.2"},"units":{"type":"String","readWrite":"R","defaultValue":"rpm"}},"attributes":{"instance":"8455","type":"doublebyteInt"}},{"description":"Get the RPM.","name":"HoldingRegister_8454","tag":null,"properties":{"value":{"type":"Float","readWrite":"R","minimum":"0.00","maximum":"300","defaultValue":"0.00","size":"4","word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":"3.2"},"units":{"type":"String","readWrite":"R","defaultValue":"Volts"}},"attributes":{"instance":"8454","type":"doublebyteInt"}},{"description":"The status of the device.","name":"HoldingRegister_2331","tag":null,"properties":{"value":{"type":"Boolean","readWrite":"R","minimum":null,"maximum":null,"defaultValue":"false","size":"1","word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":null},"units":{"type":"String","readWrite":"R","defaultValue":"OFF"}},"attributes":{"instance":"8454","type":"doublebyteInt"}},{"description":"whether generating random value in each collection cycle","name":"enableRandomization","tag":null,"properties":{"value":{"type":"Boolean","readWrite":"W","minimum":null,"maximum":null,"defaultValue":"true","size":"1","word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":null},"units":{"type":"String","readWrite":"R","defaultValue":"Random"}},"attributes":null},{"description":"the frequency of collection","name":"collectionFrequency","tag":null,"properties":{"value":{"type":"Integer","readWrite":"W","minimum":"0","maximum":"600","defaultValue":"15","size":"4","word":"2","lsb":null,"mask":"0x00","shift":"0","scale":"1.0","offset":"0.0","base":"0","assertion":null,"signed":true,"precision":"3"},"units":{"type":"String","readWrite":"R","defaultValue":"Seconds"}},"attributes":null}],"deviceCommands":[{"name":"OutputVoltage","get":[{"index":null,"operation":"get","object":"HoldingRegister_8455","parameter":"OutputVoltage","resource":null,"secondary":[],"mappings":{}}],"set":[{"index":null,"operation":"set","object":"enableRandomization","parameter":"enableRandomization","resource":null,"secondary":[],"mappings":{}},{"index":null,"operation":"set","object":"collectionFrequency","parameter":"collectionFrequency","resource":null,"secondary":[],"mappings":{}}]},{"name":"RPM","get":[{"index":null,"operation":"get","object":"HoldingRegister_8454","parameter":"RPM","resource":null,"secondary":[],"mappings":{}}],"set":[{"index":null,"operation":"set","object":"enableRandomization","parameter":"enableRandomization","resource":null,"secondary":[],"mappings":{}},{"index":null,"operation":"set","object":"collectionFrequency","parameter":"collectionFrequency","resource":null,"secondary":[],"mappings":{}}]},{"name":"Status","get":[{"index":null,"operation":"get","object":"HoldingRegister_2331","parameter":"Status","resource":null,"secondary":[],"mappings":{}}],"set":[{"index":null,"operation":"set","object":"enableRandomization","parameter":"enableRandomization","resource":null,"secondary":[],"mappings":{}},{"index":null,"operation":"set","object":"collectionFrequency","parameter":"collectionFrequency","resource":null,"secondary":[],"mappings":{}}]}],"coreCommands":[{"created":1552129341567,"modified":0,"origin":0,"id":"21548fad-69ff-4671-8f4b-3a7d5ab18dec","name":"OutputVoltage","get":{"path":"/api/v1/device/{deviceId}/OutputVoltage","responses":[{"code":"200","description":null,"expectedValues":["HoldingRegister_8455"]},{"code":"503","description":"service unavailable","expectedValues":[]}]},"put":{"path":"/api/v1/device/{deviceId}/OutputVoltage","responses":[{"code":"204","description":"valid and accepted","expectedValues":[]},{"code":"400","description":"bad request","expectedValues":[]},{"code":"503","description":"service unavailable","expectedValues":[]}],"parameterNames":["enableRandomization","collectionFrequency"]}},{"created":1552129341568,"modified":0,"origin":0,"id":"67360b6e-b426-443d-b53c-6490237456e4","name":"Status","get":{"path":"/api/v1/device/{deviceId}/Status","responses":[{"code":"200","description":null,"expectedValues":["HoldingRegister_2331"]},{"code":"503","description":"service unavailable","expectedValues":[]}]},"put":{"path":"/api/v1/device/{deviceId}/Status","responses":[{"code":"204","description":"valid and accepted","expectedValues":[]},{"code":"400","description":"bad request","expectedValues":[]},{"code":"503","description":"service unavailable","expectedValues":[]}],"parameterNames":["enableRandomization","collectionFrequency"]}},{"created":1552129341568,"modified":0,"origin":0,"id":"3d81d381-5055-4928-a8bf-d120d646a2c2","name":"RPM","get":{"path":"/api/v1/device/{deviceId}/RPM","responses":[{"code":"200","description":null,"expectedValues":["HoldingRegister_8454"]},{"code":"503","description":"service unavailable","expectedValues":[]}]},"put":{"path":"/api/v1/device/{deviceId}/RPM","responses":[{"code":"204","description":"valid and accepted","expectedValues":[]},{"code":"400","description":"bad request","expectedValues":[]},{"code":"503","description":"service unavailable","expectedValues":[]}],"parameterNames":["enableRandomization","collectionFrequency"]}}]}`
	json.Unmarshal([]byte(deviceProfileDataRandomBoolGenerator), &DeviceProfileRandomBoolGenerator)
	json.Unmarshal([]byte(deviceProfileDataRandomIntegerGenerator), &DeviceProfileRandomIntegerGenerator)
	json.Unmarshal([]byte(deviceProfileDataRandomUnsignedIntegerGenerator), &DeviceProfileRandomUnsignedGenerator)
	json.Unmarshal([]byte(deviceProfileDataRandomFloatGenerator), &DeviceProfileRandomFloatGenerator)
	json.Unmarshal([]byte(deviceProfileDataRandomFloatGenerator), &DuplicateDeviceProfileRandomFloatGenerator)
	json.Unmarshal([]byte(newDeviceProfileData), &NewDeviceProfile)
}
