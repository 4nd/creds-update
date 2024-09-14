import './app.css'
import 'vite/modulepreload-polyfill'

import Alpine from 'alpinejs'
import axios from 'axios'

window.Alpine = Alpine
Alpine.start()

async function completeCredential(id) {
    return await axios.post('/creds/complete', {id: id})
        .then((response) => response)
        .catch((err) => console.log(err));
}

window.completeCredential = completeCredential