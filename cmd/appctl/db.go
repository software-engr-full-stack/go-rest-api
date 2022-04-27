package main

import (
    "go-rest-api/app"
    "go-rest-api/sqlite"
)

func dbSQLiteSetup(cfg app.ConfigType) (defErr error) {
    db := sqlite.NewDB(cfg)
    if err := db.Open(); err != nil {
        return err
    }
    defer func() {
        defErr = db.Close()
    }()

    return nil
}

func dbSQLiteDrop(cfg app.ConfigType) (defErr error) {
    db := sqlite.NewDB(cfg)
    if err := db.Drop(); err != nil {
        return err
    }
    defer func() {
        defErr = db.Close()
    }()

    return nil
}
