let attention = Prompt();

// Validation and error checking
(function () {
    'use strict';
    window.addEventListener('load', function () {
        // Fetch all the forms we want to apply custom Bootstrap validation styles to
        var forms = document.getElementsByClassName('needs-validation');
        // Loop over them and prevent submission
        var validation = Array.prototype.filter.call(forms, function (form) {
            form.addEventListener('submit', function (event) {
                if (form.checkValidity() === false) {
                    event.preventDefault();
                    event.stopPropagation();
                }
                form.classList.add('was-validated');
            }, false);
        });
    }, false);
})();

// Notification banners with the notie library
function notify(message, messageType) {
    notie.alert({
        type: messageType,
        text: message
    })
}

// Custom date picker with the vanillajs-datepicker library 
const elem = document.getElementById('reservationDates');
const rangepicker = new DateRangePicker(elem, {
    format: "yyyy-mm-dd"
});

// Modal function with the sweetalert2 library
function notifyModal(title, text, icon, confirmationButtonText) {
    Swal.fire({
        title: title,
        html: text,
        icon: icon,
        confirmButtonText: confirmationButtonText
    })
}

// JavaScript module for prompting users
function Prompt() {
    let toast = function () {
        console.log("clicked button and got toast")
    }
    let success = function () {
        console.log("clicked button and got success")
    }
    return {
        toast: toast,
        success: success,
    }
}

// Test Button
document.getElementById("btnTest").addEventListener("click", function () {
    // notify("This is my Message", "warning");
    // notifyModal("title", "<em>Hello World!</em>", "success", "Text for my Button");
    attention.success();
})