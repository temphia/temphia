export const GenerateSRC = ({
  tenant_id,
  did,
  adapter_type,
  adapter_editor_token,
  domain_name,
  base_url,
}) => {
  const opt = JSON.stringify({
    adapter_editor_token,
    base_url,
    adapter_type,
    tenant_id,
    domain_name,
  });

  return `<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <script>
      window["__loader_options__"] = ${opt};
    </script>
    
    <script src="/z/assets/build/adapter_editor_bootloader.js"></script>
    <script src="/z/api/${tenant_id}/v2/adapter_editor/serve/${did}/main.js"></script>
    <link href="/z/api/${tenant_id}/v2/adapter_editor/serve/${did}/main.css" rel="stylesheet" />
  </head>
  <body>
    <div id="adapter-editor-root" style="height: 100vh"></div>
  </body>
  </html>`;
};
