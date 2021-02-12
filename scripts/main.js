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

// JavaScript module for alerts and toasts
function Prompt() {
    // Toasts
    let toast = function (c) {
        const {
            message = "",
            icon = "success",
            position = "top-end",
        } = c;
        const Toast = Swal.mixin({
            toast: true,
            title: message,
            position: position,
            icon: icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({

        })
    }
    let success = function (c) {
        const {
            message = "",
            title = "",
            footer = ""
        } = c

        Swal.fire({
            icon: 'success',
            title: title,
            text: message,
            footer: footer
        })
    }
    let error = function (c) {
        const {
            message = "",
            title = "",
            footer = ""
        } = c

        Swal.fire({
            icon: 'error',
            title: title,
            text: message,
            footer: footer
        })
    }
    // Custom alert modal
    async function custom(c) {
        const {
            message = "",
            title = "",
        } = c;

        const { value: formValues } = await Swal.fire({
            title: title,
            html: message,
            backdrop: false,
            focusConfirm: false,
            showCancelButton: true,
            willOpen: () => {
                const elem = document.getElementById('reservation-dates-modal');
                const rp = new DateRangePicker(elem, {
                    format: 'yyyy-mm-dd',
                    showOnFocus: true,
                })
            },
            preConfirm: () => {
                console.log("this worked");
                return [
                    document.getElementById('start').value,
                    document.getElementById('end').value
                ]
            },
            didOpen: () => {
                document.getElementById('start').removeAttribute('disabled');
                document.getElementById('end').removeAttribute('disabled');
            }
        })

        if (formValues) {
            Swal.fire(JSON.stringify(formValues))
        }
    }
    return {
        toast: toast,
        success: success,
        error: error,
        custom: custom,
    }
}

// Test Button
document.getElementById("btnTest").addEventListener("click", function () {
    let html = `
        <form id="check-availability-form action="" method="POST" class="needs-validation" novalidate>
            <div class="form-row">
                <div class="col">
                    <div class="form-row" id="reservation-dates-modal">
                        <div class="col">
                            <input disabled required type="text" class="form-control" id="start" placeholder="Arrival Date" />
                        </div>
                        <div class="col">
                            <input disabled required type="text" class="form-control" id="end" placeholder="Departure Date" />
                        </div>
                    </div>
                </div>
            </div>
        </form>
    `
    attention.custom({ message: html, title: "Choose your dates" });
})