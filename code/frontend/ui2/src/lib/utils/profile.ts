const textColor = "#ffffff";
const backgroundColor = "#1f6feb";
const fontFamily = "Arial";
const fontSize = 40;
const fontWeight = "normal";

const shortName = (name: string) => {
  let n_arr = name.split(" ");
  let n_res = "";
  if (n_arr.length <= 2) {
    for (let i = 0; i < n_arr.length; i++) {
      n_res += n_arr[i][0].toUpperCase();
    }
  } else {
    for (let i = 0; i < 2; i++) {
      n_res += n_arr[i][0].toUpperCase();
    }
  }
  return n_res;
};

export const profileSvg = (name: string) => {
  return `<svg xmlns='http://www.w3.org/2000/svg' xmlns:xlink='http://www.w3.org/1999/xlink' viewBox='0 0 100 100' width='100' height='100' style='font-weight: ${fontWeight};'><rect width='100' height='100' x='0' y='0' fill='${backgroundColor}'></rect><text x='50%' y='50%' alignment-baseline='central' text-anchor='middle' font-family='${fontFamily}' font-size='${fontSize}' fill='${textColor}' dominant-baseline='middle'>${shortName(
    name
  )}</text></svg>`;
};
