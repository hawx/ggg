package parts

const Description = `
<div class="repo {{if .IsPrivate}}private{{end}}">
  <div class="repo-header">
    <h1><a href="/{{.Name}}">{{.Name}}</a></h1>

    {{ if $.LoggedIn }}
    <div class="buttons">
      <a href="/{{.Name}}/edit">edit</a>
      <a href="/{{.Name}}/delete">delete</a>
    </div>
    {{ end }}
  </div>

  {{if .Web}}
  <i class="fa fa-external-link" aria-hidden="true"></i>
  <a class="external-link" href="{{.Web}}">{{.Web}}</a>
  {{end}}
  <p>{{.Description}}</p>

  <pre class="clone"><code>git clone {{.Url}}/{{.CloneUrl}}</code></pre>
</div>
`
