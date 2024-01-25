function outsideClickListener(event, elementRect) {
    if (event.x < elementRect.x ||
        event.x > elementRect.x + elementRect.width ||
        event.y < elementRect.y ||
        event.y > elementRect.y + elementRect.height
    ) {
        toggleUserDropdown();
    }
}

function toggleNav() {
    const navBar = document.getElementById("nav-bar");
    if (!navBar) return;

    if (!navBar.hasAttribute("open")) {
        navBar.setAttribute("open", "");
    } else {
        navBar.removeAttribute("open");
    }
}

let dropdownRect;
const dropdownListener = event => outsideClickListener(event, dropdownRect);

function toggleUserDropdown() {
    dropdown = document.getElementById("user-dropdown");
    if (!dropdown) return;

    dropdownRect = dropdown.getBoundingClientRect();

    if (!dropdown.hasAttribute("open")) {
        dropdown.setAttribute("open", "");
        document.addEventListener('click', dropdownListener);
    } else {
        dropdown.removeAttribute("open");
        document.removeEventListener('click', dropdownListener);
    }
}

function showError(id, event) {
    const element = document.getElementById(id);
    if (!element) return;

    element.innerHTML = event.detail.xhr.response;
}

function toggleModal(id) {
    const modal = document.getElementById(id);
    if (!modal) return;

    if (modal.hasAttribute("open")) {
        modal.close();
    } else {
        modal.showModal();
    }
}

function closeModalOnSuccess(id, event) {
    const modal = document.getElementById(id);
    if (!modal) return;

    if (event.successful) {
        modal.close();
    }
}

function getUpdateTransactionId() {
    return localStorage.getItem('updateTransactionId');
}
