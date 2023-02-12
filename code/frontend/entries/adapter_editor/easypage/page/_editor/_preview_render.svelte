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
  .marked {
    color: #333;
    font-family: helvetica, arial, sans-serif;
    line-height: 1.5;
    margin: 0 auto;
    max-width: 40em;
    padding: 0 1em;
  }

  :global(.marked h1, .marked h2, .marked h3, .marked h4, .marked h5, .marked h6) {
    margin: 1.25em 0 0.5em 0;
    line-height: 1.2;
    font-weight: bold;
  }

  :global(.marked h1) {
    font-size: larger;
  }

  :global(.marked h2) {
    font-size: large;
  }

  :global(.marked h3) {
    font-size: medium;
  }

  :global(p + p) {
    text-indent: 1.5em;
  }

  :global(.marked li) {
    list-style: disc;
  }

  :global(.marked blockquote) {
    background-color: #dadada;
  }

  :global(.marked pre) {
    background-color: #f3f1f1;
    border-radius: 2%;
  }

  /* ************* */

  :global(.marked a:link) {
    color: #00e;
  }
  :global(.marked a:visited) {
    color: #518;
  }
  :global(.marked a:focus, a:hover) {
    color: #03f;
  }
  :global(a:active) {
    color: #e00;
  }
  :global(
      .marked h1 a:empty:before,
      .marked h2 a:empty:before,
      .marked h3 a:empty:before,
      .marked h4 a:empty:before,
      .marked h5 a:empty:before,
      .marked h6 a:empty:before
    ) {
    content: "#";
  }

  :global(
      .marked h1 a:empty,
      .marked h2 a:empty,
      .marked h3 a:empty,
      .marked h4 a:empty,
      .marked h5 a:empty,
      .marked h6 a:empty
    ) {
    visibility: hidden;
    padding-left: 0.25em;
  }

  :global(
      .marked h1:hover a:empty,
      .marked h2:hover a:empty,
      .marked h3:hover a:empty,
      .marked h4:hover a:empty,
      .marked h5:hover a:empty,
      .marked h6:hover a:empty
    ) {
    visibility: visible;
  }
  :global(.marked img) {
    max-width: 100%;
  }
  :global(.marked figure) {
    margin: 1em 0;
    text-align: center;
  }
  :global(.marked figcaption) {
    font-size: small;
  }

  :global(.marked pre, .marked code, .marked samp, .marked kbd) {
    font-family: monospace, monospace;
  }

  :global(.marked pre, .marked code, .marked kbd) {
    color: #050;
  }

  :global(.marked samp) {
    color: #730;
  }
  :global(.marked pre, .marked blockquote) {
    background: #eee;
    padding: 0.5em;
  }

  :global(.marked pre) {
    overflow: auto;
  }

  :global(.marked blockquote) {
    border-left: medium solid #ccc;
    margin: 1em 0;
  }
  :global(.marked blockquote :first-child) {
    margin-top: 0;
  }
  :global(.marked blockquote :last-child) {
    margin-bottom: 0;
  }
  :global(.marked table) {
    border-collapse: collapse;
  }
  :global(.marked th, .marked td) {
    border: thin solid #999;
    padding: 0.3em 0.4em;
    text-align: left;
  }

  /* ************* */

  @media (prefers-color-scheme: dark) {
    .marked {
      background: #111;
      color: #bbb;
    }
    :global(.marked a:link) {
      color: #9bf;
    }

    :global(.marked a:visited) {
      color: #caf;
    }

    :global(.marked a:focus, .marked a:hover) {
      color: #9cf;
    }
    :global(.marked a:active) {
      color: #f99;
    }
    :global(.marked pre,.marked code,.marked kbd) {
      color: #9c6;
      font-family: monospace, monospace;
    }
    :global(.marked samp) {
      color: #db0;
    }
    :global(.marked pre, .marked blockquote) {
      background: #000;
    }
    :global(.marked blockquote) {
      border-color: #333;
    }
    :global(.marked th, .marked td) {
      border-color: #666;
    }
  }
</style>
