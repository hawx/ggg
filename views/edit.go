package views

const edit = `<!DOCTYPE html>
<html>
  <head>
    <title>git</title>
    <link rel="stylesheet" href="/assets/styles.css"></link>
  </head>
  <body>
    <header>
      <a href="/">home</a>
      <h1>git</h1>
    </header>

    <form method="POST" action="/edit/{{.Name}}">
      <label for="name">Name</label>
      <input name="name" id="name" type="text" value="{{.Name}}" disabled="disabled" />

      <label for="web">Web</label>
      <input name="web" id="web" type="text" value="{{.Web}}" />

      <label for="description">Description</label>
      <input name="description" id="description" type="text" value="{{.Description}}" />

      <label for="tags">Tags</label>
      <input name="tags" id="tags" type="text" value="{{.Tags}}" />

      <label for="private">Private?</label>
      <input name="private" id="private" type="checkbox" value="private" {{if .IsPrivate}}checked="checked"{{end}} />
      <br/>

      <input type="submit" value="Save" />
    </form>
  </body>
</html>`
