package views

const repo = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <title>{{.Name}}</title>
    <link rel="stylesheet" href="/assets/styles.css"></link>
  </head>
  <body class="single">
    <div>
      <div class="repo {{if .IsPrivate}}private{{end}}">
        <h1><a href="/">↺</a>&ensp;<a href="/{{.Name}}">{{.Name}}</a></h1>
        {{if .Web}}&rarr; <a href="{{.Web}}">{{.Web}}</a>{{end}}
        <p>{{.Description}}</p>
      </div>

      <hr/>

      <div class="clone">
        <span>git clone</span>
        <code>{{.Url}}/{{.CloneUrl}}</code>
      </div>
    </div>

    <aside>
      <figure>
        <figcaption>
          <h3>{{.ReadmeName}}</h3>
        </figcaption>

        <article>
          {{.ReadmeContents}}
        </article>
      </figure>
    </aside>

  </body>
</html>`
