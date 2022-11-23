
const app = Vue.createApp({
    template: '<h1>Realtt Time Forum {{useriNimi}}<h1>',  // {{ }} sees on viide muutujale mille annam return
                                                        // see template saaks olla ka html-is selle divi sees, mille id on app
    //data on funtsioon ja see tagastab objekti 
    data() {
        return{
            useriNimi: 'Koodija',
        }
    },
})

// here we will mount this app to div we made in html
app.mount('#app')
