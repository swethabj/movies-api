module github.com/swethabj/movies-api

go 1.23.3

// go get github.com/gofiber/fiber/v2@v2.52.9     // Fiber (v2.x)
// go get github.com/sirupsen/logrus@v1.9.3       // Logrus
// go get gorm.io/driver/mysql@v1.6.0             // GORM MySQL Driver
// go get gorm.io/gorm@v1.30.3                    // GORM ORM
// go get github.com/joho/godotenv@v1.5.1         // Godotenv
// go mod tidy ## Dont run else unused depenpencies will go
// go get github.com/swaggo/swag/cmd/swag  // documentation
// go get github.com/gofiber/swagger.      // documentation

require (
	github.com/gofiber/fiber/v2 v2.52.9
	github.com/gofiber/swagger v1.1.1
	github.com/joho/godotenv v1.5.1
	github.com/sirupsen/logrus v1.9.3
	github.com/swaggo/swag v1.16.6
	gorm.io/driver/mysql v1.6.0
	gorm.io/gorm v1.30.3
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/swaggo/files/v2 v2.0.2 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
