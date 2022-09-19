export const CtypeShortText = "shorttext"
export const CtypePhone = "phonenumber"
export const CtypeSelect = "select"
export const CtypeRFormula = "rowformula"
export const CtypeFile = "file"
export const CtypeMultiFile = "multifile"
export const CtypeCheckBox = "checkbox"
export const CtypeCurrency = "currency"
export const CtypeNumber = "number"
export const CtypeLocation = "location"
export const CtypeDateTime = "datetime"

export const CtypeMultSelect = "multiselect"
export const CtypeLongText = "longtext"
export const CtypeSingleUser = "singleuser"
export const CtypeMultiUser = "multiuser"
export const CtypeEmail = "email"
export const CtypeJSON = "json"
export const CtypeRangeNumber = "rangenumber"
export const CtypeColor = "color"


// meta keys
export const KeyPrimary = "__id"
export const KeyVersion = "__version"
export const KeyModSig = "__mod_sig"

export const RefHardPriId = "hard_pri"
export const RefSoftPriId = "soft_pri"
export const RefHardText = "hard_text"
export const RefSoftText = "soft_text"
export const RefHardMulti = "hard_multi"


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
}


export const CtypeFilterConds = {
    [CtypeShortText]: ["equal", "not_equal", "in", "not_in"],
    [CtypePhone]: ["equal", "not_equal", "in", "not_in"],
    [CtypeSelect]: ["equal", "not_equal", "in", "not_in"],
    [CtypeRFormula]: [],
    [CtypeFile]: ["equal", "not_equal"],
    [CtypeMultiFile]: [],
    [CtypeCheckBox]: ["equal", "not_equal"], // is_null
    [CtypeCurrency]: [
        "equal",
        "less_than",
        "not_equal",
        "greater_than",
        "less_than_or_equal",
        "greater_than_or_equal",
    ],
    [CtypeNumber]: [
        "equal",
        "not_equal",
        "less_than",
        "greater_than",
        "less_than_or_equal",
        "greater_than_or_equal",
    ],
    [CtypeLocation]: ["equal", "not_equal"], // around(50m), not_around(50m)
    [CtypeDateTime]: [
        "equal",
        "not_equal",
        "in",
        "not_in",
        "less_than",
        "greater_than",
        "less_than_or_equal",
        "greater_than_or_equal",
    ], // between
};
