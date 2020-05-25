package utils

import (
	"strconv"
	"strings"
)

func contains(arr []string, value string) bool {
	for _, element := range arr {
		if element == value {
			return true
		}
	}
	return false
}

func parser(strContent string) ([]string, bool) {
	stringArr := strings.Split(strContent, seperator)
	if len(stringArr) < 3 {
		return nil, false
	}

	// remove first 2 element (*1 $2)
	stringArr = stringArr[2:]
	newStringArr := make([]string, 0)
	for _, elm := range stringArr { // remove space
		if elm[0] == '$' {
			continue
		}

		newStringArr = append(newStringArr, elm)
	}

	// check command exists
	if !contains(commandList, strings.ToUpper(newStringArr[0])) {
		return nil, false
	}

	return newStringArr, true
}

// CommandPing ...
func CommandPing(strContent *[]string) string {
	if len(*strContent) == 2 {
		return prefString + (*strContent)[1]
	}

	return prefString + "PONG"
}

// CommandGet ...
func CommandGet(strContent *[]string, mapObj *RedisMap) string {
	if len(*strContent) == 1 {
		return msgErrorDefault
	}

	if val, ok := mapObj.Get((*strContent)[1]); ok {
		return prefBulk + strconv.Itoa(len(val)) + prefCRLF + val
	}

	return prefBulk + "-1"
}

// CommandSet ...
func CommandSet(strContent *[]string, mapObj *RedisMap) string {
	if len(*strContent) < 3 { // minimum requirement: set variable value
		return msgErrorDefault
	}

	if len(*strContent) == 4 {
		if strings.ToUpper((*strContent)[3]) == "NX" { // Only set the key if it does not already exist
			if _, ok := mapObj.Get((*strContent)[1]); ok {
				return msgNil
			}

			mapObj.Set((*strContent)[1], (*strContent)[2])
			return msgOK
		} else if strings.ToUpper((*strContent)[3]) == "XX" { // Only set the key if it already exist.
			if _, ok := mapObj.Get((*strContent)[1]); ok {
				mapObj.Set((*strContent)[1], (*strContent)[2])
				return msgOK
			}

			return msgNil
		}
	} else {
		mapObj.Set((*strContent)[1], (*strContent)[2])
		return msgOK
	}

	return msgNil
}

// CommandDel ...
func CommandDel(strContent *[]string, mapObj *RedisMap) string {
	if len(*strContent) == 1 {
		return msgErrorDefault
	}

	delNum := 0
	for i := 1; i < len(*strContent); i++ {
		if _, ok := mapObj.Get((*strContent)[i]); ok {
			mapObj.Delete((*strContent)[i])
			delNum++
		}
	}

	return prefInteger + strconv.Itoa(delNum)
}

// CommandIncrby ...
func CommandIncrby(strContent *[]string, mapObj *RedisMap) string {
	if len(*strContent) != 3 {
		return msgErrorDefault
	}

	result, isSuccess := mapObj.Increment((*strContent)[1], (*strContent)[2])
	if isSuccess {
		return prefInteger + result
	}

	return prefString + result
}

// DoLogic redis basic operation
func DoLogic(prmString string, mapObj *RedisMap) string {
	strContent, isSuccess := parser(prmString)
	if !isSuccess {
		return msgErrorDefault + prefCRLF
	}

	// fmt.Println(strContent)
	var response string
	rCommand := strings.ToUpper(strContent[0])
	switch rCommand {
	case "PING":
		response = CommandPing(&strContent)
	case "GET":
		response = CommandGet(&strContent, mapObj)
	case "SET":
		response = CommandSet(&strContent, mapObj)
	case "DEL":
		response = CommandDel(&strContent, mapObj)
	case "INCRBY":
		response = CommandIncrby(&strContent, mapObj)
	default:
		return response
	}

	return response + prefCRLF
}
