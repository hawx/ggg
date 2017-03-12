package views

const files = `<figure>
  <figcaption>
    <h3>Files</h3>
  </figcaption>

  <ul class="files">

    {{ if ne .ParentDir "" }}
    {{ if eq .ParentDir "." }}
    <li><a href="/{{$.Name}}">..</a></li>
    {{ else }}
    <li><a href="/{{$.Name}}/tree/{{.ParentDir}}">..</a></li>
    {{ end }}
    {{ end }}

    {{ range .Files }}

    {{ if .IsDir }}
    <li><a href="/{{$.Name}}/tree/{{.Path}}">{{.Name}}</a></li>
    {{ else if .IsSubmodule }}
    <li>{{.Name}}</li>
    {{ else }}
    <li><a href="/{{$.Name}}/blob/{{.Path}}">{{.Name}}</a></li>
    {{ end }}

    {{ end }}
  </ul>
</figure>`
