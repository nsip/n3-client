// https://www.html5rocks.com/en/tutorials/file/dndfiles/

// Check for the various File API support.
// if (window.File && window.FileReader && window.FileList && window.Blob) {
//     alert('Great success! All the File APIs are supported.')
// } else {
//     alert('The File APIs are not fully supported in this browser.');
// }

window.onload = function () {
    // document.getElementById('fileupload').addEventListener('change', handleFileSelect, false);
    $('#fileupload').change(handleFileSelect);
}

var filename = "";
var F;

function handleFileSelect(evt) {
    var files = evt.target.files; // FileList object    
    // files is a FileList of File objects. List some properties.

    var output = [];
    for (var i = 0, f; f = files[i]; i++) {
        output.push('<li><strong>', escape(f.name), '</strong> (', f.type || 'n/a', ') - ',
            f.size, ' bytes, last modified: ',
            f.lastModifiedDate ? f.lastModifiedDate.toLocaleDateString() : 'n/a',
            '</li>');

    }
    // document.getElementById('file').innerHTML = '<ul>' + output.join('') + '</ul>';
    $('#file').html('<ul>' + output.join('') + '</ul>');

    F = files[0];

    filename = $('#fileupload').val();
    if (!filename || !(filename.endsWith('.xml') || filename.endsWith('.json'))) {
        $('#pub').prop('disabled', true);
    } else {
        $('#pub').prop('disabled', false);
    }
}

function pub() {

    var user = $('#uname1').val() === "" ? "null" : $('#uname1').val();
    var pwd = $('#pwd1').val() === "" ? "null" : $('#pwd1').val();
    if ($('#dflt1').val() === "") {
        alert('input default data object root name')
        return
    }
    var dfltRoot = $('#dflt1').val()

    var formdata = {
        "ID": "e56dc44f-9080-41a5-82d0-a323472016c0",
        "actId": "actId",
        "proId": "proId",
        "channel_price": "channel_price",
        "sale_num": "sale_num",
        "restrict_num": "restrict_num"
    };

    var reader = new FileReader();
    reader.readAsText(F, 'UTF-8');
    reader.onload = shipOff;
    function shipOff(event) {
        var mydata = event.target.result;
        console.log(mydata);

        $.ajax({
            url: 'http://192.168.76.37:1323/api/v0/pub?dfltRoot=' + dfltRoot,
            username: user,
            password: pwd,
            type: 'POST',
            contentType: "application/json; charset=utf-8",
            data: mydata, //JSON.stringify(formdata),
            cache: false,
            dataType: 'json',
            traditional: true,
            crossDomain: true,
            success: function (data) {
                console.log(data);
            },
            error: function (jqXHR, textStatus, errorThrown) {
                // alert('An error occurred... Look at the console (F12 or Ctrl+Shift+I, Console tab) for more information!');
                // $('#result').html('<p>status code: ' + jqXHR.status + '</p><p>errorThrown: ' + errorThrown + '</p><p>jqXHR.responseText:</p><div>' + jqXHR.responseText + '</div>');
                // console.log('jqXHR:');
                console.log(jqXHR.responseText);
                // console.log('textStatus:');
                // console.log(textStatus);
                // console.log('errorThrown:');
                // console.log(errorThrown);
            },
            complete: function () {
                // Handle the complete event
                // alert("ajax completed");
            }
        });
    }
}