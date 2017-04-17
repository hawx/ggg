package views

const create = `<!DOCTYPE html>
<html>
  <head>
    {{ template "meta" "create" }}
  </head>
  <body>
    {{ template "header" . }}

    <div class="container">
      <form method="POST" action="/-/create">
        <label for="name">Name</label>
        <input name="name" id="name" type="text" />

        <label for="web">Web</label>
        <input name="web" id="web" type="text" />

        <label for="description">Description</label>
        <textarea name="description" id="description"></textarea>

        <label for="private">Private?</label>
        <input name="private" id="private" type="checkbox" value="private" />
        <br/>

        <input type="submit" value="Create" />
      </form>
    </div>
  </body>
</html>`
