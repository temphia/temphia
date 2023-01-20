<script lang="ts">
  export let code;

  let _code = window["marked"].parse(code, {
    langPrefix: "marked-code",
    headerPrefix: "marked-head",
  });
</script>

<div class="p-2 marked border bg-white h-min">
  {@html _code}
</div>

<style>
  :global(.marked-code) {
    @font-face {
      font-family: "Faster One";
      font-style: normal;
      font-weight: 400;
      src: local("Faster One"), local("FasterOne-Regular"),
        url(http://themes.googleusercontent.com/static/fonts/fasterone/v1/R45a6_a3GGq28ugU8uxGWb3hpw3pgy2gAi-Ip7WPMi0.woff)
          format("woff");
    }

    @font-face {
      font-family: "Oleo Script Swash Caps";
      font-style: normal;
      font-weight: 400;
      src: local("Oleo Script Swash Caps"), local("OleoScriptSwashCaps-Regular"),
        url(http://themes.googleusercontent.com/static/fonts/oleoscriptswashcaps/v1/vdWhGqsBUAP-FF3NOYTe4pakf2VKGDa1aHLWGf__tlnr7w4p9aSvGirXi6XmeXNA.woff)
          format("woff");
    }

    @font-face {
      font-family: "Noticia Text";
      font-style: normal;
      font-weight: 400;
      src: local("Noticia Text"), local("NoticiaText-Regular"),
        url(http://themes.googleusercontent.com/static/fonts/noticiatext/v3/wdyV6x3eKpdeUPQ7BJ5uUKRDOzjiPcYnFooOUGCOsRk.woff)
          format("woff");
    }
  }

  .marked {
    color: black;
    font-size: 1.5em;
    font-family: "Noticia Text", serif;
    margin: 50px auto;
    width: 850px;
  }

  :global(.marked h1, h2, h3, h4) {
    font-weight: bold;
  }

  :global(.marked h1) {
    font-size: larger;
  }
  

  :global(.marked p + p) {
    text-indent: 1.5em;
  }

  :global(.marked blockquote) {
    background-color: #dadada;
  }

  :global(.marked pre) {
    background-color: #f3f1f1;
    border-radius: 2%;
  }

  :global(a) {
    color: #ca2cc3;
  }

  :global(a:hover) {
    color: #880075;
  }

  :global(a:visited) {
    color: #880075;
  }

  :global(img) {
    max-width: 100%;
    height: auto;
  }

  :global(ul) {
    list-style-type: square;
  }

  :global(li > ul) {
    list-style-type: disc;
  }

  :global(li > ul > li > ul) {
    list-style-type: circle;
  }

  :global(a.footnote) {
    font-size: 50%;
    vertical-align: super;
  }

  :global(hr) {
    padding: 0;
    margin: 0;
    border: none;
    text-align: center;
    color: black;
  }

  :global(hr:after) {
    content: "\2724\a0 \2724\a0 \2724";
  }

  :global(.poetry pre code) {
    font-family: Georgia, Garamond, serif;
    font-style: italic;
    font-size: 110% !important;
    line-height: 1.6em;
    display: block;
    margin-left: 1em;
  }

  :global(code, pre) {
    overflow: auto;
  }

  :global(div.footnotes hr) {
    margin: 1em 0;
    width: 75%;
    border-style: dashed;
    border-width: 1px 0 0 0;
    border-color: black;
  }

  :global(div.footnotes hr:after) {
    content: "";
  }

  :global(div.footnotes p) {
    display: inline;
  }
</style>
