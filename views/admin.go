package views

const admin = `<!DOCTYPE html>
<html>
  <head>
    <title>{{Title}}</title>
    <link rel="stylesheet" href="/assets/styles.css"></link>
  </head>
  <body>
    <header>
      <a href="/sign-out">sign-out</a>
      <a href="/create">create</a>
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
          <a href="/edit/{{Name}}">edit</a>
          <a href="/delete/{{Name}}">delete</a>
        </li>
      {{/Repos}}
    </ul>
  </body>
</html>`
