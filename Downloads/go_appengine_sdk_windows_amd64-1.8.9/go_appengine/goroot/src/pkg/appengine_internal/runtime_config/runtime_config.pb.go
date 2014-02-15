// Code generated by protoc-gen-go.
// DO NOT EDIT!

/*
Package appengine_tools_devappserver2 is a generated protocol buffer package.

It is generated from these files:
	appengine_internal/runtime_config

It has these top-level messages:
	Config
	PhpConfig
	PythonConfig
	CloudSQL
	Library
	Environ
*/
package appengine_tools_devappserver2

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Config struct {
	AppId            []byte        `protobuf:"bytes,1,req,name=app_id" json:"app_id,omitempty"`
	VersionId        []byte        `protobuf:"bytes,2,req,name=version_id" json:"version_id,omitempty"`
	ApplicationRoot  []byte        `protobuf:"bytes,3,req,name=application_root" json:"application_root,omitempty"`
	Threadsafe       *bool         `protobuf:"varint,4,opt,name=threadsafe,def=0" json:"threadsafe,omitempty"`
	ApiHost          *string       `protobuf:"bytes,17,opt,name=api_host,def=localhost" json:"api_host,omitempty"`
	ApiPort          *int32        `protobuf:"varint,5,req,name=api_port" json:"api_port,omitempty"`
	Libraries        []*Library    `protobuf:"bytes,6,rep,name=libraries" json:"libraries,omitempty"`
	SkipFiles        *string       `protobuf:"bytes,7,opt,name=skip_files,def=^$" json:"skip_files,omitempty"`
	StaticFiles      *string       `protobuf:"bytes,8,opt,name=static_files,def=^$" json:"static_files,omitempty"`
	PythonConfig     *PythonConfig `protobuf:"bytes,14,opt,name=python_config" json:"python_config,omitempty"`
	PhpConfig        *PhpConfig    `protobuf:"bytes,9,opt,name=php_config" json:"php_config,omitempty"`
	Environ          []*Environ    `protobuf:"bytes,10,rep,name=environ" json:"environ,omitempty"`
	CloudSqlConfig   *CloudSQL     `protobuf:"bytes,11,opt,name=cloud_sql_config" json:"cloud_sql_config,omitempty"`
	Datacenter       *string       `protobuf:"bytes,12,req,name=datacenter" json:"datacenter,omitempty"`
	InstanceId       *string       `protobuf:"bytes,13,req,name=instance_id" json:"instance_id,omitempty"`
	StderrLogLevel   *int64        `protobuf:"varint,15,opt,name=stderr_log_level,def=1" json:"stderr_log_level,omitempty"`
	AuthDomain       *string       `protobuf:"bytes,16,req,name=auth_domain" json:"auth_domain,omitempty"`
	MaxInstances     *int32        `protobuf:"varint,18,opt,name=max_instances" json:"max_instances,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}

const Default_Config_Threadsafe bool = false
const Default_Config_ApiHost string = "localhost"
const Default_Config_SkipFiles string = "^$"
const Default_Config_StaticFiles string = "^$"
const Default_Config_StderrLogLevel int64 = 1

func (m *Config) GetAppId() []byte {
	if m != nil {
		return m.AppId
	}
	return nil
}

func (m *Config) GetVersionId() []byte {
	if m != nil {
		return m.VersionId
	}
	return nil
}

func (m *Config) GetApplicationRoot() []byte {
	if m != nil {
		return m.ApplicationRoot
	}
	return nil
}

func (m *Config) GetThreadsafe() bool {
	if m != nil && m.Threadsafe != nil {
		return *m.Threadsafe
	}
	return Default_Config_Threadsafe
}

func (m *Config) GetApiHost() string {
	if m != nil && m.ApiHost != nil {
		return *m.ApiHost
	}
	return Default_Config_ApiHost
}

func (m *Config) GetApiPort() int32 {
	if m != nil && m.ApiPort != nil {
		return *m.ApiPort
	}
	return 0
}

func (m *Config) GetLibraries() []*Library {
	if m != nil {
		return m.Libraries
	}
	return nil
}

func (m *Config) GetSkipFiles() string {
	if m != nil && m.SkipFiles != nil {
		return *m.SkipFiles
	}
	return Default_Config_SkipFiles
}

func (m *Config) GetStaticFiles() string {
	if m != nil && m.StaticFiles != nil {
		return *m.StaticFiles
	}
	return Default_Config_StaticFiles
}

func (m *Config) GetPythonConfig() *PythonConfig {
	if m != nil {
		return m.PythonConfig
	}
	return nil
}

func (m *Config) GetPhpConfig() *PhpConfig {
	if m != nil {
		return m.PhpConfig
	}
	return nil
}

func (m *Config) GetEnviron() []*Environ {
	if m != nil {
		return m.Environ
	}
	return nil
}

func (m *Config) GetCloudSqlConfig() *CloudSQL {
	if m != nil {
		return m.CloudSqlConfig
	}
	return nil
}

func (m *Config) GetDatacenter() string {
	if m != nil && m.Datacenter != nil {
		return *m.Datacenter
	}
	return ""
}

func (m *Config) GetInstanceId() string {
	if m != nil && m.InstanceId != nil {
		return *m.InstanceId
	}
	return ""
}

func (m *Config) GetStderrLogLevel() int64 {
	if m != nil && m.StderrLogLevel != nil {
		return *m.StderrLogLevel
	}
	return Default_Config_StderrLogLevel
}

func (m *Config) GetAuthDomain() string {
	if m != nil && m.AuthDomain != nil {
		return *m.AuthDomain
	}
	return ""
}

func (m *Config) GetMaxInstances() int32 {
	if m != nil && m.MaxInstances != nil {
		return *m.MaxInstances
	}
	return 0
}

type PhpConfig struct {
	PhpExecutablePath []byte `protobuf:"bytes,1,opt,name=php_executable_path" json:"php_executable_path,omitempty"`
	EnableDebugger    *bool  `protobuf:"varint,3,req,name=enable_debugger" json:"enable_debugger,omitempty"`
	XXX_unrecognized  []byte `json:"-"`
}

func (m *PhpConfig) Reset()         { *m = PhpConfig{} }
func (m *PhpConfig) String() string { return proto.CompactTextString(m) }
func (*PhpConfig) ProtoMessage()    {}

func (m *PhpConfig) GetPhpExecutablePath() []byte {
	if m != nil {
		return m.PhpExecutablePath
	}
	return nil
}

func (m *PhpConfig) GetEnableDebugger() bool {
	if m != nil && m.EnableDebugger != nil {
		return *m.EnableDebugger
	}
	return false
}

type PythonConfig struct {
	StartupScript    *string `protobuf:"bytes,1,opt,name=startup_script" json:"startup_script,omitempty"`
	StartupArgs      *string `protobuf:"bytes,2,opt,name=startup_args" json:"startup_args,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PythonConfig) Reset()         { *m = PythonConfig{} }
func (m *PythonConfig) String() string { return proto.CompactTextString(m) }
func (*PythonConfig) ProtoMessage()    {}

func (m *PythonConfig) GetStartupScript() string {
	if m != nil && m.StartupScript != nil {
		return *m.StartupScript
	}
	return ""
}

func (m *PythonConfig) GetStartupArgs() string {
	if m != nil && m.StartupArgs != nil {
		return *m.StartupArgs
	}
	return ""
}

type CloudSQL struct {
	MysqlHost        *string `protobuf:"bytes,1,req,name=mysql_host" json:"mysql_host,omitempty"`
	MysqlPort        *int32  `protobuf:"varint,2,req,name=mysql_port" json:"mysql_port,omitempty"`
	MysqlUser        *string `protobuf:"bytes,3,req,name=mysql_user" json:"mysql_user,omitempty"`
	MysqlPassword    *string `protobuf:"bytes,4,req,name=mysql_password" json:"mysql_password,omitempty"`
	MysqlSocket      *string `protobuf:"bytes,5,opt,name=mysql_socket" json:"mysql_socket,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CloudSQL) Reset()         { *m = CloudSQL{} }
func (m *CloudSQL) String() string { return proto.CompactTextString(m) }
func (*CloudSQL) ProtoMessage()    {}

func (m *CloudSQL) GetMysqlHost() string {
	if m != nil && m.MysqlHost != nil {
		return *m.MysqlHost
	}
	return ""
}

func (m *CloudSQL) GetMysqlPort() int32 {
	if m != nil && m.MysqlPort != nil {
		return *m.MysqlPort
	}
	return 0
}

func (m *CloudSQL) GetMysqlUser() string {
	if m != nil && m.MysqlUser != nil {
		return *m.MysqlUser
	}
	return ""
}

func (m *CloudSQL) GetMysqlPassword() string {
	if m != nil && m.MysqlPassword != nil {
		return *m.MysqlPassword
	}
	return ""
}

func (m *CloudSQL) GetMysqlSocket() string {
	if m != nil && m.MysqlSocket != nil {
		return *m.MysqlSocket
	}
	return ""
}

type Library struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Version          *string `protobuf:"bytes,2,req,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Library) Reset()         { *m = Library{} }
func (m *Library) String() string { return proto.CompactTextString(m) }
func (*Library) ProtoMessage()    {}

func (m *Library) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Library) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

type Environ struct {
	Key              []byte `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            []byte `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Environ) Reset()         { *m = Environ{} }
func (m *Environ) String() string { return proto.CompactTextString(m) }
func (*Environ) ProtoMessage()    {}

func (m *Environ) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Environ) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
}
