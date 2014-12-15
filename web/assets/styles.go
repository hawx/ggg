package assets

const styles = `
html, body {
    margin: 0;
    padding: 0;
}

body {
    font: 14px/1.3 Verdana, Geneva, sans-serif;
}

.container {
    max-width: 40em;
    margin: 0 auto;
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

header {
    margin: 0;
    background: #eee;
    font-size: 1em;
    border-bottom: 1px solid #ddd;
}

header > .container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
}

header h1, header > .container > a {
    margin: 0;
    padding: 1.3rem;
    height: 1.3rem;
    line-height: 1.3rem;
    display: inline-block;
}

header h1 {
    font-size: 1.5em;
    padding-left: 0;
    margin-left: .5rem;
    font-weight: bold;
    align-self: flex-start;
}

header h1 a {
    color: #000;
    text-decoration: none;
}

header > .container > a {
    font-size: 1.1em;
    text-decoration: none;
    margin-left: auto;
    color: #333;
}

header > .container > a + a {
    margin-left: 0;
}

.repos {
    width: auto;
    list-style: none;
    padding: 0;
    margin: 2.6rem 0;
}

.repos .repo {
    border-bottom: 1px solid #ddd;
    padding: 0 .5rem;
    width: auto;
}

.repos .repo:last-child {
    border-bottom: none;
}

.repo {
    margin: 1.3rem 0;
    width: 100%;
    position: relative;
}

.repo h1 {
    font-size: 1.2em;
}

.repo h1 a {
    text-decoration: none;
}

.repo .buttons {
    display: inline;
    position: absolute;
    left: -100px;
    top: 0;
    z-index: 999;
}

.repo.private a {
    color: hsl(20, 51%, 44%);
}

.clone {
    margin: 1.3rem 0;
}

.clone span {
    font-style: italic;
}

.clone code {
    background: #efefef;
}

hr {
    margin: 1.3rem auto;
    border: none;
    height: 0;
    max-width: 10rem;
    width: 66%;
    border-bottom: 1px solid #ccc;
}

.single > div {
    margin: 2.6rem;
}

figure {
    margin: 2.6rem;
    border: 1px solid #bbb;
}

figure figcaption {
    border-bottom: 1px solid #ddd;
    padding: 0.65rem 0.65rem;
    margin: 0 0.65rem;
}

figure figcaption h3 {
    margin: 0;
    font-size: .75rem;
    letter-spacing: .01rem;
    font-variant: small-caps;
}

figure article {
    padding: 0.65rem 1.3rem;
}

figure article pre.full {
    padding: 0;
    background: transparent;
}


/* flex */
.single {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
}

.single > div {
    flex: 1;
}

.single > aside {
    flex: 2;
}

form {
    margin: 1.3rem 0.5rem;
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

@media screen and (max-width: 60rem) {
    body.single {
        flex-direction: column;
    }

    .single > div {
        margin-bottom: 0;
    }

    aside {
        margin: 0;
    }
}

@media screen and (max-width: 40rem) {
    .single  > div {
        margin: 1.3rem 1.3rem 0 1.3rem;
    }

    aside figure {
        margin: 1.3rem;
    }
}

figure {
    background: white;
    color: black;
}

article {
    font: 13px/1.3em Menlo, monospace;
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
    background: #eee;
}

article pre > code {
    padding: 0;
    background: none;
}

article pre {
    background: #eee;
    padding: .7em 1.3em;
    font-family: Menlo, monospace;
    white-space: pre-wrap;
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
`
