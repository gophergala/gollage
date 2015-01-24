// Websocket object
var conn;

$(function () {
  initWebsockets();
});

function initWebsockets() {
  // Establish a WebSocket connection
  if (window["WebSocket"]) {
      conn = new WebSocket("ws://" + host + "/ws");
      conn.onerror = function(evt) {
        console.log(evt);
      }
      conn.onclose = function(evt) {
        console.log(evt);
      }
      conn.onmessage = function(evt) { // Message received. evt.data is something
        // Parse the JSON out of the data
        var data = JSON.parse(evt.data);
      }
  } else {
      // Your browser does not support WebSockets
  }
}

// Stolen from http://abandon.ie/notebook/simple-file-uploads-using-jquery-ajax
// ---- Everything below this line is AJAX file upload shit ----

// Variable to store your file
var file;

// Add events
$('input[type=file]').on('change', prepareUpload);

// Grab the files and set them to our variable
function prepareUpload(event)
{
  file = event.target.files[0];
}

$('.image-form').on('submit', uploadFiles);

// Catch the form submit and upload the files
function uploadFiles(event) {
  event.stopPropagation(); // Stop stuff happening
  event.preventDefault(); // Totally stop stuff happening

  // START A LOADING SPINNER HERE

  var url = $(event).attr('action');
  // Create a formdata object and add the files
  var data = new FormData();
  data.file = file;

  $.ajax({
    url: url,
    type: 'POST',
    data: data,
    cache: false,
    dataType: 'json',
    processData: false, // Don't process the files
    contentType: false, // Set content type to false as jQuery will tell the server its a query string request
    success: function(data, textStatus, jqXHR)
  {
    if(typeof data.error === 'undefined')
  {
    // Success so call function to process the form
    submitForm(event, data);
  }
    else
  {
    // Handle errors here
    console.log('ERRORS: ' + data.error);
  }
  },
    error: function(jqXHR, textStatus, errorThrown)
    {
      // Handle errors here
      console.log('ERRORS: ' + textStatus);
      // STOP LOADING SPINNER
    }
  });
}

function submitForm(event, data)
{
  // Create a jQuery object from the form
    $form = $(event.target);

    // Serialize the form data
    var formData = $form.serialize();

    // You should sterilise the file names
    formData = formData + '&filenames[]=' + data.file;

    $.ajax({
        url: 'submit.php',
        type: 'POST',
        data: formData,
        cache: false,
        dataType: 'json',
        success: function(data, textStatus, jqXHR)
        {
            if(typeof data.error === 'undefined')
            {
                // Success so call function to process the form
                console.log('SUCCESS: ' + data.success);
            }
            else
            {
                // Handle errors here
                console.log('ERRORS: ' + data.error);
            }
        },
        error: function(jqXHR, textStatus, errorThrown)
        {
            // Handle errors here
            console.log('ERRORS: ' + textStatus);
        },
        complete: function()
        {
            // STOP LOADING SPINNER
        }
    });
}
