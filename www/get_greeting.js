
window.onbeforeunload = function () {
    console.log("bye111111")
}

function greeting() {

    var user = $('#uname0').val() === "" ? "null" : $('#uname0').val();
    var pwd = $('#pwd0').val() === "" ? "null" : $('#pwd0').val();

    $.ajax({
        url: 'http://192.168.76.37:1323/api/v0/greeting',
        username: user,
        password: pwd,
        type: 'GET',
        contentType: "application/json",
        data: '',
        dataType: 'json',
        crossDomain: true,
        success: function (data) {
            console.log(data);            
        },
        error: function (jqXHR, textStatus, errorThrown) {
            console.log(jqXHR.responseText);
        },
        complete: function () {
            // Handle the complete event
            // alert("ajax completed");
        }
        // beforeSend: function (xhr) {
        //     xhr.setRequestHeader ("Authorization", "Basic " + btoa("user" + ":" + "user"));
        // },
    });
}