package views

const list = `<!DOCTYPE html>
<html>
  <head>
    {{ template "meta" .Title }}
  </head>
  <body>
    {{ template "header" . }}

    <div class="container filter">
      <input id="filter" type="search" placeholder="Filter..." tabindex="1" />
    </div>

    <div class="container">
      <ul class="repos">
        {{ range .Repos }}
        <li class="repo {{if .IsPrivate}}private{{end}}">
          <div class="repo-header">
            <h1><a href="/{{.Name}}">{{.Name}}</a></h1>
          </div>

          <p>{{.Description}}</p>
        </li>
        {{ end }}
      </ul>
    </div>

    <script src="/assets/filter.js"></script>
  </body>
</html>`
