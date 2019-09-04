# myprojects-api

This is the main repository for the backend of the `myprojects` project. This project will allow software developers to showcase their software projects. This web application will allow the host (software developer, or system administrator), to edit, delete and create projects. A project will be able to have images attached to it, and embedded within the content.

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