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
