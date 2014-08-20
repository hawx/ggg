package assets

const Styles = `
html, body {
  height: 100%;
  width: 100%;
  margin: 0;
  padding: 0;
}

body {
  font: 13px/1.3em Menlo, monospace;
  padding: 1em 2em;
  max-width: 44em;
  width: auto;
  height: auto;
}

h1 {
  font-size: 16px;
  margin: 2rem 0 1rem;
}

h2 {
  font-size: 14px;
  margin: 1.5rem 0 0.5rem;
}

h3 {
  font-size: 13px;
  font-weight: normal;
  text-decoration: underline;
  margin: 1.5rem 0 0.5rem;
}

ul {
  padding-left: 1.3em;
}

li.pad, p {
  margin: 1.3em 0;
}

code {
  font-family: Menlo, monospace;
  background: #eee;
}

pre > code {
  padding: 0;
  background: none;
}

pre {
  background: #eee;
  padding: .7em 1.3em;
  font-family: Menlo, monospace;
}

a {
  color: #761410;
}

a:hover {
  color: #cb231c;
}

#meta {
  margin-top: 2em;
  color: #666;
}

footer {
  margin-top: 2em;
}

footer hr {
  border: none;
  height: 1px;
  width: 7rem;
  margin: 2.5rem 0 0.5rem;
  background: #aaa;
}

.repos {
    list-style: none;
    margin-left: 0;
    padding-left: 0;
}

.repo {
  padding: 0 1.3rem;
  margin: 1.3rem 0;
  position: relative;
  border: 1px solid transparent;
}

.repo.private {
  border: 1px solid #aaa;
}

.name a {
    text-decoration: underline;
}

header > a {
  float: right;
  margin-top: 1px;
  margin-left: 1em;
}

input[type=text] {
  display: block;
  border: none;
  border-bottom: 1px dotted #aaa;
  margin: 0.3em 0 1.3em;
  width: 100%;
}

input[type=submit] {
  margin: 1rem 0;
}

label {
  color: #aaa;
  font-style: italic;
  margin: 0.3em 0;
}

.repo > .buttons {
  position: absolute;
  margin-left: 1em;
  top: 1.3rem;
  right: 1.3rem;
}

`
