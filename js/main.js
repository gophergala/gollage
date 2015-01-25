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
$(function () {
  $('input[type=file]').on('change', prepareUpload);
  $('.image-form').on('submit', uploadFiles);
});

// Grab the files and set them to our variable
function prepareUpload(event)
{
  // Only get one file
  file = event.target.files[0];
}

// Catch the form submit and upload the files
function uploadFiles(event) {
  event.stopPropagation(); // Stop stuff happening
  event.preventDefault(); // Totally stop stuff happening

  // START A LOADING SPINNER HERE

  var url = $(event).attr('action');
  // Create a formdata object and add the files
  var data = new FormData();
  data.append("file", file);

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
