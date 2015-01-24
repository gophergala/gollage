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
