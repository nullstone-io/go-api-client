package find

import (
	"errors"
	"fmt"
)

func IsMissingResource(err error) bool {
	var mre MissingResourceError
	if errors.As(err, &mre) {
		return mre.IsMissing()
	}
	return false
}

type MissingResourceError interface {
	IsMissing() bool
}

type AppDoesNotExistError struct{ AppName string }

func (AppDoesNotExistError) IsMissing() bool { return true }
func (e AppDoesNotExistError) Error() string {
	return fmt.Sprintf("application %q does not exist", e.AppName)
}

type StackDoesNotExistError struct{ StackName string }

func (StackDoesNotExistError) IsMissing() bool { return true }
func (e StackDoesNotExistError) Error() string {
	return fmt.Sprintf("stack %q does not exist", e.StackName)
}

type StackIdDoesNotExistError struct{ StackId int64 }

func (StackIdDoesNotExistError) IsMissing() bool { return true }
func (e StackIdDoesNotExistError) Error() string {
	return fmt.Sprintf("stack %d does not exist", e.StackId)
}

type EnvDoesNotExistError struct {
	StackName string
	EnvName   string
}

func (EnvDoesNotExistError) IsMissing() bool { return true }
func (e EnvDoesNotExistError) Error() string {
	return fmt.Sprintf("environment %s/%s does not exist", e.StackName, e.EnvName)
}

type EnvIdDoesNotExistError struct {
	StackName string
	EnvId     int64
}

func (EnvIdDoesNotExistError) IsMissing() bool { return true }
func (e EnvIdDoesNotExistError) Error() string {
	return fmt.Sprintf("environment %s/%d does not exist", e.StackName, e.EnvId)
}

type BlockDoesNotExistError struct {
	StackName string
	BlockName string
}

func (BlockDoesNotExistError) IsMissing() bool { return true }
func (e BlockDoesNotExistError) Error() string {
	return fmt.Sprintf("block %s/%s does not exist", e.StackName, e.BlockName)
}

type BlockIdDoesNotExistError struct {
	StackName string
	BlockId   int64
}

func (BlockIdDoesNotExistError) IsMissing() bool { return true }
func (e BlockIdDoesNotExistError) Error() string {
	return fmt.Sprintf("block %s/%d does not exist", e.StackName, e.BlockId)
}
