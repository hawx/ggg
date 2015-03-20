package assets

const core = `
function post(data) {
    return new Promise(function(resolve, reject) {
        var req = new XMLHttpRequest();
        req.open("POST", "/-/sign-in");
        req.setRequestHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8");
        req.onload = function(e) {
            if (req.status == 200) {
                resolve(req.response);
            } else {
                reject(Error(req.statusText));
            }
        };

        req.onerror = function(e) {
            reject(Error("network error"));
        };

        req.send(data);
    });
}

var button = document.getElementById("browserid");
if (button != void 0) {
    button.addEventListener("click", function() {
        navigator.id.get(function(assertion) {
            if (assertion !== null) {
                post("assertion=" + assertion)
                    .then(function() {
                        window.location.reload();
                    }, function(err) {
                        alert("sign-in failed: " + err);
                    });
            }
        });
    });
}

var filter = document.getElementById("filter");
var list = document.querySelector(".repos");
filter.addEventListener("keyup", function() {
    var value = filter.value.toUpperCase();
    var items = list.getElementsByTagName("li");

    for (var i = 0; i < items.length; i++) {
        var item = items[i];
        var name = item.querySelector("h1 a");

        if (name.innerHTML.toUpperCase().indexOf(value) != -1) {
            item.style.display = "list-item";
        } else {
            item.style.display = "none";
        }
    }
});
`
