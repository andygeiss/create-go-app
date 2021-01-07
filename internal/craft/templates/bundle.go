package templates

// BundleAppJs ...
var BundleAppJs = `
// View ...
class View extends Component {
    // Constructor ...
    constructor(viewModel) {
        super();
        this.viewModel = viewModel;
        // Add event listeners
        this.on("status done", (data) => {
            this.render();
        });
        // Add DOM event listeners
        document.querySelector("#btn_status").addEventListener("click", (evt) => {
            evt.preventDefault();
            this.viewModel.status({
                key: "value"
            });
        });
        // Initial rendering
        this.render();
    }
    // render ...
    render() {
        // Read the current state
        let obj = this.viewModel.getState("status result");
        if (typeof obj === "undefined") {
            obj = { text: "default here" };
        }
        // Modify the HTML output
        document.querySelector("#txt_status").value = obj.text;
    }
}

const view = new View(viewModel);`

// BundleAppCSS ...
var BundleAppCSS = `/* reset */
* {
    border: none;
    box-sizing: border-box;
    font: inherit;
    margin: 0;
    outline: none;
    padding: 0;
    text-decoration: none;
    transition: all .35s;
}

body {
    display: grid;
    font-size: 20pt;
    grid-template-columns: 1fr;
    padding: 1rem;
}

`

// BundleComponentJs ...
var BundleComponentJs = `// Code generated by {{ .Generator }} {{ .Version }} ({{ .Build }}); DO NOT EDIT

// Component introduces a state which can be accessed via a get/set.
// Additionally it provides event handling by receiving events via on and emitting events via emit.
class Component {

    // At first we construct an HTML element and set the state to an empty object.
    constructor() {
        this.state = {};
    }

    // emit dispatches a specific event with corresponding data.
    emit(event, data) {
        // We use window for dispatching events globally.
        // Thus, we don't need "bubbles" to propagate events up through the DOM.
        window.dispatchEvent(new CustomEvent(event, {
            detail: {
                output: data
            }
        }))
    }

    // getState reads a state value by a given key.
    getState(key) {
        return this.state[key];
    }

    // on adds a listener for a specific event.
    on(event, fn) {
        window.addEventListener(event, (e) => {
            // Only the object data (detail) is necessary for this kind of event.
            fn(e.detail.output);
        });
    }

    // setState writes a state key, value pair.
    setState(key, val) {
        this.state[key] = val;
    }
}

`

// BundleIndexHTML ...
var BundleIndexHTML = `<!DOCTYPE html>
<html lang="en" xmlns="http://www.w3.org/1999/html">

<head>
  <meta charset="utf-8" />
  <meta content="text/html; charset=utf-8" http-equiv="content-type" />
  <meta content="no-cache, no-store, must-revalidate" http-equiv="Cache-Control" />
  <meta content="no-cache" http-equiv="Pragma" />
  <meta content="0" http-equiv="Expires" />
  <meta name="viewport" content="width=device-width,initial-scale=1" />
  <link rel="stylesheet" href="bundle.css" />
  <title> {{ .Name }} </title>
</head>

<body>
  <input id="txt_status" type="text" />
  <button id="btn_status">Status</button>
  <script src="bundle.js"></script>
</body>

</html>`