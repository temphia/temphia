import type grapesjs from "grapesjs";
import grapejs from "grapesjs";
import webpagePlugin from "grapesjs-preset-webpage";
import basicPlugin from "grapesjs-blocks-basic";
import gjsForms from "grapesjs-plugin-forms";
import navPlugin from "grapesjs-navbar";
import customCodePlugin from "grapesjs-custom-code";
import blkFlexboxPlugin from "grapesjs-blocks-flexbox";
import stgrPlugin from "grapesjs-style-gradient";
import styleFilter from "grapesjs-style-filter";
import tabPlugin from "grapesjs-tabs";
import toolTipPlugin from "grapesjs-tooltip";
import twPlugin from "grapesjs-tailwind";

export const editorFactory = (elem, slug, customPlugin) => {
  return grapejs.init({
    container: elem,
    plugins: [
      webpagePlugin,
      basicPlugin,
      gjsForms,
      navPlugin,
      stgrPlugin,
      styleFilter,
      customCodePlugin,
      blkFlexboxPlugin,
      "grapesjs-lory-slider",
      tabPlugin,
      toolTipPlugin,
      twPlugin,
      customPlugin,
    ],

    pluginsOpts: {},
    storageManager: {
      type: "easypage-store",
      stepsBeforeSave: 0,
      options: {
        "easypage-store": {
          page_slug: slug,
        },
      },
    },
  });
};

export const extractHtml = (editor: grapesjs.Editor) => {
  return editor.Pages.getAll().map((page) => {
    const component = page.getMainComponent();
    return {
      html: editor.getHtml({ component }),
      css: editor.getCss({ component }),
    };
  });
};
