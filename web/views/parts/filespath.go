package parts

const FilesPath = `<figure>
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
</figure>`
