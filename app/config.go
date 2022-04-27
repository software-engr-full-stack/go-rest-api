package app

import (
    "os"
    "strings"
    "fmt"
    "path/filepath"

    "github.com/pkg/errors"
    yaml "gopkg.in/yaml.v2"
)

type ConfigDatabaseType struct {
    Host         string `yaml:"host"`
    Name         string `yaml:"name"`
    Adapter      string `yaml:"adapter"`
    Encoding     string `yaml:"encoding"`
    MaxOpenConns uint   `yaml:"max_open_conns"`
    SearchPath   string `yaml:"search_path"`
    Pool         uint   `yaml:"pool"`
    User         string `yaml:"user"`
    Password     string `yaml:"password"`
    Port         int    `yaml:"port"`
}

type unmarshalType struct {
    Database     map[string]ConfigDatabaseType `yaml:"database"`
}

type ConfigType struct {
    IsDevelopment bool
    IsTest        bool
    IsProduction  bool
    Env           string

    AppDir        string

    Database      ConfigDatabaseType
}

const (
    ENV_VAR         = "GO_ENV"
    ENV_DEVELOPMENT = "development"
    ENV_TEST        = "test"
    ENV_PRODUCTION  = "production"
)

func NewConfig(configFileName string) (ConfigType, error) {
    var empty ConfigType

    var cfg ConfigType
    env := strings.TrimSpace(os.Getenv(ENV_VAR))
    switch env {
    case ENV_DEVELOPMENT, "":
        cfg.IsDevelopment = true
        cfg.Env = ENV_DEVELOPMENT
    case ENV_TEST:
        cfg.IsTest = true
        cfg.Env = env
    case ENV_PRODUCTION:
        cfg.IsProduction = true
        cfg.Env = env
    default:
        return empty, errors.WithStack(fmt.Errorf("invalid env %#v", env))
    }

    fileData, err := os.ReadFile(configFileName)
    if err != nil {
        return empty, errors.WithStack(err)
    }

    var unmarshal unmarshalType
    err = yaml.Unmarshal(fileData, &unmarshal)
    if err != nil {
        return empty, errors.WithStack(err)
    }

    if db, ok := unmarshal.Database[cfg.Env]; !ok {
        return empty, errors.WithStack(fmt.Errorf("no matching database for env %#v", cfg.Env))
    } else {
        cfg.Database = db
    }

    cfg.AppDir = filepath.Join(filepath.Dir(configFileName), "..")

    return cfg, nil
}
