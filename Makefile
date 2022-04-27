db-setup:
	go run cmd/appctl/*.go db-setup --config ./secrets/config.yml

db-drop:
	go run cmd/appctl/*.go db-drop --config ./secrets/config.yml

db-reset: db-drop db-setup

.PHONY: db-setup db-drop db-reset
