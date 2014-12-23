package views

const admin = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="/assets/styles.css"></link>
  </head>
  <body>
    <header>
      <div class="container">
        <h1><a href="/">{{.Title}}</a></h1>
        <a href="/create">create</a>
        <a href="/sign-out">sign-out</a>
      </div>
    </header>

    <div class="container">
      <ul class="repos">
        {{range .Repos}}
        <li class="repo {{if .IsPrivate}}private{{end}}">
          <h1><a href="/{{.Name}}">{{.Name}}</a></h1>
          <p>{{.Description}}</p>
          <div class="buttons">
            <a href="/edit/{{.Name}}">edit</a>
            <a href="/delete/{{.Name}}">delete</a>
          </div>
        </li>
        {{end}}
      </ul>
    </div>
  </body>
</html>`
