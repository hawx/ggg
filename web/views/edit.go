package views

const edit = `<!DOCTYPE html>
<html>
  <head>
    {{ template "meta" "edit" }}
  </head>
  <body>
    {{ template "header" . }}

    <div class="container">
      <div class="repo">
        <h1><a href="/{{.Name}}">{{.Name}}</a></h1>
      </div>

      <form method="POST" action="/{{.Name}}/edit">
        <input name="name" id="name" type="hidden" value="{{.Name}}" />

        <label for="web">Web</label>
        <input name="web" id="web" type="text" value="{{.Web}}" />

        <label for="description">Description</label>
        <textarea name="description" id="description">{{.Description}}</textarea>

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
