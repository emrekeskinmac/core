package service

// Task describes a service task.
type Task struct {
	// Key is the key of task.
	Key string `hash:"name:1"`

	// Name is the name of task.
	Name string `hash:"name:2"`

	// Description is the description of task.
	Description string `hash:"name:3"`

	// Inputs are the definition of the execution inputs of task.
	Inputs []*Parameter `hash:"name:4"`

	// Outputs are the definition of the execution results of task.
	Outputs []*Output `hash:"name:5"`

	// serviceName is the task's service's name.
	serviceName string `hash:"-"`
}

// Output describes task output.
type Output struct {
	// Key is the key of output.
	Key string `hash:"name:1"`

	// Name is the name of task output.
	Name string `hash:"name:2"`

	// Description is the description of task output.
	Description string `hash:"name:3"`

	// Data holds the output parameters of a task output.
	Data []*Parameter `hash:"name:4"`

	// taskKey is the output's task's key.
	taskKey string `hash:"-"`

	// serviceName is the output's service's name.
	serviceName string `hash:"-"`
}

// GetTask returns task taskKey of service.
func (s *Service) GetTask(taskKey string) (*Task, error) {
	for _, task := range s.Tasks {
		if task.Key == taskKey {
			task.serviceName = s.Name
			return task, nil
		}
	}
	return nil, &TaskNotFoundError{
		TaskKey:     taskKey,
		ServiceName: s.Name,
	}
}

// GetInputParameter returns input inputKey parameter of task.
func (t *Task) GetInputParameter(inputKey string) (*Parameter, error) {
	for _, input := range t.Inputs {
		if input.Key == inputKey {
			return input, nil
		}
	}
	return nil, &TaskInputNotFoundError{
		TaskKey:      t.Key,
		TaskInputKey: inputKey,
		ServiceName:  t.serviceName,
	}
}

// ValidateInputs produces warnings for task inputs that doesn't satisfy their parameter schemas.
func (t *Task) ValidateInputs(taskInputs map[string]interface{}) []*ParameterWarning {
	return validateParametersSchema(t.Inputs, taskInputs)
}

// RequireInputs requires task inputs to be matched with parameter schemas.
func (t *Task) RequireInputs(taskInputs map[string]interface{}) error {
	warnings := t.ValidateInputs(taskInputs)
	if len(warnings) > 0 {
		return &InvalidTaskInputError{
			TaskKey:     t.Key,
			ServiceName: t.serviceName,
			Warnings:    warnings,
		}
	}
	return nil
}

// GetOutput returns output outputKey of task.
func (t *Task) GetOutput(outputKey string) (*Output, error) {
	for _, output := range t.Outputs {
		if output.Key == outputKey {
			output.taskKey = t.Key
			output.serviceName = t.serviceName
			return output, nil
		}
	}
	return nil, &TaskOutputNotFoundError{
		TaskKey:       t.Key,
		TaskOutputKey: outputKey,
		ServiceName:   t.serviceName,
	}
}

// ValidateData produces warnings for task outputs that doesn't satisfy their parameter schemas.
func (o *Output) ValidateData(outputData map[string]interface{}) []*ParameterWarning {
	return validateParametersSchema(o.Data, outputData)
}

// RequireData requires task outputs to be matched with parameter schemas.
func (o *Output) RequireData(outputData map[string]interface{}) error {
	warnings := o.ValidateData(outputData)
	if len(warnings) > 0 {
		return &InvalidTaskOutputError{
			TaskKey:       o.taskKey,
			TaskOutputKey: o.Key,
			ServiceName:   o.serviceName,
			Warnings:      warnings,
		}
	}
	return nil
}
