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
