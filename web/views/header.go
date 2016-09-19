package views

const header = `<header>
  <h1><a href="/">{{.Title}}</a></h1>
  {{ if .LoggedIn }}
  <a href="/-/create">create</a>
  <a href="/-/sign-out">sign-out</a>
  {{ else }}
  <a href="/-/sign-in" title="Sign-in">sign-in</a>
  {{ end }}
</header> `
