package views

const list = `<!DOCTYPE html>
<html>
  <head>
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="/assets/styles.css"></link>
  </head>
  <body>
    <header>
      <h1>{{.Title}}</h1>
      <div>
        <a id="browserid" href="#" title="Sign-in with Persona">sign-in</a>
      </div>
    </header>

    <ul class="repos">
      {{range .Repos}}
        <li class="repo">
          <div>
            <h1><a href="/{{.Name}}">{{.Name}}</a></h1>
            <p>{{.Description}}</p>
          </div>
          <div>
            <p class="meta updated">Updated <time datetime="2014-...">3 days ago</time></p>
            <ul class="meta tags">
              {{range .TagsList}}
                <li><a href="#tagged-{{.}}">{{.}}</a></li>
              {{end}}
            </ul>
          </div>
        </li>
      {{end}}
    </ul>

    <footer>
      <p><a href="https://github.com/hawx/ggg">ggg</a> is a git repo lister.</p>
    </footer>

    <script src="http://code.jquery.com/jquery-2.1.1.min.js"></script>
    <script src="https://login.persona.org/include.js"></script>
    <script src="/assets/core.js"></script>
  </body>
</html>`
