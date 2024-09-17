const logout = document.querySelector("#logout");

let apiResponses = {
  SUCCESS_UPDATE: "file was uploaded successfully",
  CREATED_USER: "user was created successfully",
  LOGIN_MSG: "user logged successfully",
  LOGOUT_MSG: "user logged out successfully",
  USER_EXISTS: "user already exists",
  ACCESS_DENIED: "access denied, wrong credentials or user do not exists",
  NOT_FOUND: "page not found",
  SERVER_ERR: "an error ocurred, try again later",
  BAD_REQUEST: "error, bad request",
};

logout.addEventListener("click", async () => {
  await fetch("/api/v1/auth/logout");
  window.location.href = "/";
});
