$(document).ready(function () {

    $("#loginBtn").click(login);


});

function login() {
    var loginForm = $("#loginForm");
    $("#invalidCredentials").slideUp();
    $("#loginBtn").addClass("disabled");

    if (!loginForm.isValid()) {
        loginForm.addClass("was-validated");
    } else {
        loginForm.removeClass("was-validated");
        loginPostCredentials($("#emailInput").val(), $("#passwordInput").val(),
            function () {
                $("#loginBtn").removeClass("disabled");
                $("#invalidCredentials").slideDown();
            }
        )
    }
}

