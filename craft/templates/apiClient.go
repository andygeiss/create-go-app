package templates

// APIClient ...
var APIClient = `// Code generated by {{ .Generator }} {{ .Version }} ({{ .Build }}); DO NOT EDIT

const {{ lc .Name }} = {
	{{ range $i, $name := .Services }}
	async {{ lc $name }}(params) {
		let config = {
			method: "POST",
			headers: {
				"accept": "application/json",
				"content-type": "application/json"
			},
			body: JSON.stringify(params)
		};
		let response = await fetch("/{{ lc $name }}", config);
		let data = await response.json();
		return data;
	}
	{{ end }}
};
`
