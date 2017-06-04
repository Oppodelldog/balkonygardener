package sensor

import (
	"testing"
	"reflect"
)

func Test_parseSensorInfo_constructive(t *testing.T) {
	type TestDataSetup struct {
		input       string
		expectation *SensorInfo
	}
	testDataSetups := []TestDataSetup{
		{input: "A0:12345", expectation: &SensorInfo{Name: "A0", Value: 12345}},
		{input: "A0:12345.987", expectation: &SensorInfo{Name: "A0", Value: 12345.987}},
	}

	for _, testDataSetup := range testDataSetups {

		err, sensorInfo := parseSensorInfo(testDataSetup.input)

		if err != nil {
			t.Fatalf("did not expect an error, but got: %v", err)
		}

		if reflect.DeepEqual(sensorInfo, testDataSetup.expectation) == false {
			t.Fatalf("expected %v, but got %v", testDataSetup.expectation, sensorInfo)
		}
	}
}

func Test_parseSensorInfo_destructive(t *testing.T) {
	type TestDataSetup struct {
		input       string
	}
	testDataSetups := []TestDataSetup{
		{input: "A0:345"},
		{input: "A0:AAA.AAA"},
		{input: "A012345"},
		{input: "A012345.987"},
		{input: ".987"},
		{input: "."},
		{input: ":."},
		{input: ""},
	}

	for _, testDataSetup := range testDataSetups {

		err, sensorInfo := parseSensorInfo(testDataSetup.input)

		if sensorInfo != nil {
			t.Fatalf("expected nil sensorInfo, but got: %v", sensorInfo)
		}
		if err == nil {
			t.Fatalf("expected an error, but got: %v", err)
		}
	}
}
