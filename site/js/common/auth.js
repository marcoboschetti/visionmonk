var totalFailsOnPost = 0;
var user_data;
var user_token;

// Global https redirection
if (location.protocol !== 'https:' && location.href.indexOf("localhost") == -1) {
    location.replace(`https:${location.href.substring(location.protocol.length)}`);
}

function loadUser() {
    // Get token, get user or create new one
    user_token = localStorage.getItem("vm_user_token");
    if(!user_token){
        console.log("No user token");
        window.location.replace("/site/pages/login.html");
        return;
    }

    // Add default auth header
    $.ajaxSetup({
        headers: { 'x-auth-token': user_token }
    });
    
    // Use cached user for render. Do not persist
    var tmp_user_data_json = localStorage.getItem("tmp_user_data_json");
    if (tmp_user_data_json) {
        tmp_user_data = JSON.parse(tmp_user_data_json);
        renderUserData(tmp_user_data);
    }

    // Force load user_data
    getUserData(function () { });

    $("#confirmLogoutBtn").click(function () {
        localStorage.removeItem("vm_user_token");
        localStorage.removeItem("tmp_user_data_json");
        window.location.replace("/site/pages/login.html");
        return;
    });
};

function getUserData(callback) {
    var wrappedCallback = wrapCallback(callback);

    // Cache
    if (user_data) {
        wrappedCallback(user_data);
        return;
    }  

    $.get("/api/p/user", function (data) {
        user_data = data.user
        localStorage.setItem("tmp_user_data_json", JSON.stringify(user_data));
        return wrappedCallback(user_data);
    }).fail(function (err) {
        console.log("Error using persisted user_token",user_token,":",err);
        localStorage.removeItem("vm_user_token");
        localStorage.removeItem("tmp_user_data_json");
        window.location.replace("/site/pages/login.html");
        return;
    });
   
}

function wrapCallback(callback) {
    return function (user) {
        renderUserData(user);
        callback(user);
    }
}

function renderUserData(user) {
    $("#nav_user_name").html(user.first_name + " " + user.last_name);
    if (user.shop)
        $("#nav_shop_name").html(user.shop.name)
}

function loginPostCredentials(email, password, invalidCredentialsCallback) {

    authStr = btoa(email + "&!&" + password)
    $.ajaxSetup({ headers: { 'x-auth-token-bearer': authStr } });

    $.get("/api/login", function (data) {
        console.log(data);
        localStorage.setItem("vm_user_token", data.secret_token);
        window.location.replace("/");
    }).fail(function (err) {
        if (err.status == 401) {
            invalidCredentialsCallback();
        }
    });;
}