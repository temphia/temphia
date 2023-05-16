package easypage

import (
	"bytes"

	"github.com/tidwall/gjson"
	"github.com/yuin/goldmark"
)

func processPost(path, val string) ([]byte, error) {

	md := (gjson.Get(val, "code").String())

	var buf bytes.Buffer

	buf.WriteString(`
	<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width,initial-scale=1" />
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/spcss@0.9.0">
  </head>	
  <body>
  	<div class="marked">
	`)

	if err := goldmark.Convert([]byte(md), &buf); err != nil {
		return nil, err
	}

	buf.WriteString(`
			</div>
		<body>
	</html>`)

	return buf.Bytes(), nil

}

func processPage(path, val string) ([]byte, error) {

	htmlArray := gjson.Get(val, "gen_html.0").String()

	html := (gjson.Get(htmlArray, "html").String())
	css := (gjson.Get(htmlArray, "css").String())

	var buf bytes.Buffer

	buf.WriteString(`
	<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width,initial-scale=1" />
	<script src="/z/assets/lib/tailwind.js"></script>
	
    <script defer src="fixme.js"></script>
	<link rel="stylesheet" href="fixme.css" />
  </head>	
	`)

	buf.WriteString(html)
	buf.WriteString("<style>")
	buf.WriteString(css)
	buf.WriteString("</style>")
	buf.WriteString("</html>")

	return buf.Bytes(), nil

}
