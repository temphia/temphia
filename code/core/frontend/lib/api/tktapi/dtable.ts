export class DtableTktAPI {
    ticket: string;
    base_url: string;
    constructor(base_url: string, ticket: string) {
      this.ticket = ticket;
      this.base_url = base_url;
    }

    // async new_row(tid: string, data: any) {
    //   return this.post(`/dtable_ops/${tid}/row`, data);
    // }
    // async get_row(tid: string, rid: number) {
    //   return this.get(`/dtable_ops/${tid}/row/${rid}`);
    // }
    // async update_row(tid: string, rid: number, data: any) {
    //   return this.post(`/dtable_ops/${tid}/row/${rid}`, data);
    // }
    // async delete_row(tid: string, rid: number) {
    //   return this.delete(`/dtable_ops/${tid}/row/${rid}`);
    // }
    // async simple_query(tid: string, data?: any) {
    //   if (!data) {
    //     data = {};
    //   }
    //   return this.post(`/dtable_ops/${tid}/simple_query`, data);
    // }
  
    // async fts_query(tid: string, str: string) {
    //   return this.post(`/dtable_ops/${tid}/fts_query`, {
    //     qstr: str,
    //   });
    // }
  
    // async ref_load(tid: string, data: any) {
    //   return this.post(`/dtable_ops/${tid}/ref_load`, data);
    // }
  
    // async ref_resolve(tid: string, data: any) {
    //   return this.post(`/dtable_ops/${tid}/ref_resolve`, data);
    // }
  
    // async rev_ref_load(tid: string, data) {
    //   return this.post(`/dtable_ops/${tid}/rev_ref_load`, data);
    // }
  
    // async list_activity(tid: string, rowid: number) {
    //   return this.get(`/dtable_ops/${tid}/activity/${rowid}`);
    // }
  
    // async comment_row(tid: string, rowid: number, msg: string) {
    //   return this.post(`/dtable_ops/${tid}/activity/${rowid}`, {
    //     message: msg,
    //   });
    // }
}  


