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

        if (filename.length >= 12) {
            let splitname = filename.split(".");
            filename = splitname[0].substring(0, 12) + "... ." + splitname[1];
        }

        uploadFile(filename);
    }
};

function uploadFile(name) {
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "api/v1/uploads/create");
    xhr.upload.addEventListener("progress", ({loaded, total}) => {
        let fileLoaded = Math.floor((loaded / total) * 100); // getting percentage of loaded file size
        let fileTotal = Math.floor(total / 1000);// getting filesize in KB from bytes
        let fileSize;
        (fileTotal > 1024) ? fileSize = fileTotal + "KB": fileSize = (loaded / (1024 * 1024)).toFixed(2) + "MB";
        let progressHTML = `<li class="row">
                                <i class="fas fa-file-alt"></i>
                                <div class="content">
                                    <div class="details">
                                        <span class="name">${name} * Uploading</span>
                                        <span class="percent">${fileLoaded}%</span>
                                    </div>
                                    <div class="progress-bar">
                                        <div class="progress"></div>
                                    </div>
                                </div>
                            </li>`;
        progressArea.innerHTML = progressHTML;
        progressArea.querySelector(".progress").style.width = `${fileLoaded}%`;

        if (loaded === total) {
            progressArea.innerHTML = "";
            let uploadedHTML = `<li class="row">
                                <i class="fas fa-file-alt"></i>
                                <div class="content">
                                    <div class="details">
                                        <span class="name">${name} * Uploaded</span>
                                        <span class="size">${fileSize}</span>
                                    </div>
                                    <div class="progress-bar">
                                        <div class="progress"></div>
                                    </div>
                                </div>
                                <i class="fas fa-check"></i>
                            </li>`;
            uploadedArea.insertAdjacentHTML("afterbegin", uploadedHTML);
        }
    });

    let formData = new FormData(form);
    xhr.send(formData);
}
