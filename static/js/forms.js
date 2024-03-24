class Form {
    constructor(action, method, formElement, expectedResponse) {
        this.action = action;
        this.method = method;
        this.formElement = formElement;
        this.expectedResponse = expectedResponse;
    }

    activate(callback) {
        this.formElement.addEventListener("submit", async() => {
            event.preventDefault();
            const body = new FormData(this.formElement);
    
            const response = await fetch(this.action, {
                method: this.method,
                body
            });
    
            if (response.status === this.expectedResponse) {
                callback();
            }
        });
    }
}

const login = document.querySelector("#loginform");
const register = document.querySelector("#registerform");

const loginForm = new Form("/api/v1/auth/login", "post", login, 200);
const registerForm = new Form("/api/v1/auth/register", "post", register, 201);

loginForm.activate(() => document.location = "/user");
registerForm.activate(() => document.location = "/user");
