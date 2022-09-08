

// all elements
export const ctypeFileDecode = (files: string): string[] => commaArrayDecode(files)

export const commaArrayDecode = (files: string): string[] => {
    if (!files) {
        return []
    }
    if (!files.includes(',')) {
        return [files]
    }
    return files.split(',')
}

export const commaArryEncode = (files: string[]): string => {
    if (files.length === 1) {
        return files[0]
    }
    return files.join()
}


export {

    CtypeShortText,
    CtypePhone,
    CtypeSelect,
    CtypeRFormula,
    CtypeFile,
    CtypeMultiFile,
    CtypeCheckBox,
    CtypeCurrency,
    CtypeNumber,
    CtypeLocation,
    CtypeDateTime,
    CtypeMultSelect,
    CtypeLongText,
    CtypeSingleUser,
    CtypeMultiUser,
    CtypeEmail,
    CtypeJSON,
    CtypeRangeNumber,
    CtypeColor,

    KeyPrimary,
    KeyVersion,
    KeyModSig,

    RefHardPriId,
    RefSoftPriId,
    RefHardText,
    RefSoftText,
    RefHardMulti,

    CtypeConvertables
} from "../../../../../../lib/core/dyntypes";




