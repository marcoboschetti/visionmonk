$(document).ready(function () {

    // Define CRUD endpoints
    var getEndpoint = "/api/p/shop_products";
    var createEndpoint = "/api/p/shop_product";
    var updateEndpoint = "/api/p/shop_product/";
    var deleteEndpoint = "/api/p/shop_product/";
    var entityDisplayName = "producto";

    // Table display columns
    var columns = [
        { title: "SKU", data: "sku" },
        {
            data: null,
            title:"Producto",
            render: function (data, type, row) {
                return `<img class="product-table-img" src="`+data.catalog_product.image_base_64+`">`+ data.catalog_product.title;
            }
        },
        // { title: "Código de barras", data: "barcode" },
        { title: "Marca", data: "catalog_product.brand" },
        {
            data: null,
            title:"Color - Tamaño",
            render: function (data, type, row) {
                return data.catalog_product.color+` - `+ data.catalog_product.size;
            }
        },
        {
            data: null,
            render: function (data, type, row) {
                return `<a class="d-none d-sm-inline-block btn btn-primary table-btn crud-view-btn">
            <i class="fas fa-eye text-white-80"></i> 
          </a>
          <a class="d-none d-sm-inline-block btn btn-success table-btn crud-edit-btn">
            <i class="fas fa-edit text-white-80"></i> 
          </a>
          <a class="d-none d-sm-inline-block btn btn-danger table-btn crud-delete-btn">
            <i class="fas fa-trash text-white-80"></i> 
          </a>`;
            }
        },
    ];

    setupModalListeners();

    // Configure endpoints, callbacks, etc
    setupCrud(
        getEndpoint, createEndpoint, updateEndpoint, deleteEndpoint,
        entityDisplayName, columns,
        updateProductModal, loadProductInputData, populateDeleteModal
    );
});

function updateProductModal(p) {
    var c = p.catalog_product;
    if(!c){
        c = {};
    }
    lastLoadedImageBase64 = c.image_base_64;

    $("#skuInput").val(p.sku)
    $("#priceInput").val(p.price_cts/100)
    $("#inventoryInput").val(p.inventory)

    $("#titleInput").val(c.title)
    $("#descriptionInput").val(c.description)
    $("#barcodeInput").val(c.barcode)
    $("#categoryInput").val(c.category)
    $("#brandInput").val(c.brand)
    $("#colorInput").val(c.color)
    $("#sizeInput").val(c.size)

    $("#productImagePreview").attr("src", c.image_base_64);
    if (c.image_base_64) {
        $("#productImagePreview").slideDown();
    } else {
        $("#productImagePreview").slideUp();
    }

    $("#barcodeVideoContainer").slideUp();
}


function loadProductInputData() {
    var catalog = {
        "title": $("#titleInput").val(),
        "description": $("#descriptionInput").val(),
        "barcode": $("#barcodeInput").val(),
        "image_base_64": lastLoadedImageBase64,
        "category": $("#categoryInput").val(),
        "brand": $("#brandInput").val(),
        "color": $("#colorInput").val(),
        "size": $("#sizeInput").val(),
    }
    return {
        "sku": $("#skuInput").val(),
        "inventory": parseInt($("#inventoryInput").val()),
        "price_cts": parseInt($("#priceInput").val())*100,
        "catalog_product": catalog,
    }
}

function populateDeleteModal(data) {
    selectedDeleteID = data.id;
    $("#deleteProductName").html(data.catalog_product.title +" ("+data.sku+")");
}

var maxSize = 1000000 // 1MB
var isBarcodeVisible;
var lastLoadedImageBase64;
function setupModalListeners() {
    $('#imageBase64Input').bind('change', function () {
        if (this.files.length == 0) { return; }

        $("#productImagePreview").slideUp();
        $("#imageFileSizeError").slideUp();

        var file = this.files[0];
        if (file.size > maxSize) {
            $("#imageFileSizeError").slideDown();
            return;
        }

        getBase64(file, function (base64) {
            $("#productImagePreview").attr("src", base64);
            $("#productImagePreview").slideDown();
            lastLoadedImageBase64 = base64;
        });
    });

    $("#iterateScannerBtn").click(iterateInputDevice);

    $("#barcodeScannerBtn").click(function () {
        if (isBarcodeVisible) {
            $("#barcodeVideoContainer").slideUp();
            resetDecoder();
            isBarcodeVisible = false;
        } else {
            loadDecoder("barcodeVideo", true,
                function (ans) {
                    $("#barcodeInput").val(ans);
                    $("#barcodeVideoContainer").slideUp();
                },
                function (scanerSources) {
                    console.log("scanerSources", scanerSources);
                }
            );
            $("#barcodeVideoContainer").slideDown();
            isBarcodeVisible = true;
        }
    })
}