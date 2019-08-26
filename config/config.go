package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type JsonWorkflowDefinition struct {
	Package Package
	Tasks []JsonTaskRaw `json:"tasks"`
}
type JsonTaskRaw struct {
	Type TaskType
	Name string
	Package *Package
	Data json.RawMessage
}

type TaskType string

const (
	Simple TaskType = "simple"
)

type WorkflowDefinition struct{
	Package Package
	Types []ParameterType
	Parameters []Parameter
	Packages []Package
	TaskDefinitions []TaskDefinition
}

type TaskDefinition struct {
	Type TaskType
	Name string
	Package *Package
	NextTaskName *string
	Task interface{}
}
type Parameter struct {
	ParameterType ParameterType `json:"parameter_type"`
	Name string
}

type Package struct {
	Name string
	Alias string
}
type ParameterType struct {
	Type string
	Package *Package

}
type SimpleTask struct {
	Arguments []Parameter
	Returns []Parameter
}

func NewConfig(cfgFile string) WorkflowDefinition {
	file, _ := ioutil.ReadFile(cfgFile)

	data := JsonWorkflowDefinition{}

	_ = json.Unmarshal([]byte(file), &data)

	var definitions []TaskDefinition

	//todo check that there's no "Done" task! this name is taken
	dataTypesMap := map[string]ParameterType{}
	var dataTypes []ParameterType

	dataParameterMap := map[string]Parameter{}
	var dataParameters []Parameter

	packagesMap := map[string]Package{}
	var packages []Package

	for i, taskData := range data.Tasks {
		var nextTaskName *string = nil
		if i != len(data.Tasks) - 1 {
			nextTaskName = &data.Tasks[i+1].Name
		}
		if taskData.Type == Simple {
			task := SimpleTask{}
			_ = json.Unmarshal(taskData.Data, &task)
			definitions = append(definitions, TaskDefinition{
				taskData.Type, taskData.Name, taskData.Package, nextTaskName, task})


			for _, arg := range task.Arguments {
				dataTypesMap[arg.ParameterType.Type] = arg.ParameterType
				dataParameterMap[arg.Name] = arg
				if arg.ParameterType.Package != nil {
					packagesMap[arg.ParameterType.Package.Alias] = *arg.ParameterType.Package
				}
			}

			for _, arg := range task.Returns {
				dataTypesMap[arg.ParameterType.Type] = arg.ParameterType
				dataParameterMap[arg.Name] = arg
				if arg.ParameterType.Package != nil {
					packagesMap[arg.ParameterType.Package.Alias] = *arg.ParameterType.Package
				}
			}
		} else {
			panic(errors.New("unknown task type"))
		}
	}


	for _, pack := range packagesMap {
		packages = append(packages, pack)
	}

	for _, parameter := range dataParameterMap {
		dataParameters = append(dataParameters, parameter)
	}

	for _, dType := range dataTypesMap {
		dataTypes = append(dataTypes, dType)
	}

	workflowDer := WorkflowDefinition{
		Package: data.Package,
		Types: dataTypes,
		Parameters: dataParameters,
		TaskDefinitions: definitions,
		Packages: packages,
	}

	return workflowDer
}
