package parts

const File = `<figure>
  <figcaption>
    <h3>
      <span>{{.FileName}}</span>
    </h3>
  </figcaption>

  <article>
    <pre class="content">{{.FileContents}}</pre>
  </article>
</figure>`
