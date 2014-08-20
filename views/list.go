package views

const list = `<!DOCTYPE html>
<html>
  <head>
    <title>{{Title}}</title>
    <link rel="stylesheet" href="/assets/styles.css"></link>
  </head>
  <body>
    <header>
      <a id="browserid" href="#" title="Sign-in with Persona">sign-in</a>
      <h1>{{Title}}</h1>
    </header>

    <p>{{Description}}</p>

    <ul class="repos">
      {{#Repos}}
        <li class="repo">
          {{#Web}}
            <h2 class="name"><a href="{{Web}}">{{Name}}</a></h2>
          {{/Web}}
          {{^Web}}
            <h2 class="name">{{Name}}</h2>
          {{/Web}}
          <p class="description">{{Description}}</p>
          <pre><code data-lang="shell">$ git clone {{Url}}/{{Name}}.git</code></pre>
        </li>
      {{/Repos}}
    </ul>

    <script src="http://code.jquery.com/jquery-2.1.1.min.js"></script>
    <script src="https://login.persona.org/include.js"></script>
    <script src="/assets/core.js"></script>
  </body>
</html>`
