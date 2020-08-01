## How to setup Beego API project

1. First, set the environment variable (bash profile) `export GO111MODULE=on` 

2. Execute command `go get -u github.com/beego/bee` in order to get access to `bee` command

3. Run command `bee api <folder_project_name>` for bootstrapping the initial content of a beego project

4. cd to the directory of the project and run `go mod init`

5. Run `go mod tidy` to get all the dependencies

## Run the project 

Use this command to build/run and refresh the swagger documentation:

`bee run -downdoc=true -gendoc=true`

Open your browser and enter following in the address bar: 

http://localhost:8080/swagger/

### Resources

You can check out more about Beego on following link:

https://github.com/beego/bee

https://beego.me/blog/beego_api
