package parts

const File = `<figure>
  <figcaption>
    <h3>
      <span>{{.FileName}}</span>
    </h3>
  </figcaption>

  <article>
    <pre class="content"><code class="{{.FileLang}}">{{.FileContents}}</code></pre>
  </article>
</figure>`
