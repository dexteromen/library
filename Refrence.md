# Link

https://medium.com/@rashid14713524/a-simple-todo-crud-api-using-golang-gin-gorm-and-postgresql-981a9bde0c4d

https://www.bacancytechnology.com/blog/golang-jwt

https://auth0.com/blog/authentication-in-golang/

https://app.studyraid.com/en/read/5926/130188/working-with-cookies

# Swagger

1. Install

```
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

```

2. Add Swag Annotations to Your Handlers

3. swag init

4. http://localhost:8080/swagger/index.html

# install swag

1. Install Swag Properly
   Run the following command to install swag globally:

go install github.com/swaggo/swag/cmd/swag@latest
After installation, swag should be available in your GOPATH/bin.

2. Add Go Bin Directory to PATH
   If swag is still not found, it might not be in your PATH. Add it manually:

export PATH=$PATH:$(go env GOPATH)/bin
To make this change permanent, add the line above to your shell configuration file:

For Bash:
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
For Zsh:
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
source ~/.zshrc 3. Verify Installation
Run:
swag --version
If it prints the version, swag is installed correctly.

4. Generate Swagger Docs
   Now you can run:

swag init
This should generate the docs/ folder.
