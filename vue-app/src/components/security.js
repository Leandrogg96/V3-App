import {store} from './store.js'
import router from './../router/index.js'

let Security = {
    // Make sure the user is athenticate
    requireToken: function() {
        if(store.token === '') {
            router.push("/login");
            return false
        }
    },

    // Create request options and send them back
    requestOptions: function(payload) {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");
        headers.append("Authorization", "Bearer " + store.token);

        return {
            method: "POST",
            body: JSON.stringify(payload),
            headers: headers,
        }
    }
}

export default Security;