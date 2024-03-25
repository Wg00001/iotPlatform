package api

import (
	"encoding/json"
	"errors"
	"iotPlatform/common"
	"iotPlatform/helper"
)

func CreateAuthUser(in *CreateAuthUserRequest) error {
	data, _ := json.Marshal(in)
	rep, err := helper.HttpPost(common.EmqxAddr+"/authentication/password_based%3Abuilt_in_database/users", data)
	if err != nil {
		return err
	}
	reply := new(CreateAuthUserReply)
	err = json.Unmarshal(rep, reply)
	if err != nil {
		return errors.New("error client exit")
	}
	return nil
}

func DeleteAuthUser(clientId string) error {
	rep, err := helper.HttpDelete(common.EmqxAddr+"/authentication/password_based%3Abuilt_in_database/users/"+clientId, []byte{})
	if err != nil {
		return err
	}
	if len(rep) > 0 {
		return errors.New("error client not found")
	}
	return nil
}
