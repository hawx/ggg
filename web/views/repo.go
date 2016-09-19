package views

const repo = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>{{.Name}}</title>
    <link rel="stylesheet" href="/assets/styles.css"></link>
  </head>
  <body>
    {{ template "header" . }}

    <div class="container">
      <div class="repo {{if .IsPrivate}}private{{end}}">
        <h1><a href="/{{.Name}}">{{.Name}}</a></h1>
        {{if .Web}}&rarr; <a href="{{.Web}}">{{.Web}}</a>{{end}}
        <p>{{.Description}}</p>
        {{ if $.LoggedIn }}
        <div class="buttons">
          <a href="/{{.Name}}/edit">edit</a>
          <a href="/{{.Name}}/delete">delete</a>
        </div>
        {{ end }}

        <pre class="clone"><code>git clone {{.Url}}/{{.CloneUrl}}</code></pre>
      </div>
    </div>

    <aside>
      <figure>
        {{if .IsEmpty}}
        <figcaption>
          <h3>Empty</h3>
        </figcaption>

        <article>
          <p>Maybe try pushing...</p>
        </article>
        {{else}}
        <figcaption>
          <h3>{{.ReadmeName}}</h3>
        </figcaption>

        <article>
          {{.ReadmeContents}}
        </article>
        {{end}}
      </figure>
    </aside>

  </body>
</html>`
