// Code generated by resources. DO NOT EDIT.
package resources

import (
	"errors"

	"go.mondoo.com/cnquery/providers/plugin"
	"go.mondoo.com/cnquery/providers/proto"
	"go.mondoo.com/cnquery/types"
)

var newResource = map[string]func(runtime *plugin.Runtime, args map[string]interface{}) (plugin.Resource, error){
	"command": NewCommand,
}

// CreateResource is used by the runtime of this plugin
func CreateResource(runtime *plugin.Runtime, name string, args map[string]interface{}) (plugin.Resource, error) {
	f, ok := newResource[name]
	if !ok {
		return nil, errors.New("cannot find resource " + name + " in os provider")
	}

	return f(runtime, args)
}

var getDataFields = map[string]func(r plugin.Resource) *proto.DataRes{
	"command.command": func(r plugin.Resource) *proto.DataRes {
		return (r.(*mqlCommand).GetCommand()).ToDataRes(types.String)
	},
	"command.stdout": func(r plugin.Resource) *proto.DataRes {
		return (r.(*mqlCommand).GetStdout()).ToDataRes(types.String)
	},
	"command.stderr": func(r plugin.Resource) *proto.DataRes {
		return (r.(*mqlCommand).GetStderr()).ToDataRes(types.String)
	},
	"command.exitcode": func(r plugin.Resource) *proto.DataRes {
		return (r.(*mqlCommand).GetExitcode()).ToDataRes(types.Int)
	},
}

func GetData(resource plugin.Resource, field string, args map[string]interface{}) *proto.DataRes {
	f, ok := getDataFields[resource.MqlName()+"."+field]
	if !ok {
		return &proto.DataRes{Error: "cannot find '" + field + "' in resource '" + resource.MqlName() + "'"}
	}

	return f(resource)
}

var setDataFields = map[string]func(r plugin.Resource, v interface{}) bool {
	"command.command": func(r plugin.Resource, v interface{}) bool {
		var ok bool
		r.(*mqlCommand).Command, ok = plugin.RawToTValue[string](v)
		return ok
	},
	"command.stdout": func(r plugin.Resource, v interface{}) bool {
		var ok bool
		r.(*mqlCommand).Stdout, ok = plugin.RawToTValue[string](v)
		return ok
	},
	"command.stderr": func(r plugin.Resource, v interface{}) bool {
		var ok bool
		r.(*mqlCommand).Stderr, ok = plugin.RawToTValue[string](v)
		return ok
	},
	"command.exitcode": func(r plugin.Resource, v interface{}) bool {
		var ok bool
		r.(*mqlCommand).Exitcode, ok = plugin.RawToTValue[int64](v)
		return ok
	},
}

func SetData(resource plugin.Resource, field string, val interface{}) error {
	f, ok := setDataFields[resource.MqlName() + "." + field]
	if !ok {
		return errors.New("cannot set '"+field+"' in resource '"+resource.MqlName()+"', field not found")
	}

	if ok := f(resource, val); !ok {
		return errors.New("cannot set '"+field+"' in resource '"+resource.MqlName()+"', type does not match")
	}
	return nil
}

// mqlCommand for the command resource
type mqlCommand struct {
	MqlRuntime *plugin.Runtime
	_id string
	mqlCommandInternal

	Command plugin.TValue[string]
	Stdout plugin.TValue[string]
	Stderr plugin.TValue[string]
	Exitcode plugin.TValue[int64]
}

// NewCommand creates a new instance of this resource
func NewCommand(runtime *plugin.Runtime, args map[string]interface{}) (plugin.Resource, error) {
	res := &mqlCommand{
		MqlRuntime: runtime,
	}

	var err error
	// to override args, implement: init(args map[string]interface{}) (map[string]interface{}, *mqlCommand, error)

	for k, v := range args {
		if err = SetData(res, k, v); err != nil {
			return res, err
		}
	}

	res._id, err = res.id()
	return res, err
}

func (c *mqlCommand) MqlName() string {
	return "command"
}

func (c *mqlCommand) MqlID() string {
	return c._id
}

func (c *mqlCommand) GetCommand() *plugin.TValue[string] {
	return &c.Command
}

func (c *mqlCommand) GetStdout() *plugin.TValue[string] {
	return plugin.GetOrCompute[string](&c.Stdout, func() (string, error) {
		vargCommand := c.GetCommand()
		if vargCommand.Error != nil {
			return "", vargCommand.Error
		}
		return c.stdout(vargCommand.Data)
	})
}

func (c *mqlCommand) GetStderr() *plugin.TValue[string] {
	return plugin.GetOrCompute[string](&c.Stderr, func() (string, error) {
		vargCommand := c.GetCommand()
		if vargCommand.Error != nil {
			return "", vargCommand.Error
		}
		return c.stderr(vargCommand.Data)
	})
}

func (c *mqlCommand) GetExitcode() *plugin.TValue[int64] {
	return plugin.GetOrCompute[int64](&c.Exitcode, func() (int64, error) {
		vargCommand := c.GetCommand()
		if vargCommand.Error != nil {
			return 0, vargCommand.Error
		}
		return c.exitcode(vargCommand.Data)
	})
}
