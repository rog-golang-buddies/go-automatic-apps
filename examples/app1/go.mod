module app1

replace github.com/rog-golang-buddies/go-automatic-apps => ../../

go 1.19

require (
	github.com/rog-golang-buddies/go-automatic-apps v1.1.0
	gorm.io/driver/sqlite v1.3.6
	gorm.io/gorm v1.23.8
)

require (
	github.com/go-chi/chi/v5 v5.0.7 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.14 // indirect
	github.com/rs/cors v1.8.2 // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect
)
