// all elements
export const ctypeFileDecode = (files: string): string[] =>
  commaArrayDecode(files);

export const commaArrayDecode = (files: string): string[] => {
  if (!files) {
    return [];
  }
  if (!files.includes(",")) {
    return [files];
  }
  return files.split(",");
};

export const commaArryEncode = (files: string[]): string => {
  if (files.length === 1) {
    return files[0];
  }
  return files.join();
};

export const CtypeShortText = "shorttext";
export const CtypePhone = "phonenumber";
export const CtypeSelect = "select";
export const CtypeRFormula = "rowformula";
export const CtypeFile = "file";
export const CtypeMultiFile = "multifile";
export const CtypeCheckBox = "checkbox";
export const CtypeCurrency = "currency";
export const CtypeNumber = "number";
export const CtypeLocation = "location";
export const CtypeDateTime = "datetime";

export const CtypeMultSelect = "multiselect";
export const CtypeLongText = "longtext";
export const CtypeSingleUser = "singleuser";
export const CtypeMultiUser = "multiuser";
export const CtypeEmail = "email";
export const CtypeJSON = "json";
export const CtypeRangeNumber = "rangenumber";
export const CtypeRatings = "ratings";
export const CtypeColor = "color";

// meta keys
export const KeyPrimary = "__id";
export const KeyVersion = "__version";
export const KeyModSig = "__mod_sig";

export const RefHardPriId = "hard_pri";
export const RefSoftPriId = "soft_pri";
export const RefHardText = "hard_text";
export const RefSoftText = "soft_text";
export const RefHardMulti = "hard_multi";

export const CtypeConvertables = {
  [CtypeShortText]: [CtypeLongText],
  [CtypePhone]: [CtypeShortText],
  [CtypeSelect]: [CtypeShortText, CtypeMultSelect],
  [CtypeRFormula]: [CtypeShortText],
  [CtypeFile]: [CtypeShortText, CtypeMultiFile],
  [CtypeMultiFile]: [CtypeShortText],
  [CtypeCheckBox]: [],
  [CtypeCurrency]: [CtypeNumber],
  [CtypeNumber]: [CtypeCurrency],
  [CtypeLocation]: [],
  [CtypeDateTime]: [],
  [CtypeMultSelect]: [CtypeShortText],
  [CtypeLongText]: [CtypeShortText],
  [CtypeSingleUser]: [CtypeShortText, CtypeMultiUser],
  [CtypeMultiUser]: [CtypeShortText],
  [CtypeEmail]: [CtypeShortText],
  [CtypeJSON]: [CtypeShortText],
  [CtypeRangeNumber]: [],
  [CtypeColor]: [CtypeShortText],
};


export const FilterEqual     = "equal"
export const FilterNotEqual  = "not_equal"
export const FilterIn        = "in"
export const FilterNotIn     = "not_in"
export const FilterNumIn     = "num_in"
export const FilterNumNotIn  = "num_not_in"
export const FilterLT        = "less_than"
export const FilterGT        = "greater_than"
export const FilterLTE       = "less_than_or_equal"
export const FilterGTE       = "greater_than_or_equal"
export const FilterAround    = "around"
export const FilterNotAround = "not_around"
export const FilterLike      = "like"
export const FilterNotLike   = "not_like"
export const FilterILike     = "insensitive_like"
export const FilterNotILike  = "not_insensitive_like"
export const FilterRgex      = "regex"
export const FilterNotRgex   = "not_regex"
export const FilterBefore    = "before"
export const FilterAfter     = "after"
export const FilterContains   = "contains"
export const FilterHasPrefix  = "has_prefix"
export const FilterHasSuffix  = "has_suffix"
export const FilterIsNull     = "is_null"
export const FilterIsNotNull  = "is_not_null"
export const FilterBetween    = "between"
export const FilterNotBetween = "not_between"

export const CtypeIcons = {
  [CtypeShortText]: "annotation",
  [CtypePhone]: "phone",
  [CtypeSelect]: "tag",
  [CtypeRFormula]: "variable",
  [CtypeFile]: "document",
  [CtypeMultiFile]: "document",
  [CtypeCheckBox]: "check-circle",
  [CtypeCurrency]: "currency-dollar",
  [CtypeNumber]: "hashtag",
  [CtypeLocation]: "location-marker",
  [CtypeDateTime]: "calendar",
  [CtypeMultSelect]: "tag",
  [CtypeLongText]: "document-text",
  [CtypeSingleUser]: "user-circle",
  [CtypeMultiUser]: "user-group",
  [CtypeEmail]: "at-symbol",
  [CtypeJSON]: "code",
  [CtypeRangeNumber]: "calculator",
  [CtypeColor]: "color-swatch",
};

export const CtypeShapes = {
  text: [
    CtypeShortText,
    CtypePhone,
    CtypeSelect,
    CtypeFile,
    CtypeLongText,
    CtypeSingleUser,
    CtypeEmail,
    CtypeColor,
  ],
  number: [CtypeCurrency, CtypeNumber, CtypeRangeNumber],
};
