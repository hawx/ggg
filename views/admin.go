package views

const admin = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="/assets/styles.css"></link>
  </head>
  <body>
    <header>
      <h1>{{.Title}}</h1>
      <div>
        <a href="/create">create</a>
        <a href="/sign-out">sign-out</a>
      </div>
    </header>

    <ul class="repos">
      {{range .Repos}}
        <li class="repo {{if .IsPrivate}}private{{end}}">
          <div>
            <h1><a href="/{{.Name}}">{{.Name}}</a></h1>
            <p>{{.Description}}</p>
          </div>
          <div>
            <p class="meta updated">Updated <time datetime="{{.LastUpdate}}">3 days ago</time></p>
            <ul class="meta tags">
              {{range .TagsList}}
                <li><a href="#tagged-{{.}}">{{.}}</a></li>
              {{end}}
            </ul>
            <div class="buttons">
              <a href="/edit/{{.Name}}">edit</a>
              <a href="/delete/{{.Name}}">delete</a>
            </div>
          </div>
        </li>
      {{end}}
    </ul>
  </body>
</html>`
