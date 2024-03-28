const logout = document.querySelector("#logout");

logout.addEventListener("click",async () => {
    await fetch("/api/v1/auth/logout");
    window.location.href = "/";
});
