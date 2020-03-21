<html lang="en">

<head>
  <meta charset="utf-8">

  <title>Knowledge</title>
  <meta name="description" content="Aurelio Ogliari's knowledge repository">
  <meta name="author" content="Aurelio">

</head>

<body>
  <h1>Knowledge Stats</h1>
  <h2>List of my topics of interest during the last 30 days</h2>
  <ol>
    {{range .}}
    <li><strong>{{ .Topic }}</strong>: {{ .Count }}</li>
    {{end}}
  </ol>
</body>

</html>
