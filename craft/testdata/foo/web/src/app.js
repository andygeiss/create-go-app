
// View ...
class View extends FlatElement {
    // constructor ...
    constructor(viewModel) {
        super();
        this.viewModel = viewModel;
        // add event listeners
        this.on("service done", (data) => {
            this.render();
        });
        // initial rendering
        this.render();
    }
    // render ...
    render() {
        // read the state
        let value = this.viewModel.getState("service result");
        if (typeof value === "undefined") {
            value = "default here";
        }
        // modify the HTML output
        document.querySelector("selector").value = value;
        // add event listeners
        document.querySelector("#btn_service").addEventListener("click", (evt) => {
            evt.preventDefault();
            // service should be replaced with the corresponding service API function.
            this.viewModel.service({
                key: "value"
            });
        });
    }
}

const view = new View(viewModel);