import type { EasypageService } from "../../service/easypage";
import type grapesjs from "grapesjs";

export const easyPageStore =
  (service: EasypageService) => (editor: grapesjs.Editor) => {
    console.log("@grapejs_editor", editor);

    editor.Panels.addButton("options", {
      id: "the_save_button",
      className: "saveButton fa fa-floppy-o",
      command: (editor) => {
        editor.store({});
      },
      attributes: { title: "Save" },
      active: true,
    });

    editor.Panels.addButton("options", {
      id: "the_go_home",
      className: "goHome fa fa-home",
      command: async (editor) => {
        location.hash = "/";
      },
      attributes: { title: "Home" },
      active: true,
    });

    editor.Storage.add("easypage-store", {
      async load(options = {}) {
        const resp = await service.getPageData(options["page_slug"]);
        if (!resp.ok) {
          console.log("Err", resp);
          return {};
        }

        const data = resp.data;

        if (data === `{"pages":[]}`) {
          return;
        }

        try {
          return JSON.parse(data);
        } catch (error) {}
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
