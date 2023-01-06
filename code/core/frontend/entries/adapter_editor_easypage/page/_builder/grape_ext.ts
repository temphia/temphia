import type { EasypageService } from "../../service/easypage";
import type grapesjs from "grapesjs";

export const easyPageStore =
  (service: EasypageService) => (editor: grapesjs.Editor) => {
    editor.Storage.add("easypage-store", {
      async load(options = {}) {
        const resp = await service.getPageData(options["page_slug"]);
        if (!resp.ok) {
          console.log("Err", resp);
          return {};
        }
        return JSON.parse(resp.data || "{}");
      },

      async store(data, options = {}) {
        data["gen_html"] = extractHtml(editor);

        const resp = await service.setPageData(
          options["page_slug"],
          JSON.stringify(data)
        );
        if (!resp.ok) {
          console.log("Err", resp);
          return;
        }
      },
    });
  };

const extractHtml = (editor: grapesjs.Editor) => {
  return editor.Pages.getAll().map((page) => {
    const component = page.getMainComponent();
    return {
      html: editor.getHtml({ component }),
      css: editor.getCss({ component }),
    };
  });
};
