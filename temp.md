# main.go

# models/users.go
const userPwPepper = "secret-random-string"
const hmacSecretKey = "secret-hmac-key"

# models/service.go
db, err := gorm.Open("postgres", connectionInfo)
db.LogMode(true)
