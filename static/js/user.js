const form = document.querySelector("form");
const fileInput = form.querySelector(".file-input");
const progressArea = document.querySelector(".progress-area");
const uploadedArea = document.querySelector(".uploaded-area");

form.addEventListener("click", () => {
  fileInput.click();
});

fileInput.onchange = ({ target }) => {
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
  xhr.open("POST", "api/v1/uploads/");
  xhr.upload.addEventListener("progress", ({ loaded, total }) => {
    const KB = 1 << 10;
    const MB = 1 << 20;
    let fileLoaded = Math.floor((loaded / total) * 100); // getting percentage of loaded file size
    let fileTotal = Math.floor(total);
    let fileSize;

    if (fileTotal > MB) {
      fileSize = Math.floor(fileTotal / MB) + "MB";
    } else if (fileTotal > KB) {
      fileSize = Math.floor(fileTotal / KB) + "KB";
    } else {
      fileSize = fileTotal + "bytes";
    }

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
                                <i class="fas fa-download"></i>
                            </li>`;
      uploadedArea.insertAdjacentHTML("afterend", uploadedHTML);
      uploadedArea.querySelectorAll(".name").forEach((v) =>
        v.style.color = "#C5D1DE"
      );
    }
  });

  let formData = new FormData(form);
  xhr.send(formData);
}
