module github.com/UpskillBD/BE-TrainerDash

replace github.com/UpskillBD/BE-TrainerDash/cmd => ./cmd

replace github.com/UpskillBD/BE-TrainerDash/api => ./api

replace github.com/UpskillBD/BE-TrainerDash/config => ./config

replace github.com/UpskillBD/BE-TrainerDash/db => ./db

replace github.com/UpskillBD/BE-TrainerDash/models => ./models

go 1.15

require (
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-gonic/gin v1.6.3
	github.com/go-pg/pg v8.0.7+incompatible
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.1
	mellium.im/sasl v0.2.1 // indirect
)
