$(document).ready(function () {
    window.onhashchange = function () {
        var page = document.location.hash.substring(1);
        moveToPage(page);
    };

    loadPageFromHash();
    loadUser();
});

function loadPageFromHash(){
    var page = document.location.hash;
    if(page){
        page = page.substring(1);
    }else{
        page = "page_index";
    }
    moveToPage(page);
}

function moveToPage(page){
    $(".nav-item.active").removeClass("active");
    $('a[href="#'+page+'"]').first().closest('.nav-item').addClass("active");
    page = page.split("?")[0];
    $("#page_content").load("/site/pages/" + page + ".html");
}

function _ajax_request(url, data, callback, type, method) {
    if (jQuery.isFunction(data)) {
        callback = data;
        data = {};
    }
    return jQuery.ajax({
        type: method,
        url: url,
        data: data,
        success: callback,
        dataType: type
        });
}

jQuery.extend({
    put: function(url, data, callback, type) {
        return _ajax_request(url, data, callback, type, 'PUT');
    },
    delete: function(url, data, callback, type) {
        return _ajax_request(url, data, callback, type, 'DELETE');
    }
});