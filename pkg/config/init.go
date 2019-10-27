package config

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

type Config interface {
	GetEnableDebugLogs() bool
	GetDSN() string
	GetNausysAddr() string
	GetNausysUser() string
	GetNausysPass() string
}

func (d *Desc) GetEnableDebugLogs() bool {
	return d.EnableDebugLogs
}

func (d *Desc) GetDSN() string {
	return d.DSN
}

func (d *Desc) GetNausysAddr() string {
	return d.NausysAddr
}

func (d *Desc) GetNausysUser() string {
	return d.NausysUser
}

func (d *Desc) GetNausysPass() string {
	return d.NausysPass
}

func Init(configPath string) (Config, error) {
	f, err := os.Open(configPath)
	if err != nil {
		return nil, errors.Wrap(err, "open config file")
	}
	defer f.Close()
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrapf(err, "read config file content")
	}
	c := &Desc{}
	if _, err := toml.Decode(string(content), c); err != nil {
		return nil, errors.Wrapf(err, "decode config file content")
	}
	if err := validate(c); err != nil {
		return nil, errors.Wrapf(err, "validate config")
	}
	return c, nil
}

func validate(d *Desc) error {
	if d.DSN == "" {
		return errors.Errorf("%q missing", "dsn")
	}
	if d.NausysAddr == "" {
		return errors.Errorf("%q missing", "nausys_addr")
	}
	if d.NausysUser == "" {
		return errors.Errorf("%q missing", "nausys_user")
	}
	if d.NausysPass == "" {
		return errors.Errorf("%q missing", "nausys_pass")
	}
	return nil
}
