// place files you want to import through the `$lib` alias in this folder.

export class WSService {
  state: number;
  constructor() {
    this.state = 1;
    setInterval(this.handle, 1000);
  }

  handle = () => {
    this.state = this.state + 1;
    console.log("@state", this.state)
  };
}
