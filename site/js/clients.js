$(document).ready(function () {

    // Define CRUD endpoints
    var getEndpoint = "/api/p/clients";
    var createEndpoint = "/api/p/client";
    var updateEndpoint = "/api/p/client/";
    var deleteEndpoint = "/api/p/client/";
    var entityDisplayName = "cliente";

    // Table display columns
    var columns = [
        { title: "Nombre", data: "first_name" },
        { title: "Apellido", data: "last_name" },
        { title: "Email", data: "email" },
        { title: "DNI/CUIL", data: "identification" },
        { title: "Teléfono", data: "phone" },
        {
            data: null,
            render: function (data, type, row) {
                return `<a class="d-none d-sm-inline-block btn btn-primary table-btn crud-view-btn">
            <i class="fas fa-eye text-white-80"></i> 
          </a>`;
            }
        },
        {
            data: null,
            render: function (data, type, row) {
                return `<a class="d-none d-sm-inline-block btn btn-success table-btn crud-edit-btn">
            <i class="fas fa-edit text-white-80"></i> 
          </a>`;
            }
        },
        {
            data: null,
            render: function (data, type, row) {
                return `<a class="d-none d-sm-inline-block btn btn-danger table-btn crud-delete-btn">
            <i class="fas fa-trash text-white-80"></i> 
          </a>`;
            }
        },
    ];

    // Configure endpoints, callbacks, etc
    setupCrud(
        getEndpoint, createEndpoint, updateEndpoint, deleteEndpoint,
        entityDisplayName, columns,
        updateClientModal, loadClientInputData, populateDeleteModal
    );
});

function updateClientModal(c) {
    $("#firstNameInput").val(c.first_name)
    $("#lastNameInput").val(c.last_name)
    $("#identificationInput").val(c.identification)
    $("#socialWorkInput").val(c.social_work)
    $("#socialWorkNumberInput").val(c.social_work_number)
    $("#phoneInput").val(c.phone)
    $("#emailInput").val(c.email)
    $("#addessInput").val(c.addess)
    $("#localityInput").val(c.locality)
    $("#neighborhoodInput").val(c.neighborhood)
    $("#zipCodeInput").val(c.zipcode)

    var aditional_info = c.aditional_info;
    if (!aditional_info) {
        aditional_info = "";
    }
    $("#aditionalInfoInput").html(aditional_info)

    $("#finalConsumerInput").prop('checked', c.final_consumer == true);
}


function loadClientInputData() {
    return {
        "first_name": $("#firstNameInput").val(),
        "last_name": $("#lastNameInput").val(),
        "identification": $("#identificationInput").val(),
        "social_work": $("#socialWorkInput").val(),
        "social_work_number": $("#socialWorkNumberInput").val(),
        "final_consumer": $("#finalConsumerInput").is(":checked"),
        "phone": $("#phoneInput").val(),
        "email": $("#emailInput").val(),
        "addess": $("#addessInput").val(),
        "locality": $("#localityInput").val(),
        "neighborhood": $("#neighborhoodInput").val(),
        "zipcode": $("#zipCodeInput").val(),
        "aditional_info": $("#aditionalInfoInput").html(),
    }
}

function populateDeleteModal(data) {
    selectedDeleteID = data.id;
    $("#deleteClientName").html(data.first_name + " " + data.last_name);
    $("#deleteClientID").html(data.identification);
}