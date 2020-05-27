package transfer

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Driver for database
type Driver struct {
	Driver   string `mapstructure:"driver"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	SSLMode  string `mapstructure:"ssl_mode"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// Query database query
type Query struct {
	Q    interface{} `mapstructure:"q"`
	Page uint        `mapstructure:"page"`
	Size uint        `mapstructure:"size"`
}

// UnmarshalQuery implements the json.Marshaler interface.
func (q Query) UnmarshalQuery(v interface{}) error {

	byteData, err := json.Marshal(q)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(byteData), &v)
}

// DatabaseOptions for transfer
type DatabaseOptions struct {
	Driver    Driver
	TableName string
	Mapping   Mapping
}

// GenerateSourceTransfer return source transfer
func GenerateSourceTransfer(args *DatabaseOptions) (source Source, err error) {

	switch args.Driver.Driver {
	case "mongodb":
		source, err = NewMongoDB(args)
	case "mysql":
		source, err = NewMySQL(args)
	}

	return
}

// GenerateTargetTransfer return target transfer
func GenerateTargetTransfer(args *DatabaseOptions) (target Target, err error) {

	switch args.Driver.Driver {
	case "mongodb":
		target, err = NewMongoDB(args)
	case "mysql":
		target, err = NewMySQL(args)
	}

	return
}