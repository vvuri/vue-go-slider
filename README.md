# vue-go-slider

Steps
1. git clone
2. $ go mod init vue-go-slider
3. create main.go
4. $ go get github.com/gofiber/fiber/v2
5. $ go mod vendor
6. Write API func in main on https://docs.gofiber.io/
7. $ go get github.com/lpernett/godotenv
- Env: https://pkg.go.dev/github.com/lpernett/godotenv#section-readme
- Add file .env file
8. MongoDB
- $ go get go.mongodb.org/mongo-driver/mongo
- https://www.mongodb.com/try/download/community-kubernetes-operator
- download msi 7.0.11 and install -> c:/ext/...
- MongoDB Shell 2.2.6 - download and install - no need install
- Compass - New database: vue-go-slider. collection: slides
- Connection string:   mongodb://localhost:27017/
- add Shell to path
9. Remove env from git
- gitignore .env
- git rm .env --cached
