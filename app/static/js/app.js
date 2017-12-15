var editor = CodeMirror.fromTextArea(document.getElementById("code"), {
    lineNumbers: true,
    mode: "yaml",
    gutters: ["CodeMirror-lint-markers"],
    lint: true,
    theme: "liquibyte"
});
var btn = document.getElementById("submit");
var result = document.getElementById("result-container");
var clearBtn = document.getElementById("reset");

clearBtn.addEventListener('click', function() {
    editor.setValue("")
});
btn.addEventListener('click', function() {
    // console.log(editor.getValue());
    while (result.firstChild) {
        result.removeChild(result.firstChild);
    }
    if (editor.getValue() != "") {
        var opts = {
            method: 'POST',
            body: editor.getValue(),
            headers: {"Content-Type": "text/plain; charset=utf-8"}
        };
        fetch('/api/val', opts).then(function (response) {
            return response.json();
        }).then(function (body) {
            result.style.display = "block";
            var node = document.createElement("pre");
            var textnode = document.createTextNode(JSON.stringify(body, null, 2));
            node.appendChild(textnode);
            result.appendChild(node);
        });
    }
});


// var request = new XMLHttpRequest();

// request.onreadystatechange = function() {
//     if(request.readyState === 4) {
//         if(request.status === 200) {
//             //  result = request.responseText;
//             var node = document.createElement("pre");
//             var textnode = document.createTextNode(request.responseText);
//             node.appendChild(textnode);
//             result.appendChild(node);
//             request = new XMLHttpRequest();
//             request.open('POST', "/api/val");
//             result.style.display = "block";
//         } else {
//             result = "server unreachable."
//             request = new XMLHttpRequest();
//             request.open('POST', "/api/val");
//             result.style.display = "block";
//         }
//     }
// }
// request.open('POST', "/api/val");

// btn.addEventListener('click', function() {
//     console.log(editor.getValue());
//     request.send();
// });