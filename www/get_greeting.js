
window.onbeforeunload = function () {
    confirm('leave')
}

function greeting() {

    var user = $('#uname0').val() === "" ? 'null' : $('#uname0').val();
    var pwd = $('#pwd0').val() === "" ? 'null' : $('#pwd0').val();

    var ip = location.host;
    $.ajax({
        url: 'http://' + ip + '/api/v0.1.0/greeting',
        username: user,
        password: pwd,
        type: 'GET',
        contentType: 'application/json',
        data: '',
        dataType: 'json',
        crossDomain: true,
        success: function (data) {
            console.log(data);
            alert(data);
        },
        error: function (jqXHR, textStatus, errorThrown) {
            console.log(jqXHR.responseText);
            alert(jqXHR.responseText);
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