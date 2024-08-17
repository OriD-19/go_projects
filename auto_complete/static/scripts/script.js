function fillSearch(clickedElement) {
    const el = document.getElementById("search");

    el.value = clickedElement.textContent.trim();
}
