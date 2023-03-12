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

function updateYearValue() {
    document.getElementById("yearValue").innerHTML = document.getElementById("year").value;
}

function uncheckOthers(currentCheckbox) {
    let checkboxes = document.getElementsByName(currentCheckbox.name);

    for (let i = 0; i < checkboxes.length; i++) {
        if (checkboxes[i] !== currentCheckbox && checkboxes[i].checked) {
            checkboxes[i].checked = false;
        }
    }
}

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