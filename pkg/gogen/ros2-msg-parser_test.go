/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
    http://www.apache.org/licenses/LICENSE-2.0
*/

package gogen

import (
	"crypto/md5"
	"encoding/hex"
	"testing"

	"github.com/bradleyjkemp/cupaloy/v2"
	. "github.com/smartystreets/goconvey/convey"
)

func matchSnapshot(name string, x interface{}) {
	sum := md5.Sum([]byte(name))
	So(cupaloy.SnapshotMulti(hex.EncodeToString(sum[:]), x), ShouldBeNil)
}

var nilAry []string

func TestParseROS2Field(t *testing.T) {
	SetDefaultFailureMode(FailureContinues)

	testFunc := func(description string, line string, ros2msg *ROS2Message) {
		testName := ros2msg.RosPackage + "." + ros2msg.RosMsgName + " " + line
		if description != "" {
			testName += " : " + description
		}
		Convey(testName, func() {
			m, err := ParseROS2MessageRow(line, ros2msg)
			So(err, ShouldBeNil)
			matchSnapshot(line, m)
		})
	}

	Convey("Parse ROS2 Fields", t, func() {
		testFunc("",
			"unique_identifier_msgs/UUID goal_id",
			ROS2MessageNew("action_msgs", "GoalInfo"),
		)
		testFunc("",
			"string[] full_node_names",
			ROS2MessageNew("composition_interfaces", "ListNodes_Response"),
		)
		testFunc("",
			"float64[3] float64_values_default [3.1415, 0.0, -3.1415]",
			ROS2MessageNew("test_msgs", "Arrays_Response"),
		)
		testFunc("",
			"BasicTypes[3] basic_types_values",
			ROS2MessageNew("test_msgs", "Arrays_Response"),
		)
		testFunc("",
			"bool[3] bool_values_default [false, true, false]",
			ROS2MessageNew("test_msgs", "Arrays_Response"),
		)
		testFunc("",
			`string[3] string_values_default ["", "max value", "min value"]`,
			ROS2MessageNew("test_msgs", "Arrays_Response"),
		)
		testFunc("Fields with comments containing '=' do not get identified as Constants",
			`float32 v_ref                      # ADC channel voltage reference, use to calculate LSB voltage(lsb=scale/resolution)`,
			ROS2MessageNew("px4_msgs", "AdcReport"),
		)
		testFunc("Array size int is big enough to store the C array size.",
			`uint8[512] junk`,
			ROS2MessageNew("px4_msgs", "OrbTestLarge"),
		)
		testFunc(`C struct has reserved Go keyword .test. Accessed with ._test instead.`,
			`uint8 type`,
			ROS2MessageNew("sensor_msgs", "JoyFeedback"),
		)
		testFunc(`Bounded sequence.`,
			`BasicTypes[<=3] basic_types_values`,
			ROS2MessageNew("test_msgs", "BoundedSequences"),
		)
		testFunc(`Bounded sequence with defaults.`,
			`int8[<=3] int8_values_default [0, 127, -128]`,
			ROS2MessageNew("test_msgs", "BoundedSequences"),
		)
		testFunc(`Bounded string with defaults.`,
			`string<=22 bounded_string_value "this is yet another"`,
			ROS2MessageNew("test_msgs", "Strings"),
		)
		testFunc(`Bounded string.`,
			`string<=22 bounded_string_value`,
			ROS2MessageNew("test_msgs", "Strings"),
		)
	})

	Convey("Parse ROS2 Constants", t, func() {
		testFunc("",
			"uint8 NUM_ACTUATOR_CONTROLS = 8",
			ROS2MessageNew("px4_msgs", "ActuatorControls0"),
		)
		testFunc("",
			"uint8 TYPE_LED    = 0",
			ROS2MessageNew("sensor_msgs", "JoyFeedback"),
		)
		testFunc("",
			"uint8 BATTERY_WARNING_CRITICAL = 2        # critical voltage, return / abort immediately",
			ROS2MessageNew("px4_msgs", "BatteryStatus"),
		)
		testFunc("",
			"byte BYTE_CONST=50",
			ROS2MessageNew("test_msgs", "Constants"),
		)
	})

	/*
		ROS2 snake-case-camel-case -conversions use different rules than
		"github.com/stoewer/go-strcase"
		or
		"github.com/iancoleman/strcase"
	*/
	Convey("Case transformations", t, func() {
		So(camelToSnake("ColorRGBA"), ShouldEqual, "color_rgba")
		So(camelToSnake("MultiDOFJointTrajectoryPoint"), ShouldEqual, "multi_dof_joint_trajectory_point")
		So(camelToSnake("TFMessage"), ShouldEqual, "tf_message")
		So(camelToSnake("WStrings"), ShouldEqual, "w_strings")
		So(camelToSnake("Float32MultiArray"), ShouldEqual, "float32_multi_array")
		So(camelToSnake("PointCloud2"), ShouldEqual, "point_cloud2")
		So(camelToSnake("GoalID"), ShouldEqual, "goal_id")
		So(camelToSnake("WString"), ShouldEqual, "w_string")
		So(camelToSnake("TF2Error"), ShouldEqual, "tf2_error")
	})

	Convey("Defaults string parser", t, func() {
		So(splitMsgDefaultArrayValues("int", ``), ShouldResemble, nilAry)
		So(splitMsgDefaultArrayValues("int", `[]`), ShouldResemble, nilAry)
		So(splitMsgDefaultArrayValues("int", `[1,2,3]`), ShouldResemble, []string{`1`, `2`, `3`})
		So(splitMsgDefaultArrayValues("string", `["", "this is a", "test msg"]`), ShouldResemble, []string{`""`, `"this is a"`, `"test msg"`})
		So(splitMsgDefaultArrayValues("string", `[1  ,  2 ,   "3"]`), ShouldResemble, []string{`"1  "`, `"2 "`, `"3"`})
		So(defaultValueSanitizer("string", `"Hello world!"`), ShouldEqual, `"Hello world!"`)
		So(defaultValueSanitizer("string", `"Hello\"world!"`), ShouldEqual, `"Hello\"world!"`)
	})

	Convey("defaultCode() generator", t, func() {
		So(defaultCode(&ROS2Field{
			TypeArray:    "[3]",
			ArraySize:    3,
			DefaultValue: "",
			PkgName:      "StringValues",
			PkgIsLocal:   false,
			RosType:      "string",
			CType:        "String",
			GoType:       "String",
			RosName:      "string_values",
			GoName:       "StringValues",
			CName:        "string_values",
			Comment:      "",
		}), ShouldEqual, `t.StringValues[0].SetDefaults(nil)
	t.StringValues[1].SetDefaults(nil)
	t.StringValues[2].SetDefaults(nil)
	`)
	})
}

