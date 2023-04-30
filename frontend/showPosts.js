//Recover all Posts and print to console.log

let ws = new WebSocket("ws://localhost:5000/ws");
ws.onopen = function() {
    console.log("Connected to Websocket");
};
ws.onmessage = function(event) {
    console.log(event.data);
};

ws.onerror = function(event) {
    console.log("[ERROR websocket]: " + event.data);
};
fetch("http://localhost:5000/posts", {
    method: "GET",
    headers: {
        "Content-Type": "application/json"
    }
}).then(function(response) {
    return response.json()
}).then(function(json) {
    console.log(json)
});
