package views

const file = `<figure>
  <figcaption>
    <h3>
      <a href="/{{.Name}}">{{.Name}}</a>
      <span>/</span>

      {{ range .DirParts }}
      <a href="/{{$.Name}}/tree{{.Path}}">{{.Name}}</a>
      <span>/</span>
      {{ end }}
      <span>{{.FileName}}</span>
    </h3>
  </figcaption>

  <article>
    {{.FileContents}}
  </article>
</figure>`
