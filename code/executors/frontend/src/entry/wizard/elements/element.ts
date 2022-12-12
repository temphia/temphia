import * as Elem from "../service/wizard_types";

import JsonSingleSelect from "./json/single_select.svelte";
import JsonMultiSelect from "./json/multi_select.svelte";
import JsonSingleNested from "./json/single_nested.svelte";
import JsonMultiNested from "./json/multi_nested.svelte";
import JsonSingleInline from "./json/single_inline.svelte";
import JsonMultiInline from "./json/multi_inline.svelte";
import Image from "./file/image.svelte";
import BasicElement from "./basic/basic.svelte";

const AllElements = {
  [Elem.JSON_MULTI_SELECT]: JsonMultiSelect,
  [Elem.JSON_MULTI_INLINE]: JsonMultiInline,
  [Elem.JSON_MULTI_NESTED]: JsonMultiNested,
  [Elem.JSON_SINGLE_SELECT]: JsonSingleSelect,
  [Elem.JSON_SINGLE_INLINE]: JsonSingleInline,
  [Elem.JSON_SINGLE_NESTED]: JsonSingleNested,
  [Elem.IMAGE]: Image,
};

const NestableElements = {
  [Elem.JSON_MULTI_SELECT]: JsonMultiSelect,
  [Elem.JSON_MULTI_INLINE]: JsonMultiInline,
  [Elem.JSON_SINGLE_SELECT]: JsonSingleSelect,
  [Elem.JSON_SINGLE_INLINE]: JsonSingleInline,
  [Elem.IMAGE]: Image,
};

export { BasicElement, AllElements, NestableElements };
