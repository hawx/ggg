package parts

const Files = `<figure>
  <figcaption>
    <h3 class="path">
      <a href="/{{.Name}}">{{.Name}}</a>
      <span>/</span>

      {{ range .DirParts }}
      <a href="/{{$.Name}}/tree{{.Path}}">{{.Name}}</a>
      <span>/</span>
      {{ end }}
    </h3>
  </figcaption>

  <ul class="files">

    {{ if ne .ParentDir "" }}
    {{ if eq .ParentDir "." }}
    <li>
      <i class="fa" aria-hidden="true"></i>
      <a href="/{{$.Name}}">..</a>
    </li>
    {{ else }}
    <li>
      <i class="fa" aria-hidden="true"></i>
      <a href="/{{$.Name}}/tree/{{.ParentDir}}">..</a>
    </li>
    {{ end }}
    {{ end }}

    {{ range .Files }}

    {{ if .IsDir }}
    <li>
      <i class="fa fa-folder" aria-hidden="true"></i>
      <a href="/{{$.Name}}/tree/{{.Path}}">{{.Name}}</a>
    </li>
    {{ else if .IsSubmodule }}
    <li>{{.Name}}</li>
    {{ else }}
    <li>
      <i class="fa fa-file-text-o" aria-hidden="true"></i>
      <a href="/{{$.Name}}/blob/{{.Path}}">{{.Name}}</a>
    </li>
    {{ end }}

    {{ end }}
  </ul>
</figure>`
