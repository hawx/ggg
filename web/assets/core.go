package assets

const core = `
function post(data) {
    return new Promise(function(resolve, reject) {
        var req = new XMLHttpRequest();
        req.open("POST", "/sign-in");
        req.setRequestHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
        req.onload = function(e) {
            if (req.status == 200) {
                resolve(req.response);
            } else {
                reject(Error(req.statusText));
            }
        };

        req.onerror = function(e) {
            reject(Error("network error"));
        }

        req.send(data);
    });
}

var button = document.getElementById("browserid");
button.addEventListener("click", function() {
    navigator.id.get(function(assertion) {
        if (assertion !== null) {
            post("assertion=" + assertion)
                .then(function() {
                    window.location.reload();
                }, function(err) {
                    alert("sign-in failed: " + err)
                });
        }
    });
});
`
