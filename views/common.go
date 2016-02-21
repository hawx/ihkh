package views

const pre = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<!--

    ( )      / __      ___      __      ___   / //
    / /      //   ) ) //   ) ) //  ) ) //   ) / // //   / /
    / /      //   / / //   / / //      //   / / // ((___/ /
    / /      //   / / ((___( ( //      ((___/ / //      / /

    / ___       __      ___                      / __      ___      __
    //\ \     //   ) ) //   ) ) //  / /  / /     //   ) ) //___) ) //  ) )
    //  \ \   //   / / //   / / //  / /  / /     //   / / //       //
    //    \ \ //   / / ((___/ / ((__( (__/ /     //   / / ((____   //

  -->
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
  <head>
    <script type="text/javascript">var _sf_startpt=(new Date()).getTime()</script>
    <title>{{.Title}}</title>
    <meta http-equiv="content-type" content="text/html; charset=utf-8" />
    <meta http-equiv="imagetoolbar" content="no" />
    <meta name="warning" content="HC SVNT DRACONES" />
    <meta name="viewport" content="width=500;initial-scale=1" />
    <style>
      html{color:#000;background:#FFF;}body,div,dl,dt,dd,ul,ol,li,h1,h2,h3,h4,h5,h6,pre,code,form,fieldset,legend,input,button,textarea,p,blockquote,th,td{margin:0;padding:0;}table{border-collapse:collapse;border-spacing:0;}fieldset,img{border:0;}address,caption,cite,code,dfn,em,strong,th,var,optgroup{font-style:inherit;font-weight:inherit;}del,ins{text-decoration:none;}li{list-style:none;}caption,th{text-align:left;}h1,h2,h3,h4,h5,h6{font-size:100%;font-weight:normal;}q:before,q:after{content:'';}abbr,acronym{border:0;font-variant:normal;}sup{vertical-align:baseline;}sub{vertical-align:baseline;}legend{color:#000;}input,button,textarea,select,optgroup,option{font-family:inherit;font-size:inherit;font-style:inherit;font-weight:inherit;}input,button,textarea,select{*font-size:100%;}

      div.regular { font: 12px/20px helvetica,sans-serif; padding: 2em; }
      div.regular a { color: #000; }
      div.regular a:hover { text-decoration: none; }
      div.regular p, div.regular h1 { margin-bottom: 1.5em; }
      div.regular strong { font-weight: bold; }
      div.regular em { font-family: courier; font-style: italic; font-size: 11px; background: #e3e3e3; margin-left: 0.4em; }
      div.regular li.footer { color: #ddd; font-size: 11px; margin-top: 2em; }
      div.regular li.footer a { color: #ddd; }
      div.regular li.footer a:hover { text-decoration: none; }

      div.faq h2 { font-weight: bold; }
      div.faq p { width: 400px; margin-bottom: 2em; }

      div.settings form { margin-bottom: 1.5em; ;}
      div.settings form p { margin-top: 1.5em; margin-bottom: 0; }
      div.settings form input[type='radio'] { margin-left: 2em; }

      div.photos { margin-top: 30px; }
      div.photos ul { display: flex; justify-content: center; flex-direction: column; }
      div.photos li { margin: 0 auto; line-height: 0; padding: 1rem 0; }
      div.photos img { max-width: 95vw; max-height: 95vh; }
      div.photos li.pagers, div.photos li.footer { color: #333; text-align: center; position: relative; font: 14px helvetica,sans-serif; }
      div.photos li.pagers div.left { position: absolute; bottom: 0; left: 0; }
      div.photos li.pagers div.right { position: absolute; bottom: 0; right: 0; }
      div.photos li.pagers a { color: #333; }
      div.photos li.pagers a:hover { text-decoration: none; }
      div.photos li.pagers a.homelink { font-size: 12px; text-transform: uppercase; }
      div.photos li.footer { color: #ddd; font-size: 11px; }
      div.photos li.footer a { color: #ddd; }
      div.photos li.footer a:hover { text-decoration: none; }
    </style>
  </head>
  <body>
    `

const post = `
  </body>
</html>`