func TestCErrorTypeParser(t *testing.T) {
	SetDefaultFailureMode(FailureContinues)
	Convey("", t, func() {
		et, err := ParseROS2ErrorType("/// Success return code.")
		So(et, ShouldBeNil)
		So(err, ShouldBeNil)
		So(ros2errorTypesCommentsBuffer.String(), ShouldEqual, "Success return code.")

		et, err = ParseROS2ErrorType("#define RCL_RET_OK RMW_RET_OK")
		So(et, ShouldResemble, &ROS2ErrorType{
			Name:      "RCL_RET_OK",
			Rcl_ret_t: "",
			Reference: "RMW_RET_OK",
			Comment:   "Success return code.",
		})
		So(err, ShouldBeNil)
		So(ros2errorTypesCommentsBuffer.Len(), ShouldEqual, 0)

		et, err = ParseROS2ErrorType("/// This comment is flushed because it is not part of a continuous stream.")
		So(et, ShouldBeNil)
		So(err, ShouldBeNil)
		So(ros2errorTypesCommentsBuffer.Len(), ShouldBeGreaterThan, 0)

		et, err = ParseROS2ErrorType("")
		So(et, ShouldBeNil)
		So(err, ShouldBeNil)
		So(ros2errorTypesCommentsBuffer.String(), ShouldEqual, "")

		et, err = ParseROS2ErrorType("#define RCL_RET_NOT_INIT 101")
		So(et, ShouldResemble, &ROS2ErrorType{
			Name:      "RCL_RET_NOT_INIT",
			Rcl_ret_t: "101",
			Reference: "",
			Comment:   "",
		})
		So(err, ShouldBeNil)
		So(ros2errorTypesCommentsBuffer.Len(), ShouldEqual, 0)
	})
}

/*
func TestSerDesSimple(t *testing.T) {

	Convey("example_interfaces/Float32MultiArray.msg", t, func() {
		ob := test_msgs.Arrays{}
	})
}*/

func TestBlacklist(t *testing.T) {
	SetDefaultFailureMode(FailureContinues)

	Convey("Blacklist", t, func() {
		skip, blacklistEntry := blacklisted("/opt/ros/foxy/this-is-a-test-blacklist-entry-do-not-remove-used-for-internal-testing/msgs/Lol.msg")
		So(skip, ShouldBeTrue)
		So(blacklistEntry, ShouldEqual, "this-is-a-test-blacklist-entry-do-not-remove-used-for-internal-testing")
	})
}
