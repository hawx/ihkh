package views

const tags = pre + `<div class="regular">
  <h1>Tags</h1>

  <ul>
    {{ range .Tags }}
    <li>
      <a href="/tags/{{.}}">{{.}}</a>
    </li>
    {{ end }}

    <li class="footer">
      <p><a href="http://hawx.me/code/ihkh">ihkh</a> based on <a href="http://ihardlyknowher.com">ihardlyknowher</a></p>
    </li>
  </ul>
</div>` + post
