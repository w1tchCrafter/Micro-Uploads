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
