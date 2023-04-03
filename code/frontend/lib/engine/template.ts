import type { LoaderOptions } from "./plug";

export interface BuilderOptions {
  api_base_url: string;
  entry_name: string;
  plug: string;
  agent: string;
  token: string;
  js_plug_script: string; // client.js
  exec_loader: string; // executor.js
  style_file: string;
  ext_scripts?: object;
  parent_secret: string;
  startup_payload?: any;
  tenant_id: string;
  bootloader: string;
}

export class TemplatedBuilder {
  opts: BuilderOptions;
  buffer: string;

  link_scripts: [boolean, string][];
  link_styles: string[];

  constructor(opts: BuilderOptions) {
    this.opts = opts;

    this.link_scripts = [];
    this.link_styles = [];
  }

  build = () => {
    console.log("@building_using (TemplatedBuilder)")

    this.build_loader_scripts();
    this.build_loader_styles();
    this.build_main_scripts();
    this.build_main_styles()


    this.join();

    return this.buffer
  };

  build_loader_scripts = () => {
    const opts = this.opts;
    const s = `${opts.api_base_url}/engine/plug/${opts.plug}/agent/${opts.agent}/executor/${opts.exec_loader}/loader.js`;
    this.link_scripts.push([false, s]);
  };

  build_loader_styles = () => {
    const opts = this.opts;
    this.link_styles.push(`${opts.api_base_url}/engine/plug/${opts.plug}/agent/${opts.agent}/executor/${opts.exec_loader}/loader.css`)
  };

  build_main_scripts = () => {
    const scripts = this.opts.js_plug_script.split(",").map((script) => {
      const isMod = script.endsWith(".mjs");

      if (script.startsWith("http://") || script.startsWith("https://")) {
        return [isMod, script];
      } else if (script.startsWith("//lib")) {
        return [isMod, `/z/assets/lib${script.replace("//lib", "")}`];
      } else {
        const opts = this.opts;
        const s = `${opts.api_base_url}/engine/plug/${opts.plug}/agent/${opts.agent}/serve/${script}`;
        return [isMod, s];
      }
    });

    this.link_scripts = [...this.link_scripts, ...(scripts as [boolean, string][])];
  };

  build_main_styles = () => {
    const styles = this.opts.style_file.split(",").map((style) => {
      if (style.startsWith("http://") || style.startsWith("https://")) {
        return style;
      } else if (style.startsWith("//lib")) {
        return `/z/assets/lib${style.replace("//lib", "")}`;
      } else {
        const opts = this.opts;
        return `${opts.api_base_url}/engine/plug/${opts.plug}/agent/${opts.agent}/serve/${style}`;
      }
    });

    this.link_styles = [...this.link_styles,...styles];
  };

  join = () => {
    this.buffer = `<!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>${this.opts.plug}</title>
        <script>window["__loader_options__"] = ${this.loader_options()}
        </script>
       
        <script>
        ${this.opts.bootloader}
        </script>`;

    this.link_scripts.forEach((script) => {
      this.buffer =
        this.buffer +
        `<script ${script[0] ? 'type="module"' : ""} src="${script[1]}"> </script>
        `;
    });

    this.link_styles.forEach((style) => {
      this.buffer =
        this.buffer + `<link href="${style}" rel="stylesheet"></link>
        `;
    });

    this.buffer =
      this.buffer +
      ` 
    </head>
    <body>
    <div id="plugroot" style="height:100vh;"></div>
    </body>
    </html>`;
  };

  private loader_options = () => {
    const lopts = {
      token: this.opts.token,
      plug: this.opts.plug,
      agent: this.opts.agent,
      api_base_url: this.opts.api_base_url,
      entry: this.opts.entry_name,
      exec_loader: this.opts.exec_loader,
      parent_secret: this.opts.parent_secret,
      startup_payload: this.opts.startup_payload,
      tenant_id: this.opts.tenant_id,
    } as LoaderOptions;

    return JSON.stringify(lopts);
  };



}
