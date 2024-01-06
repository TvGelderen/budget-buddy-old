function showError(id, event) {
    const element = document.getElementById(id);
    if (!element) return;

    element.innerHTML = event.detail.xhr.response;
}
