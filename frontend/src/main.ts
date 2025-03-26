import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import PrimeVue from 'primevue/config';
import ToastService from 'primevue/toastservice';
import Material from '@primeuix/themes/material';

const app = createApp(App)
app.use(PrimeVue, {
    theme: {
        preset: Material,
        options: {
            prefix: 'p',
            darkModeSelector: 'system',
            cssLayer: {
                name: 'primevue',
                order: 'theme, base, primevue'
            }
        }
    }
})
app.use(ToastService)
app.mount('#app')