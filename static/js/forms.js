class Form {
    constructor(action, method, formElement, expectedResponse) {
        this.action = action;
        this.method = method;
        this.formElement = formElement;
        this.expectedResponse = expectedResponse;
    }

    activate(success, err) {
        this.formElement.addEventListener("submit", async() => {
            event.preventDefault();
            let body = new FormData(this.formElement);
    
            let response = await fetch(this.action, {
                method: this.method,
                body
            });
    
            if (response.status === this.expectedResponse) {
                success();
            } else {
                let json = await response.json();
                err(json);
            }
        });
    }
}

const login = document.querySelector("#loginform");
const register = document.querySelector("#registerform");

let loginForm = new Form("/api/v1/auth/login", "post", login, 200);
let registerForm = new Form("/api/v1/auth/register", "post", register, 201);

loginForm.activate(() => document.location = "/user", json => {
    let err = json["error"];
    
    switch (err) {
        case ACCESS_DENIED:
            // do something
            break;

        case BAD_REQUEST:
            // do something else
            break;

        case SERVER_ERR:
            // do way something else
            break;

        default:
            break;
    }
});

registerForm.activate(() => document.location = "/user", json => {
    let err = json["error"];
    
    switch (err) {
        case USER_EXISTS:
            // do something
            break;

        case BAD_REQUEST:
            // do something else
            break;

        case SERVER_ERR:
            // do way something else
            break;

        default:
            break;
    }
});
