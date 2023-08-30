import { AuthAPI } from "$lib/services/apiv2/auth";
import { authURL, SiteUtils } from "$lib/utils/site";


export class LoginService {
    api: AuthAPI
    site_utils: SiteUtils
    constructor() {
        this.api = new AuthAPI(authURL())
        this.site_utils = new SiteUtils()
    }

    async init() {
        // this.api.list_methods()
        if (this.site_utils.isLogged()) {
            this.site_utils.gotoPortalPage()
            return
        }
    }

    async loginWithPassword(user_ident, password) {
        const resp = await this.api.login_next({
            user_ident: user_ident,
            password: password,
        })

        if (resp.data["ok"]) {
            this.site_utils.setAuthedData({
                user_token: resp.data["user_token"],
            })

            this.site_utils.gotoPortalPage()
            return
        }


        return resp.data
    }
}