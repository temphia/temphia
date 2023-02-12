import type { LoaderOptions } from "./plug";

interface iframeBuildOptions {
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

export const iframeTemplateBuild = (opts: iframeBuildOptions) => {
  let execscript = "";
  if (opts.exec_loader) {
    execscript = `
            <script src="${opts.api_base_url}engine/plug/${opts.plug}/agent/${opts.agent}/executor/${opts.exec_loader}/loader.js"></script>
            <link href="${opts.api_base_url}engine/plug/${opts.plug}/agent/${opts.agent}/executor/${opts.exec_loader}/loader.css" rel="stylesheet" ></link>
        `;
  }

  return `<!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>${opts.plug}</title>
        <script>window["__loader_options__"] = ${JSON.stringify(derive(opts))}
        </script>
       
        <script>
        ${opts.bootloader}
        </script>
      
        ${execscript}

        <script src="${opts.api_base_url}engine/plug/${opts.plug}/agent/${
    opts.agent
  }/serve/${opts.js_plug_script}"></script>
        <link href="${opts.api_base_url}engine/plug/${opts.plug}/agent/${
    opts.agent
  }/serve/${opts.style_file}" rel="stylesheet" ></link>

    </head>
    <body>
    <div id="plugroot" style="height:100vh;"></div>
    </body>
    </html>`;
};

const derive = (opts: iframeBuildOptions): LoaderOptions => ({
  token: opts.token,
  plug: opts.plug,
  agent: opts.agent,
  api_base_url: opts.api_base_url,
  entry: opts.entry_name,
  exec_loader: opts.exec_loader,
  parent_secret: opts.parent_secret,
  startup_payload: opts.startup_payload,
  tenant_id: opts.tenant_id,
});
