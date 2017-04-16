package assets

const styles = `
html, body {
    margin: 0;
    padding: 0;
}

body {
    font: 100%/1.3 Verdana, sans-serif;
    overflow-y: scroll;
}

.container {
    max-width: 40rem;
    margin: 1rem;
}

.container.is-wide {
    max-width: 60rem;
}

aside {
    max-width: 40rem;
    margin: 1rem 1.5rem;
}

.container:before, .container:after {
    clear: both;
    content: " ";
    display: table;
}

a {
    color: hsl(220, 51%, 44%);
}

a:hover {
    color: hsl(208, 56%, 38%);
}

a.external-link {
    font-style: italic;
}

header {
    margin: 0 1rem;
    font-size: 1rem;
    border-bottom: 1px solid #ddd;
}

header h1, header > a {
    margin: 0;
    padding: 1.3rem;
    height: 1.3rem;
    line-height: 1.3rem;
    display: inline-block;
}

header h1 {
    font-size: 1.5rem;
    padding-left: 0;
    margin-left: .5rem;
    font-weight: bold;
    align-self: flex-start;
}

header h1 a {
    color: #000;
    text-decoration: none;
}

.repos {
    width: auto;
    list-style: none;
    padding: 0;
    margin: 0;
}

.repos .repo {
    border-bottom: 1px dotted #ddd;
    width: auto;
}

.repos .repo:last-child {
    border-bottom: none;
}

.repo {
    margin: 1.3rem 0;
    padding: 0 .5rem;
    position: relative;
}

.repo h1 {
    font-size: 1.2rem;
    margin: 0;
}

.repo h1 a {
    text-decoration: none;
}

.repo.private a {
    color: hsl(20, 51%, 44%);
}

.repo-header {
    display: flex;
    justify-content: space-between;
    margin: 1rem 0;
    line-height: 1rem;
}

.clone {
    margin: 1.3rem 0;
    background: #efefef;
    padding: 1rem;
}

.filter {
    padding: 1rem 0px;
    background: #fefefe;
    border-bottom: 1px solid #eee;
}

.filter input[type=search] {
    display: block;
    border: none;
    width: 100%;
    padding: 0 0.5rem;
    font: 1rem Verdana, Geneva, sans-serif;
    margin: 0;
    background: none repeat scroll 0% 0% transparent;
}

hr {
    margin: 1.3rem auto;
    border: none;
    height: 0;
    max-width: 10rem;
    width: 66%;
    border-bottom: 1px solid #ccc;
}

figure {
    border: 1px solid #bbb;
    margin: 1rem .5rem;
    padding: 0;
}

figure figcaption {
    border-bottom: 1px solid #ddd;
    padding: 0.65rem 0.65rem;
    margin: 0 0.65rem;
}

figure figcaption h3 {
    margin: 0;
    font-size: .75rem;
}

figure article {
    padding: 0.65rem 1.3rem;
}

figure article pre.full {
    padding: 0;
    background: transparent;
}

form {
    margin: 1.3rem 0.5rem;
}

input[type=text] {
    display: block;
    border: none;
    border-bottom: 1px dotted #aaa;
    margin: 0.3rem 0 1.3rem;
    width: 100%;
}

textarea {
    display: block;
    border: 1px dotted #aaa;
    margin: 0.3rem 0 1.3rem;
    width: 100%;
    max-width: 100%;
    min-width: 100%;
}

input[type=submit] {
    margin: 1rem 0;
}

label {
    color: #aaa;
    font-style: italic;
    margin: 0.3rem 0;
}

figure {
    background: white;
    color: black;
}

article {
    font: 13px/1.3 Menlo, monospace;
}

article h1:first-child {
    margin-top: 1rem;
}

article h1 {
    font-size: 16px;
    margin: 2rem 0 1rem;
}

article h2 {
    font-size: 14px;
    margin: 1.5rem 0 0.5rem;
}

article h3 {
    font-size: 13px;
    font-weight: normal;
    text-decoration: underline;
    margin: 1.5rem 0 0.5rem;
}

article ul {
    padding-left: 1.3em;
}

article p {
    margin: 1.3em 0;
}

article code {
    font-family: Menlo, monospace;
    color: rgb(90, 10, 20);
}

article pre > code {
    padding: 0;
    background: none;
    color: rgb(90, 10, 20);
}

article pre {
    padding: .7em 1.3em;
    font-family: Menlo, monospace;
    white-space: pre-wrap;
    color: rgb(90, 10, 20);
}

article pre.content {
    color: black;
    padding: 0;
}

article blockquote {
    margin-left: 0;
    padding-left: 1.3em;
    border-left: 1px dotted #aaa;
    color: #666;
}

article hr {
    border: none;
    height: 1px;
    width: 7rem;
    margin: 2.5rem 0 0.5rem;
    background: #aaa;
}

article table {
    border-spacing: collapse;
    border-collapse: collapse;
}

article th, article td {
    padding: 0.3rem 0.5rem;
    margin: 0;
}

article th {
    font-weight: bold;
    text-align: left;
    border-bottom: 1px solid #eee;
}

.files {
    list-style: none;
    padding: 0;
    margin: 0;
}

.files li {
    margin: 0 .65rem;
    padding: .65rem;
    border-bottom: 1px dotted #ddd;
    font-size: .75rem;
}

.files li:last-child {
    border-bottom: none;
}

.files .fa {
    margin-right: .3rem;
    height: 14px;
    width: 14px;
}
`
