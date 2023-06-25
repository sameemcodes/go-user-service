package utils

import "go-tools/constants"

func HandleError(msg string, err error) {
	if err != nil {
		constants.Logger.Errorf(msg+" \n", err)
	}
}
