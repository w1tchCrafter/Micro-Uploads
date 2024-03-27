const form = document.querySelector("form");
const fileInput = form.querySelector(".file-input");
const progressArea = document.querySelector(".progress-area");
const uploadedArea = document.querySelector(".uploaded-area");

form.addEventListener("click", () => {
    fileInput.click();
});

fileInput.onchange = ({target}) => {
    let [file] = target.files; // selecting target.files[0]

    if (file) {
        let filename = file.name;
        uploadFile(filename);
    }
};

function uploadFile(name) {
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "api/v1/uploads/create");
    xhr.upload.addEventListener("progress", ({loaded, total}) => {
        let fileLoaded = Math.floor((loaded / total) * 100); // getting percentage of loaded file size
        let fileTotal = Math.floor(total / 1000);// getting filesize in KB from bytes
        let progressHTML = `<li class="row">
                                <i class="fas fa-file-alt"></i>
                                <div class="content">
                                    <div class="details">
                                        <span class="name">${name} * Uploading</span>
                                        <span class="percent">${fileLoaded}%</span>
                                    </div>
                                    <div class="progress-bar">
                                        <div class="progress" style="width: ${fileLoaded}%></div>
                                    </div>
                                </div>
                            </li>`;
        let uploadedHTML = `<li class="row">
                                <i class="fas fa-file-alt"></i>
                                <div class="content">
                                    <div class="details">
                                        <span class="name">placeHolder_img2 * Uploaded</span>
                                        <span class="size">70Kb</span>
                                    </div>
                                    <div class="progress-bar">
                                        <div class="progress"></div>
                                    </div>
                                </div>
                                <i class="fas fa-check"></i>
                            </li>`;
        progressArea.innerHTML = progressHTML;
    });

    let formData = new FormData(form);
    xhr.send(formData);
}
