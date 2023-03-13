$(document).ready(function () {

    $("#registerNewAccountBtn").click(postNewUser);

    $('#passwordRepeatInput').keyup(function () {
        'use strict';
    
        if ($('#passwordInput').val() === $(this).val()) {
            $('#pass_hint').html('match');
            this.setCustomValidity('');
        } else {
            $('#pass_hint').html('mismatch');
            this.setCustomValidity('Passwords must match');
        }
    });
});

function postNewUser() {
    var registerForm = $("#registerForm");
    if (!registerForm.isValid()) {
        registerForm.addClass("was-validated");
    } else {
        $("#invalidRegister").slideUp();
        $("#registerNewAccountBtn").addClass("disabled");
    
        registerForm.removeClass("was-validated");

        var payload = {
            shop_token: $("#shopInvitationToken").val(),
            firstname: $("#firstNameInput").val(),
            lastname: $("#lastNameInput").val(),
            email: $("#emailInput").val(),
            password: $("#passwordInput").val(),
        }

        $.post("/api/user/new", JSON.stringify(payload), function (data) {
            loginPostCredentials(payload.email,payload.password);            
        }).fail(function (err) {
            var text = err.responseJSON.message;
            if(text == "email info already exists"){
                text = "El email ya se encuentra registrado."
            }else if(text == "invalid shop token"){
                text = "El código de invitación es incorrecto."
            }
            $("#invalidRegister").html(text);
            $("#invalidRegister").slideDown();
            $("#registerNewAccountBtn").removeClass("disabled");
        });;
    }

}

