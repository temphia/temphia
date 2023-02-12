export const FieldShortText = "shorttext";
export const FieldLongText = "longtext";
export const FieldEmail = "email";
export const FieldPhone = "phone";
export const FieldCheckbox = "checkbox";
export const FieldSelect = "select";
export const FieldNumber = "number";
export const FieldRange = "range";
export const FieldColor = "color";


/*

  executor
    plug.default.main
    plug.default.extension1

*/

export const data = {
  items: [
    {
      name: "title",
      type: "shorttext",
      options: [],
      html_attr: {},
      
    },
    {
      name: "info",
      type: "longtext",
      info: "what its about?"
    },
    {
      name: "done",
      type: "checkbox",
    },
  ],
  data: {},
  on_load: "",
  on_submit: "",
};
