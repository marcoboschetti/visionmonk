// State
var selectedEditID;
var selectedDeleteID;
var displayName;

// Local Setup functions
var updateEntityModal;
var getModalPayload;
var populateDeleteModal;

// API endpoints
var getEntitiesEndpoint;
var updateEntitiesEndpoint;
var createEntitiesEndpoint;
var deleteEntitiesEndpoint;

function setupCrud(getEndpoint, createEndpoint, updateEndpoint, deleteEndpoint,
    entityDisplayName, columns,
    updateEntityModalFunc, getModalPayloadFunc, populateDeleteModalFunc) {

    displayName = entityDisplayName;

    getEntitiesEndpoint = getEndpoint;
    createEntitiesEndpoint = createEndpoint;
    updateEntitiesEndpoint = updateEndpoint;
    deleteEntitiesEndpoint = deleteEndpoint;

    updateEntityModal = updateEntityModalFunc;
    getModalPayload = getModalPayloadFunc;
    populateDeleteModal = populateDeleteModalFunc;

    getUserData(populateTable(columns));

    $("#addEntityBtn").click(function () { openEditModal({}, true, true) });
    $("#createEntityBtn").click(postNewEntity);
    $("#updateEntityBtn").click(postUpdateEntity);
    $("#confirmDeleteEntityBtn").click(postDeleteEntity);
}

function populateTable(columns) {
    return function (user) {
        $.get(getEntitiesEndpoint, function (data) {
            console.log(data.records);
            var table = $("#crudTable").DataTable({
                data: data.records,
                language: {
                    url: '/site/libraries/dataTables/es-ES.json'
                },
                columns: columns
            });
            $('#crudTable').on('error.dt', function (e, settings, techNote, message) {
                console.log('An error has been reported by DataTables: ', message);
            });

            $('#crudTable tbody').on('click', '.crud-view-btn', function () {
                var data = table.row($(this).parents('tr')).data();
                openEditModal(data, false, false);
            });
            $('#crudTable tbody').on('click', '.crud-edit-btn', function () {
                var data = table.row($(this).parents('tr')).data();
                openEditModal(data, false, true);
            });
            $('#crudTable tbody').on('click', '.crud-delete-btn', function () {
                var data = table.row($(this).parents('tr')).data();
                populateDeleteModal(data);
                $("#deleteEntityModal").modal('show');
            });

            $("#loadingContainer").slideUp();
            if (!data || !data.records || !data.records.length) {
                $("#noDataContainer").slideDown();
            } else {
                $("#dataContainer").slideDown();
            }
        })
    }
}

function openEditModal(entity, isNewEntity, allowEdit) {
    selectedEditID = entity.id;

    // Reset modal to input new entity
    updateEntityModal(entity);

    // Change title and action buttons
    $("#updateEntityBtn").parent().hide();
    $("#createEntityBtn").parent().hide();
    $("#barcodeScannerBtn").hide();

    if (isNewEntity) {
        $("#crudModalTitle").html("Agregar nuevo " + displayName);
        $("#createEntityBtn").parent().show();
        $("#barcodeScannerBtn").show();
    } else if (allowEdit) {
        $("#crudModalTitle").html("Editar " + displayName);
        $("#updateEntityBtn").parent().show();
        $("#barcodeScannerBtn").show();
    } else {
        $("#crudModalTitle").html("Ver " + displayName);
    }

    if (allowEdit || isNewEntity) {
        $('.form-control').removeAttr('readonly');
        $(".form-check-input").removeAttr('disabled');
        $(".custom-file-input").removeAttr('disabled');
        $('.__toolbar').show();
        $("[data-tiny-editor]").attr("contenteditable", true);
    } else {
        $('.form-control').attr('readonly', '');
        $(".form-check-input").attr('disabled', '');
        $(".custom-file-input").attr('disabled', '');
        $('.__toolbar').hide();
        $("[data-tiny-editor]").attr("contenteditable", false);
    }

    $('#crudModal').modal('show');
}

function postNewEntity() {
    $("#createEntityBtn").addClass("disabled");

    var payload = getModalPayload();
    $.post(createEntitiesEndpoint, JSON.stringify(payload), function (data) {
        $('#crudModal').modal('hide');
        $('body').removeClass('modal-open');
        $('.modal-backdrop').remove();
        loadPageFromHash();
    });
}

function postUpdateEntity() {
    $("#updateEntityBtn").addClass("disabled");

    var payload = getModalPayload();
    $.put(updateEntitiesEndpoint + selectedEditID, JSON.stringify(payload), function (data) {
        $('#crudModal').modal('hide');
        $('body').removeClass('modal-open');
        $('.modal-backdrop').remove();
        loadPageFromHash();
    });
}

function postDeleteEntity() {
    $("#confirmDeleteEntityBtn").addClass("disabled");

    $.delete(deleteEntitiesEndpoint + selectedDeleteID, function (data) {
        $('#crudModal').modal('hide');
        $('body').removeClass('modal-open');
        $('.modal-backdrop').remove();
        loadPageFromHash();
    });
}
