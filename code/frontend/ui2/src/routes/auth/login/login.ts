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
            console.log("@NEXT1")

            const nexttoken = resp.data["next_token"]

            const resp1 = await this.api.login_submit({
                next_token: nexttoken,
            })

            if (!resp1.data["ok"]) {
                return resp1.data
            }
            const resp2 = await this.api.finish({
                preauthed_token: resp1.data["preauthed_token"],
            })

            console.log("@FINISH", resp2.data)

            if (!resp2.ok) {
                return resp2.data
            }

            this.site_utils.setAuthedData({
                tenant_id: resp2.data["tenant_id"],
                user_token: resp2.data["user_token"],
            })

            this.site_utils.gotoPortalPage()

        } else {
            return resp.data
        }
    }
}