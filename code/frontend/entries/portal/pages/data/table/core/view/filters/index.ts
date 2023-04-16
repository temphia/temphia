import FilterCheckbox from "./_checkbox.svelte";
import FilterDate from "./_date.svelte";
import FilterNumber from "./_number.svelte";
import FilterText from "./_text.svelte";
import FilterDateBi from "./_date_bi.svelte";
import FilterBiNumber from "./_number_bi.svelte";
import FilterMultiText from "./_text_multi.svelte";
import FilterSelect from "./_select.svelte";
import * as f from "../../fields/field";

export {
  FilterCheckbox,
  FilterDate,
  FilterNumber,
  FilterText,
  FilterDateBi,
  FilterBiNumber,
  FilterMultiText,
  FilterSelect,
};



export const CtypeFilterConds = {
  [f.CtypeShortText]: {
    [f.FilterIsNull]: [null],
    [f.FilterIsNotNull]: [null],
    [f.FilterIn]: [FilterMultiText],
    [f.FilterNotIn]: [FilterMultiText],
    [f.FilterEqual]: [FilterText],
    [f.FilterContains]: [FilterText],
    [f.FilterHasPrefix]: [FilterText],
    [f.FilterHasSuffix]: [FilterText],
  },

  [f.CtypeSelect]: {
    [f.FilterIsNull]: [null],
    [f.FilterIsNotNull]: [null],
    [f.FilterEqual]: [FilterSelect],
    [f.FilterNotEqual]: [FilterSelect],
    [f.FilterIn]: [FilterMultiText],
    [f.FilterNotIn]: [FilterMultiText],
  },

  [f.CtypeFile]: {
    [f.FilterIsNull]: [null],
    [f.FilterIsNotNull]: [null],
    [f.FilterEqual]: [FilterText],
    [f.FilterNotEqual]: [FilterText],
    [f.FilterContains]: [FilterText],
  },
  [f.CtypeMultiFile]: {
    [f.FilterIsNull]: [null],
    [f.FilterIsNotNull]: [null],
    [f.FilterContains]: [FilterText],
  },
  [f.CtypeCheckBox]: {
    [f.FilterIsNull]: [null],
    [f.FilterIsNotNull]: [null],
    [f.FilterEqual]: [FilterCheckbox],
    [f.FilterNotEqual]: [FilterCheckbox],
  },

  [f.CtypeCurrency]: {
    [f.FilterIsNull]: [null],
    [f.FilterIsNotNull]: [null],
    [f.FilterEqual]: [FilterNumber],
    [f.FilterLT]: [FilterNumber],
    [f.FilterNotEqual]: [FilterNumber],
    [f.FilterGT]: [FilterNumber],
    [f.FilterLTE]: [FilterNumber],
    [f.FilterGTE]: [FilterNumber],
  },
  [f.CtypeNumber]: {
    [f.FilterIsNull]: [null],
    [f.FilterIsNotNull]: [null],
    [f.FilterIn]: [null],
    [f.FilterNotIn]: [null],
    [f.FilterEqual]: [FilterNumber],
    [f.FilterLT]: [FilterNumber],
    [f.FilterNotEqual]: [FilterNumber],
    [f.FilterGT]: [FilterNumber],
    [f.FilterLTE]: [FilterNumber],
    [f.FilterGTE]: [FilterNumber],
    [f.FilterBetween]: [FilterBiNumber],
  },

  [f.CtypeLocation]: {
    [f.FilterIsNull]: [null],
    [f.FilterIsNotNull]: [null],
    [f.FilterEqual]: [null],
    [f.FilterNotEqual]: [null],
    [f.FilterAround]: [null],
    [f.FilterNotAround]: [null],
  },
  [f.CtypeDateTime]: {
    [f.FilterIsNull]: [null],
    [f.FilterIsNotNull]: [null],
    [f.FilterEqual]: [FilterDate],
    [f.FilterNotEqual]: [FilterDate],
    [f.FilterBetween]: [FilterDateBi],
    [f.FilterBefore]: [FilterDate],
    [f.FilterAfter]: [FilterDate],
  },

  [f.CtypeRFormula]: {},

  // newer types

  [f.CtypeMultSelect]: {
    [f.FilterIsNull]: [null],
    [f.FilterIsNotNull]: [null],
    [f.FilterIn]: [FilterMultiText],
    [f.FilterNotIn]: [FilterMultiText],
    [f.FilterEqual]: [FilterText],
    [f.FilterContains]: [FilterText],
    [f.FilterHasPrefix]: [FilterText],
    [f.FilterHasSuffix]: [FilterText],
  },
  [f.CtypeLongText]: {
    [f.FilterIsNull]: [null],
    [f.FilterIsNotNull]: [null],
    [f.FilterIn]: [FilterMultiText],
    [f.FilterNotIn]: [FilterMultiText],
    [f.FilterEqual]: [FilterText],
    [f.FilterContains]: [FilterText],
    [f.FilterHasPrefix]: [FilterText],
    [f.FilterHasSuffix]: [FilterText],
  },
  [f.CtypeSingleUser]: {
    [f.FilterIsNull]: [null],
    [f.FilterIsNotNull]: [null],
    [f.FilterIn]: [FilterMultiText],
    [f.FilterNotIn]: [FilterMultiText],
    [f.FilterEqual]: [FilterText],
    [f.FilterContains]: [FilterText],
    [f.FilterHasPrefix]: [FilterText],
    [f.FilterHasSuffix]: [FilterText],
  },
  [f.CtypeMultiUser]: {
    [f.FilterIsNull]: [null],
    [f.FilterIsNotNull]: [null],
    [f.FilterIn]: [FilterMultiText],
    [f.FilterNotIn]: [FilterMultiText],
    [f.FilterEqual]: [FilterText],
    [f.FilterContains]: [FilterText],
  },
  [f.CtypeEmail]: {
    [f.FilterIsNull]: [null],
    [f.FilterIsNotNull]: [null],
    [f.FilterIn]: [FilterMultiText],
    [f.FilterNotIn]: [FilterMultiText],
    [f.FilterEqual]: [FilterText],
    [f.FilterContains]: [FilterText],
    [f.FilterHasPrefix]: [FilterText],
    [f.FilterHasSuffix]: [FilterText],
  },
  [f.CtypeJSON]: {
    [f.FilterIsNull]: [null],
    [f.FilterIsNotNull]: [null],
    [f.FilterIn]: [FilterMultiText],
    [f.FilterNotIn]: [FilterMultiText],
    [f.FilterEqual]: [FilterText],
    [f.FilterContains]: [FilterText],
  },
  [f.CtypeRangeNumber]: {
    [f.FilterIsNull]: [null],
    [f.FilterIsNotNull]: [null],
    [f.FilterIn]: [null],
    [f.FilterNotIn]: [null],
    [f.FilterEqual]: [FilterNumber],
    [f.FilterLT]: [FilterNumber],
    [f.FilterNotEqual]: [FilterNumber],
    [f.FilterGT]: [FilterNumber],
    [f.FilterLTE]: [FilterNumber],
    [f.FilterGTE]: [FilterNumber],
    [f.FilterBetween]: [FilterBiNumber],
  },

  [f.CtypeColor]: {
    [f.FilterIsNull]: [null],
    [f.FilterIsNotNull]: [null],
    [f.FilterIn]: [FilterMultiText],
    [f.FilterNotIn]: [FilterMultiText],
    [f.FilterEqual]: [FilterText],
  },
  [f.CtypePhone]: {
    [f.FilterIsNull]: [null],
    [f.FilterIsNotNull]: [null],
    [f.FilterIn]: [FilterMultiText],
    [f.FilterNotIn]: [FilterMultiText],
    [f.FilterEqual]: [FilterText],
    [f.FilterContains]: [FilterText],
    [f.FilterHasPrefix]: [FilterText],
    [f.FilterHasSuffix]: [FilterText],
  },
};