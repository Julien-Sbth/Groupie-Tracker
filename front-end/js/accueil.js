function myFunction() {
    var input, filter, ul, li, a, i, txtValue;
    input = document.getElementById("myInput");
    filter = input.value.toUpperCase();
    ul = document.getElementById("myUL");
    li = ul.getElementsByTagName("li");
    for (i = 0; i < li.length; i++) {
        a = li[i].getElementsByTagName("a")[0];
        txtValue = a.textContent || a.innerText;
        if (txtValue.toUpperCase().indexOf(filter) > -1) {
            li[i].style.display = "";
        } else {
            li[i].style.display = "none";
        }
    }
}
function uncheckOthers(checkbox) {
    if (checkbox.checked) {
        // uncheck all other checkboxes with the same name
        document.querySelectorAll(`input[name="${checkbox.name}"]:not([value="${checkbox.value}"])`).forEach(cb => cb.checked = false);
    }
}
function updateYearValue() {
    // update the value next to the slider
    document.getElementById("yearValue").textContent = document.getElementById("year").value;
}


function uncheckOthers(currentCheckbox) {
    let checkboxes = document.getElementsByName(currentCheckbox.name);

    for (let i = 0; i < checkboxes.length; i++) {
        if (checkboxes[i] !== currentCheckbox && checkboxes[i].checked) {
            checkboxes[i].checked = false;
        }
    }
}
const checkboxes = document.querySelectorAll('input[type="checkbox"][name="option1"]:checked');
const selectedValues = [];
checkboxes.forEach((checkbox) => {
    selectedValues.push(checkbox.value);
});

const yearInput = document.querySelector('input[name="year"]');
const selectedYear = yearInput.value;

document.querySelector("form").addEventListener("submit", function(event) {
    event.preventDefault(); // prevent default form submission
    const formData = new FormData(this);
    const persons = formData.getAll("persons");
    const year = formData.get("year");
    const urlSearchParams = new URLSearchParams();
    if (persons.length > 0) {
        urlSearchParams.append("persons", persons.join(","));
    }
    if (year) {
        urlSearchParams.append("year", year);
    }
    // redirect to the search results page with the filter parameters as query string
    window.location.href = `/search-results?${urlSearchParams.toString()}`;
});


const bouton = document.getElementById('mon-bouton');
const zoneTexte = document.getElementById('ma-zone-texte');

zoneTexte.style.display = 'none';

let isHidden = true;

bouton.addEventListener('click', function() {
    if (isHidden) {
        zoneTexte.style.display = 'block';
        isHidden = false;
        bouton.style.display = 'none';
    } else {
        zoneTexte.style.display = 'none';
        isHidden = true;
    }
});

const input = document.getElementById('year');
input.addEventListener('click', function() {
    if (isHidden) {
        bouton.style.display = 'block';
    }
});