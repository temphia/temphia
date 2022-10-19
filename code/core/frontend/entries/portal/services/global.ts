declare global {
  interface Window {
    showModal(c: any, p: object): void;
    closeModal(): void;
  }
}

export const showBigModal = (_compo: any, _props: object) => {
  if (!window.showModal) {
    console.warn("BigModal not mounted");
    return;
  }

  window.showModal(_compo, _props);
};

export const closeBigModal = () => {
  if (!window.closeModal) {
    console.warn("BigModal not mounted");
    return;
  }
  
  window.closeModal();
};
