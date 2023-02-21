"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.totalUniqueSlugs = exports.generateSlug = void 0;
const words_1 = require("./words");
// source => https://github.com/nas5w/random-word-slugs
const DEFAULT_NUMBER_OF_WORDS = 3;
function generateSlug(numberOfWords, options) {
    const numWords = numberOfWords || DEFAULT_NUMBER_OF_WORDS;
    const defaultOptions = {
        partsOfSpeech: getDefaultPartsOfSpeech(numWords),
        categories: {},
        format: "kebab",
    };
    const opts = {
        ...defaultOptions,
        ...options,
    };
    const words = [];
    for (let i = 0; i < numWords; i++) {
        const partOfSpeech = opts.partsOfSpeech[i];
        const candidates = words_1.getWordsByCategory(opts.partsOfSpeech[i], opts.categories[partOfSpeech]);
        const rand = candidates[Math.floor(Math.random() * candidates.length)];
        words.push(rand);
    }
    return formatter(words, opts.format);
}
exports.generateSlug = generateSlug;
function getDefaultPartsOfSpeech(length) {
    const partsOfSpeech = [];
    for (let i = 0; i < length - 1; i++) {
        partsOfSpeech.push("adjective");
    }
    partsOfSpeech.push("noun");
    return partsOfSpeech;
}
function formatter(arr, format) {
    if (format === "kebab") {
        return arr.join("-").toLowerCase();
    }
    if (format === "camel") {
        return arr
            .map((el, i) => {
            if (i === 0)
                return el.toLowerCase();
            return el[0].toUpperCase() + el.slice(1).toLowerCase();
        })
            .join("");
    }
    if (format === "lower") {
        return arr.join(" ").toLowerCase();
    }
    if (format === "sentence") {
        return arr
            .map((el, i) => {
            if (i === 0) {
                return el[0].toUpperCase() + el.slice(1).toLowerCase();
            }
            return el;
        })
            .join(" ");
    }
    return arr
        .map((el) => {
        return el[0].toUpperCase() + el.slice(1).toLowerCase();
    })
        .join(" ");
}
function totalUniqueSlugs(numberOfWords, options) {
    const numAdjectives = words_1.getWordsByCategory("adjective", options?.categories?.adjective).length;
    const numNouns = words_1.getWordsByCategory("noun", options?.categories?.noun).length;
    const nums = {
        adjective: numAdjectives,
        noun: numNouns,
    };
    const numWords = numberOfWords || DEFAULT_NUMBER_OF_WORDS;
    const partsOfSpeech = options?.partsOfSpeech || getDefaultPartsOfSpeech(numWords);
    let combos = 1;
    for (let i = 0; i < numWords; i++) {
        combos *= nums[partsOfSpeech[i]];
    }
    return combos;
}
exports.totalUniqueSlugs = totalUniqueSlugs;
