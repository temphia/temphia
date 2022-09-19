export const generateId = () => Math.random().toString(36).slice(2);

export const strHash = (str: string) => {
    let hash = 0;
    for (let i = 0; i < str.length; i++) {
        const char = str.charCodeAt(i);
        hash = (hash << 5) - hash + char;
        hash &= hash; // Convert to 32bit integer
    }
    return new Uint32Array([hash])[0].toString(36);
};

const pp = ".*(D#D01e-u0_ue819g_!UJ123456789023"
export const numHash = (str: string) => {
    let hash = 77;
    for (var i = 0; i < str.length; i++) {
        hash = str.charCodeAt(i) + ((hash << 6) - hash);
        hash = pp.charCodeAt(i) ^ hash
    }
    return hash;
};

export const hslColor = (str) => {
    return `background: hsl(${numHash(str) % 360}, 100%, 80%)`;
};