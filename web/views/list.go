package views

const list = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="/assets/styles.css"></link>
  </head>
  <body>
    <header>
      <div class="container">
        <h1><a href="/">{{.Title}}</a></h1>
        <a id="browserid" href="#" title="Sign-in with Persona">sign-in</a>
      </div>
    </header>

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

    <script src="http://code.jquery.com/jquery-2.1.1.min.js"></script>
    <script src="https://login.persona.org/include.js"></script>
    <script src="/assets/core.js"></script>
  </body>
</html>`
