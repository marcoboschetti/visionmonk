
// Plugin-free Form validation
$.fn.isValid = function(){
    var validate = true;
    this.each(function(){
        if(this.checkValidity()==false){
            validate = false;
        }
    });
    return validate;
};

function getBase64(file, callback) {
    var reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = function () {
        callback(reader.result);
    };
    reader.onerror = function (error) {
      console.log('Error reading file: ', error);
    };
 }