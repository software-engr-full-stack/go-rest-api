package main

import (
    "go-rest-api/app"
    "go-rest-api/sqlite"
)

func dbSetup(cfg app.ConfigType) (defErr error) {
    db := sqlite.NewDB(cfg)
    if err := db.Open(); err != nil {
        return err
    }
    defer func() {
        defErr = db.Close()
    }()

    return nil
}
