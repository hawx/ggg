package views

const edit = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>edit</title>
    <link rel="stylesheet" href="/assets/styles.css"></link>
  </head>
  <body>
    <header>
      <div class="container">
        <h1>edit</h1>
        <a href="/">home</a>
        <a href="/sign-out">sign-out</a>
      </div>
    </header>

    <div class="container">
      <form method="POST" action="/{{.Name}}/edit">
        <label for="name">Name</label>
        <input name="name" id="name" type="text" value="{{.Name}}" disabled="disabled" />

        <label for="web">Web</label>
        <input name="web" id="web" type="text" value="{{.Web}}" />

        <label for="description">Description</label>
        <input name="description" id="description" type="text" value="{{.Description}}" />

        <label for="branch">Default Branch</label>
        <select name="branch">
          {{range .Branches}}
          <option value="{{.}}" {{if eq . $.Branch}}selected="selected"{{end}}>{{.}}</option>
          {{end}}
        </select>
        <br/>

        <label for="private">Private?</label>
        <input name="private" id="private" type="checkbox" value="private" {{if .IsPrivate}}checked="checked"{{end}} />
        <br/>

        <input type="submit" value="Save" />
      </form>
    </div>
  </body>
</html>`
