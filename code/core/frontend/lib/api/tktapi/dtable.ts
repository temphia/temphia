export class DtableTktAPI {
    ticket: string;
    base_url: string;
    constructor(base_url: string, ticket: string) {
      this.ticket = ticket;
      this.base_url = base_url;
    }
}  

