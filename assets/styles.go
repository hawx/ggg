package assets

const styles = `
body {
    font: 14px/1.3 Verdana, Geneva, sans-serif;
    margin: 2.6rem auto 0;
    padding: 0 1.3rem;
    max-width: 40em;
}

body.single {
    max-width: 80em;
}

header {
    margin: 0;
    display: flex;
    justify-content: space-between;
    align-items: baseline;
}

header h1 {
    font-size: 1.5em;
}

header div a {
    margin-left: 1em;
}

.repos {
    width: 100%;
    list-style: none;
    padding: 0;
    margin: 1.3rem 0;
}

.repos .repo {
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

.repo .meta {
    color: #666;
}

.repos .repo .meta.updated {
    text-align: right;
}

.repo .meta.tags {
    list-style: none;
    padding: 0;
    display: flex;
    flex-direction: row;
    justify-content: flex-end;
}

.repo .meta.tags li {
    margin-left: .5rem;
}

.repo .meta.tags li a {
    text-decoration: none;
    color: #A51;
}

.clone {
    margin: 1.3rem 0;
}

footer {
    margin: 1.3rem;
    font-style: italic;
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
    width: 66%;
    border-bottom: 1px solid #ccc;
}

figure {
   margin: 0;
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


/* flex */
.repos .repo {
    display: flex;
    flex-direction: row;
    width: 100%;
    justify-content: space-between;
}

.single {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
}

.single > div {
    flex: 1;
    margin-right: 2.6rem;
}

.single > aside {
    flex: 2;
}

.repo .buttons {
  position: absolute;
  right: -100px;
  top: 14px;
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
        max-width: 40rem;
        flex-direction: column;
    }

    .clone, body.single > div hr {
        display: none;
    }

    aside {
        margin: 2.6rem 0;
    }
}
`
