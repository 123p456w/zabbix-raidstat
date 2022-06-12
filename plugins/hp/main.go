package main

import (
	"fmt"

	functions ".."
)

// GetControllersIDs - get number of controllers in the system
func GetControllersIDs(execPath string) []string {
	inputData := functions.GetCommandOutput(execPath, "ctrl", "all", "show")
	return functions.GetRegexpAllSubmatch(inputData, "in Slot (.*?)[\\s]")
}

// GetLogicalDrivesIDs - get number of logical drives for controller with ID 'controllerID'
func GetLogicalDrivesIDs(execPath string, controllerID string) []string {
	inputData := functions.GetCommandOutput(execPath, "ctrl", fmt.Sprintf("slot=%s", controllerID), "ld", "all", "show")
	return functions.GetRegexpAllSubmatch(inputData, "logicaldrive (.*?)[\\s]")
}

// GetPhysicalDrivesIDs - get number of physical drives for controller with ID 'controllerID'
func GetPhysicalDrivesIDs(execPath string, controllerID string) []string {
	inputData := functions.GetCommandOutput(execPath, "ctrl", fmt.Sprintf("slot=%s", controllerID), "pd", "all", "show")
	return functions.GetRegexpAllSubmatch(inputData, "physicaldrive (.*?)[\\s]")
}

// GetControllerStatus - get controller status
func GetControllerStatus(execPath string, controllerID string, indent int) []byte {
	type ReturnData struct {
		Status        string `json:"status"`
		Model         string `json:"model"`
		CacheStatus   string `json:"cachestatus"`
		BatteryStatus string `json:"batterystatus"`
	}

	inputData := functions.GetCommandOutput(execPath, "ctrl", fmt.Sprintf("slot=%s", controllerID), "show", "status")
	status := functions.GetRegexpSubmatch(inputData, "Controller Status *: (.*)")
	model := functions.GetRegexpSubmatch(inputData, "(.*) in Slot")
	cachestatus := functions.GetRegexpSubmatch(inputData, "Cache Status *: (.*)")
	batteryStatus := functions.GetRegexpSubmatch(inputData, "Battery/Capacitor Status *: (.*)")

	data := ReturnData{
		Status:        functions.TrimSpacesLeftAndRight(status),
		Model:         functions.TrimSpacesLeftAndRight(model),
		CacheStatus:   functions.TrimSpacesLeftAndRight(cachestatus),
		BatteryStatus: functions.TrimSpacesLeftAndRight(batteryStatus),
	}

	return append(functions.MarshallJSON(data, indent), "\n"...)
}

// GetLDStatus - get logical drive status
func GetLDStatus(execPath string, controllerID string, deviceID string, indent int) []byte {
	type ReturnData struct {
		Status string `json:"status"`
		Size   string `json:"size"`
		MediaErrors string `json:"mediaerrors"`
		Caching string `json:"caching"`
		ParityInitStatus string `json:"parityinitstatus"`
		ParityInitProgress string `json:"parityinitprogress"`
		RAIDLevel string `json:"raidlevel"`
	}

	inputData := functions.GetCommandOutput(execPath, "ctrl", fmt.Sprintf("slot=%s", controllerID), "ld", deviceID, "show", "detail")
	status := functions.GetRegexpSubmatch(inputData, "Status *: (.*)")
	size := functions.GetRegexpSubmatch(inputData, "Size *: (.*)")
	mediaerrors := functions.GetRegexpSubmatch(inputData, "Unrecoverable Media Errors *: (.*)")
	caching := functions.GetRegexpSubmatch(inputData, "Caching *: (.*)")
	paritystatus := functions.GetRegexpSubmatch(inputData, "Parity Initialization Status *: (.*)")
	parityprogress := functions.GetRegexpSubmatch(inputData, "Parity Initialization Progress *: (.*)")
	raidlevel := functions.GetRegexpSubmatch(inputData, "Fault Tolerance *: (.*)")

	data := ReturnData{
		Status: functions.TrimSpacesLeftAndRight(status),
		Size:   functions.TrimSpacesLeftAndRight(size),
		MediaErrors:   functions.TrimSpacesLeftAndRight(mediaerrors),
		Caching: functions.TrimSpacesLeftAndRight(caching),
		ParityInitStatus: functions.TrimSpacesLeftAndRight(paritystatus),
		ParityInitProgress: functions.TrimSpacesLeftAndRight(parityprogress),
		RAIDLevel: functions.TrimSpacesLeftAndRight(raidlevel),
	}

	return append(functions.MarshallJSON(data, indent), "\n"...)
}

// GetPDStatus - get physical drive status
func GetPDStatus(execPath string, controllerID string, deviceID string, indent int) []byte {
	type ReturnData struct {
		Status             string `json:"status"`
		Model              string `json:"model"`
		Size               string `json:"size"`
		CurrentTemperature string `json:"currenttemperature"`
		MaximumTemperature string `json:"maximumtemperature"`
		Usageremaining     string `json:"usageremaining"`
		PowerOnHours       string `json:"poweronhours"`
		DriveType          string `json:"drivetype"`
	}

	inputData := functions.GetCommandOutput(execPath, "ctrl", fmt.Sprintf("slot=%s", controllerID), "pd", deviceID, "show", "detail")
	status := functions.GetRegexpSubmatch(inputData, "[\\s]{2}Status: (.*)")
	model := functions.GetRegexpSubmatch(inputData, "Model: (.*)")
	size := functions.GetRegexpSubmatch(inputData, "[\\s]{2}Size: (.*)")
	currentTemperature := functions.GetRegexpSubmatch(inputData, "Current Temperature \\(C\\): (.*)")
	maximumTemperature := functions.GetRegexpSubmatch(inputData, "Maximum Temperature \\(C\\): (.*)")
	usageremaining := functions.GetRegexpSubmatch(inputData, "Usage remaining: (.*)")
	poweronhours := functions.GetRegexpSubmatch(inputData, "Power On Hours: (.*)")
	drivetype := functions.GetRegexpSubmatch(inputData, "Drive Type: (.*)")

	data := ReturnData{
		Status:             functions.TrimSpacesLeftAndRight(status),
		Model:              functions.TrimSpacesLeftAndRight(model),
		Size:               functions.TrimSpacesLeftAndRight(size),
		CurrentTemperature: functions.TrimSpacesLeftAndRight(currentTemperature),
		MaximumTemperature: functions.TrimSpacesLeftAndRight(maximumTemperature),
		Usageremaining:     functions.TrimSpacesLeftAndRight(usageremaining),
		PowerOnHours:       functions.TrimSpacesLeftAndRight(poweronhours),
		DriveType:          functions.TrimSpacesLeftAndRight(drivetype),
	}

	return append(functions.MarshallJSON(data, indent), "\n"...)
}

func main() {}
