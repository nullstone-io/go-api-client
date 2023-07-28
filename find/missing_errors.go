package find

import "fmt"

type AppDoesNotExistError struct{ AppName string }

func (e AppDoesNotExistError) Error() string {
	return fmt.Sprintf("application %q does not exist", e.AppName)
}

type StackDoesNotExistError struct{ StackName string }

func (e StackDoesNotExistError) Error() string {
	return fmt.Sprintf("stack %q does not exist", e.StackName)
}

type StackIdDoesNotExistError struct{ StackId int64 }

func (e StackIdDoesNotExistError) Error() string {
	return fmt.Sprintf("stack %d does not exist", e.StackId)
}

type EnvDoesNotExistError struct {
	StackName string
	EnvName   string
}

func (e EnvDoesNotExistError) Error() string {
	return fmt.Sprintf("environment %s/%s does not exist", e.StackName, e.EnvName)
}
