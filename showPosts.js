//Recover all Posts and print to console.log


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
