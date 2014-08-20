package views

const create = `<!DOCTYPE html>
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

    <form method="POST" action="/create">
      <label for="name">Name</label>
      <input name="name" id="name" type="text" />

      <label for="web">Web</label>
      <input name="web" id="web" type="text" />

      <label for="description">Description</label>
      <input name="description" id="description" type="text" />

      <label for="private">Private?</label>
      <input name="private" id="private" type="checkbox" value="private" />
      <br/>

      <input type="submit" value="Create" />
    </form>
  </body>
</html>`
