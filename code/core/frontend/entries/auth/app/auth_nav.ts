export class AuthNav {
  nav_options?: any;
  err_message: string;
  constructor() {}

  goto = (target: string, opts?: any) => {
    this.nav_options = opts;
    window.location.hash = target;
  };

  goto_login_page = () => {
    this.goto("/login/");
  };

  goto_login_next_stage = (opts: any) => {
    this.goto("/login/next_stage", opts);
  };

  goto_alt_first_stage = (opts: any) => {
    this.goto("/alt/first_stage", opts);
  };

  goto_alt_second_stage = (opts: any) => {
    this.goto("/alt/second_stage", opts);
  };

  goto_prehook_page = (opts: any) => {
    this.goto("/prehook", opts);
  };

  goto_final_page = () => {
    this.goto("/final");
  };

  goto_error_page = (reason: string, opts?: any) => {
    this.err_message = reason;
    this.goto("/error", opts);
  };
}
