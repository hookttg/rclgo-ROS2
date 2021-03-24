package gogen

import (
	"container/list"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func GenerateMsgTypeNameAliasDispatcher(mds *list.List, destPathPkgRoot string) error {
	mdsArray := ROS2MessageListToArray(mds)

	destFilePath := filepath.Join(destPathPkgRoot, "..", "ros2_type_dispatcher", "ros2_msgs_static_imports.go")

	_, err := os.Stat(destFilePath)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("'%s' is missing? It should exist relative the the given destination path '%s'", destFilePath, destPathPkgRoot)
		err = os.MkdirAll(filepath.Dir(destFilePath), os.ModePerm)
		if err != nil {
			return err
		}
	}
	destFile, err := os.Create(destFilePath)
	if err != nil {
		return err
	}

	deduplicatedPackagesMap := make(map[string]int)
	for _, v := range mdsArray {
		deduplicatedPackagesMap[v.RosPackage] += 1
	}

	ros2MsgsStaticImportsDispatcherTemplate.Funcs(template.FuncMap{
		"lc": strings.ToLower,
	})
	return ros2MsgsStaticImportsDispatcherTemplate.Execute(destFile, map[string]interface{}{
		"Mds":             mdsArray,
		"Ros2PackagesMap": deduplicatedPackagesMap,
	})
}

func GenerateGolangTypeFromROS2MessagePath(sourcePath string, destPathPkgRoot string) (*ROS2Message, error) {
	md := &ROS2Message{}

	err := ParseMessageMetadataFromPath(sourcePath, md)
	if err != nil {
		return md, err
	}

	content, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		return md, err
	}

	err = ParseROS2Message(md, string(content))
	if err != nil {
		return md, err
	}

	destFile, err := CreateTargetGolangTypeFile(destPathPkgRoot, md)
	if err != nil {
		return md, err
	}

	err = ros2MsgToGolangTypeTemplate.Execute(destFile, md)
	if err != nil {
		return md, err
	}
	return md, nil
}

func ParseMessageMetadataFromPath(p string, md *ROS2Message) error {
	var dirs []string
	var ros2msgName string
	var ros2pkgName string

	dirs = strings.Split(p, "/")
	ros2msgName = strings.TrimSuffix(filepath.Base(p), ".msg")

	if len(dirs) >= 2 {
		ros2pkgName = dirs[len(dirs)-3]
	} else {
		return fmt.Errorf("Path '%s' cannot be parsed for ROS2 package name!", p)
	}

	md.RosMsgName = ros2msgName
	md.RosPackage = ros2pkgName
	md.Url = p

	return nil
}

func CreateTargetGolangTypeFile(destPathPkgRoot string, md *ROS2Message) (*os.File, error) {
	destFilePath := filepath.Join(destPathPkgRoot, md.RosPackage, "msg", md.RosMsgName+".go")
	destFileDir := filepath.Join(destPathPkgRoot, md.RosPackage, "msg")
	_, err := os.Stat(destFileDir)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Creating directory '%s'", destFileDir)
		err = os.MkdirAll(destFileDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	destFile, err := os.Create(destFilePath)
	if err != nil {
		return destFile, err
	}
	return destFile, err
}

func ROS2MessageListToArray(l *list.List) []*ROS2Message {
	mdsArray := make([]*ROS2Message, l.Len())
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		md, ok := e.Value.(*ROS2Message)
		if !ok {
			panic(fmt.Sprintf("ROS2MessageListToArray():> One of the ROS2Messages at index '%d' value '%+v' is not a ROS2Message!", i, e.Value))
		}
		mdsArray[i] = md
		i++
	}
	return mdsArray
}
