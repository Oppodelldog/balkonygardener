package sensor

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func parseSensorInfo(rawMessage string) (error, *Info) {
	sensorInfo := &Info{}
	parts := strings.Split(rawMessage, ":")
	if len(parts) != 2 {
		return errors.New(fmt.Sprintf("could not parse arduino info from: '%s'. Expected 2 parts, but got %v", rawMessage, len(parts))), nil
	}

	if len(parts[0]) != 2 {
		return errors.New(fmt.Sprintf("could not parse arduino info from: '%s'. First part must be 2 characters long but was %v", rawMessage, len(parts[0]))), nil
	}
	sensorInfo.Name = parts[0]

	if strings.Contains(parts[1], ".") {
	} else {
		if len(parts[1]) != 5 {
			return errors.New(fmt.Sprintf("could not parse arduino info from: '%s'. Second part must be 5 characters long but was %v", rawMessage, len(parts[1]))), nil
		}
	}
	f, err := strconv.ParseFloat(parts[1], 10)
	if err != nil {
		return err, nil
	}
	sensorInfo.Value = f

	return nil, sensorInfo
}
