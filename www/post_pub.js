// https://www.html5rocks.com/en/tutorials/file/dndfiles/

// Check for the various File API support.
// if (window.File && window.FileReader && window.FileList && window.Blob) {
//     alert('Great success! All the File APIs are supported.')
// } else {
//     alert('The File APIs are not fully supported in this browser.');
// }

window.onload = function () {

    var user, pwd, formdata, approval, info, filename, root;

    // document.getElementById('selectfile').addEventListener('change', handleFileSelect, false);
    $('#selectfile').change(function (evt) {
        var files = evt.target.files; // files is a FileList of File objects. List some properties.

        var output = [];
        for (var i = 0, f; f = files[i]; i++) {
            output.push('<li><strong>', escape(f.name), '</strong> (', f.type || 'n/a', ') - ',
                f.size, ' bytes, last modified: ',
                f.lastModifiedDate ? f.lastModifiedDate.toLocaleDateString() : 'n/a',
                '</li>');
        }
        $('#info').html('<ul>' + output.join('') + '</ul>');

        filename = $('#selectfile').val();
        var nopub = (!filename || !(filename.endsWith('.xml') || filename.endsWith('.json')) || $('#root').val() === "")
        $('#pub').prop('disabled', nopub);
    });

    $('#root').on('input', function (e) {
        filename = $('#selectfile').val();
        var nopub = (!filename || !(filename.endsWith('.xml') || filename.endsWith('.json')) || $('#root').val() === "")
        $('#pub').prop('disabled', nopub);
    });

    $("#uploadform").on('submit', function (e) {
        e.preventDefault();  // avoid to execute the actual submit of the form

        user = $('#uname').val() === "" ? "null" : $('#uname').val();
        pwd = $('#pwd').val() === "" ? "null" : $('#pwd').val();
        root = $('#root').val();

        formdata = new FormData();
        formdata.append('username', user);
        formdata.append('password', pwd);
        formdata.append('root', root);
        jQuery.each(jQuery('#selectfile')[0].files, function (i, file) {
            formdata.append('file', file);
        });

        approval = confirm('Upload [' + filename + '] as object [' + root + '] to NIAS3?');
        if (!approval) {
            info = 'upload canceled';
            return;
        }
    });

    $("#uploadform").submit(function () { // intercepts the submit event

        if (!approval) {
            $('#info').html('<ul> ' + info + ' </ul>');
            return;
        }

        // ***
        $('#selectfile').prop('disabled', true);
        $('#pub').prop('disabled', true);
        $('#uploadwaiting').show();
        // ***

        var ip = location.host;
        $.ajax({ // make an AJAX request
            type: "POST",
            method: "POST",
            url: 'http://' + ip + '/file/v0.1.0/upload',
            username: user,
            password: pwd,
            contentType: false,
            data: formdata, // $("#uploadform").serialize(), // serializes the form's elements
            cache: false,
            processData: false,
            crossDomain: true,
            success: function (data) {
                console.log(data);
                $('#info').html('<ul>' + data + '</ul>');
                $('#root').val('');
                $('#selectfile').val('');
            },
            error: function (jqXHR, textStatus, errorThrown) {
                console.log(jqXHR.responseText);
                $('#info').html('<ul>' + jqXHR.responseText + '</ul>');
                $('#pub').prop('disabled', false);
            },
            complete: function () {
                $('#selectfile').prop('disabled', false);
                $('#uploadwaiting').hide();
            }
        });
    });
}

// function pub() {

//     var user = $('#uname1').val() === "" ? "null" : $('#uname1').val();
//     var pwd = $('#pwd1').val() === "" ? "null" : $('#pwd1').val();
//     if ($('#dflt1').val() === "") {
//         alert('input default data object root name')
//         return
//     }
//     var dfltRoot = $('#dflt1').val()

//     var formdata = {
//         "ID": "e56dc44f-9080-41a5-82d0-a323472016c0",
//         "actId": "actId",
//         "proId": "proId",
//         "channel_price": "channel_price",
//         "sale_num": "sale_num",
//         "restrict_num": "restrict_num"
//     };

//     var reader = new FileReader();
//     reader.readAsText(F, 'UTF-8');
//     reader.onload = shipOff;
//     function shipOff(event) {
//         var mydata = event.target.result;
//         console.log(mydata);

//         var ip = location.host;
//         $.ajax({
//             url: 'http://' + ip + '/api/v0/pub?dfltRoot=' + dfltRoot,
//             username: user,
//             password: pwd,
//             type: 'POST',
//             contentType: "application/json; charset=utf-8",
//             data: mydata, //JSON.stringify(formdata),
//             cache: false,
//             dataType: 'json',
//             traditional: true,
//             crossDomain: true,
//             success: function (data) {
//                 console.log(data);
//             },
//             error: function (jqXHR, textStatus, errorThrown) {
//                 // alert('An error occurred... Look at the console (F12 or Ctrl+Shift+I, Console tab) for more information!');
//                 // $('#result').html('<p>status code: ' + jqXHR.status + '</p><p>errorThrown: ' + errorThrown + '</p><p>jqXHR.responseText:</p><div>' + jqXHR.responseText + '</div>');
//                 // console.log('jqXHR:');
//                 console.log(jqXHR.responseText);
//                 // console.log('textStatus:');
//                 // console.log(textStatus);
//                 // console.log('errorThrown:');
//                 // console.log(errorThrown);
//             },
//             complete: function () {
//                 // Handle the complete event
//                 // alert("ajax completed");
//             }
//         });
//     }
// }