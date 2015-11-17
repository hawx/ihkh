package views

const photostream = pre + `<div class="photos">
  <ul>
    {{ range .Photos }}
    <li class="photo" id="photo_{{.Id}}" style="width: {{.Width}}px;">
      <a href="{{$.UserInfo.PhotosUrl}}{{.Id}}/">
        <img src="{{.Src}}" alt="{{.Id}}" style="width: {{.Width}}px; {{ if .Height }}height: {{.Height}}px{{ end }}" />
      </a>
    </li>
    {{ end }}

    <li class="pagers" style="width: {{.Width}}px;">
      {{ if .PrevPage }}
      <div class="left">
        <a href="{{.PrevPage}}" class="backlink">&larr; previous</a>
      </div>
      {{ end }}

      {{ if .NextPage }}
      <div class="right">
        <a href="{{.NextPage}}">next &rarr;</a>
      </div>
      {{ end }}
    </li>

    <li class="footer">
      <p>all images &copy; <a href="{{.UserInfo.ProfileUrl}}">{{.UserInfo.Name}}</a></p>
      <p><a href="http://hawx.me/code/ihkh">ihkh</a> based on <a href="http://ihardlyknowher.com">ihardlyknowher</a></p>
    </li>
  </ul>
</div>` + post
