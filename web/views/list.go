package views

const list = `<!DOCTYPE html>
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
        <a href="/sign-in" title="Sign-in with Persona">sign-in</a>
      </div>
    </header>

    <div class="filter">
      <div class="container">
        <input id="filter" type="search" placeholder="Filter..." tabindex="1" />
      </div>
    </div>

    <div class="container">
      <ul class="repos">
        {{range .Repos}}
        <li class="repo">
          <h1><a href="/{{.Name}}">{{.Name}}</a></h1>
          <p>{{.Description}}</p>
        </li>
        {{end}}
      </ul>
    </div>

    <script src="/assets/core.js"></script>
  </body>
</html>`
