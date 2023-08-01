import type { EasypageService } from "../../service/easypage";
import type grapesjs from "grapesjs";
import { extractHtml } from "./core";

export const template_ext =
  (service: EasypageService, initialData: string) =>
  (editor: grapesjs.Editor) => {
    editor.loadProjectData(initialData);

    editor.Panels.addButton("options", {
      id: "the_go_home",
      className: "goHome fa fa-home",
      command: async (editor) => {
        location.hash = "/";
      },
      attributes: { title: "Home" },
      active: true,
    });

    editor.Panels.addButton("options", {
      id: "the_save_button",
      className: "saveButton fa fa-floppy-o",
      command: (editor) => {
        console.log("@save_here");
      },
      attributes: { title: "Save" },
      active: true,
    });
  };
