$(document).ready(function () {
    getUserData(function(user_data){
        console.log("render shop data",user_data.shop)
        $("#countShopUsers").html(user_data.shop.users_count);
        $("#countClients").html(user_data.shop.clients_count);
    });
});
