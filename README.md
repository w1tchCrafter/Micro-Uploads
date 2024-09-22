# Micro Uploads

Micro Uploads is a minimalist file upload website built using Go. It allows users to upload files and manage them easily.
Features

- Simple and minimalistic user interface.
- Easy file uploading and management.
- Built with Go for fast performance.

Getting Started

To build the project, run:

```bash
make
```

To run tests, use:

```bash
make test
```

## Usage

Clone this repository.  
Build the project using make.  
Run the compiled binary.  
Visit the website in your browser.  
Upload files using the provided interface.  

## Api Documentation

### Endpoints

### Home page 
- Endpoint: `Get /`
- Description: return the html of the homepage
- Response:
    - should always return html and a 200 status code

### user page
- Endpoint: `Get /user`
- Description: returns the personal user page
- Response:
    - Return the page and status code 200 on success
    - Return status code 403 in case of authentication error
    - Return status code 500 on server error

### register
- Endpoint: `POST /register`
- Description: Request the server tp register a new user
- Request Body (form):
    ```json
    {
        "username" : "John Doe",
        "password" : "12345678"
    }
    ```
- Response: 
    - Return status code 400 if requested username already exists in the database
    - Return status code 500 on server error
    - Return status code 2101 on success

### login
- Endpoint: `POST /login`
- Description: start a session on the server
- Request Body (form):
    ```json
    {
        "username" : "anonymous",
        "password" : "876543210"
    }
    ```
- Response:
    - Return status code 200 on success
    - Return status code 401 on failure
    - Return status code 400 on wrong data format

### logout
- Endpoint: `GET /logout`
- Description: logs the user out
- Response:
    - Return status code 200 on success
    - Return status code 500 on server error


### uploads
- Endpoint: `POST /uploads`
- Description: uploads a file to the server via form
- Response:
    - Return status code 201 on success
    - Return status code 400 on wrong data format
    - Return status code 500 on server error


### getfile
- Endpoint: `GET /uploads/{filename}`
- Description: download/open the given file from the server
- Response:
    - Return status code 200 on success
    - Return status code 404 if not found

### delete
- Endpoint: `DELETE /uploads/{filename}`
- Description: download/open the given file from the server
- Response:
    - Return status code 200 on success
    - Return status code 491 if the user do not own the file
    - Return status code 404 if not found
    - Return status code 500 on server error

## Contributing

Contributions are welcome! If you'd like to contribute to Micro Uploads, please fork the repository and submit a pull request.
