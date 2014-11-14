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
      <header class="repo">
        <div>
          <h1><a href="/">↺</a>&ensp;<a href="/{{.Name}}">{{.Name}}</a></h1>
          <p>{{.Description}}</p>
        </div>
        <div>
          <p class="meta updated">Updated <time datetime="2014-08-05T10:11:11Z">3 days ago</time></p>
        </div>
      </header>

      <hr/>

      <div class="clone">
        <span>git clone</span>
        <code>https://git.hawx.me/{{.CloneUrl}}</code>
      </div>

    </div>

    <aside>
      <figure>
        <figcaption>
          <h3>README.md</h3>
        </figcaption>

        <article>
          {{.Readme}}
        </article>
      </figure>
    </aside>

  </body>
</html>`
