package libtest

import (
    "runtime"
    "path/filepath"
    "os"

    "go-rest-api/app"
)

func Config() (app.ConfigType, error){
    _, filename, _, _ := runtime.Caller(0) //nolint:golangcilint,dogsled

    configFileFullPath := filepath.Join(filepath.Dir(filename), "..", "..", "secrets", "config.yml")

    var empty app.ConfigType
    err := os.Setenv(app.ENV_VAR, app.ENV_TEST)
    if err != nil {
        return empty, err
    }

    cfg, err := app.NewConfig(configFileFullPath)
    if err != nil {
        return empty, err
    }

    return cfg, nil
}
