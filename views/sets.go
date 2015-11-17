package views

const sets = pre + `<div class="regular">
  <h1>Sets</h1>

  <ul>
    {{ range .Sets }}
    <li>
      <a href="/sets/{{.Id}}">{{.Title}}</a>
    </li>
    {{ end }}

    <li class="footer">
      <p><a href="http://hawx.me/code/ihkh">ihkh</a> based on <a href="http://ihardlyknowher.com">ihardlyknowher</a></p>
    </li>
  </ul>
</div>` + post
