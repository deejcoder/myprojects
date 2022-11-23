# myprojects-api

This is the main repository for the backend of the `myprojects` project. This project will allow software developers to showcase their software projects. This web application will allow the host (software developer, or system administrator), to edit, delete and create projects. A project will be able to have images attached to it, and embedded within the content.
For the frontend, refer to [this](https://github.com/deejcoder/myprojects-http) repository.

## services

The API will provide RESTful replies to the frontend in React. The API solely manages the communication to the database, as well as authorization. 

### replies
An example response from the API is below, where all API responses will be in the same format. The logic behind replies are located in the `reply` folder.
```json
{
    "ok": false,
    "errorCode": 1,
    "message": "Validation errors when updating the project",
    "body": {
        "validationErrors": [
            {"error": "Title should be between 5 and 80 characters", "path": "title"}
        ],
        "data": null
    }
}
```

### authorization
The client must send a login request with the secret key provided in the backend's configuration. This is basic authorization which can be extended in the future. The backend will then reply with a JWT token if it's correct. The client may validate this token at any time for validity. Whenever accessing resources that require authorization, the token should be included in the `authorization` header of the HTTP request in the form, `Bearer {token}`. Refer to `api/auth.go` & `api/endpoints.go`.

### storage
The backend uses MongoDB as a NoSQL DMS. All database logic has been thrown into the folder, `storage`, where schemas have their own file which also includes validation.

### configuration
To configure this web app, you may modify `config.example.yaml`, and rename it to `config.yaml`, however avoid sharing this file.

## deployment
First build the Go project on the local machine,
`GOOS=linux GOARCH=amd64 go build` and transfer it to the remote machine.

Then execute the following commands on the remote machine,
```bash
mkdir -p /home/apps/myprojects
mv /path/to/myprojects ./ // (executable)
mv /path/to/config.yaml ./
chown u+x ./myprojects

// install mongodb
apt-get install gnupg2
wget -qO - https://www.mongodb.org/static/pgp/server-4.2.asc | sudo apt-key add -
echo "deb http://repo.mongodb.org/apt/debian stretch/mongodb-org/4.2 main" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.2.list
apt-get update

// if using Debian Buster, you require packages from Debian Stretch
echo "deb http://deb.debian.org/debian/ stretch main" | tee /etc/apt/sources.list.d/debian-stretch.list
apt get update
apt get install libcurl3

apt-get install -y mongodb-org

// restart and enable mongodb service
systemctl start mongod
systemctl enable mongod

// TODO: install and configure PM2 for the API, for now use nohup to serve the API ...
nohup ./myprojects serve &

// allow port 8080
sudo ufw allow 8080
```

You should now be able to visit the API by visiting http://{domain}:8080
HTTPS will be implemented in the future.
