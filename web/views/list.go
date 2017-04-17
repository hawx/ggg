package views

const list = `<!DOCTYPE html>
<html>
  <head>
    {{ template "meta" .Title }}
  </head>
  <body>
    {{ template "header" . }}

    <div class="container">
      <ul class="repos">
        {{ range .Repos }}
        <li class="repo {{if .IsPrivate}}private{{end}}">
          <div class="repo-header">
            <h1><a href="/{{.Name}}">{{.Name}}</a></h1>

            {{ if $.LoggedIn }}
            <div class="buttons">
              <a href="/{{.Name}}/edit">edit</a>
              <a href="/{{.Name}}/delete">delete</a>
            </div>
            {{ end }}
          </div>

          <p>{{.Description}}</p>
        </li>
        {{ end }}
      </ul>
    </div>
  </body>
</html>`
