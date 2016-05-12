package views

const photostream = pre + `<div class="photos">
  <ul>
    {{ range .Photos }}
    <li class="photo" id="photo_{{.Id}}">
      <a href="{{$.UserInfo.PhotosUrl}}{{.Id}}/">
        <img src="{{.Src}}" alt="{{.Id}}" />
      </a>
    </li>
    {{ end }}

    <li class="pagers" style="width: 500px;">
      {{ if .PrevPage }}
      <div class="left">
        <a id="prev" href="{{.PrevPage}}" class="backlink">&larr; previous</a>
      </div>
      {{ end }}

      {{ if .NextPage }}
      <div class="right">
        <a id="next" href="{{.NextPage}}">next &rarr;</a>
      </div>
      {{ end }}
    </li>

    <li class="footer">
      <p>all images &copy; <a href="{{.UserInfo.ProfileUrl}}">{{.UserInfo.Name}}</a></p>
      <p><a href="http://hawx.me/code/ihkh">ihkh</a> based on <a href="http://ihardlyknowher.com">ihardlyknowher</a></p>
    </li>
  </ul>
</div>` + scripts + post

const scripts = `<script type="text/javascript">
window.onload = function() {
  var idx = 0,
      els = document.getElementsByClassName('photo'),
      len = els.length,
      nxt = false,
      prv = false;

  function showCurrent() {
    els[idx].scrollIntoView(true);
  }

  function nextPage() {
    var next = document.getElementById('next');
    if (next) next.click();
  }

  function prevPage() {
    var prev = document.getElementById('prev')
    if (prev) prev.click();
  }

  function handleKeyPress(e) {
    var ch = String.fromCharCode(e.keyCode || e.charCode);

    switch (ch) {
      case 'j':
        idx++;
        if (idx >= len) {
          idx = len - 1;
          if (nxt) { nextPage(); }
          nxt = true;
          prv = false;
        }
        break;

      case 'k':
        idx--;
        if (idx < 0) {
          idx = 0;
          if (prv) { prevPage(); }
          nxt = false;
          prv = true;
        }
        break;

      default:
        nxt = false;
        prv = false;
    }

    showCurrent();
  }

  document.onkeypress = handleKeyPress;
}
</script>`
